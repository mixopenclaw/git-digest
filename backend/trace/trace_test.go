package trace

import (
	"context"
	"testing"
)

func TestInitTracer(t *testing.T) {
	if err := InitTracer(context.Background()); err != nil {
		t.Fatalf("InitTracer failed: %v", err)
	}
}

func TestShutdownTracer(t *testing.T) {
	if err := ShutdownTracer(context.Background()); err != nil {
		t.Fatalf("ShutdownTracer failed: %v", err)
	}
}
