package operationstatus

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatusClient struct {
	Client *resourcemanager.Client
}

func NewOperationStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*OperationStatusClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "operationstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationStatusClient: %+v", err)
	}

	return &OperationStatusClient{
		Client: client,
	}, nil
}
