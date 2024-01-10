package v2022_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/networkfunction/2022-11-01/azuretrafficcollectors"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/networkfunction/2022-11-01/collectorpolicies"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureTrafficCollectors *azuretrafficcollectors.AzureTrafficCollectorsClient
	CollectorPolicies      *collectorpolicies.CollectorPoliciesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	azureTrafficCollectorsClient, err := azuretrafficcollectors.NewAzureTrafficCollectorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureTrafficCollectors client: %+v", err)
	}
	configureFunc(azureTrafficCollectorsClient.Client)

	collectorPoliciesClient, err := collectorpolicies.NewCollectorPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CollectorPolicies client: %+v", err)
	}
	configureFunc(collectorPoliciesClient.Client)

	return &Client{
		AzureTrafficCollectors: azureTrafficCollectorsClient,
		CollectorPolicies:      collectorPoliciesClient,
	}, nil
}
