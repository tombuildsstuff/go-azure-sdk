package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2021-06-01/policyassignments"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2021-06-01/policydefinitions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/resources/2021-06-01/policysetdefinitions"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	PolicyAssignments    *policyassignments.PolicyAssignmentsClient
	PolicyDefinitions    *policydefinitions.PolicyDefinitionsClient
	PolicySetDefinitions *policysetdefinitions.PolicySetDefinitionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	policyAssignmentsClient, err := policyassignments.NewPolicyAssignmentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyAssignments client: %+v", err)
	}
	configureFunc(policyAssignmentsClient.Client)

	policyDefinitionsClient, err := policydefinitions.NewPolicyDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicyDefinitions client: %+v", err)
	}
	configureFunc(policyDefinitionsClient.Client)

	policySetDefinitionsClient, err := policysetdefinitions.NewPolicySetDefinitionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PolicySetDefinitions client: %+v", err)
	}
	configureFunc(policySetDefinitionsClient.Client)

	return &Client{
		PolicyAssignments:    policyAssignmentsClient,
		PolicyDefinitions:    policyDefinitionsClient,
		PolicySetDefinitions: policySetDefinitionsClient,
	}, nil
}
