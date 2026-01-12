package cmd

import (
	"context"
	file "devctl/internal/fileutil"
	"fmt"
	"os"
)

func runFile(args []string) {
	if len(args) < 1 {
		fmt.Println("Error! Missing option.")
		os.Exit(1)
	}

	switch args[0] {
	case "exists":
		runFileExists(args[1:])
	default:
		fmt.Println("Unknown option:", args[0])
		os.Exit(1)
	}
}

func runFileExists(args []string) {
	if len(args) < 1 {
		fmt.Println("Error! Missing path.")
		os.Exit(1)
	}

	path := args[0]
	baseURL := "http://localhost:8080"

	exists, err := file.Exists(context.Background(), baseURL, path)
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}

	if exists {
		fmt.Println("file exists")
	} else {
		fmt.Println("file does not exist")
	}
}
