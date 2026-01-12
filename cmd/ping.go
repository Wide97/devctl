package cmd

import (
	"context"
	"devctl/internal/ping"
	"devctl/pkg/output"
	"os"
)

func runPing(args []string) {
	// Base URL di esempio; in produzione si potrebbe leggere da flag o config
	baseURL := "http://localhost:8080"

	err := ping.Check(context.Background(), baseURL)
	if err != nil {
		output.Error(err.Error())
		os.Exit(1)
	}

	output.OK("service reachable")
}
