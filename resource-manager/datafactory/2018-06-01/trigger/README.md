
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/datafactory/2018-06-01/trigger` Documentation

The `trigger` SDK allows for interaction with the Azure Resource Manager Service `datafactory` (API Version `2018-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/datafactory/2018-06-01/trigger"
```


### Client Initialization

```go
client := trigger.NewTriggerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `TriggerClient.Get`

```go
ctx := context.TODO()
id := trigger.NewTriggerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "factoryValue", "triggerValue")

read, err := client.Get(ctx, id, trigger.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
