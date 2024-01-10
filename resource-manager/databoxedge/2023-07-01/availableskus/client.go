package availableskus

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AvailableSkusClient struct {
	Client *resourcemanager.Client
}

func NewAvailableSkusClientWithBaseURI(sdkApi sdkEnv.Api) (*AvailableSkusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "availableskus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AvailableSkusClient: %+v", err)
	}

	return &AvailableSkusClient{
		Client: client,
	}, nil
}
