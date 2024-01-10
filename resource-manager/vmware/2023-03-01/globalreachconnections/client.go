package globalreachconnections

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GlobalReachConnectionsClient struct {
	Client *resourcemanager.Client
}

func NewGlobalReachConnectionsClientWithBaseURI(sdkApi sdkEnv.Api) (*GlobalReachConnectionsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "globalreachconnections", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GlobalReachConnectionsClient: %+v", err)
	}

	return &GlobalReachConnectionsClient{
		Client: client,
	}, nil
}
