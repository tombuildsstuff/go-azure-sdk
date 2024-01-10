package changedatacapture

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureClient struct {
	Client *resourcemanager.Client
}

func NewChangeDataCaptureClientWithBaseURI(sdkApi sdkEnv.Api) (*ChangeDataCaptureClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "changedatacapture", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChangeDataCaptureClient: %+v", err)
	}

	return &ChangeDataCaptureClient{
		Client: client,
	}, nil
}
