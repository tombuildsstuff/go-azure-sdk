package hcrpreports

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HCRPReportsClient struct {
	Client *resourcemanager.Client
}

func NewHCRPReportsClientWithBaseURI(sdkApi sdkEnv.Api) (*HCRPReportsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "hcrpreports", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HCRPReportsClient: %+v", err)
	}

	return &HCRPReportsClient{
		Client: client,
	}, nil
}
