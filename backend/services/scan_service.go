package services

import "context"

// ScanService is a thin abstraction to queue and manage scan jobs.
type ScanService struct{}

// NewScanService returns a new ScanService.
func NewScanService() *ScanService { return &ScanService{} }

// QueueScan enqueues a scan request and returns a job ID (placeholder).
func (s *ScanService) QueueScan(ctx context.Context, path string, opts map[string]string) (string, error) {
	// TODO: implement real queue/storage integration
	return "todo-job-id", nil
}
