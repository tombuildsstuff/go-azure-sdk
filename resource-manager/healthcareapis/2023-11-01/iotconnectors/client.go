package iotconnectors

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotConnectorsClient struct {
	Client *resourcemanager.Client
}

func NewIotConnectorsClientWithBaseURI(sdkApi sdkEnv.Api) (*IotConnectorsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "iotconnectors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IotConnectorsClient: %+v", err)
	}

	return &IotConnectorsClient{
		Client: client,
	}, nil
}
