package cmd

import (
	"context"
	"devctl/internal/ping"
	"devctl/pkg/output"
	"os"
)

func runPing(args []string) error {
	// Base URL di esempio; in produzione si potrebbe leggere da flag o config
	baseURL := "http://localhost:8080"

	err := ping.Check(context.Background(), baseURL)
	if err != nil {
		output.PrintError(err.Error())
		os.Exit(1)
	}

	output.PrintSuccess("Service Reachable")

	return err
}
