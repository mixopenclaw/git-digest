package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	w := httptest.NewRecorder()

	HealthCheck(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	ct := resp.Header.Get("Content-Type")
	if ct != "application/json" {
		t.Fatalf("expected application/json, got %s", ct)
	}

	var body HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if body.Status != "ok" {
		t.Fatalf("expected status=ok, got %s", body.Status)
	}
}

func TestPostScan_ValidRequest(t *testing.T) {
	payload := `{"path":"/repo","options":{"format":"json"}}`
	req := httptest.NewRequest(http.MethodPost, "/api/scan", strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	PostScan(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body ScanResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if body.JobID == "" {
		t.Fatal("expected non-empty job_id")
	}
}

func TestPostScan_InvalidPayload(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/scan", strings.NewReader("not json"))
	w := httptest.NewRecorder()

	PostScan(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", resp.StatusCode)
	}
}

func TestGetScanResults(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/scan/test-id-123/results", nil)
	w := httptest.NewRecorder()

	GetScanResults(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var body ResultsResponse
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if body.ScanID != "test-id-123" {
		t.Fatalf("expected scan_id=test-id-123, got %s", body.ScanID)
	}
	if body.Results == nil {
		t.Fatal("expected non-nil results slice")
	}
}
