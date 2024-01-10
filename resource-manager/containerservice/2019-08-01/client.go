package v2019_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2019-08-01/agentpools"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2019-08-01/containerservices"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerservice/2019-08-01/managedclusters"
)

type Client struct {
	AgentPools        *agentpools.AgentPoolsClient
	ContainerServices *containerservices.ContainerServicesClient
	ManagedClusters   *managedclusters.ManagedClustersClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	agentPoolsClient := agentpools.NewAgentPoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&agentPoolsClient.Client)

	containerServicesClient := containerservices.NewContainerServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&containerServicesClient.Client)

	managedClustersClient := managedclusters.NewManagedClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&managedClustersClient.Client)

	return Client{
		AgentPools:        &agentPoolsClient,
		ContainerServices: &containerServicesClient,
		ManagedClusters:   &managedClustersClient,
	}
}
