package logging

import (
	"context"
	"fmt"
	"time"

	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/middleware"
	"github.com/dayu-go/gkit/transport"
)

func Server(logger log.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (resp interface{}, err error) {
			var (
				kind      string
				operation string
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				kind = info.Kind().String()
				operation = info.Operation()
			}
			_, err = handler(ctx, req)
			level, stack := extractError(err)
			log.DefaultLogger.Log(level,
				"kind", "server",
				"component", kind,
				"operation", operation,
				"stack", stack,
				"latency", time.Since(startTime).Seconds(),
			)
			return
		}
	}
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}
