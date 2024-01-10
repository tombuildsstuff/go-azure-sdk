
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/locationcapabilities` Documentation

The `locationcapabilities` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/locationcapabilities"
```


### Client Initialization

```go
client := locationcapabilities.NewLocationCapabilitiesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `LocationCapabilitiesClient.CapabilitiesListByLocation`

```go
ctx := context.TODO()
id := locationcapabilities.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.CapabilitiesListByLocation(ctx, id, locationcapabilities.DefaultCapabilitiesListByLocationOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
