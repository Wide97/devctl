package sysinfo

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetInfoSuccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"os": "linux",
			"arch": "amd64",
			"runtime": "go1.22"
		}`))
	}))
	defer server.Close()

	os.Setenv("DEVCTL_BASE_URL", server.URL)
	defer os.Unsetenv("DEVCTL_BASE_URL")

	info, err := GetInfo(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if info.OS != "linux" {
		t.Errorf("expected linux, got %s", info.OS)
	}
}

func TestGetInfoMissingFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	os.Setenv("DEVCTL_BASE_URL", server.URL)
	defer os.Unsetenv("DEVCTL_BASE_URL")

	_, err := GetInfo(context.Background())
	if err == nil {
		t.Fatal("expected error, got nil")
	}
}
