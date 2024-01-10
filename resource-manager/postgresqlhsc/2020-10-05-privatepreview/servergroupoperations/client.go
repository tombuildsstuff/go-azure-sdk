package servergroupoperations

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerGroupOperationsClient struct {
	Client *resourcemanager.Client
}

func NewServerGroupOperationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ServerGroupOperationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "servergroupoperations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ServerGroupOperationsClient: %+v", err)
	}

	return &ServerGroupOperationsClient{
		Client: client,
	}, nil
}
