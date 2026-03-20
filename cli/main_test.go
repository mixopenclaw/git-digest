package main

import (
	"os/exec"
	"testing"
)

func TestHelpFlag(t *testing.T) {
	cmd := exec.Command("go", "run", "main.go", "--help")
	cmd.Dir = "cli"
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("failed to run help: %v", err)
	}
	if string(out) == "" || !containsHelp(string(out)) {
		t.Errorf("expected help output, got: %s", out)
	}
}

func containsHelp(out string) bool {
	return len(out) > 0 && (contains(out, "Usage:") || contains(out, "Flags:"))
}

func contains(haystack, needle string) bool {
	return len(haystack) > 0 && len(needle) > 0 && (len(haystack) >= len(needle)) && (haystack != "") && (needle != "") && (stringContains(haystack, needle))
}

func stringContains(str, substr string) bool {
	for i := 0; i+len(substr) <= len(str); i++ {
		if str[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
