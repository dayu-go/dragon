package middleware

import (
	"context"
	"fmt"
	"testing"
)

// go test -v *.go -test.run=^TestChain$
func TestChain(t *testing.T) {
	e := Chain(
		mymiddleware("first"),
		mymiddleware("second"),
		mymiddleware("third"),
	)(myHandler)
	e(context.TODO(), nil)
	fmt.Println("===========================")
	e(context.TODO(), nil)
}

func myHandler(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my handler!")
	return struct{}{}, nil
}

func middleware1() Middleware {
	return func(next Handler) Handler {
		return nil
	}
}

func mymiddleware(s string) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")
			fmt.Println(s, "end")
			return next(ctx, request)
		}
	}
}
