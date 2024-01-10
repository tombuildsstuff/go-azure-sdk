package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/databricks/2023-05-01/accessconnector"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccessConnector *accessconnector.AccessConnectorClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accessConnectorClient, err := accessconnector.NewAccessConnectorClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccessConnector client: %+v", err)
	}
	configureFunc(accessConnectorClient.Client)

	return &Client{
		AccessConnector: accessConnectorClient,
	}, nil
}
