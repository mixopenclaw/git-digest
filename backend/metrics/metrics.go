package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	ScansTotal = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "app_scans_total",
		Help: "Total number of scan jobs queued",
	})
	ScanErrors = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "app_scan_errors_total",
		Help: "Total number of scan errors",
	})
)

func init() {
	prometheus.MustRegister(ScansTotal, ScanErrors)
}

// MetricsHandler returns the prometheus metrics handler (promhttp).
func MetricsHandler() http.Handler {
	return promhttp.Handler()
}
