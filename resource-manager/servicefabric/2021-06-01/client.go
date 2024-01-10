package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/application"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/applicationtype"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/applicationtypeversion"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/cluster"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/clusterversion"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/listupgradableversions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/servicefabric/2021-06-01/service"
)

type Client struct {
	Application            *application.ApplicationClient
	ApplicationType        *applicationtype.ApplicationTypeClient
	ApplicationTypeVersion *applicationtypeversion.ApplicationTypeVersionClient
	Cluster                *cluster.ClusterClient
	ClusterVersion         *clusterversion.ClusterVersionClient
	ListUpgradableVersions *listupgradableversions.ListUpgradableVersionsClient
	Service                *service.ServiceClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	applicationClient := application.NewApplicationClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationClient.Client)

	applicationTypeClient := applicationtype.NewApplicationTypeClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationTypeClient.Client)

	applicationTypeVersionClient := applicationtypeversion.NewApplicationTypeVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&applicationTypeVersionClient.Client)

	clusterClient := cluster.NewClusterClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterClient.Client)

	clusterVersionClient := clusterversion.NewClusterVersionClientWithBaseURI(endpoint)
	configureAuthFunc(&clusterVersionClient.Client)

	listUpgradableVersionsClient := listupgradableversions.NewListUpgradableVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&listUpgradableVersionsClient.Client)

	serviceClient := service.NewServiceClientWithBaseURI(endpoint)
	configureAuthFunc(&serviceClient.Client)

	return Client{
		Application:            &applicationClient,
		ApplicationType:        &applicationTypeClient,
		ApplicationTypeVersion: &applicationTypeVersionClient,
		Cluster:                &clusterClient,
		ClusterVersion:         &clusterVersionClient,
		ListUpgradableVersions: &listUpgradableVersionsClient,
		Service:                &serviceClient,
	}
}
