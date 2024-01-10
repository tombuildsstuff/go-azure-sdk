package v2023_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/agentpools"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/maintenanceconfigurations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/managedclusters"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/resolveprivatelinkserviceid"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/snapshots"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2023-10-01/trustedaccess"
)

type Client struct {
	AgentPools                  *agentpools.AgentPoolsClient
	MaintenanceConfigurations   *maintenanceconfigurations.MaintenanceConfigurationsClient
	ManagedClusters             *managedclusters.ManagedClustersClient
	PrivateEndpointConnections  *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources        *privatelinkresources.PrivateLinkResourcesClient
	ResolvePrivateLinkServiceId *resolveprivatelinkserviceid.ResolvePrivateLinkServiceIdClient
	Snapshots                   *snapshots.SnapshotsClient
	TrustedAccess               *trustedaccess.TrustedAccessClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	agentPoolsClient := agentpools.NewAgentPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&agentPoolsClient.Client)

	maintenanceConfigurationsClient := maintenanceconfigurations.NewMaintenanceConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&maintenanceConfigurationsClient.Client)

	managedClustersClient := managedclusters.NewManagedClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&managedClustersClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	resolvePrivateLinkServiceIdClient := resolveprivatelinkserviceid.NewResolvePrivateLinkServiceIdClientWithBaseURI(endpoint)
	configureAuthFunc(&resolvePrivateLinkServiceIdClient.Client)

	snapshotsClient := snapshots.NewSnapshotsClientWithBaseURI(endpoint)
	configureAuthFunc(&snapshotsClient.Client)

	trustedAccessClient := trustedaccess.NewTrustedAccessClientWithBaseURI(endpoint)
	configureAuthFunc(&trustedAccessClient.Client)

	return Client{
		AgentPools:                  &agentPoolsClient,
		MaintenanceConfigurations:   &maintenanceConfigurationsClient,
		ManagedClusters:             &managedClustersClient,
		PrivateEndpointConnections:  &privateEndpointConnectionsClient,
		PrivateLinkResources:        &privateLinkResourcesClient,
		ResolvePrivateLinkServiceId: &resolvePrivateLinkServiceIdClient,
		Snapshots:                   &snapshotsClient,
		TrustedAccess:               &trustedAccessClient,
	}
}
