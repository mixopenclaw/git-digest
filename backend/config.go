package main

// Config holds application-wide configuration values used by background jobs.
type Config struct {
	ArtifactRetentionDays  int
	CleanupIntervalMinutes int
}

// DefaultConfig returns a sensible default configuration.
func DefaultConfig() *Config {
	return &Config{
		ArtifactRetentionDays:  30,
		CleanupIntervalMinutes: 60,
	}
}
