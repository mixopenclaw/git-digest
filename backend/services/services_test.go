package services

import (
	"context"
	"testing"
)

func TestScanServiceQueueScan(t *testing.T) {
	svc := NewScanService()
	jobID, err := svc.QueueScan(context.Background(), "/repo", map[string]string{"format": "json"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if jobID == "" {
		t.Fatal("expected non-empty job ID")
	}
}

func TestResultStoreGetResults(t *testing.T) {
	store := NewResultStore()
	results, nextPage, err := store.GetResults(context.Background(), "scan-1", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if results == nil {
		t.Fatal("expected non-nil results")
	}
	if nextPage != "" {
		t.Fatalf("expected empty nextPage, got %s", nextPage)
	}
}
