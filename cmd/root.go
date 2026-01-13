package cmd

import (
	"devctl/internal/config"
	"errors"
	"os"
)

func Execute() error {
	if err := config.LoadDefault(); err != nil {
		return err
	}

	if len(os.Args) < 2 {
		return help()
	}

	switch os.Args[1] {
	case "help", "-h", "--help":
		return help()
	case "doctor":
		return Doctor()
	case "sys":
		if len(os.Args) >= 3 && os.Args[2] == "info" {
			return SysInfo()
		}
		return errors.New("usage: devctl sys info")
	case "calc":
		return Calc()
	case "ping":
		return runPing(os.Args[2:])
	case "file":
		return runFile(os.Args[2:])
	default:
		return errors.New("unknown command")
	}
}

//gestito da Wide
//root minimale di gestione error per ora
