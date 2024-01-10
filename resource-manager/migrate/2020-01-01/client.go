package v2020_01_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervcluster"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervhost"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervjobs"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervmachines"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervrunasaccounts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/hypervsites"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/jobs"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/machines"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/runasaccounts"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/sites"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/migrate/2020-01-01/vcenter"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	HyperVCluster       *hypervcluster.HyperVClusterClient
	HyperVHost          *hypervhost.HyperVHostClient
	HyperVJobs          *hypervjobs.HyperVJobsClient
	HyperVMachines      *hypervmachines.HyperVMachinesClient
	HyperVRunAsAccounts *hypervrunasaccounts.HyperVRunAsAccountsClient
	HyperVSites         *hypervsites.HyperVSitesClient
	Jobs                *jobs.JobsClient
	Machines            *machines.MachinesClient
	RunAsAccounts       *runasaccounts.RunAsAccountsClient
	Sites               *sites.SitesClient
	VCenter             *vcenter.VCenterClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	hyperVClusterClient, err := hypervcluster.NewHyperVClusterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HyperVCluster client: %+v", err)
	}
	configureFunc(hyperVClusterClient.Client)

	hyperVHostClient, err := hypervhost.NewHyperVHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HyperVHost client: %+v", err)
	}
	configureFunc(hyperVHostClient.Client)

	hyperVJobsClient, err := hypervjobs.NewHyperVJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HyperVJobs client: %+v", err)
	}
	configureFunc(hyperVJobsClient.Client)

	hyperVMachinesClient, err := hypervmachines.NewHyperVMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HyperVMachines client: %+v", err)
	}
	configureFunc(hyperVMachinesClient.Client)

	hyperVRunAsAccountsClient, err := hypervrunasaccounts.NewHyperVRunAsAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HyperVRunAsAccounts client: %+v", err)
	}
	configureFunc(hyperVRunAsAccountsClient.Client)

	hyperVSitesClient, err := hypervsites.NewHyperVSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HyperVSites client: %+v", err)
	}
	configureFunc(hyperVSitesClient.Client)

	jobsClient, err := jobs.NewJobsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Jobs client: %+v", err)
	}
	configureFunc(jobsClient.Client)

	machinesClient, err := machines.NewMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Machines client: %+v", err)
	}
	configureFunc(machinesClient.Client)

	runAsAccountsClient, err := runasaccounts.NewRunAsAccountsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RunAsAccounts client: %+v", err)
	}
	configureFunc(runAsAccountsClient.Client)

	sitesClient, err := sites.NewSitesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Sites client: %+v", err)
	}
	configureFunc(sitesClient.Client)

	vCenterClient, err := vcenter.NewVCenterClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VCenter client: %+v", err)
	}
	configureFunc(vCenterClient.Client)

	return &Client{
		HyperVCluster:       hyperVClusterClient,
		HyperVHost:          hyperVHostClient,
		HyperVJobs:          hyperVJobsClient,
		HyperVMachines:      hyperVMachinesClient,
		HyperVRunAsAccounts: hyperVRunAsAccountsClient,
		HyperVSites:         hyperVSitesClient,
		Jobs:                jobsClient,
		Machines:            machinesClient,
		RunAsAccounts:       runAsAccountsClient,
		Sites:               sitesClient,
		VCenter:             vCenterClient,
	}, nil
}
