package balances

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2019-10-01"

func userAgent() string {
	return fmt.Sprintf("tombuildsstuff/go-azure-sdk/balances/%s", defaultApiVersion)
}
