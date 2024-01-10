package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/frontdoor/2022-05-01/webapplicationfirewallmanagedrulesets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/frontdoor/2022-05-01/webapplicationfirewallpolicies"
)

type Client struct {
	WebApplicationFirewallManagedRuleSets *webapplicationfirewallmanagedrulesets.WebApplicationFirewallManagedRuleSetsClient
	WebApplicationFirewallPolicies        *webapplicationfirewallpolicies.WebApplicationFirewallPoliciesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	webApplicationFirewallManagedRuleSetsClient := webapplicationfirewallmanagedrulesets.NewWebApplicationFirewallManagedRuleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&webApplicationFirewallManagedRuleSetsClient.Client)

	webApplicationFirewallPoliciesClient := webapplicationfirewallpolicies.NewWebApplicationFirewallPoliciesClientWithBaseURI(endpoint)
	configureAuthFunc(&webApplicationFirewallPoliciesClient.Client)

	return Client{
		WebApplicationFirewallManagedRuleSets: &webApplicationFirewallManagedRuleSetsClient,
		WebApplicationFirewallPolicies:        &webApplicationFirewallPoliciesClient,
	}
}
