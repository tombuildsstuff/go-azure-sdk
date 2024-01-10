package queueserviceproperties

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueueServicePropertiesClient struct {
	Client *resourcemanager.Client
}

func NewQueueServicePropertiesClientWithBaseURI(sdkApi sdkEnv.Api) (*QueueServicePropertiesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "queueserviceproperties", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QueueServicePropertiesClient: %+v", err)
	}

	return &QueueServicePropertiesClient{
		Client: client,
	}, nil
}
