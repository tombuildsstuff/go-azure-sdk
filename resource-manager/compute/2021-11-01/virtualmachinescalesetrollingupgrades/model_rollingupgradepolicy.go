package virtualmachinescalesetrollingupgrades

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RollingUpgradePolicy struct {
	EnableCrossZoneUpgrade              *bool   `json:"enableCrossZoneUpgrade,omitempty"`
	MaxBatchInstancePercent             *int64  `json:"maxBatchInstancePercent,omitempty"`
	MaxUnhealthyInstancePercent         *int64  `json:"maxUnhealthyInstancePercent,omitempty"`
	MaxUnhealthyUpgradedInstancePercent *int64  `json:"maxUnhealthyUpgradedInstancePercent,omitempty"`
	PauseTimeBetweenBatches             *string `json:"pauseTimeBetweenBatches,omitempty"`
	PrioritizeUnhealthyInstances        *bool   `json:"prioritizeUnhealthyInstances,omitempty"`
}
