package config

// Decoder is config decoder.
type Decoder func(*KeyValue, map[string]interface{}) error

// Option is config option.
type Option func(*options)

type options struct {
	source  Source
	decoder Decoder
}

// WithSource with config source.
func WithSource(s Source) Option {
	return func(o *options) {
		o.source = s
	}
}

// WithDecoder with config decoder.
func WithDecoder(d Decoder) Option {
	return func(o *options) {
		o.decoder = d
	}
}
