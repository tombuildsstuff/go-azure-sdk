
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/databoxedge/2023-07-01/nodes` Documentation

The `nodes` SDK allows for interaction with the Azure Resource Manager Service `databoxedge` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/databoxedge/2023-07-01/nodes"
```


### Client Initialization

```go
client := nodes.NewNodesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `NodesClient.ListByDataBoxEdgeDevice`

```go
ctx := context.TODO()
id := nodes.NewDataBoxEdgeDeviceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "dataBoxEdgeDeviceValue")

// alternatively `client.ListByDataBoxEdgeDevice(ctx, id)` can be used to do batched pagination
items, err := client.ListByDataBoxEdgeDeviceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
