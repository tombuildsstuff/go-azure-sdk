package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/healthbot/2023-05-01/healthbots"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Healthbots *healthbots.HealthbotsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	healthbotsClient, err := healthbots.NewHealthbotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Healthbots client: %+v", err)
	}
	configureFunc(healthbotsClient.Client)

	return &Client{
		Healthbots: healthbotsClient,
	}, nil
}
