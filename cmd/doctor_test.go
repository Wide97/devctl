package cmd

import (
	"devctl/internal/mockserver"
	"net/http/httptest"
	"os"
	"testing"
)

func TestDoctorMissingEnv(t *testing.T) {
	_ = os.Unsetenv("DEVCTL_BASE_URL")

	if err := Doctor(); err == nil {
		t.Fatalf("expected error when DEVCTL_BASE_URL is not set")
	}
}

func TestDoctorOK(t *testing.T) {
	ts := httptest.NewServer(mockserver.NewMux())
	defer ts.Close()

	if err := os.Setenv("DEVCTL_BASE_URL", ts.URL); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() { _ = os.Unsetenv("DEVCTL_BASE_URL") })

	if err := Doctor(); err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}
