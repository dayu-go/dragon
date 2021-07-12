package config

// KeyValue is config key value.
type KeyValue struct {
	Key    string
	Value  []byte
	Format string
}

type Source interface {
	Load() ([]*KeyValue, error)
}
