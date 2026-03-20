package main

import "testing"

func TestScanCommand_stub(t *testing.T) {
	if err := scanCommand("./", "json"); err != nil {
		t.Fatalf("scanCommand returned error: %v", err)
	}
}
