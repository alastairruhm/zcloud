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
	Use:   "version [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: printVersion,
}

func printVersion(cmd *cobra.Command, args []string) {
	cmd.Printf("zcloud client tool version %s\n", VERSION)
}

func init() {
	RootCmd.AddCommand(VersionCmd)
}
