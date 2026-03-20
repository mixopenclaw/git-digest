package main

import "testing"

func TestMapErrorToExitCode(t *testing.T) {
	if MapErrorToExitCode(newErr("x")) != 1 {
		t.Fatalf("expected default exit code 1")
	}
	if MapErrorToExitCode(ExitError{Code: 42, Err: newErr("y")}) != 42 {
		t.Fatalf("expected exit code 42")
	}
}
