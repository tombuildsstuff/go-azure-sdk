package apisbytag

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApisByTagClient struct {
	Client *resourcemanager.Client
}

func NewApisByTagClientWithBaseURI(sdkApi sdkEnv.Api) (*ApisByTagClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "apisbytag", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApisByTagClient: %+v", err)
	}

	return &ApisByTagClient{
		Client: client,
	}, nil
}
