package integrationserviceenvironmentskus

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationServiceEnvironmentSkusClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationServiceEnvironmentSkusClientWithBaseURI(sdkApi sdkEnv.Api) (*IntegrationServiceEnvironmentSkusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "integrationserviceenvironmentskus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationServiceEnvironmentSkusClient: %+v", err)
	}

	return &IntegrationServiceEnvironmentSkusClient{
		Client: client,
	}, nil
}
