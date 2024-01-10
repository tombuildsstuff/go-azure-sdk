
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/kusto` Documentation

The `kusto` SDK allows for interaction with the Azure Resource Manager Service `kusto` (API Version `2023-05-02`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/kusto/2023-05-02/kusto"
```


### Client Initialization

```go
client := kusto.NewKustoClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `KustoClient.ClustersListSkus`

```go
ctx := context.TODO()
id := kusto.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ClustersListSkus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `KustoClient.SkusList`

```go
ctx := context.TODO()
id := kusto.NewLocationID("12345678-1234-9876-4563-123456789012", "locationValue")

read, err := client.SkusList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
