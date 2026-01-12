package cmd

import (
	"context"
	file "devctl/internal/fileutil"
	"fmt"
	"os"
)

func runFile(args []string) error {
	if len(args) < 1 {
		fmt.Println("Error! Missing option.")
		os.Exit(1)
	}

	switch args[0] {
	case "exists":
		runFileExists(args[1:])
	case "ls":
		runFileLs(args[1:])
	default:
		fmt.Println("Unknown option:", args[0])
		os.Exit(1)
	}

	return nil
}

func runFileExists(args []string) error {
	if len(args) < 1 {
		fmt.Println("Error! Missing path.")
		os.Exit(1)
	}

	path := args[0]
	baseURL := "http://localhost:8080"

	exists, err := file.Exists(context.Background(), baseURL, path)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("file exists")
	} else {
		fmt.Println("file does not exist")
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
		fmt.Println(f)
	}

	return nil
}
