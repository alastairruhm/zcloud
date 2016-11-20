package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// RootCmd the cobra root command
var RootCmd = &cobra.Command{
	Use:   "zcloud",
	Short: "a CLI app for openstack client",
	Long: `zcloud is a CLI app for openstack client.
This application is created for speeding up the daily operation on cloud servers on openstack
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
	RootCmd.PersistentFlags().StringVarP(&Username, "username", "", "", "openstack user's name")
	RootCmd.PersistentFlags().StringVarP(&Password, "password", "", "", "openstack user's name")
	RootCmd.AddCommand(ServerCmd)
}
