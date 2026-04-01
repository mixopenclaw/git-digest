package scanner

import (
	"testing"
)

func TestNewScanner(t *testing.T) {
	s := NewScanner()
	if s == nil {
		t.Fatal("NewScanner returned nil")
	}
}

func TestNoopScanReturnsEmptyFindings(t *testing.T) {
	s := NewScanner()
	findings, err := s.Scan("/some/path", map[string]string{"format": "json"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(findings) != 0 {
		t.Fatalf("expected 0 findings, got %d", len(findings))
	}
}

func TestFindingFields(t *testing.T) {
	f := Finding{
		File:    "main.go",
		Message: "TODO: implement feature",
		Meta:    map[string]interface{}{"severity": "low"},
	}
	if f.File != "main.go" {
		t.Errorf("expected File=main.go, got %s", f.File)
	}
	if f.Message != "TODO: implement feature" {
		t.Errorf("unexpected Message: %s", f.Message)
	}
	if f.Meta["severity"] != "low" {
		t.Errorf("unexpected Meta severity: %v", f.Meta["severity"])
	}
}
