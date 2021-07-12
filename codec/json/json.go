package json

import (
	"encoding/json"

	"github.com/dayu-go/gkit/codec"
)

func init() {
	codec.RegisterCodec(code{})
}

type code struct{}

func (code) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (code) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (code) Name() string {
	return "json"
}
