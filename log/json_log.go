package log

import (
	"bytes"
	"encoding"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
)

type jsonLogger struct {
	opts Options
}

// NewjsonLoggerger returns a Logger that encodes keyvals to the Writer as a
// single JSON object. Each log event produces no more than one call to
// w.Write. The passed Writer must be safe for concurrent use by multiple
// goroutines if the returned Logger will be used concurrently.
func NewJsonLogger(opts ...Option) Logger {
	// Default options
	options := Options{
		Level:           LevelInfo,
		Fields:          make(map[string]interface{}),
		Out:             &bytes.Buffer{},
		CallerSkipCount: 2,
	}
	d := &jsonLogger{opts: options}
	for _, o := range opts {
		o(&d.opts)
	}
	return d
}

func (l *jsonLogger) Fields(fields map[string]interface{}) Logger {
	l.opts.Fields = copyFields(fields)
	return l
}

func (l *jsonLogger) Log(level Level, kvs ...interface{}) {

	if !l.opts.Level.Enabled(level) {
		return
	}

	m := copyFields(l.opts.Fields)

	// level
	m["level"] = level.String()

	// caller
	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		m["caller"] = fmt.Sprintf("%s:%d", logCallerfilePath(file), line)
	}

	// kvs
	for i := 0; i < len(kvs); i += 2 {
		k := kvs[i]
		var v interface{} = ErrMissingValue
		if i+1 < len(kvs) {
			v = kvs[i+1]
		}
		Merge(m, k, v)
	}

	enc := json.NewEncoder(l.opts.Out)
	enc.SetEscapeHTML(false)
	enc.Encode(m)
}

func Merge(dst map[string]interface{}, k, v interface{}) {
	var key string
	switch x := k.(type) {
	case string:
		key = x
	case fmt.Stringer:
		key = safeString(x)
	default:
		key = fmt.Sprint(x)
	}

	// We want json.Marshaler and encoding.TextMarshaller to take priority over
	// err.Error() and v.String(). But json.Marshall (called later) does that by
	// default so we force a no-op if it's one of those 2 case.
	switch x := v.(type) {
	case json.Marshaler:
	case encoding.TextMarshaler:
	case error:
		v = safeError(x)
	case fmt.Stringer:
		v = safeString(x)

	}
	dst[key] = v
}

// TODO: 函数定义命名返回值 和 不命名返回值对defer里面对返回值赋值的影响
func safeString(str fmt.Stringer) (s string) {
	defer func() {
		if err := recover(); err != nil {
			if v := reflect.ValueOf(str); v.Kind() == reflect.Ptr && v.IsNil() {
				s = "NULL"
			} else {
				panic(err)
			}
		}
	}()
	s = str.String()
	return
}

func safeError(err error) interface{} {
	var e interface{}
	defer func() {
		if err := recover(); err != nil {
			if v := reflect.ValueOf(e); v.Kind() == reflect.Ptr && v.IsNil() {
				e = nil
			} else {
				panic(err)
			}
		}
	}()
	e = err.Error()
	return e
}
