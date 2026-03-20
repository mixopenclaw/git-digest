package main

import (
	"os"
	"testing"
)

func TestLoadConfig_env(t *testing.T) {
	os.Setenv("GITDIGEST_TOKEN", "test-token")
	cfg, err := LoadConfig("")
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}
	if cfg.Token != "test-token" {
		t.Fatalf("expected token 'test-token', got '%s'", cfg.Token)
	}
}
