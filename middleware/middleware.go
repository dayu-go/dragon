package middleware

import "context"

type Handler func(ctx context.Context, req interface{}) (res interface{}, err error)

type Middleware func(Handler) Handler

func Chain(outer Middleware, others ...Middleware) Middleware {
	return func(next Handler) Handler {
		for i := range others {
			m := others[len(others)-1-i]
			next = m(next)
		}
		return outer(next)
	}
}
