package delegationsettings

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegationSettingsClient struct {
	Client *resourcemanager.Client
}

func NewDelegationSettingsClientWithBaseURI(sdkApi sdkEnv.Api) (*DelegationSettingsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "delegationsettings", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DelegationSettingsClient: %+v", err)
	}

	return &DelegationSettingsClient{
		Client: client,
	}, nil
}
