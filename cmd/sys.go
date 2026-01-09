<<<<<<< HEAD
package cmd

import (
	"context"

	"devctl/internal/sysinfo"
	"devctl/pkg/output"

	"github.com/spf13/cobra"
)

var sysCmd = &cobra.Command{
	Use:   "sys",
	Short: "System related commands",
}

var sysInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show system information of the target service",
	Run: func(cmd *cobra.Command, args []string) {
		info, err := sysinfo.GetInfo(context.Background())
		if err != nil {
			output.PrintError(err)
			return
		}

		output.PrintKeyValue(map[string]string{
			"OS":      info.OS,
			"Arch":    info.Arch,
			"Runtime": info.Runtime,
		})
	},
}

func init() {
	sysCmd.AddCommand(sysInfoCmd)
	rootCmd.AddCommand(sysCmd)
}
=======
package cmd
>>>>>>> origin/main
