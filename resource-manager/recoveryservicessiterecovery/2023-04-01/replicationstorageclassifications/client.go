package replicationstorageclassifications

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicationStorageClassificationsClient struct {
	Client *resourcemanager.Client
}

func NewReplicationStorageClassificationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReplicationStorageClassificationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "replicationstorageclassifications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReplicationStorageClassificationsClient: %+v", err)
	}

	return &ReplicationStorageClassificationsClient{
		Client: client,
	}, nil
}
