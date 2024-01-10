package policyfragment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tombuildsstuff/go-azure-sdk/sdk/client"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacePolicyFragmentListReferencesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *ResourceCollection
}

type WorkspacePolicyFragmentListReferencesOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultWorkspacePolicyFragmentListReferencesOperationOptions() WorkspacePolicyFragmentListReferencesOperationOptions {
	return WorkspacePolicyFragmentListReferencesOperationOptions{}
}

func (o WorkspacePolicyFragmentListReferencesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WorkspacePolicyFragmentListReferencesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o WorkspacePolicyFragmentListReferencesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// WorkspacePolicyFragmentListReferences ...
func (c PolicyFragmentClient) WorkspacePolicyFragmentListReferences(ctx context.Context, id WorkspacePolicyFragmentId, options WorkspacePolicyFragmentListReferencesOperationOptions) (result WorkspacePolicyFragmentListReferencesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/listReferences", id.ID()),
		OptionsObject: options,
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

	if err = resp.Unmarshal(&result.Model); err != nil {
		return
	}

	return
}
