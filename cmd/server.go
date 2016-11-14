package cmd

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/spf13/cobra"
	"os"
)

// ServerCmd sub-command of zcloud about server
var ServerCmd = &cobra.Command{
	Use:   "server [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
}

var serverListCmd = &cobra.Command{
	Use:   "list [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: serverList,
}

func serverList(cmd *cobra.Command, args []string) {
	openstackClient := &OpenstackClient{
		host:       Host,
		username:   Username,
		password:   Password,
		tenantName: project,
	}
	opts := servers.ListOpts{}
	serverList, err := openstackClient.ServerList(opts)
	if err != nil {
		er(err)
	}

	for _, s := range serverList {
		cmd.Printf("%s\n", s.Name)
	}
}

var serverCreateCmd = &cobra.Command{
	Use:   "create [string to echo]",
	Short: "Echo anything to the screen",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: serverCreate,
}

func serverCreate(cmd *cobra.Command, args []string) {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://" + Host + ":5000/v2.0",
		Username:         Username,
		Password:         Password,
		TenantName:       project,
	}

	provider, err := openstack.AuthenticatedClient(authOpts)

	if err != nil {
		cmd.Printf("Authentication error: %s", err)
	}

	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
	// Region: "RegionOne",
	})

	if err != nil {
		cmd.Printf("Client init error: %s", err)
	}

	networks := [1]servers.Network{servers.Network{UUID: "66cfd9c4-173d-4322-a0d9-7f226079bc04"}}

	server, err := servers.Create(computeClient, servers.CreateOpts{
		Name:       name,
		FlavorName: flavor,
		ImageName:  "centos-7-1503-minimal[100G]",
		Networks:   networks[:],
	}).Extract()

	if err != nil {
		cmd.Printf("server creation error: %s", err)
	}
	cmd.Printf("server %s is created", server.Name)
}

var serverDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a specific server",
	Long: `echo is for echoing anything back.
    Echo echo’s.
    `,
	Run: serverDelete,
}

func serverDelete(cmd *cobra.Command, args []string) {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://" + Host + ":5000/v2.0",
		Username:         Username,
		Password:         Password,
		TenantID:         project,
	}

	provider, err := openstack.AuthenticatedClient(authOpts)

	if err != nil {
		cmd.Printf("Authentication error: %s", err)
		os.Exit(1)
	}
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		cmd.Printf("Client init error: %s", err)
		os.Exit(1)
	}
	serverID, err := servers.IDFromName(computeClient, name)
	if err != nil {
		cmd.Printf("Get server id err: %s\n", err)
		os.Exit(1)
	}
	result := servers.Delete(computeClient, serverID)
	if result.Err != nil {
		cmd.Printf("server delelte action error: %v\n", result.Err)
		os.Exit(1)
	}
	cmd.Printf("server %s is deleted\n", name)
}

var (
	// server name
	name string
	// project name or id
	project string
	// network name or id
	network string
	// flavor name
	flavor string
)

func init() {
	ServerCmd.AddCommand(serverCreateCmd)
	ServerCmd.AddCommand(serverListCmd)
	ServerCmd.AddCommand(serverDeleteCmd)
	ServerCmd.PersistentFlags().StringVarP(&name, "name", "n", "", "server name")
	ServerCmd.PersistentFlags().StringVarP(&project, "project", "p", "", "project name")
	ServerCmd.PersistentFlags().StringVarP(&network, "network", "", "", "network name")
	serverCreateCmd.Flags().StringVarP(&flavor, "flavor", "", "", "flavor name")
	ServerCmd.MarkPersistentFlagRequired("project")
}
