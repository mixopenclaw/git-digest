package db

import (
	"context"
	"database/sql"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func setupTestDB(t *testing.T) *Store {
	t.Helper()
	database, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	// Apply migration
	_, err = database.Exec(`
		CREATE TABLE IF NOT EXISTS jobs (
			id TEXT PRIMARY KEY,
			status TEXT NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
		CREATE TABLE IF NOT EXISTS results (
			id TEXT PRIMARY KEY,
			job_id TEXT NOT NULL REFERENCES jobs(id),
			payload JSON,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		);
	`)
	if err != nil {
		t.Fatalf("failed to run migration: %v", err)
	}

	return NewStore(database)
}

func TestStorePing(t *testing.T) {
	store := setupTestDB(t)
	if err := store.Ping(context.Background()); err != nil {
		t.Fatalf("ping failed: %v", err)
	}
}

func TestStoreInsertAndGetJob(t *testing.T) {
	store := setupTestDB(t)
	ctx := context.Background()

	if err := store.InsertJob(ctx, "job-1", "pending"); err != nil {
		t.Fatalf("InsertJob failed: %v", err)
	}

	status, err := store.GetJobStatus(ctx, "job-1")
	if err != nil {
		t.Fatalf("GetJobStatus failed: %v", err)
	}
	if status != "pending" {
		t.Fatalf("expected status=pending, got %s", status)
	}
}

func TestStoreGetJobNotFound(t *testing.T) {
	store := setupTestDB(t)
	ctx := context.Background()

	_, err := store.GetJobStatus(ctx, "nonexistent")
	if err == nil {
		t.Fatal("expected error for nonexistent job")
	}
}
