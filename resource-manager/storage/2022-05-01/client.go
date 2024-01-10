package v2022_05_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/blobcontainers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/blobinventorypolicies"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/blobservice"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/deletedaccounts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/encryptionscopes"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/fileservice"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/fileshares"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/localusers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/managementpolicies"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/objectreplicationpolicies"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/privateendpointconnections"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/privatelinkresources"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/queueservice"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/queueserviceproperties"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/skus"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/storageaccounts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/tableservice"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/storage/2022-05-01/tableserviceproperties"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	BlobContainers             *blobcontainers.BlobContainersClient
	BlobInventoryPolicies      *blobinventorypolicies.BlobInventoryPoliciesClient
	BlobService                *blobservice.BlobServiceClient
	DeletedAccounts            *deletedaccounts.DeletedAccountsClient
	EncryptionScopes           *encryptionscopes.EncryptionScopesClient
	FileService                *fileservice.FileServiceClient
	FileShares                 *fileshares.FileSharesClient
	LocalUsers                 *localusers.LocalUsersClient
	ManagementPolicies         *managementpolicies.ManagementPoliciesClient
	ObjectReplicationPolicies  *objectreplicationpolicies.ObjectReplicationPoliciesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	QueueService               *queueservice.QueueServiceClient
	QueueServiceProperties     *queueserviceproperties.QueueServicePropertiesClient
	Skus                       *skus.SkusClient
	StorageAccounts            *storageaccounts.StorageAccountsClient
	TableService               *tableservice.TableServiceClient
	TableServiceProperties     *tableserviceproperties.TableServicePropertiesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	blobContainersClient, err := blobcontainers.NewBlobContainersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BlobContainers client: %+v", err)
	}
	configureFunc(blobContainersClient.Client)

	blobInventoryPoliciesClient, err := blobinventorypolicies.NewBlobInventoryPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BlobInventoryPolicies client: %+v", err)
	}
	configureFunc(blobInventoryPoliciesClient.Client)

	blobServiceClient, err := blobservice.NewBlobServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BlobService client: %+v", err)
	}
	configureFunc(blobServiceClient.Client)

	deletedAccountsClient, err := deletedaccounts.NewDeletedAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedAccounts client: %+v", err)
	}
	configureFunc(deletedAccountsClient.Client)

	encryptionScopesClient, err := encryptionscopes.NewEncryptionScopesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building EncryptionScopes client: %+v", err)
	}
	configureFunc(encryptionScopesClient.Client)

	fileServiceClient, err := fileservice.NewFileServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileService client: %+v", err)
	}
	configureFunc(fileServiceClient.Client)

	fileSharesClient, err := fileshares.NewFileSharesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileShares client: %+v", err)
	}
	configureFunc(fileSharesClient.Client)

	localUsersClient, err := localusers.NewLocalUsersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocalUsers client: %+v", err)
	}
	configureFunc(localUsersClient.Client)

	managementPoliciesClient, err := managementpolicies.NewManagementPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagementPolicies client: %+v", err)
	}
	configureFunc(managementPoliciesClient.Client)

	objectReplicationPoliciesClient, err := objectreplicationpolicies.NewObjectReplicationPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ObjectReplicationPolicies client: %+v", err)
	}
	configureFunc(objectReplicationPoliciesClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	queueServiceClient, err := queueservice.NewQueueServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueueService client: %+v", err)
	}
	configureFunc(queueServiceClient.Client)

	queueServicePropertiesClient, err := queueserviceproperties.NewQueueServicePropertiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building QueueServiceProperties client: %+v", err)
	}
	configureFunc(queueServicePropertiesClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	storageAccountsClient, err := storageaccounts.NewStorageAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageAccounts client: %+v", err)
	}
	configureFunc(storageAccountsClient.Client)

	tableServiceClient, err := tableservice.NewTableServiceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TableService client: %+v", err)
	}
	configureFunc(tableServiceClient.Client)

	tableServicePropertiesClient, err := tableserviceproperties.NewTableServicePropertiesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TableServiceProperties client: %+v", err)
	}
	configureFunc(tableServicePropertiesClient.Client)

	return &Client{
		BlobContainers:             blobContainersClient,
		BlobInventoryPolicies:      blobInventoryPoliciesClient,
		BlobService:                blobServiceClient,
		DeletedAccounts:            deletedAccountsClient,
		EncryptionScopes:           encryptionScopesClient,
		FileService:                fileServiceClient,
		FileShares:                 fileSharesClient,
		LocalUsers:                 localUsersClient,
		ManagementPolicies:         managementPoliciesClient,
		ObjectReplicationPolicies:  objectReplicationPoliciesClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		QueueService:               queueServiceClient,
		QueueServiceProperties:     queueServicePropertiesClient,
		Skus:                       skusClient,
		StorageAccounts:            storageAccountsClient,
		TableService:               tableServiceClient,
		TableServiceProperties:     tableServicePropertiesClient,
	}, nil
}
