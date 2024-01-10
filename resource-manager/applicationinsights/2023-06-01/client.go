package v2023_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/applicationinsights/2023-06-01/workbooksapis"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	WorkbooksAPIs *workbooksapis.WorkbooksAPIsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	workbooksAPIsClient, err := workbooksapis.NewWorkbooksAPIsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building WorkbooksAPIs client: %+v", err)
	}
	configureFunc(workbooksAPIsClient.Client)

	return &Client{
		WorkbooksAPIs: workbooksAPIsClient,
	}, nil
}
