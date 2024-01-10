package integrationaccountsessions

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationAccountSessionsClient struct {
	Client *resourcemanager.Client
}

func NewIntegrationAccountSessionsClientWithBaseURI(sdkApi sdkEnv.Api) (*IntegrationAccountSessionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "integrationaccountsessions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationAccountSessionsClient: %+v", err)
	}

	return &IntegrationAccountSessionsClient{
		Client: client,
	}, nil
}
