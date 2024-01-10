package notificationrecipientuser

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NotificationRecipientUserClient struct {
	Client *resourcemanager.Client
}

func NewNotificationRecipientUserClientWithBaseURI(sdkApi sdkEnv.Api) (*NotificationRecipientUserClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "notificationrecipientuser", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NotificationRecipientUserClient: %+v", err)
	}

	return &NotificationRecipientUserClient{
		Client: client,
	}, nil
}
