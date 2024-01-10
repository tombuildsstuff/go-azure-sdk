package datasources

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSourcesClient struct {
	Client *resourcemanager.Client
}

func NewDataSourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*DataSourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "datasources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSourcesClient: %+v", err)
	}

	return &DataSourcesClient{
		Client: client,
	}, nil
}
