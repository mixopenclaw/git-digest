package main

import "testing"

func TestLogger(t *testing.T) {
	l := NewLogger(true)
	l.Info("info")
	l.Debug("debug")
}
