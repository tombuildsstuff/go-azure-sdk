package quota

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaClient struct {
	Client *resourcemanager.Client
}

func NewQuotaClientWithBaseURI(sdkApi sdkEnv.Api) (*QuotaClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "quota", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating QuotaClient: %+v", err)
	}

	return &QuotaClient{
		Client: client,
	}, nil
}
