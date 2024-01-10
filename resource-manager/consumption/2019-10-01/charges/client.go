package charges

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChargesClient struct {
	Client *resourcemanager.Client
}

func NewChargesClientWithBaseURI(sdkApi sdkEnv.Api) (*ChargesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "charges", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChargesClient: %+v", err)
	}

	return &ChargesClient{
		Client: client,
	}, nil
}
