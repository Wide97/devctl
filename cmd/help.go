package cmd

import "devctl/pkg/output"

func help() error {
	helpText := `Usage:
  devctl help
  devctl doctor
  devctl sys info
  devctl ping
  devctl file exists <path>
  devctl file ls [path]
  devctl calc <add|sub|mul|div> <val1> <val2>

Environment:
  DEVCTL_BASE_URL (can be set via .env)
`
	output.PrintSuccess(helpText)
	return nil
}
