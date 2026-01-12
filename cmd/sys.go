package cmd

import (
	"context"
	"devctl/internal/sysinfo"
	"devctl/pkg/output"
	"errors"
	"os"
)

func SysInfo() error {
	if len(os.Args) < 3 {
		return errors.New("usage: devctl sys info")
	}

	if os.Args[1] != "sys" || os.Args[2] != "info" {
		return errors.New("unknown command")
	}

	info, err := sysinfo.GetInfo(context.Background())
	if err != nil {
		output.PrintError(err.Error())
		return err
	}

	output.PrintSuccess("OS: " + info.OS)
	output.PrintSuccess("Arch: " + info.Arch)
	output.PrintSuccess("Runtime: " + info.Runtime)

	return nil
}
