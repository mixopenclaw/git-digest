package main

import (
	"os"
)

// Config holds CLI configuration values.
type Config struct {
	Token string
}

// LoadConfig loads configuration from environment variables or a file path (stubbed).
func LoadConfig(path string) (*Config, error) {
	// For now, prefer env var GITDIGEST_TOKEN
	return &Config{Token: os.Getenv("GITDIGEST_TOKEN")}, nil
}
