package cmd

import (
	"context"
	"devctl/internal/config"
	file "devctl/internal/fileutil"
	"devctl/pkg/output"
	"errors"
	"fmt"
)

func runFile(args []string) error {
	if len(args) < 1 {
		return errors.New("Error! Usage file <exists|ls>.")
	}

	switch args[0] {
	case "exists":
		return runFileExists(args[1:])
	case "ls":
		return runFileLs(args[1:])
	default:
		return fmt.Errorf("invalid number: %s", args[0])
	}
}

func runFileExists(args []string) error {
	if len(args) < 1 {
		return errors.New("Error! Missing path.")
	}

	path := args[0]
	baseURL, err := config.BaseURL()
	if err != nil {
		return err
	}

	exists, err := file.Exists(context.Background(), baseURL, path)
	if err != nil {
		return err
	}

	if exists {
		output.PrintSuccess("file exists")
	} else {
		output.PrintSuccess("file does not exist")
	}

	return nil
}

func runFileLs(args []string) error {
	path := "."
	if len(args) >= 1 {
		path = args[0]
	}

	files, err := file.ListFiles(path)
	if err != nil {
		return err
	}

	for _, f := range files {
		output.PrintSuccess(f)
	}

	return nil
}
