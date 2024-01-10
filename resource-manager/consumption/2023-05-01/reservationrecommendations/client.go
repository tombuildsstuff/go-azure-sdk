package reservationrecommendations

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendationsClient struct {
	Client *resourcemanager.Client
}

func NewReservationRecommendationsClientWithBaseURI(sdkApi sdkEnv.Api) (*ReservationRecommendationsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "reservationrecommendations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ReservationRecommendationsClient: %+v", err)
	}

	return &ReservationRecommendationsClient{
		Client: client,
	}, nil
}
