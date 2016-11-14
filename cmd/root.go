package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd the cobra root command
var RootCmd = &cobra.Command{
	Use:   "zcloud CLI tool",
	Short: "Short description",
	Long: `Longer description.. 
            feel free to use a few lines here.
            `,
}

var (
	// Verbose flag set to show detail message
	Verbose bool
	// Host is openstack serivce host
	Host string
	// Username is the name of user in openstack
	Username string
	// Password is the password of user in openstack
	Password string
)

func er(msg interface{}) {
	fmt.Println("Error:", msg)
	os.Exit(-1)
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	RootCmd.PersistentFlags().StringVarP(&Host, "host", "", "", "openstack auth service host")
	RootCmd.MarkPersistentFlagRequired("host")
	RootCmd.PersistentFlags().StringVarP(&Username, "username", "", "", "openstack use's name")
	RootCmd.PersistentFlags().StringVarP(&Password, "password", "", "", "openstack use's name")
	RootCmd.AddCommand(ServerCmd)
}
