package services

import "context"

// ResultStore provides persistence for scan results (placeholder).
type ResultStore struct{}

func NewResultStore() *ResultStore { return &ResultStore{} }

// GetResults returns results for a scan id with simple pagination.
func (s *ResultStore) GetResults(ctx context.Context, scanID string, pageToken string) ([]interface{}, string, error) {
	// TODO: implement storage-backed query
	return []interface{}{}, "", nil
}
