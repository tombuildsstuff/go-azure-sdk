package delete

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DELETEClient struct {
	Client *resourcemanager.Client
}

func NewDELETEClientWithBaseURI(sdkApi sdkEnv.Api) (*DELETEClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "delete", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DELETEClient: %+v", err)
	}

	return &DELETEClient{
		Client: client,
	}, nil
}
