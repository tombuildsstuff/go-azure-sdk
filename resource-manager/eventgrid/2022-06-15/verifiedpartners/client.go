package verifiedpartners

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VerifiedPartnersClient struct {
	Client *resourcemanager.Client
}

func NewVerifiedPartnersClientWithBaseURI(sdkApi sdkEnv.Api) (*VerifiedPartnersClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "verifiedpartners", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VerifiedPartnersClient: %+v", err)
	}

	return &VerifiedPartnersClient{
		Client: client,
	}, nil
}
