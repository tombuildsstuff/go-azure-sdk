package devices

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesClient struct {
	Client *resourcemanager.Client
}

func NewDevicesClientWithBaseURI(sdkApi sdkEnv.Api) (*DevicesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "devices", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DevicesClient: %+v", err)
	}

	return &DevicesClient{
		Client: client,
	}, nil
}
