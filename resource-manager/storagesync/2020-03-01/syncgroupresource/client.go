package syncgroupresource

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupResourceClient struct {
	Client *resourcemanager.Client
}

func NewSyncGroupResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*SyncGroupResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "syncgroupresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SyncGroupResourceClient: %+v", err)
	}

	return &SyncGroupResourceClient{
		Client: client,
	}, nil
}
