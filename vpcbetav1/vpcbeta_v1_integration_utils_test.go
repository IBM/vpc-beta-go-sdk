//go:build integration
// +build integration

/**
 * (C) Copyright IBM Corp. 2020, 2021, 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package vpcbetav1_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
)

/**
 * REST methods
 *
 */
const (
	POST   = http.MethodPost
	GET    = http.MethodGet
	DELETE = http.MethodDelete
	PUT    = http.MethodPut
	PATCH  = http.MethodPatch
)

// InstantiateVPCService - Instantiate VPC Gen2 service
func InstantiateVPCService() *vpcbetav1.VpcbetaV1 {
	service, serviceErr := vpcbetav1.NewVpcbetaV1UsingExternalConfig(
		&vpcbetav1.VpcbetaV1Options{
			ServiceName: "vpcbetaint",
		},
	)
	// Check successful instantiation
	if serviceErr != nil {
		fmt.Println("Gen2 Service creation failed.", serviceErr)
		return nil
	}
	// return new vpc gen2 service
	return service
}

/**
 * Regions and Zones
 *
 */

// ListRegions - List all regions
// GET
// /regions
func ListRegions(gen2 *vpcbetav1.VpcbetaV1) (regions *vpcbetav1.RegionCollection, response *core.DetailedResponse, err error) {
	listRegionsOptions := &vpcbetav1.ListRegionsOptions{}
	regions, response, err = gen2.ListRegions(listRegionsOptions)
	return
}

// GetRegion - GETd
// /regions/{name}
// Retrieve a region
func GetRegion(vpcService *vpcbetav1.VpcbetaV1, name string) (region *vpcbetav1.Region, response *core.DetailedResponse, err error) {
	getRegionOptions := &vpcbetav1.GetRegionOptions{}
	getRegionOptions.SetName(name)
	region, response, err = vpcService.GetRegion(getRegionOptions)
	return
}

// ListZones - GET
// /regions/{region_name}/zones
// List all zones in a region
func ListZones(vpcService *vpcbetav1.VpcbetaV1, regionName string) (zones *vpcbetav1.ZoneCollection, response *core.DetailedResponse, err error) {
	listZonesOptions := &vpcbetav1.ListRegionZonesOptions{}
	listZonesOptions.SetRegionName(regionName)
	zones, response, err = vpcService.ListRegionZones(listZonesOptions)
	return
}

// GetZone - GET
// /regions/{region_name}/zones/{zone_name}
// Retrieve a zone
func GetZone(vpcService *vpcbetav1.VpcbetaV1, regionName, zoneName string) (zone *vpcbetav1.Zone, response *core.DetailedResponse, err error) {
	getZoneOptions := &vpcbetav1.GetRegionZoneOptions{}
	getZoneOptions.SetRegionName(regionName)
	getZoneOptions.SetName(zoneName)
	zone, response, err = vpcService.GetRegionZone(getZoneOptions)
	return
}

/**
 * Floating IPs
 */

// GetFloatingIPsList - GET
// /floating_ips
// List all floating IPs
func GetFloatingIPsList(vpcService *vpcbetav1.VpcbetaV1) (floatingIPs *vpcbetav1.FloatingIPCollection, response *core.DetailedResponse, err error) {
	listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()
	floatingIPs, response, err = vpcService.ListFloatingIps(listFloatingIpsOptions)
	return
}

// GetFloatingIP - GET
// /floating_ips/{id}
// Retrieve the specified floating IP
func GetFloatingIP(vpcService *vpcbetav1.VpcbetaV1, id string) (floatingIP *vpcbetav1.FloatingIP, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetFloatingIPOptions(id)
	floatingIP, response, err = vpcService.GetFloatingIP(options)
	return
}

// ReleaseFloatingIP - DELETE
// /floating_ips/{id}
// Release the specified floating IP
func ReleaseFloatingIP(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteFloatingIPOptions(id)
	response, err = vpcService.DeleteFloatingIP(options)
	return response, err
}

// UpdateFloatingIP - PATCH
// /floating_ips/{id}
// Update the specified floating IP
func UpdateFloatingIP(vpcService *vpcbetav1.VpcbetaV1, id, name string) (floatingIP *vpcbetav1.FloatingIP, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.FloatingIPPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateFloatingIPOptions{
		ID:              &id,
		FloatingIPPatch: patchBody,
	}

	floatingIP, response, err = vpcService.UpdateFloatingIP(options)
	return
}

// CreateFloatingIP - POST
// /floating_ips
// Reserve a floating IP
func CreateFloatingIP(vpcService *vpcbetav1.VpcbetaV1, zone, name string) (floatingIP *vpcbetav1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateFloatingIPOptions{}
	options.SetFloatingIPPrototype(&vpcbetav1.FloatingIPPrototype{
		Name: &name,
		Zone: &vpcbetav1.ZoneIdentity{
			Name: &zone,
		},
	})
	floatingIP, response, err = vpcService.CreateFloatingIP(options)
	return
}

/**
 * SSH Keys
 *
 */

// ListKeys - GET
// /keys
// List all keys
func ListKeys(vpcService *vpcbetav1.VpcbetaV1) (keys *vpcbetav1.KeyCollection, response *core.DetailedResponse, err error) {
	listKeysOptions := &vpcbetav1.ListKeysOptions{}
	keys, response, err = vpcService.ListKeys(listKeysOptions)
	return
}

// GetSSHKey - GET
// /keys/{id}
// Retrieve specified key
func GetSSHKey(vpcService *vpcbetav1.VpcbetaV1, id string) (key *vpcbetav1.Key, response *core.DetailedResponse, err error) {
	getKeyOptions := &vpcbetav1.GetKeyOptions{}
	getKeyOptions.SetID(id)
	key, response, err = vpcService.GetKey(getKeyOptions)
	return
}

// UpdateSSHKey - PATCH
// /keys/{id}
// Update specified key
func UpdateSSHKey(vpcService *vpcbetav1.VpcbetaV1, id, name string) (key *vpcbetav1.Key, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.KeyPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	updateKeyOptions := vpcService.NewUpdateKeyOptions(id, patchBody)
	key, response, err = vpcService.UpdateKey(updateKeyOptions)
	return
}

// DeleteSSHKey - DELETE
// /keys/{id}
// Delete specified key
func DeleteSSHKey(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	deleteKeyOptions := &vpcbetav1.DeleteKeyOptions{}
	deleteKeyOptions.SetID(id)
	response, err = vpcService.DeleteKey(deleteKeyOptions)
	return response, err
}

// CreateSSHKey - POST
// /keys
// Create a key
func CreateSSHKey(vpcService *vpcbetav1.VpcbetaV1, name, publicKey string) (key *vpcbetav1.Key, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateKeyOptions{}
	options.SetName(name)
	options.SetPublicKey(publicKey)
	key, response, err = vpcService.CreateKey(options)
	return
}

/**
 * VPC
 *
 */

// GetVPCsList - GET
// /vpcs
// List all VPCs
func ListVpcs(vpcService *vpcbetav1.VpcbetaV1) (vpcs *vpcbetav1.VPCCollection, response *core.DetailedResponse, err error) {
	listVpcsOptions := &vpcbetav1.ListVpcsOptions{}
	vpcs, response, err = vpcService.ListVpcs(listVpcsOptions)
	return
}

// GetVPC - GET
// /vpcs/{id}
// Retrieve specified VPC
func GetVPC(vpcService *vpcbetav1.VpcbetaV1, id string) (vpc *vpcbetav1.VPC, response *core.DetailedResponse, err error) {
	getVpcOptions := &vpcbetav1.GetVPCOptions{}
	getVpcOptions.SetID(id)
	vpc, response, err = vpcService.GetVPC(getVpcOptions)
	return
}

// DeleteVPC - DELETE
// /vpcs/{id}
// Delete specified VPC
func DeleteVPC(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	deleteVpcOptions := &vpcbetav1.DeleteVPCOptions{}
	deleteVpcOptions.SetID(id)
	response, err = vpcService.DeleteVPC(deleteVpcOptions)
	return response, err
}

// UpdateVPC - PATCH
// /vpcs/{id}
// Update specified VPC
func UpdateVPC(vpcService *vpcbetav1.VpcbetaV1, id, name string) (vpc *vpcbetav1.VPC, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.VPCPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := vpcService.NewUpdateVPCOptions(id, patchBody)
	vpc, response, err = vpcService.UpdateVPC(options)
	return
}

// CreateVPC - POST
// /vpcs
// Create a VPC
func CreateVPC(vpcService *vpcbetav1.VpcbetaV1, name, resourceGroup string) (vpc *vpcbetav1.VPC, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateVPCOptions{}

	options.SetResourceGroup(&vpcbetav1.ResourceGroupIdentity{
		ID: &resourceGroup,
	})
	options.SetName(name)
	vpc, response, err = vpcService.CreateVPC(options)
	return
}

/**
 * VPC default Security group
 * Getting default security group for a vpc with id
 */

// GetVPCDefaultSecurityGroup - GET
// /vpcs/{id}/default_security_group
// Retrieve a VPC's default security group
func GetVPCDefaultSecurityGroup(vpcService *vpcbetav1.VpcbetaV1, id string) (defaultSg *vpcbetav1.DefaultSecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVPCDefaultSecurityGroupOptions{}
	options.SetID(id)
	defaultSg, response, err = vpcService.GetVPCDefaultSecurityGroup(options)
	return
}

/**
 * VPC default ACL
 * Getting default security group for a vpc with id
 */

// GetVPCDefaultACL - GET
// /vpcs/{id}/default_network_acl
// Retrieve a VPC's default network acl
func GetVPCDefaultACL(vpcService *vpcbetav1.VpcbetaV1, id string) (defaultACL *vpcbetav1.DefaultNetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVPCDefaultNetworkACLOptions{}
	options.SetID(id)
	defaultACL, response, err = vpcService.GetVPCDefaultNetworkACL(options)
	return
}

/**
 * VPC address prefix
 *
 */

// ListVpcAddressPrefixes - GET
// /vpcs/{vpc_id}/address_prefixes
// List all address pool prefixes for a VPC
func ListVpcAddressPrefixes(vpcService *vpcbetav1.VpcbetaV1, vpcID string) (addressPrefixes *vpcbetav1.AddressPrefixCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVPCAddressPrefixesOptions{}
	options.SetVPCID(vpcID)
	addressPrefixes, response, err = vpcService.ListVPCAddressPrefixes(options)
	return
}

// GetVpcAddressPrefix - GET
// /vpcs/{vpc_id}/address_prefixes/{id}
// Retrieve specified address pool prefix
func GetVpcAddressPrefix(vpcService *vpcbetav1.VpcbetaV1, vpcID, addressPrefixID string) (addressPrefix *vpcbetav1.AddressPrefix, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVPCAddressPrefixOptions{}
	options.SetVPCID(vpcID)
	options.SetID(addressPrefixID)
	addressPrefix, response, err = vpcService.GetVPCAddressPrefix(options)
	return
}

// CreateVpcAddressPrefix - POST
// /vpcs/{vpc_id}/address_prefixes
// Create an address pool prefix
func CreateVpcAddressPrefix(vpcService *vpcbetav1.VpcbetaV1, vpcID, zone, cidr, name string) (addressPrefix *vpcbetav1.AddressPrefix, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateVPCAddressPrefixOptions{}
	options.SetVPCID(vpcID)
	options.SetCIDR(cidr)
	options.SetName(name)
	options.SetZone(&vpcbetav1.ZoneIdentity{
		Name: &zone,
	})
	addressPrefix, response, err = vpcService.CreateVPCAddressPrefix(options)
	return
}

// DeleteVpcAddressPrefix - DELETE
// /vpcs/{vpc_id}/address_prefixes/{id}
// Delete specified address pool prefix
func DeleteVpcAddressPrefix(vpcService *vpcbetav1.VpcbetaV1, vpcID, addressPrefixID string) (response *core.DetailedResponse, err error) {
	deleteVpcAddressPrefixOptions := &vpcbetav1.DeleteVPCAddressPrefixOptions{}
	deleteVpcAddressPrefixOptions.SetVPCID(vpcID)
	deleteVpcAddressPrefixOptions.SetID(addressPrefixID)
	response, err = vpcService.DeleteVPCAddressPrefix(deleteVpcAddressPrefixOptions)
	return response, err
}

// UpdateVpcAddressPrefix - PATCH
// /vpcs/{vpc_id}/address_prefixes/{id}
// Update an address pool prefix
func UpdateVpcAddressPrefix(vpcService *vpcbetav1.VpcbetaV1, vpcID, addressPrefixID, name string) (addressPrefix *vpcbetav1.AddressPrefix, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := vpcService.NewUpdateVPCAddressPrefixOptions(vpcID, addressPrefixID, patchBody)
	addressPrefix, response, err = vpcService.UpdateVPCAddressPrefix(options)
	return
}

/**
 * VPC routes
 *
 */

// ListVpcRoutes - GET
// /vpcs/{vpc_id}/routes
// List all user-defined routes for a VPC
func ListVpcRoutes(vpcService *vpcbetav1.VpcbetaV1, vpcID string) (routes *vpcbetav1.RouteCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVPCRoutesOptions{}
	options.SetVPCID(vpcID)
	routes, response, err = vpcService.ListVPCRoutes(options)
	return
}

// GetVpcRoute - GET
// /vpcs/{vpc_id}/routes/{id}
// Retrieve the specified route
func GetVpcRoute(vpcService *vpcbetav1.VpcbetaV1, vpcID, routeID string) (route *vpcbetav1.Route, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	route, response, err = vpcService.GetVPCRoute(options)
	return
}

// CreateVpcRoute - POST
// /vpcs/{vpc_id}/routes
// Create a route on your VPC
func CreateVpcRoute(vpcService *vpcbetav1.VpcbetaV1, vpcID, zone, destination, nextHopAddress, name string) (route *vpcbetav1.Route, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetName(name)
	options.SetZone(&vpcbetav1.ZoneIdentity{
		Name: &zone,
	})
	options.SetNextHop(&vpcbetav1.RoutePrototypeNextHop{
		Address: &nextHopAddress,
	})
	options.SetDestination(destination)
	route, response, err = vpcService.CreateVPCRoute(options)
	return
}

// DeleteVpcRoute - DELETE
// /vpcs/{vpc_id}/routes/{id}
// Delete the specified route
func DeleteVpcRoute(vpcService *vpcbetav1.VpcbetaV1, vpcID, routeID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteVPCRouteOptions{}
	options.SetVPCID(vpcID)
	options.SetID(routeID)
	response, err = vpcService.DeleteVPCRoute(options)
	return response, err
}

// UpdateVpcRoute - PATCH
// /vpcs/{vpc_id}/routes/{id}
// Update a route
func UpdateVpcRoute(vpcService *vpcbetav1.VpcbetaV1, vpcID, routeID, name string) (route *vpcbetav1.Route, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.RoutePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateVPCRouteOptions{
		RoutePatch: patchBody,
		VPCID:      &vpcID,
		ID:         &routeID,
	}
	route, response, err = vpcService.UpdateVPCRoute(options)
	return
}

/**
 * Volumes
 *
 */

// ListVolumeProfiles - GET
// /volume/profiles
// List all volume profiles
func ListVolumeProfiles(vpcService *vpcbetav1.VpcbetaV1) (profiles *vpcbetav1.VolumeProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVolumeProfilesOptions{}
	profiles, response, err = vpcService.ListVolumeProfiles(options)
	return
}

// GetVolumeProfile - GET
// /volume/profiles/{name}
// Retrieve specified volume profile
func GetVolumeProfile(vpcService *vpcbetav1.VpcbetaV1, profileName string) (profile *vpcbetav1.VolumeProfile, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVolumeProfileOptions{}
	options.SetName(profileName)
	profile, response, err = vpcService.GetVolumeProfile(options)
	return
}

// ListVolumes - GET
// /volumes
// List all volumes
func ListVolumes(vpcService *vpcbetav1.VpcbetaV1) (volumes *vpcbetav1.VolumeCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVolumesOptions{}
	volumes, response, err = vpcService.ListVolumes(options)
	return
}

// GetVolume - GET
// /volumes/{id}
// Retrieve specified volume
func GetVolume(vpcService *vpcbetav1.VpcbetaV1, volumeID string) (volume *vpcbetav1.Volume, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVolumeOptions{}
	options.SetID(volumeID)
	volume, response, err = vpcService.GetVolume(options)
	return
}

// DeleteVolume - DELETE
// /volumes/{id}
// Delete specified volume
func DeleteVolume(vpcService *vpcbetav1.VpcbetaV1, id, ifMatch string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteVolumeOptions{}
	options.SetID(id)
	options.SetIfMatch(ifMatch)
	response, err = vpcService.DeleteVolume(options)
	return response, err
}

// UpdateVolume - PATCH
// /volumes/{id}
// Update specified volume
func UpdateVolume(vpcService *vpcbetav1.VpcbetaV1, userTags []string, id, name, ifMatch string) (volume *vpcbetav1.Volume, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.VolumePatch{
		Name:     &name,
		UserTags: userTags,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateVolumeOptions{
		VolumePatch: patchBody,
		ID:          &id,
		IfMatch:     &ifMatch,
	}
	volume, response, err = vpcService.UpdateVolume(options)
	return
}

// CreateVolume - POST
// /volumes
// Create a volume
func CreateVolume(vpcService *vpcbetav1.VpcbetaV1, name, profileName, zoneName string, capacity int64) (volume *vpcbetav1.Volume, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateVolumeOptions{}
	options.SetVolumePrototype(&vpcbetav1.VolumePrototype{
		Capacity: core.Int64Ptr(capacity),
		Zone: &vpcbetav1.ZoneIdentity{
			Name: &zoneName,
		},
		Profile: &vpcbetav1.VolumeProfileIdentity{
			Name: &profileName,
		},
		Name: &name,
	})
	volume, response, err = vpcService.CreateVolume(options)
	return
}

/**
 * Subnets
 *
 */

// ListSubnets - GET
// /subnets
// List all subnets
func ListSubnets(vpcService *vpcbetav1.VpcbetaV1) (subnets *vpcbetav1.SubnetCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListSubnetsOptions{}
	subnets, response, err = vpcService.ListSubnets(options)
	return
}

// GetSubnet - GET
// /subnets/{id}
// Retrieve specified subnet
func GetSubnet(vpcService *vpcbetav1.VpcbetaV1, subnetID string) (subnet *vpcbetav1.Subnet, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetSubnetOptions{}
	options.SetID(subnetID)
	subnet, response, err = vpcService.GetSubnet(options)
	return
}

// DeleteSubnet - DELETE
// /subnets/{id}
// Delete specified subnet
func DeleteSubnet(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteSubnetOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteSubnet(options)
	return response, err
}

// UpdateSubnet - PATCH
// /subnets/{id}
// Update specified subnet
func UpdateSubnet(vpcService *vpcbetav1.VpcbetaV1, id, name string) (subnet *vpcbetav1.Subnet, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.SubnetPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateSubnetOptions{
		SubnetPatch: patchBody,
	}
	options.SetID(id)
	subnet, response, err = vpcService.UpdateSubnet(options)
	return
}

// CreateSubnet - POST
// /subnets
// Create a subnet
func CreateSubnet(vpcService *vpcbetav1.VpcbetaV1, vpcID, name, zone string, mock bool) (subnet *vpcbetav1.Subnet, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateSubnetOptions{}
	if mock {
		options.SetSubnetPrototype(&vpcbetav1.SubnetPrototype{
			Ipv4CIDRBlock: core.StringPtr("10.243.0.0/24"),
			Name:          &name,
			VPC: &vpcbetav1.VPCIdentity{
				ID: &vpcID,
			},
			Zone: &vpcbetav1.ZoneIdentity{
				Name: &zone,
			},
		})
	} else {
		options.SetSubnetPrototype(&vpcbetav1.SubnetPrototype{
			Name: &name,
			VPC: &vpcbetav1.VPCIdentity{
				ID: &vpcID,
			},
			Zone: &vpcbetav1.ZoneIdentity{
				Name: &zone,
			},
			TotalIpv4AddressCount: core.Int64Ptr(128),
		})
	}
	subnet, response, err = vpcService.CreateSubnet(options)
	return
}

// GetSubnetNetworkAcl -GET
// /subnets/{id}/network_acl
// Retrieve a subnet's attached network ACL
func GetSubnetNetworkAcl(vpcService *vpcbetav1.VpcbetaV1, subnetID string) (subnetACL *vpcbetav1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetSubnetNetworkACLOptions{}
	options.SetID(subnetID)
	subnetACL, response, err = vpcService.GetSubnetNetworkACL(options)
	return
}

// SetSubnetNetworkAclBinding - PUT
// /subnets/{id}/network_acl
// Attach a network ACL to a subnet
func SetSubnetNetworkAclBinding(vpcService *vpcbetav1.VpcbetaV1, subnetID, id string) (nacl *vpcbetav1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ReplaceSubnetNetworkACLOptions{}
	options.SetID(subnetID)
	options.SetNetworkACLIdentity(&vpcbetav1.NetworkACLIdentity{ID: &id})
	nacl, response, err = vpcService.ReplaceSubnetNetworkACL(options)
	return
}

// DeleteSubnetPublicGatewayBinding - DELETE
// /subnets/{id}/public_gateway
// Detach a public gateway from a subnet
func DeleteSubnetPublicGatewayBinding(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.UnsetSubnetPublicGatewayOptions{}
	options.SetID(id)
	response, err = vpcService.UnsetSubnetPublicGateway(options)
	return response, err
}

// GetSubnetPublicGateway - GET
// /subnets/{id}/public_gateway
// Retrieve a subnet's attached public gateway
func GetSubnetPublicGateway(vpcService *vpcbetav1.VpcbetaV1, subnetID string) (pgw *vpcbetav1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	pgw, response, err = vpcService.GetSubnetPublicGateway(options)
	return
}

// SetSubnetPublicGatewayBinding - PUT
// /subnets/{id}/public_gateway
// Attach a public gateway to a subnet
func CreateSubnetPublicGatewayBinding(vpcService *vpcbetav1.VpcbetaV1, subnetID, id string) (pgw *vpcbetav1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.SetSubnetPublicGatewayOptions{}
	options.SetID(subnetID)
	options.SetPublicGatewayIdentity(&vpcbetav1.PublicGatewayIdentity{ID: &id})
	pgw, response, err = vpcService.SetSubnetPublicGateway(options)
	return
}

/**
 * Subnet Reserved IPs
 *
 */

// GET
// /subnets/{subnet_id}/reserved_ips
// List all reserved IPs in a subnet
func ListSubnetReservedIps(vpcService *vpcbetav1.VpcbetaV1, subnetId string) (reservedIPCollection *vpcbetav1.ReservedIPCollection, response *core.DetailedResponse, err error) {
	listSubnetReservedIpsOptions := vpcService.NewListSubnetReservedIpsOptions(
		subnetId,
	)
	reservedIPCollection, response, err = vpcService.ListSubnetReservedIps(listSubnetReservedIpsOptions)
	return
}

// POST
// /subnets/{subnet_id}/reserved_ips
// Reserve an IP in a subnet
func CreateSubnetReservedIP(vpcService *vpcbetav1.VpcbetaV1, subnetId, name string) (reservedIP *vpcbetav1.ReservedIP, response *core.DetailedResponse, err error) {
	createSubnetReservedIPOptions := &vpcbetav1.CreateSubnetReservedIPOptions{
		SubnetID: &subnetId,
		Name:     &name,
	}
	reservedIP, response, err = vpcService.CreateSubnetReservedIP(createSubnetReservedIPOptions)
	return
}

// DELETE
// /subnets/{subnet_id}/reserved_ips/{id}
// Release specified reserved IP
func DeleteSubnetReservedIP(vpcService *vpcbetav1.VpcbetaV1, subnetId, reservedIPId string) (response *core.DetailedResponse, err error) {
	deleteSubnetReservedIPOptions := vpcService.NewDeleteSubnetReservedIPOptions(
		subnetId,
		reservedIPId,
	)

	response, err = vpcService.DeleteSubnetReservedIP(deleteSubnetReservedIPOptions)
	return response, err
}

// GET
// /subnets/{subnet_id}/reserved_ips/{id}
// Retrieve specified reserved IP
func GetSubnetReservedIP(vpcService *vpcbetav1.VpcbetaV1, subnetId, reservedIPId string) (reservedIP *vpcbetav1.ReservedIP, response *core.DetailedResponse, err error) {
	getSubnetReservedIPOptions := vpcService.NewGetSubnetReservedIPOptions(
		subnetId,
		reservedIPId,
	)
	reservedIP, response, err = vpcService.GetSubnetReservedIP(getSubnetReservedIPOptions)
	return
}

// PATCH
// /subnets/{subnet_id}/reserved_ips/{id}
// Update specified reserved IP
func UpdateSubnetReservedIP(vpcService *vpcbetav1.VpcbetaV1, subnetId, reservedIPId, name string) (reservedIP *vpcbetav1.ReservedIP, response *core.DetailedResponse, err error) {
	reservedIPPatchModel := &vpcbetav1.ReservedIPPatch{
		Name: &name,
	}
	reservedIPPatchModelAsPatch, _ := reservedIPPatchModel.AsPatch()

	updateSubnetReservedIPOptions := vpcService.NewUpdateSubnetReservedIPOptions(
		subnetId,
		reservedIPId,
		reservedIPPatchModelAsPatch,
	)
	reservedIP, response, err = vpcService.UpdateSubnetReservedIP(updateSubnetReservedIPOptions)
	return
}

/**
 * Images
 *
 */

// ListImages - GET
// /images
// List all images
func ListImages(vpcService *vpcbetav1.VpcbetaV1, visibility string) (images *vpcbetav1.ImageCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListImagesOptions{}
	options.SetVisibility(visibility)
	images, response, err = vpcService.ListImages(options)
	return
}

// GetImage - GET
// /images/{id}
// Retrieve the specified image
func GetImage(vpcService *vpcbetav1.VpcbetaV1, imageID string) (image *vpcbetav1.Image, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetImageOptions{}
	options.SetID(imageID)
	image, response, err = vpcService.GetImage(options)
	return
}

// DeleteImage DELETE
// /images/{id}
// Delete specified image
func DeleteImage(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteImageOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteImage(options)
	return response, err
}

// UpdateImage PATCH
// /images/{id}
// Update specified image
func UpdateImage(vpcService *vpcbetav1.VpcbetaV1, id, name string) (image *vpcbetav1.Image, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.ImagePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateImageOptions{
		ImagePatch: patchBody,
	}
	options.SetID(id)
	image, response, err = vpcService.UpdateImage(options)
	return
}

func CreateImage(vpcService *vpcbetav1.VpcbetaV1, name string) (image *vpcbetav1.Image, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateImageOptions{}
	cosID := "cos://cos-location-of-image-file"
	options.SetImagePrototype(&vpcbetav1.ImagePrototype{
		Name: &name,
		File: &vpcbetav1.ImageFilePrototype{
			Href: &cosID,
		},
	})
	image, response, err = vpcService.CreateImage(options)
	return
}

func ListOperatingSystems(vpcService *vpcbetav1.VpcbetaV1) (operatingSystems *vpcbetav1.OperatingSystemCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListOperatingSystemsOptions{}
	operatingSystems, response, err = vpcService.ListOperatingSystems(options)
	return
}

func GetOperatingSystem(vpcService *vpcbetav1.VpcbetaV1, osName string) (os *vpcbetav1.OperatingSystem, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetOperatingSystemOptions{}
	options.SetName(osName)
	os, response, err = vpcService.GetOperatingSystem(options)
	return
}

/**
 * Instances
 *
 */

// ListInstanceProfiles - GET
// /instance/profiles
// List all instance profiles
func ListInstanceProfiles(vpcService *vpcbetav1.VpcbetaV1) (profiles *vpcbetav1.InstanceProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceProfilesOptions{}
	profiles, response, err = vpcService.ListInstanceProfiles(options)
	return
}

// GetInstanceProfile - GET
// /instance/profiles/{name}
// Retrieve specified instance profile
func GetInstanceProfile(vpcService *vpcbetav1.VpcbetaV1, profileName string) (profile *vpcbetav1.InstanceProfile, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceProfileOptions{}
	options.SetName(profileName)
	profile, response, err = vpcService.GetInstanceProfile(options)
	return
}

// ListInstances GET
// /instances
// List all instances
func ListInstances(vpcService *vpcbetav1.VpcbetaV1) (instances *vpcbetav1.InstanceCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstancesOptions{}
	instances, response, err = vpcService.ListInstances(options)
	return
}

// GetInstance GET
// instances/{id}
// Retrieve an instance
func GetInstance(vpcService *vpcbetav1.VpcbetaV1, instanceID string) (instance *vpcbetav1.Instance, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceOptions{}
	options.SetID(instanceID)
	instance, response, err = vpcService.GetInstance(options)
	return
}

// DeleteInstance DELETE
// /instances/{id}
// Delete specified instance
func DeleteInstance(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstance(options)
	return response, err
}

// UpdateInstance PATCH
// /instances/{id}
// Update specified instance
func UpdateInstance(vpcService *vpcbetav1.VpcbetaV1, id, name string) (instance *vpcbetav1.Instance, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.InstancePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceOptions{
		InstancePatch: patchBody,
		ID:            &id,
	}
	instance, response, err = vpcService.UpdateInstance(options)
	return
}

// CreateInstance POST
// /instances/{instance_id}
// Create an instance action
func CreateInstance(vpcService *vpcbetav1.VpcbetaV1, name, profileName, imageID, zoneName, subnetID, sshkeyID, vpcID string) (instance *vpcbetav1.Instance, response *core.DetailedResponse, err error) {
	volumeProfileIdentityModel := new(vpcbetav1.VolumeProfileIdentityByName)
	volumeProfileIdentityModel.Name = core.StringPtr("general-purpose")

	volume := new(vpcbetav1.VolumeAttachmentPrototypeVolumeVolumePrototypeInstanceContext)
	volume.Capacity = core.Int64Ptr(int64(100))
	volume.Name = core.StringPtr("my-volume")
	volume.Profile = volumeProfileIdentityModel

	volumeAttachmentPrototypeModel := new(vpcbetav1.VolumeAttachmentPrototype)
	volumeAttachmentPrototypeModel.DeleteVolumeOnInstanceDelete = core.BoolPtr(true)
	volumeAttachmentPrototypeModel.Name = core.StringPtr("my-volume-attachment")
	volumeAttachmentPrototypeModel.Volume = volume

	options := &vpcbetav1.CreateInstanceOptions{}
	options.SetInstancePrototype(&vpcbetav1.InstancePrototypeInstanceByImage{
		Keys:                    []vpcbetav1.KeyIdentityIntf{&vpcbetav1.KeyIdentity{ID: &sshkeyID}},
		Name:                    &name,
		NetworkInterfaces:       []vpcbetav1.NetworkInterfacePrototype{},
		Profile:                 &vpcbetav1.InstanceProfileIdentity{Name: &profileName},
		VolumeAttachments:       []vpcbetav1.VolumeAttachmentPrototype{*volumeAttachmentPrototypeModel},
		VPC:                     &vpcbetav1.VPCIdentity{ID: &vpcID},
		Image:                   &vpcbetav1.ImageIdentity{ID: &imageID},
		PrimaryNetworkInterface: &vpcbetav1.NetworkInterfacePrototype{Subnet: &vpcbetav1.SubnetIdentity{ID: &subnetID}},
		Zone:                    &vpcbetav1.ZoneIdentity{Name: &zoneName},
	})

	instance, response, err = vpcService.CreateInstance(options)
	return
}

// CreateInstanceAction PATCH
// /instances/{instance_id}/actions
// Update specified instance
func CreateInstanceAction(vpcService *vpcbetav1.VpcbetaV1, instanceID, typeOfAction string) (action *vpcbetav1.InstanceAction, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateInstanceActionOptions{}
	options.SetInstanceID(instanceID)
	options.SetType(typeOfAction)
	action, response, err = vpcService.CreateInstanceAction(options)
	return
}

// GetInstanceInitialization GET
// /instances/{id}/initialization
// Retrieve configuration used to initialize the instance.
func GetInstanceInitialization(vpcService *vpcbetav1.VpcbetaV1, instanceID string) (initData *vpcbetav1.InstanceInitialization, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceInitializationOptions{}
	options.SetID(instanceID)
	initData, response, err = vpcService.GetInstanceInitialization(options)
	return
}

// ListNetworkInterfaces GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaces(vpcService *vpcbetav1.VpcbetaV1, id string) (networkInterfaces *vpcbetav1.NetworkInterfaceUnpaginatedCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceNetworkInterfacesOptions{}
	options.SetInstanceID(id)
	networkInterfaces, response, err = vpcService.ListInstanceNetworkInterfaces(options)
	return
}

// CreateNetworkInterface POST
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func CreateNetworkInterface(vpcService *vpcbetav1.VpcbetaV1, id, subnetID string) (networkInterface *vpcbetav1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateInstanceNetworkInterfaceOptions{}
	options.SetInstanceID(id)
	options.SetName("eth1")
	options.SetSubnet(&vpcbetav1.SubnetIdentityByID{
		ID: &subnetID,
	})
	networkInterface, response, err = vpcService.CreateInstanceNetworkInterface(options)
	return
}

// DeleteNetworkInterface Delete
// /instances/{instance_id}/network_interfaces/{id}
// Retrieve specified network interface
func DeleteNetworkInterface(vpcService *vpcbetav1.VpcbetaV1, instanceID, vnicID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceNetworkInterfaceOptions{}
	options.SetID(vnicID)
	options.SetInstanceID(instanceID)
	response, err = vpcService.DeleteInstanceNetworkInterface(options)
	return response, err
}

// GetNetworkInterface GET
// /instances/{instance_id}/network_interfaces/{id}
// Retrieve specified network interface
func GetNetworkInterface(vpcService *vpcbetav1.VpcbetaV1, instanceID, networkID string) (networkInterface *vpcbetav1.NetworkInterface, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceNetworkInterfaceOptions{}
	options.SetID(networkID)
	options.SetInstanceID(instanceID)
	networkInterface, response, err = vpcService.GetInstanceNetworkInterface(options)
	return
}

// UpdateNetworkInterface PATCH
// /instances/{instance_id}/network_interfaces/{id}
// Update a network interface
func UpdateNetworkInterface(vpcService *vpcbetav1.VpcbetaV1, instanceID, networkID, name string) (networkInterface *vpcbetav1.NetworkInterface, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.NetworkInterfacePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceNetworkInterfaceOptions{
		NetworkInterfacePatch: patchBody,
		ID:                    &networkID,
		InstanceID:            &instanceID,
	}
	networkInterface, response, err = vpcService.UpdateInstanceNetworkInterface(options)
	return
}

// ListNetworkInterfaceFloatingIps GET
// /instances/{instance_id}/network_interfaces
// List all network interfaces on an instance
func ListNetworkInterfaceFloatingIps(vpcService *vpcbetav1.VpcbetaV1, instanceID, networkID string) (fips *vpcbetav1.FloatingIPUnpaginatedCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceNetworkInterfaceFloatingIpsOptions{}
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	fips, response, err = vpcService.ListInstanceNetworkInterfaceFloatingIps(options)
	return
}

// GetNetworkInterfaceFloatingIp GET
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips
// List all floating IPs associated with a network interface
func GetNetworkInterfaceFloatingIp(vpcService *vpcbetav1.VpcbetaV1, instanceID, networkID, fipID string) (fip *vpcbetav1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	fip, response, err = vpcService.GetInstanceNetworkInterfaceFloatingIP(options)
	return
}

// DeleteNetworkInterfaceFloatingIpBinding DELETE
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Disassociate specified floating IP
func DeleteNetworkInterfaceFloatingIpBinding(vpcService *vpcbetav1.VpcbetaV1, instanceID, networkID, fipID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.RemoveInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	response, err = vpcService.RemoveInstanceNetworkInterfaceFloatingIP(options)
	return response, err
}

// CreateNetworkInterfaceFloatingIpBinding PUT
// /instances/{instance_id}/network_interfaces/{network_interface_id}/floating_ips/{id}
// Associate a floating IP with a network interface
func CreateNetworkInterfaceFloatingIpBinding(vpcService *vpcbetav1.VpcbetaV1, instanceID, networkID, fipID string) (fip *vpcbetav1.FloatingIP, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.AddInstanceNetworkInterfaceFloatingIPOptions{}
	options.SetID(fipID)
	options.SetInstanceID(instanceID)
	options.SetNetworkInterfaceID(networkID)
	fip, response, err = vpcService.AddInstanceNetworkInterfaceFloatingIP(options)
	return
}

// ListVolumeAttachments GET
// /instances/{instance_id}/volume_attachments
// List all volumes attached to an instance
func ListVolumeAttachments(vpcService *vpcbetav1.VpcbetaV1, id string) (volumeAttachments *vpcbetav1.VolumeAttachmentCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceVolumeAttachmentsOptions{}
	options.SetInstanceID(id)
	volumeAttachments, response, err = vpcService.ListInstanceVolumeAttachments(options)
	return
}

// CreateVolumeAttachment POST
// /instances/{instance_id}/volume_attachments
// Create a volume attachment, connecting a volume to an instance
func CreateVolumeAttachment(vpcService *vpcbetav1.VpcbetaV1, instanceID, volumeID, name string) (volumeAttachment *vpcbetav1.VolumeAttachment, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetVolume(&vpcbetav1.VolumeAttachmentPrototypeVolume{
		ID: &volumeID,
	})
	options.SetName(name)
	options.SetDeleteVolumeOnInstanceDelete(false)
	volumeAttachment, response, err = vpcService.CreateInstanceVolumeAttachment(options)
	return
}

// DeleteVolumeAttachment DELETE
// /instances/{instance_id}/volume_attachments/{id}
// Delete a volume attachment, detaching a volume from an instance
func DeleteVolumeAttachment(vpcService *vpcbetav1.VpcbetaV1, instanceID, volumeID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceVolumeAttachmentOptions{}
	options.SetID(volumeID)
	options.SetInstanceID(instanceID)
	response, err = vpcService.DeleteInstanceVolumeAttachment(options)
	return response, err
}

// GetVolumeAttachment GET
// /instances/{instance_id}/volume_attachments/{id}
// Retrieve specified volume attachment
func GetVolumeAttachment(vpcService *vpcbetav1.VpcbetaV1, instanceID, volumeID string) (volumeAttachment *vpcbetav1.VolumeAttachment, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceVolumeAttachmentOptions{}
	options.SetInstanceID(instanceID)
	options.SetID(volumeID)
	volumeAttachment, response, err = vpcService.GetInstanceVolumeAttachment(options)
	return
}

// UpdateVolumeAttachment PATCH
// /instances/{instance_id}/volume_attachments/{id}
// Update a volume attachment
func UpdateVolumeAttachment(vpcService *vpcbetav1.VpcbetaV1, instanceID, volumeID, name string) (volumeAttachment *vpcbetav1.VolumeAttachment, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.VolumeAttachmentPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceVolumeAttachmentOptions{
		VolumeAttachmentPatch: patchBody,
		InstanceID:            &instanceID,
		ID:                    &volumeID,
	}
	volumeAttachment, response, err = vpcService.UpdateInstanceVolumeAttachment(options)
	return
}

/**
 * Public Gateway
 *
 */

// ListPublicGateways GET
// /public_gateways
// List all public gateways
func ListPublicGateways(vpcService *vpcbetav1.VpcbetaV1) (pgws *vpcbetav1.PublicGatewayCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListPublicGatewaysOptions{}
	pgws, response, err = vpcService.ListPublicGateways(options)
	return
}

// CreatePublicGateway POST
// /public_gateways
// Create a public gateway
func CreatePublicGateway(vpcService *vpcbetav1.VpcbetaV1, name, vpcID, zoneName string) (pgw *vpcbetav1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreatePublicGatewayOptions{}
	options.SetVPC(&vpcbetav1.VPCIdentity{
		ID: &vpcID,
	})
	options.SetZone(&vpcbetav1.ZoneIdentity{
		Name: &zoneName,
	})
	pgw, response, err = vpcService.CreatePublicGateway(options)
	return
}

// DeletePublicGateway DELETE
// /public_gateways/{id}
// Delete specified public gateway
func DeletePublicGateway(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeletePublicGatewayOptions{}
	options.SetID(id)
	response, err = vpcService.DeletePublicGateway(options)
	return response, err
}

// GetPublicGateway GET
// /public_gateways/{id}
// Retrieve specified public gateway
func GetPublicGateway(vpcService *vpcbetav1.VpcbetaV1, id string) (pgw *vpcbetav1.PublicGateway, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetPublicGatewayOptions{}
	options.SetID(id)
	pgw, response, err = vpcService.GetPublicGateway(options)
	return
}

// UpdatePublicGateway PATCH
// /public_gateways/{id}
// Update a public gateway's name
func UpdatePublicGateway(vpcService *vpcbetav1.VpcbetaV1, id, name string) (pgw *vpcbetav1.PublicGateway, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.PublicGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdatePublicGatewayOptions{
		PublicGatewayPatch: patchBody,
		ID:                 &id,
	}
	pgw, response, err = vpcService.UpdatePublicGateway(options)
	return
}

/**
 * Network ACLs not available in gen2 currently
 *
 */

// ListNetworkAcls - GET
// /network_acls
// List all network ACLs
func ListNetworkAcls(vpcService *vpcbetav1.VpcbetaV1) (networkACLs *vpcbetav1.NetworkACLCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListNetworkAclsOptions{}
	networkACLs, response, err = vpcService.ListNetworkAcls(options)
	return
}

// CreateNetworkAcl - POST
// /network_acls
// Create a network ACL
func CreateNetworkAcl(vpcService *vpcbetav1.VpcbetaV1, name, copyableAclID, vpcID string) (networkACL *vpcbetav1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateNetworkACLOptions{}
	options.SetNetworkACLPrototype(&vpcbetav1.NetworkACLPrototype{
		Name: &name,
		SourceNetworkACL: &vpcbetav1.NetworkACLIdentity{
			ID: &copyableAclID,
		},
		VPC: &vpcbetav1.VPCIdentity{
			ID: &vpcID,
		},
	})
	networkACL, response, err = vpcService.CreateNetworkACL(options)
	return
}

// DeleteNetworkAcl - DELETE
// /network_acls/{id}
// Delete specified network ACL
func DeleteNetworkAcl(vpcService *vpcbetav1.VpcbetaV1, ID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteNetworkACLOptions{}
	options.SetID(ID)
	response, err = vpcService.DeleteNetworkACL(options)
	return response, err
}

// GetNetworkAcl - GET
// /network_acls/{id}
// Retrieve specified network ACL
func GetNetworkAcl(vpcService *vpcbetav1.VpcbetaV1, ID string) (networkACL *vpcbetav1.NetworkACL, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetNetworkACLOptions{}
	options.SetID(ID)
	networkACL, response, err = vpcService.GetNetworkACL(options)
	return
}

// UpdateNetworkAcl PATCH
// /network_acls/{id}
// Update a network ACL
func UpdateNetworkAcl(vpcService *vpcbetav1.VpcbetaV1, id, name string) (networkACL *vpcbetav1.NetworkACL, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.NetworkACLPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateNetworkACLOptions{
		NetworkACLPatch: patchBody,
		ID:              &id,
	}
	networkACL, response, err = vpcService.UpdateNetworkACL(options)
	return
}

// ListNetworkAclRules - GET
// /network_acls/{network_acl_id}/rules
// List all rules for a network ACL
func ListNetworkAclRules(vpcService *vpcbetav1.VpcbetaV1, aclID string) (networkACL *vpcbetav1.NetworkACLRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListNetworkACLRulesOptions{}
	options.SetNetworkACLID(aclID)
	networkACL, response, err = vpcService.ListNetworkACLRules(options)
	return
}

// CreateNetworkAclRule - POST
// /network_acls/{network_acl_id}/rules
// Create a rule
func CreateNetworkAclRule(vpcService *vpcbetav1.VpcbetaV1, name, aclID string) (rules vpcbetav1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateNetworkACLRuleOptions{}
	options.SetNetworkACLID(aclID)
	options.SetNetworkACLRulePrototype(&vpcbetav1.NetworkACLRulePrototype{
		Action:      core.StringPtr("allow"),
		Direction:   core.StringPtr("inbound"),
		Destination: core.StringPtr("0.0.0.0/0"),
		Source:      core.StringPtr("0.0.0.0/0"),
		Protocol:    core.StringPtr("all"),
		Name:        &name,
	})
	rules, response, err = vpcService.CreateNetworkACLRule(options)
	return
}

// DeleteNetworkAclRule DELETE
// /network_acls/{network_acl_id}/rules/{id}
// Delete specified rule
func DeleteNetworkAclRule(vpcService *vpcbetav1.VpcbetaV1, aclID, ruleID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	response, err = vpcService.DeleteNetworkACLRule(options)
	return response, err
}

// GetNetworkAclRule GET
// /network_acls/{network_acl_id}/rules/{id}
// Retrieve specified rule
func GetNetworkAclRule(vpcService *vpcbetav1.VpcbetaV1, aclID, ruleID string) (rule vpcbetav1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetNetworkACLRuleOptions{}
	options.SetID(ruleID)
	options.SetNetworkACLID(aclID)
	rule, response, err = vpcService.GetNetworkACLRule(options)
	return
}

// UpdateNetworkAclRule PATCH
// /network_acls/{network_acl_id}/rules/{id}
// Update a rule
func UpdateNetworkAclRule(vpcService *vpcbetav1.VpcbetaV1, aclID, ruleID, name string) (rule vpcbetav1.NetworkACLRuleIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.NetworkACLRulePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateNetworkACLRuleOptions{
		NetworkACLRulePatch: patchBody,
		ID:                  &ruleID,
		NetworkACLID:        &aclID,
	}
	rule, response, err = vpcService.UpdateNetworkACLRule(options)
	return
}

/**
 * Security Groups
 *
 */

// ListSecurityGroups GET
// /security_groups
// List all security groups
func ListSecurityGroups(vpcService *vpcbetav1.VpcbetaV1) (securityGroups *vpcbetav1.SecurityGroupCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListSecurityGroupsOptions{}
	securityGroups, response, err = vpcService.ListSecurityGroups(options)
	return
}

// CreateSecurityGroup POST
// /security_groups
// Create a security group
func CreateSecurityGroup(vpcService *vpcbetav1.VpcbetaV1, name, vpcID string) (securityGroup *vpcbetav1.SecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateSecurityGroupOptions{}
	options.SetVPC(&vpcbetav1.VPCIdentity{
		ID: &vpcID,
	})
	options.SetName(name)
	securityGroup, response, err = vpcService.CreateSecurityGroup(options)
	return
}

// DeleteSecurityGroup DELETE
// /security_groups/{id}
// Delete a security group
func DeleteSecurityGroup(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteSecurityGroupOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteSecurityGroup(options)
	return response, err
}

// GetSecurityGroup GET
// /security_groups/{id}
// Retrieve a security group
func GetSecurityGroup(vpcService *vpcbetav1.VpcbetaV1, id string) (securityGroup *vpcbetav1.SecurityGroup, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetSecurityGroupOptions{}
	options.SetID(id)
	securityGroup, response, err = vpcService.GetSecurityGroup(options)
	return
}

// UpdateSecurityGroup PATCH
// /security_groups/{id}
// Update a security group
func UpdateSecurityGroup(vpcService *vpcbetav1.VpcbetaV1, id, name string) (securityGroup *vpcbetav1.SecurityGroup, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.SecurityGroupPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateSecurityGroupOptions{
		SecurityGroupPatch: patchBody,
		ID:                 &id,
	}
	securityGroup, response, err = vpcService.UpdateSecurityGroup(options)
	return
}

// ListSecurityGroupTarget GET
// /security_groups/{security_group_id}/targets
// List all the targets of a security group
func ListSecurityGroupTargets(vpcService *vpcbetav1.VpcbetaV1, id string) (targets *vpcbetav1.SecurityGroupTargetCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListSecurityGroupTargetsOptions{}
	options.SetSecurityGroupID(id)
	targets, response, err = vpcService.ListSecurityGroupTargets(options)
	return
}

// CreateSecurityGroupTarget POST
// /security_groups/{security_group_id}/targets/{id}
// Create a security group target
func CreateSecurityGroupTarget(vpcService *vpcbetav1.VpcbetaV1, sgID string, targetID string) (sg vpcbetav1.SecurityGroupTargetReferenceIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateSecurityGroupTargetBindingOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(targetID)
	sg, response, err = vpcService.CreateSecurityGroupTargetBinding(options)
	return
}

// DeleteSecurityGroupRule DELETE
// /security_groups/{security_group_id}/targets/{id}
// Delete a security group target
func DeleteSecurityGroupTarget(vpcService *vpcbetav1.VpcbetaV1, sgID, targetID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteSecurityGroupTargetBindingOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(targetID)
	response, err = vpcService.DeleteSecurityGroupTargetBinding(options)
	return response, err
}

// GetSecurityGroupRule GET
// /security_groups/{security_group_id}/targets/{id}
// Retrieve a security group target
func GetSecurityGroupTarget(vpcService *vpcbetav1.VpcbetaV1, sgID, targetID string) (sgTarget vpcbetav1.SecurityGroupTargetReferenceIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetSecurityGroupTargetOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(targetID)
	sgTarget, response, err = vpcService.GetSecurityGroupTarget(options)
	return
}

// ListSecurityGroupRules GET
// /security_groups/{security_group_id}/rules
// List all the rules of a security group
func ListSecurityGroupRules(vpcService *vpcbetav1.VpcbetaV1, id string) (rules *vpcbetav1.SecurityGroupRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListSecurityGroupRulesOptions{}
	options.SetSecurityGroupID(id)
	rules, response, err = vpcService.ListSecurityGroupRules(options)
	return
}

// CreateSecurityGroupRule POST
// /security_groups/{security_group_id}/rules
// Create a security group rule
func CreateSecurityGroupRule(vpcService *vpcbetav1.VpcbetaV1, sgID string) (rule vpcbetav1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetSecurityGroupRulePrototype(&vpcbetav1.SecurityGroupRulePrototype{
		Direction: core.StringPtr("inbound"),
		Protocol:  core.StringPtr("all"),
		IPVersion: core.StringPtr("ipv4"),
	})
	rule, response, err = vpcService.CreateSecurityGroupRule(options)
	return
}

// DeleteSecurityGroupRule DELETE
// /security_groups/{security_group_id}/rules/{id}
// Delete a security group rule
func DeleteSecurityGroupRule(vpcService *vpcbetav1.VpcbetaV1, sgID, sgRuleID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	response, err = vpcService.DeleteSecurityGroupRule(options)
	return response, err
}

// GetSecurityGroupRule GET
// /security_groups/{security_group_id}/rules/{id}
// Retrieve a security group rule
func GetSecurityGroupRule(vpcService *vpcbetav1.VpcbetaV1, sgID, sgRuleID string) (rule vpcbetav1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetSecurityGroupRuleOptions{}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	rule, response, err = vpcService.GetSecurityGroupRule(options)
	return
}

// UpdateSecurityGroupRule PATCH
// /security_groups/{security_group_id}/rules/{id}
// Update a security group rule
func UpdateSecurityGroupRule(vpcService *vpcbetav1.VpcbetaV1, sgID, sgRuleID string) (rule vpcbetav1.SecurityGroupRuleIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.SecurityGroupRulePatch{
		Remote: &vpcbetav1.SecurityGroupRuleRemotePatch{
			Address: core.StringPtr("1.1.1.11"),
		},
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateSecurityGroupRuleOptions{
		SecurityGroupRulePatch: patchBody,
	}
	options.SetSecurityGroupID(sgID)
	options.SetID(sgRuleID)
	rule, response, err = vpcService.UpdateSecurityGroupRule(options)
	return
}

/**
 * Load Balancers
 *
 */
func ListLoadBalancerProfiles(vpcService *vpcbetav1.VpcbetaV1) (profiles *vpcbetav1.LoadBalancerProfileCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancerProfilesOptions{}
	profiles, response, err = vpcService.ListLoadBalancerProfiles(options)
	return
}

func GetLoadBalancerProfile(vpcService *vpcbetav1.VpcbetaV1, name string) (profile *vpcbetav1.LoadBalancerProfile, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerProfileOptions{}
	options.SetName(name)
	profile, response, err = vpcService.GetLoadBalancerProfile(options)
	return
}

// ListLoadBalancers GET
// /load_balancers
// List all load balancers
func ListLoadBalancers(vpcService *vpcbetav1.VpcbetaV1) (lbs *vpcbetav1.LoadBalancerCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancersOptions{}
	lbs, response, err = vpcService.ListLoadBalancers(options)
	return
}

// CreateLoadBalancer POST
// /load_balancers
// Create and provision a load balancer
func CreateLoadBalancer(vpcService *vpcbetav1.VpcbetaV1, name, subnetID string) (lb *vpcbetav1.LoadBalancer, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateLoadBalancerOptions{}
	options.SetIsPublic(true)
	options.SetName(name)
	var subnetArray = []vpcbetav1.SubnetIdentityIntf{
		&vpcbetav1.SubnetIdentity{
			ID: &subnetID,
		},
	}
	options.SetSubnets(subnetArray)
	options.SetProfile(&vpcbetav1.LoadBalancerProfileIdentityByName{Name: core.StringPtr("network_small")})
	lb, response, err = vpcService.CreateLoadBalancer(options)
	return
}

// DeleteLoadBalancer DELETE
// /load_balancers/{id}
// Delete a load balancer
func DeleteLoadBalancer(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	deleteVpcOptions := &vpcbetav1.DeleteLoadBalancerOptions{}
	deleteVpcOptions.SetID(id)
	response, err = vpcService.DeleteLoadBalancer(deleteVpcOptions)
	return response, err
}

// GetLoadBalancer GET
// /load_balancers/{id}
// Retrieve a load balancer
func GetLoadBalancer(vpcService *vpcbetav1.VpcbetaV1, id string) (lb *vpcbetav1.LoadBalancer, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerOptions{}
	options.SetID(id)
	lb, response, err = vpcService.GetLoadBalancer(options)
	return
}

// UpdateLoadBalancer PATCH
// /load_balancers/{id}
// Update a load balancer
func UpdateLoadBalancer(vpcService *vpcbetav1.VpcbetaV1, id, name string) (lb *vpcbetav1.LoadBalancer, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.AddressPrefixPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateLoadBalancerOptions{
		LoadBalancerPatch: patchBody,
		ID:                &id,
	}
	lb, response, err = vpcService.UpdateLoadBalancer(options)
	return
}

// GetLoadBalancerStatistics GET
// /load_balancers/{id}/statistics
// List statistics of a load balancer
func GetLoadBalancerStatistics(vpcService *vpcbetav1.VpcbetaV1, id string) (lb *vpcbetav1.LoadBalancerStatistics, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerStatisticsOptions{}
	options.SetID(id)
	lb, response, err = vpcService.GetLoadBalancerStatistics(options)
	return
}

// ListLoadBalancerListeners GET
// /load_balancers/{load_balancer_id}/listeners
// List all listeners of the load balancer
func ListLoadBalancerListeners(vpcService *vpcbetav1.VpcbetaV1, id string) (listeners *vpcbetav1.LoadBalancerListenerCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancerListenersOptions{}
	options.SetLoadBalancerID(id)
	listeners, response, err = vpcService.ListLoadBalancerListeners(options)
	return
}

// CreateLoadBalancerListener POST
// /load_balancers/{load_balancer_id}/listeners
// Create a listener
func CreateLoadBalancerListener(vpcService *vpcbetav1.VpcbetaV1, lbID string) (listener *vpcbetav1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPort(rand.Int63n(100))
	options.SetProtocol("http")
	listener, response, err = vpcService.CreateLoadBalancerListener(options)
	return
}

// DeleteLoadBalancerListener DELETE
// /load_balancers/{load_balancer_id}/listeners/{id}
// Delete a listener
func DeleteLoadBalancerListener(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	response, err = vpcService.DeleteLoadBalancerListener(options)
	return response, err
}

// GetLoadBalancerListener GET
// /load_balancers/{load_balancer_id}/listeners/{id}
// Retrieve a listener
func GetLoadBalancerListener(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID string) (listener *vpcbetav1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerListenerOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(listenerID)
	listener, response, err = vpcService.GetLoadBalancerListener(options)
	return
}

// UpdateLoadBalancerListener PATCH
// /load_balancers/{load_balancer_id}/listeners/{id}
// Update a listener
func UpdateLoadBalancerListener(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID string) (listener *vpcbetav1.LoadBalancerListener, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.LoadBalancerListenerPatch{
		Protocol: core.StringPtr("http"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateLoadBalancerListenerOptions{
		LoadBalancerListenerPatch: patchBody,
		LoadBalancerID:            &lbID,
		ID:                        &listenerID,
	}
	listener, response, err = vpcService.UpdateLoadBalancerListener(options)
	return
}

// ListLoadBalancerListenerPolicies GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
// List all policies of the load balancer listener
func ListLoadBalancerListenerPolicies(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID string) (policies *vpcbetav1.LoadBalancerListenerPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancerListenerPoliciesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	policies, response, err = vpcService.ListLoadBalancerListenerPolicies(options)
	return
}

// CreateLoadBalancerListenerPolicy POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies
func CreateLoadBalancerListenerPolicy(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID string) (policy *vpcbetav1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPriority(2)
	options.SetAction("reject")
	policy, response, err = vpcService.CreateLoadBalancerListenerPolicy(options)
	return
}

// DeleteLoadBalancerListenerPolicy DELETE
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Delete a policy of the load balancer listener
func DeleteLoadBalancerListenerPolicy(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	response, err = vpcService.DeleteLoadBalancerListenerPolicy(options)
	return response, err
}

// GetLoadBalancerListenerPolicy GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Retrieve a policy of the load balancer listener
func GetLoadBalancerListenerPolicy(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID string) (policy *vpcbetav1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerListenerPolicyOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)
	policy, response, err = vpcService.GetLoadBalancerListenerPolicy(options)
	return
}

// UpdateLoadBalancerListenerPolicy PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{id}
// Update a policy of the load balancer listener
func UpdateLoadBalancerListenerPolicy(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID, targetPoolID string) (policy *vpcbetav1.LoadBalancerListenerPolicy, response *core.DetailedResponse, err error) {
	target := &vpcbetav1.LoadBalancerListenerPolicyTargetPatch{
		ID: &targetPoolID,
	}
	body := &vpcbetav1.LoadBalancerListenerPolicyPatch{
		Name:     core.StringPtr("my-loadblanacer-listner-policy"),
		Priority: core.Int64Ptr(4),
		Target:   target,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateLoadBalancerListenerPolicyOptions{
		LoadBalancerListenerPolicyPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetID(policyID)

	policy, response, err = vpcService.UpdateLoadBalancerListenerPolicy(options)
	return
}

// ListLoadBalancerListenerPolicyRules GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// List all rules of the load balancer listener policy
func ListLoadBalancerListenerPolicyRules(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID string) (rules *vpcbetav1.LoadBalancerListenerPolicyRuleCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancerListenerPolicyRulesOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	rules, response, err = vpcService.ListLoadBalancerListenerPolicyRules(options)
	return
}

// CreateLoadBalancerListenerPolicyRule POST
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules
// Create a rule for the load balancer listener policy
func CreateLoadBalancerListenerPolicyRule(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID string) (rule *vpcbetav1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetCondition("contains")
	options.SetType("hostname")
	options.SetValue("one")
	rule, response, err = vpcService.CreateLoadBalancerListenerPolicyRule(options)
	return
}

// DeleteLoadBalancerListenerPolicyRule DELETE
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Delete a rule from the load balancer listener policy
func DeleteLoadBalancerListenerPolicyRule(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID, ruleID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	response, err = vpcService.DeleteLoadBalancerListenerPolicyRule(options)
	return response, err
}

// GetLoadBalancerListenerPolicyRule GET
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Retrieve a rule of the load balancer listener policy
func GetLoadBalancerListenerPolicyRule(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID, ruleID string) (rule *vpcbetav1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerListenerPolicyRuleOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	rule, response, err = vpcService.GetLoadBalancerListenerPolicyRule(options)
	return
}

// UpdateLoadBalancerListenerPolicyRule PATCH
// /load_balancers/{load_balancer_id}/listeners/{listener_id}/policies/{policy_id}/rules/{id}
// Update a rule of the load balancer listener policy
func UpdateLoadBalancerListenerPolicyRule(vpcService *vpcbetav1.VpcbetaV1, lbID, listenerID, policyID, ruleID string) (rule *vpcbetav1.LoadBalancerListenerPolicyRule, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.LoadBalancerListenerPolicyRulePatch{
		Condition: core.StringPtr("equals"),
		Type:      core.StringPtr("header"),
		Value:     core.StringPtr("1"),
		Field:     core.StringPtr("field-1"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateLoadBalancerListenerPolicyRuleOptions{
		LoadBalancerListenerPolicyRulePatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetListenerID(listenerID)
	options.SetPolicyID(policyID)
	options.SetID(ruleID)
	rule, response, err = vpcService.UpdateLoadBalancerListenerPolicyRule(options)
	return
}

// ListLoadBalancerPools GET
// /load_balancers/{load_balancer_id}/pools
// List all pools of the load balancer
func ListLoadBalancerPools(vpcService *vpcbetav1.VpcbetaV1, id string) (pools *vpcbetav1.LoadBalancerPoolCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancerPoolsOptions{}
	options.SetLoadBalancerID(id)
	pools, response, err = vpcService.ListLoadBalancerPools(options)
	return
}

// CreateLoadBalancerPool POST
// /load_balancers/{load_balancer_id}/pools
// Create a load balancer pool
func CreateLoadBalancerPool(vpcService *vpcbetav1.VpcbetaV1, lbID, name string) (pool *vpcbetav1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetAlgorithm("round_robin")
	options.SetHealthMonitor(&vpcbetav1.LoadBalancerPoolHealthMonitorPrototype{
		Delay:      core.Int64Ptr(5),
		MaxRetries: core.Int64Ptr(2),
		Timeout:    core.Int64Ptr(4),
		Type:       core.StringPtr("http"),
	})
	options.SetName(name)
	options.SetProtocol("http")
	pool, response, err = vpcService.CreateLoadBalancerPool(options)
	return
}

// DeleteLoadBalancerPool DELETE
// /load_balancers/{load_balancer_id}/pools/{id}
// Delete a pool
func DeleteLoadBalancerPool(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	response, err = vpcService.DeleteLoadBalancerPool(options)
	return response, err
}

// GetLoadBalancerPool GET
// /load_balancers/{load_balancer_id}/pools/{id}
// Retrieve a load balancer pool
func GetLoadBalancerPool(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID string) (pool *vpcbetav1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerPoolOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	pool, response, err = vpcService.GetLoadBalancerPool(options)
	return
}

// UpdateLoadBalancerPool PATCH
// /load_balancers/{load_balancer_id}/pools/{id}
// Update a load balancer pool
func UpdateLoadBalancerPool(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID string) (pool *vpcbetav1.LoadBalancerPool, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.LoadBalancerPoolPatch{
		Protocol: core.StringPtr("tcp"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateLoadBalancerPoolOptions{
		LoadBalancerPoolPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetID(poolID)
	pool, response, err = vpcService.UpdateLoadBalancerPool(options)
	return
}

// ListLoadBalancerPoolMembers GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// List all members of the load balancer pool
func ListLoadBalancerPoolMembers(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID string) (members *vpcbetav1.LoadBalancerPoolMemberCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	members, response, err = vpcService.ListLoadBalancerPoolMembers(options)
	return
}

// CreateLoadBalancerPoolMember POST
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Create a member in the load balancer pool
func CreateLoadBalancerPoolMember(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID string) (member *vpcbetav1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetPort(1234)
	options.SetTarget(&vpcbetav1.LoadBalancerPoolMemberTargetPrototype{
		Address: core.StringPtr("12.12.0.0"),
	})
	member, response, err = vpcService.CreateLoadBalancerPoolMember(options)
	return
}

// UpdateLoadBalancerPoolMembers PUT
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members
// Update members of the load balancer pool
func UpdateLoadBalancerPoolMembers(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID string) (member *vpcbetav1.LoadBalancerPoolMemberCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ReplaceLoadBalancerPoolMembersOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetMembers([]vpcbetav1.LoadBalancerPoolMemberPrototype{
		{
			Port: core.Int64Ptr(2345),
			Target: &vpcbetav1.LoadBalancerPoolMemberTargetPrototype{
				Address: core.StringPtr("13.13.0.0"),
			},
		},
	})
	member, response, err = vpcService.ReplaceLoadBalancerPoolMembers(options)
	return
}

// DeleteLoadBalancerPoolMember DELETE
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Delete a member from the load balancer pool
func DeleteLoadBalancerPoolMember(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID, memberID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	response, err = vpcService.DeleteLoadBalancerPoolMember(options)
	return response, err
}

// GetLoadBalancerPoolMember GET
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
// Retrieve a member in the load balancer pool
func GetLoadBalancerPoolMember(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID, memberID string) (member *vpcbetav1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetLoadBalancerPoolMemberOptions{}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	member, response, err = vpcService.GetLoadBalancerPoolMember(options)
	return
}

// UpdateLoadBalancerPoolMember PATCH
// /load_balancers/{load_balancer_id}/pools/{pool_id}/members/{id}
func UpdateLoadBalancerPoolMember(vpcService *vpcbetav1.VpcbetaV1, lbID, poolID, memberID string) (member *vpcbetav1.LoadBalancerPoolMember, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.LoadBalancerPoolMemberPatch{
		Port: core.Int64Ptr(3434),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateLoadBalancerPoolMemberOptions{
		LoadBalancerPoolMemberPatch: patchBody,
	}
	options.SetLoadBalancerID(lbID)
	options.SetPoolID(poolID)
	options.SetID(memberID)
	member, response, err = vpcService.UpdateLoadBalancerPoolMember(options)
	return
}

/**
 * VPN
 *
 */

// ListIkePolicies GET
// /ike_policies
// List all IKE policies
func ListIkePolicies(vpcService *vpcbetav1.VpcbetaV1) (ikePolicies *vpcbetav1.IkePolicyCollection, response *core.DetailedResponse, err error) {
	options := vpcService.NewListIkePoliciesOptions()
	ikePolicies, response, err = vpcService.ListIkePolicies(options)
	return
}

// CreateIkePolicy POST
// /ike_policies
// Create an IKE policy
func CreateIkePolicy(vpcService *vpcbetav1.VpcbetaV1, name string) (ikePolicy *vpcbetav1.IkePolicy, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateIkePolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("sha256")
	options.SetDhGroup(14)
	options.SetEncryptionAlgorithm("aes128")
	options.SetIkeVersion(1)
	ikePolicy, response, err = vpcService.CreateIkePolicy(options)
	return
}

// DeleteIkePolicy DELETE
// /ike_policies/{id}
// Delete an IKE policy
func DeleteIkePolicy(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteIkePolicyOptions(id)
	response, err = vpcService.DeleteIkePolicy(options)
	return response, err
}

// GetIkePolicy GET
// /ike_policies/{id}
// Retrieve the specified IKE policy
func GetIkePolicy(vpcService *vpcbetav1.VpcbetaV1, id string) (ikePolicy *vpcbetav1.IkePolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetIkePolicyOptions(id)
	ikePolicy, response, err = vpcService.GetIkePolicy(options)
	return
}

// UpdateIkePolicy PATCH
// /ike_policies/{id}
// Update an IKE policy
func UpdateIkePolicy(vpcService *vpcbetav1.VpcbetaV1, id string) (ikePolicy *vpcbetav1.IkePolicy, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.IkePolicyPatch{
		Name:    core.StringPtr("go-ike-policy-2"),
		DhGroup: core.Int64Ptr(15),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateIkePolicyOptions{
		ID:             &id,
		IkePolicyPatch: patchBody,
	}
	ikePolicy, response, err = vpcService.UpdateIkePolicy(options)
	return
}

// ListVPNGatewayIkePolicyConnections GET
// /ike_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListVPNGatewayIkePolicyConnections(vpcService *vpcbetav1.VpcbetaV1, id string) (connections *vpcbetav1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListIkePolicyConnectionsOptions{
		ID: &id,
	}
	connections, response, err = vpcService.ListIkePolicyConnections(options)
	return
}

// ListIpsecPolicies GET
// /ipsec_policies
// List all IPsec policies
func ListIpsecPolicies(vpcService *vpcbetav1.VpcbetaV1) (ipsecPolicies *vpcbetav1.IPsecPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListIpsecPoliciesOptions{}
	ipsecPolicies, response, err = vpcService.ListIpsecPolicies(options)
	return
}

// CreateIpsecPolicy POST
// /ipsec_policies
// Create an IPsec policy
func CreateIpsecPolicy(vpcService *vpcbetav1.VpcbetaV1, name string) (ipsecPolicy *vpcbetav1.IPsecPolicy, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateIpsecPolicyOptions{}
	options.SetName(name)
	options.SetAuthenticationAlgorithm("sha256")
	options.SetEncryptionAlgorithm("aes128")
	options.SetPfs("disabled")
	ipsecPolicy, response, err = vpcService.CreateIpsecPolicy(options)
	return
}

// DeleteIpsecPolicy DELETE
// /ipsec_policies/{id}
// Delete an IPsec policy
func DeleteIpsecPolicy(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteIpsecPolicyOptions(id)
	response, err = vpcService.DeleteIpsecPolicy(options)
	return response, err
}

// GetIpsecPolicy GET
// /ipsec_policies/{id}
// Retrieve the specified IPsec policy
func GetIpsecPolicy(vpcService *vpcbetav1.VpcbetaV1, id string) (ipsecPolicy *vpcbetav1.IPsecPolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetIpsecPolicyOptions(id)
	ipsecPolicy, response, err = vpcService.GetIpsecPolicy(options)
	return
}

// UpdateIpsecPolicy PATCH
// /ipsec_policies/{id}
// Update an IPsec policy
func UpdateIpsecPolicy(vpcService *vpcbetav1.VpcbetaV1, id string) (ipsecPolicy *vpcbetav1.IPsecPolicy, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.IPsecPolicyPatch{
		EncryptionAlgorithm: core.StringPtr("3des"),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateIpsecPolicyOptions{
		ID:               &id,
		IPsecPolicyPatch: patchBody,
	}
	ipsecPolicy, response, err = vpcService.UpdateIpsecPolicy(options)
	return
}

// ListVPNGatewayIpsecPolicyConnections GET
// /ipsec_policies/{id}/connections
// Lists all the connections that use the specified policy
func ListIpsecPolicyConnections(vpcService *vpcbetav1.VpcbetaV1, id string) (connections *vpcbetav1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListIpsecPolicyConnectionsOptions{
		ID: &id,
	}
	connections, response, err = vpcService.ListIpsecPolicyConnections(options)
	return
}

func ListShareProfiles(vpcService *vpcbetav1.VpcbetaV1) (result *vpcbetav1.ShareProfileCollection, response *core.DetailedResponse, err error) {
	listShareProfilesOptions := &vpcbetav1.ListShareProfilesOptions{}

	result, response, err = vpcService.ListShareProfiles(listShareProfilesOptions)
	return
}

func GetShareProfile(vpcService *vpcbetav1.VpcbetaV1, name string) (result *vpcbetav1.ShareProfile, response *core.DetailedResponse, err error) {
	getShareProfileOptions := &vpcbetav1.GetShareProfileOptions{
		Name: &name,
	}

	result, response, err = vpcService.GetShareProfile(getShareProfileOptions)
	return
}

func ListShares(vpcService *vpcbetav1.VpcbetaV1) (result *vpcbetav1.ShareCollection, response *core.DetailedResponse, err error) {
	listSharesOptions := &vpcbetav1.ListSharesOptions{}

	result, response, err = vpcService.ListShares(listSharesOptions)

	return
}

func CreateShare(vpcService *vpcbetav1.VpcbetaV1, createdVPCID, name *string) (result *vpcbetav1.Share, response *core.DetailedResponse, err error) {
	shareProfileIdentityModel := &vpcbetav1.ShareProfileIdentityByName{
		Name: core.StringPtr("tier-3iops"),
	}

	vpcIdentityModel := &vpcbetav1.VPCIdentityByID{
		ID: createdVPCID,
	}

	shareTargetPrototypeModel := &vpcbetav1.ShareMountTargetPrototype{
		Name: core.StringPtr("my-share-target" + timestamp),
		VPC:  vpcIdentityModel,
	}

	zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
		Name: core.StringPtr("us-south-1"),
	}

	sharePrototypeShareContextModel := &vpcbetav1.SharePrototypeShareContext{
		Iops:                core.Int64Ptr(int64(100)),
		Name:                core.StringPtr("my-share" + timestamp),
		Profile:             shareProfileIdentityModel,
		ReplicationCronSpec: core.StringPtr("0 */5 * * *"),
		Targets:             []vpcbetav1.ShareMountTargetPrototype{*shareTargetPrototypeModel},
		UserTags:            []string{"testString"},
		Zone:                zoneIdentityModel,
	}

	encryptionKeyIdentityModel := &vpcbetav1.EncryptionKeyIdentityByCRN{
		CRN: core.StringPtr("crn:v1:bluemix:public:kms:us-south:a/dffc98a0f1f0f95f6613b3b752286b87:e4a29d1a-2ef0-42a6-8fd2-350deb1c647e:key:5437653b-c4b1-447f-9646-b2a2a4cd6179"),
	}

	shareInitialOwnerModel := &vpcbetav1.ShareInitialOwner{
		Gid: core.Int64Ptr(int64(50)),
		Uid: core.Int64Ptr(int64(50)),
	}

	sharePrototypeModel := &vpcbetav1.SharePrototypeShareBySize{
		Iops:          core.Int64Ptr(int64(100)),
		Name:          name,
		Profile:       shareProfileIdentityModel,
		ReplicaShare:  sharePrototypeShareContextModel,
		Targets:       []vpcbetav1.ShareMountTargetPrototype{*shareTargetPrototypeModel},
		UserTags:      []string{"my-share-tag"},
		Zone:          zoneIdentityModel,
		EncryptionKey: encryptionKeyIdentityModel,
		InitialOwner:  shareInitialOwnerModel,
		Size:          core.Int64Ptr(int64(200)),
	}

	createShareOptions := &vpcbetav1.CreateShareOptions{
		SharePrototype: sharePrototypeModel,
	}

	result, response, err = vpcService.CreateShare(createShareOptions)
	return
}

func GetShare(vpcService *vpcbetav1.VpcbetaV1, shareID *string) (result *vpcbetav1.Share, response *core.DetailedResponse, err error) {
	getShareOptions := &vpcbetav1.GetShareOptions{
		ID: shareID,
	}

	result, response, err = vpcService.GetShare(getShareOptions)
	return
}

func UpdateShare(vpcService *vpcbetav1.VpcbetaV1, shareID, etag, name *string) (result *vpcbetav1.Share, response *core.DetailedResponse, err error) {

	sharePatchModel := &vpcbetav1.SharePatch{
		Name: name,
	}
	sharePatchModelAsPatch, _ := sharePatchModel.AsPatch()

	updateShareOptions := &vpcbetav1.UpdateShareOptions{
		ID:         shareID,
		SharePatch: sharePatchModelAsPatch,
		IfMatch:    etag,
	}

	result, response, err = vpcService.UpdateShare(updateShareOptions)
	return
}

func GetShareSource(vpcService *vpcbetav1.VpcbetaV1, shareID *string) (result *vpcbetav1.Share, response *core.DetailedResponse, err error) {
	getShareSourceOptions := &vpcbetav1.GetShareSourceOptions{
		ShareID: shareID,
	}

	result, response, err = vpcService.GetShareSource(getShareSourceOptions)
	return
}

func FailoverShare(vpcService *vpcbetav1.VpcbetaV1, shareID *string) (response *core.DetailedResponse, err error) {
	failoverShareOptions := &vpcbetav1.FailoverShareOptions{
		ShareID:        shareID,
		FallbackPolicy: core.StringPtr("split"),
		Timeout:        core.Int64Ptr(int64(600)),
	}

	response, err = vpcService.FailoverShare(failoverShareOptions)
	return
}

func ListShareTargets(vpcService *vpcbetav1.VpcbetaV1, shareID *string) (result *vpcbetav1.ShareMountTargetCollection, response *core.DetailedResponse, err error) {
	listShareTargetsOptions := &vpcbetav1.ListShareMountTargetsOptions{
		ShareID: shareID,
	}

	result, response, err = vpcService.ListShareMountTargets(listShareTargetsOptions)
	return
}

func CreateShareTarget(vpcService *vpcbetav1.VpcbetaV1, vpcID, shareID, name *string) (result *vpcbetav1.ShareMountTarget, response *core.DetailedResponse, err error) {
	vpcIdentityModel := &vpcbetav1.VPCIdentityByID{
		ID: vpcID,
	}

	createShareTargetOptions := &vpcbetav1.CreateShareMountTargetOptions{
		ShareID: shareID,
		VPC:     vpcIdentityModel,
		Name:    core.StringPtr("my-share-target" + timestamp),
	}

	result, response, err = vpcService.CreateShareMountTarget(createShareTargetOptions)
	return
}

func GetShareTarget(vpcService *vpcbetav1.VpcbetaV1, shareID, targetID *string) (result *vpcbetav1.ShareMountTarget, response *core.DetailedResponse, err error) {
	getShareTargetOptions := &vpcbetav1.GetShareMountTargetOptions{
		ShareID: shareID,
		ID:      targetID,
	}

	result, response, err = vpcService.GetShareMountTarget(getShareTargetOptions)
	return
}

func UpdateShareTarget(vpcService *vpcbetav1.VpcbetaV1, shareID, targetID, name *string) (result *vpcbetav1.ShareMountTarget, response *core.DetailedResponse, err error) {
	shareTargetPatchModel := &vpcbetav1.ShareMountTargetPatch{
		Name: name,
	}
	shareTargetPatchModelAsPatch, _ := shareTargetPatchModel.AsPatch()

	updateShareTargetOptions := &vpcbetav1.UpdateShareMountTargetOptions{
		ShareID:               shareID,
		ID:                    targetID,
		ShareMountTargetPatch: shareTargetPatchModelAsPatch,
	}

	result, response, err = vpcService.UpdateShareMountTarget(updateShareTargetOptions)
	return
}

func DeleteShareTarget(vpcService *vpcbetav1.VpcbetaV1, shareID, targetID *string) (result *vpcbetav1.ShareMountTarget, response *core.DetailedResponse, err error) {
	deleteShareTargetOptions := &vpcbetav1.DeleteShareMountTargetOptions{
		ShareID: shareID,
		ID:      targetID,
	}

	result, response, err = vpcService.DeleteShareMountTarget(deleteShareTargetOptions)
	return
}

func DeleteShareSource(vpcService *vpcbetav1.VpcbetaV1, shareID *string) (response *core.DetailedResponse, err error) {
	deleteShareSourceOptions := &vpcbetav1.DeleteShareSourceOptions{
		ShareID: shareID,
	}

	response, err = vpcService.DeleteShareSource(deleteShareSourceOptions)
	return
}

func DeleteShare(vpcService *vpcbetav1.VpcbetaV1, shareID, etag *string) (result *vpcbetav1.Share, response *core.DetailedResponse, err error) {
	deleteShareOptions := &vpcbetav1.DeleteShareOptions{
		ID:      shareID,
		IfMatch: etag,
	}

	result, response, err = vpcService.DeleteShare(deleteShareOptions)
	return
}

// ListVPNGateways GET
// /VPN_gateways
// List all VPN gateways
func ListVPNGateways(vpcService *vpcbetav1.VpcbetaV1) (gateways *vpcbetav1.VPNGatewayCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVPNGatewaysOptions{}
	gateways, response, err = vpcService.ListVPNGateways(options)
	return
}

// CreateVPNGateway POST
// /VPN_gateways
// Create a VPN gateway
func CreateVPNGateway(vpcService *vpcbetav1.VpcbetaV1, subnetID, name string) (gateway vpcbetav1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateVPNGatewayOptions{
		VPNGatewayPrototype: &vpcbetav1.VPNGatewayPrototype{
			Name: &name,
			Subnet: &vpcbetav1.SubnetIdentity{
				ID: &subnetID,
			},
		},
	}
	gateway, response, err = vpcService.CreateVPNGateway(options)
	return
}

// DeleteVPNGateway DELETE
// /VPN_gateways/{id}
// Delete a VPN gateway
func DeleteVPNGateway(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteVPNGatewayOptions(id)
	response, err = vpcService.DeleteVPNGateway(options)
	return response, err
}

// GetVPNGateway GET
// /VPN_gateways/{id}
// Retrieve the specified VPN gateway
func GetVPNGateway(vpcService *vpcbetav1.VpcbetaV1, id string) (gateway vpcbetav1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetVPNGatewayOptions(id)
	gateway, response, err = vpcService.GetVPNGateway(options)
	return
}

// UpdateVPNGateway PATCH
// /VPN_gateways/{id}
// Update a VPN gateway
func UpdateVPNGateway(vpcService *vpcbetav1.VpcbetaV1, id, name string) (gateway vpcbetav1.VPNGatewayIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.VPNGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateVPNGatewayOptions{
		ID:              &id,
		VPNGatewayPatch: patchBody,
	}
	gateway, response, err = vpcService.UpdateVPNGateway(options)
	return
}

// ListVPNGatewayConnections GET
// /VPN_gateways/{VPN_gateway_id}/connections
// List all the connections of a VPN gateway
func ListVPNGatewayConnections(vpcService *vpcbetav1.VpcbetaV1, gatewayID string) (connections *vpcbetav1.VPNGatewayConnectionCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVPNGatewayConnectionsOptions{}
	options.SetVPNGatewayID(gatewayID)
	connections, response, err = vpcService.ListVPNGatewayConnections(options)
	return
}

// CreateVPNGatewayConnection POST
// /VPN_gateways/{VPN_gateway_id}/connections
// Create a VPN connection
func CreateVPNGatewayConnection(vpcService *vpcbetav1.VpcbetaV1, gatewayID, name string) (connections vpcbetav1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateVPNGatewayConnectionOptions{}
	peerAddress := "192.168.0.1"
	psk := "pre-shared-key"
	local := []string{"192.132.0.0/28"}
	peer := []string{"197.155.0.0/28"}
	options.SetVPNGatewayConnectionPrototype(&vpcbetav1.VPNGatewayConnectionPrototype{
		Name:        &name,
		PeerAddress: &peerAddress,
		Psk:         &psk,
		LocalCIDRs:  local,
		PeerCIDRs:   peer,
	})
	options.SetVPNGatewayID(gatewayID)
	connections, response, err = vpcService.CreateVPNGatewayConnection(options)
	return
}

// DeleteVPNGatewayConnection DELETE
// /VPN_gateways/{VPN_gateway_id}/connections/{id}
// Delete a VPN connection
func DeleteVPNGatewayConnection(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteVPNGatewayConnectionOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	response, err = vpcService.DeleteVPNGatewayConnection(options)
	return response, err
}

// GetVPNGatewayConnection GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}
// Retrieve the specified VPN connection
func GetVPNGatewayConnection(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID string) (connection vpcbetav1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetVPNGatewayConnectionOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	connection, response, err = vpcService.GetVPNGatewayConnection(options)
	return
}

// UpdateVPNGatewayConnection PATCH
// /VPN_gateways/{VPN_gateway_id}/connections/{id}
// Update a VPN connection
func UpdateVPNGatewayConnection(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, name string) (connection vpcbetav1.VPNGatewayConnectionIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.VPNGatewayConnectionPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateVPNGatewayConnectionOptions{
		ID:                        &connID,
		VPNGatewayID:              &gatewayID,
		VPNGatewayConnectionPatch: patchBody,
	}
	connection, response, err = vpcService.UpdateVPNGatewayConnection(options)
	return
}

// ListVPNGatewayConnectionLocalCIDRs GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs
// List all local CIDRs for a resource
func ListVPNGatewayConnectionLocalCIDRs(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID string) (localCIDRs *vpcbetav1.VPNGatewayConnectionLocalCIDRs, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVPNGatewayConnectionLocalCIDRsOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	localCIDRs, response, err = vpcService.ListVPNGatewayConnectionLocalCIDRs(options)
	return
}

// DeleteVPNGatewayConnectionLocalCIDR DELETE
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVPNGatewayConnectionLocalCIDR(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.RemoveVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.RemoveVPNGatewayConnectionLocalCIDR(options)
	return response, err
}

// GetVPNGatewayConnectionLocalCIDR GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs
// Check if a specific CIDR exists on a specific resource
func CheckVPNGatewayConnectionLocalCIDR(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CheckVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.CheckVPNGatewayConnectionLocalCIDR(options)
	return response, err
}

// SetVPNGatewayConnectionLocalCIDR - PUT
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/local_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVPNGatewayConnectionLocalCIDR(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.AddVPNGatewayConnectionLocalCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.AddVPNGatewayConnectionLocalCIDR(options)
	return response, err
}

// ListVPNGatewayConnectionPeerCIDRs GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs
// List all peer CIDRs for a resource
func ListVPNGatewayConnectionPeerCIDRs(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID string) (peerCIDRs *vpcbetav1.VPNGatewayConnectionPeerCIDRs, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListVPNGatewayConnectionPeerCIDRsOptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	peerCIDRs, response, err = vpcService.ListVPNGatewayConnectionPeerCIDRs(options)
	return
}

// DeleteVPNGatewayConnectionPeerCIDR DELETE
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Remove a CIDR from a resource
func DeleteVPNGatewayConnectionPeerCIDR(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.RemoveVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.RemoveVPNGatewayConnectionPeerCIDR(options)
	return response, err
}

// GetVPNGatewayConnectionPeerCIDR GET
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Check if a specific CIDR exists on a specific resource
func CheckVPNGatewayConnectionPeerCIDR(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CheckVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.CheckVPNGatewayConnectionPeerCIDR(options)
	return response, err
}

// SetVPNGatewayConnectionPeerCIDR - PUT
// /VPN_gateways/{VPN_gateway_id}/connections/{id}/peer_cidrs/{prefix_address}/{prefix_length}
// Set a CIDR on a resource
func SetVPNGatewayConnectionPeerCIDR(vpcService *vpcbetav1.VpcbetaV1, gatewayID, connID, prefixAdd, prefixLen string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.AddVPNGatewayConnectionPeerCIDROptions{}
	options.SetVPNGatewayID(gatewayID)
	options.SetID(connID)
	options.SetCIDRPrefix(prefixAdd)
	options.SetPrefixLength(prefixLen)
	response, err = vpcService.AddVPNGatewayConnectionPeerCIDR(options)
	return response, err
}

// Flow Logs
// ListFlowLogCollectors - GET
// /flow_log_collectors
// List all flow log collectors
func ListFlowLogCollectors(vpcService *vpcbetav1.VpcbetaV1) (flowLogs *vpcbetav1.FlowLogCollectorCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListFlowLogCollectorsOptions{}
	flowLogs, response, err = vpcService.ListFlowLogCollectors(options)
	return
}

// GetFlowLogCollector - GET
// /flow_log_collectors/{id}
// Retrieve the specified flow log collector
func GetFlowLogCollector(vpcService *vpcbetav1.VpcbetaV1, id string) (flowLog *vpcbetav1.FlowLogCollector, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetFlowLogCollectorOptions{}
	options.SetID(id)
	flowLog, response, err = vpcService.GetFlowLogCollector(options)
	return
}

// DeleteFlowLogCollector DELETE
// /flow_log_collectors/{id}
// Delete specified flow_log_collector
func DeleteFlowLogCollector(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteFlowLogCollectorOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteFlowLogCollector(options)
	return response, err
}

// UpdateFlowLogCollector PATCH
// /flow_log_collectors/{id}
// Update specified flow log collector
func UpdateFlowLogCollector(vpcService *vpcbetav1.VpcbetaV1, id, name string) (flowLog *vpcbetav1.FlowLogCollector, response *core.DetailedResponse, err error) {
	model := &vpcbetav1.FlowLogCollectorPatch{
		Name: &name,
	}
	patch, _ := model.AsPatch()
	options := &vpcbetav1.UpdateFlowLogCollectorOptions{
		FlowLogCollectorPatch: patch,
		ID:                    &id,
	}
	flowLog, response, err = vpcService.UpdateFlowLogCollector(options)
	return
}

func CreateFlowLogCollector(vpcService *vpcbetav1.VpcbetaV1, name, bucketName, vpcId string) (flowLog *vpcbetav1.FlowLogCollector, response *core.DetailedResponse, err error) {

	options := &vpcbetav1.CreateFlowLogCollectorOptions{}
	options.SetName(name)
	options.SetTarget(&vpcbetav1.FlowLogCollectorTargetPrototype{
		ID: &vpcId,
	})
	options.SetStorageBucket(&vpcbetav1.LegacyCloudObjectStorageBucketIdentity{
		Name: &bucketName,
	})
	flowLog, response, err = vpcService.CreateFlowLogCollector(options)
	return
}

/**
 * Autoscale
 *
 */
// GET
// /instance/templates
// Get instance templates.
func ListInstanceTemplates(vpcService *vpcbetav1.VpcbetaV1) (templates *vpcbetav1.InstanceTemplateCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceTemplatesOptions{}
	templates, response, err = vpcService.ListInstanceTemplates(options)
	return
}

// POST
// /instance/templates
// Create an instance template
func CreateInstanceTemplate(vpcService *vpcbetav1.VpcbetaV1, name, imageID, profileName, zoneName, subnetID, vpcID string) (template vpcbetav1.InstanceTemplateIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.CreateInstanceTemplateOptions{}
	options.SetInstanceTemplatePrototype(&vpcbetav1.InstanceTemplatePrototype{
		Name: &name,
		Image: &vpcbetav1.ImageIdentity{
			ID: &imageID,
		},
		Profile: &vpcbetav1.InstanceProfileIdentity{
			Name: &profileName,
		},
		Zone: &vpcbetav1.ZoneIdentity{
			Name: &zoneName,
		},
		PrimaryNetworkInterface: &vpcbetav1.NetworkInterfacePrototype{
			Subnet: &vpcbetav1.SubnetIdentity{
				ID: &subnetID,
			},
		},
		VPC: &vpcbetav1.VPCIdentity{
			ID: &vpcID,
		},
	})
	template, response, err = vpcService.CreateInstanceTemplate(options)
	return
}

// DELETE
// /instance/templates/{id}
// Delete specified instance template
func DeleteInstanceTemplate(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceTemplateOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstanceTemplate(options)
	return response, err
}

// GET
// /instance/templates/{id}
// Retrieve specified instance template
func GetInstanceTemplate(vpcService *vpcbetav1.VpcbetaV1, id string) (template vpcbetav1.InstanceTemplateIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceTemplateOptions{}
	options.SetID(id)
	template, response, err = vpcService.GetInstanceTemplate(options)
	return
}

// PATCH
// /instance/templates/{id}
// Update specified instance template
func UpdateInstanceTemplate(vpcService *vpcbetav1.VpcbetaV1, id, name string) (template vpcbetav1.InstanceTemplateIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.InstanceTemplatePatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceTemplateOptions{
		InstanceTemplatePatch: patchBody,
	}
	options.SetID(id)
	template, response, err = vpcService.UpdateInstanceTemplate(options)
	return
}

// GET
// /instance_groups
// List all instance groups
func ListInstanceGroups(vpcService *vpcbetav1.VpcbetaV1) (templates *vpcbetav1.InstanceGroupCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceGroupsOptions{}
	templates, response, err = vpcService.ListInstanceGroups(options)
	return
}

// POST
// /instance_groups
// Create an instance group
func CreateInstanceGroup(vpcService *vpcbetav1.VpcbetaV1, instanceID, name, subnetID string, membership int64) (template *vpcbetav1.InstanceGroup, response *core.DetailedResponse, err error) {

	options := &vpcbetav1.CreateInstanceGroupOptions{
		InstanceTemplate: &vpcbetav1.InstanceTemplateIdentityByID{
			ID: &instanceID,
		},
		Subnets: []vpcbetav1.SubnetIdentityIntf{
			&vpcbetav1.SubnetIdentity{
				ID: &subnetID,
			},
		},
		Name:            &name,
		MembershipCount: &membership,
	}
	template, response, err = vpcService.CreateInstanceGroup(options)
	return
}

// DELETE
// /instance_groups/{id}
// Delete specified instance group

func DeleteInstanceGroup(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupOptions{}
	options.SetID(id)
	response, err = vpcService.DeleteInstanceGroup(options)
	return response, err
}

// GET
// /instance_groups/{id}
// Retrieve specified instance group
func GetInstanceGroup(vpcService *vpcbetav1.VpcbetaV1, id string) (ig *vpcbetav1.InstanceGroup, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceGroupOptions{}
	options.SetID(id)
	ig, response, err = vpcService.GetInstanceGroup(options)
	return
}

// PATCH
// /instance_groups/{id}
// Update specified instance group
func UpdateInstanceGroup(vpcService *vpcbetav1.VpcbetaV1, id, name string) (template *vpcbetav1.InstanceGroup, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.InstanceGroupPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceGroupOptions{
		InstanceGroupPatch: patchBody,
	}
	options.SetID(id)
	template, response, err = vpcService.UpdateInstanceGroup(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/load_balancer
// Delete specified instance group load balancer
func DeleteInstanceGroupLoadBalancer(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupLoadBalancerOptions{}
	options.SetInstanceGroupID(id)
	response, err = vpcService.DeleteInstanceGroupLoadBalancer(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers
// List all managers for an instance group
func ListInstanceGroupManagers(vpcService *vpcbetav1.VpcbetaV1, id string) (templates *vpcbetav1.InstanceGroupManagerCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceGroupManagersOptions{}
	options.SetInstanceGroupID(id)
	templates, response, err = vpcService.ListInstanceGroupManagers(options)
	return
}

// POST
// /instance_groups/{instance_group_id}/managers
// Create an instance group manager
func CreateInstanceGroupManager(vpcService *vpcbetav1.VpcbetaV1, gID, name string) (manager vpcbetav1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {

	options := &vpcbetav1.CreateInstanceGroupManagerOptions{
		InstanceGroupManagerPrototype: &vpcbetav1.InstanceGroupManagerPrototype{
			Name:               &name,
			ManagerType:        core.StringPtr("autoscale"),
			MaxMembershipCount: core.Int64Ptr(2),
		},
	}
	options.SetInstanceGroupID(gID)
	manager, response, err = vpcService.CreateInstanceGroupManager(options)
	return
}

// POST
// /instance_groups/{instance_group_id}/managers
// Create an instance group manager
func CreateInstanceGroupManagerScheduled(vpcService *vpcbetav1.VpcbetaV1, gID, name string) (manager vpcbetav1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {

	options := &vpcbetav1.CreateInstanceGroupManagerOptions{
		InstanceGroupManagerPrototype: &vpcbetav1.InstanceGroupManagerPrototype{
			Name:        &name,
			ManagerType: core.StringPtr("scheduled"),
			// ManagementEnabled: true
		},
	}
	options.SetInstanceGroupID(gID)
	manager, response, err = vpcService.CreateInstanceGroupManager(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/managers/{id}
// Delete specified instance group manager

func DeleteInstanceGroupManager(vpcService *vpcbetav1.VpcbetaV1, gID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupManagerOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	response, err = vpcService.DeleteInstanceGroupManager(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers/{id}
// Retrieve specified instance group

func GetInstanceGroupManager(vpcService *vpcbetav1.VpcbetaV1, gID, id, name string) (manager vpcbetav1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceGroupManagerOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	manager, response, err = vpcService.GetInstanceGroupManager(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/managers/{id}
// Update specified instance group manager
func UpdateInstanceGroupManager(vpcService *vpcbetav1.VpcbetaV1, gID, id, name string) (manager vpcbetav1.InstanceGroupManagerIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.InstanceGroupManagerPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceGroupManagerOptions{
		InstanceGroupManagerPatch: patchBody,
	}
	options.SetInstanceGroupID(gID)
	options.SetID(id)
	manager, response, err = vpcService.UpdateInstanceGroupManager(options)
	return
}

// GET
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/actions
// List all actions for an instance group manager
func ListInstanceGroupManagerActions(vpcService *vpcbetav1.VpcbetaV1, gID, mID string) (actions *vpcbetav1.InstanceGroupManagerActionsCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceGroupManagerActionsOptions{}
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	actions, response, err = vpcService.ListInstanceGroupManagerActions(options)
	return
}

// POST
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/actions
// Create an instance group manager action
func CreateInstanceGroupManagerAction(vpcService *vpcbetav1.VpcbetaV1, gID, mID, name string, membershipCount int64) (actions vpcbetav1.InstanceGroupManagerActionIntf, response *core.DetailedResponse, err error) {
	instanceGroupManagerScheduledActionGroupPrototype := vpcbetav1.InstanceGroupManagerScheduledActionGroupPrototype{}
	instanceGroupManagerScheduledActionGroupPrototype.MembershipCount = &membershipCount
	options := &vpcbetav1.CreateInstanceGroupManagerActionOptions{
		InstanceGroupManagerActionPrototype: &vpcbetav1.InstanceGroupManagerActionPrototype{
			CronSpec: core.StringPtr("*/5 1,2,3 * * *"),
			Group:    &instanceGroupManagerScheduledActionGroupPrototype,
		},
	}
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	actions, response, err = vpcService.CreateInstanceGroupManagerAction(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/actions/{id}
// Delete specified instance group manager action

func DeleteInstanceGroupManagerAction(vpcService *vpcbetav1.VpcbetaV1, gID, mID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupManagerActionOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	response, err = vpcService.DeleteInstanceGroupManagerAction(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/actions/{id}
// Retrieve specified instance group manager action

func GetInstanceGroupManagerAction(vpcService *vpcbetav1.VpcbetaV1, gID, mID, id string) (action vpcbetav1.InstanceGroupManagerActionIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceGroupManagerActionOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	action, response, err = vpcService.GetInstanceGroupManagerAction(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/actions/{id}
// Update specified instance group manager action
func UpdateInstanceGroupManagerAction(vpcService *vpcbetav1.VpcbetaV1, igID, mID, id, name string) (action vpcbetav1.InstanceGroupManagerActionIntf, response *core.DetailedResponse, err error) {
	instanceGroupManagerActionPatchModel := &vpcbetav1.InstanceGroupManagerActionPatch{}
	instanceGroupManagerActionPatchModel.CronSpec = core.StringPtr("*/5 1,2,3 * * *")
	instanceGroupManagerActionPatch, _ := instanceGroupManagerActionPatchModel.AsPatch()
	options := &vpcbetav1.UpdateInstanceGroupManagerActionOptions{}
	options.SetInstanceGroupID(igID)
	options.SetInstanceGroupManagerID(mID)
	options.SetID(id)
	options.InstanceGroupManagerActionPatch = instanceGroupManagerActionPatch

	action, response, err = vpcService.UpdateInstanceGroupManagerAction(options)
	return
}

// GET
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies
// List all policies for an instance group manager
func ListInstanceGroupManagerPolicies(vpcService *vpcbetav1.VpcbetaV1, gID, mID string) (policies *vpcbetav1.InstanceGroupManagerPolicyCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceGroupManagerPoliciesOptions{}
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	policies, response, err = vpcService.ListInstanceGroupManagerPolicies(options)
	return
}

// POST
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies
// Create an instance group manager policy
func CreateInstanceGroupManagerPolicy(vpcService *vpcbetav1.VpcbetaV1, gID, mID, name string) (policy vpcbetav1.InstanceGroupManagerPolicyIntf, response *core.DetailedResponse, err error) {

	options := &vpcbetav1.CreateInstanceGroupManagerPolicyOptions{
		InstanceGroupManagerPolicyPrototype: &vpcbetav1.InstanceGroupManagerPolicyPrototype{
			Name:        &name,
			MetricType:  core.StringPtr("cpu"),
			MetricValue: core.Int64Ptr(50),
			PolicyType:  core.StringPtr("target"),
		},
	}
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	policy, response, err = vpcService.CreateInstanceGroupManagerPolicy(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies/{id}
// Delete specified instance group manager policy

func DeleteInstanceGroupManagerPolicy(vpcService *vpcbetav1.VpcbetaV1, gID, mID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupManagerPolicyOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	response, err = vpcService.DeleteInstanceGroupManagerPolicy(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies/{id}
// Retrieve specified instance group manager policy

func GetInstanceGroupManagerPolicy(vpcService *vpcbetav1.VpcbetaV1, gID, mID, id string) (template vpcbetav1.InstanceGroupManagerPolicyIntf, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceGroupManagerPolicyOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(gID)
	options.SetInstanceGroupManagerID(mID)
	template, response, err = vpcService.GetInstanceGroupManagerPolicy(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/managers/{instance_group_manager_id}/policies/{id}
// Update specified instance group manager policy
func UpdateInstanceGroupManagerPolicy(vpcService *vpcbetav1.VpcbetaV1, igID, mID, id, name string) (template vpcbetav1.InstanceGroupManagerPolicyIntf, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.InstanceGroupManagerPolicyPatch{
		MetricType:  core.StringPtr("cpu"),
		MetricValue: core.Int64Ptr(80),
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceGroupManagerPolicyOptions{
		InstanceGroupManagerPolicyPatch: patchBody,
	}
	options.SetID(id)
	options.SetInstanceGroupID(igID)
	options.SetInstanceGroupManagerID(mID)
	template, response, err = vpcService.UpdateInstanceGroupManagerPolicy(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/memberships
// Delete all memberships from the instance group

func DeleteInstanceGroupMemberships(vpcService *vpcbetav1.VpcbetaV1, igID string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupMembershipsOptions{}
	options.SetInstanceGroupID(igID)
	response, err = vpcService.DeleteInstanceGroupMemberships(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/memberships
// List all memberships for the instance group
func ListInstanceGroupMemberships(vpcService *vpcbetav1.VpcbetaV1, igID string) (members *vpcbetav1.InstanceGroupMembershipCollection, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.ListInstanceGroupMembershipsOptions{}
	options.SetInstanceGroupID(igID)
	members, response, err = vpcService.ListInstanceGroupMemberships(options)
	return
}

// DELETE
// /instance_groups/{instance_group_id}/memberships/{id}
// Delete specified instance group membership
func DeleteInstanceGroupMembership(vpcService *vpcbetav1.VpcbetaV1, igID, id string) (response *core.DetailedResponse, err error) {
	options := &vpcbetav1.DeleteInstanceGroupMembershipOptions{}
	options.SetInstanceGroupID(igID)
	options.SetID(id)
	response, err = vpcService.DeleteInstanceGroupMembership(options)
	return response, err
}

// GET
// /instance_groups/{instance_group_id}/memberships/{id}
// Retrieve specified instance group membership
func GetInstanceGroupMembership(vpcService *vpcbetav1.VpcbetaV1, igID, id string) (member *vpcbetav1.InstanceGroupMembership, response *core.DetailedResponse, err error) {
	options := &vpcbetav1.GetInstanceGroupMembershipOptions{}
	options.SetID(id)
	options.SetInstanceGroupID(igID)
	member, response, err = vpcService.GetInstanceGroupMembership(options)
	return
}

// PATCH
// /instance_groups/{instance_group_id}/memberships/{id}
// Update specified instance group membership
func UpdateInstanceGroupMembership(vpcService *vpcbetav1.VpcbetaV1, igID, id, name string) (membership *vpcbetav1.InstanceGroupMembership, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.InstanceGroupMembershipPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateInstanceGroupMembershipOptions{
		InstanceGroupMembershipPatch: patchBody,
	}
	options.SetID(id)
	options.SetInstanceGroupID(igID)
	membership, response, err = vpcService.UpdateInstanceGroupMembership(options)
	return
}

/**
 * Endpoint Gateways
 */

// ListEndpointGateways - GET
// /endpoint_gateway
// List all Endpoint Gateways
func ListEndpointGateways(vpcService *vpcbetav1.VpcbetaV1) (endpointGateways *vpcbetav1.EndpointGatewayCollection, response *core.DetailedResponse, err error) {
	listEndpointGatewaysOptions := vpcService.NewListEndpointGatewaysOptions()
	endpointGateways, response, err = vpcService.ListEndpointGateways(listEndpointGatewaysOptions)
	return
}

// GetEndpointGateway - GET
// /endpoint_gateway/{id}
// Retrieve the specified Endpoint Gateway
func GetEndpointGateway(vpcService *vpcbetav1.VpcbetaV1, id string) (endpointGateway *vpcbetav1.EndpointGateway, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetEndpointGatewayOptions(id)
	endpointGateway, response, err = vpcService.GetEndpointGateway(options)
	return
}

// DeleteEndpointGateway - DELETE
// /endpoint_gateway/{id}
// Delete the specified Endpoint Gateway
func DeleteEndpointGateway(vpcService *vpcbetav1.VpcbetaV1, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewDeleteEndpointGatewayOptions(id)
	response, err = vpcService.DeleteEndpointGateway(options)
	return response, err
}

// UpdateEndpointGateway - PATCH
// /endpoint_gateway/{id}
// Update the specified Endpoint Gateway
func UpdateEndpointGateway(vpcService *vpcbetav1.VpcbetaV1, id, name string) (endpointGateway *vpcbetav1.EndpointGateway, response *core.DetailedResponse, err error) {
	body := &vpcbetav1.EndpointGatewayPatch{
		Name: &name,
	}
	patchBody, _ := body.AsPatch()
	options := &vpcbetav1.UpdateEndpointGatewayOptions{
		ID:                   &id,
		EndpointGatewayPatch: patchBody,
	}

	endpointGateway, response, err = vpcService.UpdateEndpointGateway(options)
	return
}

// CreateEndpointGateway - POST
// /endpoint_gateway
// Reserve a Endpoint Gateway
func CreateEndpointGateway(vpcService *vpcbetav1.VpcbetaV1, vpcId string) (endpointGateway *vpcbetav1.EndpointGateway, response *core.DetailedResponse, err error) {
	endpointGatewayTargetPrototypeModel := &vpcbetav1.EndpointGatewayTargetPrototypeProviderInfrastructureServiceIdentityProviderInfrastructureServiceIdentityByName{
		ResourceType: core.StringPtr("provider_infrastructure_service"),
		Name:         core.StringPtr("ibm-ntp-server"),
	}

	vpcIdentityModel := &vpcbetav1.VPCIdentityByID{
		ID: &vpcId,
	}

	options := vpcService.NewCreateEndpointGatewayOptions(
		endpointGatewayTargetPrototypeModel,
		vpcIdentityModel,
	)

	endpointGateway, response, err = vpcService.CreateEndpointGateway(options)
	return
}

// GET
// /endpoint_gateways/{endpoint_gateway_id}/ips
// List all reserved IPs bound to an endpoint gateway
func ListEndpointGatewayIps(vpcService *vpcbetav1.VpcbetaV1, endpointGatewayId string) (reserveIPList *vpcbetav1.ReservedIPCollectionEndpointGatewayContext, response *core.DetailedResponse, err error) {
	options := vpcService.NewListEndpointGatewayIpsOptions(
		endpointGatewayId,
	)
	reserveIPList, response, err = vpcService.ListEndpointGatewayIps(options)
	return
}

// DELETE
// /endpoint_gateways/{endpoint_gateway_id}/ips/{id}
// Unbind a reserved IP from an endpoint gateway
func RemoveEndpointGatewayIP(vpcService *vpcbetav1.VpcbetaV1, endpointGatewayId, id string) (response *core.DetailedResponse, err error) {
	options := vpcService.NewRemoveEndpointGatewayIPOptions(endpointGatewayId, id)
	response, err = vpcService.RemoveEndpointGatewayIP(options)
	return response, err
}

// GET
// /endpoint_gateways/{endpoint_gateway_id}/ips/{id}
// Retrieve a reserved IP bound to an endpoint gateway
func GetEndpointGatewayIP(vpcService *vpcbetav1.VpcbetaV1, endpointGatewayId, id string) (reserveIP *vpcbetav1.ReservedIP, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetEndpointGatewayIPOptions(
		endpointGatewayId,
		id,
	)
	reserveIP, response, err = vpcService.GetEndpointGatewayIP(options)
	return
}

// PUT
// /endpoint_gateways/{endpoint_gateway_id}/ips/{id}
// Bind a reserved IP to an endpoint gateway
func AddEndpointGatewayIP(vpcService *vpcbetav1.VpcbetaV1, endpointGatewayId, id string) (reserveIP *vpcbetav1.ReservedIP, response *core.DetailedResponse, err error) {
	addEndpointGatewayIPOptions := vpcService.NewAddEndpointGatewayIPOptions(
		endpointGatewayId,
		id,
	)

	reserveIP, response, err = vpcService.AddEndpointGatewayIP(addEndpointGatewayIPOptions)
	return
}

/**
 * Routing Tables
 */

// GET
// /subnets/{id}/routing_table
// Retrieve a subnet's attached routing table
func GetSubnetRoutingTable(vpcService *vpcbetav1.VpcbetaV1, subnetId string) (routingTable *vpcbetav1.RoutingTable, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetSubnetRoutingTableOptions(
		subnetId,
	)
	routingTable, response, err = vpcService.GetSubnetRoutingTable(options)
	return
}

// PUT
// /subnets/{id}/routing_table
// Attach a routing table to a subnet
func ReplaceSubnetRoutingTable(vpcService *vpcbetav1.VpcbetaV1, subnetId, routingTableId string) (routingTable *vpcbetav1.RoutingTable, response *core.DetailedResponse, err error) {
	routingTableIdentityModel := &vpcbetav1.RoutingTableIdentityByID{
		ID: &routingTableId,
	}
	options := vpcService.NewReplaceSubnetRoutingTableOptions(
		subnetId,
		routingTableIdentityModel,
	)
	routingTable, response, err = vpcService.ReplaceSubnetRoutingTable(options)
	return
}

// GET
// /vpcs/{id}/default_routing_table
// Retrieve a VPC's default routing table
func GetVPCDefaultRoutingTable(vpcService *vpcbetav1.VpcbetaV1, vpcId string) (defaultRoutingTable *vpcbetav1.DefaultRoutingTable, response *core.DetailedResponse, err error) {
	getVPCDefaultRoutingTableOptions := vpcService.NewGetVPCDefaultRoutingTableOptions(
		vpcId,
	)

	defaultRoutingTable, response, err = vpcService.GetVPCDefaultRoutingTable(getVPCDefaultRoutingTableOptions)
	return
}

// GET
// /vpcs/{vpc_id}/routing_tables
// List all routing tables for a VPC
func ListVPCRoutingTables(vpcService *vpcbetav1.VpcbetaV1, vpcId string) (routingTableCollection *vpcbetav1.RoutingTableCollection, response *core.DetailedResponse, err error) {
	listVPCRoutingTablesOptions := vpcService.NewListVPCRoutingTablesOptions(
		vpcId,
	)

	routingTableCollection, response, err = vpcService.ListVPCRoutingTables(listVPCRoutingTablesOptions)
	return
}

// POST
// /vpcs/{vpc_id}/routing_tables
// Create a VPC routing table
func CreateVPCRoutingTable(vpcService *vpcbetav1.VpcbetaV1, vpcId, name, zoneName string) (routingTable *vpcbetav1.RoutingTable, response *core.DetailedResponse, err error) {
	routeNextHopPrototypeModel := &vpcbetav1.RoutePrototypeNextHopRouteNextHopPrototypeRouteNextHopIP{
		Address: core.StringPtr("192.168.3.4"),
	}

	zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
		Name: &zoneName,
	}

	routePrototypeModel := &vpcbetav1.RoutePrototype{
		Action:      core.StringPtr("delegate"),
		Destination: core.StringPtr("192.168.3.0/24"),
		NextHop:     routeNextHopPrototypeModel,
		Zone:        zoneIdentityModel,
	}

	createVPCRoutingTableOptions := &vpcbetav1.CreateVPCRoutingTableOptions{
		VPCID:  &vpcId,
		Name:   &name,
		Routes: []vpcbetav1.RoutePrototype{*routePrototypeModel},
	}

	routingTable, response, err = vpcService.CreateVPCRoutingTable(createVPCRoutingTableOptions)
	return
}

// DELETE
// /vpcs/{vpc_id}/routing_tables/{id}
// Delete specified VPC routing table
func DeleteVPCRoutingTable(vpcService *vpcbetav1.VpcbetaV1, vpcId, id string) (response *core.DetailedResponse, err error) {
	deleteVPCRoutingTableOptions := vpcService.NewDeleteVPCRoutingTableOptions(vpcId, id)
	response, err = vpcService.DeleteVPCRoutingTable(deleteVPCRoutingTableOptions)
	return response, err
}

// GET
// /vpcs/{vpc_id}/routing_tables/{id}
// Retrieve specified VPC routing table
func GetVPCRoutingTable(vpcService *vpcbetav1.VpcbetaV1, vpcId, id string) (routingTable *vpcbetav1.RoutingTable, response *core.DetailedResponse, err error) {
	getVPCRoutingTableOptions := vpcService.NewGetVPCRoutingTableOptions(vpcId, id)
	routingTable, response, err = vpcService.GetVPCRoutingTable(getVPCRoutingTableOptions)
	return
}

// PATCH
// /vpcs/{vpc_id}/routing_tables/{id}
// Update specified VPC routing table
func UpdateVPCRoutingTable(vpcService *vpcbetav1.VpcbetaV1, vpcId, id, name string) (routingTable *vpcbetav1.RoutingTable, response *core.DetailedResponse, err error) {
	routingTablePatchModel := &vpcbetav1.RoutingTablePatch{
		Name: &name,
	}
	routingTablePatchModelAsPatch, _ := routingTablePatchModel.AsPatch()

	updateVPCRoutingTableOptions := &vpcbetav1.UpdateVPCRoutingTableOptions{
		VPCID:             &vpcId,
		ID:                &id,
		RoutingTablePatch: routingTablePatchModelAsPatch,
	}

	routingTable, response, err = vpcService.UpdateVPCRoutingTable(updateVPCRoutingTableOptions)
	return
}

// GET
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes
// List all the routes of a VPC routing table
func ListVPCRoutingTableRoutes(vpcService *vpcbetav1.VpcbetaV1, vpcId, routingTableID string) (routeCollection *vpcbetav1.RouteCollection, response *core.DetailedResponse, err error) {
	listVPCRoutingTableRoutesOptions := &vpcbetav1.ListVPCRoutingTableRoutesOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableID,
	}

	routeCollection, response, err = vpcService.ListVPCRoutingTableRoutes(listVPCRoutingTableRoutesOptions)
	return
}

// POST
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes
// Create a VPC route
func CreateVPCRoutingTableRoute(vpcService *vpcbetav1.VpcbetaV1, vpcId, routingTableId, zone string) (route *vpcbetav1.Route, response *core.DetailedResponse, err error) {
	routeNextHopPrototypeModel := &vpcbetav1.RoutePrototypeNextHopRouteNextHopPrototypeRouteNextHopIP{
		Address: core.StringPtr("192.168.3.4"),
	}

	zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
		Name: &zone,
	}

	createVPCRoutingTableRouteOptions := &vpcbetav1.CreateVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		Destination:    core.StringPtr("192.168.3.0/24"),
		NextHop:        routeNextHopPrototypeModel,
		Zone:           zoneIdentityModel,
		Action:         core.StringPtr("delegate"),
	}

	route, response, err = vpcService.CreateVPCRoutingTableRoute(createVPCRoutingTableRouteOptions)
	return
}

// DELETE
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes/{id}
// Delete the specified VPC route
func DeleteVPCRoutingTableRoute(vpcService *vpcbetav1.VpcbetaV1, vpcId, routingTableId, id string) (response *core.DetailedResponse, err error) {
	deleteVPCRoutingTableRouteOptions := &vpcbetav1.DeleteVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		ID:             &id,
	}

	response, err = vpcService.DeleteVPCRoutingTableRoute(deleteVPCRoutingTableRouteOptions)
	return response, err
}

// GET
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes/{id}
// Retrieve the specified VPC route
func GetVPCRoutingTableRoute(vpcService *vpcbetav1.VpcbetaV1, vpcId, routingTableId, id string) (route *vpcbetav1.Route, response *core.DetailedResponse, err error) {
	getVPCRoutingTableRouteOptions := &vpcbetav1.GetVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		ID:             &id,
	}
	route, response, err = vpcService.GetVPCRoutingTableRoute(getVPCRoutingTableRouteOptions)
	return
}

// PATCH
// /vpcs/{vpc_id}/routing_tables/{routing_table_id}/routes/{id}
// Update the specified VPC route
func UpdateVPCRoutingTableRoute(vpcService *vpcbetav1.VpcbetaV1, vpcId, routingTableId, id, name string) (route *vpcbetav1.Route, response *core.DetailedResponse, err error) {
	routePatchModel := &vpcbetav1.RoutePatch{
		Name: &name,
	}
	routePatchModelAsPatch, _ := routePatchModel.AsPatch()

	updateVPCRoutingTableRouteOptions := &vpcbetav1.UpdateVPCRoutingTableRouteOptions{
		VPCID:          &vpcId,
		RoutingTableID: &routingTableId,
		ID:             &id,
		RoutePatch:     routePatchModelAsPatch,
	}

	route, response, err = vpcService.UpdateVPCRoutingTableRoute(updateVPCRoutingTableRouteOptions)
	return
}

// GET
// /dedicated_host/groups
// List all dedicated host groups
func ListDedicatedHostGroups(vpcService *vpcbetav1.VpcbetaV1) (dedicatedHostGroupCollection *vpcbetav1.DedicatedHostGroupCollection, response *core.DetailedResponse, err error) {
	listDedicatedHostGroupsOptions := &vpcbetav1.ListDedicatedHostGroupsOptions{}

	dedicatedHostGroupCollection, response, err = vpcService.ListDedicatedHostGroups(listDedicatedHostGroupsOptions)
	return
}

// POST
// /dedicated_host/groups
// Create a dedicated host group
func CreateDedicatedHostGroup(vpcService *vpcbetav1.VpcbetaV1, name, zone string) (dedicatedHostGroup *vpcbetav1.DedicatedHostGroup, response *core.DetailedResponse, err error) {
	zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
		Name: &zone,
	}
	createDedicatedHostGroupOptions := &vpcbetav1.CreateDedicatedHostGroupOptions{
		Class:  core.StringPtr("mx2"),
		Family: core.StringPtr("balanced"),
		Name:   &name,
		Zone:   zoneIdentityModel,
	}

	dedicatedHostGroup, response, err = vpcService.CreateDedicatedHostGroup(createDedicatedHostGroupOptions)
	return
}

// DELETE
// /dedicated_host/groups/{id}
// Delete specified dedicated host group
func DeleteDedicatedHostGroup(vpcService *vpcbetav1.VpcbetaV1, id *string) (response *core.DetailedResponse, err error) {
	getDedicatedHostGroupOptions := &vpcbetav1.DeleteDedicatedHostGroupOptions{
		ID: id,
	}
	response, err = vpcService.DeleteDedicatedHostGroup(getDedicatedHostGroupOptions)
	return
}

// GET
// /dedicated_host/groups/{id}
// Retrieve a dedicated host group
func GetDedicatedHostGroup(vpcService *vpcbetav1.VpcbetaV1, id *string) (dedicatedHostGroup *vpcbetav1.DedicatedHostGroup, response *core.DetailedResponse, err error) {
	getDedicatedHostGroupOptions := &vpcbetav1.GetDedicatedHostGroupOptions{
		ID: id,
	}

	dedicatedHostGroup, response, err = vpcService.GetDedicatedHostGroup(getDedicatedHostGroupOptions)
	return
}

// PATCH
// /dedicated_host/groups/{id}
// Update specified dedicated host group
func UpdateDedicatedHostGroup(vpcService *vpcbetav1.VpcbetaV1, id, name *string) (dedicatedHostGroup *vpcbetav1.DedicatedHostGroup, response *core.DetailedResponse, err error) {
	model := &vpcbetav1.DedicatedHostGroupPatch{
		Name: name,
	}
	patch, _ := model.AsPatch()
	getDedicatedHostGroupOptions := &vpcbetav1.UpdateDedicatedHostGroupOptions{
		ID:                      id,
		DedicatedHostGroupPatch: patch,
	}

	dedicatedHostGroup, response, err = vpcService.UpdateDedicatedHostGroup(getDedicatedHostGroupOptions)
	return
}

// GET
// /dedicated_host/profiles
// List all dedicated host profiles
func ListDedicatedHostProfiles(vpcService *vpcbetav1.VpcbetaV1) (result *vpcbetav1.DedicatedHostProfileCollection, response *core.DetailedResponse, err error) {
	listDedicatedHostProfilesOptions := &vpcbetav1.ListDedicatedHostProfilesOptions{}

	result, response, err = vpcService.ListDedicatedHostProfiles(listDedicatedHostProfilesOptions)
	return
}

// GET
// /dedicated_host/profiles/{name}
// Retrieve specified dedicated host profile
func GetDedicatedHostProfile(vpcService *vpcbetav1.VpcbetaV1, name *string) (result *vpcbetav1.DedicatedHostProfile, response *core.DetailedResponse, err error) {
	getDedicatedHostProfileOptions := &vpcbetav1.GetDedicatedHostProfileOptions{
		Name: name,
	}
	result, response, err = vpcService.GetDedicatedHostProfile(getDedicatedHostProfileOptions)
	return
}

// GET
// /dedicated_hosts
// List all dedicated hosts
func ListDedicatedHosts(vpcService *vpcbetav1.VpcbetaV1) (dedicatedHostCollection *vpcbetav1.DedicatedHostCollection, response *core.DetailedResponse, err error) {
	listDedicatedHostsOptions := &vpcbetav1.ListDedicatedHostsOptions{}

	dedicatedHostCollection, response, err = vpcService.ListDedicatedHosts(listDedicatedHostsOptions)
	return
}

// POST
// /dedicated_hosts
// Create a dedicated host
func CreateDedicatedHost(vpcService *vpcbetav1.VpcbetaV1, name, profile, groupID *string) (dedicatedHost *vpcbetav1.DedicatedHost, response *core.DetailedResponse, err error) {
	fmt.Println(" sd", *name, *groupID)
	dedicatedHostProfileIdentityModel := &vpcbetav1.DedicatedHostProfileIdentityByName{
		Name: profile,
	}

	dedicatedHostGroupIdentityModel := &vpcbetav1.DedicatedHostGroupIdentityByID{
		ID: groupID,
	}

	dedicatedHostPrototypeModel := &vpcbetav1.DedicatedHostPrototypeDedicatedHostByGroup{
		Name:    name,
		Profile: dedicatedHostProfileIdentityModel,
		Group:   dedicatedHostGroupIdentityModel,
	}

	createDedicatedHostOptions := &vpcbetav1.CreateDedicatedHostOptions{
		DedicatedHostPrototype: dedicatedHostPrototypeModel,
	}

	dedicatedHost, response, err = vpcService.CreateDedicatedHost(createDedicatedHostOptions)
	return
}

// DELETE
// /dedicated_hosts/{id}
// Delete specified dedicated host
func DeleteDedicatedHost(vpcService *vpcbetav1.VpcbetaV1, id *string) (response *core.DetailedResponse, err error) {
	getDedicatedHostOptions := &vpcbetav1.DeleteDedicatedHostOptions{
		ID: id,
	}
	response, err = vpcService.DeleteDedicatedHost(getDedicatedHostOptions)
	return
}

// GET
// /dedicated_hosts/{id}
// Retrieve a dedicated host
func GetDedicatedHost(vpcService *vpcbetav1.VpcbetaV1, id string) (dedicatedHost *vpcbetav1.DedicatedHost, response *core.DetailedResponse, err error) {
	getDedicatedHostOptions := &vpcbetav1.GetDedicatedHostOptions{
		ID: &id,
	}

	dedicatedHost, response, err = vpcService.GetDedicatedHost(getDedicatedHostOptions)
	return
}

// PATCH
// /dedicated_hosts/{id}
// Update specified dedicated host
func UpdateDedicatedHost(vpcService *vpcbetav1.VpcbetaV1, name, id *string) (dedicatedHost *vpcbetav1.DedicatedHost, response *core.DetailedResponse, err error) {
	model := &vpcbetav1.DedicatedHostPatch{
		Name: name,
	}
	patch, _ := model.AsPatch()
	getDedicatedHostOptions := &vpcbetav1.UpdateDedicatedHostOptions{
		ID:                 id,
		DedicatedHostPatch: patch,
	}

	dedicatedHost, response, err = vpcService.UpdateDedicatedHost(getDedicatedHostOptions)
	return
}

// Snapshots

func ListSnapshots(vpcService *vpcbetav1.VpcbetaV1) (snapshots *vpcbetav1.SnapshotCollection, response *core.DetailedResponse, err error) {
	options := vpcService.NewListSnapshotsOptions()
	snapshots, response, err = vpcService.ListSnapshots(options)
	return
}

func CreateSnapshot(vpcService *vpcbetav1.VpcbetaV1, volumeID, name string) (snapshot *vpcbetav1.Snapshot, response *core.DetailedResponse, err error) {
	volumeIdentityModel := &vpcbetav1.VolumeIdentity{
		ID: &volumeID,
	}
	snapshotPrototypeModel := &vpcbetav1.SnapshotPrototypeSnapshotBySourceVolume{
		Name:         core.StringPtr("my-snapshot-1"),
		SourceVolume: volumeIdentityModel,
	}

	options := &vpcbetav1.CreateSnapshotOptions{
		SnapshotPrototype: snapshotPrototypeModel,
	}

	snapshot, response, err = vpcService.CreateSnapshot(options)
	return
}

func DeleteSnapshot(vpcService *vpcbetav1.VpcbetaV1, snapshotId, ifMatch string) (response *core.DetailedResponse, err error) {
	deleteSnapshotOptions := vpcService.NewDeleteSnapshotOptions(
		snapshotId,
	)
	deleteSnapshotOptions.SetIfMatch(ifMatch)
	response, err = vpcService.DeleteSnapshot(deleteSnapshotOptions)
	return response, err
}

func DeleteSnapshots(vpcService *vpcbetav1.VpcbetaV1, volumeID string) (response *core.DetailedResponse, err error) {
	deleteSnapshotsOptions := vpcService.NewDeleteSnapshotsOptions(
		volumeID,
	)
	response, err = vpcService.DeleteSnapshots(deleteSnapshotsOptions)
	return response, err
}

func GetSnapshot(vpcService *vpcbetav1.VpcbetaV1, snapshotId string) (snapshot *vpcbetav1.Snapshot, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetSnapshotOptions(
		snapshotId,
	)
	snapshot, response, err = vpcService.GetSnapshot(options)
	return
}

func UpdateSnapshot(vpcService *vpcbetav1.VpcbetaV1, userTags []string, snapshotId, name, ifMatch string) (snapshot *vpcbetav1.Snapshot, response *core.DetailedResponse, err error) {
	snapshotPatchModel := &vpcbetav1.SnapshotPatch{
		Name:     &name,
		UserTags: userTags,
	}
	snapshotPatchModelAsPatch, _ := snapshotPatchModel.AsPatch()
	updateSnapshotOptions := &vpcbetav1.UpdateSnapshotOptions{
		ID:            &snapshotId,
		SnapshotPatch: snapshotPatchModelAsPatch,
		IfMatch:       &ifMatch,
	}
	snapshot, response, err = vpcService.UpdateSnapshot(updateSnapshotOptions)
	return
}

// Placement Groups
func ListPlacementGroups(vpcService *vpcbetav1.VpcbetaV1) (placementGroupCollection *vpcbetav1.PlacementGroupCollection, response *core.DetailedResponse, err error) {
	listPlacementGroupsOptions := &vpcbetav1.ListPlacementGroupsOptions{}

	placementGroupCollection, response, err = vpcService.ListPlacementGroups(listPlacementGroupsOptions)
	return
}

func CreatePlacementGroup(vpcService *vpcbetav1.VpcbetaV1, strategy, name, resgroup string) (placementGroup *vpcbetav1.PlacementGroup, response *core.DetailedResponse, err error) {
	resourceGroupIdentityModel := &vpcbetav1.ResourceGroupIdentityByID{
		ID: &resgroup,
	}
	createPlacementGroupOptions := &vpcbetav1.CreatePlacementGroupOptions{
		Strategy:      &strategy,
		Name:          &name,
		ResourceGroup: resourceGroupIdentityModel,
	}

	placementGroup, response, err = vpcService.CreatePlacementGroup(createPlacementGroupOptions)
	return
}

func GetPlacementGroup(vpcService *vpcbetav1.VpcbetaV1, pgID string) (placementGroup *vpcbetav1.PlacementGroup, response *core.DetailedResponse, err error) {
	getPlacementGroupOptions := &vpcbetav1.GetPlacementGroupOptions{
		ID: &pgID,
	}

	placementGroup, response, err = vpcService.GetPlacementGroup(getPlacementGroupOptions)
	return
}

func UpdatePlacementGroup(vpcService *vpcbetav1.VpcbetaV1, pgID, name string) (placementGroup *vpcbetav1.PlacementGroup, response *core.DetailedResponse, err error) {
	placementGroupPatchModel := &vpcbetav1.PlacementGroupPatch{
		Name: &name,
	}
	placementGroupPatchModelAsPatch, _ := placementGroupPatchModel.AsPatch()

	updatePlacementGroupOptions := &vpcbetav1.UpdatePlacementGroupOptions{
		ID:                  &pgID,
		PlacementGroupPatch: placementGroupPatchModelAsPatch,
	}

	placementGroup, response, err = vpcService.UpdatePlacementGroup(updatePlacementGroupOptions)
	return
}

func DeletePlacementGroup(vpcService *vpcbetav1.VpcbetaV1, pgID string) (response *core.DetailedResponse, err error) {
	deletePlacementGroupOptions := &vpcbetav1.DeletePlacementGroupOptions{
		ID: &pgID,
	}

	response, err = vpcService.DeletePlacementGroup(deletePlacementGroupOptions)
	return
}

// Print - Marshal JSON and print
func Print(printObject interface{}) {
	p, _ := json.MarshalIndent(printObject, "", "\t")
	fmt.Println(string(p))
}

// PollInstance - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollInstance(vpcService *vpcbetav1.VpcbetaV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetInstance(vpcService, ID)
			fmt.Println("Current status of VSI - ", *res.Status)
			fmt.Println("Expected status of VSI - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving instance ID %s with err error message: %s", ID, err)
				return false
			}
			if *res.Status == "pending" {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++

		}
	}
}

// PollSubnet - poll and check the status of VSI before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollSubnet(vpcService *vpcbetav1.VpcbetaV1, ID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetSubnet(vpcService, ID)
			fmt.Println("Current status of Subnet - ", *res.Status)
			fmt.Println("Expected status of Subnet - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving subnet ID %s with err error message: %s", ID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++

		}
	}
}

// PollVolAttachment - poll and check the status of Volume attachment before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollVolAttachment(vpcService *vpcbetav1.VpcbetaV1, vpcID, volAttachmentID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVolumeAttachment(vpcService, vpcID, volAttachmentID)
			fmt.Println("Current status of attachment - ", *res.Status)
			fmt.Println("Expected status of attachment - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving volume attachment ID %s with err error message: %s", vpcID, err)
				return false
			}
			if *res.Status == status {
				fmt.Println("Received expected status - ", *res.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// PollLB - poll and check the status of LB Listener before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollLB(vpcService *vpcbetav1.VpcbetaV1, lbID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetLoadBalancer(vpcService, lbID)
			fmt.Println("Current status of load balancer - ", *res.ProvisioningStatus)
			fmt.Println("Expected status of load balancer - ", status)
			if err != nil && res == nil {
				fmt.Printf("err error: Retrieving load balancer ID %s with err error message: %s", lbID, err)
				return false
			}
			if *res.ProvisioningStatus == status {
				fmt.Println("Received expected status - ", *res.ProvisioningStatus)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// PollVPNGateway - poll and check the status of VPNGateway before performing any action
// ID - resource ID
// status - expected status/ poll until this status is returned
// pollFrequency - number of times polling happens
func PollVPNGateway(vpcService *vpcbetav1.VpcbetaV1, gatewayID, status string, pollFrequency int) bool {
	count := 1
	for {
		if count < pollFrequency {
			res, _, err := GetVPNGateway(vpcService, gatewayID)
			vpn := res.(*vpcbetav1.VPNGateway)
			fmt.Println("Current status of VPNGateway - ", *vpn.Status)
			fmt.Println("Expected status of VPNGateway - ", status)
			if err != nil && vpn == nil {
				fmt.Printf("err error: Retrieving VPNGateway ID %s with err error message: %s", gatewayID, err)
				return false
			}
			if *vpn.Status == status {
				fmt.Println("Received expected status - ", *vpn.Status)
				return true
			}
			fmt.Printf("Waiting (60 sec) for resource to change status. Attempt - %d", count)
			time.Sleep(60 * time.Second)
			count++
		}
	}
}

// Backup Policies

func ListBackupPolicies(vpcService *vpcbetav1.VpcbetaV1) (backupPoliciesCollection *vpcbetav1.BackupPolicyCollection, response *core.DetailedResponse, err error) {
	options := vpcService.NewListBackupPoliciesOptions()
	backupPoliciesCollection, response, err = vpcService.ListBackupPolicies(options)
	return
}
func ListBackupPolicyPlans(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID string) (backupPolicyPlanCollection *vpcbetav1.BackupPolicyPlanCollection, response *core.DetailedResponse, err error) {
	options := vpcService.NewListBackupPolicyPlansOptions(
		backupPolicyID,
	)
	backupPolicyPlanCollection, response, err = vpcService.ListBackupPolicyPlans(options)
	return
}

func CreateBackupPolicy(vpcService *vpcbetav1.VpcbetaV1, name string, userTags []string) (backupPolicy *vpcbetav1.BackupPolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewCreateBackupPolicyOptions()
	options.SetName(name)
	options.SetMatchUserTags(userTags)

	backupPolicy, response, err = vpcService.CreateBackupPolicy(options)
	return
}

func CreateBackupPolicyPlan(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID, cronSpec, name string) (backupPolicyPlan *vpcbetav1.BackupPolicyPlan, response *core.DetailedResponse, err error) {
	options := vpcService.NewCreateBackupPolicyPlanOptions(
		backupPolicyID,
		cronSpec,
	)
	options.SetName(name)

	backupPolicyPlan, response, err = vpcService.CreateBackupPolicyPlan(options)
	return
}

func DeleteBackupPolicyPlan(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID, backupPolicyPlanID, ifMatch string) (response *core.DetailedResponse, err error) {
	deleteBackupPolicyPlanOptions := vpcService.NewDeleteBackupPolicyPlanOptions(
		backupPolicyID,
		backupPolicyPlanID,
	)
	_, response, err = vpcService.DeleteBackupPolicyPlan(deleteBackupPolicyPlanOptions)
	return response, err
}

func DeleteBackupPolicy(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID, ifMatch string) (response *core.DetailedResponse, err error) {
	deleteBackupPolicyOptions := vpcService.NewDeleteBackupPolicyOptions(
		backupPolicyID,
	)
	_, response, err = vpcService.DeleteBackupPolicy(deleteBackupPolicyOptions)
	return response, err
}

func GetBackupPolicyPlan(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID, backupPolicyPlanID string) (backupPolicyPlan *vpcbetav1.BackupPolicyPlan, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetBackupPolicyPlanOptions(
		backupPolicyID,
		backupPolicyPlanID,
	)
	backupPolicyPlan, response, err = vpcService.GetBackupPolicyPlan(options)
	return
}
func GetBackupPolicy(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID string) (backupPolicy *vpcbetav1.BackupPolicy, response *core.DetailedResponse, err error) {
	options := vpcService.NewGetBackupPolicyOptions(
		backupPolicyID,
	)
	backupPolicy, response, err = vpcService.GetBackupPolicy(options)
	return
}

func UpdateBackupPolicyPlan(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID, backupPolicyPlanID, name, ifMatch string) (backupPolicyPlan *vpcbetav1.BackupPolicyPlan, response *core.DetailedResponse, err error) {
	backupPolicyPlanPatchModel := &vpcbetav1.BackupPolicyPlanPatch{
		Name: &name,
	}
	backupPolicyPlanPatchModelAsPatch, _ := backupPolicyPlanPatchModel.AsPatch()
	updateBackupPolicyPlanOptions := vpcService.NewUpdateBackupPolicyPlanOptions(
		backupPolicyID,
		backupPolicyPlanID,
		backupPolicyPlanPatchModelAsPatch,
	)
	updateBackupPolicyPlanOptions.SetIfMatch(ifMatch)
	backupPolicyPlan, response, err = vpcService.UpdateBackupPolicyPlan(updateBackupPolicyPlanOptions)
	return
}
func UpdateBackupPolicy(vpcService *vpcbetav1.VpcbetaV1, backupPolicyID, name, ifMatch string) (backupPolicy *vpcbetav1.BackupPolicy, response *core.DetailedResponse, err error) {
	backupPolicyPatchModel := &vpcbetav1.BackupPolicyPatch{
		Name: &name,
	}
	backupPolicyPatchModelAsPatch, _ := backupPolicyPatchModel.AsPatch()
	updateBackupPolicyOptions := vpcService.NewUpdateBackupPolicyOptions(
		backupPolicyID,
		backupPolicyPatchModelAsPatch,
	)
	updateBackupPolicyOptions.SetIfMatch(ifMatch)
	backupPolicy, response, err = vpcService.UpdateBackupPolicy(updateBackupPolicyOptions)
	return
}

// VPN servers

func ListVPNServers(vpcService *vpcbetav1.VpcbetaV1) (vpnServerCollection *vpcbetav1.VPNServerCollection, response *core.DetailedResponse, err error) {
	listVPNServersOptions := vpcService.NewListVPNServersOptions()
	vpnServerCollection, response, err = vpcService.ListVPNServers(listVPNServersOptions)
	return
}
func ListVPNServerClients(vpcService *vpcbetav1.VpcbetaV1, vpnServerID string) (vpnServerClientCollection *vpcbetav1.VPNServerClientCollection, response *core.DetailedResponse, err error) {
	listVPNServerClientsOptions := vpcService.NewListVPNServerClientsOptions(
		vpnServerID,
	)
	vpnServerClientCollection, response, err = vpcService.ListVPNServerClients(listVPNServerClientsOptions)
	return
}
func ListVPNServerRoutes(vpcService *vpcbetav1.VpcbetaV1, vpnServerID string) (vpnServerRouteCollection *vpcbetav1.VPNServerRouteCollection, response *core.DetailedResponse, err error) {
	listVPNServerRoutesOptions := vpcService.NewListVPNServerRoutesOptions(
		vpnServerID,
	)
	vpnServerRouteCollection, response, err = vpcService.ListVPNServerRoutes(listVPNServerRoutesOptions)
	return
}

func CreateVPNServer(vpcService *vpcbetav1.VpcbetaV1, subnetID, crn, providerType, method, name, clientIPPool string) (vpnServer *vpcbetav1.VPNServer, response *core.DetailedResponse, err error) {

	certificateInstanceIdentityModel := &vpcbetav1.CertificateInstanceIdentityByCRN{
		CRN: core.StringPtr(crn),
	}

	vpnServerAuthenticationByUsernameIDProviderModel := &vpcbetav1.VPNServerAuthenticationByUsernameIDProviderByIam{
		ProviderType: core.StringPtr(providerType),
	}

	vpnServerAuthenticationPrototypeModel := &vpcbetav1.VPNServerAuthenticationPrototypeVPNServerAuthenticationByUsernamePrototype{
		Method:           core.StringPtr(method),
		IdentityProvider: vpnServerAuthenticationByUsernameIDProviderModel,
	}

	subnetIdentityModel := &vpcbetav1.SubnetIdentityByID{
		ID: core.StringPtr(subnetID),
	}

	options := vpcService.NewCreateVPNServerOptions(
		certificateInstanceIdentityModel,
		[]vpcbetav1.VPNServerAuthenticationPrototypeIntf{vpnServerAuthenticationPrototypeModel},
		clientIPPool,
		[]vpcbetav1.SubnetIdentityIntf{subnetIdentityModel},
	)
	options.SetName(name)

	vpnServer, response, err = vpcService.CreateVPNServer(options)
	return
}

func DisconnectVPNClient(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, vpnClientID string) (response *core.DetailedResponse, err error) {
	disconnectVPNClientOptions := vpcService.NewDisconnectVPNClientOptions(
		vpnServerID,
		vpnClientID,
	)

	response, err = vpcService.DisconnectVPNClient(disconnectVPNClientOptions)
	return
}
func CreateVPNServerRoute(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, destination, name string) (vpnServerRoute *vpcbetav1.VPNServerRoute, response *core.DetailedResponse, err error) {
	createVPNServerRouteOptions := vpcService.NewCreateVPNServerRouteOptions(
		vpnServerID,
		destination,
	)
	createVPNServerRouteOptions.SetName(name)

	vpnServerRoute, response, err = vpcService.CreateVPNServerRoute(createVPNServerRouteOptions)
	return
}

func DeleteVPNServerClient(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, vpnClientID string) (response *core.DetailedResponse, err error) {
	deleteVPNServerClientOptions := vpcService.NewDeleteVPNServerClientOptions(
		vpnServerID,
		vpnClientID,
	)

	response, err = vpcService.DeleteVPNServerClient(deleteVPNServerClientOptions)
	return response, err
}

func DeleteVPNServerRoute(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, vpnServerRouteID string) (response *core.DetailedResponse, err error) {
	deleteVPNServerRouteOptions := vpcService.NewDeleteVPNServerRouteOptions(
		vpnServerID,
		vpnServerRouteID,
	)

	response, err = vpcService.DeleteVPNServerRoute(deleteVPNServerRouteOptions)
	return response, err
}
func DeleteVPNServer(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, ifMatchVPNServer string) (response *core.DetailedResponse, err error) {
	deleteVPNServerOptions := vpcService.NewDeleteVPNServerOptions(
		vpnServerID,
	)
	deleteVPNServerOptions.SetIfMatch(ifMatchVPNServer)

	response, err = vpcService.DeleteVPNServer(deleteVPNServerOptions)
	return response, err
}

func GetVPNServer(vpcService *vpcbetav1.VpcbetaV1, vpnServerID string) (vpnServer *vpcbetav1.VPNServer, response *core.DetailedResponse, err error) {
	getVPNServerOptions := vpcService.NewGetVPNServerOptions(
		vpnServerID,
	)

	vpnServer, response, err = vpcService.GetVPNServer(getVPNServerOptions)
	return
}
func GetVPNServerClient(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, vpnClientID string) (vpnServerClient *vpcbetav1.VPNServerClient, response *core.DetailedResponse, err error) {
	getVPNServerClientOptions := vpcService.NewGetVPNServerClientOptions(
		vpnServerID,
		vpnClientID,
	)

	vpnServerClient, response, err = vpcService.GetVPNServerClient(getVPNServerClientOptions)
	return
}
func GetVPNServerRoute(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, vpnServerRouteID string) (vpnServerRoute *vpcbetav1.VPNServerRoute, response *core.DetailedResponse, err error) {
	getVPNServerRouteOptions := vpcService.NewGetVPNServerRouteOptions(
		vpnServerID,
		vpnServerRouteID,
	)

	vpnServerRoute, response, err = vpcService.GetVPNServerRoute(getVPNServerRouteOptions)
	return
}
func GetVPNServerClientConfiguration(vpcService *vpcbetav1.VpcbetaV1, vpnServerID string) (vpnServerClientConfiguration *string, response *core.DetailedResponse, err error) {
	getVPNServerClientConfigurationOptions := vpcService.NewGetVPNServerClientConfigurationOptions(
		vpnServerID,
	)

	vpnServerClientConfiguration, response, err = vpcService.GetVPNServerClientConfiguration(getVPNServerClientConfigurationOptions)
	return
}

func UpdateVPNServerRoute(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, vpnServerRouteID, name, ifMatch string) (vpnServerRoute *vpcbetav1.VPNServerRoute, response *core.DetailedResponse, err error) {
	vpnServerRoutePatchModel := &vpcbetav1.VPNServerRoutePatch{
		Name: &[]string{name}[0],
	}
	vpnServerRoutePatchModelAsPatch, asPatchErr := vpnServerRoutePatchModel.AsPatch()
	if asPatchErr != nil {
		panic(asPatchErr)
	}

	updateVPNServerRouteOptions := vpcService.NewUpdateVPNServerRouteOptions(
		vpnServerID,
		vpnServerRouteID,
		vpnServerRoutePatchModelAsPatch,
	)

	vpnServerRoute, response, err = vpcService.UpdateVPNServerRoute(updateVPNServerRouteOptions)
	return
}
func UpdateVPNServer(vpcService *vpcbetav1.VpcbetaV1, vpnServerID, name, ifMatchVPNServer string) (vpnServer *vpcbetav1.VPNServer, response *core.DetailedResponse, err error) {
	vpnServerPatchModel := &vpcbetav1.VPNServerPatch{
		Name: &[]string{name}[0],
	}
	vpnServerPatchModelAsPatch, asPatchErr := vpnServerPatchModel.AsPatch()
	if asPatchErr != nil {
		panic(asPatchErr)
	}
	updateVPNServerOptions := vpcService.NewUpdateVPNServerOptions(
		vpnServerID,
		vpnServerPatchModelAsPatch,
	)
	updateVPNServerOptions.SetIfMatch(ifMatchVPNServer)

	vpnServer, response, err = vpcService.UpdateVPNServer(updateVPNServerOptions)
	return
}
