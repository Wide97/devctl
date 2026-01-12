package cmd

import (
	"context"
	"devctl/internal/config"
	"devctl/internal/ping"
	"devctl/pkg/output"
)

func runPing(args []string) error {
	baseURL, err := config.BaseURL()
	if err != nil {
		return err
	}

	if err := ping.Check(context.Background(), baseURL); err != nil {
		output.PrintError(err.Error())
		return err
	}

	output.PrintSuccess("Service Reachable")
	return nil
}
