package resolveprivatelinkserviceid

import "fmt"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

const defaultApiVersion = "2023-03-02-preview"

func userAgent() string {
	return fmt.Sprintf("tombuildsstuff/go-azure-sdk/resolveprivatelinkserviceid/%s", defaultApiVersion)
}
