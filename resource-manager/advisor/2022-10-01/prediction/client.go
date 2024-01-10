package prediction

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictionClient struct {
	Client *resourcemanager.Client
}

func NewPredictionClientWithBaseURI(sdkApi sdkEnv.Api) (*PredictionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "prediction", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PredictionClient: %+v", err)
	}

	return &PredictionClient{
		Client: client,
	}, nil
}
