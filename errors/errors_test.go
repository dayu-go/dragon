package errors

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
)

// go test -v *.go -test.run=TestError
func TestError(t *testing.T) {
	var base *Error
	err := New(http.StatusBadRequest, "reason", "message")
	err2 := New(http.StatusBadRequest, "reason", "message")

	werr := fmt.Errorf("wrap %w", err)

	if errors.Is(err, new(Error)) {
		t.Errorf("should not be equal: %v", err)
	}

	if !errors.Is(werr, err) {
		t.Errorf("should be equal: %v", err)
	}

	if errors.Is(werr, err2) {
		t.Errorf("should be equal: %v", err)
	}

	if !errors.As(err, &base) {
		t.Errorf("shoudl be matchs: %v", err)
	}
}
