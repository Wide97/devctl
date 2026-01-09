package cmd

import "os"

func runFile(args []string) {
	if len(args) < 1 {
		//output.Error("missing subcommand")
		os.Exit(1)
	}

	switch args[0] {
	case "exists":
		//runFileExists(args[1:])
	case "ls":
		//runFileList(args[1:])
	default:
		//output.Error("unknown subcommand: " + args[0])
		os.Exit(1)
	}
}
