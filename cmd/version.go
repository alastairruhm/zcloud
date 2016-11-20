package cmd

import (
	"github.com/spf13/cobra"
)

var (
	// VERSION is set during build
	VERSION = "0.0.1"
)

// VersionCmd show version of zcloud tool
var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "show zcloud version",
	Long: `
    `,
	Run: printVersion,
}

func printVersion(cmd *cobra.Command, args []string) {
	cmd.Println("zcloud client tool version ", VERSION)
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
