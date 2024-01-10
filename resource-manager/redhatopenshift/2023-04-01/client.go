package v2023_04_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/machinepools"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/openshiftclusters"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/openshiftversions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/secrets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/syncidentityproviders"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/redhatopenshift/2023-04-01/syncsets"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	MachinePools          *machinepools.MachinePoolsClient
	OpenShiftClusters     *openshiftclusters.OpenShiftClustersClient
	OpenShiftVersions     *openshiftversions.OpenShiftVersionsClient
	Secrets               *secrets.SecretsClient
	SyncIdentityProviders *syncidentityproviders.SyncIdentityProvidersClient
	SyncSets              *syncsets.SyncSetsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	machinePoolsClient, err := machinepools.NewMachinePoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachinePools client: %+v", err)
	}
	configureFunc(machinePoolsClient.Client)

	openShiftClustersClient, err := openshiftclusters.NewOpenShiftClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OpenShiftClusters client: %+v", err)
	}
	configureFunc(openShiftClustersClient.Client)

	openShiftVersionsClient, err := openshiftversions.NewOpenShiftVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OpenShiftVersions client: %+v", err)
	}
	configureFunc(openShiftVersionsClient.Client)

	secretsClient, err := secrets.NewSecretsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Secrets client: %+v", err)
	}
	configureFunc(secretsClient.Client)

	syncIdentityProvidersClient, err := syncidentityproviders.NewSyncIdentityProvidersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SyncIdentityProviders client: %+v", err)
	}
	configureFunc(syncIdentityProvidersClient.Client)

	syncSetsClient, err := syncsets.NewSyncSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SyncSets client: %+v", err)
	}
	configureFunc(syncSetsClient.Client)

	return &Client{
		MachinePools:          machinePoolsClient,
		OpenShiftClusters:     openShiftClustersClient,
		OpenShiftVersions:     openShiftVersionsClient,
		Secrets:               secretsClient,
		SyncIdentityProviders: syncIdentityProvidersClient,
		SyncSets:              syncSetsClient,
	}, nil
}
