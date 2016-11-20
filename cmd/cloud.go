// Package cmd provides the cobra cmd
package cmd

import (
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	"github.com/gophercloud/gophercloud/pagination"
)

// OpenstackClient is the client of openstack service
type OpenstackClient struct {
	host       string
	username   string
	password   string
	tenantName string
	// computeClient *gophercloud.ServiceClient
}

// AuthOpts yields out auth options
func (o *OpenstackClient) AuthOpts() gophercloud.AuthOptions {
	authOpts := gophercloud.AuthOptions{
		IdentityEndpoint: "http://" + o.host + ":5000/v2.0",
		Username:         o.username,
		Password:         o.password,
		TenantName:       o.tenantName,
	}
	return authOpts
}

// ComputeClient yields opentack compute client
func (o *OpenstackClient) ComputeClient() (*gophercloud.ServiceClient, error) {
	authOpts := o.AuthOpts()
	provider, err := openstack.AuthenticatedClient(authOpts)

	if err != nil {
		// fmt.Printf("Authentication error: %s", err)
		return nil, err
	}
	computeClient, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
	// Region: "RegionOne",
	})
	if err != nil {
		// fmt.Printf("Client init error: %s", err)
		return nil, err
	}
	return computeClient, nil
}

// ServerList get the lis of server info in specific tenant
func (o *OpenstackClient) ServerList(opts servers.ListOpts) ([]servers.Server, error) {
	computeClient, err := o.ComputeClient()
	if err != nil {
		return nil, err
	}
	var serverListResult []servers.Server
	pager := servers.List(computeClient, opts)
	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		serverList, err := servers.ExtractServers(page)
		if err != nil {
			return false, err
		}
		for _, s := range serverList {
			serverListResult = append(serverListResult, s)
		}
		return true, nil
	})
	return serverListResult, err
}

// ServerCreate create the server with provided options
func (o *OpenstackClient) ServerCreate(opts servers.CreateOpts) (*servers.Server, error) {
	computeClient, err := o.ComputeClient()
	var s *servers.Server
	if err != nil {
		return s, err
	}
	s, err = servers.Create(computeClient, opts).Extract()
	if err != nil {
		return s, err
	}
	return s, nil
}
