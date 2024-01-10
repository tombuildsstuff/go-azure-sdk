package v2022_12_12

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-12-12/dpscertificate"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/deviceprovisioningservices/2022-12-12/iotdpsresource"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	DpsCertificate *dpscertificate.DpsCertificateClient
	IotDpsResource *iotdpsresource.IotDpsResourceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dpsCertificateClient, err := dpscertificate.NewDpsCertificateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DpsCertificate client: %+v", err)
	}
	configureFunc(dpsCertificateClient.Client)

	iotDpsResourceClient, err := iotdpsresource.NewIotDpsResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IotDpsResource client: %+v", err)
	}
	configureFunc(iotDpsResourceClient.Client)

	return &Client{
		DpsCertificate: dpsCertificateClient,
		IotDpsResource: iotDpsResourceClient,
	}, nil
}
