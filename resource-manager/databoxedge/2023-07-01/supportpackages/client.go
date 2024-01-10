package supportpackages

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportPackagesClient struct {
	Client *resourcemanager.Client
}

func NewSupportPackagesClientWithBaseURI(sdkApi sdkEnv.Api) (*SupportPackagesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "supportpackages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SupportPackagesClient: %+v", err)
	}

	return &SupportPackagesClient{
		Client: client,
	}, nil
}
