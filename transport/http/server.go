package http

import (
	"context"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/dayu-go/gkit/errors"

	"github.com/dayu-go/gkit/log"
	"github.com/dayu-go/gkit/middleware"
	"github.com/dayu-go/gkit/transport"
	"github.com/gorilla/mux"
)

// Server is a HTTP server wrapper.
type Server struct {
	*http.Server
	lis      net.Listener
	endpoint *url.URL
	once     sync.Once
	network  string
	address  string
	timeout  time.Duration
	ms       []middleware.Middleware
	router   *mux.Router
	log      *log.Helper
	err      error
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

	srv.Server = &http.Server{Handler: srv}
	srv.router = mux.NewRouter()
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
	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()
	if s.timeout > 0 {
		ctx, cancel = context.WithTimeout(ctx, s.timeout)
		defer cancel()
	}
	pathTemplate := req.URL.Path
	if route := mux.CurrentRoute(req); route != nil {
		// /path/123 -> /path/{id}
		pathTemplate, _ = route.GetPathTemplate()
	}

	tr := &Transport{
		endpoint:     s.endpoint.String(),
		operation:    pathTemplate,
		reqHeader:    headerCarrier(req.Header),
		replyHeader:  headerCarrier(res.Header()),
		request:      req,
		pathTemplate: pathTemplate,
	}
	ctx = transport.NewServerContext(ctx, tr)
	// s.router.ServeHTTP(res, req.WithContext(ctx))

	// middleware
	h := func(ctx context.Context, req interface{}) (interface{}, error) {
		s.router.ServeHTTP(res, req.(*http.Request))
		return res, nil
	}
	if len(s.ms) > 0 {
		h = middleware.Chain(s.ms...)(h)
	}
	h(ctx, req)
}

// Endpoint return a real address to registry endpoint.
// examples:
// http://127.0.0.1:8000?isSecure=false
func (s *Server) Endpoint() (*url.URL, error) {
	s.once.Do(func() {
		lis, err := net.Listen(s.network, s.address)
		if err != nil {
			s.err = err
			return
		}
		addr, err := transport.Extract(s.address, lis)
		if err != nil {
			lis.Close()
			s.err = err
			return
		}
		s.lis = lis
		s.endpoint = &url.URL{Scheme: "http", Host: addr}
	})
	if s.err != nil {
		return nil, s.err
	}
	return s.endpoint, nil
}

// Start start the HTTP server.
func (s *Server) Start(ctx context.Context) error {
	if _, err := s.Endpoint(); err != nil {
		return err
	}
	s.BaseContext = func(net.Listener) context.Context {
		return ctx
	}
	s.log.Infof("[HTTP] server listening on: %s", s.lis.Addr().String())
	if err := s.Serve(s.lis); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

// Stop stop the HTTP server.
func (s *Server) Stop(ctx context.Context) error {
	s.log.Info("[HTTP] server stopping")
	return s.Shutdown(context.Background())
}
