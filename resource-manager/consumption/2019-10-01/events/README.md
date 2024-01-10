
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2019-10-01/events` Documentation

The `events` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2019-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2019-10-01/events"
```


### Client Initialization

```go
client := events.NewEventsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EventsClient.List`

```go
ctx := context.TODO()
id := events.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.List(ctx, id, events.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, events.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
