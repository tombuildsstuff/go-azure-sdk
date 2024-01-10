
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsan` Documentation

The `elasticsan` SDK allows for interaction with the Azure Resource Manager Service `elasticsan` (API Version `2021-11-20-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/elasticsan/2021-11-20-preview/elasticsan"
```


### Client Initialization

```go
client := elasticsan.NewElasticSanClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ElasticSanClient.ListByResourceGroup`

```go
ctx := context.TODO()
id := elasticsan.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

// alternatively `client.ListByResourceGroup(ctx, id)` can be used to do batched pagination
items, err := client.ListByResourceGroupComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
