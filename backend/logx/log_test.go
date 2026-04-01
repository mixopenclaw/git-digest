package logx

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestInfof(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	Infof("hello %s", "world")

	output := buf.String()
	if !strings.Contains(output, "INFO: hello world") {
		t.Errorf("expected INFO prefix, got: %s", output)
	}
}

func TestErrorf(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	Errorf("something %s", "failed")

	output := buf.String()
	if !strings.Contains(output, "ERROR: something failed") {
		t.Errorf("expected ERROR prefix, got: %s", output)
	}
}

func TestDebugf(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer log.SetOutput(nil)

	Debugf("detail %d", 42)

	output := buf.String()
	if !strings.Contains(output, "DEBUG: detail 42") {
		t.Errorf("expected DEBUG prefix, got: %s", output)
	}
}
