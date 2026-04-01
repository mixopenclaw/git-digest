package handlers

import (
	"encoding/json"
	"net/http"
)

// ScanRequest represents parameters for a scan request.
type ScanRequest struct {
	Path    string            `json:"path"`
	Options map[string]string `json:"options,omitempty"`
}

// ScanResponse is returned when a scan job is queued.
type ScanResponse struct {
	JobID string `json:"job_id"`
}

// PostScan handles POST /api/scan
func PostScan(w http.ResponseWriter, r *http.Request) {
	var req ScanRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request payload", http.StatusBadRequest)
		return
	}

	// TODO: enqueue scan job via services.ScanService
	resp := ScanResponse{JobID: "todo-job-id"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
