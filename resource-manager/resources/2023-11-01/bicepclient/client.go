package bicepclient

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BicepClientClient struct {
	Client *resourcemanager.Client
}

func NewBicepClientClientWithBaseURI(sdkApi sdkEnv.Api) (*BicepClientClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "bicepclient", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BicepClientClient: %+v", err)
	}

	return &BicepClientClient{
		Client: client,
	}, nil
}
