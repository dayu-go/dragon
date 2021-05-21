package http

import (
	"time"

	"github.com/dayu-go/dragon/log"
)

// ServerOption is HTTP server option.
type ServerOption func(*Server)

// Network with server network
func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// Address with server address
func Addr(address string) ServerOption {
	return func(s *Server) {
		s.address = address
	}
}

// Timeout with server timeout
func Timeout(t time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = t
	}
}

func Logger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(log.DefaultLogger)
	}
}
