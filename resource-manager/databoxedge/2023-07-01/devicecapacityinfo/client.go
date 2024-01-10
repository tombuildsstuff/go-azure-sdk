package devicecapacityinfo

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCapacityInfoClient struct {
	Client *resourcemanager.Client
}

func NewDeviceCapacityInfoClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCapacityInfoClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "devicecapacityinfo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCapacityInfoClient: %+v", err)
	}

	return &DeviceCapacityInfoClient{
		Client: client,
	}, nil
}
