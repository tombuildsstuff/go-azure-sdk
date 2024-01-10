package datacontainer

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataContainerClient struct {
	Client *resourcemanager.Client
}

func NewDataContainerClientWithBaseURI(sdkApi sdkEnv.Api) (*DataContainerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "datacontainer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataContainerClient: %+v", err)
	}

	return &DataContainerClient{
		Client: client,
	}, nil
}
