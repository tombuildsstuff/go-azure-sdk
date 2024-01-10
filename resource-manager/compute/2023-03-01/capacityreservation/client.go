package capacityreservation

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CapacityReservationClient struct {
	Client *resourcemanager.Client
}

func NewCapacityReservationClientWithBaseURI(sdkApi sdkEnv.Api) (*CapacityReservationClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "capacityreservation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CapacityReservationClient: %+v", err)
	}

	return &CapacityReservationClient{
		Client: client,
	}, nil
}
