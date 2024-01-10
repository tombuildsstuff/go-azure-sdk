package advisors

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvisorsClient struct {
	Client *resourcemanager.Client
}

func NewAdvisorsClientWithBaseURI(sdkApi sdkEnv.Api) (*AdvisorsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "advisors", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdvisorsClient: %+v", err)
	}

	return &AdvisorsClient{
		Client: client,
	}, nil
}
