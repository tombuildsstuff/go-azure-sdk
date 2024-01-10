package deletedservers

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedServersClient struct {
	Client *resourcemanager.Client
}

func NewDeletedServersClientWithBaseURI(sdkApi sdkEnv.Api) (*DeletedServersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "deletedservers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedServersClient: %+v", err)
	}

	return &DeletedServersClient{
		Client: client,
	}, nil
}
