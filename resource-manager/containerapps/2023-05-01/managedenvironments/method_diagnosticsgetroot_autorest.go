package managedenvironments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiagnosticsGetRootOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedEnvironment
}

// DiagnosticsGetRoot ...
func (c ManagedEnvironmentsClient) DiagnosticsGetRoot(ctx context.Context, id ManagedEnvironmentId) (result DiagnosticsGetRootOperationResponse, err error) {
	req, err := c.preparerForDiagnosticsGetRoot(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedenvironments.ManagedEnvironmentsClient", "DiagnosticsGetRoot", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedenvironments.ManagedEnvironmentsClient", "DiagnosticsGetRoot", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDiagnosticsGetRoot(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedenvironments.ManagedEnvironmentsClient", "DiagnosticsGetRoot", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDiagnosticsGetRoot prepares the DiagnosticsGetRoot request.
func (c ManagedEnvironmentsClient) preparerForDiagnosticsGetRoot(ctx context.Context, id ManagedEnvironmentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detectorProperties/rootApi", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDiagnosticsGetRoot handles the response to the DiagnosticsGetRoot request. The method always
// closes the http.Response Body.
func (c ManagedEnvironmentsClient) responderForDiagnosticsGetRoot(resp *http.Response) (result DiagnosticsGetRootOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
