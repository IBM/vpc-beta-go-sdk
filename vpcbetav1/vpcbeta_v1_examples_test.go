//go:build examples
// +build examples

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
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-beta-go-sdk/vpcbetav1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"os"
	"reflect"
	"strconv"
	"time"
)

var (
	vpcService                        *vpcbetav1.VpcbetaV1
	serviceErr                        error
	configLoaded                      bool = false
	externalConfigFile                     = "../vpc.env"
	backupPolicyID                    string
	backupPolicyPlanID                string
	backupPolicyJobID                 string
	vpcID                             string
	subnetID                          string
	keyID                             string
	imageID                           string
	imageExportJobID                  string
	instanceID                        string
	addressPrefixID                   string
	routingTableID                    string
	routeID                           string
	eth2ID                            string
	floatingIPID                      string
	volumeID                          string
	snapshotID                        string
	volumeAttachmentID                string
	reservedIPID                      string
	reservedIPID2                     string
	ifMatchVolume                     string
	ifMatchBackupPolicy               string
	ifMatchSnapshot                   string
	ifMatchVPNServer                  string
	instanceTemplateID                string
	instanceGroupID                   string
	instanceGroupManagerID            string
	instanceGroupManagerPolicyID      string
	instanceGroupManagerActionID      string
	instanceGroupMembershipID         string
	dedicatedHostGroupID              string
	virtualNetworkInterfaceId         string
	dedicatedHostID                   string
	publicGatewayID                   string
	diskID                            string
	dhID                              string
	securityGroupID                   string
	ikePolicyID                       string
	ipsecPolicyID                     string
	securityGroupRuleID               string
	networkACLID                      string
	targetID                          string
	networkACLRuleID                  string
	vpnGatewayConnectionID            string
	vpnGatewayID                      string
	endpointGatewayID                 string
	placementGroupID                  string
	loadBalancerID                    string
	listenerID                        string
	policyID                          string
	policyRuleID                      string
	poolID                            string
	poolMemberID                      string
	endpointGatewayTargetID           string
	flowLogID                         string
	dhProfile                         string
	operatingSystemName               string
	instanceProfileName               string
	timestamp                         = strconv.FormatInt(tunix, 10)
	tunix                             = time.Now().Unix()
	zone                              *string
	resourceGroupID                   *string
	bareMetalServerProfileName        string
	bareMetalServerId                 string
	bareMetalServerDiskId             string
	bareMetalServerNetworkInterfaceId string
	vpnClientID                       string
	vpnServerRouteID                  string
	vpnServerID                       string
	createdShareID                    string
	createdShare1ID                   string
	createdShareTargetID              string
	createdReplicaShareID             string
	createdReplicaShare1ID            string
	shareProfileName                  string
	createdShareETag                  string
	createdPPSGID                     string
	createdPPSGCRN                    string
	createdPPSGAPID                   string
	EndpointGatewayBindingID          string
)

func skipTest() {
	if !configLoaded {
		Skip("External configuration is not available, skipping tests...")
	}
}

func getName(rtype string) string {
	return "gsdk-" + rtype + "-" + timestamp
}

var _ = Describe(`VpcbetaV1 Examples Tests`, func() {
	Describe(`External configuration`, func() {

		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}
			if err = os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile); err == nil {
				configLoaded = true
			}
			Expect(err).To(BeNil())
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			skipTest()
		})
		It("Successfully construct the service client instance", func() {

			// begin-common

			vpcService, serviceErr = vpcbetav1.NewVpcbetaV1UsingExternalConfig(
				&vpcbetav1.VpcbetaV1Options{
					ServiceName: "vpcbetaint",
				},
			)
			if serviceErr != nil {
				fmt.Println("Gen2 Service creation failed.", serviceErr)
			}

			// end-common

			Expect(vpcService).ToNot(BeNil())
		})
	})
	Describe(`Variable setting`, func() {
		BeforeEach(func() {
			skipTest()
		})
		It("Setting up required variable", func() {
			listSubnetsOptions := &vpcbetav1.ListSubnetsOptions{}

			subnetCollection, _, err := vpcService.ListSubnets(listSubnetsOptions)
			zone = subnetCollection.Subnets[0].Zone.Name
			resourceGroupID = subnetCollection.Subnets[0].ResourceGroup.ID
			Expect(subnetCollection).ToNot(BeNil())
			Expect(zone).ToNot(BeNil())
			Expect(resourceGroupID).ToNot(BeNil())
			Expect(err).To(BeNil())

		})
	})

	Describe(`VpcbetaV1 request examples`, func() {
		BeforeEach(func() {
			skipTest()
		})
		It(`ListVpcs request example`, func() {
			fmt.Println("\nListVpcs() result:")

			// begin-list_vpcs
			listVpcsOptions := &vpcbetav1.ListVpcsOptions{}
			vpcs, response, err := vpcService.ListVpcs(listVpcsOptions)

			// end-list_vpcs
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpcs).ToNot(BeNil())

		})
		It(`CreateVPC request example`, func() {
			fmt.Println("\nCreateVPC() result:")

			classicAccess := true
			manual := "manual"
			// begin-create_vpc

			options := &vpcbetav1.CreateVPCOptions{
				ResourceGroup: &vpcbetav1.ResourceGroupIdentity{
					ID: resourceGroupID,
				},
				Name:                    &[]string{"my-vpc"}[0],
				ClassicAccess:           &classicAccess,
				AddressPrefixManagement: &manual,
			}
			vpc, response, err := vpcService.CreateVPC(options)

			// end-create_vpc
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpc).ToNot(BeNil())
			vpcID = *vpc.ID
		})
		It(`GetVPC request example`, func() {
			fmt.Println("\nGetVPC() result:")
			// begin-get_vpc

			getVpcOptions := &vpcbetav1.GetVPCOptions{
				ID: &vpcID,
			}
			vpc, response, err := vpcService.GetVPC(getVpcOptions)
			// end-get_vpc
			if err != nil {
				panic(err)
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
		It(`UpdateVPC request example`, func() {
			fmt.Println("\nUpdateVPC() result:")
			// begin-update_vpc

			options := &vpcbetav1.UpdateVPCOptions{
				ID: &vpcID,
			}
			vpcPatchModel := &vpcbetav1.VPCPatch{
				Name: &[]string{"my-vpc-modified"}[0],
			}
			vpcPatch, asPatchErr := vpcPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VPCPatch = vpcPatch
			vpc, response, err := vpcService.UpdateVPC(options)

			// end-update_vpc
			if err != nil {
				panic(err)
			}

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpc).ToNot(BeNil())

		})
		It(`GetVPCDefaultNetworkACL request example`, func() {
			fmt.Println("\nGetVPCDefaultNetworkACL() result:")
			// begin-get_vpc_default_network_acl

			options := &vpcbetav1.GetVPCDefaultNetworkACLOptions{}
			options.SetID(vpcID)
			defaultACL, response, err := vpcService.GetVPCDefaultNetworkACL(options)

			// end-get_vpc_default_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultACL).ToNot(BeNil())

		})
		It(`GetVPCDefaultRoutingTable request example`, func() {
			fmt.Println("\nGetVPCDefaultRoutingTable() result:")
			// begin-get_vpc_default_routing_table

			options := vpcService.NewGetVPCDefaultRoutingTableOptions(vpcID)
			defaultRoutingTable, response, err := vpcService.GetVPCDefaultRoutingTable(options)

			// end-get_vpc_default_routing_table

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultRoutingTable).ToNot(BeNil())

		})
		It(`GetVPCDefaultSecurityGroup request example`, func() {
			fmt.Println("\nGetVPCDefaultSecurityGroup() result:")
			// begin-get_vpc_default_security_group

			options := &vpcbetav1.GetVPCDefaultSecurityGroupOptions{}
			options.SetID(vpcID)
			defaultSG, response, err := vpcService.GetVPCDefaultSecurityGroup(options)
			// end-get_vpc_default_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(defaultSG).ToNot(BeNil())

		})
		It(`ListVPCAddressPrefixes request example`, func() {
			fmt.Println("\nListVPCAddressPrefixes() result:")
			// begin-list_vpc_address_prefixes
			listVpcAddressPrefixesOptions := &vpcbetav1.ListVPCAddressPrefixesOptions{}
			listVpcAddressPrefixesOptions.SetVPCID(vpcID)
			addressPrefixes, response, err :=
				vpcService.ListVPCAddressPrefixes(listVpcAddressPrefixesOptions)

			// end-list_vpc_address_prefixes
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefixes).ToNot(BeNil())

		})
		It(`CreateVPCAddressPrefix request example`, func() {
			fmt.Println("\nCreateVPCAddressPrefix() result:")
			// begin-create_vpc_address_prefix

			options := &vpcbetav1.CreateVPCAddressPrefixOptions{}
			options.SetVPCID(vpcID)
			options.SetCIDR("10.0.0.0/24")
			options.SetName("my-address-prefix")
			options.SetZone(&vpcbetav1.ZoneIdentity{
				Name: zone,
			})
			addressPrefix, response, err := vpcService.CreateVPCAddressPrefix(options)
			// end-create_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(addressPrefix).ToNot(BeNil())
			addressPrefixID = *addressPrefix.ID

		})
		It(`GetVPCAddressPrefix request example`, func() {
			fmt.Println("\nGetVPCAddressPrefix() result:")
			// begin-get_vpc_address_prefix

			getVpcAddressPrefixOptions := &vpcbetav1.GetVPCAddressPrefixOptions{}
			getVpcAddressPrefixOptions.SetVPCID(vpcID)
			getVpcAddressPrefixOptions.SetID(addressPrefixID)
			addressPrefix, response, err :=
				vpcService.GetVPCAddressPrefix(getVpcAddressPrefixOptions)

			// end-get_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`UpdateVPCAddressPrefix request example`, func() {
			fmt.Println("\nUpdateVPCAddressPrefix() result:")
			isDefault := true
			// begin-update_vpc_address_prefix
			options := &vpcbetav1.UpdateVPCAddressPrefixOptions{}
			options.SetVPCID(vpcID)
			options.SetID(addressPrefixID)
			addressPrefixPatchModel := &vpcbetav1.AddressPrefixPatch{}
			addressPrefixPatchModel.Name = &[]string{"my-address-prefix-updated"}[0]
			addressPrefixPatchModel.IsDefault = &isDefault
			addressPrefixPatch, asPatchErr := addressPrefixPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.AddressPrefixPatch = addressPrefixPatch
			addressPrefix, response, err := vpcService.UpdateVPCAddressPrefix(options)

			// end-update_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(addressPrefix).ToNot(BeNil())

		})
		It(`ListVPCRoutingTables request example`, func() {
			fmt.Println("\nListVPCRoutingTables() result:")
			// begin-list_vpc_routing_tables

			options := vpcService.NewListVPCRoutingTablesOptions(vpcID)
			routingTableCollection, response, err := vpcService.ListVPCRoutingTables(options)

			// end-list_vpc_routing_tables
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTableCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoutingTable request example`, func() {
			fmt.Println("\nCreateVPCRoutingTable() result:")
			routeName := "my-route"
			action := "delegate"
			// begin-create_vpc_routing_table
			routePrototypeModel := &vpcbetav1.RoutePrototype{
				Action: &action,
				NextHop: &vpcbetav1.RoutePrototypeNextHopRouteNextHopPrototypeRouteNextHopIP{
					Address: &[]string{"192.168.3.4"}[0],
				},
				Name:        &routeName,
				Destination: &[]string{"192.168.3.0/24"}[0],
				Zone: &vpcbetav1.ZoneIdentityByName{
					Name: zone,
				},
			}
			name := "my-routing-table"
			options := &vpcbetav1.CreateVPCRoutingTableOptions{
				VPCID:  &vpcID,
				Name:   &name,
				Routes: []vpcbetav1.RoutePrototype{*routePrototypeModel},
			}
			routingTable, response, err := vpcService.CreateVPCRoutingTable(options)
			// end-create_vpc_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())
			routingTableID = *routingTable.ID
		})
		It(`GetVPCRoutingTable request example`, func() {
			fmt.Println("\nGetVPCRoutingTable() result:")
			// begin-get_vpc_routing_table

			options := &vpcbetav1.GetVPCRoutingTableOptions{
				VPCID: &vpcID,
				ID:    &routingTableID,
			}
			routingTable, response, err := vpcService.GetVPCRoutingTable(options)
			// end-get_vpc_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())
		})
		It(`UpdateVPCRoutingTable request example`, func() {
			fmt.Println("\nUpdateVPCRoutingTable() result:")
			// begin-update_vpc_routing_table

			name := "my-routing-table"
			routingTablePatchModel := &vpcbetav1.RoutingTablePatch{
				Name: &name,
			}
			routingTablePatchModelAsPatch, asPatchErr := routingTablePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcbetav1.UpdateVPCRoutingTableOptions{
				VPCID:             &vpcID,
				ID:                &routingTableID,
				RoutingTablePatch: routingTablePatchModelAsPatch,
			}
			routingTable, response, err := vpcService.UpdateVPCRoutingTable(options)

			// end-update_vpc_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ListVPCRoutingTableRoutes request example`, func() {
			fmt.Println("\nListVPCRoutingTableRoutes() result:")
			// begin-list_vpc_routing_table_routes

			options := vpcService.NewListVPCRoutingTableRoutesOptions(
				vpcID,
				routingTableID,
			)
			routeCollection, response, err := vpcService.ListVPCRoutingTableRoutes(options)

			// end-list_vpc_routing_table_routes
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routeCollection).ToNot(BeNil())

		})
		It(`CreateVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nCreateVPCRoutingTableRoute() result:")
			destination := "192.168.77.0/24"
			address := "192.168.3.7"
			// begin-create_vpc_routing_table_route
			zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
				Name: zone,
			}
			options := &vpcbetav1.CreateVPCRoutingTableRouteOptions{
				VPCID:          &vpcID,
				RoutingTableID: &routingTableID,
				Destination:    &destination,
				Zone:           zoneIdentityModel,
				NextHop: &vpcbetav1.RoutePrototypeNextHopRouteNextHopPrototypeRouteNextHopIP{
					Address: &address,
				},
			}
			route, response, err := vpcService.CreateVPCRoutingTableRoute(options)

			// end-create_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(route).ToNot(BeNil())
			routeID = *route.ID
		})
		It(`GetVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nGetVPCRoutingTableRoute() result:")
			// begin-get_vpc_routing_table_route

			options := vpcService.NewGetVPCRoutingTableRouteOptions(
				vpcID,
				routingTableID,
				routeID,
			)
			route, response, err := vpcService.GetVPCRoutingTableRoute(options)

			// end-get_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`UpdateVPCRoutingTableRoute request example`, func() {
			fmt.Println("\nUpdateVPCRoutingTableRoute() result:")
			// begin-update_vpc_routing_table_route

			name := "my-route-updated"
			routePatchModel := &vpcbetav1.RoutePatch{
				Name: &name,
			}
			routePatchModelAsPatch, asPatchErr := routePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcbetav1.UpdateVPCRoutingTableRouteOptions{
				VPCID:          &vpcID,
				RoutingTableID: &routingTableID,
				ID:             &routeID,
				RoutePatch:     routePatchModelAsPatch,
			}
			route, response, err := vpcService.UpdateVPCRoutingTableRoute(options)

			// end-update_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(route).ToNot(BeNil())

		})
		It(`ListSubnets request example`, func() {
			fmt.Println("\nListSubnets() result:")
			// begin-list_subnets

			options := &vpcbetav1.ListSubnetsOptions{}
			subnets, response, err := vpcService.ListSubnets(options)

			// end-list_subnets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnets).ToNot(BeNil())

		})
		It(`CreateSubnet request example`, func() {
			fmt.Println("\nCreateSubnet() result:")
			cidrBlock := "10.0.1.0/24"
			// begin-create_subnet

			options := &vpcbetav1.CreateSubnetOptions{}
			options.SetSubnetPrototype(&vpcbetav1.SubnetPrototype{
				Ipv4CIDRBlock: &cidrBlock,
				Name:          &[]string{"my-subnet"}[0],
				VPC: &vpcbetav1.VPCIdentity{
					ID: &vpcID,
				},
				Zone: &vpcbetav1.ZoneIdentity{
					Name: zone,
				},
			})
			subnet, response, err := vpcService.CreateSubnet(options)

			// end-create_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(subnet).ToNot(BeNil())
			subnetID = *subnet.ID
		})
		It(`GetSubnet request example`, func() {
			fmt.Println("\nGetSubnet() result:")
			// begin-get_subnet

			options := &vpcbetav1.GetSubnetOptions{}
			options.SetID(subnetID)
			subnet, response, err := vpcService.GetSubnet(options)

			// end-get_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
		It(`UpdateSubnet request example`, func() {
			fmt.Println("\nUpdateSubnet() result:")
			name := getName("subnet")
			networkAclId := &networkACLID
			routingTableId := &[]string{""}[0]
			// begin-update_subnet

			options := &vpcbetav1.UpdateSubnetOptions{}
			options.SetID(subnetID)
			subnetPatchModel := &vpcbetav1.SubnetPatch{}
			subnetPatchModel.Name = &name
			subnetPatchModel.NetworkACL = &vpcbetav1.NetworkACLIdentity{
				ID: networkAclId,
			}
			routingTableIdentityModel := new(vpcbetav1.RoutingTableIdentityByID)
			routingTableIdentityModel.ID = routingTableId
			subnetPatchModel.RoutingTable = routingTableIdentityModel
			subnetPatch, asPatchErr := subnetPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SubnetPatch = subnetPatch
			subnet, response, err := vpcService.UpdateSubnet(options)

			// end-update_subnet

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(subnet).ToNot(BeNil())

		})
		It(`ReplaceSubnetNetworkACL request example`, func() {
			fmt.Println("\nReplaceSubnetNetworkACL() result:")
			vpcIDentityModel := &vpcbetav1.VPCIdentityByID{
				ID: &vpcID,
			}
			networkACLPrototypeModel := &vpcbetav1.NetworkACLPrototypeNetworkACLByRules{
				Name: &[]string{"my-network-acl"}[0],
				VPC:  vpcIDentityModel,
			}
			createNetworkACLOptions := vpcService.NewCreateNetworkACLOptions(networkACLPrototypeModel)

			networkACL, _, _ := vpcService.CreateNetworkACL(createNetworkACLOptions)
			Expect(networkACL).ToNot(BeNil())
			networkACLID := networkACL.ID
			// begin-replace_subnet_network_acl

			options := &vpcbetav1.ReplaceSubnetNetworkACLOptions{}
			options.SetID(subnetID)
			options.SetNetworkACLIdentity(&vpcbetav1.NetworkACLIdentity{
				ID: networkACLID,
			})
			networkACL, response, err := vpcService.ReplaceSubnetNetworkACL(options)

			// end-replace_subnet_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`GetSubnetNetworkACL request example`, func() {
			fmt.Println("\nGetSubnetNetworkACL() result:")
			// begin-get_subnet_network_acl

			options := &vpcbetav1.GetSubnetNetworkACLOptions{}
			options.SetID(subnetID)
			acls, response, err := vpcService.GetSubnetNetworkACL(options)

			// end-get_subnet_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(acls).ToNot(BeNil())

		})
		It(`SetSubnetPublicGateway request example`, func() {
			fmt.Println("\nSetSubnetPublicGateway() result:")
			vpcIDentityModel := &vpcbetav1.VPCIdentityByID{
				ID: &vpcID,
			}

			zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
				Name: zone,
			}

			createPublicGatewayOptions := vpcService.NewCreatePublicGatewayOptions(
				vpcIDentityModel,
				zoneIdentityModel,
			)

			publicGateway, _, err := vpcService.CreatePublicGateway(createPublicGatewayOptions)
			if err != nil {
				panic(err)
			}
			Expect(publicGateway).ToNot(BeNil())

			// begin-set_subnet_public_gateway

			options := &vpcbetav1.SetSubnetPublicGatewayOptions{}
			options.SetID(subnetID)
			options.SetPublicGatewayIdentity(&vpcbetav1.PublicGatewayIdentity{
				ID: publicGateway.ID,
			})
			publicGateway, response, err := vpcService.SetSubnetPublicGateway(options)
			// end-set_subnet_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())
		})
		It(`GetSubnetPublicGateway request example`, func() {
			fmt.Println("\nGetSubnetPublicGateway() result:")
			// begin-get_subnet_public_gateway

			options := &vpcbetav1.GetSubnetPublicGatewayOptions{}
			options.SetID(subnetID)
			publicGateway, response, err := vpcService.GetSubnetPublicGateway(options)

			// end-get_subnet_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})

		It(`UnsetSubnetPublicGateway request example`, func() {
			// begin-unset_subnet_public_gateway

			options := vpcService.NewUnsetSubnetPublicGatewayOptions(
				subnetID,
			)

			response, err := vpcService.UnsetSubnetPublicGateway(options)

			// end-unset_subnet_public_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nUnsetSubnetPublicGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ReplaceSubnetRoutingTable request example`, func() {
			fmt.Println("\nReplaceSubnetRoutingTable() result:")
			// begin-replace_subnet_routing_table

			routingTableIdentityModel := &vpcbetav1.RoutingTableIdentityByID{
				ID: &routingTableID,
			}
			replaceSubnetRoutingTableOptions := vpcService.NewReplaceSubnetRoutingTableOptions(
				subnetID,
				routingTableIdentityModel,
			)
			routingTable, response, err := vpcService.ReplaceSubnetRoutingTable(
				replaceSubnetRoutingTableOptions,
			)

			// end-replace_subnet_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`GetSubnetRoutingTable request example`, func() {
			fmt.Println("\nGetSubnetRoutingTable() result:")
			// begin-get_subnet_routing_table
			options := vpcService.NewGetSubnetRoutingTableOptions(subnetID)
			routingTable, response, err := vpcService.GetSubnetRoutingTable(options)

			// end-get_subnet_routing_table
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(routingTable).ToNot(BeNil())

		})
		It(`ListSubnetReservedIps request example`, func() {
			fmt.Println("\nListSubnetReservedIps() result:")
			// begin-list_subnet_reserved_ips

			options := vpcService.NewListSubnetReservedIpsOptions(subnetID)
			reservedIPCollection, response, err := vpcService.ListSubnetReservedIps(options)

			// end-list_subnet_reserved_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPCollection).ToNot(BeNil())

		})
		It(`CreateSubnetReservedIP request example`, func() {
			fmt.Println("\nCreateSubnetReservedIP() result:")
			name := getName("subnetRip")
			// begin-create_subnet_reserved_ip

			options := vpcService.NewCreateSubnetReservedIPOptions(subnetID)
			options.Name = &name
			reservedIP, response, err := vpcService.CreateSubnetReservedIP(options)

			// end-create_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())
			reservedIPID = *reservedIP.ID

		})
		It(`GetSubnetReservedIP request example`, func() {
			fmt.Println("\nGetSubnetReservedIP() result:")
			// begin-get_subnet_reserved_ip

			options := vpcService.NewGetSubnetReservedIPOptions(subnetID, reservedIPID)
			reservedIP, response, err := vpcService.GetSubnetReservedIP(options)

			// end-get_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`UpdateSubnetReservedIP request example`, func() {
			fmt.Println("\nUpdateSubnetReservedIP() result:")
			name := getName("subnetRip")
			// begin-update_subnet_reserved_ip

			options := &vpcbetav1.UpdateSubnetReservedIPOptions{}

			patchBody := new(vpcbetav1.ReservedIPPatch)
			patchBody.Name = &name
			patchBody.AutoDelete = &[]bool{true}[0]
			reservedIPPatch, asPatchErr := patchBody.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SetReservedIPPatch(reservedIPPatch)
			options.SetID(reservedIPID)
			options.SetSubnetID(subnetID)
			reservedIP, response, err := vpcService.UpdateSubnetReservedIP(options)

			// end-update_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`DeleteSubnetReservedIP request example`, func() {
			// begin-delete_subnet_reserved_ip
			deleteSubnetReservedIPOptions := vpcService.NewDeleteSubnetReservedIPOptions(
				subnetID,
				reservedIPID,
			)

			response, err := vpcService.DeleteSubnetReservedIP(deleteSubnetReservedIPOptions)

			// end-delete_subnet_reserved_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSubnetReservedIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
			// create another reserved ip for endpoint gateway
			name := getName("subnetRip")
			options := vpcService.NewCreateSubnetReservedIPOptions(subnetID)
			options.Name = &name
			reservedIP, response, err := vpcService.CreateSubnetReservedIP(options)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())
			reservedIPID = *reservedIP.ID
		})
		It(`ListImages request example`, func() {
			fmt.Println("\nListImages() result:")
			// begin-list_images
			options := &vpcbetav1.ListImagesOptions{}
			options.SetVisibility("private")
			images, response, err := vpcService.ListImages(options)

			// end-list_images
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(images).ToNot(BeNil())
		})
		It(`CreateImage request example`, func() {
			fmt.Println("\nCreateImage() result:")
			name := getName("image")
			// begin-create_image

			operatingSystemIdentityModel := &vpcbetav1.OperatingSystemIdentityByName{
				Name: &[]string{"debian-9-amd64"}[0],
			}

			options := &vpcbetav1.CreateImageOptions{}
			cosID := "cos://us-south/my-bucket/my-image.qcow2"
			options.SetImagePrototype(&vpcbetav1.ImagePrototype{
				Name: &name,
				File: &vpcbetav1.ImageFilePrototype{
					Href: &cosID,
				},
				OperatingSystem: operatingSystemIdentityModel,
			})
			image, response, err := vpcService.CreateImage(options)

			// end-create_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(image).ToNot(BeNil())
			imageID = *image.ID
		})
		It(`GetImage request example`, func() {
			fmt.Println("\nGetImage() result:")
			// begin-get_image
			options := &vpcbetav1.GetImageOptions{}
			options.SetID(imageID)
			image, response, err := vpcService.GetImage(options)
			// end-get_image
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
		It(`UpdateImage request example`, func() {
			fmt.Println("\nUpdateImage() result:")
			name := getName("image")
			// begin-update_image

			options := &vpcbetav1.UpdateImageOptions{}
			options.SetID(imageID)
			imagePatchModel := &vpcbetav1.ImagePatch{
				Name: &name,
			}
			imagePatch, asPatchErr := imagePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.ImagePatch = imagePatch
			image, response, err := vpcService.UpdateImage(options)

			// end-update_image
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(image).ToNot(BeNil())

		})
		It(`DeprecateImage request example`, func() {
			fmt.Println("\nDeprecateImage result:")
			// begin-deprecate_image

			deprecateImageOptions := vpcService.NewDeprecateImageOptions(
				imageID,
			)

			response, err := vpcService.DeprecateImage(deprecateImageOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeprecateImage(): %d\n", response.StatusCode)
			}

			// end-deprecate_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ObsoleteImage request example`, func() {
			fmt.Println("\nObsoleteImage result:")
			// begin-obsolete_image

			obsoleteImageOptions := vpcService.NewObsoleteImageOptions(
				imageID,
			)

			response, err := vpcService.ObsoleteImage(obsoleteImageOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from ObsoleteImage(): %d\n", response.StatusCode)
			}

			// end-obsolete_image

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
		It(`ListImageExportJobs request example`, func() {
			fmt.Println("\nListImageExportJobs() result:")
			// begin-list_image_export_jobs

			listImageExportJobsOptions := vpcService.NewListImageExportJobsOptions(
				imageID,
			)

			imageExportJobUnpaginatedCollection, response, err := vpcService.ListImageExportJobs(listImageExportJobsOptions)
			if err != nil {
				panic(err)
			}

			// end-list_image_export_jobs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageExportJobUnpaginatedCollection).ToNot(BeNil())
		})
		It(`CreateImageExportJob request example`, func() {
			fmt.Println("\nCreateImageExportJob() result:")
			//name := getName("image-export-job")
			// begin-create_image_export_job

			cloudObjectStorageBucketIdentityModel := &vpcbetav1.CloudObjectStorageBucketIdentityCloudObjectStorageBucketIdentityByName{
				Name: &[]string{"bucket-27200-lwx4cfvcue"}[0],
			}

			createImageExportJobOptions := &vpcbetav1.CreateImageExportJobOptions{
				ImageID:       &imageID,
				StorageBucket: cloudObjectStorageBucketIdentityModel,
				Format:        &[]string{"qcow2"}[0],
				Name:          &[]string{"my-image-export-job"}[0],
			}

			imageExportJob, response, err := vpcService.CreateImageExportJob(createImageExportJobOptions)
			if err != nil {
				panic(err)
			}

			// end-create_image_export_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(imageExportJob).ToNot(BeNil())
			imageExportJobID = *imageExportJob.ID
		})
		It(`GetImageExportJob request example`, func() {
			fmt.Println("\nGetImageExportJob() result:")
			// begin-get_image_export_job

			getImageExportJobOptions := &vpcbetav1.GetImageExportJobOptions{
				ImageID: &imageID,
				ID:      &imageExportJobID,
			}

			imageExportJob, response, err := vpcService.GetImageExportJob(getImageExportJobOptions)
			if err != nil {
				panic(err)
			}

			// end-get_image_export_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageExportJob).ToNot(BeNil())
		})
		It(`UpdateImageExportJob request example`, func() {
			fmt.Println("\nUpdateImageExportJob() result:")
			//name := getName("image-export-job-updated")
			// begin-update_image_export_job

			imageExportJobPatchModel := &vpcbetav1.ImageExportJobPatch{
				Name: &[]string{"image-export-job-updated"}[0],
			}
			imageExportJobPatchModelAsPatch, _ := imageExportJobPatchModel.AsPatch()

			updateImageExportJobOptions := &vpcbetav1.UpdateImageExportJobOptions{
				ImageID:             &imageID,
				ID:                  &imageExportJobID,
				ImageExportJobPatch: imageExportJobPatchModelAsPatch,
			}

			imageExportJob, response, err := vpcService.UpdateImageExportJob(updateImageExportJobOptions)
			if err != nil {
				panic(err)
			}

			// end-update_image_export_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(imageExportJob).ToNot(BeNil())
		})
		It(`ListOperatingSystems request example`, func() {
			fmt.Println("\nListOperatingSystems() result:")
			// begin-list_operating_systems

			options := &vpcbetav1.ListOperatingSystemsOptions{}
			operatingSystems, response, err := vpcService.ListOperatingSystems(options)

			// end-list_operating_systems
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystems).ToNot(BeNil())
			operatingSystemName = *operatingSystems.OperatingSystems[0].Name
		})
		It(`GetOperatingSystem request example`, func() {
			fmt.Println("\nGetOperatingSystem() result:")
			// begin-get_operating_system

			options := &vpcbetav1.GetOperatingSystemOptions{}
			options.SetName(operatingSystemName)
			operatingSystem, response, err := vpcService.GetOperatingSystem(options)

			// end-get_operating_system
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(operatingSystem).ToNot(BeNil())

		})
		It(`ListKeys request example`, func() {
			fmt.Println("\nListKeys() result:")
			// begin-list_keys

			listKeysOptions := &vpcbetav1.ListKeysOptions{}
			keys, response, err := vpcService.ListKeys(listKeysOptions)

			// end-list_keys

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(keys).ToNot(BeNil())

		})
		It(`CreateKey request example`, func() {
			fmt.Println("\nCreateKey() result:")
			name := getName("sshkey")
			publicKey := "AAAAB3NzaC1yc2EAAAADAQABAAABAQDDGe50Bxa5T5NDddrrtbx2Y4/VGbiCgXqnBsYToIUKoFSHTQl5IX3PasGnneKanhcLwWz5M5MoCRvhxTp66NKzIfAz7r+FX9rxgR+ZgcM253YAqOVeIpOU408simDZKriTlN8kYsXL7P34tsWuAJf4MgZtJAQxous/2byetpdCv8ddnT4X3ltOg9w+LqSCPYfNivqH00Eh7S1Ldz7I8aw5WOp5a+sQFP/RbwfpwHp+ny7DfeIOokcuI42tJkoBn7UsLTVpCSmXr2EDRlSWe/1M/iHNRBzaT3CK0+SwZWd2AEjePxSnWKNGIEUJDlUYp7hKhiQcgT5ZAnWU121oc5En"

			// begin-create_key
			options := &vpcbetav1.CreateKeyOptions{}
			options.SetName(name)
			options.SetPublicKey(publicKey)
			key, response, err := vpcService.CreateKey(options)

			// end-create_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(key).ToNot(BeNil())
			keyID = *key.ID
		})
		It(`GetKey request example`, func() {
			fmt.Println("\nGetKey() result:")
			// begin-get_key

			getKeyOptions := &vpcbetav1.GetKeyOptions{}
			getKeyOptions.SetID(keyID)
			key, response, err := vpcService.GetKey(getKeyOptions)

			// end-get_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
		It(`UpdateKey request example`, func() {
			fmt.Println("\nUpdateKey() result:")
			// begin-update_key

			options := &vpcbetav1.UpdateKeyOptions{}
			options.SetID(keyID)
			keyPatchModel := &vpcbetav1.KeyPatch{
				Name: &[]string{"my-key-modified"}[0],
			}
			keyPatch, asPatchErr := keyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.KeyPatch = keyPatch
			key, response, err := vpcService.UpdateKey(options)

			// end-update_key

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(key).ToNot(BeNil())

		})
		It(`ListFloatingIps request example`, func() {
			fmt.Println("\nListFloatingIps() result:")
			// begin-list_floating_ips
			listFloatingIpsOptions := vpcService.NewListFloatingIpsOptions()
			floatingIPs, response, err :=
				vpcService.ListFloatingIps(listFloatingIpsOptions)

			// end-list_floating_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPs).ToNot(BeNil())

		})
		It(`CreateFloatingIP request example`, func() {
			fmt.Println("\nCreateFloatingIP() result:")
			name := getName("floatingIP")
			// begin-create_floating_ip

			options := &vpcbetav1.CreateFloatingIPOptions{}
			options.SetFloatingIPPrototype(&vpcbetav1.FloatingIPPrototype{
				Name: &name,
				Zone: &vpcbetav1.ZoneIdentity{
					Name: zone,
				},
			})
			floatingIP, response, err := vpcService.CreateFloatingIP(options)

			// end-create_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())
			floatingIPID = *floatingIP.ID
		})
		It(`GetFloatingIP request example`, func() {
			fmt.Println("\nGetFloatingIP() result:")
			// begin-get_floating_ip

			options := vpcService.NewGetFloatingIPOptions(floatingIPID)
			floatingIP, response, err := vpcService.GetFloatingIP(options)

			// end-get_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`UpdateFloatingIP request example`, func() {
			name := getName("fip")
			fmt.Println("\nUpdateFloatingIP() result:")
			// begin-update_floating_ip

			floatingIPPatchModel := &vpcbetav1.FloatingIPPatch{
				Name: &name,
			}
			floatingIPPatchModelAsPatch, asPatchErr := floatingIPPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}

			updateFloatingIPOptions := vpcService.NewUpdateFloatingIPOptions(
				floatingIPID,
				floatingIPPatchModelAsPatch,
			)

			floatingIP, response, err := vpcService.UpdateFloatingIP(updateFloatingIPOptions)

			// end-update_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`ListVolumes request example`, func() {
			fmt.Println("\nListVolumes() result:")
			// begin-list_volumes

			options := &vpcbetav1.ListVolumesOptions{}
			volumes, response, err := vpcService.ListVolumes(options)

			// end-list_volumes
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumes).ToNot(BeNil())

		})
		It(`CreateVolume request example`, func() {
			fmt.Println("\nCreateVolume() result:")
			name := getName("vol")
			// begin-create_volume
			options := &vpcbetav1.CreateVolumeOptions{}
			options.SetVolumePrototype(&vpcbetav1.VolumePrototype{
				Capacity: &[]int64{100}[0],
				Zone: &vpcbetav1.ZoneIdentity{
					Name: zone,
				},
				Profile: &vpcbetav1.VolumeProfileIdentity{
					Name: &[]string{"general-purpose"}[0],
				},
				Name: &name,
			})
			volume, response, err := vpcService.CreateVolume(options)
			// end-create_volume
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volume).ToNot(BeNil())
			volumeID = *volume.ID
		})
		It(`GetVolume request example`, func() {
			fmt.Println("\nGetVolume() result:")
			// begin-get_volume

			options := &vpcbetav1.GetVolumeOptions{}
			options.SetID(volumeID)
			volume, response, err := vpcService.GetVolume(options)

			// end-get_volume
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())
			ifMatchVolume = response.GetHeaders()["Etag"][0]
		})
		It(`UpdateVolume request example`, func() {
			fmt.Println("\nUpdateVolume() result:")
			name := getName("vol")
			userTags := []string{"usertag-vol-1"}
			// begin-update_volume

			options := &vpcbetav1.UpdateVolumeOptions{}
			options.SetID(volumeID)
			options.SetIfMatch(ifMatchVolume)
			volumePatchModel := &vpcbetav1.VolumePatch{
				Name:     &name,
				UserTags: userTags,
			}
			volumePatch, asPatchErr := volumePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VolumePatch = volumePatch
			volume, response, err := vpcService.UpdateVolume(options)
			// end-update_volume
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volume).ToNot(BeNil())

		})
		It(`ListInstanceProfiles request example`, func() {
			fmt.Println("\nListInstanceProfiles() result:")
			// begin-list_instance_profiles

			options := &vpcbetav1.ListInstanceProfilesOptions{}
			profiles, response, err := vpcService.ListInstanceProfiles(options)

			// end-list_instance_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())
			instanceProfileName = *profiles.Profiles[0].Name
		})
		It(`GetInstanceProfile request example`, func() {
			fmt.Println("\nGetInstanceProfile() result:")
			// begin-get_instance_profile

			options := &vpcbetav1.GetInstanceProfileOptions{}
			options.SetName(instanceProfileName)
			profile, response, err := vpcService.GetInstanceProfile(options)
			// end-get_instance_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ListInstanceTemplates request example`, func() {
			fmt.Println("\nListInstanceTemplates() result:")
			// begin-list_instance_templates

			options := &vpcbetav1.ListInstanceTemplatesOptions{}
			instanceTemplates, response, err := vpcService.ListInstanceTemplates(options)

			// end-list_instance_templates
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplates).ToNot(BeNil())

		})
		It(`CreateInstanceTemplate request example`, func() {
			fmt.Println("\nCreateInstanceTemplate() result:")
			name := getName("template")
			instanceProfile := []string{"bx2d-2x8"}[0]
			// begin-create_instance_template

			options := &vpcbetav1.CreateInstanceTemplateOptions{}
			options.SetInstanceTemplatePrototype(&vpcbetav1.InstanceTemplatePrototype{
				Name: &name,
				Image: &vpcbetav1.ImageIdentity{
					ID: &imageID,
				},
				Profile: &vpcbetav1.InstanceProfileIdentity{
					Name: &instanceProfile,
				},
				Zone: &vpcbetav1.ZoneIdentity{
					Name: zone,
				},
				PrimaryNetworkInterface: &vpcbetav1.NetworkInterfacePrototype{
					Subnet: &vpcbetav1.SubnetIdentity{
						ID: &subnetID,
					},
				},
				Keys: []vpcbetav1.KeyIdentityIntf{
					&vpcbetav1.KeyIdentity{
						ID: &keyID,
					},
				},
				VPC: &vpcbetav1.VPCIdentity{
					ID: &vpcID,
				},
			})
			instanceTemplate, response, err := vpcService.CreateInstanceTemplate(options)

			// end-create_instance_template
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceTemplate).ToNot(BeNil())
			instanceTemplateID = *instanceTemplate.(*vpcbetav1.InstanceTemplate).ID
		})
		It(`GetInstanceTemplate request example`, func() {
			fmt.Println("\nGetInstanceTemplate() result:")
			// begin-get_instance_template

			options := &vpcbetav1.GetInstanceTemplateOptions{}
			options.SetID(instanceTemplateID)
			instanceTemplate, response, err := vpcService.GetInstanceTemplate(options)

			// end-get_instance_template
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`UpdateInstanceTemplate request example`, func() {
			fmt.Println("\nUpdateInstanceTemplate() result:")
			name := getName("template")
			// begin-update_instance_template

			options := &vpcbetav1.UpdateInstanceTemplateOptions{}
			options.SetID(instanceTemplateID)
			instanceTemplatePatchModel := &vpcbetav1.InstanceTemplatePatch{
				Name: &name,
			}
			instanceTemplatePatch, asPatchErr := instanceTemplatePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceTemplatePatch = instanceTemplatePatch
			instanceTemplate, response, err := vpcService.UpdateInstanceTemplate(options)

			// end-update_instance_template
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceTemplate).ToNot(BeNil())

		})
		It(`ListInstances request example`, func() {
			fmt.Println("\nListInstances() result:")
			// begin-list_instances

			options := &vpcbetav1.ListInstancesOptions{}
			instances, response, err := vpcService.ListInstances(options)

			// end-list_instances
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instances).ToNot(BeNil())

		})
		It(`CreateInstance request example`, func() {
			fmt.Println("\nCreateInstance() result:")
			crn := "crn:[...]"
			// begin-create_instance
			keyIDentityModel := &vpcbetav1.KeyIdentityByID{
				ID: &keyID,
			}
			instanceProfileIdentityModel := &vpcbetav1.InstanceProfileIdentity{
				Name: &[]string{"bx2d-2x8"}[0],
			}
			encryptionKeyIdentityModel := &vpcbetav1.EncryptionKeyIdentityByCRN{
				CRN: &crn,
			}
			volumeProfileIdentityModel := &vpcbetav1.VolumeProfileIdentityByName{
				Name: &[]string{"5iops-tier"}[0],
			}
			volumeAttachmentPrototypeVolumeModel := &vpcbetav1.VolumeAttachmentPrototypeVolumeVolumePrototypeInstanceContext{
				Name:          &[]string{"my-instance-modified"}[0],
				Capacity:      &[]int64{100}[0],
				EncryptionKey: encryptionKeyIdentityModel,
				Profile:       volumeProfileIdentityModel,
			}
			volumeAttachmentPrototypeModel := &vpcbetav1.VolumeAttachmentPrototype{
				Volume: volumeAttachmentPrototypeVolumeModel,
			}
			vpcIDentityModel := &vpcbetav1.VPCIdentityByID{
				ID: &vpcID,
			}
			imageIDentityModel := &vpcbetav1.ImageIdentityByID{
				ID: &imageID,
			}
			subnetIDentityModel := &vpcbetav1.SubnetIdentityByID{
				ID: &subnetID,
			}
			networkInterfacePrototypeModel := &vpcbetav1.NetworkInterfacePrototype{
				Name:   &[]string{"my-instance-modified"}[0],
				Subnet: subnetIDentityModel,
			}
			zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
				Name: zone,
			}
			instancePrototypeModel := &vpcbetav1.InstancePrototypeInstanceByImage{
				Keys:                    []vpcbetav1.KeyIdentityIntf{keyIDentityModel},
				Name:                    &[]string{"my-instance-modified"}[0],
				Profile:                 instanceProfileIdentityModel,
				VolumeAttachments:       []vpcbetav1.VolumeAttachmentPrototype{*volumeAttachmentPrototypeModel},
				VPC:                     vpcIDentityModel,
				Image:                   imageIDentityModel,
				PrimaryNetworkInterface: networkInterfacePrototypeModel,
				Zone:                    zoneIdentityModel,
				ConfidentialComputeMode: &[]string{"sgx"}[0],
				EnableSecureBoot:        &[]bool{true}[0],
			}
			createInstanceOptions := vpcService.NewCreateInstanceOptions(
				instancePrototypeModel,
			)
			instance, response, err := vpcService.CreateInstance(createInstanceOptions)

			// end-create_instance
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instance).ToNot(BeNil())
			instanceID = *instance.ID
		})
		It(`GetInstance request example`, func() {
			fmt.Println("\nGetInstance() result:")
			// begin-get_instance

			options := &vpcbetav1.GetInstanceOptions{}
			options.SetID(instanceID)
			instance, response, err := vpcService.GetInstance(options)

			// end-get_instance
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`UpdateInstance request example`, func() {
			fmt.Println("\nUpdateInstance() result:")
			// begin-update_instance

			options := &vpcbetav1.UpdateInstanceOptions{
				ID: &instanceID,
			}
			instancePatchModel := &vpcbetav1.InstancePatch{
				Name: &[]string{"my-instance-modified"}[0],
			}
			instancePatch, asPatchErr := instancePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstancePatch = instancePatch
			instance, response, err := vpcService.UpdateInstance(options)

			// end-update_instance
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instance).ToNot(BeNil())

		})
		It(`GetInstanceInitialization request example`, func() {
			fmt.Println("\nGetInstanceInitialization() result:")
			// begin-get_instance_initialization
			options := &vpcbetav1.GetInstanceInitializationOptions{}
			options.SetID(instanceID)
			initData, response, err := vpcService.GetInstanceInitialization(options)

			// end-get_instance_initialization
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(initData).ToNot(BeNil())

		})
		It(`CreateInstanceAction request example`, func() {
			fmt.Println("\nCreateInstanceAction() result:")
			// begin-create_instance_action

			options := &vpcbetav1.CreateInstanceActionOptions{}
			options.SetInstanceID(instanceID)
			options.SetType("stop")
			action, response, err := vpcService.CreateInstanceAction(options)

			// end-create_instance_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(action).ToNot(BeNil())

		})
		It(`CreateInstanceConsoleAccessToken request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nCreateInstanceConsoleAccessToken() result:")
			// begin-create_instance_console_access_token
			options := &vpcbetav1.CreateInstanceConsoleAccessTokenOptions{
				InstanceID:  &instanceID,
				ConsoleType: &[]string{"serial"}[0],
			}

			instanceConsoleAccessToken, response, err :=
				vpcService.CreateInstanceConsoleAccessToken(options)

			// end-create_instance_console_access_token
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceConsoleAccessToken).ToNot(BeNil())

		})
		It(`ListInstanceDisks request example`, func() {
			fmt.Println("\nListInstanceDisks() result:")
			// begin-list_instance_disks

			listInstanceDisksOptions := vpcService.NewListInstanceDisksOptions(
				instanceID,
			)
			instanceDisksCollection, response, err :=
				vpcService.ListInstanceDisks(listInstanceDisksOptions)

			// end-list_instance_disks
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisksCollection).ToNot(BeNil())
			diskID = *instanceDisksCollection.Disks[0].ID
		})
		It(`GetInstanceDisk request example`, func() {
			fmt.Println("\nGetInstanceDisk() result:")
			// begin-get_instance_disk

			options := vpcService.NewGetInstanceDiskOptions(
				instanceID,
				diskID,
			)
			instanceDisk, response, err := vpcService.GetInstanceDisk(options)

			// end-get_instance_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
		It(`UpdateInstanceDisk request example`, func() {
			fmt.Println("\nUpdateInstanceDisk() result:")
			name := getName("disk")
			// begin-update_instance_disk

			instanceDiskPatchModel := &vpcbetav1.InstanceDiskPatch{
				Name: &name,
			}
			instanceDiskPatchModelAsPatch, asPatchErr := instanceDiskPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := vpcService.NewUpdateInstanceDiskOptions(
				instanceID,
				diskID,
				instanceDiskPatchModelAsPatch,
			)
			instanceDisk, response, err := vpcService.UpdateInstanceDisk(options)

			// end-update_instance_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceDisk).ToNot(BeNil())

		})
		It(`ListInstanceNetworkInterfaces request example`, func() {
			fmt.Println("\nListInstanceNetworkInterfaces() result:")
			// begin-list_instance_network_interfaces

			options := &vpcbetav1.ListInstanceNetworkInterfacesOptions{}
			options.SetInstanceID(instanceID)
			networkInterfaces, response, err := vpcService.ListInstanceNetworkInterfaces(options)

			// end-list_instance_network_interfaces
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterfaces).ToNot(BeNil())

		})
		It(`CreateInstanceNetworkInterface request example`, func() {
			fmt.Println("\nCreateInstanceNetworkInterface() result:")
			// begin-create_instance_network_interface

			options := &vpcbetav1.CreateInstanceNetworkInterfaceOptions{}
			options.SetInstanceID(instanceID)
			options.SetName("eth1")
			options.SetSubnet(&vpcbetav1.SubnetIdentityByID{
				ID: &subnetID,
			})
			networkInterface, response, err := vpcService.CreateInstanceNetworkInterface(options)

			// end-create_instance_network_interface
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkInterface).ToNot(BeNil())
			eth2ID = *networkInterface.ID
		})
		It(`GetInstanceNetworkInterface request example`, func() {
			fmt.Println("\nGetInstanceNetworkInterface() result:")
			// begin-get_instance_network_interface

			options := &vpcbetav1.GetInstanceNetworkInterfaceOptions{}
			options.SetID(eth2ID)
			options.SetInstanceID(instanceID)
			networkInterface, response, err := vpcService.GetInstanceNetworkInterface(options)

			// end-get_instance_network_interface
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`UpdateInstanceNetworkInterface request example`, func() {
			fmt.Println("\nUpdateInstanceNetworkInterface() result:")
			name := getName("nic")
			ipSpoofing := true
			// begin-update_instance_network_interface

			options := &vpcbetav1.UpdateInstanceNetworkInterfaceOptions{
				InstanceID: &instanceID,
				ID:         &eth2ID,
			}
			instancePatchModel := &vpcbetav1.NetworkInterfacePatch{
				Name:            &name,
				AllowIPSpoofing: &ipSpoofing,
			}
			networkInterfacePatch, asPatchErr := instancePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.NetworkInterfacePatch = networkInterfacePatch
			networkInterface, response, err := vpcService.UpdateInstanceNetworkInterface(options)

			// end-update_instance_network_interface
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkInterface).ToNot(BeNil())

		})
		It(`AddInstanceNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nAddInstanceNetworkInterfaceFloatingIP() result:")
			// begin-add_instance_network_interface_floating_ip

			options := &vpcbetav1.AddInstanceNetworkInterfaceFloatingIPOptions{}
			options.SetID(floatingIPID)
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			floatingIP, response, err :=
				vpcService.AddInstanceNetworkInterfaceFloatingIP(options)

			// end-add_instance_network_interface_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`ListInstanceNetworkInterfaceFloatingIps request example`, func() {
			fmt.Println("\nListInstanceNetworkInterfaceFloatingIps() result:")
			// begin-list_instance_network_interface_floating_ips

			options := &vpcbetav1.ListInstanceNetworkInterfaceFloatingIpsOptions{}
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			floatingIPs, response, err :=
				vpcService.ListInstanceNetworkInterfaceFloatingIps(options)

			// end-list_instance_network_interface_floating_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPs).ToNot(BeNil())

		})

		It(`GetInstanceNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nGetInstanceNetworkInterfaceFloatingIP() result:")
			// begin-get_instance_network_interface_floating_ip

			options := &vpcbetav1.GetInstanceNetworkInterfaceFloatingIPOptions{}
			options.SetID(floatingIPID)
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			floatingIP, response, err :=
				vpcService.GetInstanceNetworkInterfaceFloatingIP(options)

			// end-get_instance_network_interface_floating_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})

		It(`ListInstanceVolumeAttachments request example`, func() {
			fmt.Println("\nListInstanceVolumeAttachments() result:")
			// begin-list_instance_volume_attachments

			options := &vpcbetav1.ListInstanceVolumeAttachmentsOptions{}
			options.SetInstanceID(instanceID)
			volumeAttachments, response, err := vpcService.ListInstanceVolumeAttachments(
				options)

			// end-list_instance_volume_attachments
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachments).ToNot(BeNil())

		})
		It(`CreateInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nCreateInstanceVolumeAttachment() result:")
			// begin-create_instance_volume_attachment

			volumeAttachmentPrototypeVolumeModel := &vpcbetav1.VolumeAttachmentPrototypeVolumeVolumeIdentityVolumeIdentityByID{
				ID: &volumeID,
			}

			options := vpcService.NewCreateInstanceVolumeAttachmentOptions(
				instanceID,
				volumeAttachmentPrototypeVolumeModel,
			)

			volumeAttachment, response, err := vpcService.CreateInstanceVolumeAttachment(options)

			// end-create_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(volumeAttachment).ToNot(BeNil())
			volumeAttachmentID = *volumeAttachment.ID
		})
		It(`GetInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nGetInstanceVolumeAttachment() result:")
			// begin-get_instance_volume_attachment

			options := &vpcbetav1.GetInstanceVolumeAttachmentOptions{}
			options.SetInstanceID(instanceID)
			options.SetID(volumeAttachmentID)
			volumeAttachment, response, err := vpcService.GetInstanceVolumeAttachment(options)

			// end-get_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`UpdateInstanceVolumeAttachment request example`, func() {
			fmt.Println("\nUpdateInstanceVolumeAttachment() result:")
			name := getName("vol-att")
			// begin-update_instance_volume_attachment

			options := &vpcbetav1.UpdateInstanceVolumeAttachmentOptions{}
			volumeAttachmentPatchModel := &vpcbetav1.VolumeAttachmentPatch{
				Name: &name,
			}
			volumeAttachmentPatch, asPatchErr := volumeAttachmentPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SetInstanceID(instanceID)
			options.SetID(volumeAttachmentID)
			options.SetVolumeAttachmentPatch(volumeAttachmentPatch)
			volumeAttachment, response, err := vpcService.UpdateInstanceVolumeAttachment(options)

			// end-update_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(volumeAttachment).ToNot(BeNil())

		})
		It(`ListInstanceGroups request example`, func() {
			fmt.Println("\nListInstanceGroups() result:")
			// begin-list_instance_groups

			options := &vpcbetav1.ListInstanceGroupsOptions{}
			instanceGroups, response, err := vpcService.ListInstanceGroups(options)

			// end-list_instance_groups
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroups).ToNot(BeNil())

		})
		It(`CreateInstanceGroup request example`, func() {
			fmt.Println("\nCreateInstanceGroup() result:")
			name := getName("ig")
			// begin-create_instance_group

			options := &vpcbetav1.CreateInstanceGroupOptions{
				InstanceTemplate: &vpcbetav1.InstanceTemplateIdentity{
					ID: &instanceTemplateID,
				},
			}
			options.SetName(name)
			var subnetArray = []vpcbetav1.SubnetIdentityIntf{
				&vpcbetav1.SubnetIdentity{
					ID: &subnetID,
				},
			}
			options.SetSubnets(subnetArray)
			instanceGroup, response, err := vpcService.CreateInstanceGroup(options)

			// end-create_instance_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroup).ToNot(BeNil())
			instanceGroupID = *instanceGroup.ID
		})
		It(`GetInstanceGroup request example`, func() {
			fmt.Println("\nGetInstanceGroup() result:")
			// begin-get_instance_group

			options := &vpcbetav1.GetInstanceGroupOptions{}
			options.SetID(instanceGroupID)
			instanceGroup, response, err := vpcService.GetInstanceGroup(options)

			// end-get_instance_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`UpdateInstanceGroup request example`, func() {
			fmt.Println("\nUpdateInstanceGroup() result:")
			name := getName("ig")
			// begin-update_instance_group

			options := &vpcbetav1.UpdateInstanceGroupOptions{}
			options.SetID(instanceGroupID)
			instanceGroupPatchModel := vpcbetav1.InstanceGroupPatch{}
			instanceGroupPatchModel.Name = &name
			instanceGroupPatchModel.InstanceTemplate = &vpcbetav1.InstanceTemplateIdentity{
				ID: &instanceTemplateID,
			}
			instanceGroupPatchModel.MembershipCount = &[]int64{5}[0]
			instanceGroupPatch, asPatchErr := instanceGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceGroupPatch = instanceGroupPatch
			instanceGroup, response, err := vpcService.UpdateInstanceGroup(options)

			// end-update_instance_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroup).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagers request example`, func() {
			fmt.Println("\nListInstanceGroupManagers() result:")
			// begin-list_instance_group_managers

			options := &vpcbetav1.ListInstanceGroupManagersOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			instanceGroupManagers, response, err :=
				vpcService.ListInstanceGroupManagers(options)

			// end-list_instance_group_managers
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagers).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManager request example`, func() {
			fmt.Println("\nCreateInstanceGroupManager() result:")
			// begin-create_instance_group_manager

			prototype := &vpcbetav1.InstanceGroupManagerPrototype{
				ManagerType:        &[]string{"autoscale"}[0],
				MaxMembershipCount: &[]int64{5}[0],
			}
			options := &vpcbetav1.CreateInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerPrototype(prototype)
			instanceGroupManagerIntf, response, err :=
				vpcService.CreateInstanceGroupManager(options)
			instanceGroupManager := instanceGroupManagerIntf.(*vpcbetav1.InstanceGroupManager)

			// end-create_instance_group_manager
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManager).ToNot(BeNil())
			instanceGroupManagerID = *instanceGroupManager.ID
		})
		It(`GetInstanceGroupManager request example`, func() {
			fmt.Println("\nGetInstanceGroupManager() result:")
			// begin-get_instance_group_manager

			options := &vpcbetav1.GetInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupManagerID)
			instanceGroupManager, response, err := vpcService.GetInstanceGroupManager(options)

			// end-get_instance_group_manager
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManager request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManager() result:")
			name := getName("manager")
			// begin-update_instance_group_manager
			instanceGroupManagerPatchModel := &vpcbetav1.InstanceGroupManagerPatch{}
			instanceGroupManagerPatchModel.Name = &name
			instanceGroupManagerPatch, asPatchErr := instanceGroupManagerPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcbetav1.UpdateInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupManagerID)
			options.InstanceGroupManagerPatch = instanceGroupManagerPatch
			instanceGroupManager, response, err :=
				vpcService.UpdateInstanceGroupManager(options)

			// end-update_instance_group_manager
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManager).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagerActions request example`, func() {
			fmt.Println("\nListInstanceGroupManagerActions() result:")
			// begin-list_instance_group_manager_actions

			options := &vpcbetav1.ListInstanceGroupManagerActionsOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			instanceGroupManagerActions, response, err :=
				vpcService.ListInstanceGroupManagerActions(options)

			// end-list_instance_group_manager_actions
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerActions).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nCreateInstanceGroupManagerAction() result:")
			name := getName("igAction")
			// begin-create_instance_group_manager_action

			cronSpec := &[]string{"*/5 1,2,3 * * *"}[0]
			instanceGroupManagerScheduledActionGroupPrototypeModel :=
				&vpcbetav1.InstanceGroupManagerScheduledActionGroupPrototype{
					MembershipCount: &[]int64{5}[0],
				}
			instanceGroupManagerActionPrototypeModel :=
				&vpcbetav1.InstanceGroupManagerActionPrototypeScheduledActionPrototype{
					Name:     &name,
					CronSpec: cronSpec,
					Group:    instanceGroupManagerScheduledActionGroupPrototypeModel,
				}
			createInstanceGroupManagerActionOptions :=
				&vpcbetav1.CreateInstanceGroupManagerActionOptions{
					InstanceGroupID:                     &instanceGroupID,
					InstanceGroupManagerID:              &instanceGroupManagerID,
					InstanceGroupManagerActionPrototype: instanceGroupManagerActionPrototypeModel,
				}
			instanceGroupManagerActionIntf, response, err :=
				vpcService.CreateInstanceGroupManagerAction(
					createInstanceGroupManagerActionOptions,
				)
			instanceGroupManagerAction := instanceGroupManagerActionIntf.(*vpcbetav1.InstanceGroupManagerAction)
			// end-create_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerAction).ToNot(BeNil())
			instanceGroupManagerActionID = *instanceGroupManagerAction.ID
		})
		It(`GetInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nGetInstanceGroupManagerAction() result:")
			// begin-get_instance_group_manager_action

			options := &vpcbetav1.GetInstanceGroupManagerActionOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerActionID)
			instanceGroupManagerAction, response, err :=
				vpcService.GetInstanceGroupManagerAction(options)

			// end-get_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManagerAction request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManagerAction() result:")
			name := getName("igManager")
			// begin-update_instance_group_manager_action
			cronSpec := &[]string{"*/5 1,2,3 * * *"}[0]
			instanceGroupManagerScheduledActionGroupPatchModel :=
				&vpcbetav1.InstanceGroupManagerActionGroupPatch{
					MembershipCount: &[]int64{5}[0],
				}
			instanceGroupManagerActionPatchModel :=
				&vpcbetav1.InstanceGroupManagerActionPatch{
					Name:     &name,
					CronSpec: cronSpec,
					Group:    instanceGroupManagerScheduledActionGroupPatchModel,
				}
			instanceGroupManagerActionPatchModelAsPatch, asPatchErr :=
				instanceGroupManagerActionPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options :=
				&vpcbetav1.UpdateInstanceGroupManagerActionOptions{
					InstanceGroupID:                 &instanceGroupID,
					InstanceGroupManagerID:          &instanceGroupManagerID,
					ID:                              &instanceGroupManagerActionID,
					InstanceGroupManagerActionPatch: instanceGroupManagerActionPatchModelAsPatch,
				}

			instanceGroupManagerAction, response, err := vpcService.UpdateInstanceGroupManagerAction(options)

			// end-update_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerAction).ToNot(BeNil())

		})
		It(`ListInstanceGroupManagerPolicies request example`, func() {
			fmt.Println("\nListInstanceGroupManagerPolicies() result:")
			// begin-list_instance_group_manager_policies

			options := &vpcbetav1.ListInstanceGroupManagerPoliciesOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			instanceGroupManagerPolicies, response, err :=
				vpcService.ListInstanceGroupManagerPolicies(options)

			// end-list_instance_group_manager_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicies).ToNot(BeNil())

		})
		It(`CreateInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nCreateInstanceGroupManagerPolicy() result:")
			// begin-create_instance_group_manager_policy

			prototype := &vpcbetav1.InstanceGroupManagerPolicyPrototype{
				PolicyType:  &[]string{"target"}[0],
				MetricType:  &[]string{"cpu"}[0],
				MetricValue: &[]int64{20}[0],
			}
			options := &vpcbetav1.CreateInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetInstanceGroupManagerPolicyPrototype(prototype)
			instanceGroupManagerPolicyIntf, response, err :=
				vpcService.CreateInstanceGroupManagerPolicy(options)
			instanceGroupManagerPolicy := instanceGroupManagerPolicyIntf.(*vpcbetav1.InstanceGroupManagerPolicy)
			// end-create_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())
			instanceGroupManagerPolicyID = *instanceGroupManagerPolicy.ID
		})
		It(`GetInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nGetInstanceGroupManagerPolicy() result:")
			// begin-get_instance_group_manager_policy

			options := &vpcbetav1.GetInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerPolicyID)
			instanceGroupManagerPolicy, response, err :=
				vpcService.GetInstanceGroupManagerPolicy(options)

			// end-get_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupManagerPolicy request example`, func() {
			fmt.Println("\nUpdateInstanceGroupManagerPolicy() result:")
			name := getName("igPolicy")
			// begin-update_instance_group_manager_policy

			options := &vpcbetav1.UpdateInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerPolicyID)
			instanceGroupManagerPolicyPatchModel := &vpcbetav1.InstanceGroupManagerPolicyPatch{}
			instanceGroupManagerPolicyPatchModel.Name = &name
			instanceGroupManagerPolicyPatchModel.MetricType = &[]string{"cpu"}[0]
			instanceGroupManagerPolicyPatchModel.MetricValue = &[]int64{70}[0]
			instanceGroupManagerPolicyPatchModelAsPatch, asPatchErr :=
				instanceGroupManagerPolicyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceGroupManagerPolicyPatch =
				instanceGroupManagerPolicyPatchModelAsPatch
			instanceGroupManagerPolicy, response, err :=
				vpcService.UpdateInstanceGroupManagerPolicy(options)

			// end-update_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupManagerPolicy).ToNot(BeNil())

		})
		It(`ListInstanceGroupMemberships request example`, func() {
			fmt.Println("\nListInstanceGroupMemberships() result:")
			// begin-list_instance_group_memberships

			options := &vpcbetav1.ListInstanceGroupMembershipsOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			instanceGroupMemberships, response, err :=
				vpcService.ListInstanceGroupMemberships(options)

			// end-list_instance_group_memberships
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMemberships).ToNot(BeNil())
			instanceGroupMembershipID = *instanceGroupMemberships.Memberships[0].ID
		})
		It(`GetInstanceGroupMembership request example`, func() {
			fmt.Println("\nGetInstanceGroupMembership() result:")
			// begin-get_instance_group_membership

			options := &vpcbetav1.GetInstanceGroupMembershipOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupMembershipID)
			instanceGroupMembership, response, err :=
				vpcService.GetInstanceGroupMembership(options)

			// end-get_instance_group_membership
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
		It(`UpdateInstanceGroupMembership request example`, func() {
			fmt.Println("\nUpdateInstanceGroupMembership() result:")
			name := getName("membership")
			// begin-update_instance_group_membership

			options := &vpcbetav1.UpdateInstanceGroupMembershipOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupMembershipID)
			instanceGroupMembershipPatchModel := &vpcbetav1.InstanceGroupMembershipPatch{}
			instanceGroupMembershipPatchModel.Name = &name
			instanceGroupMembershipPatch, asPatchErr := instanceGroupMembershipPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.InstanceGroupMembershipPatch = instanceGroupMembershipPatch
			instanceGroupMembership, response, err :=
				vpcService.UpdateInstanceGroupMembership(options)

			// end-update_instance_group_membership
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(instanceGroupMembership).ToNot(BeNil())

		})
		It(`ListDedicatedHostGroups request example`, func() {
			fmt.Println("\nListDedicatedHostGroups() result:")
			// begin-list_dedicated_host_groups

			options := vpcService.NewListDedicatedHostGroupsOptions()
			dedicatedHostGroups, response, err :=
				vpcService.ListDedicatedHostGroups(options)

			// end-list_dedicated_host_groups
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroups).ToNot(BeNil())

		})
		It(`CreateDedicatedHostGroup request example`, func() {
			fmt.Println("\nCreateDedicatedHostGroup() result:")
			name := getName("dhg")
			// begin-create_dedicated_host_group

			options := &vpcbetav1.CreateDedicatedHostGroupOptions{
				Name:   &name,
				Class:  &[]string{"mx2d"}[0],
				Family: &[]string{"memory-disk"}[0],
				Zone: &vpcbetav1.ZoneIdentity{
					Name: zone,
				},
			}
			dedicatedHostGroup, response, err := vpcService.CreateDedicatedHostGroup(options)

			// end-create_dedicated_host_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHostGroup).ToNot(BeNil())
			dedicatedHostGroupID = *dedicatedHostGroup.ID
		})
		It(`GetDedicatedHostGroup request example`, func() {
			fmt.Println("\nGetDedicatedHostGroup() result:")
			// begin-get_dedicated_host_group

			options := vpcService.NewGetDedicatedHostGroupOptions(dedicatedHostGroupID)
			dedicatedHostGroup, response, err := vpcService.GetDedicatedHostGroup(options)

			// end-get_dedicated_host_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`UpdateDedicatedHostGroup request example`, func() {
			fmt.Println("\nUpdateDedicatedHostGroup() result:")
			name := getName("dhg")
			// begin-update_dedicated_host_group

			dedicatedHostGroupPatchModel := &vpcbetav1.DedicatedHostGroupPatch{
				Name: &name,
			}
			dedicatedHostGroupPatchModelAsPatch, asPatchErr := dedicatedHostGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}

			updateDedicatedHostGroupOptions := vpcService.NewUpdateDedicatedHostGroupOptions(
				dedicatedHostGroupID,
				dedicatedHostGroupPatchModelAsPatch,
			)

			dedicatedHostGroup, response, err := vpcService.UpdateDedicatedHostGroup(updateDedicatedHostGroupOptions)

			// end-update_dedicated_host_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostGroup).ToNot(BeNil())

		})
		It(`ListDedicatedHostProfiles request example`, func() {
			fmt.Println("\nListDedicatedHostProfiles() result:")
			// begin-list_dedicated_host_profiles

			options := &vpcbetav1.ListDedicatedHostProfilesOptions{}
			profiles, response, err := vpcService.ListDedicatedHostProfiles(options)

			// end-list_dedicated_host_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())
			dhProfile = *profiles.Profiles[0].Name
		})
		It(`GetDedicatedHostProfile request example`, func() {
			fmt.Println("\nGetDedicatedHostProfile() result:")
			// begin-get_dedicated_host_profile

			options := &vpcbetav1.GetDedicatedHostProfileOptions{}
			options.SetName(dhProfile)
			profile, response, err := vpcService.GetDedicatedHostProfile(options)

			// end-get_dedicated_host_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ListDedicatedHosts request example`, func() {
			fmt.Println("\nListDedicatedHosts() result:")
			// begin-list_dedicated_hosts

			options := vpcService.NewListDedicatedHostsOptions()
			dedicatedHosts, response, err :=
				vpcService.ListDedicatedHosts(options)

			// end-list_dedicated_hosts
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHosts).ToNot(BeNil())

		})
		It(`CreateDedicatedHost request example`, func() {
			fmt.Println("\nCreateDedicatedHost() result:")
			name := getName("dh")
			profile := "mx2d-host-152x1216"
			// begin-create_dedicated_host

			options := &vpcbetav1.CreateDedicatedHostOptions{}
			options.SetDedicatedHostPrototype(&vpcbetav1.DedicatedHostPrototype{
				Name: &name,
				Profile: &vpcbetav1.DedicatedHostProfileIdentity{
					Name: &profile,
				},
				Group: &vpcbetav1.DedicatedHostGroupIdentity{
					ID: &dedicatedHostGroupID,
				},
			})
			dedicatedHost, response, err := vpcService.CreateDedicatedHost(options)

			// end-create_dedicated_host
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(dedicatedHost).ToNot(BeNil())
			dedicatedHostID = *dedicatedHost.ID
		})
		It(`GetDedicatedHost request example`, func() {
			fmt.Println("\nGetDedicatedHost() result:")
			// begin-get_dedicated_host

			options := vpcService.NewGetDedicatedHostOptions(dedicatedHostID)
			dedicatedHost, response, err := vpcService.GetDedicatedHost(options)

			// end-get_dedicated_host
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`ListDedicatedHostDisks request example`, func() {
			fmt.Println("\nListDedicatedHostDisks() result:")
			options := vpcService.NewListDedicatedHostsOptions()
			dedicatedHosts, response, err :=
				vpcService.ListDedicatedHosts(options)
			for i := range dedicatedHosts.DedicatedHosts {
				if len(dedicatedHosts.DedicatedHosts[i].Disks) > 0 {
					dhID = *dedicatedHosts.DedicatedHosts[i].ID
					break
				}
			}
			// begin-list_dedicated_host_disks

			listDedicatedHostDisksOptions := vpcService.NewListDedicatedHostDisksOptions(
				dhID,
			)
			dedicatedHostDiskCollection, response, err :=
				vpcService.ListDedicatedHostDisks(listDedicatedHostDisksOptions)

			// end-list_dedicated_host_disks
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDiskCollection).ToNot(BeNil())
			diskID = *dedicatedHostDiskCollection.Disks[0].ID
		})
		It(`GetDedicatedHostDisk request example`, func() {
			fmt.Println("\nGetDedicatedHostDisk() result:")
			// begin-get_dedicated_host_disk

			getDedicatedHostDiskOptions := vpcService.NewGetDedicatedHostDiskOptions(
				dhID,
				diskID,
			)
			dedicatedHostDisk, response, err :=
				vpcService.GetDedicatedHostDisk(getDedicatedHostDiskOptions)

			// end-get_dedicated_host_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
		It(`UpdateDedicatedHostDisk request example`, func() {
			fmt.Println("\nUpdateDedicatedHostDisk() result:")
			name := getName("dhdisk")
			// begin-update_dedicated_host_disk

			dedicatedHostDiskPatchModel := &vpcbetav1.DedicatedHostDiskPatch{
				Name: &name,
			}
			dedicatedHostDiskPatchModelAsPatch, asPatchErr := dedicatedHostDiskPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := vpcService.NewUpdateDedicatedHostDiskOptions(
				dhID,
				diskID,
				dedicatedHostDiskPatchModelAsPatch,
			)
			dedicatedHostDisk, response, err := vpcService.UpdateDedicatedHostDisk(options)

			// end-update_dedicated_host_disk
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHostDisk).ToNot(BeNil())

		})
		It(`UpdateDedicatedHost request example`, func() {
			fmt.Println("\nUpdateDedicatedHost() result:")
			name := getName("dh")
			insPlacement := false
			// begin-update_dedicated_host
			options := &vpcbetav1.UpdateDedicatedHostOptions{
				ID: &dedicatedHostID,
			}
			dedicatedHostPatchModel := &vpcbetav1.DedicatedHostPatch{
				Name:                     &name,
				InstancePlacementEnabled: &insPlacement,
			}
			dedicatedHostPatchModelAsPatch, asPatchErr := dedicatedHostPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.DedicatedHostPatch = dedicatedHostPatchModelAsPatch
			dedicatedHost, response, err := vpcService.UpdateDedicatedHost(options)
			// end-update_dedicated_host
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(dedicatedHost).ToNot(BeNil())

		})
		It(`ListVolumeProfiles request example`, func() {
			fmt.Println("\nListVolumeProfiles() result:")
			// begin-list_volume_profiles

			options := &vpcbetav1.ListVolumeProfilesOptions{}
			profiles, response, err := vpcService.ListVolumeProfiles(options)

			// end-list_volume_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())

		})
		It(`GetVolumeProfile request example`, func() {
			fmt.Println("\nGetVolumeProfile() result:")
			// begin-get_volume_profile

			options := &vpcbetav1.GetVolumeProfileOptions{}
			options.SetName("10iops-tier")
			profile, response, err := vpcService.GetVolumeProfile(options)

			// end-get_volume_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})

		It(`ListShareProfiles request example`, func() {
			fmt.Println("\nListShareProfiles() result:")
			// begin-list_share_profiles

			listShareProfilesOptions := vpcService.NewListShareProfilesOptions()
			listShareProfilesOptions.SetSort("name")

			shareProfileCollection, response, err := vpcService.ListShareProfiles(listShareProfilesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_share_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareProfileCollection).ToNot(BeNil())
			shareProfileName = *shareProfileCollection.Profiles[0].Name
		})
		It(`GetShareProfile request example`, func() {
			fmt.Println("\nGetShareProfile() result:")
			// begin-get_share_profile

			getShareProfileOptions := vpcService.NewGetShareProfileOptions(
				shareProfileName,
			)

			shareProfile, response, err := vpcService.GetShareProfile(getShareProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-get_share_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareProfile).ToNot(BeNil())

		})
		It(`ListShares request example`, func() {
			fmt.Println("\nListShares() result:")
			// begin-list_shares

			listSharesOptions := vpcService.NewListSharesOptions()
			listSharesOptions.SetSort("name")

			shareCollection, response, err := vpcService.ListShares(listSharesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_shares

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareCollection).ToNot(BeNil())

		})
		It(`CreateShare request example`, func() {
			fmt.Println("\nCreateShare() result:")
			// begin-create_share

			vpcIdentityModel := &vpcbetav1.VPCIdentityByID{
				ID: &vpcID,
			}
			shareProfileIdentityModel := &vpcbetav1.ShareProfileIdentityByName{
				Name: core.StringPtr("tier-3iops"),
			}
			zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
				Name: core.StringPtr("us-east-1"),
			}
			shareTargetPrototypeModel := &vpcbetav1.ShareMountTargetPrototype{
				Name: core.StringPtr("my-share-target"),
				VPC:  vpcIdentityModel,
			}
			sharePrototypeShareContextModel := &vpcbetav1.SharePrototypeShareContext{
				Iops:                core.Int64Ptr(int64(100)),
				Name:                core.StringPtr("my-replica-share"),
				Profile:             shareProfileIdentityModel,
				ReplicationCronSpec: core.StringPtr("0 */5 * * *"),
				MountTargets:        []vpcbetav1.ShareMountTargetPrototypeIntf{shareTargetPrototypeModel},
				UserTags:            []string{"my-share-tag"},
				Zone:                zoneIdentityModel,
			}
			sharePrototype := &vpcbetav1.SharePrototype{
				AccessControlMode: &[]string{"security_group"}[0],
				Name:              core.StringPtr("my-share"),
				Size:              core.Int64Ptr(int64(600)),
				Profile:           shareProfileIdentityModel,
				ReplicaShare:      sharePrototypeShareContextModel,
				Zone:              zoneIdentityModel,
			}
			createShareOptions := vpcService.NewCreateShareOptions(sharePrototype)

			share, response, err := vpcService.CreateShare(createShareOptions)
			if err != nil {
				panic(err)
			}

			// end-create_share

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(share).ToNot(BeNil())
			createdShareID = *share.ID
			createdReplicaShareID = *share.ReplicaShare.ID

			shareTargetPrototypeModel1 := &vpcbetav1.ShareMountTargetPrototype{
				Name: core.StringPtr("my-share-target-1"),
				VPC:  vpcIdentityModel,
			}
			sharePrototypeShareContextModel1 := &vpcbetav1.SharePrototypeShareContext{
				Iops:                core.Int64Ptr(int64(100)),
				Name:                core.StringPtr("my-replica-share-1"),
				Profile:             shareProfileIdentityModel,
				ReplicationCronSpec: core.StringPtr("0 */5 * * *"),
				MountTargets:        []vpcbetav1.ShareMountTargetPrototypeIntf{shareTargetPrototypeModel1},
				UserTags:            []string{"my-share-tag-1"},
				Zone:                zoneIdentityModel,
			}
			sharePrototype1 := &vpcbetav1.SharePrototype{
				Name:         core.StringPtr("my-share-1"),
				Size:         core.Int64Ptr(int64(600)),
				Profile:      shareProfileIdentityModel,
				ReplicaShare: sharePrototypeShareContextModel1,
				Zone:         zoneIdentityModel,
			}

			createShareOptions.SetSharePrototype(sharePrototype1)
			share, response, err = vpcService.CreateShare(createShareOptions)
			createdShare1ID = *share.ID
			createdReplicaShare1ID = *share.ReplicaShare.ID
		})
		It(`GetShare request example`, func() {
			fmt.Println("\nGetShare() result:")
			// begin-get_share

			getShareOptions := vpcService.NewGetShareOptions(
				createdShareID,
			)

			share, response, err := vpcService.GetShare(getShareOptions)
			if err != nil {
				panic(err)
			}
			createdShareETag = response.GetHeaders()["Etag"][0]
			// end-get_share

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(share).ToNot(BeNil())

		})
		It(`UpdateShare request example`, func() {
			fmt.Println("\nUpdateShare() result:")
			// begin-update_share

			sharePatchModel := &vpcbetav1.SharePatch{
				Name: core.StringPtr("my-share-updated"),
			}
			sharePatchModelAsPatch, asPatchErr := sharePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateShareOptions := vpcService.NewUpdateShareOptions(
				createdShareID,
				sharePatchModelAsPatch,
			)
			updateShareOptions.SetIfMatch(createdShareETag)

			share, response, err := vpcService.UpdateShare(updateShareOptions)
			if err != nil {
				panic(err)
			}

			// end-update_share

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(share).ToNot(BeNil())

		})
		It(`FailoverShare request example`, func() {
			// begin-failover_share

			failoverShareOptions := vpcService.NewFailoverShareOptions(
				createdShareID,
			)

			response, err := vpcService.FailoverShare(failoverShareOptions)
			if err != nil {
				panic(err)
			}

			// end-failover_share
			fmt.Printf("\nFailoverShare() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`GetShareSource request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nGetShareSource() result:")
			// begin-get_share_source

			getShareSourceOptions := vpcService.NewGetShareSourceOptions(
				createdReplicaShareID,
			)

			share, response, err := vpcService.GetShareSource(getShareSourceOptions)
			if err != nil {
				panic(err)
			}

			// end-get_share_source

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(share).ToNot(BeNil())

		})
		It(`ListShareMountTargets request example`, func() {
			fmt.Println("\nListShareMountTargets() result:")

			// begin-list_share_mount_targets

			listShareMountTargetsOptions := &vpcbetav1.ListShareMountTargetsOptions{
				ShareID: &createdShareID,
			}

			shareMountTargetCollection, response, err := vpcService.ListShareMountTargets(listShareMountTargetsOptions)
			if err != nil {
				panic(err)
			}

			// end-list_share_mount_targets

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareMountTargetCollection).ToNot(BeNil())
		})
		It(`CreateShareMountTarget request example`, func() {
			fmt.Println("\nCreateShareMountTarget() result:")

			// begin-create_share_mount_target
			virtualNetworkInterfacePrimaryIPPrototype := &vpcbetav1.VirtualNetworkInterfacePrimaryIPPrototype{
				Address:    &[]string{"10.0.1.3"}[0],
				AutoDelete: &[]bool{true}[0],
				Name:       &[]string{"my-reserved-ip"}[0],
			}

			subnetIdentityModel := &vpcbetav1.SubnetIdentityByID{
				ID: &subnetID,
			}

			shareMountTargetVirtualNetworkInterfacePrototype := &vpcbetav1.ShareMountTargetVirtualNetworkInterfacePrototype{
				Name:      &[]string{"my-virtual-network-interface"}[0],
				PrimaryIP: virtualNetworkInterfacePrimaryIPPrototype,
				Subnet:    subnetIdentityModel,
			}

			shareMountTargetPrototypeModel := &vpcbetav1.ShareMountTargetPrototypeShareMountTargetByAccessControlModeSecurityGroup{
				Name:                    &[]string{"my-share-mount-target"}[0],
				TransitEncryption:       &[]string{"user_managed"}[0],
				VirtualNetworkInterface: shareMountTargetVirtualNetworkInterfacePrototype,
			}

			createShareMountTargetOptions := &vpcbetav1.CreateShareMountTargetOptions{
				ShareID:                   &createdReplicaShare1ID,
				ShareMountTargetPrototype: shareMountTargetPrototypeModel,
			}

			shareMountTarget, response, err := vpcService.CreateShareMountTarget(createShareMountTargetOptions)
			if err != nil {
				panic(err)
			}

			// end-create_share_mount_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(shareMountTarget).ToNot(BeNil())
			createdShareTargetID = *shareMountTarget.ID

		})
		It(`GetShareMountTarget request example`, func() {
			fmt.Println("\nGetShareMountTarget() result:")

			shareMountTargetId := createdShareTargetID
			// begin-get_share_mount_target

			getShareMountTargetOptions := &vpcbetav1.GetShareMountTargetOptions{
				ShareID: &createdReplicaShare1ID,
				ID:      &shareMountTargetId,
			}

			shareMountTarget, response, err := vpcService.GetShareMountTarget(getShareMountTargetOptions)
			if err != nil {
				panic(err)
			}
			// end-get_share_mount_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareMountTarget).ToNot(BeNil())
		})
		It(`UpdateShareMountTarget request example`, func() {
			fmt.Println("\nUpdateShareMountTarget() result:")

			shareMountTargetId := createdShareTargetID
			// begin-update_share_mount_target

			shareMountTargetPatchModel := &vpcbetav1.ShareMountTargetPatch{
				Name: &[]string{"my-share-mount-target-updated"}[0],
			}
			shareMountTargetPatchModelAsPatch, asPatchErr := shareMountTargetPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateShareMountTargetOptions := &vpcbetav1.UpdateShareMountTargetOptions{
				ShareID:               &createdReplicaShare1ID,
				ID:                    &shareMountTargetId,
				ShareMountTargetPatch: shareMountTargetPatchModelAsPatch,
			}

			shareMountTarget, response, err := vpcService.UpdateShareMountTarget(updateShareMountTargetOptions)
			if err != nil {
				panic(err)
			}

			// end-update_share_mount_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(shareMountTarget).ToNot(BeNil())
		})

		It(`ListVirtualNetworkInterfaces request example`, func() {
			fmt.Println("\nListVirtualNetworkInterfaces() result:")
			// begin-list_virtual_network_interfaces
			listVirtualNetworkInterfacesOptions := &vpcbetav1.ListVirtualNetworkInterfacesOptions{}

			virtualNetworkInterfaceCollection, response, err := vpcService.ListVirtualNetworkInterfaces(listVirtualNetworkInterfacesOptions)

			// end-list_virtual_network_interfaces
			virtualNetworkInterfaceId = *virtualNetworkInterfaceCollection.VirtualNetworkInterfaces[0].ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(virtualNetworkInterfaceCollection).ToNot(BeNil())
		})
		It(`GetVirtualNetworkInterface request example`, func() {
			fmt.Println("\nGetVirtualNetworkInterface() result:")
			// begin-get_virtual_network_interface

			getVirtualNetworkInterfaceOptions := &vpcbetav1.GetVirtualNetworkInterfaceOptions{
				ID: &virtualNetworkInterfaceId,
			}

			virtualNetworkInterface, response, err := vpcService.GetVirtualNetworkInterface(getVirtualNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-get_virtual_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(virtualNetworkInterface).ToNot(BeNil())
		})
		It(`UpdateVirtualNetworkInterface request example`, func() {
			fmt.Println("\nUpdateVirtualNetworkInterface() result:")
			// begin-update_virtual_network_interface

			virtualNetworkInterfacePatchModel := &vpcbetav1.VirtualNetworkInterfacePatch{
				Name: &[]string{"my-virtual-network-interface-updated"}[0],
			}
			virtualNetworkInterfacePatchModelAsPatch, asPatchErr := virtualNetworkInterfacePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVirtualNetworkInterfaceOptions := &vpcbetav1.UpdateVirtualNetworkInterfaceOptions{
				ID:                           &virtualNetworkInterfaceId,
				VirtualNetworkInterfacePatch: virtualNetworkInterfacePatchModelAsPatch,
			}

			virtualNetworkInterface, response, err := vpcService.UpdateVirtualNetworkInterface(updateVirtualNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-update_virtual_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(virtualNetworkInterface).ToNot(BeNil())
		})

		It(`DeleteShareMountTarget request example`, func() {
			fmt.Println("\nDeleteShareMountTarget() result:")

			shareMountTargetId := createdShareTargetID
			// begin-delete_share_mount_target

			deleteShareMountTargetOptions := &vpcbetav1.DeleteShareMountTargetOptions{
				ShareID: &createdReplicaShare1ID,
				ID:      &shareMountTargetId,
			}

			shareMountTarget, response, err := vpcService.DeleteShareMountTarget(deleteShareMountTargetOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_share_mount_target

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(shareMountTarget).ToNot(BeNil())
		})

		It(`DeleteShareSource request example`, func() {
			// begin-delete_share_source

			deleteShareSourceOptions := vpcService.NewDeleteShareSourceOptions(
				createdShare1ID,
			)

			response, err := vpcService.DeleteShareSource(deleteShareSourceOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_share_source
			fmt.Printf("\nDeleteShareSource() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteShare request example`, func() {
			fmt.Println("\nDeleteShare() result:")
			// begin-delete_share

			deleteShareOptions := vpcService.NewDeleteShareOptions(
				createdShareID,
			)
			deleteShareOptions.SetIfMatch(createdShareETag)

			share, response, err := vpcService.DeleteShare(deleteShareOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_share

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(share).ToNot(BeNil())

		})

		It(`ListSnapshots request example`, func() {
			fmt.Println("\nListSnapshots() result:")
			// begin-list_snapshots

			options := &vpcbetav1.ListSnapshotsOptions{}
			snapshotCollection, response, err := vpcService.ListSnapshots(options)

			// end-list_snapshots
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshotCollection).ToNot(BeNil())

		})
		It(`CreateSnapshot request example`, func() {
			fmt.Println("\nCreateSnapshot() result:")
			name := getName("snapshotone")
			volumeIdentityModel := &vpcbetav1.VolumeIdentityByID{
				ID: &volumeID,
			}
			secondSnap := &vpcbetav1.SnapshotPrototypeSnapshotBySourceVolume{
				Name:         &name,
				SourceVolume: volumeIdentityModel,
			}
			secondCreateSnapshotOptions := vpcService.NewCreateSnapshotOptions(
				secondSnap,
			)
			_, _, err := vpcService.CreateSnapshot(secondCreateSnapshotOptions)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			name = getName("snapshottwo")
			// begin-create_snapshot
			options := &vpcbetav1.SnapshotPrototypeSnapshotBySourceVolume{
				Name:         &name,
				SourceVolume: volumeIdentityModel,
			}
			createSnapshotOptions := vpcService.NewCreateSnapshotOptions(
				options,
			)
			snapshot, response, err := vpcService.CreateSnapshot(createSnapshotOptions)

			// end-create_snapshot
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(snapshot).ToNot(BeNil())
			snapshotID = *snapshot.ID
		})
		It(`GetSnapshot request example`, func() {
			fmt.Println("\nGetSnapshot() result:")
			// begin-get_snapshot

			options := &vpcbetav1.GetSnapshotOptions{
				ID: &snapshotID,
			}
			snapshot, response, err := vpcService.GetSnapshot(options)

			// end-get_snapshot
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())
			ifMatchSnapshot = response.GetHeaders()["Etag"][0]
		})
		It(`UpdateSnapshot request example`, func() {
			fmt.Println("\nUpdateSnapshot() result:")
			name := getName("snapshot")
			userTags := []string{"usertag-snap-1"}
			// begin-update_snapshot

			snapshotPatchModel := &vpcbetav1.SnapshotPatch{
				Name:     &name,
				UserTags: userTags,
			}
			snapshotPatchModelAsPatch, asPatchErr := snapshotPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			updateSnapshotOptions := &vpcbetav1.UpdateSnapshotOptions{
				ID:            &snapshotID,
				SnapshotPatch: snapshotPatchModelAsPatch,
				IfMatch:       &ifMatchSnapshot,
			}
			snapshot, response, err := vpcService.UpdateSnapshot(updateSnapshotOptions)

			// end-update_snapshot
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())

			Expect(response.StatusCode).To(Equal(200))
			Expect(snapshot).ToNot(BeNil())

		})
		It(`ListRegions request example`, func() {
			fmt.Println("\nListRegions() result:")
			// begin-list_regions

			listRegionsOptions := &vpcbetav1.ListRegionsOptions{}
			regions, response, err := vpcService.ListRegions(listRegionsOptions)

			// end-list_regions
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(regions).ToNot(BeNil())

		})
		It(`GetRegion request example`, func() {
			fmt.Println("\nGetRegion() result:")
			// begin-get_region

			getRegionOptions := &vpcbetav1.GetRegionOptions{}
			getRegionOptions.SetName("us-east")
			region, response, err := vpcService.GetRegion(getRegionOptions)

			// end-get_region
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(region).ToNot(BeNil())

		})
		It(`ListRegionZones request example`, func() {
			fmt.Println("\nListRegionZones() result:")
			// begin-list_region_zones

			listZonesOptions := &vpcbetav1.ListRegionZonesOptions{}
			listZonesOptions.SetRegionName("us-east")
			zones, response, err := vpcService.ListRegionZones(listZonesOptions)
			// end-list_region_zones
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zones).ToNot(BeNil())

		})
		It(`GetRegionZone request example`, func() {
			fmt.Println("\nGetRegionZone() result:")
			// begin-get_region_zone

			getZoneOptions := &vpcbetav1.GetRegionZoneOptions{}
			getZoneOptions.SetRegionName("us-east")
			getZoneOptions.SetName("us-east-1")
			zone, response, err := vpcService.GetRegionZone(getZoneOptions)

			// end-get_region_zone
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(zone).ToNot(BeNil())

		})
		It(`ListPublicGateways request example`, func() {
			fmt.Println("\nListPublicGateways() result:")
			// begin-list_public_gateways

			options := &vpcbetav1.ListPublicGatewaysOptions{}
			publicGateways, response, err := vpcService.ListPublicGateways(options)

			// end-list_public_gateways
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateways).ToNot(BeNil())

		})
		It(`CreatePublicGateway request example`, func() {
			fmt.Println("\nCreatePublicGateway() result:")
			// begin-create_public_gateway

			options := &vpcbetav1.CreatePublicGatewayOptions{}
			options.SetVPC(&vpcbetav1.VPCIdentity{
				ID: &vpcID,
			})
			options.SetZone(&vpcbetav1.ZoneIdentity{
				Name: zone,
			})
			publicGateway, response, err := vpcService.CreatePublicGateway(options)

			// end-create_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(publicGateway).ToNot(BeNil())
			publicGatewayID = *publicGateway.ID
		})
		It(`GetPublicGateway request example`, func() {
			fmt.Println("\nGetPublicGateway() result:")
			// begin-get_public_gateway

			options := &vpcbetav1.GetPublicGatewayOptions{}
			options.SetID(publicGatewayID)
			publicGateway, response, err := vpcService.GetPublicGateway(options)

			// end-get_public_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`UpdatePublicGateway request example`, func() {
			fmt.Println("\nUpdatePublicGateway() result:")
			name := getName("pgw")
			// begin-update_public_gateway

			options := &vpcbetav1.UpdatePublicGatewayOptions{}
			options.SetID(publicGatewayID)
			PublicGatewayPatchModel := &vpcbetav1.PublicGatewayPatch{
				Name: &name,
			}
			PublicGatewayPatch, asPatchErr := PublicGatewayPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.PublicGatewayPatch = PublicGatewayPatch
			publicGateway, response, err := vpcService.UpdatePublicGateway(options)
			// end-update_public_gateway
			if err != nil {
				panic(err)
			} // 	Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(publicGateway).ToNot(BeNil())

		})
		It(`ListNetworkAcls request example`, func() {
			fmt.Println("\nListNetworkAcls() result:")
			// begin-list_network_acls

			options := &vpcbetav1.ListNetworkAclsOptions{}
			networkACLCollection, response, err := vpcService.ListNetworkAcls(options)

			// end-list_network_acls
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLCollection).ToNot(BeNil())

		})
		It(`CreateNetworkACL request example`, func() {
			fmt.Println("\nCreateNetworkACL() result:")
			name := getName("acl")
			// begin-create_network_acl
			options := &vpcbetav1.CreateNetworkACLOptions{}
			options.SetNetworkACLPrototype(&vpcbetav1.NetworkACLPrototype{
				Name: &name,
				VPC: &vpcbetav1.VPCIdentity{
					ID: &vpcID,
				},
			})
			networkACL, response, err := vpcService.CreateNetworkACL(options)
			// end-create_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACL).ToNot(BeNil())
			networkACLID = *networkACL.ID
		})
		It(`GetNetworkACL request example`, func() {
			fmt.Println("\nGetNetworkACL() result:")
			// begin-get_network_acl

			options := &vpcbetav1.GetNetworkACLOptions{}
			options.SetID(networkACLID)
			networkACL, response, err := vpcService.GetNetworkACL(options)

			// end-get_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`UpdateNetworkACL request example`, func() {
			fmt.Println("\nUpdateNetworkACL() result:")
			name := getName("acl")
			// begin-update_network_acl

			options := &vpcbetav1.UpdateNetworkACLOptions{}
			options.SetID(networkACLID)
			networkACLPatchModel := &vpcbetav1.NetworkACLPatch{
				Name: &name,
			}
			networkACLPatch, asPatchErr := networkACLPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.NetworkACLPatch = networkACLPatch
			networkACL, response, err := vpcService.UpdateNetworkACL(options)

			// end-update_network_acl
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACL).ToNot(BeNil())

		})
		It(`ListNetworkACLRules request example`, func() {
			fmt.Println("\nListNetworkACLRules() result:")
			// begin-list_network_acl_rules

			options := &vpcbetav1.ListNetworkACLRulesOptions{}
			options.SetNetworkACLID(networkACLID)
			networkACLRules, response, err := vpcService.ListNetworkACLRules(options)

			// end-list_network_acl_rules
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRules).ToNot(BeNil())

		})
		It(`CreateNetworkACLRule request example`, func() {
			fmt.Println("\nCreateNetworkACLRule() result:")
			name := getName("aclrule")
			// begin-create_network_acl_rule
			options := &vpcbetav1.CreateNetworkACLRuleOptions{}
			options.SetNetworkACLID(networkACLID)
			options.SetNetworkACLRulePrototype(&vpcbetav1.NetworkACLRulePrototype{
				Action:      &[]string{"allow"}[0],
				Destination: &[]string{"192.168.3.2/32"}[0],
				Direction:   &[]string{"inbound"}[0],
				Source:      &[]string{"192.168.3.2/32"}[0],
				Protocol:    &[]string{"all"}[0],
				Name:        &name,
			})
			networkACLRuleIntf, response, err := vpcService.CreateNetworkACLRule(options)
			networkACLRule := networkACLRuleIntf.(*vpcbetav1.NetworkACLRuleNetworkACLRuleProtocolAll)
			// end-create_network_acl_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(networkACLRule).ToNot(BeNil())
			networkACLRuleID = *networkACLRule.ID
		})
		It(`GetNetworkACLRule request example`, func() {
			fmt.Println("\nGetNetworkACLRule() result:")
			// begin-get_network_acl_rule

			options := &vpcbetav1.GetNetworkACLRuleOptions{}
			options.SetID(networkACLRuleID)
			options.SetNetworkACLID(networkACLID)
			networkACLRule, response, err := vpcService.GetNetworkACLRule(options)

			// end-get_network_acl_rule

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`UpdateNetworkACLRule request example`, func() {
			fmt.Println("\nUpdateNetworkACLRule() result:")
			name := getName("aclrule")
			// begin-update_network_acl_rule
			options := &vpcbetav1.UpdateNetworkACLRuleOptions{}
			options.SetID(networkACLRuleID)
			options.SetNetworkACLID(networkACLID)
			networkACLRulePatchModel := &vpcbetav1.NetworkACLRulePatch{
				Name: &name,
			}
			networkACLRulePatch, asPatchErr := networkACLRulePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.NetworkACLRulePatch = networkACLRulePatch
			networkACLRule, response, err := vpcService.UpdateNetworkACLRule(options)
			// end-update_network_acl_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(networkACLRule).ToNot(BeNil())

		})
		It(`ListSecurityGroups request example`, func() {
			fmt.Println("\nListSecurityGroups() result:")
			// begin-list_security_groups

			options := &vpcbetav1.ListSecurityGroupsOptions{}
			securityGroups, response, err := vpcService.ListSecurityGroups(options)

			// end-list_security_groups
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroups).ToNot(BeNil())

		})
		It(`CreateSecurityGroup request example`, func() {
			fmt.Println("\nCreateSecurityGroup() result:")
			name := getName("sg")
			// begin-create_security_group

			options := &vpcbetav1.CreateSecurityGroupOptions{}
			options.SetVPC(&vpcbetav1.VPCIdentity{
				ID: &vpcID,
			})
			options.SetName(name)
			securityGroup, response, err := vpcService.CreateSecurityGroup(options)

			// end-create_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroup).ToNot(BeNil())
			securityGroupID = *securityGroup.ID
		})
		It(`GetSecurityGroup request example`, func() {
			fmt.Println("\nGetSecurityGroup() result:")
			// begin-get_security_group

			options := &vpcbetav1.GetSecurityGroupOptions{}
			options.SetID(securityGroupID)
			securityGroup, response, err := vpcService.GetSecurityGroup(options)

			// end-get_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`UpdateSecurityGroup request example`, func() {
			fmt.Println("\nUpdateSecurityGroup() result:")
			name := getName("sg")
			// begin-update_security_group
			options := &vpcbetav1.UpdateSecurityGroupOptions{}
			options.SetID(securityGroupID)
			securityGroupPatchModel := &vpcbetav1.SecurityGroupPatch{
				Name: &name,
			}
			securityGroupPatch, asPatchErr := securityGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SecurityGroupPatch = securityGroupPatch
			securityGroup, response, err := vpcService.UpdateSecurityGroup(options)

			// end-update_security_group
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroup).ToNot(BeNil())

		})
		It(`ListSecurityGroupRules request example`, func() {
			fmt.Println("\nListSecurityGroupRules() result:")
			// begin-list_security_group_rules

			options := &vpcbetav1.ListSecurityGroupRulesOptions{}
			options.SetSecurityGroupID(securityGroupID)
			rules, response, err := vpcService.ListSecurityGroupRules(options)

			// end-list_security_group_rules
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())

		})
		It(`CreateSecurityGroupRule request example`, func() {
			fmt.Println("\nCreateSecurityGroupRule() result:")
			// begin-create_security_group_rule

			options := &vpcbetav1.CreateSecurityGroupRuleOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetSecurityGroupRulePrototype(&vpcbetav1.SecurityGroupRulePrototype{
				Direction: &[]string{"inbound"}[0],
				Protocol:  &[]string{"udp"}[0],
			})
			securityGroupRuleIntf, response, err := vpcService.CreateSecurityGroupRule(options)
			securityGroupRule := securityGroupRuleIntf.(*vpcbetav1.SecurityGroupRuleSecurityGroupRuleProtocolTcpudp)
			// end-create_security_group_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupRule).ToNot(BeNil())
			securityGroupRuleID = *securityGroupRule.ID
		})
		It(`GetSecurityGroupRule request example`, func() {
			fmt.Println("\nGetSecurityGroupRule() result:")
			// begin-get_security_group_rule

			options := &vpcbetav1.GetSecurityGroupRuleOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(securityGroupRuleID)
			securityGroupRule, response, err := vpcService.GetSecurityGroupRule(options)

			// end-get_security_group_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`UpdateSecurityGroupRule request example`, func() {
			fmt.Println("\nUpdateSecurityGroupRule() result:")
			// begin-update_security_group_rule

			options := &vpcbetav1.UpdateSecurityGroupRuleOptions{}
			options.SecurityGroupID = &securityGroupID
			options.ID = &securityGroupRuleID
			securityGroupRulePatchModel := &vpcbetav1.SecurityGroupRulePatch{}
			securityGroupRulePatchModel.Direction = &[]string{"inbound"}[0]

			securityGroupRulePatch, asPatchErr := securityGroupRulePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.SecurityGroupRulePatch = securityGroupRulePatch
			securityGroupRule, response, err := vpcService.UpdateSecurityGroupRule(options)

			// end-update_security_group_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(securityGroupRule).ToNot(BeNil())

		})
		It(`ListSecurityGroupTargets request example`, func() {
			fmt.Println("\nListSecurityGroupTargets() result:")
			// begin-list_security_group_targets

			options := &vpcbetav1.ListSecurityGroupTargetsOptions{}
			options.SetSecurityGroupID(securityGroupID)
			targets, response, err := vpcService.ListSecurityGroupTargets(options)

			// end-list_security_group_targets
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(targets).ToNot(BeNil())
		})
		It(`CreateSecurityGroupTargetBinding request example`, func() {
			fmt.Println("\nCreateSecurityGroupTargetBinding() result:")
			// begin-create_security_group_target_binding

			options := vpcService.NewCreateSecurityGroupTargetBindingOptions(
				securityGroupID,
				eth2ID,
			)

			securityGroupTargetReferenceIntf, response, err := vpcService.CreateSecurityGroupTargetBinding(options)
			securityGroupTargetReference := securityGroupTargetReferenceIntf.(*vpcbetav1.SecurityGroupTargetReference)

			// end-create_security_group_target_binding
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(securityGroupTargetReference).ToNot(BeNil())
			targetID = *securityGroupTargetReference.ID
		})
		It(`GetSecurityGroupTarget request example`, func() {
			fmt.Println("\nGetSecurityGroupTarget() result:")
			// begin-get_security_group_target

			options := &vpcbetav1.GetSecurityGroupTargetOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(targetID)
			target, response, err :=
				vpcService.GetSecurityGroupTarget(options)

			// end-get_security_group_target
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(target).ToNot(BeNil())

		})

		It(`ListIkePolicies request example`, func() {
			fmt.Println("\nListIkePolicies() result:")
			// begin-list_ike_policies

			options := vpcService.NewListIkePoliciesOptions()
			ikePolicies, response, err := vpcService.ListIkePolicies(options)

			// end-list_ike_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicies).ToNot(BeNil())

		})
		It(`CreateIkePolicy request example`, func() {
			fmt.Println("\nCreateIkePolicy() result:")
			name := getName("ike")
			// begin-create_ike_policy

			options := &vpcbetav1.CreateIkePolicyOptions{}
			options.SetName(name)
			options.SetAuthenticationAlgorithm("sha256")
			options.SetDhGroup(14)
			options.SetEncryptionAlgorithm("aes128")
			options.SetIkeVersion(1)
			ikePolicy, response, err := vpcService.CreateIkePolicy(options)

			// end-create_ike_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ikePolicy).ToNot(BeNil())
			ikePolicyID = *ikePolicy.ID
		})
		It(`GetIkePolicy request example`, func() {
			fmt.Println("\nGetIkePolicy() result:")
			// begin-get_ike_policy

			options := vpcService.NewGetIkePolicyOptions(ikePolicyID)
			ikePolicy, response, err := vpcService.GetIkePolicy(options)

			// end-get_ike_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`UpdateIkePolicy request example`, func() {
			fmt.Println("\nUpdateIkePolicy() result:")
			name := getName("ike")
			// begin-update_ike_policy

			options := &vpcbetav1.UpdateIkePolicyOptions{
				ID: &ikePolicyID,
			}
			ikePolicyPatchModel := &vpcbetav1.IkePolicyPatch{}
			ikePolicyPatchModel.Name = &name
			ikePolicyPatchModel.DhGroup = &[]int64{15}[0]
			ikePolicyPatch, asPatchErr := ikePolicyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.IkePolicyPatch = ikePolicyPatch
			ikePolicy, response, err := vpcService.UpdateIkePolicy(options)

			// end-update_ike_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ikePolicy).ToNot(BeNil())

		})
		It(`ListIkePolicyConnections request example`, func() {
			fmt.Println("\nListIkePolicyConnections() result:")
			// begin-list_ike_policy_connections

			options := &vpcbetav1.ListIkePolicyConnectionsOptions{
				ID: &ikePolicyID,
			}
			connections, response, err := vpcService.ListIkePolicyConnections(options)

			// end-list_ike_policy_connections
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connections).ToNot(BeNil())

		})
		It(`ListIpsecPolicies request example`, func() {
			fmt.Println("\nListIpsecPolicies() result:")
			// begin-list_ipsec_policies

			options := &vpcbetav1.ListIpsecPoliciesOptions{}
			ipsecPolicies, response, err := vpcService.ListIpsecPolicies(options)

			// end-list_ipsec_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipsecPolicies).ToNot(BeNil())

		})
		It(`CreateIpsecPolicy request example`, func() {
			fmt.Println("\nCreateIpsecPolicy() result:")
			name := getName("ipsec")
			// begin-create_ipsec_policy

			options := &vpcbetav1.CreateIpsecPolicyOptions{}
			options.SetName(name)
			options.SetAuthenticationAlgorithm("sha256")
			options.SetEncryptionAlgorithm("aes128")
			options.SetPfs("disabled")
			ipsecPolicy, response, err := vpcService.CreateIpsecPolicy(options)
			// end-create_ipsec_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(ipsecPolicy).ToNot(BeNil())
			ipsecPolicyID = *ipsecPolicy.ID
		})
		It(`GetIpsecPolicy request example`, func() {
			fmt.Println("\nGetIpsecPolicy() result:")
			// begin-get_ipsec_policy

			options := vpcService.NewGetIpsecPolicyOptions(ipsecPolicyID)
			ipsecPolicy, response, err := vpcService.GetIpsecPolicy(options)

			// end-get_ipsec_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipsecPolicy).ToNot(BeNil())

		})
		It(`UpdateIpsecPolicy request example`, func() {
			fmt.Println("\nUpdateIpsecPolicy() result:")
			name := getName("ipsec")
			// begin-update_ipsec_policy

			options := &vpcbetav1.UpdateIpsecPolicyOptions{
				ID: &ipsecPolicyID,
			}
			ipsecPolicyPatchModel := &vpcbetav1.IPsecPolicyPatch{
				Name:                    &name,
				AuthenticationAlgorithm: &[]string{"sha256"}[0],
			}
			ipsecPolicyPatch, asPatchErr := ipsecPolicyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.IPsecPolicyPatch = ipsecPolicyPatch
			ipsecPolicy, response, err := vpcService.UpdateIpsecPolicy(options)
			// end-update_ipsec_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(ipsecPolicy).ToNot(BeNil())
		})
		It(`ListIpsecPolicyConnections request example`, func() {
			fmt.Println("\nListIpsecPolicyConnections() result:")
			// begin-list_ipsec_policy_connections

			options := &vpcbetav1.ListIpsecPolicyConnectionsOptions{
				ID: &ipsecPolicyID,
			}
			connections, response, err :=
				vpcService.ListIpsecPolicyConnections(options)

			// end-list_ipsec_policy_connections
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(connections).ToNot(BeNil())

		})
		It(`ListVPNGateways request example`, func() {
			fmt.Println("\nListVPNGateways() result:")
			// begin-list_vpn_gateways

			listVPNGatewaysOptions := vpcService.NewListVPNGatewaysOptions()
			vpnGatewayCollection, response, err := vpcService.ListVPNGateways(listVPNGatewaysOptions)

			// end-list_vpn_gateways
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayCollection).ToNot(BeNil())

		})
		It(`CreateVPNGateway request example`, func() {
			fmt.Println("\nCreateVPNGateway() result:")
			name := getName("vpngateway")
			// begin-create_vpn_gateway

			vpnGatewayPrototypeModel := new(vpcbetav1.VPNGatewayPrototypeVPNGatewayRouteModePrototype)
			vpnGatewayPrototypeModel.Name = &name
			vpnGatewayPrototypeModel.Subnet = &vpcbetav1.SubnetIdentityByID{
				ID: &subnetID,
			}
			vpnGatewayPrototypeModel.Mode = &[]string{"route"}[0]

			createVPNGatewayOptionsModel := new(vpcbetav1.CreateVPNGatewayOptions)
			createVPNGatewayOptionsModel.VPNGatewayPrototype = vpnGatewayPrototypeModel
			vpnGatewayIntf, response, err := vpcService.CreateVPNGateway(createVPNGatewayOptionsModel)
			vpnGateway := vpnGatewayIntf.(*vpcbetav1.VPNGateway)
			// end-create_vpn_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGateway).ToNot(BeNil())
			vpnGatewayID = *vpnGateway.ID
		})
		It(`GetVPNGateway request example`, func() {
			fmt.Println("\nGetVPNGateway() result:")
			// begin-get_vpn_gateway

			options := vpcService.NewGetVPNGatewayOptions(vpnGatewayID)
			vpnGateway, response, err := vpcService.GetVPNGateway(options)

			// end-get_vpn_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`UpdateVPNGateway request example`, func() {
			fmt.Println("\nUpdateVPNGateway() result:")
			name := getName("vpngateway")
			// begin-update_vpn_gateway

			options := &vpcbetav1.UpdateVPNGatewayOptions{
				ID: &vpnGatewayID,
			}
			vpnGatewayPatchModel := &vpcbetav1.VPNGatewayPatch{
				Name: &name,
			}
			vpnGatewayPatch, asPatchErr := vpnGatewayPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VPNGatewayPatch = vpnGatewayPatch
			vpnGateway, response, err := vpcService.UpdateVPNGateway(options)
			// end-update_vpn_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGateway).ToNot(BeNil())

		})
		It(`ListVPNGatewayConnections request example`, func() {
			fmt.Println("\nListVPNGatewayConnections() result:")
			// begin-list_vpn_gateway_connections

			options := &vpcbetav1.ListVPNGatewayConnectionsOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			vpnGatewayConnections, response, err := vpcService.ListVPNGatewayConnections(
				options,
			)

			// end-list_vpn_gateway_connections
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnections).ToNot(BeNil())

		})
		It(`CreateVPNGatewayConnection request example`, func() {
			fmt.Println("\nCreateVPNGatewayConnection() result:")
			name := getName("vpnconnection")
			// begin-create_vpn_gateway_connection

			options := &vpcbetav1.CreateVPNGatewayConnectionOptions{
				VPNGatewayConnectionPrototype: &vpcbetav1.VPNGatewayConnectionPrototypeVPNGatewayConnectionPolicyModePrototype{
					PeerAddress: &[]string{"192.132.5.0"}[0],
					Psk:         &[]string{"lkj14b1oi0alcniejkso"}[0],
					Name:        &name,
					PeerCIDRs:   []string{"197.155.0.0/28"},
					LocalCIDRs:  []string{"192.132.0.0/28"},
				},
				VPNGatewayID: &vpnGatewayID,
			}
			vpnGatewayConnectionIntf, response, err := vpcService.CreateVPNGatewayConnection(
				options,
			)
			vpnGatewayConnection := vpnGatewayConnectionIntf.(*vpcbetav1.VPNGatewayConnection)
			// end-create_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnGatewayConnection).ToNot(BeNil())
			vpnGatewayConnectionID = *vpnGatewayConnection.ID
		})
		It(`GetVPNGatewayConnection request example`, func() {
			fmt.Println("\nGetVPNGatewayConnection() result:")
			// begin-get_vpn_gateway_connection

			options := &vpcbetav1.GetVPNGatewayConnectionOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			vpnGatewayConnection, response, err := vpcService.GetVPNGatewayConnection(options)

			// end-get_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`UpdateVPNGatewayConnection request example`, func() {
			fmt.Println("\nUpdateVPNGatewayConnection() result:")
			name := getName("vpnConnection")
			// begin-update_vpn_gateway_connection
			options := &vpcbetav1.UpdateVPNGatewayConnectionOptions{
				ID:           &vpnGatewayConnectionID,
				VPNGatewayID: &vpnGatewayID,
			}
			vpnGatewayConnectionPatchModel := &vpcbetav1.VPNGatewayConnectionPatch{}
			vpnGatewayConnectionPatchModel.Name = &name
			vpnGatewayConnectionPatchModel.PeerAddress = &[]string{"192.132.5.0"}[0]
			vpnGatewayConnectionPatchModel.Psk = &[]string{"lkj14b1oi0alcniejkso"}[0]
			vpnGatewayConnectionPatch, asPatchErr := vpnGatewayConnectionPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.VPNGatewayConnectionPatch = vpnGatewayConnectionPatch
			vpnGatewayConnection, response, err := vpcService.UpdateVPNGatewayConnection(
				options,
			)

			// end-update_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnGatewayConnection).ToNot(BeNil())

		})
		It(`AddVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-add_vpn_gateway_connection_local_cidr

			options := &vpcbetav1.AddVPNGatewayConnectionLocalCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.134.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.AddVPNGatewayConnectionLocalCIDR(options)

			// end-add_vpn_gateway_connection_local_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nAddVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListVPNGatewayConnectionLocalCidrs request example`, func() {
			fmt.Println("\nListVPNGatewayConnectionLocalCidrs() result:")
			// begin-list_vpn_gateway_connection_local_cidrs

			options := &vpcbetav1.ListVPNGatewayConnectionLocalCIDRsOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			localCIDRs, response, err :=
				vpcService.ListVPNGatewayConnectionLocalCIDRs(options)

			// end-list_vpn_gateway_connection_local_cidrs
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(localCIDRs).ToNot(BeNil())

		})
		It(`AddVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-add_vpn_gateway_connection_peer_cidr

			options := &vpcbetav1.AddVPNGatewayConnectionPeerCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.144.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.AddVPNGatewayConnectionPeerCIDR(options)

			// end-add_vpn_gateway_connection_peer_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nAddVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`CheckVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-check_vpn_gateway_connection_local_cidr

			options := &vpcbetav1.CheckVPNGatewayConnectionLocalCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.134.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.CheckVPNGatewayConnectionLocalCIDR(options)

			// end-check_vpn_gateway_connection_local_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nCheckVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListVPNGatewayConnectionPeerCidrs request example`, func() {
			fmt.Println("\nListVPNGatewayConnectionPeerCidrs() result:")
			// begin-list_vpn_gateway_connection_peer_cidrs

			options := &vpcbetav1.ListVPNGatewayConnectionPeerCIDRsOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			peerCIDRs, response, err :=
				vpcService.ListVPNGatewayConnectionPeerCIDRs(options)

			// end-list_vpn_gateway_connection_peer_cidrs
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(peerCIDRs).ToNot(BeNil())

		})
		It(`CheckVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-check_vpn_gateway_connection_peer_cidr

			options := &vpcbetav1.CheckVPNGatewayConnectionPeerCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.144.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.CheckVPNGatewayConnectionPeerCIDR(options)
			// end-check_vpn_gateway_connection_peer_cidr
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nCheckVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListBareMetalServerProfiles request example`, func() {
			fmt.Println("\nListBareMetalServerProfiles() result:")
			// begin-list_bare_metal_server_profiles

			listBareMetalServerProfilesOptions := vpcService.NewListBareMetalServerProfilesOptions()

			bareMetalServerProfileCollection, response, err := vpcService.ListBareMetalServerProfiles(listBareMetalServerProfilesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerProfileCollection).ToNot(BeNil())
			bareMetalServerProfileName = *bareMetalServerProfileCollection.Profiles[0].Name

		})
		It(`GetBareMetalServerProfile request example`, func() {
			fmt.Println("\nGetBareMetalServerProfile() result:")
			// begin-get_bare_metal_server_profile

			getBareMetalServerProfileOptions := &vpcbetav1.GetBareMetalServerProfileOptions{
				Name: &bareMetalServerProfileName,
			}

			bareMetalServerProfile, response, err := vpcService.GetBareMetalServerProfile(getBareMetalServerProfileOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_profile

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerProfile).ToNot(BeNil())

		})
		It(`ListBareMetalServers request example`, func() {
			fmt.Println("\nListBareMetalServers() result:")
			// begin-list_bare_metal_servers

			listBareMetalServersOptions := &vpcbetav1.ListBareMetalServersOptions{}

			bareMetalServerCollection, response, err := vpcService.ListBareMetalServers(listBareMetalServersOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_servers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerCollection).ToNot(BeNil())

		})
		It(`CreateBareMetalServer request example`, func() {
			fmt.Println("\nCreateBareMetalServer() result:")
			// begin-create_bare_metal_server

			imageIdentityModel := &vpcbetav1.ImageIdentityByID{
				ID: &imageID,
			}

			keyIdentityModel := &vpcbetav1.KeyIdentityByID{
				ID: &keyID,
			}

			bareMetalServerInitializationPrototypeModel := &vpcbetav1.BareMetalServerInitializationPrototype{
				Image: imageIdentityModel,
				Keys:  []vpcbetav1.KeyIdentityIntf{keyIdentityModel},
			}

			subnetIdentityModel := &vpcbetav1.SubnetIdentityByID{
				ID: &subnetID,
			}

			bareMetalServerPrimaryNetworkInterfacePrototypeModel := &vpcbetav1.BareMetalServerPrimaryNetworkInterfacePrototype{
				Subnet: subnetIdentityModel,
			}

			bareMetalServerProfileIdentityModel := &vpcbetav1.BareMetalServerProfileIdentityByName{
				Name: &bareMetalServerProfileName,
			}

			zoneIdentityModel := &vpcbetav1.ZoneIdentityByName{
				Name: zone,
			}
			bareMetalServerPrototype := &vpcbetav1.BareMetalServerPrototype{
				Initialization:          bareMetalServerInitializationPrototypeModel,
				PrimaryNetworkInterface: bareMetalServerPrimaryNetworkInterfacePrototypeModel,
				Profile:                 bareMetalServerProfileIdentityModel,
				Zone:                    zoneIdentityModel,
				Name:                    &[]string{"my-bare-metal-server"}[0],
			}
			createBareMetalServerOptions := &vpcbetav1.CreateBareMetalServerOptions{}
			createBareMetalServerOptions.BareMetalServerPrototype = bareMetalServerPrototype

			bareMetalServer, response, err := vpcService.CreateBareMetalServer(createBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-create_bare_metal_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bareMetalServer).ToNot(BeNil())
			bareMetalServerId = *bareMetalServer.ID
		})
		It(`CreateBareMetalServerConsoleAccessToken request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nCreateBareMetalServerConsoleAccessToken() result:")
			// begin-create_bare_metal_server_console_access_token

			createBareMetalServerConsoleAccessTokenOptions := &vpcbetav1.CreateBareMetalServerConsoleAccessTokenOptions{
				BareMetalServerID: &bareMetalServerId,
			}
			createBareMetalServerConsoleAccessTokenOptions.SetConsoleType("serial")

			bareMetalServerConsoleAccessToken, response, err := vpcService.CreateBareMetalServerConsoleAccessToken(createBareMetalServerConsoleAccessTokenOptions)
			if err != nil {
				panic(err)
			}

			// end-create_bare_metal_server_console_access_token

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerConsoleAccessToken).ToNot(BeNil())

		})
		It(`ListBareMetalServerDisks request example`, func() {
			fmt.Println("\nListBareMetalServerDisks() result:")
			// begin-list_bare_metal_server_disks

			listBareMetalServerDisksOptions := &vpcbetav1.ListBareMetalServerDisksOptions{
				BareMetalServerID: &bareMetalServerId,
			}

			bareMetalServerDiskCollection, response, err := vpcService.ListBareMetalServerDisks(listBareMetalServerDisksOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_disks

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDiskCollection).ToNot(BeNil())
			bareMetalServerDiskId = *bareMetalServerDiskCollection.Disks[0].ID
		})
		It(`GetBareMetalServerDisk request example`, func() {
			fmt.Println("\nGetBareMetalServerDisk() result:")
			// begin-get_bare_metal_server_disk

			getBareMetalServerDiskOptions := &vpcbetav1.GetBareMetalServerDiskOptions{
				BareMetalServerID: &bareMetalServerId,
				ID:                &bareMetalServerDiskId,
			}

			bareMetalServerDisk, response, err := vpcService.GetBareMetalServerDisk(getBareMetalServerDiskOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDisk).ToNot(BeNil())

		})
		It(`UpdateBareMetalServerDisk request example`, func() {
			fmt.Println("\nUpdateBareMetalServerDisk() result:")
			// begin-update_bare_metal_server_disk

			bareMetalServerDiskPatchModel := &vpcbetav1.BareMetalServerDiskPatch{}
			bareMetalServerDiskPatchModel.Name = &[]string{"my-bare-metal-server-disk-update"}[0]

			bareMetalServerDiskPatchModelAsPatch, asPatchErr := bareMetalServerDiskPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerDiskOptions := &vpcbetav1.UpdateBareMetalServerDiskOptions{
				BareMetalServerID:        &bareMetalServerId,
				ID:                       &bareMetalServerDiskId,
				BareMetalServerDiskPatch: bareMetalServerDiskPatchModelAsPatch,
			}

			bareMetalServerDisk, response, err := vpcService.UpdateBareMetalServerDisk(updateBareMetalServerDiskOptions)
			if err != nil {
				panic(err)
			}

			// end-update_bare_metal_server_disk

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerDisk).ToNot(BeNil())

		})
		It(`ListBareMetalServerNetworkInterfaces request example`, func() {
			fmt.Println("\nListBareMetalServerNetworkInterfaces() result:")
			// begin-list_bare_metal_server_network_interfaces

			listBareMetalServerNetworkInterfacesOptions := &vpcbetav1.ListBareMetalServerNetworkInterfacesOptions{
				BareMetalServerID: &bareMetalServerId,
			}

			bareMetalServerNetworkInterfaceCollection, response, err := vpcService.ListBareMetalServerNetworkInterfaces(listBareMetalServerNetworkInterfacesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_network_interfaces

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterfaceCollection).ToNot(BeNil())

		})
		It(`CreateBareMetalServerNetworkInterface request example`, func() {
			fmt.Println("\nCreateBareMetalServerNetworkInterface() result:")
			// begin-create_bare_metal_server_network_interface

			subnetIdentityModel := &vpcbetav1.SubnetIdentityByID{
				ID: &subnetID,
			}

			bareMetalServerNetworkInterfacePrototypeModel := &vpcbetav1.BareMetalServerNetworkInterfacePrototypeBareMetalServerNetworkInterfaceByPciPrototype{
				InterfaceType: core.StringPtr("pci"),
				Subnet:        subnetIdentityModel,
				Name:          core.StringPtr("my-metal-server-nic"),
			}

			createBareMetalServerNetworkInterfaceOptions := &vpcbetav1.CreateBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID:                        &bareMetalServerId,
				BareMetalServerNetworkInterfacePrototype: bareMetalServerNetworkInterfacePrototypeModel,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.CreateBareMetalServerNetworkInterface(createBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-create_bare_metal_server_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())
			switch reflect.TypeOf(bareMetalServerNetworkInterface).String() {
			case "*vpcbetav1.BareMetalServerNetworkInterfaceByPci":
				{
					nic := bareMetalServerNetworkInterface.(*vpcbetav1.BareMetalServerNetworkInterfaceByPci)
					bareMetalServerNetworkInterfaceId = *nic.ID
				}
			}
		})
		It(`GetBareMetalServerNetworkInterface request example`, func() {
			fmt.Println("\nGetBareMetalServerNetworkInterface() result:")
			// begin-get_bare_metal_server_network_interface

			getBareMetalServerNetworkInterfaceOptions := &vpcbetav1.GetBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: &bareMetalServerId,
				ID:                &bareMetalServerNetworkInterfaceId,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.GetBareMetalServerNetworkInterface(getBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
		It(`UpdateBareMetalServerNetworkInterface request example`, func() {
			fmt.Println("\nUpdateBareMetalServerNetworkInterface() result:")
			// begin-update_bare_metal_server_network_interface

			bareMetalServerNetworkInterfacePatchModel := &vpcbetav1.BareMetalServerNetworkInterfacePatch{
				Name: core.StringPtr("my-metal-server-nic-update"),
			}
			bareMetalServerNetworkInterfacePatchModelAsPatch, asPatchErr := bareMetalServerNetworkInterfacePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerNetworkInterfaceOptions := &vpcbetav1.UpdateBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID:                    &bareMetalServerId,
				ID:                                   &bareMetalServerNetworkInterfaceId,
				BareMetalServerNetworkInterfacePatch: bareMetalServerNetworkInterfacePatchModelAsPatch,
			}

			bareMetalServerNetworkInterface, response, err := vpcService.UpdateBareMetalServerNetworkInterface(updateBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}

			// end-update_bare_metal_server_network_interface

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerNetworkInterface).ToNot(BeNil())

		})
		It(`ListBareMetalServerNetworkInterfaceFloatingIps request example`, func() {
			fmt.Println("\nListBareMetalServerNetworkInterfaceFloatingIps() result:")
			// begin-list_bare_metal_server_network_interface_floating_ips

			listBareMetalServerNetworkInterfaceFloatingIpsOptions := &vpcbetav1.ListBareMetalServerNetworkInterfaceFloatingIpsOptions{
				BareMetalServerID:  &bareMetalServerId,
				NetworkInterfaceID: &bareMetalServerNetworkInterfaceId,
			}

			floatingIPUnpaginatedCollection, response, err := vpcService.ListBareMetalServerNetworkInterfaceFloatingIps(listBareMetalServerNetworkInterfaceFloatingIpsOptions)
			if err != nil {
				panic(err)
			}

			// end-list_bare_metal_server_network_interface_floating_ips

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIPUnpaginatedCollection).ToNot(BeNil())

		})
		It(`AddBareMetalServerNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nAddBareMetalServerNetworkInterfaceFloatingIP() result:")
			// begin-add_bare_metal_server_network_interface_floating_ip

			addBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcbetav1.AddBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  &bareMetalServerId,
				NetworkInterfaceID: &bareMetalServerNetworkInterfaceId,
				ID:                 &floatingIPID,
			}

			floatingIP, response, err := vpcService.AddBareMetalServerNetworkInterfaceFloatingIP(addBareMetalServerNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}

			// end-add_bare_metal_server_network_interface_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`GetBareMetalServerNetworkInterfaceFloatingIP request example`, func() {
			fmt.Println("\nGetBareMetalServerNetworkInterfaceFloatingIP() result:")
			// begin-get_bare_metal_server_network_interface_floating_ip

			getBareMetalServerNetworkInterfaceFloatingIPOptions := &vpcbetav1.GetBareMetalServerNetworkInterfaceFloatingIPOptions{
				BareMetalServerID:  &bareMetalServerId,
				NetworkInterfaceID: &bareMetalServerNetworkInterfaceId,
				ID:                 &floatingIPID,
			}

			floatingIP, response, err := vpcService.GetBareMetalServerNetworkInterfaceFloatingIP(getBareMetalServerNetworkInterfaceFloatingIPOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_network_interface_floating_ip

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(floatingIP).ToNot(BeNil())

		})
		It(`GetBareMetalServer request example`, func() {
			fmt.Println("\nGetBareMetalServer() result:")
			// begin-get_bare_metal_server

			getBareMetalServerOptions := &vpcbetav1.GetBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			bareMetalServer, response, err := vpcService.GetBareMetalServer(getBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServer).ToNot(BeNil())

		})
		It(`UpdateBareMetalServer request example`, func() {
			fmt.Println("\nUpdateBareMetalServer() result:")
			// begin-update_bare_metal_server

			bareMetalServerPatchModel := &vpcbetav1.BareMetalServerPatch{
				Name: core.StringPtr("my-metal-server-update"),
			}
			bareMetalServerPatchModelAsPatch, asPatchErr := bareMetalServerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBareMetalServerOptions := &vpcbetav1.UpdateBareMetalServerOptions{
				ID:                   &bareMetalServerId,
				BareMetalServerPatch: bareMetalServerPatchModelAsPatch,
			}

			bareMetalServer, response, err := vpcService.UpdateBareMetalServer(updateBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-update_bare_metal_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServer).ToNot(BeNil())

		})
		It(`GetBareMetalServerInitialization request example`, func() {
			fmt.Println("\nGetBareMetalServerInitialization() result:")
			// begin-get_bare_metal_server_initialization

			getBareMetalServerInitializationOptions := &vpcbetav1.GetBareMetalServerInitializationOptions{
				ID: &bareMetalServerId,
			}

			bareMetalServerInitialization, response, err := vpcService.GetBareMetalServerInitialization(getBareMetalServerInitializationOptions)
			if err != nil {
				panic(err)
			}

			// end-get_bare_metal_server_initialization

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(bareMetalServerInitialization).ToNot(BeNil())

		})
		It(`RestartBareMetalServer request example`, func() {
			// begin-restart_bare_metal_server

			restartBareMetalServerOptions := &vpcbetav1.RestartBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			response, err := vpcService.RestartBareMetalServer(restartBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-restart_bare_metal_server
			fmt.Printf("\nRestartBareMetalServer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`StartBareMetalServer request example`, func() {
			// begin-start_bare_metal_server

			startBareMetalServerOptions := &vpcbetav1.StartBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			response, err := vpcService.StartBareMetalServer(startBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-start_bare_metal_server
			fmt.Printf("\nStartBareMetalServer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`StopBareMetalServer request example`, func() {
			// begin-stop_bare_metal_server

			stopBareMetalServerOptions := &vpcbetav1.StopBareMetalServerOptions{
				ID:   &bareMetalServerId,
				Type: core.StringPtr("soft"),
			}

			response, err := vpcService.StopBareMetalServer(stopBareMetalServerOptions)
			if err != nil {
				panic(err)
			}

			// end-stop_bare_metal_server
			fmt.Printf("\nStopBareMetalServer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListBackupPolicies request example`, func() {
			fmt.Println("\nListBackupPolicies() result:")
			// begin-list_backup_policies

			listBackupPoliciesOptions := vpcService.NewListBackupPoliciesOptions()

			backupPolicyCollection, response, err := vpcService.ListBackupPolicies(listBackupPoliciesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_backup_policies

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyCollection).ToNot(BeNil())

		})
		It(`CreateBackupPolicy request example`, func() {
			fmt.Println("\nCreateBackupPolicy() result:")
			// begin-create_backup_policy

			userTags := []string{"tag1", "tag2"}
			name := "my-backup-policy"
			matchResourceType := "instance"
			backupPolicyPrototype := &vpcbetav1.BackupPolicyPrototype{
				MatchUserTags:     userTags,
				Name:              &name,
				MatchResourceType: &matchResourceType,
			}
			createBackupPolicyOptions := vpcService.NewCreateBackupPolicyOptions(backupPolicyPrototype)
			backupPolicyIntf, response, err := vpcService.CreateBackupPolicy(createBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-create_backup_policy
			backupPolicy := backupPolicyIntf.(*vpcbetav1.BackupPolicy)
			backupPolicyID = *backupPolicy.ID

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(backupPolicy).ToNot(BeNil())

		})
		It(`CreateBackupPolicyPlan request example`, func() {
			fmt.Println("\nCreateBackupPolicyPlan() result:")
			// begin-create_backup_policy_plan

			createBackupPolicyPlanOptions := vpcService.NewCreateBackupPolicyPlanOptions(
				backupPolicyID,
				"*/5 1,2,3 * * *",
			)
			createBackupPolicyPlanOptions.SetName("my-backup-policy-plan")

			backupPolicyPlan, response, err := vpcService.CreateBackupPolicyPlan(createBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-create_backup_policy_plan
			backupPolicyPlanID = *backupPolicyPlan.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`ListBackupPolicyPlans request example`, func() {
			fmt.Println("\nListBackupPolicyPlans() result:")
			// begin-list_backup_policy_plans

			listBackupPolicyPlansOptions := vpcService.NewListBackupPolicyPlansOptions(
				backupPolicyID,
			)

			backupPolicyPlanCollection, response, err := vpcService.ListBackupPolicyPlans(listBackupPolicyPlansOptions)
			if err != nil {
				panic(err)
			}

			// end-list_backup_policy_plans

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyPlanCollection).ToNot(BeNil())

		})

		It(`GetBackupPolicyPlan request example`, func() {
			fmt.Println("\nGetBackupPolicyPlan() result:")
			// begin-get_backup_policy_plan

			getBackupPolicyPlanOptions := vpcService.NewGetBackupPolicyPlanOptions(
				backupPolicyID,
				backupPolicyPlanID,
			)

			backupPolicyPlan, response, err := vpcService.GetBackupPolicyPlan(getBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-get_backup_policy_plan
			ifMatchBackupPolicy = response.GetHeaders()["Etag"][0]
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`UpdateBackupPolicyPlan request example`, func() {
			fmt.Println("\nUpdateBackupPolicyPlan() result:")
			// begin-update_backup_policy_plan

			backupPolicyPlanPatchModel := &vpcbetav1.BackupPolicyPlanPatch{
				Name: core.StringPtr("my-backup-plan-updated"),
			}
			backupPolicyPlanPatchModelAsPatch, asPatchErr := backupPolicyPlanPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBackupPolicyPlanOptions := vpcService.NewUpdateBackupPolicyPlanOptions(
				backupPolicyID,
				backupPolicyPlanID,
				backupPolicyPlanPatchModelAsPatch,
			)
			updateBackupPolicyPlanOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicyPlan, response, err := vpcService.UpdateBackupPolicyPlan(updateBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-update_backup_policy_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`GetBackupPolicy request example`, func() {
			fmt.Println("\nGetBackupPolicy() result:")
			// begin-get_backup_policy

			getBackupPolicyOptions := vpcService.NewGetBackupPolicyOptions(
				backupPolicyID,
			)

			backupPolicy, response, err := vpcService.GetBackupPolicy(getBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-get_backup_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicy).ToNot(BeNil())

		})
		It(`UpdateBackupPolicy request example`, func() {
			fmt.Println("\nUpdateBackupPolicy() result:")
			// begin-update_backup_policy

			backupPolicyPatchModel := &vpcbetav1.BackupPolicyPatch{
				Name: core.StringPtr("my-backup-policy-update"),
			}
			backupPolicyPatchModelAsPatch, asPatchErr := backupPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateBackupPolicyOptions := vpcService.NewUpdateBackupPolicyOptions(
				backupPolicyID,
				backupPolicyPatchModelAsPatch,
			)
			updateBackupPolicyOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicy, response, err := vpcService.UpdateBackupPolicy(updateBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-update_backup_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(backupPolicy).ToNot(BeNil())

		})

		It(`ListPlacementGroups request example`, func() {
			fmt.Println("\nListPlacementGroups() result:")
			// begin-list_placement_groups

			options := &vpcbetav1.ListPlacementGroupsOptions{}
			placementGroups, response, err := vpcService.ListPlacementGroups(options)

			// end-list_flow_log_collectors
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroups).ToNot(BeNil())

		})
		It(`CreatePlacementGroup request example`, func() {
			fmt.Println("\nCreatePlacementGroup() result:")
			name := getName("placement")
			// begin-create_flow_log_collector

			strategy := "host_spread"
			createPlacementGroupOptions := &vpcbetav1.CreatePlacementGroupOptions{
				Strategy: &strategy,
				Name:     &name,
			}
			placementGroup, response, err := vpcService.CreatePlacementGroup(createPlacementGroupOptions)

			// end-create_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(placementGroup).ToNot(BeNil())
			placementGroupID = *placementGroup.ID
		})
		It(`GetPlacementGroup request example`, func() {
			fmt.Println("\nGetPlacementGroup() result:")
			// begin-get_flow_log_collector

			getPlacementGroupOptions := &vpcbetav1.GetPlacementGroupOptions{
				ID: &placementGroupID,
			}

			placementGroup, response, err := vpcService.GetPlacementGroup(getPlacementGroupOptions)

			// end-get_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())

		})

		It(`UpdatePlacementGroup request example`, func() {
			fmt.Println("\nUpdatePlacementGroup() result:")
			name := getName("fl")
			// begin-update_flow_log_collector

			placementGroupPatchModel := &vpcbetav1.PlacementGroupPatch{
				Name: &name,
			}
			placementGroupPatchModelAsPatch, asPatchErr := placementGroupPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}

			options := &vpcbetav1.UpdatePlacementGroupOptions{
				ID:                  &placementGroupID,
				PlacementGroupPatch: placementGroupPatchModelAsPatch,
			}

			placementGroup, response, err := vpcService.UpdatePlacementGroup(options)

			// end-update_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(placementGroup).ToNot(BeNil())

		})

		It(`DeletePlacementGroup request example`, func() {
			// begin-delete_flow_log_collector

			deletePlacementGroupOptions := &vpcbetav1.DeletePlacementGroupOptions{
				ID: &placementGroupID,
			}

			response, err := vpcService.DeletePlacementGroup(deletePlacementGroupOptions)

			// end-delete_flow_log_collector
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeletePlacementGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListVPNServers request example`, func() {
			fmt.Println("\nListVPNServers() result:")
			// begin-list_vpn_servers

			listVPNServersOptions := vpcService.NewListVPNServersOptions()
			listVPNServersOptions.SetSort("name")

			vpnServerCollection, response, err := vpcService.ListVPNServers(listVPNServersOptions)
			if err != nil {
				panic(err)
			}

			// end-list_vpn_servers

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerCollection).ToNot(BeNil())

		})
		It(`CreateVPNServer request example`, func() {
			fmt.Println("\nCreateVPNServer() result:")
			// begin-create_vpn_server

			certificateInstanceIdentityModel := &vpcbetav1.CertificateInstanceIdentityByCRN{
				CRN: core.StringPtr("crn:v1:bluemix:public:secrets-manager:us-south:a/123456:36fa422d-080d-4d83-8d2d-86851b4001df:secret:2e786aab-42fa-63ed-14f8-d66d552f4dd5"),
			}

			vpnServerAuthenticationByUsernameIDProviderModel := &vpcbetav1.VPNServerAuthenticationByUsernameIDProviderByIam{
				ProviderType: core.StringPtr("iam"),
			}

			vpnServerAuthenticationPrototypeModel := &vpcbetav1.VPNServerAuthenticationPrototypeVPNServerAuthenticationByUsernamePrototype{
				Method:           core.StringPtr("certificate"),
				IdentityProvider: vpnServerAuthenticationByUsernameIDProviderModel,
			}

			subnetIdentityModel := &vpcbetav1.SubnetIdentityByID{
				ID: core.StringPtr(subnetID),
			}

			createVPNServerOptions := vpcService.NewCreateVPNServerOptions(
				certificateInstanceIdentityModel,
				[]vpcbetav1.VPNServerAuthenticationPrototypeIntf{vpnServerAuthenticationPrototypeModel},
				"172.16.0.0/16",
				[]vpcbetav1.SubnetIdentityIntf{subnetIdentityModel},
			)
			createVPNServerOptions.SetName("my-vpn-server")

			vpnServer, response, err := vpcService.CreateVPNServer(createVPNServerOptions)
			if err != nil {
				panic(err)
			}

			// end-create_vpn_server
			vpnServerID = *vpnServer.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnServer).ToNot(BeNil())

		})
		It(`GetVPNServer request example`, func() {
			fmt.Println("\nGetVPNServer() result:")
			// begin-get_vpn_server

			getVPNServerOptions := vpcService.NewGetVPNServerOptions(
				vpnServerID,
			)

			vpnServer, response, err := vpcService.GetVPNServer(getVPNServerOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server
			ifMatchVPNServer = response.GetHeaders()["Etag"][0]

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServer).ToNot(BeNil())

		})
		It(`UpdateVPNServer request example`, func() {
			fmt.Println("\nUpdateVPNServer() result:")
			// begin-update_vpn_server

			vpnServerPatchModel := &vpcbetav1.VPNServerPatch{
				Name: &[]string{"my-vpn-server-modified"}[0],
			}
			vpnServerPatchModelAsPatch, asPatchErr := vpnServerPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNServerOptions := vpcService.NewUpdateVPNServerOptions(
				vpnServerID,
				vpnServerPatchModelAsPatch,
			)
			updateVPNServerOptions.SetIfMatch(ifMatchVPNServer)

			vpnServer, response, err := vpcService.UpdateVPNServer(updateVPNServerOptions)
			if err != nil {
				panic(err)
			}

			// end-update_vpn_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServer).ToNot(BeNil())

		})
		It(`GetVPNServerClientConfiguration request example`, func() {
			fmt.Println("\nGetVPNServerClientConfiguration() result:")
			// begin-get_vpn_server_client_configuration

			getVPNServerClientConfigurationOptions := vpcService.NewGetVPNServerClientConfigurationOptions(
				vpnServerID,
			)

			vpnServerClientConfiguration, response, err := vpcService.GetVPNServerClientConfiguration(getVPNServerClientConfigurationOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server_client_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerClientConfiguration).ToNot(BeNil())

		})
		It(`ListVPNServerClients request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nListVPNServerClients() result:")
			// begin-list_vpn_server_clients

			listVPNServerClientsOptions := vpcService.NewListVPNServerClientsOptions(
				vpnServerID,
			)
			listVPNServerClientsOptions.SetSort("created_at")

			vpnServerClientCollection, response, err := vpcService.ListVPNServerClients(listVPNServerClientsOptions)
			if err != nil {
				panic(err)
			}

			// end-list_vpn_server_clients
			vpnClientID = *vpnServerClientCollection.Clients[0].ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerClientCollection).ToNot(BeNil())

		})
		It(`GetVPNServerClient request example`, func() {
			Skip("not runnin with mock")
			fmt.Println("\nGetVPNServerClient() result:")
			// begin-get_vpn_server_client

			getVPNServerClientOptions := vpcService.NewGetVPNServerClientOptions(
				vpnServerID,
				vpnClientID,
			)

			vpnServerClient, response, err := vpcService.GetVPNServerClient(getVPNServerClientOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server_client

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerClient).ToNot(BeNil())

		})
		It(`DisconnectVPNClient request example`, func() {
			Skip("not runnin with mock")
			// begin-disconnect_vpn_client

			disconnectVPNClientOptions := vpcService.NewDisconnectVPNClientOptions(
				vpnServerID,
				vpnClientID,
			)

			response, err := vpcService.DisconnectVPNClient(disconnectVPNClientOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DisconnectVPNClient(): %d\n", response.StatusCode)
			}

			// end-disconnect_vpn_client

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`ListVPNServerRoutes request example`, func() {
			fmt.Println("\nListVPNServerRoutes() result:")
			// begin-list_vpn_server_routes

			listVPNServerRoutesOptions := vpcService.NewListVPNServerRoutesOptions(
				vpnServerID,
			)
			listVPNServerRoutesOptions.SetSort("name")

			vpnServerRouteCollection, response, err := vpcService.ListVPNServerRoutes(listVPNServerRoutesOptions)
			if err != nil {
				panic(err)
			}

			// end-list_vpn_server_routes

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerRouteCollection).ToNot(BeNil())

		})
		It(`CreateVPNServerRoute request example`, func() {
			fmt.Println("\nCreateVPNServerRoute() result:")
			// begin-create_vpn_server_route

			createVPNServerRouteOptions := vpcService.NewCreateVPNServerRouteOptions(
				vpnServerID,
				"172.16.0.0/16",
			)
			createVPNServerRouteOptions.SetName("my-vpn-server-route")

			vpnServerRoute, response, err := vpcService.CreateVPNServerRoute(createVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-create_vpn_server_route
			vpnServerRouteID = *vpnServerRoute.ID
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(vpnServerRoute).ToNot(BeNil())

		})
		It(`GetVPNServerRoute request example`, func() {
			fmt.Println("\nGetVPNServerRoute() result:")
			// begin-get_vpn_server_route

			getVPNServerRouteOptions := vpcService.NewGetVPNServerRouteOptions(
				vpnServerID,
				vpnServerRouteID,
			)

			vpnServerRoute, response, err := vpcService.GetVPNServerRoute(getVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-get_vpn_server_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerRoute).ToNot(BeNil())

		})
		It(`UpdateVPNServerRoute request example`, func() {
			fmt.Println("\nUpdateVPNServerRoute() result:")
			// begin-update_vpn_server_route

			vpnServerRoutePatchModel := &vpcbetav1.VPNServerRoutePatch{
				Name: &[]string{"my-vpn-server-route-modified"}[0],
			}
			vpnServerRoutePatchModelAsPatch, asPatchErr := vpnServerRoutePatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVPNServerRouteOptions := vpcService.NewUpdateVPNServerRouteOptions(
				vpnServerID,
				vpnServerRouteID,
				vpnServerRoutePatchModelAsPatch,
			)

			vpnServerRoute, response, err := vpcService.UpdateVPNServerRoute(updateVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}

			// end-update_vpn_server_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vpnServerRoute).ToNot(BeNil())

		})
		It(`ListLoadBalancerProfiles request example`, func() {
			fmt.Println("\nListLoadBalancerProfiles() result:")
			// begin-list_load_balancer_profiles

			options := &vpcbetav1.ListLoadBalancerProfilesOptions{}
			profiles, response, err := vpcService.ListLoadBalancerProfiles(options)

			// end-list_load_balancer_profiles
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profiles).ToNot(BeNil())

		})
		It(`GetLoadBalancerProfile request example`, func() {
			fmt.Println("\nGetLoadBalancerProfile() result:")
			// begin-get_load_balancer_profile
			options := &vpcbetav1.GetLoadBalancerProfileOptions{}
			options.SetName("network-fixed")
			profile, response, err := vpcService.GetLoadBalancerProfile(options)
			// end-get_load_balancer_profile
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(profile).ToNot(BeNil())

		})
		It(`ListLoadBalancers request example`, func() {
			fmt.Println("\nListLoadBalancers() result:")
			// begin-list_load_balancers

			options := &vpcbetav1.ListLoadBalancersOptions{}
			loadBalancers, response, err := vpcService.ListLoadBalancers(options)

			// end-list_load_balancers
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancers).ToNot(BeNil())

		})
		It(`CreateLoadBalancer request example`, func() {
			fmt.Println("\nCreateLoadBalancer() result:")
			name := getName("lb")
			// begin-create_load_balancer
			loadBalancerProfileIdentityModel := &vpcbetav1.LoadBalancerProfileIdentityByName{
				Name: core.StringPtr("network-private-path"),
			}
			options := &vpcbetav1.CreateLoadBalancerOptions{
				IsPublic:      &[]bool{false}[0],
				IsPrivatePath: &[]bool{true}[0],
				Name:          &name,
				Profile:       loadBalancerProfileIdentityModel,
				Subnets: []vpcbetav1.SubnetIdentityIntf{
					&vpcbetav1.SubnetIdentity{
						ID: &subnetID,
					},
				},
			}
			loadBalancer, response, err := vpcService.CreateLoadBalancer(options)
			// end-create_load_balancer

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(loadBalancer).ToNot(BeNil())
			loadBalancerID = *loadBalancer.ID
		})
		It(`GetLoadBalancer request example`, func() {
			fmt.Println("\nGetLoadBalancer() result:")
			// begin-get_load_balancer

			options := &vpcbetav1.GetLoadBalancerOptions{
				ID: &loadBalancerID,
			}
			loadBalancer, response, err := vpcService.GetLoadBalancer(options)

			// end-get_load_balancer
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`UpdateLoadBalancer request example`, func() {
			fmt.Println("\nUpdateLoadBalancer() result:")
			name := getName("lb")
			// begin-update_load_balancer

			loadBalancerPatchModel := &vpcbetav1.LoadBalancerPatch{
				Name: &name,
			}
			loadBalancerPatchModelAsPatch, asPatchErr := loadBalancerPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			updateLoadBalancerOptions := vpcService.NewUpdateLoadBalancerOptions(
				loadBalancerID,
				loadBalancerPatchModelAsPatch,
			)

			loadBalancer, response, err := vpcService.UpdateLoadBalancer(updateLoadBalancerOptions)

			// end-update_load_balancer
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(loadBalancer).ToNot(BeNil())

		})
		It(`GetLoadBalancerStatistics request example`, func() {
			fmt.Println("\nGetLoadBalancerStatistics() result:")
			// begin-get_load_balancer_statistics

			options := &vpcbetav1.GetLoadBalancerStatisticsOptions{
				ID: &loadBalancerID,
			}
			statistics, response, err := vpcService.GetLoadBalancerStatistics(options)
			// end-get_load_balancer_statistics
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(statistics).ToNot(BeNil())

		})
		It(`ListLoadBalancerListeners request example`, func() {
			fmt.Println("\nListLoadBalancerListeners() result:")
			// begin-list_load_balancer_listeners

			options := &vpcbetav1.ListLoadBalancerListenersOptions{
				LoadBalancerID: &loadBalancerID,
			}
			listeners, response, err := vpcService.ListLoadBalancerListeners(options)

			// end-list_load_balancer_listeners
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listeners).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListener request example`, func() {
			fmt.Println("\nCreateLoadBalancerListener() result:")
			// begin-create_load_balancer_listener

			options := &vpcbetav1.CreateLoadBalancerListenerOptions{
				LoadBalancerID: &loadBalancerID,
			}
			options.SetPort(5656)
			options.SetProtocol("http")
			listener, response, err := vpcService.CreateLoadBalancerListener(options)

			// end-create_load_balancer_listener
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(listener).ToNot(BeNil())
			listenerID = *listener.ID
		})
		It(`GetLoadBalancerListener request example`, func() {
			fmt.Println("\nGetLoadBalancerListener() result:")
			// begin-get_load_balancer_listener

			options := &vpcbetav1.GetLoadBalancerListenerOptions{
				LoadBalancerID: &loadBalancerID,
				ID:             &listenerID,
			}
			listener, response, err := vpcService.GetLoadBalancerListener(options)

			// end-get_load_balancer_listener
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listener).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListener request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListener() result:")
			// begin-update_load_balancer_listener

			loadBalancerListenerPatchModel := &vpcbetav1.LoadBalancerListenerPatch{
				Port: &[]int64{5666}[0],
			}
			loadBalancerListenerPatchModelAsPatch, asPatchErr := loadBalancerListenerPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := vpcService.NewUpdateLoadBalancerListenerOptions(
				loadBalancerID,
				listenerID,
				loadBalancerListenerPatchModelAsPatch,
			)

			listener, response, err := vpcService.UpdateLoadBalancerListener(options)

			// end-update_load_balancer_listener
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listener).ToNot(BeNil())

		})
		It(`ListLoadBalancerListenerPolicies request example`, func() {
			fmt.Println("\nListLoadBalancerListenerPolicies() result:")
			// begin-list_load_balancer_listener_policies

			options := &vpcbetav1.ListLoadBalancerListenerPoliciesOptions{
				LoadBalancerID: &loadBalancerID,
				ListenerID:     &listenerID,
			}
			policies, response, err :=
				vpcService.ListLoadBalancerListenerPolicies(options)

			// end-list_load_balancer_listener_policies
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policies).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nCreateLoadBalancerListenerPolicy() result:")
			// begin-create_load_balancer_listener_policy

			options := &vpcbetav1.CreateLoadBalancerListenerPolicyOptions{
				LoadBalancerID: &loadBalancerID,
				ListenerID:     &listenerID,
			}
			options.SetPriority(2)
			options.SetAction("reject")
			policy, response, err :=
				vpcService.CreateLoadBalancerListenerPolicy(options)

			// end-create_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policy).ToNot(BeNil())
			policyID = *policy.ID
		})
		It(`GetLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nGetLoadBalancerListenerPolicy() result:")
			// begin-get_load_balancer_listener_policy

			options := &vpcbetav1.GetLoadBalancerListenerPolicyOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetID(policyID)
			policy, response, err := vpcService.GetLoadBalancerListenerPolicy(options)

			// end-get_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListenerPolicy request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListenerPolicy() result:")
			// begin-update_load_balancer_listener_policy

			options := &vpcbetav1.UpdateLoadBalancerListenerPolicyOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetID(policyID)
			policyPatchModel := &vpcbetav1.LoadBalancerListenerPolicyPatch{}
			policyPatchModel.Priority = &[]int64{5}[0]
			policyPatch, asPatchErr :=
				policyPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.LoadBalancerListenerPolicyPatch = policyPatch
			policy, response, err :=
				vpcService.UpdateLoadBalancerListenerPolicy(options)

			// end-update_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(policy).ToNot(BeNil())
		})
		It(`ListLoadBalancerListenerPolicyRules request example`, func() {
			fmt.Println("\nListLoadBalancerListenerPolicyRules() result:")
			// begin-list_load_balancer_listener_policy_rules

			options := &vpcbetav1.ListLoadBalancerListenerPolicyRulesOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			rules, response, err :=
				vpcService.ListLoadBalancerListenerPolicyRules(options)
			// end-list_load_balancer_listener_policy_rules
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rules).ToNot(BeNil())

		})
		It(`CreateLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nCreateLoadBalancerListenerPolicyRule() result:")
			// begin-create_load_balancer_listener_policy_rule
			options := &vpcbetav1.CreateLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetCondition("contains")
			options.SetType("hostname")
			options.SetValue("one")
			policyRule, response, err :=
				vpcService.CreateLoadBalancerListenerPolicyRule(options)

			// end-create_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(policyRule).ToNot(BeNil())
			policyRuleID = *policyRule.ID
		})
		It(`GetLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nGetLoadBalancerListenerPolicyRule() result:")
			// begin-get_load_balancer_listener_policy_rule

			options := &vpcbetav1.GetLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetID(policyRuleID)
			rule, response, err :=
				vpcService.GetLoadBalancerListenerPolicyRule(options)

			// end-get_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerListenerPolicyRule request example`, func() {
			fmt.Println("\nUpdateLoadBalancerListenerPolicyRule() result:")
			// begin-update_load_balancer_listener_policy_rule

			options := &vpcbetav1.UpdateLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetID(policyRuleID)
			policyRulePatchModel :=
				&vpcbetav1.LoadBalancerListenerPolicyRulePatch{
					Condition: &[]string{"contains"}[0],
					Type:      &[]string{"header"}[0],
					Value:     &[]string{"app"}[0],
					Field:     &[]string{"MY-APP-HEADER"}[0],
				}
			policyRulePatch, asPatchErr := policyRulePatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.LoadBalancerListenerPolicyRulePatch = policyRulePatch
			rule, response, err :=
				vpcService.UpdateLoadBalancerListenerPolicyRule(options)

			// end-update_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(rule).ToNot(BeNil())

		})
		It(`ListLoadBalancerPools request example`, func() {
			fmt.Println("\nListLoadBalancerPools() result:")
			// begin-list_load_balancer_pools
			options := &vpcbetav1.ListLoadBalancerPoolsOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			pools, response, err := vpcService.ListLoadBalancerPools(options)
			// end-list_load_balancer_pools
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pools).ToNot(BeNil())

		})
		It(`CreateLoadBalancerPool request example`, func() {
			fmt.Println("\nCreateLoadBalancerPool() result:")
			name := getName("pool")
			// begin-create_load_balancer_pool

			options := &vpcbetav1.CreateLoadBalancerPoolOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetAlgorithm("round_robin")
			options.SetHealthMonitor(&vpcbetav1.LoadBalancerPoolHealthMonitorPrototype{
				Delay:      &[]int64{30}[0],
				MaxRetries: &[]int64{3}[0],
				Timeout:    &[]int64{30}[0],
				Type:       &[]string{"http"}[0],
			})
			options.SetName(name)
			options.SetProtocol("http")
			pool, response, err := vpcService.CreateLoadBalancerPool(options)

			// end-create_load_balancer_pool
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(pool).ToNot(BeNil())
			poolID = *pool.ID
		})
		It(`GetLoadBalancerPool request example`, func() {
			fmt.Println("\nGetLoadBalancerPool() result:")
			// begin-get_load_balancer_pool

			options := &vpcbetav1.GetLoadBalancerPoolOptions{
				LoadBalancerID: &loadBalancerID,
				ID:             &poolID,
			}
			pool, response, err := vpcService.GetLoadBalancerPool(options)

			// end-get_load_balancer_pool
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pool).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerPool request example`, func() {
			fmt.Println("\nUpdateLoadBalancerPool() result:")
			// begin-update_load_balancer_pool

			options := &vpcbetav1.UpdateLoadBalancerPoolOptions{
				LoadBalancerID: &loadBalancerID,
				ID:             &poolID,
			}
			poolPatchModel := &vpcbetav1.LoadBalancerPoolPatch{}
			healthMonitorPatchModel := &vpcbetav1.LoadBalancerPoolHealthMonitorPatch{
				Delay:      &[]int64{30}[0],
				MaxRetries: &[]int64{3}[0],
				Timeout:    &[]int64{30}[0],
				Type:       &[]string{"http"}[0],
			}
			poolPatchModel.HealthMonitor = healthMonitorPatchModel
			sessionPersistence := &vpcbetav1.LoadBalancerPoolSessionPersistencePatch{
				Type: &[]string{"http_cookie"}[0],
			}
			poolPatchModel.SessionPersistence = sessionPersistence
			LoadBalancerPoolPatch, _ := poolPatchModel.AsPatch()
			options.LoadBalancerPoolPatch = LoadBalancerPoolPatch
			pool, response, err := vpcService.UpdateLoadBalancerPool(options)

			// end-update_load_balancer_pool
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pool).ToNot(BeNil())

		})
		It(`ListLoadBalancerPoolMembers request example`, func() {
			fmt.Println("\nListLoadBalancerPoolMembers() result:")
			// begin-list_load_balancer_pool_members

			options := &vpcbetav1.ListLoadBalancerPoolMembersOptions{
				LoadBalancerID: &loadBalancerID,
				PoolID:         &poolID,
			}
			members, response, err := vpcService.ListLoadBalancerPoolMembers(options)

			// end-list_load_balancer_pool_members
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(members).ToNot(BeNil())

		})
		It(`CreateLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nCreateLoadBalancerPoolMember() result:")
			// begin-create_load_balancer_pool_member

			options := &vpcbetav1.CreateLoadBalancerPoolMemberOptions{
				LoadBalancerID: &loadBalancerID,
				PoolID:         &poolID,
				Port:           &[]int64{1234}[0],
				Target: &vpcbetav1.LoadBalancerPoolMemberTargetPrototypeIP{
					Address: &[]string{"192.168.3.4"}[0],
				},
			}
			member, response, err := vpcService.CreateLoadBalancerPoolMember(options)
			// end-create_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(member).ToNot(BeNil())
			poolMemberID = *member.ID
		})

		It(`GetLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nGetLoadBalancerPoolMember() result:")
			// begin-get_load_balancer_pool_member

			options := &vpcbetav1.GetLoadBalancerPoolMemberOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetID(poolMemberID)
			member, response, err := vpcService.GetLoadBalancerPoolMember(options)

			// end-get_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(member).ToNot(BeNil())

		})
		It(`UpdateLoadBalancerPoolMember request example`, func() {
			fmt.Println("\nUpdateLoadBalancerPoolMember() result:")
			// begin-update_load_balancer_pool_member

			options := &vpcbetav1.UpdateLoadBalancerPoolMemberOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetID(poolMemberID)
			loadBalancerPoolMemberPatchModel := &vpcbetav1.LoadBalancerPoolMemberPatch{
				Port:   &[]int64{1235}[0],
				Weight: &[]int64{50}[0],
			}
			loadBalancerPoolMemberPatch, _ := loadBalancerPoolMemberPatchModel.AsPatch()
			options.LoadBalancerPoolMemberPatch = loadBalancerPoolMemberPatch
			member, response, err := vpcService.UpdateLoadBalancerPoolMember(options)

			// end-update_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(member).ToNot(BeNil())

		})
		It(`ReplaceLoadBalancerPoolMembers request example`, func() {
			fmt.Println("\nReplaceLoadBalancerPoolMembers() result:")
			// begin-replace_load_balancer_pool_members

			options := &vpcbetav1.ReplaceLoadBalancerPoolMembersOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetMembers([]vpcbetav1.LoadBalancerPoolMemberPrototype{
				{
					Port: &[]int64{1235}[0],
					Target: &vpcbetav1.LoadBalancerPoolMemberTargetPrototypeIP{
						Address: &[]string{"192.168.3.5"}[0],
					},
				},
			})
			members, response, err :=
				vpcService.ReplaceLoadBalancerPoolMembers(options)

			// end-replace_load_balancer_pool_members
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(members).ToNot(BeNil())
			poolMemberID = *members.Members[0].ID
		})
		It(`DeleteLoadBalancerPoolMember request example`, func() {
			// begin-delete_load_balancer_pool_member

			options := &vpcbetav1.DeleteLoadBalancerPoolMemberOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetPoolID(poolID)
			options.SetID(poolMemberID)
			response, err := vpcService.DeleteLoadBalancerPoolMember(options)

			// end-delete_load_balancer_pool_member
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerPoolMember() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerPool request example`, func() {
			// begin-delete_load_balancer_pool

			options := &vpcbetav1.DeleteLoadBalancerPoolOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetID(poolID)
			response, err := vpcService.DeleteLoadBalancerPool(options)

			// end-delete_load_balancer_pool
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerPool() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListenerPolicyRule request example`, func() {
			// begin-delete_load_balancer_listener_policy_rule

			options := &vpcbetav1.DeleteLoadBalancerListenerPolicyRuleOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetPolicyID(policyID)
			options.SetID(policyRuleID)
			response, err :=
				vpcService.DeleteLoadBalancerListenerPolicyRule(options)

			// end-delete_load_balancer_listener_policy_rule
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerListenerPolicyRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListenerPolicy request example`, func() {
			// begin-delete_load_balancer_listener_policy

			options := &vpcbetav1.DeleteLoadBalancerListenerPolicyOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetListenerID(listenerID)
			options.SetID(policyID)
			response, err := vpcService.DeleteLoadBalancerListenerPolicy(options)

			// end-delete_load_balancer_listener_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerListenerPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancerListener request example`, func() {
			// begin-delete_load_balancer_listener

			options := &vpcbetav1.DeleteLoadBalancerListenerOptions{}
			options.SetLoadBalancerID(loadBalancerID)
			options.SetID(listenerID)
			response, err := vpcService.DeleteLoadBalancerListener(options)

			// end-delete_load_balancer_listener
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancerListener() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteLoadBalancer request example`, func() {
			// begin-delete_load_balancer
			deleteVpcOptions := &vpcbetav1.DeleteLoadBalancerOptions{}
			deleteVpcOptions.SetID(loadBalancerID)
			response, err := vpcService.DeleteLoadBalancer(deleteVpcOptions)

			// end-delete_load_balancer
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteLoadBalancer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`ListPrivatePathServiceGateways request example`, func() {
			fmt.Println("\nListPrivatePathServiceGateways() result:")
			// begin-list_private_path_service_gateways
			listPrivatePathServiceGatewaysOptions := &vpcbetav1.ListPrivatePathServiceGatewaysOptions{
				Limit: core.Int64Ptr(int64(10)),
			}

			pager, err := vpcService.NewPrivatePathServiceGatewaysPager(listPrivatePathServiceGatewaysOptions)
			if err != nil {
				panic(err)
			}

			var allResults []vpcbetav1.PrivatePathServiceGateway
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}

			// end-list_private_path_service_gateways
		})
		It(`CreatePrivatePathServiceGateway request example`, func() {
			fmt.Println("\nCreatePrivatePathServiceGateway() result:")
			// begin-create_private_path_service_gateway

			loadBalancerIdentityModel := &vpcbetav1.LoadBalancerIdentityByID{
				ID: &loadBalancerID,
			}

			createPrivatePathServiceGatewayOptions := vpcService.NewCreatePrivatePathServiceGatewayOptions(
				loadBalancerIdentityModel,
				[]string{"my-service.example.com"},
			)

			privatePathServiceGateway, response, err := vpcService.CreatePrivatePathServiceGateway(createPrivatePathServiceGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-create_private_path_service_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(privatePathServiceGateway).ToNot(BeNil())
			createdPPSGID = *privatePathServiceGateway.ID
			createdPPSGCRN = *privatePathServiceGateway.CRN
		})
		It(`GetPrivatePathServiceGateway request example`, func() {
			fmt.Println("\nGetPrivatePathServiceGateway() result:")
			// begin-get_private_path_service_gateway

			getPrivatePathServiceGatewayOptions := vpcService.NewGetPrivatePathServiceGatewayOptions(
				createdPPSGID,
			)

			privatePathServiceGateway, response, err := vpcService.GetPrivatePathServiceGateway(getPrivatePathServiceGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-get_private_path_service_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(privatePathServiceGateway).ToNot(BeNil())
		})
		It(`UpdatePrivatePathServiceGateway request example`, func() {
			fmt.Println("\nUpdatePrivatePathServiceGateway() result:")
			// begin-update_private_path_service_gateway

			privatePathServiceGatewayPatchModel := &vpcbetav1.PrivatePathServiceGatewayPatch{}
			privatePathServiceGatewayPatchModelAsPatch, asPatchErr := privatePathServiceGatewayPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrivatePathServiceGatewayOptions := vpcService.NewUpdatePrivatePathServiceGatewayOptions(
				createdPPSGID,
				privatePathServiceGatewayPatchModelAsPatch,
			)

			privatePathServiceGateway, response, err := vpcService.UpdatePrivatePathServiceGateway(updatePrivatePathServiceGatewayOptions)
			if err != nil {
				panic(err)
			}

			// end-update_private_path_service_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(privatePathServiceGateway).ToNot(BeNil())
		})
		It(`ListPrivatePathServiceGatewayAccountPolicies request example`, func() {
			fmt.Println("\nListPrivatePathServiceGatewayAccountPolicies() result:")
			// begin-list_private_path_service_gateway_account_policies
			listPrivatePathServiceGatewayAccountPoliciesOptions := &vpcbetav1.ListPrivatePathServiceGatewayAccountPoliciesOptions{
				PrivatePathServiceGatewayID: &createdPPSGID,
				Limit:                       core.Int64Ptr(int64(10)),
			}

			pager, err := vpcService.NewPrivatePathServiceGatewayAccountPoliciesPager(listPrivatePathServiceGatewayAccountPoliciesOptions)
			if err != nil {
				panic(err)
			}

			var allResults []vpcbetav1.PrivatePathServiceGatewayAccountPolicy
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}

			// end-list_private_path_service_gateway_account_policies
		})
		It(`CreatePrivatePathServiceGatewayAccountPolicy request example`, func() {
			fmt.Println("\nCreatePrivatePathServiceGatewayAccountPolicy() result:")
			// begin-create_private_path_service_gateway_account_policy

			accountIdentityModel := &vpcbetav1.AccountIdentityByID{
				ID: core.StringPtr("aa2432b1fa4d4ace891e9b80fc104e34"),
			}

			createPrivatePathServiceGatewayAccountPolicyOptions := vpcService.NewCreatePrivatePathServiceGatewayAccountPolicyOptions(
				createdPPSGID,
				"deny",
				accountIdentityModel,
			)

			privatePathServiceGatewayAccountPolicy, response, err := vpcService.CreatePrivatePathServiceGatewayAccountPolicy(createPrivatePathServiceGatewayAccountPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-create_private_path_service_gateway_account_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(privatePathServiceGatewayAccountPolicy).ToNot(BeNil())
			createdPPSGAPID = *privatePathServiceGatewayAccountPolicy.ID
		})
		It(`GetPrivatePathServiceGatewayAccountPolicy request example`, func() {
			fmt.Println("\nGetPrivatePathServiceGatewayAccountPolicy() result:")
			// begin-get_private_path_service_gateway_account_policy

			getPrivatePathServiceGatewayAccountPolicyOptions := vpcService.NewGetPrivatePathServiceGatewayAccountPolicyOptions(
				createdPPSGID,
				createdPPSGAPID,
			)

			privatePathServiceGatewayAccountPolicy, response, err := vpcService.GetPrivatePathServiceGatewayAccountPolicy(getPrivatePathServiceGatewayAccountPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-get_private_path_service_gateway_account_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(privatePathServiceGatewayAccountPolicy).ToNot(BeNil())
		})
		It(`UpdatePrivatePathServiceGatewayAccountPolicy request example`, func() {
			fmt.Println("\nUpdatePrivatePathServiceGatewayAccountPolicy() result:")
			// begin-update_private_path_service_gateway_account_policy

			privatePathServiceGatewayAccountPolicyPatchModel := &vpcbetav1.PrivatePathServiceGatewayAccountPolicyPatch{}
			privatePathServiceGatewayAccountPolicyPatchModelAsPatch, asPatchErr := privatePathServiceGatewayAccountPolicyPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updatePrivatePathServiceGatewayAccountPolicyOptions := vpcService.NewUpdatePrivatePathServiceGatewayAccountPolicyOptions(
				createdPPSGID,
				createdPPSGAPID,
				privatePathServiceGatewayAccountPolicyPatchModelAsPatch,
			)

			privatePathServiceGatewayAccountPolicy, response, err := vpcService.UpdatePrivatePathServiceGatewayAccountPolicy(updatePrivatePathServiceGatewayAccountPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-update_private_path_service_gateway_account_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(privatePathServiceGatewayAccountPolicy).ToNot(BeNil())
		})

		It(`ListEndpointGateways request example`, func() {
			fmt.Println("\nListEndpointGateways() result:")
			// begin-list_endpoint_gateways

			options := vpcService.NewListEndpointGatewaysOptions()
			endpointGateways, response, err :=
				vpcService.ListEndpointGateways(options)

			// end-list_endpoint_gateways
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateways).ToNot(BeNil())

		})
		It(`CreateEndpointGateway request example`, func() {
			fmt.Println("\nCreateEndpointGateway() result:")
			name := getName("egw")
			// begin-create_endpoint_gateway

			options := &vpcbetav1.CreateEndpointGatewayOptions{}
			options.SetName(name)
			options.SetVPC(&vpcbetav1.VPCIdentity{
				ID: &vpcID,
			})

			// targetName := "ibm-ntp-server"
			providerInfrastructureService := "private_path_service_gateway"
			options.SetTarget(
				&vpcbetav1.EndpointGatewayTargetPrototype{
					ResourceType: &providerInfrastructureService,
					CRN:          &createdPPSGCRN,
				},
			)
			endpointGateway, response, err := vpcService.CreateEndpointGateway(options)

			// end-create_endpoint_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(endpointGateway).ToNot(BeNil())
			endpointGatewayID = *endpointGateway.ID
		})
		It(`ListPrivatePathServiceGatewayEndpointGatewayBindings request example`, func() {
			fmt.Println("\nListPrivatePathServiceGatewayEndpointGatewayBindings() result:")
			// begin-list_private_path_service_gateway_endpoint_gateway_bindings
			listPrivatePathServiceGatewayEndpointGatewayBindingsOptions := &vpcbetav1.ListPrivatePathServiceGatewayEndpointGatewayBindingsOptions{
				PrivatePathServiceGatewayID: &createdPPSGID,
				Limit:                       core.Int64Ptr(int64(10)),
				Status:                      core.StringPtr("denied"),
				AccountID:                   core.StringPtr("aa2432b1fa4d4ace891e9b80fc104e34"),
			}

			pager, err := vpcService.NewPrivatePathServiceGatewayEndpointGatewayBindingsPager(listPrivatePathServiceGatewayEndpointGatewayBindingsOptions)
			if err != nil {
				panic(err)
			}

			var allResults []vpcbetav1.PrivatePathServiceGatewayEndpointGatewayBinding
			for pager.HasNext() {
				nextPage, err := pager.GetNext()
				if err != nil {
					panic(err)
				}
				allResults = append(allResults, nextPage...)
			}
			EndpointGatewayBindingID = *allResults[0].ID
			// end-list_private_path_service_gateway_endpoint_gateway_bindings
		})
		It(`GetPrivatePathServiceGatewayEndpointGatewayBinding request example`, func() {
			fmt.Println("\nGetPrivatePathServiceGatewayEndpointGatewayBinding() result:")
			// begin-get_private_path_service_gateway_endpoint_gateway_binding

			getPrivatePathServiceGatewayEndpointGatewayBindingOptions := vpcService.NewGetPrivatePathServiceGatewayEndpointGatewayBindingOptions(
				createdPPSGID,
				EndpointGatewayBindingID,
			)

			privatePathServiceGatewayEndpointGatewayBinding, response, err := vpcService.GetPrivatePathServiceGatewayEndpointGatewayBinding(getPrivatePathServiceGatewayEndpointGatewayBindingOptions)
			if err != nil {
				panic(err)
			}

			// end-get_private_path_service_gateway_endpoint_gateway_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(privatePathServiceGatewayEndpointGatewayBinding).ToNot(BeNil())
		})
		It(`DenyPrivatePathServiceGatewayEndpointGatewayBinding request example`, func() {
			// begin-deny_private_path_service_gateway_endpoint_gateway_binding

			denyPrivatePathServiceGatewayEndpointGatewayBindingOptions := vpcService.NewDenyPrivatePathServiceGatewayEndpointGatewayBindingOptions(
				createdPPSGID,
				EndpointGatewayBindingID,
			)

			response, err := vpcService.DenyPrivatePathServiceGatewayEndpointGatewayBinding(denyPrivatePathServiceGatewayEndpointGatewayBindingOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from DenyPrivatePathServiceGatewayEndpointGatewayBinding(): %d\n", response.StatusCode)
			}

			// end-deny_private_path_service_gateway_endpoint_gateway_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`PermitPrivatePathServiceGatewayEndpointGatewayBinding request example`, func() {
			// begin-permit_private_path_service_gateway_endpoint_gateway_binding

			permitPrivatePathServiceGatewayEndpointGatewayBindingOptions := vpcService.NewPermitPrivatePathServiceGatewayEndpointGatewayBindingOptions(
				createdPPSGID,
				EndpointGatewayBindingID,
			)

			response, err := vpcService.PermitPrivatePathServiceGatewayEndpointGatewayBinding(permitPrivatePathServiceGatewayEndpointGatewayBindingOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from PermitPrivatePathServiceGatewayEndpointGatewayBinding(): %d\n", response.StatusCode)
			}

			// end-permit_private_path_service_gateway_endpoint_gateway_binding

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`RevokeAccountForPrivatePathServiceGateway request example`, func() {
			// begin-revoke_account_for_private_path_service_gateway

			accountIdentityModel := &vpcbetav1.AccountIdentityByID{
				ID: core.StringPtr("aa2432b1fa4d4ace891e9b80fc104e34"),
			}

			revokeAccountForPrivatePathServiceGatewayOptions := vpcService.NewRevokeAccountForPrivatePathServiceGatewayOptions(
				createdPPSGID,
				accountIdentityModel,
			)

			response, err := vpcService.RevokeAccountForPrivatePathServiceGateway(revokeAccountForPrivatePathServiceGatewayOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 200 {
				fmt.Printf("\nUnexpected response status code received from RevokeAccountForPrivatePathServiceGateway(): %d\n", response.StatusCode)
			}

			// end-revoke_account_for_private_path_service_gateway

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
		})
		It(`ListEndpointGatewayIps request example`, func() {
			fmt.Println("\nListEndpointGatewayIps() result:")
			// begin-list_endpoint_gateway_ips

			options := vpcService.NewListEndpointGatewayIpsOptions(endpointGatewayID)
			reservedIPs, response, err := vpcService.ListEndpointGatewayIps(options)

			// end-list_endpoint_gateway_ips
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIPs).ToNot(BeNil())

		})
		It(`AddEndpointGatewayIP request example`, func() {
			fmt.Println("\nAddEndpointGatewayIP() result:")
			// begin-add_endpoint_gateway_ip

			options := vpcService.NewAddEndpointGatewayIPOptions(endpointGatewayID, reservedIPID)
			reservedIP, response, err := vpcService.AddEndpointGatewayIP(options)

			// end-add_endpoint_gateway_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(reservedIP).ToNot(BeNil())
			endpointGatewayTargetID = *reservedIP.ID
		})
		It(`GetEndpointGateway request example`, func() {
			fmt.Println("\nGetEndpointGateway() result:")
			// begin-get_endpoint_gateway_ip

			options := vpcService.NewGetEndpointGatewayOptions(endpointGatewayID)
			endpointGateway, response, err := vpcService.GetEndpointGateway(options)

			// end-get_endpoint_gateway_ip
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})

		It(`GetEndpointGatewayIP request example`, func() {
			fmt.Println("\nGetEndpointGatewayIP() result:")
			// begin-get_endpoint_gateway

			options := vpcService.NewGetEndpointGatewayIPOptions(endpointGatewayID, endpointGatewayTargetID)
			reservedIP, response, err := vpcService.GetEndpointGatewayIP(options)

			// end-get_endpoint_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(reservedIP).ToNot(BeNil())

		})
		It(`UpdateEndpointGateway request example`, func() {
			fmt.Println("\nUpdateEndpointGateway() result:")
			name := getName("egw")
			// begin-update_endpoint_gateway

			endpointGatewayPatchModel := new(vpcbetav1.EndpointGatewayPatch)
			endpointGatewayPatchModel.Name = &name
			endpointGatewayPatchModelAsPatch, asPatchErr := endpointGatewayPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options := &vpcbetav1.UpdateEndpointGatewayOptions{
				ID:                   &endpointGatewayID,
				EndpointGatewayPatch: endpointGatewayPatchModelAsPatch,
			}
			endpointGateway, response, err := vpcService.UpdateEndpointGateway(options)

			// end-update_endpoint_gateway
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(endpointGateway).ToNot(BeNil())

		})

		It(`RemoveEndpointGatewayIP request example`, func() {
			// begin-remove_endpoint_gateway_ip

			removeEndpointGatewayIPOptions := vpcService.NewRemoveEndpointGatewayIPOptions(
				endpointGatewayID,
				endpointGatewayTargetID,
			)

			response, err := vpcService.RemoveEndpointGatewayIP(removeEndpointGatewayIPOptions)

			// end-remove_endpoint_gateway_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nRemoveEndpointGatewayIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteEndpointGateway request example`, func() {
			// begin-delete_endpoint_gateway

			deleteEndpointGatewayOptions := vpcService.NewDeleteEndpointGatewayOptions(
				endpointGatewayID,
			)

			response, err := vpcService.DeleteEndpointGateway(deleteEndpointGatewayOptions)

			// end-delete_endpoint_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteEndpointGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`ListFlowLogCollectors request example`, func() {
			fmt.Println("\nListFlowLogCollectors() result:")
			// begin-list_flow_log_collectors

			options := &vpcbetav1.ListFlowLogCollectorsOptions{}
			flowLogs, response, err := vpcService.ListFlowLogCollectors(options)

			// end-list_flow_log_collectors
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLogs).ToNot(BeNil())

		})
		It(`CreateFlowLogCollector request example`, func() {
			fmt.Println("\nCreateFlowLogCollector() result:")
			name := getName("flowlog")
			// begin-create_flow_log_collector

			options := &vpcbetav1.CreateFlowLogCollectorOptions{}
			options.SetName(name)
			options.SetTarget(&vpcbetav1.FlowLogCollectorTargetPrototypeVPCIdentity{
				ID: &vpcID,
			})
			options.SetStorageBucket(&vpcbetav1.LegacyCloudObjectStorageBucketIdentity{
				Name: &[]string{"bucket-name"}[0],
			})
			flowLog, response, err := vpcService.CreateFlowLogCollector(options)

			// end-create_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(flowLog).ToNot(BeNil())
			flowLogID = *flowLog.ID
		})
		It(`GetFlowLogCollector request example`, func() {
			fmt.Println("\nGetFlowLogCollector() result:")
			// begin-get_flow_log_collector

			options := &vpcbetav1.GetFlowLogCollectorOptions{}
			options.SetID(flowLogID)
			flowLog, response, err := vpcService.GetFlowLogCollector(options)

			// end-get_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLog).ToNot(BeNil())

		})

		It(`UpdateFlowLogCollector request example`, func() {
			fmt.Println("\nUpdateFlowLogCollector() result:")
			name := getName("fl")
			// begin-update_flow_log_collector

			options := &vpcbetav1.UpdateFlowLogCollectorOptions{}
			options.SetID(flowLogID)
			flowLogCollectorPatchModel := &vpcbetav1.FlowLogCollectorPatch{
				Active: &[]bool{true}[0],
				Name:   &name,
			}
			flowLogCollectorPatch, asPatchErr := flowLogCollectorPatchModel.AsPatch()
			if asPatchErr != nil {
				panic(asPatchErr)
			}
			options.FlowLogCollectorPatch = flowLogCollectorPatch
			flowLog, response, err := vpcService.UpdateFlowLogCollector(options)

			// end-update_flow_log_collector
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(flowLog).ToNot(BeNil())

		})

		It(`DeleteFlowLogCollector request example`, func() {
			// begin-delete_flow_log_collector

			deleteFlowLogCollectorOptions := vpcService.NewDeleteFlowLogCollectorOptions(
				flowLogID,
			)

			response, err := vpcService.DeleteFlowLogCollector(deleteFlowLogCollectorOptions)

			// end-delete_flow_log_collector
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteFlowLogCollector() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteBareMetalServerNetworkInterface request example`, func() {
			// begin-delete_bare_metal_server_network_interface

			deleteBareMetalServerNetworkInterfaceOptions := &vpcbetav1.DeleteBareMetalServerNetworkInterfaceOptions{
				BareMetalServerID: &bareMetalServerId,
				ID:                &bareMetalServerNetworkInterfaceId,
			}

			response, err := vpcService.DeleteBareMetalServerNetworkInterface(deleteBareMetalServerNetworkInterfaceOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteBareMetalServerNetworkInterface(): %d\n", response.StatusCode)
			}
			// end-delete_bare_metal_server_network_interface

			fmt.Printf("\nDeleteBareMetalServerNetworkInterface() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteBareMetalServer request example`, func() {
			// begin-delete_bare_metal_server

			deleteBareMetalServerOptions := &vpcbetav1.DeleteBareMetalServerOptions{
				ID: &bareMetalServerId,
			}

			response, err := vpcService.DeleteBareMetalServer(deleteBareMetalServerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteBareMetalServer(): %d\n", response.StatusCode)
			}
			// end-delete_bare_metal_server

			fmt.Printf("\nDeleteBareMetalServer() response status code: %d\n", response.StatusCode)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`RemoveVPNGatewayConnectionPeerCIDR request example`, func() {
			// begin-remove_vpn_gateway_connection_peer_cidr

			options := &vpcbetav1.RemoveVPNGatewayConnectionPeerCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.144.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.RemoveVPNGatewayConnectionPeerCIDR(options)

			// end-remove_vpn_gateway_connection_peer_cidr
			fmt.Printf("\nRemoveVPNGatewayConnectionPeerCIDR() response status code: %d\n", response.StatusCode)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`RemoveVPNGatewayConnectionLocalCIDR request example`, func() {
			// begin-remove_vpn_gateway_connection_local_cidr

			options := &vpcbetav1.RemoveVPNGatewayConnectionLocalCIDROptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			options.SetCIDRPrefix("192.134.0.0")
			options.SetPrefixLength("28")
			response, err := vpcService.RemoveVPNGatewayConnectionLocalCIDR(options)

			// end-remove_vpn_gateway_connection_local_cidr
			fmt.Printf("\nRemoveVPNGatewayConnectionLocalCIDR() response status code: %d\n", response.StatusCode)
			if err != nil {
				panic(err)
			}
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`RemoveInstanceNetworkInterfaceFloatingIP request example`, func() {
			// begin-remove_instance_network_interface_floating_ip

			options := &vpcbetav1.RemoveInstanceNetworkInterfaceFloatingIPOptions{}
			options.SetID(floatingIPID)
			options.SetInstanceID(instanceID)
			options.SetNetworkInterfaceID(eth2ID)
			response, err :=
				vpcService.RemoveInstanceNetworkInterfaceFloatingIP(options)

			// end-remove_instance_network_interface_floating_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nRemoveInstanceNetworkInterfaceFloatingIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroupTargetBinding request example`, func() {
			// begin-delete_security_group_target_binding

			options := &vpcbetav1.DeleteSecurityGroupTargetBindingOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(targetID)
			response, err :=
				vpcService.DeleteSecurityGroupTargetBinding(options)

			// end-delete_security_group_target_binding
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSecurityGroupTargetBinding() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceNetworkInterface request example`, func() {
			// begin-delete_instance_network_interface

			options := &vpcbetav1.DeleteInstanceNetworkInterfaceOptions{}
			options.SetID(eth2ID)
			options.SetInstanceID(instanceID)
			response, err := vpcService.DeleteInstanceNetworkInterface(options)

			// end-delete_instance_network_interface
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceNetworkInterface() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceVolumeAttachment request example`, func() {
			// begin-delete_instance_volume_attachment

			options := &vpcbetav1.DeleteInstanceVolumeAttachmentOptions{}
			options.SetID(volumeAttachmentID)
			options.SetInstanceID(instanceID)
			response, err := vpcService.DeleteInstanceVolumeAttachment(options)

			// end-delete_instance_volume_attachment
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceVolumeAttachment() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVolume request example`, func() {
			// begin-delete_volume

			options := &vpcbetav1.DeleteVolumeOptions{}
			options.SetID(volumeID)
			options.SetIfMatch(ifMatchVolume)
			response, err := vpcService.DeleteVolume(options)

			// end-delete_volume
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVolume() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteFloatingIP request example`, func() {
			// begin-delete_floating_ip

			options := vpcService.NewDeleteFloatingIPOptions(floatingIPID)
			response, err := vpcService.DeleteFloatingIP(options)

			// end-delete_floating_ip
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteFloatingIP() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteInstance request example`, func() {
			// begin-delete_instance

			options := &vpcbetav1.DeleteInstanceOptions{}
			options.SetID(instanceID)
			response, err := vpcService.DeleteInstance(options)
			// end-delete_instance
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstance() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteKey request example`, func() {
			// begin - delete_key

			deleteKeyOptions := &vpcbetav1.DeleteKeyOptions{}
			deleteKeyOptions.SetID(keyID)
			response, err := vpcService.DeleteKey(deleteKeyOptions)

			// end-delete_key

			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteKey() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteImageExportJob request example`, func() {
			// begin-delete_image_export_job

			deleteImageExportJobOptions := &vpcbetav1.DeleteImageExportJobOptions{
				ImageID: &imageID,
				ID:      &imageExportJobID,
			}

			response, err := vpcService.DeleteImageExportJob(deleteImageExportJobOptions)
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteImageExportJob() response status code: %d\n", response.StatusCode)
			// end-delete_image_export_job

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
		})
		It(`DeleteImage request example`, func() {
			// begin-delete_image

			options := &vpcbetav1.DeleteImageOptions{}
			options.SetID(imageID)
			response, err := vpcService.DeleteImage(options)
			// end-delete_image
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteImage() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteVPCRoutingTableRoute request example`, func() {
			// begin-delete_vpc_routing_table_route

			options := &vpcbetav1.DeleteVPCRoutingTableRouteOptions{
				VPCID:          &vpcID,
				RoutingTableID: &routingTableID,
				ID:             &routeID,
			}
			response, err := vpcService.DeleteVPCRoutingTableRoute(options)

			// end-delete_vpc_routing_table_route
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPCRoutingTableRoute() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCRoutingTable request example`, func() {
			// begin-delete_vpc_routing_table

			options := vpcService.NewDeleteVPCRoutingTableOptions(
				vpcID,
				routingTableID,
			)
			response, err := vpcService.DeleteVPCRoutingTable(options)

			// end-delete_vpc_routing_table
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPCRoutingTable() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPCAddressPrefix request example`, func() {
			// begin-delete_vpc_address_prefix

			deleteVpcAddressPrefixOptions := &vpcbetav1.DeleteVPCAddressPrefixOptions{}
			deleteVpcAddressPrefixOptions.SetVPCID(vpcID)
			deleteVpcAddressPrefixOptions.SetID(addressPrefixID)
			response, err :=
				vpcService.DeleteVPCAddressPrefix(deleteVpcAddressPrefixOptions)

			// end-delete_vpc_address_prefix
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPCAddressPrefix() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSnapshot request example`, func() {
			// begin-delete_snapshot
			options := &vpcbetav1.DeleteSnapshotOptions{
				ID:      &snapshotID,
				IfMatch: &ifMatchSnapshot,
			}
			response, err := vpcService.DeleteSnapshot(options)

			// end-delete_snapshot
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSnapshot() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSnapshots request example`, func() {
			// begin-delete_snapshots

			options := &vpcbetav1.DeleteSnapshotsOptions{
				SourceVolumeID: &volumeID,
			}
			response, err := vpcService.DeleteSnapshots(options)

			// end-delete_snapshots
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSnapshots() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteSecurityGroupRule request example`, func() {
			// begin-delete_security_group_rule

			options := &vpcbetav1.DeleteSecurityGroupRuleOptions{}
			options.SetSecurityGroupID(securityGroupID)
			options.SetID(securityGroupRuleID)
			response, err := vpcService.DeleteSecurityGroupRule(options)
			// end-delete_security_group_rule
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSecurityGroupRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSecurityGroup request example`, func() {
			// begin-delete_security_group

			options := &vpcbetav1.DeleteSecurityGroupOptions{}
			options.SetID(securityGroupID)
			response, err := vpcService.DeleteSecurityGroup(options)

			// end-delete_security_group
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSecurityGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeletePublicGateway request example`, func() {
			// begin-delete_public_gateway

			options := &vpcbetav1.DeletePublicGatewayOptions{}
			options.SetID(publicGatewayID)
			response, err := vpcService.DeletePublicGateway(options)

			// end-delete_public_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeletePublicGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteNetworkACLRule request example`, func() {
			// begin-delete_network_acl_rule

			options := &vpcbetav1.DeleteNetworkACLRuleOptions{}
			options.SetID(networkACLRuleID)
			options.SetNetworkACLID(networkACLID)
			response, err := vpcService.DeleteNetworkACLRule(options)

			// end-delete_network_acl_rule
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteNetworkACLRule() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteNetworkACL request example`, func() {
			// begin-delete_network_acl

			options := &vpcbetav1.DeleteNetworkACLOptions{}
			options.SetID(networkACLID)
			response, err := vpcService.DeleteNetworkACL(options)

			// end-delete_network_acl
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteNetworkACL() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})

		It(`DeleteInstanceGroupMembership request example`, func() {
			// begin-delete_instance_group_membership

			options := &vpcbetav1.DeleteInstanceGroupMembershipOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupMembershipID)
			response, err := vpcService.DeleteInstanceGroupMembership(options)

			// end-delete_instance_group_membership
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupMembership() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupMemberships request example`, func() {
			// begin-delete_instance_group_memberships

			options := &vpcbetav1.DeleteInstanceGroupMembershipsOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			response, err := vpcService.DeleteInstanceGroupMemberships(options)

			// end-delete_instance_group_memberships
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupMemberships() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManagerPolicy request example`, func() {
			// begin-delete_instance_group_manager_policy
			options := &vpcbetav1.DeleteInstanceGroupManagerPolicyOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerPolicyID)
			response, err := vpcService.DeleteInstanceGroupManagerPolicy(options)

			// end-delete_instance_group_manager_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupManagerPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManagerAction request example`, func() {
			// begin-delete_instance_group_manager_action

			options := &vpcbetav1.DeleteInstanceGroupManagerActionOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetInstanceGroupManagerID(instanceGroupManagerID)
			options.SetID(instanceGroupManagerActionID)
			response, err := vpcService.DeleteInstanceGroupManagerAction(options)

			// end-delete_instance_group_manager_action
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupManagerAction() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupManager request example`, func() {
			// begin-delete_instance_group_manager

			options := &vpcbetav1.DeleteInstanceGroupManagerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			options.SetID(instanceGroupManagerID)
			response, err := vpcService.DeleteInstanceGroupManager(options)

			// end-delete_instance_group_manager
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupManager() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroupLoadBalancer request example`, func() {
			// begin-delete_instance_group_load_balancer

			options := &vpcbetav1.DeleteInstanceGroupLoadBalancerOptions{}
			options.SetInstanceGroupID(instanceGroupID)
			response, err := vpcService.DeleteInstanceGroupLoadBalancer(options)

			// end-delete_instance_group_load_balancer
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroupLoadBalancer() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceGroup request example`, func() {
			// begin-delete_instance_group

			options := &vpcbetav1.DeleteInstanceGroupOptions{}
			options.SetID(instanceGroupID)
			response, err := vpcService.DeleteInstanceGroup(options)

			// end-delete_instance_group
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteInstanceTemplate request example`, func() {
			// begin-delete_instance_template

			options := &vpcbetav1.DeleteInstanceTemplateOptions{}
			options.SetID(instanceTemplateID)
			response, err := vpcService.DeleteInstanceTemplate(options)

			// end-delete_instance_template
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteInstanceTemplate() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteSubnet request example`, func() {
			// begin-delete_subnet

			options := &vpcbetav1.DeleteSubnetOptions{}
			options.SetID(subnetID)
			response, err := vpcService.DeleteSubnet(options)

			// end-delete_subnet
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteSubnet() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteDedicatedHost request example`, func() {
			// begin-delete_dedicated_host

			options := vpcService.NewDeleteDedicatedHostOptions(dedicatedHostID)
			response, err := vpcService.DeleteDedicatedHost(options)

			// end-delete_dedicated_host
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteDedicatedHost() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteDedicatedHostGroup request example`, func() {
			// begin-delete_dedicated_host_group

			options := vpcService.NewDeleteDedicatedHostGroupOptions(
				dedicatedHostGroupID,
			)
			response, err := vpcService.DeleteDedicatedHostGroup(options)

			// end-delete_dedicated_host_group
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteDedicatedHostGroup() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteIkePolicy request example`, func() {
			// begin-delete_ike_policy

			options := vpcService.NewDeleteIkePolicyOptions(ikePolicyID)
			response, err := vpcService.DeleteIkePolicy(options)

			// end-delete_ike_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteIkePolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteIpsecPolicy request example`, func() {
			// begin-delete_ipsec_policy

			options := vpcService.NewDeleteIpsecPolicyOptions(ipsecPolicyID)
			response, err := vpcService.DeleteIpsecPolicy(options)

			// end-delete_ipsec_policy
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteIpsecPolicy() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPNServerRoute request example`, func() {
			// begin-delete_vpn_server_route

			deleteVPNServerRouteOptions := vpcService.NewDeleteVPNServerRouteOptions(
				vpnServerID,
				vpnServerRouteID,
			)

			response, err := vpcService.DeleteVPNServerRoute(deleteVPNServerRouteOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteVPNServerRoute(): %d\n", response.StatusCode)
			}

			// end-delete_vpn_server_route

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteVPNServerClient request example`, func() {
			Skip("not runnin with mock")
			// begin-delete_vpn_server_client

			deleteVPNServerClientOptions := vpcService.NewDeleteVPNServerClientOptions(
				vpnServerID,
				vpnClientID,
			)

			response, err := vpcService.DeleteVPNServerClient(deleteVPNServerClientOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteVPNServerClient(): %d\n", response.StatusCode)
			}

			// end-delete_vpn_server_client

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})
		It(`DeleteVPNServer request example`, func() {
			// begin-delete_vpn_server

			deleteVPNServerOptions := vpcService.NewDeleteVPNServerOptions(
				vpnServerID,
			)
			deleteVPNServerOptions.SetIfMatch(ifMatchVPNServer)

			response, err := vpcService.DeleteVPNServer(deleteVPNServerOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 202 {
				fmt.Printf("\nUnexpected response status code received from DeleteVPNServer(): %d\n", response.StatusCode)
			}

			// end-delete_vpn_server

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`DeleteVPNGatewayConnection request example`, func() {
			// begin-delete_vpn_gateway_connection

			options := &vpcbetav1.DeleteVPNGatewayConnectionOptions{}
			options.SetVPNGatewayID(vpnGatewayID)
			options.SetID(vpnGatewayConnectionID)
			response, err := vpcService.DeleteVPNGatewayConnection(options)

			// end-delete_vpn_gateway_connection
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPNGatewayConnection() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
		It(`DeleteVPNGateway request example`, func() {
			// begin-delete_vpn_gateway

			options := vpcService.NewDeleteVPNGatewayOptions(vpnGatewayID)
			response, err := vpcService.DeleteVPNGateway(options)

			// end-delete_vpn_gateway
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPNGateway() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))

		})

		It(`DeleteBackupPolicyPlan request example`, func() {
			fmt.Println("\nDeleteBackupPolicyPlan() result:")
			// begin-delete_backup_policy_plan

			deleteBackupPolicyPlanOptions := vpcService.NewDeleteBackupPolicyPlanOptions(
				backupPolicyID,
				backupPolicyPlanID,
			)
			deleteBackupPolicyPlanOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicyPlan, response, err := vpcService.DeleteBackupPolicyPlan(deleteBackupPolicyPlanOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_backup_policy_plan

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(backupPolicyPlan).ToNot(BeNil())

		})
		It(`DeleteBackupPolicy request example`, func() {
			fmt.Println("\nDeleteBackupPolicy() result:")
			// begin-delete_backup_policy

			deleteBackupPolicyOptions := vpcService.NewDeleteBackupPolicyOptions(
				backupPolicyID,
			)
			deleteBackupPolicyOptions.SetIfMatch(ifMatchBackupPolicy)

			backupPolicy, response, err := vpcService.DeleteBackupPolicy(deleteBackupPolicyOptions)
			if err != nil {
				panic(err)
			}

			// end-delete_backup_policy

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(backupPolicy).ToNot(BeNil())

		})
		It(`DeleteVPC request example`, func() {
			// begin-delete_vpc

			deleteVpcOptions := &vpcbetav1.DeleteVPCOptions{}
			deleteVpcOptions.SetID(vpcID)
			response, err := vpcService.DeleteVPC(deleteVpcOptions)
			// end-delete_vpc
			if err != nil {
				panic(err)
			}
			fmt.Printf("\nDeleteVPC() response status code: %d\n", response.StatusCode)

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))

		})
	})

})
