package hypervsites

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PutSiteOperationResponse struct {
	HttpResponse *http.Response
	Model        *HyperVSite
}

// PutSite ...
func (c HyperVSitesClient) PutSite(ctx context.Context, id HyperVSiteId, input HyperVSite) (result PutSiteOperationResponse, err error) {
	req, err := c.preparerForPutSite(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "PutSite", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "PutSite", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForPutSite(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "hypervsites.HyperVSitesClient", "PutSite", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForPutSite prepares the PutSite request.
func (c HyperVSitesClient) preparerForPutSite(ctx context.Context, id HyperVSiteId, input HyperVSite) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForPutSite handles the response to the PutSite request. The method always
// closes the http.Response Body.
func (c HyperVSitesClient) responderForPutSite(resp *http.Response) (result PutSiteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}