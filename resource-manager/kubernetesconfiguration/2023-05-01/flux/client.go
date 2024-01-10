package flux

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FluxClient struct {
	Client *resourcemanager.Client
}

func NewFluxClientWithBaseURI(sdkApi sdkEnv.Api) (*FluxClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "flux", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating FluxClient: %+v", err)
	}

	return &FluxClient{
		Client: client,
	}, nil
}
