package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
)

// ResultsResponse represents paginated results for a scan.
type ResultsResponse struct {
	ScanID   string        `json:"scan_id"`
	Results  []interface{} `json:"results"`
	NextPage string        `json:"next_page,omitempty"`
}

// GetScanResults handles GET /api/scan/{id}/results
func GetScanResults(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	scanID := vars["id"]

	// TODO: load results from result_store
	resp := ResultsResponse{
		ScanID:  scanID,
		Results: []interface{}{},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
