package log

import (
	"fmt"
	"os"
	"sync"
	"time"
)

const (
	errorLog = "errors.log"
)

// Logger represents the logger interface while hiding the implementation.
type Logger interface {
	// Info is used to log any kind of information.
	Info(string) error

	// Debug is used by the debugger when enabled.
	Debug(string) error

	// Error is used to log errors.
	Error(error) error
}

type logger struct {
	path       string
	writeMutex sync.Mutex
}

func (l *logger) Info(entry string) error {
	msg := fmt.Sprintf("[INFO] %v: %s\n", time.Now().Format(time.RFC3339), entry)
	return l.write([]byte(msg))
}

func (l *logger) Debug(entry string) error {
	msg := fmt.Sprintf("[DEBUG] %v: %s\n", time.Now().Format(time.RFC3339), entry)
	return l.write([]byte(msg))
}

func (l *logger) Error(err error) error {
	msg := fmt.Sprintf("[ERROR] %v: %s\n", time.Now().Format(time.RFC3339), err.Error())
	return l.write([]byte(msg))
}

func (l *logger) write(logEntry []byte) error {
	// Lock access to the file.
	l.writeMutex.Lock()

	// Unlock it when we're done writing.
	defer l.writeMutex.Unlock()

	// Check if the log file exists already.
	exists, err := l.logExists()
	if err != nil {
		return err
	}

	// File pointer, we'll use it regardless if we create the file
	// or if we open it in append mode to append the log entry.
	var f *os.File

	// Log file didn't exist, creating one.
	if !exists {
		f, err = os.Create(fmt.Sprintf("%s/%s", l.path, errorLog))
		if err != nil {
			return err
		}
	} else {
		// Open the file in append mode.
		f, err = os.OpenFile(
			fmt.Sprintf("%s/%s", l.path, errorLog),
			os.O_APPEND|os.O_CREATE|os.O_WRONLY,
			0644,
		)
		if err != nil {
			return err
		}
	}

	// Close the file when we're done.
	defer f.Close()

	if _, err := f.Write(logEntry); err != nil {
		return err
	}

	return nil
}

func (l *logger) logExists() (bool, error) {
	_, err := os.Stat(fmt.Sprintf("%s/%s", l.path, errorLog))
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		// Unknown error.
		return false, err

	}

	return true, nil
}

// NewLogger returns a new logger with all dependencies set up.
func NewLogger(path string) Logger {
	return &logger{
		path: path,
	}
}
