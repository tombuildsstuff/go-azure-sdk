package v2020_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/frontdoor/2020-05-01/checkfrontdoornameavailability"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/frontdoor/2020-05-01/checkfrontdoornameavailabilitywithsubscription"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/frontdoor/2020-05-01/frontdoors"
)

type Client struct {
	CheckFrontDoorNameAvailability                 *checkfrontdoornameavailability.CheckFrontDoorNameAvailabilityClient
	CheckFrontDoorNameAvailabilityWithSubscription *checkfrontdoornameavailabilitywithsubscription.CheckFrontDoorNameAvailabilityWithSubscriptionClient
	FrontDoors                                     *frontdoors.FrontDoorsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	checkFrontDoorNameAvailabilityClient := checkfrontdoornameavailability.NewCheckFrontDoorNameAvailabilityClientWithBaseURI(endpoint)
	configureAuthFunc(&checkFrontDoorNameAvailabilityClient.Client)

	checkFrontDoorNameAvailabilityWithSubscriptionClient := checkfrontdoornameavailabilitywithsubscription.NewCheckFrontDoorNameAvailabilityWithSubscriptionClientWithBaseURI(endpoint)
	configureAuthFunc(&checkFrontDoorNameAvailabilityWithSubscriptionClient.Client)

	frontDoorsClient := frontdoors.NewFrontDoorsClientWithBaseURI(endpoint)
	configureAuthFunc(&frontDoorsClient.Client)

	return Client{
		CheckFrontDoorNameAvailability:                 &checkFrontDoorNameAvailabilityClient,
		CheckFrontDoorNameAvailabilityWithSubscription: &checkFrontDoorNameAvailabilityWithSubscriptionClient,
		FrontDoors: &frontDoorsClient,
	}
}
