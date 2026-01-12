package output

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestPrintSuccess(t *testing.T) {
	var buf bytes.Buffer
	stdout = &buf
	defer func() { stdout = os.Stdout }()

	PrintSuccess("ok")

	got := strings.TrimSpace(buf.String())
	if got != "ok" {
		t.Fatalf("expected 'ok', got '%s'", got)
	}
}

func TestPrintError(t *testing.T) {
	var buf bytes.Buffer
	stderr = &buf
	defer func() { stderr = os.Stderr }()

	PrintError("error")

	got := strings.TrimSpace(buf.String())
	if got != "error" {
		t.Fatalf("expected 'error', got '%s'", got)
	}
}
