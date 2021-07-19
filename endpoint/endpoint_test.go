package endpoint

import (
	"context"
	"fmt"
	"testing"
)

var (
	ctx = context.Background()
	req = struct{}{}
)

// go test -v *.go -test.run=^TestExampleChain$
func TestExampleChain(t *testing.T) {
	e := Chain(
		annotate("first"),
		annotate("second"),
		annotate("third"),
	)(myEndpoint)
	if _, err := e(ctx, req); err != nil {
		panic(err)
	}

	// first pre
	// second pre
	// third pre
	// my endpoint!
	// third post
	// second post
	// first post
}

func annotate(s string) Middleware {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			return next(ctx, request)
		}
	}
}

func myEndpoint(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my endpoint!")
	return struct{}{}, nil
}
