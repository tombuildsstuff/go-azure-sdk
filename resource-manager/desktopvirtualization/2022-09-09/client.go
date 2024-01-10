package v2022_09_09

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/application"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/applicationgroup"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/desktop"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/hostpool"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/msiximage"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/msixpackage"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/scalingplan"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/scalingplanpooledschedule"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/sessionhost"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/startmenuitem"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/usersession"
	"github.com/tombuildsstuff/go-azure-sdk/resource-manager/desktopvirtualization/2022-09-09/workspace"
	"github.com/tombuildsstuff/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/tombuildsstuff/go-azure-sdk/sdk/environments"
)

type Client struct {
	Application               *application.ApplicationClient
	ApplicationGroup          *applicationgroup.ApplicationGroupClient
	Desktop                   *desktop.DesktopClient
	HostPool                  *hostpool.HostPoolClient
	MSIXPackage               *msixpackage.MSIXPackageClient
	MsixImage                 *msiximage.MsixImageClient
	ScalingPlan               *scalingplan.ScalingPlanClient
	ScalingPlanPooledSchedule *scalingplanpooledschedule.ScalingPlanPooledScheduleClient
	SessionHost               *sessionhost.SessionHostClient
	StartMenuItem             *startmenuitem.StartMenuItemClient
	UserSession               *usersession.UserSessionClient
	Workspace                 *workspace.WorkspaceClient
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

	scalingPlanClient, err := scalingplan.NewScalingPlanClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlan client: %+v", err)
	}
	configureFunc(scalingPlanClient.Client)

	scalingPlanPooledScheduleClient, err := scalingplanpooledschedule.NewScalingPlanPooledScheduleClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ScalingPlanPooledSchedule client: %+v", err)
	}
	configureFunc(scalingPlanPooledScheduleClient.Client)

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
		Application:               applicationClient,
		ApplicationGroup:          applicationGroupClient,
		Desktop:                   desktopClient,
		HostPool:                  hostPoolClient,
		MSIXPackage:               mSIXPackageClient,
		MsixImage:                 msixImageClient,
		ScalingPlan:               scalingPlanClient,
		ScalingPlanPooledSchedule: scalingPlanPooledScheduleClient,
		SessionHost:               sessionHostClient,
		StartMenuItem:             startMenuItemClient,
		UserSession:               userSessionClient,
		Workspace:                 workspaceClient,
	}, nil
}
