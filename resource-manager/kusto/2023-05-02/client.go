package v2023_05_02

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/attacheddatabaseconfigurations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/clusterprincipalassignments"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/clusters"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/databaseprincipalassignments"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/databases"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/dataconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/kusto"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/managedprivateendpoints"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/outboundnetworkdependenciesendpoints"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/scripts"
)

type Client struct {
	AttachedDatabaseConfigurations       *attacheddatabaseconfigurations.AttachedDatabaseConfigurationsClient
	ClusterPrincipalAssignments          *clusterprincipalassignments.ClusterPrincipalAssignmentsClient
	Clusters                             *clusters.ClustersClient
	DataConnections                      *dataconnections.DataConnectionsClient
	DatabasePrincipalAssignments         *databaseprincipalassignments.DatabasePrincipalAssignmentsClient
	Databases                            *databases.DatabasesClient
	Kusto                                *kusto.KustoClient
	ManagedPrivateEndpoints              *managedprivateendpoints.ManagedPrivateEndpointsClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	Scripts                              *scripts.ScriptsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	attachedDatabaseConfigurationsClient := attacheddatabaseconfigurations.NewAttachedDatabaseConfigurationsClientWithBaseURI(endpoint)
	configureAuthFunc(&attachedDatabaseConfigurationsClient.Client)

	clusterPrincipalAssignmentsClient := clusterprincipalassignments.NewClusterPrincipalAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterPrincipalAssignmentsClient.Client)

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	dataConnectionsClient := dataconnections.NewDataConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&dataConnectionsClient.Client)

	databasePrincipalAssignmentsClient := databaseprincipalassignments.NewDatabasePrincipalAssignmentsClientWithBaseURI(endpoint)
	configureAuthFunc(&databasePrincipalAssignmentsClient.Client)

	databasesClient := databases.NewDatabasesClientWithBaseURI(endpoint)
	configureAuthFunc(&databasesClient.Client)

	kustoClient := kusto.NewKustoClientWithBaseURI(endpoint)
	configureAuthFunc(&kustoClient.Client)

	managedPrivateEndpointsClient := managedprivateendpoints.NewManagedPrivateEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedPrivateEndpointsClient.Client)

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	scriptsClient := scripts.NewScriptsClientWithBaseURI(endpoint)
	configureAuthFunc(&scriptsClient.Client)

	return Client{
		AttachedDatabaseConfigurations:       &attachedDatabaseConfigurationsClient,
		ClusterPrincipalAssignments:          &clusterPrincipalAssignmentsClient,
		Clusters:                             &clustersClient,
		DataConnections:                      &dataConnectionsClient,
		DatabasePrincipalAssignments:         &databasePrincipalAssignmentsClient,
		Databases:                            &databasesClient,
		Kusto:                                &kustoClient,
		ManagedPrivateEndpoints:              &managedPrivateEndpointsClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		PrivateLinkResources:                 &privateLinkResourcesClient,
		Scripts:                              &scriptsClient,
	}
}
