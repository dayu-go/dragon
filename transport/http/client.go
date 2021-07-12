package http

import "net/http"

// HTTPClient is an interface that models *http.Client.
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
