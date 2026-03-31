package backend

import (
	"net/http"
)

// Server is the main HTTP server entrypoint.
func Server() error {
	mux := http.NewServeMux()
	// TODO: wire routes from router.go
	return http.ListenAndServe(":8080", mux)
}
