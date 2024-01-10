package v2023_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/backuppolicy"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/capacitypools"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/filelocks"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/groupidlistforldapuser"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/netappaccounts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/netappresource"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/poolchange"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/resetcifspassword"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/restore"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/snapshotpolicy"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/snapshotpolicylistvolumes"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/snapshots"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/subvolumes"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/volumegroups"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/volumequotarules"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/volumes"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/volumesrelocation"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/volumesreplication"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-05-01/volumesrevert"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	BackupPolicy              *backuppolicy.BackupPolicyClient
	CapacityPools             *capacitypools.CapacityPoolsClient
	FileLocks                 *filelocks.FileLocksClient
	GroupIdListForLDAPUser    *groupidlistforldapuser.GroupIdListForLDAPUserClient
	NetAppAccounts            *netappaccounts.NetAppAccountsClient
	NetAppResource            *netappresource.NetAppResourceClient
	PoolChange                *poolchange.PoolChangeClient
	ResetCifsPassword         *resetcifspassword.ResetCifsPasswordClient
	Restore                   *restore.RestoreClient
	SnapshotPolicy            *snapshotpolicy.SnapshotPolicyClient
	SnapshotPolicyListVolumes *snapshotpolicylistvolumes.SnapshotPolicyListVolumesClient
	Snapshots                 *snapshots.SnapshotsClient
	SubVolumes                *subvolumes.SubVolumesClient
	VolumeGroups              *volumegroups.VolumeGroupsClient
	VolumeQuotaRules          *volumequotarules.VolumeQuotaRulesClient
	Volumes                   *volumes.VolumesClient
	VolumesRelocation         *volumesrelocation.VolumesRelocationClient
	VolumesReplication        *volumesreplication.VolumesReplicationClient
	VolumesRevert             *volumesrevert.VolumesRevertClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	backupPolicyClient, err := backuppolicy.NewBackupPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BackupPolicy client: %+v", err)
	}
	configureFunc(backupPolicyClient.Client)

	capacityPoolsClient, err := capacitypools.NewCapacityPoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityPools client: %+v", err)
	}
	configureFunc(capacityPoolsClient.Client)

	fileLocksClient, err := filelocks.NewFileLocksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileLocks client: %+v", err)
	}
	configureFunc(fileLocksClient.Client)

	groupIdListForLDAPUserClient, err := groupidlistforldapuser.NewGroupIdListForLDAPUserClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GroupIdListForLDAPUser client: %+v", err)
	}
	configureFunc(groupIdListForLDAPUserClient.Client)

	netAppAccountsClient, err := netappaccounts.NewNetAppAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetAppAccounts client: %+v", err)
	}
	configureFunc(netAppAccountsClient.Client)

	netAppResourceClient, err := netappresource.NewNetAppResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetAppResource client: %+v", err)
	}
	configureFunc(netAppResourceClient.Client)

	poolChangeClient, err := poolchange.NewPoolChangeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PoolChange client: %+v", err)
	}
	configureFunc(poolChangeClient.Client)

	resetCifsPasswordClient, err := resetcifspassword.NewResetCifsPasswordClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResetCifsPassword client: %+v", err)
	}
	configureFunc(resetCifsPasswordClient.Client)

	restoreClient, err := restore.NewRestoreClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Restore client: %+v", err)
	}
	configureFunc(restoreClient.Client)

	snapshotPolicyClient, err := snapshotpolicy.NewSnapshotPolicyClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicy client: %+v", err)
	}
	configureFunc(snapshotPolicyClient.Client)

	snapshotPolicyListVolumesClient, err := snapshotpolicylistvolumes.NewSnapshotPolicyListVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SnapshotPolicyListVolumes client: %+v", err)
	}
	configureFunc(snapshotPolicyListVolumesClient.Client)

	snapshotsClient, err := snapshots.NewSnapshotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Snapshots client: %+v", err)
	}
	configureFunc(snapshotsClient.Client)

	subVolumesClient, err := subvolumes.NewSubVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SubVolumes client: %+v", err)
	}
	configureFunc(subVolumesClient.Client)

	volumeGroupsClient, err := volumegroups.NewVolumeGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeGroups client: %+v", err)
	}
	configureFunc(volumeGroupsClient.Client)

	volumeQuotaRulesClient, err := volumequotarules.NewVolumeQuotaRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumeQuotaRules client: %+v", err)
	}
	configureFunc(volumeQuotaRulesClient.Client)

	volumesClient, err := volumes.NewVolumesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Volumes client: %+v", err)
	}
	configureFunc(volumesClient.Client)

	volumesRelocationClient, err := volumesrelocation.NewVolumesRelocationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumesRelocation client: %+v", err)
	}
	configureFunc(volumesRelocationClient.Client)

	volumesReplicationClient, err := volumesreplication.NewVolumesReplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumesReplication client: %+v", err)
	}
	configureFunc(volumesReplicationClient.Client)

	volumesRevertClient, err := volumesrevert.NewVolumesRevertClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VolumesRevert client: %+v", err)
	}
	configureFunc(volumesRevertClient.Client)

	return &Client{
		BackupPolicy:              backupPolicyClient,
		CapacityPools:             capacityPoolsClient,
		FileLocks:                 fileLocksClient,
		GroupIdListForLDAPUser:    groupIdListForLDAPUserClient,
		NetAppAccounts:            netAppAccountsClient,
		NetAppResource:            netAppResourceClient,
		PoolChange:                poolChangeClient,
		ResetCifsPassword:         resetCifsPasswordClient,
		Restore:                   restoreClient,
		SnapshotPolicy:            snapshotPolicyClient,
		SnapshotPolicyListVolumes: snapshotPolicyListVolumesClient,
		Snapshots:                 snapshotsClient,
		SubVolumes:                subVolumesClient,
		VolumeGroups:              volumeGroupsClient,
		VolumeQuotaRules:          volumeQuotaRulesClient,
		Volumes:                   volumesClient,
		VolumesRelocation:         volumesRelocationClient,
		VolumesReplication:        volumesReplicationClient,
		VolumesRevert:             volumesRevertClient,
	}, nil
}
