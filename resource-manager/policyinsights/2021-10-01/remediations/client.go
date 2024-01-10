package remediations

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemediationsClient struct {
	Client *resourcemanager.Client
}

func NewRemediationsClientWithBaseURI(sdkApi sdkEnv.Api) (*RemediationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "remediations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RemediationsClient: %+v", err)
	}

	return &RemediationsClient{
		Client: client,
	}, nil
}
