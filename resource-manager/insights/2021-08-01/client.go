package v2021_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2021-08-01/scheduledqueryrules"
)

type Client struct {
	ScheduledQueryRules *scheduledqueryrules.ScheduledQueryRulesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	scheduledQueryRulesClient := scheduledqueryrules.NewScheduledQueryRulesClientWithBaseURI(endpoint)
	configureAuthFunc(&scheduledQueryRulesClient.Client)

	return Client{
		ScheduledQueryRules: &scheduledQueryRulesClient,
	}
}
