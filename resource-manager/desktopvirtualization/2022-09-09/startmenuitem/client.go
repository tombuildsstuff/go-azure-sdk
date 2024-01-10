package startmenuitem

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartMenuItemClient struct {
	Client *resourcemanager.Client
}

func NewStartMenuItemClientWithBaseURI(sdkApi sdkEnv.Api) (*StartMenuItemClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "startmenuitem", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating StartMenuItemClient: %+v", err)
	}

	return &StartMenuItemClient{
		Client: client,
	}, nil
}
