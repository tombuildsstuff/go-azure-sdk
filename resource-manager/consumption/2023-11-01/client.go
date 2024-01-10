package v2023_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/aggregatedcost"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/balances"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/budgets"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/charges"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/credits"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/events"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/lots"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/marketplaces"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/pricesheet"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationdetails"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationrecommendationdetails"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationrecommendations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationsummaries"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/reservationtransactions"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-11-01/usagedetails"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	AggregatedCost                   *aggregatedcost.AggregatedCostClient
	Balances                         *balances.BalancesClient
	Budgets                          *budgets.BudgetsClient
	Charges                          *charges.ChargesClient
	Credits                          *credits.CreditsClient
	Events                           *events.EventsClient
	Lots                             *lots.LotsClient
	Marketplaces                     *marketplaces.MarketplacesClient
	PriceSheet                       *pricesheet.PriceSheetClient
	ReservationDetails               *reservationdetails.ReservationDetailsClient
	ReservationRecommendationDetails *reservationrecommendationdetails.ReservationRecommendationDetailsClient
	ReservationRecommendations       *reservationrecommendations.ReservationRecommendationsClient
	ReservationSummaries             *reservationsummaries.ReservationSummariesClient
	ReservationTransactions          *reservationtransactions.ReservationTransactionsClient
	UsageDetails                     *usagedetails.UsageDetailsClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	aggregatedCostClient, err := aggregatedcost.NewAggregatedCostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AggregatedCost client: %+v", err)
	}
	configureFunc(aggregatedCostClient.Client)

	balancesClient, err := balances.NewBalancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Balances client: %+v", err)
	}
	configureFunc(balancesClient.Client)

	budgetsClient, err := budgets.NewBudgetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Budgets client: %+v", err)
	}
	configureFunc(budgetsClient.Client)

	chargesClient, err := charges.NewChargesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Charges client: %+v", err)
	}
	configureFunc(chargesClient.Client)

	creditsClient, err := credits.NewCreditsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Credits client: %+v", err)
	}
	configureFunc(creditsClient.Client)

	eventsClient, err := events.NewEventsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Events client: %+v", err)
	}
	configureFunc(eventsClient.Client)

	lotsClient, err := lots.NewLotsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Lots client: %+v", err)
	}
	configureFunc(lotsClient.Client)

	marketplacesClient, err := marketplaces.NewMarketplacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Marketplaces client: %+v", err)
	}
	configureFunc(marketplacesClient.Client)

	priceSheetClient, err := pricesheet.NewPriceSheetClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PriceSheet client: %+v", err)
	}
	configureFunc(priceSheetClient.Client)

	reservationDetailsClient, err := reservationdetails.NewReservationDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationDetails client: %+v", err)
	}
	configureFunc(reservationDetailsClient.Client)

	reservationRecommendationDetailsClient, err := reservationrecommendationdetails.NewReservationRecommendationDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationRecommendationDetails client: %+v", err)
	}
	configureFunc(reservationRecommendationDetailsClient.Client)

	reservationRecommendationsClient, err := reservationrecommendations.NewReservationRecommendationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationRecommendations client: %+v", err)
	}
	configureFunc(reservationRecommendationsClient.Client)

	reservationSummariesClient, err := reservationsummaries.NewReservationSummariesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationSummaries client: %+v", err)
	}
	configureFunc(reservationSummariesClient.Client)

	reservationTransactionsClient, err := reservationtransactions.NewReservationTransactionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ReservationTransactions client: %+v", err)
	}
	configureFunc(reservationTransactionsClient.Client)

	usageDetailsClient, err := usagedetails.NewUsageDetailsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UsageDetails client: %+v", err)
	}
	configureFunc(usageDetailsClient.Client)

	return &Client{
		AggregatedCost:                   aggregatedCostClient,
		Balances:                         balancesClient,
		Budgets:                          budgetsClient,
		Charges:                          chargesClient,
		Credits:                          creditsClient,
		Events:                           eventsClient,
		Lots:                             lotsClient,
		Marketplaces:                     marketplacesClient,
		PriceSheet:                       priceSheetClient,
		ReservationDetails:               reservationDetailsClient,
		ReservationRecommendationDetails: reservationRecommendationDetailsClient,
		ReservationRecommendations:       reservationRecommendationsClient,
		ReservationSummaries:             reservationSummariesClient,
		ReservationTransactions:          reservationTransactionsClient,
		UsageDetails:                     usageDetailsClient,
	}, nil
}
