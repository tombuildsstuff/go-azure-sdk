
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2023-10-01/metricdefinitions` Documentation

The `metricdefinitions` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2023-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2023-10-01/metricdefinitions"
```


### Client Initialization

```go
client := metricdefinitions.NewMetricDefinitionsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MetricDefinitionsClient.List`

```go
ctx := context.TODO()
id := metricdefinitions.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id, metricdefinitions.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MetricDefinitionsClient.ListAtSubscriptionScope`

```go
ctx := context.TODO()
id := metricdefinitions.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ListAtSubscriptionScope(ctx, id, metricdefinitions.DefaultListAtSubscriptionScopeOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
