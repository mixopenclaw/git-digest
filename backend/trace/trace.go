package trace

import "context"

// InitTracer initializes tracing provider (placeholder).
func InitTracer(ctx context.Context) error {
	// TODO: wire OpenTelemetry exporter and tracer provider
	return nil
}

// ShutdownTracer cleans up tracing provider (placeholder).
func ShutdownTracer(ctx context.Context) error {
	// TODO: implement shutdown
	return nil
}
