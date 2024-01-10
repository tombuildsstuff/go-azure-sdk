package servergroupoperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/pollers"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerGroupsRestartOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ServerGroupsRestart ...
func (c ServerGroupOperationsClient) ServerGroupsRestart(ctx context.Context, id ServerGroupsv2Id) (result ServerGroupsRestartOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/restart", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// ServerGroupsRestartThenPoll performs ServerGroupsRestart then polls until it's completed
func (c ServerGroupOperationsClient) ServerGroupsRestartThenPoll(ctx context.Context, id ServerGroupsv2Id) error {
	result, err := c.ServerGroupsRestart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServerGroupsRestart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ServerGroupsRestart: %+v", err)
	}

	return nil
}
