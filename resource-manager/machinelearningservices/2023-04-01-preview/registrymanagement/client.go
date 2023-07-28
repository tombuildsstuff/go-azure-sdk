package registrymanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegistryManagementClient struct {
	Client *resourcemanager.Client
}

func NewRegistryManagementClientWithBaseURI(api environments.Api) (*RegistryManagementClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "registrymanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegistryManagementClient: %+v", err)
	}

	return &RegistryManagementClient{
		Client: client,
	}, nil
}
