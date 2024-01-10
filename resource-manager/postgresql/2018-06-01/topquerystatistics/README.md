
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/postgresql/2018-06-01/topquerystatistics` Documentation

The `topquerystatistics` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/postgresql/2018-06-01/topquerystatistics"
```


### Client Initialization

```go
client := topquerystatistics.NewTopQueryStatisticsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TopQueryStatisticsClient.Get`

```go
ctx := context.TODO()
id := topquerystatistics.NewTopQueryStatisticID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "queryStatisticIdValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `TopQueryStatisticsClient.ListByServer`

```go
ctx := context.TODO()
id := topquerystatistics.NewServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
