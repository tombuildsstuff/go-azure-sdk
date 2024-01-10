package tagresource

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TagResourceClient struct {
	Client *resourcemanager.Client
}

func NewTagResourceClientWithBaseURI(sdkApi sdkEnv.Api) (*TagResourceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "tagresource", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TagResourceClient: %+v", err)
	}

	return &TagResourceClient{
		Client: client,
	}, nil
}
