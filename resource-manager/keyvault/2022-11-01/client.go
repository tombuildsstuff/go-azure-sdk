package v2022_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/keys"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/managedhsmkeys"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/managedhsms"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/mhsmlistprivateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/mhsmprivateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/mhsmprivatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/secrets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/keyvault/2022-11-01/vaults"
)

type Client struct {
	Keys                               *keys.KeysClient
	MHSMListPrivateEndpointConnections *mhsmlistprivateendpointconnections.MHSMListPrivateEndpointConnectionsClient
	MHSMPrivateEndpointConnections     *mhsmprivateendpointconnections.MHSMPrivateEndpointConnectionsClient
	MHSMPrivateLinkResources           *mhsmprivatelinkresources.MHSMPrivateLinkResourcesClient
	ManagedHsmKeys                     *managedhsmkeys.ManagedHsmKeysClient
	ManagedHsms                        *managedhsms.ManagedHsmsClient
	PrivateEndpointConnections         *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources               *privatelinkresources.PrivateLinkResourcesClient
	Secrets                            *secrets.SecretsClient
	Vaults                             *vaults.VaultsClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	keysClient := keys.NewKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&keysClient.Client)

	mHSMListPrivateEndpointConnectionsClient := mhsmlistprivateendpointconnections.NewMHSMListPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&mHSMListPrivateEndpointConnectionsClient.Client)

	mHSMPrivateEndpointConnectionsClient := mhsmprivateendpointconnections.NewMHSMPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&mHSMPrivateEndpointConnectionsClient.Client)

	mHSMPrivateLinkResourcesClient := mhsmprivatelinkresources.NewMHSMPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&mHSMPrivateLinkResourcesClient.Client)

	managedHsmKeysClient := managedhsmkeys.NewManagedHsmKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&managedHsmKeysClient.Client)

	managedHsmsClient := managedhsms.NewManagedHsmsClientWithBaseURI(endpoint)
	configureAuthFunc(&managedHsmsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	secretsClient := secrets.NewSecretsClientWithBaseURI(endpoint)
	configureAuthFunc(&secretsClient.Client)

	vaultsClient := vaults.NewVaultsClientWithBaseURI(endpoint)
	configureAuthFunc(&vaultsClient.Client)

	return Client{
		Keys:                               &keysClient,
		MHSMListPrivateEndpointConnections: &mHSMListPrivateEndpointConnectionsClient,
		MHSMPrivateEndpointConnections:     &mHSMPrivateEndpointConnectionsClient,
		MHSMPrivateLinkResources:           &mHSMPrivateLinkResourcesClient,
		ManagedHsmKeys:                     &managedHsmKeysClient,
		ManagedHsms:                        &managedHsmsClient,
		PrivateEndpointConnections:         &privateEndpointConnectionsClient,
		PrivateLinkResources:               &privateLinkResourcesClient,
		Secrets:                            &secretsClient,
		Vaults:                             &vaultsClient,
	}
}
