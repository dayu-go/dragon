package log

import (
	"fmt"
	"runtime"
	"sort"
	"strings"
	"time"
)

type defaultLogger struct {
	opts Options
}

func (l *defaultLogger) With(kvs []interface{}) Logger {
	n := (len(kvs) + 1) / 2 // +1 to handle case when len is odd
	m := make(map[string]interface{}, n)
	for i := 0; i < len(kvs); i += 2 {
		k := kvs[i]
		var v interface{} = ErrMissingValue
		if i+1 < len(kvs) {
			v = kvs[i+1]
		}
		Merge(m, k, v)
	}
	l.opts.Fields = m
	return l
}

func (l *defaultLogger) Fields(fields map[string]interface{}) Logger {
	l.opts.Fields = copyFields(fields)
	return l
}

func (l *defaultLogger) Log(level Level, kvs ...interface{}) {
	if !l.opts.Level.Enabled(level) {
		return
	}

	fields := copyFields(l.opts.Fields)

	// level
	fields["level"] = level.String()

	// caller
	if _, file, line, ok := runtime.Caller(l.opts.CallerSkipCount); ok {
		fields["caller"] = fmt.Sprintf("%s:%d", logCallerfilePath(file), line)
	}

	// kvs
	for i := 0; i < len(kvs); i += 2 {
		k := kvs[i]
		var v interface{} = ErrMissingValue
		if i+1 < len(kvs) {
			v = kvs[i+1]
		}
		Merge(fields, k, v)
	}

	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	metadata := ""
	for _, k := range keys {
		metadata += fmt.Sprintf(" %s=%v", k, fields[k])
	}

	t := time.Now().Format(time.RFC3339Nano)
	fmt.Printf("%s %s\n", t, metadata)

}

// logCallerfilePath returns a package/file:line description of the caller,
// preserving only the leaf directory name and file name.
func logCallerfilePath(loggingFilePath string) string {
	// To make sure we trim the path correctly on Windows too, we
	// counter-intuitively need to use '/' and *not* os.PathSeparator here,
	// because the path given originates from Go stdlib, specifically
	// runtime.Caller() which (as of Mar/17) returns forward slashes even on
	// Windows.
	//
	// See https://github.com/golang/go/issues/3335
	// and https://github.com/golang/go/issues/18151
	//
	// for discussion on the issue on Go side.
	idx := strings.LastIndexByte(loggingFilePath, '/')
	if idx == -1 {
		return loggingFilePath
	}
	idx = strings.LastIndexByte(loggingFilePath[:idx], '/')
	if idx == -1 {
		return loggingFilePath
	}
	return loggingFilePath[idx+1:]
}

func copyFields(src map[string]interface{}) map[string]interface{} {
	dst := make(map[string]interface{}, len(src))
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
