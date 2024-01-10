package managedinstancedtcs

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceDtcsClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstanceDtcsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedInstanceDtcsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managedinstancedtcs", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstanceDtcsClient: %+v", err)
	}

	return &ManagedInstanceDtcsClient{
		Client: client,
	}, nil
}
