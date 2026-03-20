package main

import "fmt"

// ExitError represents an error with an exit code.
type ExitError struct{
	Code int
	Err  error
}

func (e ExitError) Error() string { return e.Err.Error() }

func MapErrorToExitCode(err error) int {
	switch err.(type) {
	case ExitError:
		return err.(ExitError).Code
	default:
		return 1
	}
}

func newErr(msg string) error { return fmt.Errorf(msg) }
