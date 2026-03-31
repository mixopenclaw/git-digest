package handlers

import (
	"net/http"
)

// Login issues a fake JWT to match API scaffold.
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"token": "stub.jwt.token"}`))
}
