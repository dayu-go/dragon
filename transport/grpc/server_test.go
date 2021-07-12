package grpc

import (
	"context"
	"testing"
	"time"
)

type testKey struct{}

// go test -v *.go -test.run=TestServer
func TestServer(t *testing.T) {
	ctx := context.Background()
	ctx = context.WithValue(ctx, testKey{}, "test")
	srv := NewServer()
	if e, err := srv.Endpoint(); err != nil || len(e) == 0 {
		t.Fatal(e, err)
	}

	go func() {
		// start server
		if err := srv.Start(ctx); err != nil {
			panic(err)
		}
	}()

	time.Sleep(1 * time.Second)
	testClient(t, srv)
	srv.Stop(ctx)

}

func testClient(t *testing.T, srv *Server) {
	u, err := srv.Endpoint()
	if err != nil {
		t.Fatal(err)
	}
	// new a gRPC client
	conn, err := DialInsecure(context.Background(), WithEndpoint(u))
	if err != nil {
		t.Fatal(err)
	}

	conn.Close()
}
