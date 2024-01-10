package virtualmachinesizes

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineSizesClient struct {
	Client *resourcemanager.Client
}

func NewVirtualMachineSizesClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualMachineSizesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "virtualmachinesizes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualMachineSizesClient: %+v", err)
	}

	return &VirtualMachineSizesClient{
		Client: client,
	}, nil
}
