package oucontainer

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OuContainerClient struct {
	Client *resourcemanager.Client
}

func NewOuContainerClientWithBaseURI(sdkApi sdkEnv.Api) (*OuContainerClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "oucontainer", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OuContainerClient: %+v", err)
	}

	return &OuContainerClient{
		Client: client,
	}, nil
}
