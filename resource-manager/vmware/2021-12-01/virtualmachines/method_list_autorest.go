package virtualmachines

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]VirtualMachine

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (ListOperationResponse, error)
}

type ListCompleteResult struct {
	Items []VirtualMachine
}

func (r ListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r ListOperationResponse) LoadMore(ctx context.Context) (resp ListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// List ...
func (c VirtualMachinesClient) List(ctx context.Context, id ClusterId) (resp ListOperationResponse, err error) {
	req, err := c.preparerForList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachines.VirtualMachinesClient", "List", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachines.VirtualMachinesClient", "List", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "virtualmachines.VirtualMachinesClient", "List", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForList prepares the List request.
func (c VirtualMachinesClient) preparerForList(ctx context.Context, id ClusterId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/virtualMachines", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForListWithNextLink prepares the List request with the given nextLink token.
func (c VirtualMachinesClient) preparerForListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForList handles the response to the List request. The method always
// closes the http.Response Body.
func (c VirtualMachinesClient) responderForList(resp *http.Response) (result ListOperationResponse, err error) {
	type page struct {
		Values   []VirtualMachine `json:"value"`
		NextLink *string          `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result ListOperationResponse, err error) {
			req, err := c.preparerForListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachines.VirtualMachinesClient", "List", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachines.VirtualMachinesClient", "List", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "virtualmachines.VirtualMachinesClient", "List", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// ListComplete retrieves all of the results into a single object
func (c VirtualMachinesClient) ListComplete(ctx context.Context, id ClusterId) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, VirtualMachineOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c VirtualMachinesClient) ListCompleteMatchingPredicate(ctx context.Context, id ClusterId, predicate VirtualMachineOperationPredicate) (resp ListCompleteResult, err error) {
	items := make([]VirtualMachine, 0)

	page, err := c.List(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := ListCompleteResult{
		Items: items,
	}
	return out, nil
}