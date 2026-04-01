package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

// ResultsResponse represents paginated results for a scan.
type ResultsResponse struct {
	ScanID   string        `json:"scan_id"`
	Results  []interface{} `json:"results"`
	NextPage string        `json:"next_page,omitempty"`
}

// GetScanResults handles GET /api/scan/{id}/results
func GetScanResults(w http.ResponseWriter, r *http.Request) {
	// Extract scan ID from URL path: /api/scan/{id}/results
	// Trim the known prefix and parse the remaining segments.
	path := strings.TrimPrefix(r.URL.Path, "/api/scan/")
	segments := strings.SplitN(path, "/", 2)
	scanID := segments[0]

	// TODO: load results from result_store
	resp := ResultsResponse{
		ScanID:  scanID,
		Results: []interface{}{},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
