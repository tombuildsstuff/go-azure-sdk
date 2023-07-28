package featurestoreentitycontainer

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturestoreEntityContainerClient struct {
	Client *resourcemanager.Client
}

func NewFeaturestoreEntityContainerClientWithBaseURI(api environments.Api) (*FeaturestoreEntityContainerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "featurestoreentitycontainer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeaturestoreEntityContainerClient: %+v", err)
	}

	return &FeaturestoreEntityContainerClient{
		Client: client,
	}, nil
}
