package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"runtime"
	"testing"
)

func TestPingOK(t *testing.T) {
	ts := httptest.NewServer(newMux())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/ping")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestSysInfoOK(t *testing.T) {
	ts := httptest.NewServer(newMux())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/sys/info")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var payload map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&payload); err != nil {
		t.Fatalf("decode error: %v", err)
	}

	if payload["os"] != runtime.GOOS {
		t.Fatalf("expected os %s, got %s", runtime.GOOS, payload["os"])
	}
	if payload["arch"] != runtime.GOARCH {
		t.Fatalf("expected arch %s, got %s", runtime.GOARCH, payload["arch"])
	}
	if payload["runtime"] == "" {
		t.Fatalf("expected runtime to be set")
	}
}

func TestFileExists(t *testing.T) {
	ts := httptest.NewServer(newMux())
	defer ts.Close()

	dir := t.TempDir()
	path := dir + string(os.PathSeparator) + "file.txt"
	if err := os.WriteFile(path, []byte("x"), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	resp, err := http.Get(ts.URL + "/file/exists?path=" + path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}
}

func TestFileMissing(t *testing.T) {
	ts := httptest.NewServer(newMux())
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/file/exists?path=/not-found")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected 404, got %d", resp.StatusCode)
	}
}
