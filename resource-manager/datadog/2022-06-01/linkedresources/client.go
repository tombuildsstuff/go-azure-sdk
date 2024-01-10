package linkedresources

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LinkedResourcesClient struct {
	Client *resourcemanager.Client
}

func NewLinkedResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*LinkedResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "linkedresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LinkedResourcesClient: %+v", err)
	}

	return &LinkedResourcesClient{
		Client: client,
	}, nil
}
