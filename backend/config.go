package backend

import (
	"fmt"
	"os"
	"strconv"
)

// Config holds all backend configuration values.
type Config struct {
	Addr           string // listen address, e.g. ":8080"
	JWTSecret      string
	DBPath         string // path to SQLite file or DSN
	StorageRoot    string // local artifact storage directory
	JobConcurrency int    // max concurrent scan jobs
	RetentionDays  int    // artifact retention in days
}

// DefaultConfig returns sane defaults for local development.
func DefaultConfig() Config {
	return Config{
		Addr:           ":8080",
		JWTSecret:      "changeme-dev-secret",
		DBPath:         "data/gitdigest.db",
		StorageRoot:    "data/artifacts",
		JobConcurrency: 4,
		RetentionDays:  30,
	}
}

// LoadConfig builds a Config from environment variables, falling back to defaults.
func LoadConfig() (Config, error) {
	c := DefaultConfig()

	if v := os.Getenv("GD_ADDR"); v != "" {
		c.Addr = v
	}
	if v := os.Getenv("GD_JWT_SECRET"); v != "" {
		c.JWTSecret = v
	}
	if v := os.Getenv("GD_DB_PATH"); v != "" {
		c.DBPath = v
	}
	if v := os.Getenv("GD_STORAGE_ROOT"); v != "" {
		c.StorageRoot = v
	}
	if v := os.Getenv("GD_JOB_CONCURRENCY"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			return c, fmt.Errorf("GD_JOB_CONCURRENCY must be a positive integer")
		}
		c.JobConcurrency = n
	}
	if v := os.Getenv("GD_RETENTION_DAYS"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 1 {
			return c, fmt.Errorf("GD_RETENTION_DAYS must be a positive integer")
		}
		c.RetentionDays = n
	}

	if err := c.Validate(); err != nil {
		return c, err
	}
	return c, nil
}

// Validate checks that required fields are set.
func (c Config) Validate() error {
	if c.JWTSecret == "" {
		return fmt.Errorf("config: JWTSecret is required")
	}
	if c.Addr == "" {
		return fmt.Errorf("config: Addr is required")
	}
	return nil
}
