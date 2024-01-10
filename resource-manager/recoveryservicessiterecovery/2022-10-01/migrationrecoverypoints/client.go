package migrationrecoverypoints

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationRecoveryPointsClient struct {
	Client *resourcemanager.Client
}

func NewMigrationRecoveryPointsClientWithBaseURI(sdkApi sdkEnv.Api) (*MigrationRecoveryPointsClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "migrationrecoverypoints", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MigrationRecoveryPointsClient: %+v", err)
	}

	return &MigrationRecoveryPointsClient{
		Client: client,
	}, nil
}
