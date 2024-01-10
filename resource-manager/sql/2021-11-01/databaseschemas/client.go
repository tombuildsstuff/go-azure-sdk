package databaseschemas

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseSchemasClient struct {
	Client *resourcemanager.Client
}

func NewDatabaseSchemasClientWithBaseURI(sdkApi sdkEnv.Api) (*DatabaseSchemasClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "databaseschemas", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatabaseSchemasClient: %+v", err)
	}

	return &DatabaseSchemasClient{
		Client: client,
	}, nil
}
