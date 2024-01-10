package communitygalleries

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleriesClient struct {
	Client *resourcemanager.Client
}

func NewCommunityGalleriesClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunityGalleriesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "communitygalleries", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunityGalleriesClient: %+v", err)
	}

	return &CommunityGalleriesClient{
		Client: client,
	}, nil
}
