package handlers

import (
	"encoding/json"
	"net/http"
)

// HealthResponse is a simple readiness/health payload.
type HealthResponse struct {
	Status  string `json:"status"`
	Details map[string]string `json:"details,omitempty"`
}

// HealthCheck handles simple readiness checks (DB/storage checks TODO).
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	resp := HealthResponse{Status: "ok"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
