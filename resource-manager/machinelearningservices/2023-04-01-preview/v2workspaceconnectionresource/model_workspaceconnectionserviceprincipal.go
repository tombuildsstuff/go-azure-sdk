package v2workspaceconnectionresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceConnectionServicePrincipal struct {
	ClientId     *string `json:"clientId,omitempty"`
	ClientSecret *string `json:"clientSecret,omitempty"`
	TenantId     *string `json:"tenantId,omitempty"`
}
