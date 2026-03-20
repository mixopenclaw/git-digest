package main

import "testing"

func TestRenderJSON(t *testing.T) {
	s := RenderJSON(map[string]string{"k":"v"})
	if s == "" {
		t.Fatalf("empty JSON")
	}
}
