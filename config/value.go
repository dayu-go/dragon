package config

import (
	"fmt"
	"reflect"
	"strconv"
	"sync/atomic"
)

type Value interface {
	Bool() (bool, error)
}

type atomicValue struct {
	atomic.Value
}

func (v *atomicValue) Bool() (bool, error) {
	switch val := v.Load().(type) {
	case bool:
		return val, nil
	case int64, float64, string:
		return strconv.ParseBool(fmt.Sprintf("%s", val))
	}
	return false, fmt.Errorf("type assert to %v failed", reflect.TypeOf(v.Load()))
}
