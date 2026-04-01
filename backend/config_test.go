package main

import "testing"

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	if cfg == nil {
		t.Fatal("DefaultConfig returned nil")
	}
	if cfg.ArtifactRetentionDays != 30 {
		t.Errorf("expected ArtifactRetentionDays=30, got %d", cfg.ArtifactRetentionDays)
	}
	if cfg.CleanupIntervalMinutes != 60 {
		t.Errorf("expected CleanupIntervalMinutes=60, got %d", cfg.CleanupIntervalMinutes)
	}
}
