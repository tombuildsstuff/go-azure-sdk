package api

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiClient struct {
	Client *resourcemanager.Client
}

func NewApiClientWithBaseURI(sdkApi sdkEnv.Api) (*ApiClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "api", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiClient: %+v", err)
	}

	return &ApiClient{
		Client: client,
	}, nil
}
