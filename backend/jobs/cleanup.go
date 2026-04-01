package jobs

import (
	"context"
	"github.com/mixopenclaw/git-digest/backend/logx"
	"time"
)

// StartCleanupWorker starts a goroutine that periodically removes old artifacts.
func StartCleanupWorker(ctx context.Context, retentionDays int, interval time.Duration) {
	go func() {
		logx.Infof("cleanup worker started: retention=%dd interval=%s", retentionDays, interval)
		ticker := time.NewTicker(interval)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				logx.Infof("cleanup worker stopping")
				return
			case <-ticker.C:
				// TODO: implement artifact deletion logic
				logx.Infof("cleanup tick: would prune artifacts older than %d days", retentionDays)
			}
		}
	}()
}
