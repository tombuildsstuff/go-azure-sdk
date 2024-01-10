package v2023_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/deviceupdates"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/privateendpointconnectionproxies"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/deviceupdate/2023-07-01/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Deviceupdates                    *deviceupdates.DeviceupdatesClient
	PrivateEndpointConnectionProxies *privateendpointconnectionproxies.PrivateEndpointConnectionProxiesClient
	PrivateEndpointConnections       *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources             *privatelinkresources.PrivateLinkResourcesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deviceupdatesClient, err := deviceupdates.NewDeviceupdatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Deviceupdates client: %+v", err)
	}
	configureFunc(deviceupdatesClient.Client)

	privateEndpointConnectionProxiesClient, err := privateendpointconnectionproxies.NewPrivateEndpointConnectionProxiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnectionProxies client: %+v", err)
	}
	configureFunc(privateEndpointConnectionProxiesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	return &Client{
		Deviceupdates:                    deviceupdatesClient,
		PrivateEndpointConnectionProxies: privateEndpointConnectionProxiesClient,
		PrivateEndpointConnections:       privateEndpointConnectionsClient,
		PrivateLinkResources:             privateLinkResourcesClient,
	}, nil
}
