package grpc

import (
	"context"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/transport"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
)

// ServerOption is gRPC server option.
type ServerOption func(o *Server)

// Network with server network.
func Network(network string) ServerOption {
	return func(s *Server) {
		s.network = network
	}
}

// Address with server address.
func Address(addr string) ServerOption {
	return func(s *Server) {
		s.address = addr
	}
}

// Timeout with server timeout.
func Timeout(timeout time.Duration) ServerOption {
	return func(s *Server) {
		s.timeout = timeout
	}
}

// Logger with server logger.
func Logger(logger log.Logger) ServerOption {
	return func(s *Server) {
		s.log = log.NewHelper(logger)
	}
}

// Server is a gRPC server wrapper
type Server struct {
	*grpc.Server
	ctx     context.Context
	lis     net.Listener
	once    sync.Once
	err     error
	network string
	address string
	timeout time.Duration
	log     *log.Helper
	health  *health.Server
}

func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		health:  health.NewServer(),
		log:     log.NewHelper(log.DefaultLogger),
	}
	for _, o := range opts {
		o(srv)
	}
	// var ints = []grpc.UnaryServerInterceptor{}
	// if len(srv.ints) > 0 {
	// 	ints = append(ints, srv.ints...)
	// }

	// grpcOpts := []grpc.ServerOption{
	// 	grpc.UnaryInterceptor(interceptor.Unary()),
	// }
	srv.Server = grpc.NewServer()

	// internal register
	// grpc_health_v1.RegisterHealthServer(srv.Server, srv.health)
	// reflection.Register(srv.Server)
	return srv
}

// Endpoint return a real address to registry endpoint.
// examples:
//   grpc://127.0.0.1:9000?isSecure=false
func (s *Server) Endpoint() (string, error) {
	addr, err := transport.Extract(s.address, s.lis)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%s", addr), nil
}

// Start start the gRPC server.
func (s *Server) Start(ctx context.Context) error {
	s.once.Do(func() {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			s.err = err
			return
		}
		s.lis = lis
	})
	if s.err != nil {
		return s.err
	}
	s.ctx = ctx
	s.log.Infof("[gRPC] server listening on: %s", s.lis.Addr().String())
	s.health.Resume()
	return s.Serve(s.lis)
}

// Stop stop the gRPC server.
func (s *Server) Stop(ctx context.Context) error {
	s.GracefulStop()
	s.health.Shutdown()
	s.log.Info("[gRPC] server stopping")
	return nil
}
