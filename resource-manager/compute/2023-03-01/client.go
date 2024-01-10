package v2023_03_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/availabilitysets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/capacityreservation"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/capacityreservationgroups"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/capacityreservations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/dedicatedhost"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/dedicatedhostgroups"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/dedicatedhosts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/images"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/loganalytics"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/proximityplacementgroups"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/restorepointcollections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/restorepoints"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/sshpublickeys"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachineextensionimages"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachineextensions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachineimages"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachineruncommands"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachines"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinescalesetextensions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinescalesetrollingupgrades"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinescalesets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinescalesetvmextensions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinescalesetvmruncommands"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinescalesetvms"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/compute/2023-03-01/virtualmachinesizes"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailabilitySets                      *availabilitysets.AvailabilitySetsClient
	CapacityReservation                   *capacityreservation.CapacityReservationClient
	CapacityReservationGroups             *capacityreservationgroups.CapacityReservationGroupsClient
	CapacityReservations                  *capacityreservations.CapacityReservationsClient
	DedicatedHost                         *dedicatedhost.DedicatedHostClient
	DedicatedHostGroups                   *dedicatedhostgroups.DedicatedHostGroupsClient
	DedicatedHosts                        *dedicatedhosts.DedicatedHostsClient
	Images                                *images.ImagesClient
	LogAnalytics                          *loganalytics.LogAnalyticsClient
	ProximityPlacementGroups              *proximityplacementgroups.ProximityPlacementGroupsClient
	RestorePointCollections               *restorepointcollections.RestorePointCollectionsClient
	RestorePoints                         *restorepoints.RestorePointsClient
	SshPublicKeys                         *sshpublickeys.SshPublicKeysClient
	VirtualMachineExtensionImages         *virtualmachineextensionimages.VirtualMachineExtensionImagesClient
	VirtualMachineExtensions              *virtualmachineextensions.VirtualMachineExtensionsClient
	VirtualMachineImages                  *virtualmachineimages.VirtualMachineImagesClient
	VirtualMachineRunCommands             *virtualmachineruncommands.VirtualMachineRunCommandsClient
	VirtualMachineScaleSetExtensions      *virtualmachinescalesetextensions.VirtualMachineScaleSetExtensionsClient
	VirtualMachineScaleSetRollingUpgrades *virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient
	VirtualMachineScaleSetVMExtensions    *virtualmachinescalesetvmextensions.VirtualMachineScaleSetVMExtensionsClient
	VirtualMachineScaleSetVMRunCommands   *virtualmachinescalesetvmruncommands.VirtualMachineScaleSetVMRunCommandsClient
	VirtualMachineScaleSetVMs             *virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient
	VirtualMachineScaleSets               *virtualmachinescalesets.VirtualMachineScaleSetsClient
	VirtualMachineSizes                   *virtualmachinesizes.VirtualMachineSizesClient
	VirtualMachines                       *virtualmachines.VirtualMachinesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	availabilitySetsClient, err := availabilitysets.NewAvailabilitySetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailabilitySets client: %+v", err)
	}
	configureFunc(availabilitySetsClient.Client)

	capacityReservationClient, err := capacityreservation.NewCapacityReservationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityReservation client: %+v", err)
	}
	configureFunc(capacityReservationClient.Client)

	capacityReservationGroupsClient, err := capacityreservationgroups.NewCapacityReservationGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityReservationGroups client: %+v", err)
	}
	configureFunc(capacityReservationGroupsClient.Client)

	capacityReservationsClient, err := capacityreservations.NewCapacityReservationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityReservations client: %+v", err)
	}
	configureFunc(capacityReservationsClient.Client)

	dedicatedHostClient, err := dedicatedhost.NewDedicatedHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHost client: %+v", err)
	}
	configureFunc(dedicatedHostClient.Client)

	dedicatedHostGroupsClient, err := dedicatedhostgroups.NewDedicatedHostGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHostGroups client: %+v", err)
	}
	configureFunc(dedicatedHostGroupsClient.Client)

	dedicatedHostsClient, err := dedicatedhosts.NewDedicatedHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHosts client: %+v", err)
	}
	configureFunc(dedicatedHostsClient.Client)

	imagesClient, err := images.NewImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Images client: %+v", err)
	}
	configureFunc(imagesClient.Client)

	logAnalyticsClient, err := loganalytics.NewLogAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogAnalytics client: %+v", err)
	}
	configureFunc(logAnalyticsClient.Client)

	proximityPlacementGroupsClient, err := proximityplacementgroups.NewProximityPlacementGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProximityPlacementGroups client: %+v", err)
	}
	configureFunc(proximityPlacementGroupsClient.Client)

	restorePointCollectionsClient, err := restorepointcollections.NewRestorePointCollectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorePointCollections client: %+v", err)
	}
	configureFunc(restorePointCollectionsClient.Client)

	restorePointsClient, err := restorepoints.NewRestorePointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorePoints client: %+v", err)
	}
	configureFunc(restorePointsClient.Client)

	sshPublicKeysClient, err := sshpublickeys.NewSshPublicKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SshPublicKeys client: %+v", err)
	}
	configureFunc(sshPublicKeysClient.Client)

	virtualMachineExtensionImagesClient, err := virtualmachineextensionimages.NewVirtualMachineExtensionImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineExtensionImages client: %+v", err)
	}
	configureFunc(virtualMachineExtensionImagesClient.Client)

	virtualMachineExtensionsClient, err := virtualmachineextensions.NewVirtualMachineExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineExtensions client: %+v", err)
	}
	configureFunc(virtualMachineExtensionsClient.Client)

	virtualMachineImagesClient, err := virtualmachineimages.NewVirtualMachineImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineImages client: %+v", err)
	}
	configureFunc(virtualMachineImagesClient.Client)

	virtualMachineRunCommandsClient, err := virtualmachineruncommands.NewVirtualMachineRunCommandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineRunCommands client: %+v", err)
	}
	configureFunc(virtualMachineRunCommandsClient.Client)

	virtualMachineScaleSetExtensionsClient, err := virtualmachinescalesetextensions.NewVirtualMachineScaleSetExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetExtensions client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetExtensionsClient.Client)

	virtualMachineScaleSetRollingUpgradesClient, err := virtualmachinescalesetrollingupgrades.NewVirtualMachineScaleSetRollingUpgradesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetRollingUpgrades client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetRollingUpgradesClient.Client)

	virtualMachineScaleSetVMExtensionsClient, err := virtualmachinescalesetvmextensions.NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetVMExtensions client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetVMExtensionsClient.Client)

	virtualMachineScaleSetVMRunCommandsClient, err := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetVMRunCommands client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetVMRunCommandsClient.Client)

	virtualMachineScaleSetVMsClient, err := virtualmachinescalesetvms.NewVirtualMachineScaleSetVMsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetVMs client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetVMsClient.Client)

	virtualMachineScaleSetsClient, err := virtualmachinescalesets.NewVirtualMachineScaleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSets client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetsClient.Client)

	virtualMachineSizesClient, err := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineSizes client: %+v", err)
	}
	configureFunc(virtualMachineSizesClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	return &Client{
		AvailabilitySets:                      availabilitySetsClient,
		CapacityReservation:                   capacityReservationClient,
		CapacityReservationGroups:             capacityReservationGroupsClient,
		CapacityReservations:                  capacityReservationsClient,
		DedicatedHost:                         dedicatedHostClient,
		DedicatedHostGroups:                   dedicatedHostGroupsClient,
		DedicatedHosts:                        dedicatedHostsClient,
		Images:                                imagesClient,
		LogAnalytics:                          logAnalyticsClient,
		ProximityPlacementGroups:              proximityPlacementGroupsClient,
		RestorePointCollections:               restorePointCollectionsClient,
		RestorePoints:                         restorePointsClient,
		SshPublicKeys:                         sshPublicKeysClient,
		VirtualMachineExtensionImages:         virtualMachineExtensionImagesClient,
		VirtualMachineExtensions:              virtualMachineExtensionsClient,
		VirtualMachineImages:                  virtualMachineImagesClient,
		VirtualMachineRunCommands:             virtualMachineRunCommandsClient,
		VirtualMachineScaleSetExtensions:      virtualMachineScaleSetExtensionsClient,
		VirtualMachineScaleSetRollingUpgrades: virtualMachineScaleSetRollingUpgradesClient,
		VirtualMachineScaleSetVMExtensions:    virtualMachineScaleSetVMExtensionsClient,
		VirtualMachineScaleSetVMRunCommands:   virtualMachineScaleSetVMRunCommandsClient,
		VirtualMachineScaleSetVMs:             virtualMachineScaleSetVMsClient,
		VirtualMachineScaleSets:               virtualMachineScaleSetsClient,
		VirtualMachineSizes:                   virtualMachineSizesClient,
		VirtualMachines:                       virtualMachinesClient,
	}, nil
}
