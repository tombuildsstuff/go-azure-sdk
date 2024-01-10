package heatmaps

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HeatMapsClient struct {
	Client *resourcemanager.Client
}

func NewHeatMapsClientWithBaseURI(sdkApi sdkEnv.Api) (*HeatMapsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "heatmaps", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HeatMapsClient: %+v", err)
	}

	return &HeatMapsClient{
		Client: client,
	}, nil
}
