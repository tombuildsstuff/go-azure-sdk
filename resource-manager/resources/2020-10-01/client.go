package v2020_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2020-10-01/deployments"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2020-10-01/deploymentscripts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2020-10-01/providers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2020-10-01/resourcegroups"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2020-10-01/resources"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeploymentScripts *deploymentscripts.DeploymentScriptsClient
	Deployments       *deployments.DeploymentsClient
	Providers         *providers.ProvidersClient
	ResourceGroups    *resourcegroups.ResourceGroupsClient
	Resources         *resources.ResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deploymentScriptsClient, err := deploymentscripts.NewDeploymentScriptsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeploymentScripts client: %+v", err)
	}
	configureFunc(deploymentScriptsClient.Client)

	deploymentsClient, err := deployments.NewDeploymentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deployments client: %+v", err)
	}
	configureFunc(deploymentsClient.Client)

	providersClient, err := providers.NewProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Providers client: %+v", err)
	}
	configureFunc(providersClient.Client)

	resourceGroupsClient, err := resourcegroups.NewResourceGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourceGroups client: %+v", err)
	}
	configureFunc(resourceGroupsClient.Client)

	resourcesClient, err := resources.NewResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Resources client: %+v", err)
	}
	configureFunc(resourcesClient.Client)

	return &Client{
		DeploymentScripts: deploymentScriptsClient,
		Deployments:       deploymentsClient,
		Providers:         providersClient,
		ResourceGroups:    resourceGroupsClient,
		Resources:         resourcesClient,
	}, nil
}
