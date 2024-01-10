package networksecurityperimeterconfigurations

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkSecurityPerimeterConfigurationsClient struct {
	Client *resourcemanager.Client
}

func NewNetworkSecurityPerimeterConfigurationsClientWithBaseURI(sdkApi sdkEnv.Api) (*NetworkSecurityPerimeterConfigurationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "networksecurityperimeterconfigurations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NetworkSecurityPerimeterConfigurationsClient: %+v", err)
	}

	return &NetworkSecurityPerimeterConfigurationsClient{
		Client: client,
	}, nil
}
