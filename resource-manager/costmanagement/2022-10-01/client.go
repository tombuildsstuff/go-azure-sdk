package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/alerts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/benefitrecommendations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/benefitutilizationsummaries"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/costdetails"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/dimensions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/exports"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/forecast"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/pricesheets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/query"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/reservedinstances"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/scheduledactions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/usagedetails"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/views"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Alerts                      *alerts.AlertsClient
	BenefitRecommendations      *benefitrecommendations.BenefitRecommendationsClient
	BenefitUtilizationSummaries *benefitutilizationsummaries.BenefitUtilizationSummariesClient
	CostDetails                 *costdetails.CostDetailsClient
	Dimensions                  *dimensions.DimensionsClient
	Exports                     *exports.ExportsClient
	Forecast                    *forecast.ForecastClient
	PriceSheets                 *pricesheets.PriceSheetsClient
	Query                       *query.QueryClient
	ReservedInstances           *reservedinstances.ReservedInstancesClient
	ScheduledActions            *scheduledactions.ScheduledActionsClient
	UsageDetails                *usagedetails.UsageDetailsClient
	Views                       *views.ViewsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	alertsClient, err := alerts.NewAlertsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Alerts client: %+v", err)
	}
	configureFunc(alertsClient.Client)

	benefitRecommendationsClient, err := benefitrecommendations.NewBenefitRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BenefitRecommendations client: %+v", err)
	}
	configureFunc(benefitRecommendationsClient.Client)

	benefitUtilizationSummariesClient, err := benefitutilizationsummaries.NewBenefitUtilizationSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building BenefitUtilizationSummaries client: %+v", err)
	}
	configureFunc(benefitUtilizationSummariesClient.Client)

	costDetailsClient, err := costdetails.NewCostDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CostDetails client: %+v", err)
	}
	configureFunc(costDetailsClient.Client)

	dimensionsClient, err := dimensions.NewDimensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Dimensions client: %+v", err)
	}
	configureFunc(dimensionsClient.Client)

	exportsClient, err := exports.NewExportsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Exports client: %+v", err)
	}
	configureFunc(exportsClient.Client)

	forecastClient, err := forecast.NewForecastClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Forecast client: %+v", err)
	}
	configureFunc(forecastClient.Client)

	priceSheetsClient, err := pricesheets.NewPriceSheetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheets client: %+v", err)
	}
	configureFunc(priceSheetsClient.Client)

	queryClient, err := query.NewQueryClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Query client: %+v", err)
	}
	configureFunc(queryClient.Client)

	reservedInstancesClient, err := reservedinstances.NewReservedInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservedInstances client: %+v", err)
	}
	configureFunc(reservedInstancesClient.Client)

	scheduledActionsClient, err := scheduledactions.NewScheduledActionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScheduledActions client: %+v", err)
	}
	configureFunc(scheduledActionsClient.Client)

	usageDetailsClient, err := usagedetails.NewUsageDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsageDetails client: %+v", err)
	}
	configureFunc(usageDetailsClient.Client)

	viewsClient, err := views.NewViewsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Views client: %+v", err)
	}
	configureFunc(viewsClient.Client)

	return &Client{
		Alerts:                      alertsClient,
		BenefitRecommendations:      benefitRecommendationsClient,
		BenefitUtilizationSummaries: benefitUtilizationSummariesClient,
		CostDetails:                 costDetailsClient,
		Dimensions:                  dimensionsClient,
		Exports:                     exportsClient,
		Forecast:                    forecastClient,
		PriceSheets:                 priceSheetsClient,
		Query:                       queryClient,
		ReservedInstances:           reservedInstancesClient,
		ScheduledActions:            scheduledActionsClient,
		UsageDetails:                usageDetailsClient,
		Views:                       viewsClient,
	}, nil
}
