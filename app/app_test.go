package app

import (
	"testing"
	"time"

	"github.com/dayu-go/gkit/transport/http"
)

// go test -v *.go -test.run=TestApp
func TestApp(t *testing.T) {
	hs := http.NewServer()
	app := New(
		Name("gkit"),
		Version("v1.0.0"),
		Server(hs),
	)
	time.AfterFunc(time.Second, func() {
		app.Stop()
	})
	if err := app.Run(); err != nil {
		t.Fatal(err)
	}
}
