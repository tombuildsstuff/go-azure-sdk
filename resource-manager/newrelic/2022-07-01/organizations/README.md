
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/newrelic/2022-07-01/organizations` Documentation

The `organizations` SDK allows for interaction with the Azure Resource Manager Service `newrelic` (API Version `2022-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/newrelic/2022-07-01/organizations"
```


### Client Initialization

```go
client := organizations.NewOrganizationsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OrganizationsClient.List`

```go
ctx := context.TODO()
id := organizations.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id, organizations.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, organizations.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
