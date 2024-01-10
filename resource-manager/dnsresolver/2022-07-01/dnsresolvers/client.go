package dnsresolvers

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DnsResolversClient struct {
	Client *resourcemanager.Client
}

func NewDnsResolversClientWithBaseURI(sdkApi sdkEnv.Api) (*DnsResolversClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "dnsresolvers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DnsResolversClient: %+v", err)
	}

	return &DnsResolversClient{
		Client: client,
	}, nil
}
