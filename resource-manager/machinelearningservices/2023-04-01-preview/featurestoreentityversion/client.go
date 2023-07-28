package featurestoreentityversion

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeaturestoreEntityVersionClient struct {
	Client *resourcemanager.Client
}

func NewFeaturestoreEntityVersionClientWithBaseURI(api environments.Api) (*FeaturestoreEntityVersionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(api, "featurestoreentityversion", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FeaturestoreEntityVersionClient: %+v", err)
	}

	return &FeaturestoreEntityVersionClient{
		Client: client,
	}, nil
}
