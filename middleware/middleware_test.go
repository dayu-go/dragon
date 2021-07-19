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
}

func myHandler(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my handler!")
	return struct{}{}, nil
}

func mymiddleware(s string) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			fmt.Println(s, "pre", request)
			defer fmt.Println(s, "post")
			fmt.Println(s, "end")
			return next(ctx, request)
		}
	}
}

// go test -v *.go -test.run=^TestAAA$
func TestAAA(t *testing.T) {
	f1 := mymiddleware("first")
	f2 := mymiddleware("second")
	f3 := mymiddleware("third")

	res, err := f1(f2(f3(myHandler)))(context.TODO(), "aa")
	t.Logf("res:%+v, err:%v", res, err)

}
