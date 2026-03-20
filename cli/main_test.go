package main

import (
	"os/exec"
	"testing"
)

func TestHelpFlag(t *testing.T) {
	// Run 'go run cli/main.go --help' from the repository root.
	cmd := exec.Command("go", "run", "cli/main.go", "--help")
	// When `go test` runs inside the package folder (cli/), relative paths
	// are resolved from that folder. Set Dir to parent (repo root) so
	// the path 'cli/main.go' is valid.
	cmd.Dir = ".."
	out, err := cmd.CombinedOutput()
	t.Logf("output:\n%s", out)
	if err != nil {
		t.Fatalf("failed to run help: %v, output: %s", err, out)
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
