package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/advisor/2020-01-01/configurations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/advisor/2020-01-01/generaterecommendations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/advisor/2020-01-01/getrecommendations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/advisor/2020-01-01/metadata"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/advisor/2020-01-01/suppressions"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Configurations          *configurations.ConfigurationsClient
	GenerateRecommendations *generaterecommendations.GenerateRecommendationsClient
	GetRecommendations      *getrecommendations.GetRecommendationsClient
	Metadata                *metadata.MetadataClient
	Suppressions            *suppressions.SuppressionsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	generateRecommendationsClient, err := generaterecommendations.NewGenerateRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GenerateRecommendations client: %+v", err)
	}
	configureFunc(generateRecommendationsClient.Client)

	getRecommendationsClient, err := getrecommendations.NewGetRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetRecommendations client: %+v", err)
	}
	configureFunc(getRecommendationsClient.Client)

	metadataClient, err := metadata.NewMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Metadata client: %+v", err)
	}
	configureFunc(metadataClient.Client)

	suppressionsClient, err := suppressions.NewSuppressionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Suppressions client: %+v", err)
	}
	configureFunc(suppressionsClient.Client)

	return &Client{
		Configurations:          configurationsClient,
		GenerateRecommendations: generateRecommendationsClient,
		GetRecommendations:      getRecommendationsClient,
		Metadata:                metadataClient,
		Suppressions:            suppressionsClient,
	}, nil
}
