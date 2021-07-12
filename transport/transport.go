package transport

import "context"

// Server is transport server
type Server interface {
	Endpoint() (string, error)
	Start(context.Context) error
	Stop(context.Context) error
}

type Transport struct {
	Kind Kind
}

// Kind defines the type of Transport
type Kind string

// Defines a set of transport kind
const (
	KindGRPC Kind = "grpc"
	KindHTTP Kind = "http"
)

type transportKey struct{}

// NewContext returns a new Context that carries value
func NewContext(ctx context.Context, tr Transport) context.Context {
	return context.WithValue(ctx, transportKey{}, tr)
}

// FromContext returns the Transport value stored in ctx, if any.
func FromContext(ctx context.Context) (ok bool, tr Transport) {
	tr, ok = ctx.Value(transportKey{}).(Transport)
	return
}
