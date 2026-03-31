package main

import (
	"net/http"

	"backend/handlers"
	"backend/logx"
	"backend/metrics"
)

// StartServer wires basic HTTP routes and starts the server.
func StartServer(addr string) error {
	http.HandleFunc("/api/scan", handlers.PostScan) // POST
	http.HandleFunc("/api/scan/", handlers.GetScanResults) // GET with mux-style handling expected elsewhere
	// Metrics
	http.Handle("/metrics", metrics.MetricsHandler())

	logx.Infof("starting server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

func main() {
	if err := StartServer(":8080"); err != nil {
		logx.Errorf("server failed: %v", err)
	}
}
