package v2023_03_31

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/communication/2023-03-31/communicationservices"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/communication/2023-03-31/domains"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/communication/2023-03-31/emailservices"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/communication/2023-03-31/senderusernames"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	CommunicationServices *communicationservices.CommunicationServicesClient
	Domains               *domains.DomainsClient
	EmailServices         *emailservices.EmailServicesClient
	SenderUsernames       *senderusernames.SenderUsernamesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	communicationServicesClient, err := communicationservices.NewCommunicationServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunicationServices client: %+v", err)
	}
	configureFunc(communicationServicesClient.Client)

	domainsClient, err := domains.NewDomainsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Domains client: %+v", err)
	}
	configureFunc(domainsClient.Client)

	emailServicesClient, err := emailservices.NewEmailServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EmailServices client: %+v", err)
	}
	configureFunc(emailServicesClient.Client)

	senderUsernamesClient, err := senderusernames.NewSenderUsernamesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SenderUsernames client: %+v", err)
	}
	configureFunc(senderUsernamesClient.Client)

	return &Client{
		CommunicationServices: communicationServicesClient,
		Domains:               domainsClient,
		EmailServices:         emailServicesClient,
		SenderUsernames:       senderUsernamesClient,
	}, nil
}
