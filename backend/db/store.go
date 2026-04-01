package db

import (
	"context"
	"database/sql"
)

// Store is a thin wrapper around *sql.DB for job/result operations.
type Store struct{
	DB *sql.DB
}

func NewStore(db *sql.DB) *Store { return &Store{DB: db} }

// Ping verifies DB connectivity.
func (s *Store) Ping(ctx context.Context) error {
	return s.DB.PingContext(ctx)
}

// InsertJob inserts a job record (placeholder implementation).
func (s *Store) InsertJob(ctx context.Context, id string, status string) error {
	_, err := s.DB.ExecContext(ctx, "INSERT INTO jobs(id, status) VALUES(?, ?)", id, status)
	return err
}

// GetJobStatus returns the status for a job id (placeholder).
func (s *Store) GetJobStatus(ctx context.Context, id string) (string, error) {
	var status string
	if err := s.DB.QueryRowContext(ctx, "SELECT status FROM jobs WHERE id = ?", id).Scan(&status); err != nil {
		return "", err
	}
	return status, nil
}
