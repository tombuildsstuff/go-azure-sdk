package namespacesprivatelinkresources

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespacesPrivateLinkResourcesClient struct {
	Client *resourcemanager.Client
}

func NewNamespacesPrivateLinkResourcesClientWithBaseURI(sdkApi sdkEnv.Api) (*NamespacesPrivateLinkResourcesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "namespacesprivatelinkresources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating NamespacesPrivateLinkResourcesClient: %+v", err)
	}

	return &NamespacesPrivateLinkResourcesClient{
		Client: client,
	}, nil
}
