
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-05-01/charges` Documentation

The `charges` SDK allows for interaction with the Azure Resource Manager Service `consumption` (API Version `2023-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/consumption/2023-05-01/charges"
```


### Client Initialization

```go
client := charges.NewChargesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ChargesClient.List`

```go
ctx := context.TODO()
id := charges.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

read, err := client.List(ctx, id, charges.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
