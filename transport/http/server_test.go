package http

import (
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
		// defer srv.Stop()

	})

	if err := srv.Start(); err != nil {
		t.Fatal(err)
	}
}
