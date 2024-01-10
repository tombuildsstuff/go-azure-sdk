package apiversionsets

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiVersionSetsClient struct {
	Client *resourcemanager.Client
}

func NewApiVersionSetsClientWithBaseURI(sdkApi sdkEnv.Api) (*ApiVersionSetsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "apiversionsets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiVersionSetsClient: %+v", err)
	}

	return &ApiVersionSetsClient{
		Client: client,
	}, nil
}
