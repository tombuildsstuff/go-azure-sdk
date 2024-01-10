package deletedservice

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServiceClient struct {
	Client *resourcemanager.Client
}

func NewDeletedServiceClientWithBaseURI(sdkApi sdkEnv.Api) (*DeletedServiceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "deletedservice", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedServiceClient: %+v", err)
	}

	return &DeletedServiceClient{
		Client: client,
	}, nil
}
