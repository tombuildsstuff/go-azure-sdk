
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/serverconnectionpolicies` Documentation

The `serverconnectionpolicies` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/serverconnectionpolicies"
```


### Client Initialization

```go
client := serverconnectionpolicies.NewServerConnectionPoliciesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerConnectionPoliciesClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := serverconnectionpolicies.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

payload := serverconnectionpolicies.ServerConnectionPolicy{
	// ...
}


if err := client.CreateOrUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```


### Example Usage: `ServerConnectionPoliciesClient.Get`

```go
ctx := context.TODO()
id := serverconnectionpolicies.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerConnectionPoliciesClient.ListByServer`

```go
ctx := context.TODO()
id := serverconnectionpolicies.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
