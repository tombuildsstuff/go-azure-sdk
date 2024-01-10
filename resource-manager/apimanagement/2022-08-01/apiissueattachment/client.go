package apiissueattachment

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiIssueAttachmentClient struct {
	Client *resourcemanager.Client
}

func NewApiIssueAttachmentClientWithBaseURI(sdkApi sdkEnv.Api) (*ApiIssueAttachmentClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "apiissueattachment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApiIssueAttachmentClient: %+v", err)
	}

	return &ApiIssueAttachmentClient{
		Client: client,
	}, nil
}
