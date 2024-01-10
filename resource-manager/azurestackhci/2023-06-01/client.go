package v2023_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/arcsettings"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/cluster"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/clusters"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/extensions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/offers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/publishers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/skuses"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/updateruns"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/updates"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/azurestackhci/2023-06-01/updatesummaries"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	ArcSettings     *arcsettings.ArcSettingsClient
	Cluster         *cluster.ClusterClient
	Clusters        *clusters.ClustersClient
	Extensions      *extensions.ExtensionsClient
	Offers          *offers.OffersClient
	Publishers      *publishers.PublishersClient
	Skuses          *skuses.SkusesClient
	UpdateRuns      *updateruns.UpdateRunsClient
	UpdateSummaries *updatesummaries.UpdateSummariesClient
	Updates         *updates.UpdatesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	arcSettingsClient, err := arcsettings.NewArcSettingsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ArcSettings client: %+v", err)
	}
	configureFunc(arcSettingsClient.Client)

	clusterClient, err := cluster.NewClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Cluster client: %+v", err)
	}
	configureFunc(clusterClient.Client)

	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	extensionsClient, err := extensions.NewExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extensions client: %+v", err)
	}
	configureFunc(extensionsClient.Client)

	offersClient, err := offers.NewOffersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Offers client: %+v", err)
	}
	configureFunc(offersClient.Client)

	publishersClient, err := publishers.NewPublishersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Publishers client: %+v", err)
	}
	configureFunc(publishersClient.Client)

	skusesClient, err := skuses.NewSkusesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skuses client: %+v", err)
	}
	configureFunc(skusesClient.Client)

	updateRunsClient, err := updateruns.NewUpdateRunsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateRuns client: %+v", err)
	}
	configureFunc(updateRunsClient.Client)

	updateSummariesClient, err := updatesummaries.NewUpdateSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateSummaries client: %+v", err)
	}
	configureFunc(updateSummariesClient.Client)

	updatesClient, err := updates.NewUpdatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Updates client: %+v", err)
	}
	configureFunc(updatesClient.Client)

	return &Client{
		ArcSettings:     arcSettingsClient,
		Cluster:         clusterClient,
		Clusters:        clustersClient,
		Extensions:      extensionsClient,
		Offers:          offersClient,
		Publishers:      publishersClient,
		Skuses:          skusesClient,
		UpdateRuns:      updateRunsClient,
		UpdateSummaries: updateSummariesClient,
		Updates:         updatesClient,
	}, nil
}
