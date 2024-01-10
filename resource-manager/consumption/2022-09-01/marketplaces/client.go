package marketplaces

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplacesClient struct {
	Client *resourcemanager.Client
}

func NewMarketplacesClientWithBaseURI(sdkApi sdkEnv.Api) (*MarketplacesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "marketplaces", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MarketplacesClient: %+v", err)
	}

	return &MarketplacesClient{
		Client: client,
	}, nil
}
