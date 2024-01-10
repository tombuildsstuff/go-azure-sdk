package syncagents

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncAgentsClient struct {
	Client *resourcemanager.Client
}

func NewSyncAgentsClientWithBaseURI(sdkApi sdkEnv.Api) (*SyncAgentsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "syncagents", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SyncAgentsClient: %+v", err)
	}

	return &SyncAgentsClient{
		Client: client,
	}, nil
}
