
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/costdetails` Documentation

The `costdetails` SDK allows for interaction with the Azure Resource Manager Service `costmanagement` (API Version `2022-10-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/costmanagement/2022-10-01/costdetails"
```


### Client Initialization

```go
client := costdetails.NewCostDetailsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CostDetailsClient.GenerateCostDetailsReportCreateOperation`

```go
ctx := context.TODO()
id := costdetails.NewScopeID("/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/some-resource-group")

payload := costdetails.GenerateCostDetailsReportRequestDefinition{
	// ...
}


if err := client.GenerateCostDetailsReportCreateOperationThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
