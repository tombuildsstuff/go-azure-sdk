
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/enrichment` Documentation

The `enrichment` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2021-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/enrichment"
```


### Client Initialization

```go
client := enrichment.NewEnrichmentClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EnrichmentClient.DomainWhoisGet`

```go
ctx := context.TODO()
id := enrichment.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.DomainWhoisGet(ctx, id, enrichment.DefaultDomainWhoisGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EnrichmentClient.IPGeodataGet`

```go
ctx := context.TODO()
id := enrichment.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.IPGeodataGet(ctx, id, enrichment.DefaultIPGeodataGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
