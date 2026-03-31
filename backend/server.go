package main

import (
	"log"
	"net/http"

	"backend/handlers"
	"backend/metrics"
)

// StartServer wires basic HTTP routes and starts the server.
func StartServer(addr string) error {
	http.HandleFunc("/api/scan", handlers.PostScan) // POST
	http.HandleFunc("/api/scan/", handlers.GetScanResults) // GET with mux-style handling expected elsewhere
	// Metrics
	http.Handle("/metrics", metrics.MetricsHandler())

	log.Printf("starting server on %s", addr)
	return http.ListenAndServe(addr, nil)
}

func main() {
	if err := StartServer(":8080"); err != nil {
		log.Fatal(err)
	}
}
