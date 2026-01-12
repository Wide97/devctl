package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"

	"devctl/internal/sysinfo"
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
		return err
	}

	fmt.Printf("OS: %s\n", info.OS)
	fmt.Printf("Arch: %s\n", info.Arch)
	fmt.Printf("Runtime: %s\n", info.Runtime)

	return nil
}
