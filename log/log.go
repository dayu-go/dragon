package log

import (
	"errors"
	"os"
)

var (
	// DefaultLogger is default logger.
	DefaultLogger = NewLogger()
)

// ErrMissingValue is appended to keyvals slices with odd length to substitute
// the missing value.
var ErrMissingValue = errors.New("(MISSING)")

// Logger is a logger interface.
type Logger interface {
	// Log writes a log entry
	Log(level Level, kvs ...interface{})
	// Fields set fields to always be logged
	Fields(fields map[string]interface{}) Logger
}

func NewLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:           LevelInfo,
		Fields:          make(map[string]interface{}),
		Out:             os.Stdout,
		CallerSkipCount: 2,
	}
	d := &defaultLogger{opts: options}
	for _, o := range opts {
		o(&d.opts)
	}
	return d
}
