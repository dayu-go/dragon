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

func (h *Helper) Debug(a ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelDebug, "msg", fmt.Sprint(a...))
}

func (h *Helper) Debugf(format string, a ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelDebug, "msg", fmt.Sprintf(format, a...))
}

func (h *Helper) Debugw(kv ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelDebug, kv...)
}

func (h *Helper) Error(a ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelDebug, "msg", fmt.Sprint(a...))
}

func (h *Helper) Errorf(format string, a ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelDebug, "msg", fmt.Sprint(a...))
}

func (h *Helper) Errorw(kv ...interface{}) {
	h.Logger.Fields(h.fields).Log(LevelError, kv...)
}

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
