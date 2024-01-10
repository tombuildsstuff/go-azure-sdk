package azureadadministrators

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADAdministratorsClient struct {
	Client *resourcemanager.Client
}

func NewAzureADAdministratorsClientWithBaseURI(sdkApi sdkEnv.Api) (*AzureADAdministratorsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "azureadadministrators", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AzureADAdministratorsClient: %+v", err)
	}

	return &AzureADAdministratorsClient{
		Client: client,
	}, nil
}
