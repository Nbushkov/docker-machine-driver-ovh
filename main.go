package main

import (
	"github.com/docker/machine/drivers/openstack"
	"github.com/docker/machine/libmachine/drivers/plugin"
)

const (
	authUrl              = "https://auth.cloud.ovh.net/v2.0/"
	imageName            = "Ubuntu 14.04"
	SSHUser              = "admin"
	networkName          = "Ext-Net"
	defaultFlavorName    = "vps-ssd-1"
	defaultRegionName    = "GRA1"
	defaultSecurityGroup = "default"
)

func main() {
	plugin.RegisterDriver(&Driver{
		Driver: openstack.NewDerivedDriver("", ""),
	})
}