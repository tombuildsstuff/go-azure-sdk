package v2021_08_08

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/alertsmanagement/2021-08-08/alertprocessingrules"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	AlertProcessingRules *alertprocessingrules.AlertProcessingRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertProcessingRulesClient, err := alertprocessingrules.NewAlertProcessingRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AlertProcessingRules client: %+v", err)
	}
	configureFunc(alertProcessingRulesClient.Client)

	return &Client{
		AlertProcessingRules: alertProcessingRulesClient,
	}, nil
}
