package log

import "fmt"

type Helper struct {
	Logger
	fields map[string]interface{}
}

func NewHelper(log Logger) *Helper {
	return &Helper{Logger: log}
}

func (h *Helper) Info(a ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelInfo, "msg", fmt.Sprint(a...))
}

func (h *Helper) Infof(format string, a ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelInfo, "msg", fmt.Sprintf(format, a...))
}

func (h *Helper) Infow(kv ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelInfo, kv...)
}

// func (h *Helper) Debug(a ...interface{}) {
// 	if !h.level.Enabled(LevelDebug) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelDebug).Log(fmt.Print(a...))
// }

// func (h *Helper) Debugf(format string, a ...interface{}) {
// 	if !h.level.Enabled(LevelDebug) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelDebug).Log("msg", fmt.Sprintf(format, a...))
// }

// func (h *Helper) Debugw(kv ...interface{}) {
// 	if !h.level.Enabled(LevelDebug) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelDebug).Log(kv...)
// }

// func (h *Helper) Warn(a ...interface{}) {
// 	if !h.level.Enabled(LevelWarn) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelWarn).Log(fmt.Print(a...))
// }

// func (h *Helper) Warnf(format string, a ...interface{}) {
// 	if !h.level.Enabled(LevelWarn) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelWarn).Log("msg", fmt.Sprintf(format, a...))
// }

// func (h *Helper) Warnw(kv ...interface{}) {
// 	if !h.level.Enabled(LevelWarn) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelWarn).Log(kv...)
// }

// func (h *Helper) Error(a ...interface{}) {
// 	if !h.level.Enabled(LevelError) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelError).Log(fmt.Print(a...))
// }

// func (h *Helper) Errorf(format string, a ...interface{}) {
// 	if !h.level.Enabled(LevelError) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelError).Log("msg", fmt.Sprintf(format, a...))
// }

// func (h *Helper) Errorw(kv ...interface{}) {
// 	if !h.level.Enabled(LevelError) {
// 		return
// 	}
// 	With(h.Logger, LevelKey, LevelError).Log(kv...)
// }

func (h *Helper) WithError(err error) *Helper {
	fields := copyFields(h.fields)
	fields["error"] = err
	return &Helper{Logger: h.Logger, fields: fields}
}

func (h *Helper) WithFields(fields map[string]interface{}) *Helper {
	nfields := copyFields(fields)
	for k, v := range h.fields {
		nfields[k] = v
	}
	return &Helper{Logger: h.Logger, fields: nfields}
}
