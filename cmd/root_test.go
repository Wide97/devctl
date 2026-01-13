package cmd

import (
	"os"
	"testing"
)

func TestExecuteNoArgs(t *testing.T) {
	withTempCwd(t)
	setArgs([]string{"devctl"})

	err := Execute()
	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}
}

func TestExecuteUnknownCommand(t *testing.T) {
	withTempCwd(t)
	setArgs([]string{"devctl", "nope"})

	err := Execute()
	if err == nil || err.Error() != "unknown command" {
		t.Fatalf("expected unknown command, got %v", err)
	}
}

func TestExecuteSysUsage(t *testing.T) {
	withTempCwd(t)
	setArgs([]string{"devctl", "sys"})

	err := Execute()
	if err == nil || err.Error() != "usage: devctl sys info" {
		t.Fatalf("expected sys usage, got %v", err)
	}
}

func withTempCwd(t *testing.T) {
	t.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("getwd: %v", err)
	}

	tmp := t.TempDir()
	if err := os.Chdir(tmp); err != nil {
		t.Fatalf("chdir: %v", err)
	}

	t.Cleanup(func() {
		_ = os.Chdir(cwd)
	})
}

func setArgs(args []string) {
	os.Args = append([]string{}, args...)
}
