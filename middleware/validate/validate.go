package validate

import (
	"context"

	"github.com/dayu-go/gkit/errors"
	"github.com/dayu-go/gkit/middleware"
)

type validator interface {
	Validate() error
}

// Validator is a validator middleware.
func Validator() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			if v, ok := req.(validator); ok {
				if err := v.Validate(); err != nil {
					return nil, errors.BadRequest("VALIDATOR", err.Error())
				}
			}
			resp, err = handler(ctx, req)
			return
		}
	}
}
