package cmd

import (
	"context"
	"devctl/internal/config"
	"devctl/internal/ping"
	"devctl/internal/sysinfo"
	"devctl/pkg/output"
	"errors"
)

func Doctor() error {
	baseURL, err := config.BaseURL()
	if err != nil {
		output.PrintError(err.Error())
		output.PrintSuccess("Tip: set DEVCTL_BASE_URL or create a .env file")
		return err
	}

	output.PrintSuccess("DEVCTL_BASE_URL=" + baseURL)

	var failed bool

	if err := ping.Check(context.Background(), baseURL); err != nil {
		output.PrintError("ping failed: " + err.Error())
		failed = true
	} else {
		output.PrintSuccess("ping OK")
	}

	info, err := sysinfo.GetInfo(context.Background())
	if err != nil {
		output.PrintError("sys info failed: " + err.Error())
		failed = true
	} else {
		output.PrintSuccess("sys info OK: " + info.OS + " " + info.Arch + " " + info.Runtime)
	}

	if failed {
		return errors.New("doctor failed")
	}
	return nil
}
