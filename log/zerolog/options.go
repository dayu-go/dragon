package zerolog

import (
	"io"

	"github.com/dayu-go/gkit/log"
)

type Mode uint8

const (
	Production Mode = iota
	Development
)

type Options struct {
	log.Options
	// Runtime mode. (Production by default)
	Mode Mode
	// TimeFormat is one of time.RFC3339, time.RFC3339Nano, time.*
	TimeFormat string
}

// WithOutput set default output writer for the logger
func WithOutput(out io.Writer) Option {
	return func(o *Options) {
		o.Out = out
	}
}

func WithLevel(level log.Level) Option {
	return func(o *Options) {
		o.Level = level
	}
}

func WithProductionMode() Option {
	return func(o *Options) {
		o.Mode = Production
	}
}

func WithDevelopmentMode() Option {
	return func(o *Options) {
		o.Mode = Development
	}
}
