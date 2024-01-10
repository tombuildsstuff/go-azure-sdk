package privatelink

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkClient struct {
	Client *resourcemanager.Client
}

func NewPrivateLinkClientWithBaseURI(sdkApi sdkEnv.Api) (*PrivateLinkClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "privatelink", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PrivateLinkClient: %+v", err)
	}

	return &PrivateLinkClient{
		Client: client,
	}, nil
}
