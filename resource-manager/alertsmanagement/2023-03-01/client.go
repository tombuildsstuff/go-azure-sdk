package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/alertsmanagement/2023-03-01/prometheusrulegroups"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	PrometheusRuleGroups *prometheusrulegroups.PrometheusRuleGroupsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	prometheusRuleGroupsClient, err := prometheusrulegroups.NewPrometheusRuleGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrometheusRuleGroups client: %+v", err)
	}
	configureFunc(prometheusRuleGroupsClient.Client)

	return &Client{
		PrometheusRuleGroups: prometheusRuleGroupsClient,
	}, nil
}
