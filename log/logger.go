package log

import (
	"encoding/json"
	"fmt"
	"time"
)

// Logger represents the logger interface while hiding the implementation.
type Logger interface {
	Log(keyvals ...interface{}) error
}

type logger struct {
	path string
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
	o, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(string(o))
	return err
}

// NewLogger returns a new logger with all dependencies set up.
func NewLogger(path string) Logger {
	return &logger{
		path: path,
	}
}
