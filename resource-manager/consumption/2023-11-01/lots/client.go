package lots

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotsClient struct {
	Client *resourcemanager.Client
}

func NewLotsClientWithBaseURI(sdkApi sdkEnv.Api) (*LotsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "lots", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LotsClient: %+v", err)
	}

	return &LotsClient{
		Client: client,
	}, nil
}
