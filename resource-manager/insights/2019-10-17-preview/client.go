package v2019_10_17_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2019-10-17-preview/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2019-10-17-preview/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2019-10-17-preview/privatelinkscopedresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2019-10-17-preview/privatelinkscopesapis"
)

type Client struct {
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	PrivateLinkScopedResources *privatelinkscopedresources.PrivateLinkScopedResourcesClient
	PrivateLinkScopesAPIs      *privatelinkscopesapis.PrivateLinkScopesAPIsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	privateLinkScopedResourcesClient := privatelinkscopedresources.NewPrivateLinkScopedResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkScopedResourcesClient.Client)

	privateLinkScopesAPIsClient := privatelinkscopesapis.NewPrivateLinkScopesAPIsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkScopesAPIsClient.Client)

	return Client{
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		PrivateLinkScopedResources: &privateLinkScopedResourcesClient,
		PrivateLinkScopesAPIs:      &privateLinkScopesAPIsClient,
	}
}
