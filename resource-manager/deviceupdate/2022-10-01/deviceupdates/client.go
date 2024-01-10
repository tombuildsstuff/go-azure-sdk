package deviceupdates

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceupdatesClient struct {
	Client *resourcemanager.Client
}

func NewDeviceupdatesClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceupdatesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "deviceupdates", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceupdatesClient: %+v", err)
	}

	return &DeviceupdatesClient{
		Client: client,
	}, nil
}
