package policyexemptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListForResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyExemption
}

type ListForResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PolicyExemption
}

type ListForResourceGroupOperationOptions struct {
	Filter *string
}

func DefaultListForResourceGroupOperationOptions() ListForResourceGroupOperationOptions {
	return ListForResourceGroupOperationOptions{}
}

func (o ListForResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListForResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListForResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListForResourceGroup ...
func (c PolicyExemptionsClient) ListForResourceGroup(ctx context.Context, id commonids.ResourceGroupId, options ListForResourceGroupOperationOptions) (result ListForResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/providers/Microsoft.Authorization/policyExemptions", id.ID()),
		OptionsObject: options,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]PolicyExemption `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListForResourceGroupComplete retrieves all the results into a single object
func (c PolicyExemptionsClient) ListForResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId, options ListForResourceGroupOperationOptions) (ListForResourceGroupCompleteResult, error) {
	return c.ListForResourceGroupCompleteMatchingPredicate(ctx, id, options, PolicyExemptionOperationPredicate{})
}

// ListForResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyExemptionsClient) ListForResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, options ListForResourceGroupOperationOptions, predicate PolicyExemptionOperationPredicate) (result ListForResourceGroupCompleteResult, err error) {
	items := make([]PolicyExemption, 0)

	resp, err := c.ListForResourceGroup(ctx, id, options)
	if err != nil {
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListForResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
