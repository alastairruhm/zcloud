package cmd

import (
	"github.com/alastairruhm/zcloud/cloud"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// ServerCmd sub-command of zcloud about server
var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "operation about server",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
}

// ServerListCmd list server in specific project(tenant)
var ServerListCmd = &cobra.Command{
	Use:   "list [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: serverList,
}

func serverList(cmd *cobra.Command, args []string) {
	openstackClient, err := cloud.NewClient(Host, Username, Password, Project)
	if err != nil {
		errOutput(cmd, err)
	}
	serverList, err := openstackClient.ServerList(servers.ListOpts{})
	if err != nil {
		errOutput(cmd, err)
	}
	cmd.Printf("%32s\t%30s\t%s\t%s\t\n", "ID", "Name", "Status", "Networks")
	cmd.Printf("---------\n")
	for _, s := range serverList {
		networks := cloud.GetServerNetworkAddr(&s)
		network := strings.Join(networks, "|")
		cmd.Printf("%32s\t%30s\t%s\t%v\t\n", s.ID, s.Name, s.Status, network)
	}
}

// ServerCreateCmd create and boot server with given options
var ServerCreateCmd = &cobra.Command{
	Use:   "create [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: serverCreate,
}

func serverCreate(cmd *cobra.Command, args []string) {
	openstackClient, err := cloud.NewClient(Host, Username, Password, Project)
	if err != nil {
		errOutput(cmd, err)
	}
	networkID, err := openstackClient.GetNetworkIDFromName(Network)
	if err != nil {
		errOutput(cmd, err)
	}
	networks := []servers.Network{servers.Network{UUID: networkID}}
	opts := servers.CreateOpts{
		Name:          Name,
		FlavorName:    Flavor,
		ImageName:     Image,
		Networks:      networks[:],
		ServiceClient: openstackClient.ComputeService,
	}
	if err := openstackClient.ServerCreate(opts); err != nil {
		cmd.Println("server create action error:", err)
		os.Exit(-1)
	}
	cmd.Printf("server %s is created\n", Name)
}

// ServerDeleteCmd delete the specific server
var ServerDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a specific server",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: serverDelete,
}

func serverDelete(cmd *cobra.Command, args []string) {
	openstackClient, err := cloud.NewClient(Host, Username, Password, Project)
	if err != nil {
		errOutput(cmd, err)
	}
	serverID, err := openstackClient.GetServerIDFromName(Name)
	if err != nil {
		errOutput(cmd, err)
	}
	err = openstackClient.ServerDelete(serverID)
	if err != nil {
		errOutput(cmd, err)
	}
	cmd.Printf("server %s is deleted\n", Name)
}

var (
	// Name of server
	Name string
	// Project name or id
	Project string
	// Network name or id
	Network string
	// Flavor name
	Flavor string
	// Image name or id
	Image string
)

func init() {
	ServerCmd.AddCommand(ServerCreateCmd)
	ServerCmd.AddCommand(ServerListCmd)
	ServerCmd.AddCommand(ServerDeleteCmd)
	ServerCmd.PersistentFlags().StringVarP(&Name, "name", "n", "", "server name")
	ServerCmd.PersistentFlags().StringVarP(&Project, "project", "p", "", "project name")
	ServerCreateCmd.Flags().StringVarP(&Flavor, "flavor", "", "", "flavor name")
	ServerCreateCmd.Flags().StringVarP(&Network, "network", "", "", "network name")
	ServerCreateCmd.Flags().StringVarP(&Image, "image", "", "", "image name")
	ServerCmd.MarkPersistentFlagRequired("project")
}
