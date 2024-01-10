package v2020_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/privatedns/2020-06-01/privatezones"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/privatedns/2020-06-01/recordsets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/privatedns/2020-06-01/virtualnetworklinks"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	PrivateZones        *privatezones.PrivateZonesClient
	RecordSets          *recordsets.RecordSetsClient
	VirtualNetworkLinks *virtualnetworklinks.VirtualNetworkLinksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	privateZonesClient, err := privatezones.NewPrivateZonesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateZones client: %+v", err)
	}
	configureFunc(privateZonesClient.Client)

	recordSetsClient, err := recordsets.NewRecordSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecordSets client: %+v", err)
	}
	configureFunc(recordSetsClient.Client)

	virtualNetworkLinksClient, err := virtualnetworklinks.NewVirtualNetworkLinksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkLinks client: %+v", err)
	}
	configureFunc(virtualNetworkLinksClient.Client)

	return &Client{
		PrivateZones:        privateZonesClient,
		RecordSets:          recordSetsClient,
		VirtualNetworkLinks: virtualNetworkLinksClient,
	}, nil
}
