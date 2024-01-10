package managedinstanceprivateendpointconnections

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstancePrivateEndpointConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewManagedInstancePrivateEndpointConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*ManagedInstancePrivateEndpointConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "managedinstanceprivateendpointconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ManagedInstancePrivateEndpointConnectionsClient: %+v", err)
	}

	return &ManagedInstancePrivateEndpointConnectionsClient{
		Client: client,
	}, nil
}
