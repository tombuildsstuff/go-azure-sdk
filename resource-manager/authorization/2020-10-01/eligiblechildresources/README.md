
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/authorization/2020-10-01/eligiblechildresources` Documentation

The `eligiblechildresources` SDK allows for interaction with the Azure Resource Manager Service `authorization` (API Version `2020-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/authorization/2020-10-01/eligiblechildresources"
```


### Client Initialization

```go
client := eligiblechildresources.NewEligibleChildResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EligibleChildResourcesClient.Get`

```go
ctx := context.TODO()
id := eligiblechildresources.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

// alternatively `client.Get(ctx, id, eligiblechildresources.DefaultGetOperationOptions())` can be used to do batched pagination
items, err := client.GetComplete(ctx, id, eligiblechildresources.DefaultGetOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
