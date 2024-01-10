package v2023_04_27

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/dynatrace/2023-04-27/monitors"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/dynatrace/2023-04-27/singlesignon"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/dynatrace/2023-04-27/tagrules"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Monitors     *monitors.MonitorsClient
	SingleSignOn *singlesignon.SingleSignOnClient
	TagRules     *tagrules.TagRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	monitorsClient, err := monitors.NewMonitorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Monitors client: %+v", err)
	}
	configureFunc(monitorsClient.Client)

	singleSignOnClient, err := singlesignon.NewSingleSignOnClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SingleSignOn client: %+v", err)
	}
	configureFunc(singleSignOnClient.Client)

	tagRulesClient, err := tagrules.NewTagRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TagRules client: %+v", err)
	}
	configureFunc(tagRulesClient.Client)

	return &Client{
		Monitors:     monitorsClient,
		SingleSignOn: singleSignOnClient,
		TagRules:     tagRulesClient,
	}, nil
}
