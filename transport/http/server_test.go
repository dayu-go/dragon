package http

import (
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"
)

type testData struct {
	Path string `json:"path"`
}

// go test -v *.go -test.run=^TestServer$
func TestServer(t *testing.T) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		data := testData{Path: r.RequestURI}
		json.NewEncoder(w).Encode(data)
	}

	srv := NewServer()
	srv.HandleFunc("/index", fn)

	time.AfterFunc(time.Second, func() {
		defer srv.Stop(context.TODO())

	})

	if err := srv.Start(context.Background()); err != nil {
		t.Fatal(err)
	}
}
