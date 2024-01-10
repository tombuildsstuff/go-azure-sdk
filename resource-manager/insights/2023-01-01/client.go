package v2023_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2023-01-01/actiongroupsapis"
)

type Client struct {
	ActionGroupsAPIs *actiongroupsapis.ActionGroupsAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	actionGroupsAPIsClient := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&actionGroupsAPIsClient.Client)

	return Client{
		ActionGroupsAPIs: &actionGroupsAPIsClient,
	}
}
