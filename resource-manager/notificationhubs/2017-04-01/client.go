package v2017_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/notificationhubs/2017-04-01/namespaces"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/notificationhubs/2017-04-01/notificationhubs"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Namespaces       *namespaces.NamespacesClient
	NotificationHubs *notificationhubs.NotificationHubsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	namespacesClient, err := namespaces.NewNamespacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Namespaces client: %+v", err)
	}
	configureFunc(namespacesClient.Client)

	notificationHubsClient, err := notificationhubs.NewNotificationHubsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotificationHubs client: %+v", err)
	}
	configureFunc(notificationHubsClient.Client)

	return &Client{
		Namespaces:       namespacesClient,
		NotificationHubs: notificationHubsClient,
	}, nil
}
