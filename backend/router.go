package backend

import (
	"net/http"
)

// NewRouter returns an HTTP handler with all API routes.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	// TODO: add handlers for /health, /login etc.
	return mux
}
