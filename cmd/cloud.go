// Package cmd provides the cobra cmd
package cmd

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"strings"
)

// IOpenstackClient dscribe the methods which openstack client behaves
type IOpenstackClient interface {
}

// OpenstackClient is the client of openstack service
type OpenstackClient struct {
	Host           string
	Username       string
	Password       string
	TenantName     string
	ComputeService *gophercloud.ServiceClient
}

// NewClient create a new client with given parameter
func NewClient(host string, username string, password string, tenantName string) (*OpenstackClient, error) {
	client := OpenstackClient{Host: host, Username: username, Password: password, TenantName: tenantName}
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://" + client.Host + ":5000/v2.0",
		Username:         client.Username,
		Password:         client.Password,
		TenantName:       client.TenantName,
	}
	provider, err := openstack.AuthenticatedClient(authOpts)
	if err != nil {
		return nil, err
	}
	client.ComputeService, err = openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// ServerList get the lis of server info in specific tenant
func (o *OpenstackClient) ServerList(opts servers.ListOpts) ([]servers.Server, error) {
	var serverListResult []servers.Server
	allPager, err := servers.List(o.ComputeService, opts).AllPages()
	if err != nil {
		return nil, err
	}
	serverListResult, err = servers.ExtractServers(allPager)
	if err != nil {
		return nil, err
	}
	return serverListResult, nil
}

// GetServerNetworkAddr extract network list from server info
func (o *OpenstackClient) GetServerNetworkAddr(server *servers.Server) (map[string][]servers.Address, error) {
	allPages, err := servers.ListAddresses(o.ComputeService, server.ID).AllPages()
	if err != nil {
		return nil, err
	}
	networks, err := servers.ExtractAddresses(allPages)
	if err != nil {
		return nil, err
	}
	return networks, nil
}

// GetServerNetworkAddr return network ip of specific server
func GetServerNetworkAddr(s *servers.Server) []string {
	var networks []string
	network = ""
	for key := range s.Addresses {
		var ips []string
		for _, networkInterface := range s.Addresses[key].([]interface{}) {
			if networkInterface.(map[string]interface{})["version"].(float64) == 4 {
				ips = append(ips, networkInterface.(map[string]interface{})["addr"].(string))
			}
		}
		ip := strings.Join(ips, ",")
		network = network + key + "=" + ip
		networks = append(networks, network)
	}
	return networks
}

// ServerCreate create the server with provided options
func (o *OpenstackClient) ServerCreate(opts servers.CreateOpts) (*servers.Server, error) {
	// computeClient, err := o.ComputeClient()
	var s *servers.Server
	s, err := servers.Create(o.ComputeService, opts).Extract()
	if err != nil {
		return s, err
	}
	return s, nil
}
