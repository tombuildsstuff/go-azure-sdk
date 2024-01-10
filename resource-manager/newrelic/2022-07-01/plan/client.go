package plan

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlanClient struct {
	Client *resourcemanager.Client
}

func NewPlanClientWithBaseURI(sdkApi sdkEnv.Api) (*PlanClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "plan", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PlanClient: %+v", err)
	}

	return &PlanClient{
		Client: client,
	}, nil
}
