package http

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/transport"
	"github.com/gorilla/mux"
)

// Server is a HTTP server wrapper.
type Server struct {
	*http.Server
	lis     net.Listener
	network string
	address string
	timeout time.Duration
	router  *mux.Router
	log     *log.Helper
}

// NewServer creates an HTTP server by options.
func NewServer(opts ...ServerOption) *Server {
	srv := &Server{
		network: "tcp",
		address: ":0",
		timeout: 1 * time.Second,
		log:     log.NewHelper(log.DefaultLogger),
	}
	for _, o := range opts {
		o(srv)
	}
	srv.router = mux.NewRouter()
	srv.Server = &http.Server{Handler: srv}
	return srv
}

// Handle registers a new route with a matcher for the URL path.
func (s *Server) Handle(path string, h http.Handler) {
	s.router.Handle(path, h)
}

// HandlePrefix registers a new route with a matcher for the URL path prefix.
func (s *Server) HandlePrefix(prefix string, h http.Handler) {
	s.router.PathPrefix(prefix).Handler(h)
}

// HandleFunc registers a new route with a matcher for the URL path.
func (s *Server) HandleFunc(path string, h http.HandlerFunc) {
	s.router.HandleFunc(path, h)
}

// ServeHTTP should write reply headers and data to the ResponseWriter and then return.
func (s *Server) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(res, req)
}

// Endpoint return a real address to registry endpoint.
// examples:
// http://127.0.0.1:8000?isSecure=false
func (s *Server) Endpoint() (string, error) {
	addr, err := transport.Extract(s.address, s.lis)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("http://%s", addr), nil
}

// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	lis, err := net.Listen(s.network, s.address)
	if err != nil {
		return err
	}
	s.lis = lis
	s.log.Infof("[HTTP] server listening on: %s", lis.Addr().String())
	if err := s.Serve(lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("[HTTP] server stopping")
	return s.Shutdown(context.Background())
}
