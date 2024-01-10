package v2018_06_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/checknameavailability"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/configurations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/databases"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/firewallrules"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/locationbasedperformancetier"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/logfiles"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/recoverableservers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/replicas"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/serverbasedperformancetier"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/serverrestart"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/servers"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/serversecurityalertpolicies"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/updateconfigurations"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/mariadb/2018-06-01-preview/virtualnetworkrules"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	CheckNameAvailability        *checknameavailability.CheckNameAvailabilityClient
	Configurations               *configurations.ConfigurationsClient
	Databases                    *databases.DatabasesClient
	FirewallRules                *firewallrules.FirewallRulesClient
	LocationBasedPerformanceTier *locationbasedperformancetier.LocationBasedPerformanceTierClient
	LogFiles                     *logfiles.LogFilesClient
	RecoverableServers           *recoverableservers.RecoverableServersClient
	Replicas                     *replicas.ReplicasClient
	ServerBasedPerformanceTier   *serverbasedperformancetier.ServerBasedPerformanceTierClient
	ServerRestart                *serverrestart.ServerRestartClient
	ServerSecurityAlertPolicies  *serversecurityalertpolicies.ServerSecurityAlertPoliciesClient
	Servers                      *servers.ServersClient
	UpdateConfigurations         *updateconfigurations.UpdateConfigurationsClient
	VirtualNetworkRules          *virtualnetworkrules.VirtualNetworkRulesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	checkNameAvailabilityClient, err := checknameavailability.NewCheckNameAvailabilityClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CheckNameAvailability client: %+v", err)
	}
	configureFunc(checkNameAvailabilityClient.Client)

	configurationsClient, err := configurations.NewConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Configurations client: %+v", err)
	}
	configureFunc(configurationsClient.Client)

	databasesClient, err := databases.NewDatabasesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Databases client: %+v", err)
	}
	configureFunc(databasesClient.Client)

	firewallRulesClient, err := firewallrules.NewFirewallRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FirewallRules client: %+v", err)
	}
	configureFunc(firewallRulesClient.Client)

	locationBasedPerformanceTierClient, err := locationbasedperformancetier.NewLocationBasedPerformanceTierClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LocationBasedPerformanceTier client: %+v", err)
	}
	configureFunc(locationBasedPerformanceTierClient.Client)

	logFilesClient, err := logfiles.NewLogFilesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogFiles client: %+v", err)
	}
	configureFunc(logFilesClient.Client)

	recoverableServersClient, err := recoverableservers.NewRecoverableServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RecoverableServers client: %+v", err)
	}
	configureFunc(recoverableServersClient.Client)

	replicasClient, err := replicas.NewReplicasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Replicas client: %+v", err)
	}
	configureFunc(replicasClient.Client)

	serverBasedPerformanceTierClient, err := serverbasedperformancetier.NewServerBasedPerformanceTierClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerBasedPerformanceTier client: %+v", err)
	}
	configureFunc(serverBasedPerformanceTierClient.Client)

	serverRestartClient, err := serverrestart.NewServerRestartClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerRestart client: %+v", err)
	}
	configureFunc(serverRestartClient.Client)

	serverSecurityAlertPoliciesClient, err := serversecurityalertpolicies.NewServerSecurityAlertPoliciesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServerSecurityAlertPolicies client: %+v", err)
	}
	configureFunc(serverSecurityAlertPoliciesClient.Client)

	serversClient, err := servers.NewServersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Servers client: %+v", err)
	}
	configureFunc(serversClient.Client)

	updateConfigurationsClient, err := updateconfigurations.NewUpdateConfigurationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UpdateConfigurations client: %+v", err)
	}
	configureFunc(updateConfigurationsClient.Client)

	virtualNetworkRulesClient, err := virtualnetworkrules.NewVirtualNetworkRulesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworkRules client: %+v", err)
	}
	configureFunc(virtualNetworkRulesClient.Client)

	return &Client{
		CheckNameAvailability:        checkNameAvailabilityClient,
		Configurations:               configurationsClient,
		Databases:                    databasesClient,
		FirewallRules:                firewallRulesClient,
		LocationBasedPerformanceTier: locationBasedPerformanceTierClient,
		LogFiles:                     logFilesClient,
		RecoverableServers:           recoverableServersClient,
		Replicas:                     replicasClient,
		ServerBasedPerformanceTier:   serverBasedPerformanceTierClient,
		ServerRestart:                serverRestartClient,
		ServerSecurityAlertPolicies:  serverSecurityAlertPoliciesClient,
		Servers:                      serversClient,
		UpdateConfigurations:         updateConfigurationsClient,
		VirtualNetworkRules:          virtualNetworkRulesClient,
	}, nil
}
