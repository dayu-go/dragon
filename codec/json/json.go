package json

import (
	"encoding/json"
	"reflect"

	"github.com/dayu-go/dragon/codec"
)

// Name is the name registered for the json codec.
const Name = "json"

func init() {
	codec.RegisterCodec(Codec{})
}

// codec is a Codec implementation with json.
type Codec struct{}

func (Codec) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (Codec) Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v)
	for rv.Kind() == reflect.Ptr {
		if rv.IsNil() {
			rv.Set(reflect.New(rv.Type().Elem()))
		}
		rv = rv.Elem()
	}
	return json.Unmarshal(data, v)
}

func (Codec) Name() string {
	return Name
}
