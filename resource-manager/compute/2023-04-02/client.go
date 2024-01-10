package v2023_04_02

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-04-02/diskaccesses"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-04-02/diskencryptionsets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-04-02/disks"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-04-02/incrementalrestorepoints"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-04-02/snapshots"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	DiskAccesses             *diskaccesses.DiskAccessesClient
	DiskEncryptionSets       *diskencryptionsets.DiskEncryptionSetsClient
	Disks                    *disks.DisksClient
	IncrementalRestorePoints *incrementalrestorepoints.IncrementalRestorePointsClient
	Snapshots                *snapshots.SnapshotsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	diskAccessesClient, err := diskaccesses.NewDiskAccessesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiskAccesses client: %+v", err)
	}
	configureFunc(diskAccessesClient.Client)

	diskEncryptionSetsClient, err := diskencryptionsets.NewDiskEncryptionSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DiskEncryptionSets client: %+v", err)
	}
	configureFunc(diskEncryptionSetsClient.Client)

	disksClient, err := disks.NewDisksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Disks client: %+v", err)
	}
	configureFunc(disksClient.Client)

	incrementalRestorePointsClient, err := incrementalrestorepoints.NewIncrementalRestorePointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building IncrementalRestorePoints client: %+v", err)
	}
	configureFunc(incrementalRestorePointsClient.Client)

	snapshotsClient, err := snapshots.NewSnapshotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Snapshots client: %+v", err)
	}
	configureFunc(snapshotsClient.Client)

	return &Client{
		DiskAccesses:             diskAccessesClient,
		DiskEncryptionSets:       diskEncryptionSetsClient,
		Disks:                    disksClient,
		IncrementalRestorePoints: incrementalRestorePointsClient,
		Snapshots:                snapshotsClient,
	}, nil
}
