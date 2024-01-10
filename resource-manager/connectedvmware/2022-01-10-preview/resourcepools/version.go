package resourcepools

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2022-01-10-preview"

func userAgent() string {
	return fmt.Sprintf("tombuildsstuff/go-azure-sdk/resourcepools/%s", defaultApiVersion)
}
