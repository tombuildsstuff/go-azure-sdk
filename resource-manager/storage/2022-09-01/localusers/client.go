package localusers

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocalUsersClient struct {
	Client *resourcemanager.Client
}

func NewLocalUsersClientWithBaseURI(sdkApi sdkEnv.Api) (*LocalUsersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "localusers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocalUsersClient: %+v", err)
	}

	return &LocalUsersClient{
		Client: client,
	}, nil
}
