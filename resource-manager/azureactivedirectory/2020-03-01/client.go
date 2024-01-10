package v2020_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privatelinkforazuread"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azureactivedirectory/2020-03-01/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkForAzureAd      *privatelinkforazuread.PrivateLinkForAzureAdClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkForAzureAdClient, err := privatelinkforazuread.NewPrivateLinkForAzureAdClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkForAzureAd client: %+v", err)
	}
	configureFunc(privateLinkForAzureAdClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	return &Client{
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkForAzureAd:      privateLinkForAzureAdClient,
		PrivateLinkResources:       privateLinkResourcesClient,
	}, nil
}
