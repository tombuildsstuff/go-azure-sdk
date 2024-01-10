package packetcoredataplane

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PacketCoreDataPlaneClient struct {
	Client *resourcemanager.Client
}

func NewPacketCoreDataPlaneClientWithBaseURI(sdkApi sdkEnv.Api) (*PacketCoreDataPlaneClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "packetcoredataplane", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PacketCoreDataPlaneClient: %+v", err)
	}

	return &PacketCoreDataPlaneClient{
		Client: client,
	}, nil
}
