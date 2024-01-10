package dataversionregistry

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataVersionRegistryClient struct {
	Client *resourcemanager.Client
}

func NewDataVersionRegistryClientWithBaseURI(sdkApi sdkEnv.Api) (*DataVersionRegistryClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dataversionregistry", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataVersionRegistryClient: %+v", err)
	}

	return &DataVersionRegistryClient{
		Client: client,
	}, nil
}
