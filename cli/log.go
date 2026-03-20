package main

import "fmt"

// Logger is a minimal structured logger for CLI.
type Logger struct{
	Verbose bool
}

func NewLogger(verbose bool) *Logger { return &Logger{Verbose: verbose} }

func (l *Logger) Info(msg string) { fmt.Println("INFO:", msg) }

func (l *Logger) Debug(msg string) { if l != nil && l.Verbose { fmt.Println("DEBUG:", msg) } }
