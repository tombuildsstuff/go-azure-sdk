package waitstatistics

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WaitStatisticsClient struct {
	Client *resourcemanager.Client
}

func NewWaitStatisticsClientWithBaseURI(sdkApi sdkEnv.Api) (*WaitStatisticsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "waitstatistics", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WaitStatisticsClient: %+v", err)
	}

	return &WaitStatisticsClient{
		Client: client,
	}, nil
}
