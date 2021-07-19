package log

import (
	"fmt"
	"strings"
)

// Level is a logger level.
type Level int8

// LevelKey is logger level key.
const LevelKey = "level"

const (
	// LevelDebug is logger debug level.
	LevelDebug Level = iota
	// LevelInfo is logger info level.
	LevelInfo
	// LevelWarn is logger warn level.
	LevelWarn
	// LevelError is logger error level.
	LevelError
	// LevelFatal is logger fatal level.
	LevelFatal
)

func (l Level) String() string {
	switch l {
	case LevelDebug:
		return "DEBUG"
	case LevelInfo:
		return "INFO"
	case LevelWarn:
		return "WARN"
	case LevelError:
		return "ERROR"
	case LevelFatal:
		return "FATAL"
	default:
		return ""
	}
}

// ParseLevel parses a level string into a logger Level value.
func ParseLevel(s string) Level {
	switch strings.ToUpper(s) {
	case "DEBUG":
		return LevelDebug
	case "INFO":
		return LevelInfo
	case "WARN":
		return LevelWarn
	case "ERROR":
		return LevelError
	case "FATAL":
		return LevelFatal
	}
	return LevelInfo
}

func (l Level) Enabled(lv Level) bool {
	return lv >= l
}

func Info(a ...interface{}) {
	DefaultLogger.Log(LevelInfo, "msg", fmt.Sprint(a...))
}

func Infof(format string, a ...interface{}) {
	DefaultLogger.Log(LevelInfo, "msg", fmt.Sprintf(format, a...))
}

func Debug(a ...interface{}) {
	DefaultLogger.Log(LevelDebug, "msg", fmt.Sprint(a...))
}

func Debugf(format string, a ...interface{}) {
	DefaultLogger.Log(LevelDebug, "msg", fmt.Sprintf(format, a...))
}

func Warn(a ...interface{}) {
	DefaultLogger.Log(LevelWarn, "msg", fmt.Sprint(a...))
}

func Warnf(format string, a ...interface{}) {
	DefaultLogger.Log(LevelWarn, "msg", fmt.Sprintf(format, a...))
}

func Error(a ...interface{}) {
	DefaultLogger.Log(LevelError, "msg", fmt.Sprint(a...))
}

func Errorf(format string, a ...interface{}) {
	DefaultLogger.Log(LevelError, "msg", fmt.Sprintf(format, a...))
}

func Fatal(a ...interface{}) {
	DefaultLogger.Log(LevelFatal, "msg", fmt.Sprint(a...))
}

func Fatalf(format string, a ...interface{}) {
	DefaultLogger.Log(LevelFatal, "msg", fmt.Sprintf(format, a...))
}
