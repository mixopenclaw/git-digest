package main

import (
	"context"
	"net/http"
	"time"

	"github.com/mixopenclaw/git-digest/backend/handlers"
	"github.com/mixopenclaw/git-digest/backend/logx"
	"github.com/mixopenclaw/git-digest/backend/metrics"
	"github.com/mixopenclaw/git-digest/backend/trace"
)

// StartServer wires basic HTTP routes and starts the server.
func StartServer(addr string) error {
	http.HandleFunc("/api/scan", handlers.PostScan)        // POST
	http.HandleFunc("/api/scan/", handlers.GetScanResults) // GET with mux-style handling expected elsewhere
	// Metrics
	http.Handle("/metrics", metrics.MetricsHandler())

	logx.Infof("starting server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := trace.InitTracer(ctx); err != nil {
		logx.Errorf("tracer init failed: %v", err)
	}

	if err := StartServer(":8080"); err != nil {
		logx.Errorf("server failed: %v", err)
	}

	if err := trace.ShutdownTracer(ctx); err != nil {
		logx.Errorf("tracer shutdown failed: %v", err)
	}
}
