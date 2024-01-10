package forecast

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastClient struct {
	Client *resourcemanager.Client
}

func NewForecastClientWithBaseURI(sdkApi sdkEnv.Api) (*ForecastClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "forecast", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ForecastClient: %+v", err)
	}

	return &ForecastClient{
		Client: client,
	}, nil
}
