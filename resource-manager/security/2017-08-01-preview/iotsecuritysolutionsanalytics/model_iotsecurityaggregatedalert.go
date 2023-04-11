package iotsecuritysolutionsanalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IoTSecurityAggregatedAlert struct {
	Id         *string                               `json:"id,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	Properties *IoTSecurityAggregatedAlertProperties `json:"properties,omitempty"`
	Tags       *map[string]string                    `json:"tags,omitempty"`
	Type       *string                               `json:"type,omitempty"`
}
