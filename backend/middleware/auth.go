package middleware

import (
	"net/http"
)

// JWTAuth is a stub middleware.
func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: parse and verify JWT
		next.ServeHTTP(w, r)
	})
}
