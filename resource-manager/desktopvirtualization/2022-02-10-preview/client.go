package v2022_02_10_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/application"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/applicationgroup"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/desktop"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/hostpool"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/msiximage"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/msixpackage"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/privatelink"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/scalingplan"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/sessionhost"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/startmenuitem"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/usersession"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-02-10-preview/workspace"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Application      *application.ApplicationClient
	ApplicationGroup *applicationgroup.ApplicationGroupClient
	Desktop          *desktop.DesktopClient
	HostPool         *hostpool.HostPoolClient
	MSIXPackage      *msixpackage.MSIXPackageClient
	MsixImage        *msiximage.MsixImageClient
	PrivateLink      *privatelink.PrivateLinkClient
	ScalingPlan      *scalingplan.ScalingPlanClient
	SessionHost      *sessionhost.SessionHostClient
	StartMenuItem    *startmenuitem.StartMenuItemClient
	UserSession      *usersession.UserSessionClient
	Workspace        *workspace.WorkspaceClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	applicationClient, err := application.NewApplicationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Application client: %+v", err)
	}
	configureFunc(applicationClient.Client)

	applicationGroupClient, err := applicationgroup.NewApplicationGroupClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ApplicationGroup client: %+v", err)
	}
	configureFunc(applicationGroupClient.Client)

	desktopClient, err := desktop.NewDesktopClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Desktop client: %+v", err)
	}
	configureFunc(desktopClient.Client)

	hostPoolClient, err := hostpool.NewHostPoolClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HostPool client: %+v", err)
	}
	configureFunc(hostPoolClient.Client)

	mSIXPackageClient, err := msixpackage.NewMSIXPackageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MSIXPackage client: %+v", err)
	}
	configureFunc(mSIXPackageClient.Client)

	msixImageClient, err := msiximage.NewMsixImageClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MsixImage client: %+v", err)
	}
	configureFunc(msixImageClient.Client)

	privateLinkClient, err := privatelink.NewPrivateLinkClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLink client: %+v", err)
	}
	configureFunc(privateLinkClient.Client)

	scalingPlanClient, err := scalingplan.NewScalingPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlan client: %+v", err)
	}
	configureFunc(scalingPlanClient.Client)

	sessionHostClient, err := sessionhost.NewSessionHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SessionHost client: %+v", err)
	}
	configureFunc(sessionHostClient.Client)

	startMenuItemClient, err := startmenuitem.NewStartMenuItemClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StartMenuItem client: %+v", err)
	}
	configureFunc(startMenuItemClient.Client)

	userSessionClient, err := usersession.NewUserSessionClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building UserSession client: %+v", err)
	}
	configureFunc(userSessionClient.Client)

	workspaceClient, err := workspace.NewWorkspaceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspace client: %+v", err)
	}
	configureFunc(workspaceClient.Client)

	return &Client{
		Application:      applicationClient,
		ApplicationGroup: applicationGroupClient,
		Desktop:          desktopClient,
		HostPool:         hostPoolClient,
		MSIXPackage:      mSIXPackageClient,
		MsixImage:        msixImageClient,
		PrivateLink:      privateLinkClient,
		ScalingPlan:      scalingPlanClient,
		SessionHost:      sessionHostClient,
		StartMenuItem:    startMenuItemClient,
		UserSession:      userSessionClient,
		Workspace:        workspaceClient,
	}, nil
}
