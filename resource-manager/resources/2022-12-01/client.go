package v2022_12_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2022-12-01/subscriptions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2022-12-01/tenants"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Subscriptions *subscriptions.SubscriptionsClient
	Tenants       *tenants.TenantsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	subscriptionsClient, err := subscriptions.NewSubscriptionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Subscriptions client: %+v", err)
	}
	configureFunc(subscriptionsClient.Client)

	tenantsClient, err := tenants.NewTenantsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tenants client: %+v", err)
	}
	configureFunc(tenantsClient.Client)

	return &Client{
		Subscriptions: subscriptionsClient,
		Tenants:       tenantsClient,
	}, nil
}
