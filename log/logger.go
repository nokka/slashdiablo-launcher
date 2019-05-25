package log

import (
	"encoding/json"
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
	Log(keyvals ...interface{}) error
}

type logger struct {
	path       string
	writeMutex sync.Mutex
}

// Log will log the given keyvals.
func (l *logger) Log(keyvals ...interface{}) error {
	data := map[string]interface{}{
		"timestamp": time.Now(),
	}
	written := make(map[string]struct{})

	for i, l := 0, len(keyvals); i < l; i += 2 {
		if k, ok := keyvals[i].(string); ok {
			if _, ok := written[k]; ok {
				continue
			}
			written[k] = struct{}{}

			v := keyvals[i+1]
			switch t := v.(type) {
			case func() interface{}:
				v = t()
			case time.Duration:
				v = t.Seconds()
			case error:
				v = map[string]interface{}{
					"error": t.Error(),
				}
			}
			data[k] = v
		}
	}

	// Indent to readable JSON.
	formatted, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}

	// Write to file on disk.
	if err := l.write(formatted); err != nil {
		return err
	}

	return err
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
