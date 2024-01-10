package blobinventorypolicies

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobInventoryPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewBlobInventoryPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*BlobInventoryPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "blobinventorypolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BlobInventoryPoliciesClient: %+v", err)
	}

	return &BlobInventoryPoliciesClient{
		Client: client,
	}, nil
}
