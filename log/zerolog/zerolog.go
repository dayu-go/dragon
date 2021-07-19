package zerolog

import (
	"os"
	"time"

	"github.com/dayu-go/gkit/log"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
)

type Logger struct {
	zLog zerolog.Logger
	opts Options
}

type Option func(*Options)

func NewLogger(opts ...Option) log.Logger {
	l := &Logger{
		zLog: zerolog.New(os.Stdout),
	}
	for _, o := range opts {
		o(&l.opts)
	}

	// RESET
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	zerolog.CallerSkipFrameCount = 4

	switch l.opts.Mode {
	case Development:
		l.zLog = zerolog.New(os.Stderr).Level(zerolog.DebugLevel).With().Timestamp().Stack().Caller().Logger()
	default: // Production
		l.zLog = zerolog.New(l.opts.Out).Level(zerolog.InfoLevel).With().Timestamp().Stack().Caller().Logger()
	}

	// Set log Level if not default
	if l.opts.Level != 0 {
		zerolog.SetGlobalLevel(loggerToZerologLevel(l.opts.Level))
		l.zLog = l.zLog.Level(loggerToZerologLevel(l.opts.Level))
	}

	// Setting timeFormat
	if len(l.opts.TimeFormat) > 0 {
		zerolog.TimeFieldFormat = l.opts.TimeFormat
	}

	return l
}

func (l *Logger) Log(level log.Level, kvs ...interface{}) {
	if !l.opts.Level.Enabled(level) {
		return
	}

	fields := copyFields(l.opts.Fields)

	// level
	fields["level"] = level.String()

	// caller
	// fields["caller"] = log.Caller(l.opts.CallerSkipCount)

	// kvs
	for i := 0; i < len(kvs); i += 2 {
		k := kvs[i]
		var v interface{} = log.ErrMissingValue
		if i+1 < len(kvs) {
			v = kvs[i+1]
		}
		log.Merge(fields, k, v)
	}
	l.zLog.Log().Fields(fields).Msg("")
}

func (l *Logger) Fields(fields map[string]interface{}) log.Logger {
	l.opts.Fields = copyFields(fields)
	return l
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

func loggerToZerologLevel(level log.Level) zerolog.Level {
	switch level {
	case log.LevelDebug:
		return zerolog.DebugLevel
	case log.LevelInfo:
		return zerolog.InfoLevel
	case log.LevelWarn:
		return zerolog.WarnLevel
	case log.LevelError:
		return zerolog.ErrorLevel
	default:
		return zerolog.InfoLevel
	}
}
