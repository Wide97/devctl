package main

import (
	"devctl/cmd"
	"devctl/internal/mockserver"
	"devctl/pkg/output"
	"net/http/httptest"
	"os"
	"path/filepath"
)

func main() {
	server := httptest.NewServer(mockserver.NewMux())
	defer server.Close()

	if err := os.Setenv("DEVCTL_BASE_URL", server.URL); err != nil {
		output.PrintError("failed to set DEVCTL_BASE_URL: " + err.Error())
		os.Exit(1)
	}

	output.PrintSuccess("Mock server: " + server.URL)

	runCmd([]string{"devctl", "sys", "info"})
	runCmd([]string{"devctl", "ping"})

	tmp := os.TempDir()
	runCmd([]string{"devctl", "file", "ls", tmp})

	path := filepath.Join(tmp, "devctl_demo.txt")
	if err := os.WriteFile(path, []byte("demo"), 0644); err == nil {
		runCmd([]string{"devctl", "file", "exists", path})
	}
}

func runCmd(args []string) {
	os.Args = append([]string{}, args...)
	if err := cmd.Execute(); err != nil {
		output.PrintError(err.Error())
	}
}
