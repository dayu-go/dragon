package yaml

import (
	"github.com/dayu-go/gkit/codec"
	"gopkg.in/yaml.v3"
)

func init() {
	codec.RegisterCodec(code{})
}

// codec is a Codec implementation with json.
type code struct{}

func (code) Marshal(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

func (code) Unmarshal(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func (code) Name() string {
	return "yaml"
}
