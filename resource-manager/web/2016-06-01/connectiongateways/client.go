package connectiongateways

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionGatewaysClient struct {
	Client *resourcemanager.Client
}

func NewConnectionGatewaysClientWithBaseURI(sdkApi sdkEnv.Api) (*ConnectionGatewaysClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "connectiongateways", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConnectionGatewaysClient: %+v", err)
	}

	return &ConnectionGatewaysClient{
		Client: client,
	}, nil
}
