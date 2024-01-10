package communitygalleryimages

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunityGalleryImagesClient struct {
	Client *resourcemanager.Client
}

func NewCommunityGalleryImagesClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunityGalleryImagesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "communitygalleryimages", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunityGalleryImagesClient: %+v", err)
	}

	return &CommunityGalleryImagesClient{
		Client: client,
	}, nil
}
