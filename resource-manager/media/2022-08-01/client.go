package v2022_08_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/accountfilters"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/assetsandassetfilters"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/contentkeypolicies"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/liveevents"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/liveoutputs"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/streamingendpoints"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/media/2022-08-01/streamingpoliciesandstreaminglocators"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	AccountFilters                        *accountfilters.AccountFiltersClient
	AssetsAndAssetFilters                 *assetsandassetfilters.AssetsAndAssetFiltersClient
	ContentKeyPolicies                    *contentkeypolicies.ContentKeyPoliciesClient
	LiveEvents                            *liveevents.LiveEventsClient
	LiveOutputs                           *liveoutputs.LiveOutputsClient
	StreamingEndpoints                    *streamingendpoints.StreamingEndpointsClient
	StreamingPoliciesAndStreamingLocators *streamingpoliciesandstreaminglocators.StreamingPoliciesAndStreamingLocatorsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accountFiltersClient, err := accountfilters.NewAccountFiltersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AccountFilters client: %+v", err)
	}
	configureFunc(accountFiltersClient.Client)

	assetsAndAssetFiltersClient, err := assetsandassetfilters.NewAssetsAndAssetFiltersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AssetsAndAssetFilters client: %+v", err)
	}
	configureFunc(assetsAndAssetFiltersClient.Client)

	contentKeyPoliciesClient, err := contentkeypolicies.NewContentKeyPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ContentKeyPolicies client: %+v", err)
	}
	configureFunc(contentKeyPoliciesClient.Client)

	liveEventsClient, err := liveevents.NewLiveEventsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LiveEvents client: %+v", err)
	}
	configureFunc(liveEventsClient.Client)

	liveOutputsClient, err := liveoutputs.NewLiveOutputsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LiveOutputs client: %+v", err)
	}
	configureFunc(liveOutputsClient.Client)

	streamingEndpointsClient, err := streamingendpoints.NewStreamingEndpointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StreamingEndpoints client: %+v", err)
	}
	configureFunc(streamingEndpointsClient.Client)

	streamingPoliciesAndStreamingLocatorsClient, err := streamingpoliciesandstreaminglocators.NewStreamingPoliciesAndStreamingLocatorsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StreamingPoliciesAndStreamingLocators client: %+v", err)
	}
	configureFunc(streamingPoliciesAndStreamingLocatorsClient.Client)

	return &Client{
		AccountFilters:                        accountFiltersClient,
		AssetsAndAssetFilters:                 assetsAndAssetFiltersClient,
		ContentKeyPolicies:                    contentKeyPoliciesClient,
		LiveEvents:                            liveEventsClient,
		LiveOutputs:                           liveOutputsClient,
		StreamingEndpoints:                    streamingEndpointsClient,
		StreamingPoliciesAndStreamingLocators: streamingPoliciesAndStreamingLocatorsClient,
	}, nil
}
