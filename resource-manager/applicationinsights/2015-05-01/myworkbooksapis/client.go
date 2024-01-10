package myworkbooksapis

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MyworkbooksAPIsClient struct {
	Client *resourcemanager.Client
}

func NewMyworkbooksAPIsClientWithBaseURI(sdkApi sdkEnv.Api) (*MyworkbooksAPIsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "myworkbooksapis", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MyworkbooksAPIsClient: %+v", err)
	}

	return &MyworkbooksAPIsClient{
		Client: client,
	}, nil
}
