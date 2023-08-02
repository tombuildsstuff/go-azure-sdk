package registrymanagement

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Registry struct {
	DiscoveryUrl                  *string                              `json:"discoveryUrl,omitempty"`
	IntellectualPropertyPublisher *string                              `json:"intellectualPropertyPublisher,omitempty"`
	ManagedResourceGroup          *ArmResourceId                       `json:"managedResourceGroup,omitempty"`
	MlFlowRegistryUri             *string                              `json:"mlFlowRegistryUri,omitempty"`
	PrivateEndpointConnections    *[]RegistryPrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	PublicNetworkAccess           *string                              `json:"publicNetworkAccess,omitempty"`
	RegionDetails                 *[]RegistryRegionArmDetails          `json:"regionDetails,omitempty"`
}
