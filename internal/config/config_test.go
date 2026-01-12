package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadDotEnvMissingFile(t *testing.T) {
	err := LoadDotEnv(filepath.Join(t.TempDir(), ".env"))
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestLoadDotEnvSetsValues(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")
	content := "# comment\nexport DEVCTL_BASE_URL=\"http://localhost:8080\"\nFOO=bar\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	if err := LoadDotEnv(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := os.Getenv("DEVCTL_BASE_URL"); got != "http://localhost:8080" {
		t.Fatalf("expected DEVCTL_BASE_URL to be set, got %q", got)
	}
	if got := os.Getenv("FOO"); got != "bar" {
		t.Fatalf("expected FOO to be set, got %q", got)
	}
}

func TestLoadDotEnvDoesNotOverride(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")
	content := "FOO=fromfile\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	if err := os.Setenv("FOO", "existing"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() { _ = os.Unsetenv("FOO") })

	if err := LoadDotEnv(path); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if got := os.Getenv("FOO"); got != "existing" {
		t.Fatalf("expected FOO to remain, got %q", got)
	}
}

func TestLoadDotEnvInvalidLine(t *testing.T) {
	path := filepath.Join(t.TempDir(), ".env")
	content := "NOT_A_PAIR\n"
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		t.Fatalf("write file: %v", err)
	}

	if err := LoadDotEnv(path); err == nil {
		t.Fatalf("expected error for invalid line")
	}
}

func TestBaseURL(t *testing.T) {
	if err := os.Unsetenv("DEVCTL_BASE_URL"); err != nil {
		t.Fatalf("unsetenv: %v", err)
	}

	if _, err := BaseURL(); err == nil {
		t.Fatalf("expected error when DEVCTL_BASE_URL not set")
	}

	if err := os.Setenv("DEVCTL_BASE_URL", "http://example"); err != nil {
		t.Fatalf("setenv: %v", err)
	}
	t.Cleanup(func() { _ = os.Unsetenv("DEVCTL_BASE_URL") })

	got, err := BaseURL()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "http://example" {
		t.Fatalf("expected http://example, got %q", got)
	}
}
