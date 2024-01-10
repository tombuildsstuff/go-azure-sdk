package apirevision

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiRevisionClient struct {
	Client *resourcemanager.Client
}

func NewApiRevisionClientWithBaseURI(sdkApi sdkEnv.Api) (*ApiRevisionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "apirevision", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiRevisionClient: %+v", err)
	}

	return &ApiRevisionClient{
		Client: client,
	}, nil
}
