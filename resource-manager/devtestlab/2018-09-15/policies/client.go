package policies

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoliciesClient struct {
	Client *resourcemanager.Client
}

func NewPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*PoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "policies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PoliciesClient: %+v", err)
	}

	return &PoliciesClient{
		Client: client,
	}, nil
}
