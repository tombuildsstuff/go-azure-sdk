package changedetection

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDetectionClient struct {
	Client *resourcemanager.Client
}

func NewChangeDetectionClientWithBaseURI(sdkApi sdkEnv.Api) (*ChangeDetectionClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "changedetection", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChangeDetectionClient: %+v", err)
	}

	return &ChangeDetectionClient{
		Client: client,
	}, nil
}
