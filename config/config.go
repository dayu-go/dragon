package config

import (
	// init encoding

	"encoding/json"
	"fmt"

	"github.com/dayu-go/gkit/codec"
	_ "github.com/dayu-go/gkit/codec/json"
	_ "github.com/dayu-go/gkit/codec/yaml"
)

// Config is a config interface.
type Config interface {
	Load() error
	Scan(v interface{}) error
	Source() ([]byte, error)
}

type config struct {
	opts   options
	values map[string]interface{}
}

// New new a config with options.
func New(opts ...Option) Config {
	options := options{
		decoder: func(kv *KeyValue, v map[string]interface{}) error {
			if codec := codec.GetCodec(kv.Format); codec != nil {
				return codec.Unmarshal(kv.Value, &v)
			}
			return fmt.Errorf("unsupported key: %s format: %s", kv.Key, kv.Format)
		},
	}
	for _, o := range opts {
		o(&options)
	}
	return &config{
		opts:   options,
		values: make(map[string]interface{}),
	}

}

func (c *config) Load() error {
	kvs, err := c.opts.source.Load()
	if err != nil {
		return err
	}
	merged, err := cloneMap(c.values)
	if err != nil {
		return err
	}
	for _, kv := range kvs {
		next := make(map[string]interface{})
		if err := c.opts.decoder(kv, next); err != nil {
			return err
		}
		for k, v := range next {
			merged[k] = v
		}
	}
	c.values = merged
	return nil
}

func (r *config) Source() ([]byte, error) {
	return json.Marshal(r.values)
}

func (c *config) Scan(v interface{}) error {
	data, err := c.Source()
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func cloneMap(src map[string]interface{}) (map[string]interface{}, error) {
	data, err := json.Marshal(src)
	if err != nil {
		return nil, err
	}
	dst := make(map[string]interface{})
	if err = json.Unmarshal(data, &dst); err != nil {
		return nil, err
	}
	return dst, nil
}
