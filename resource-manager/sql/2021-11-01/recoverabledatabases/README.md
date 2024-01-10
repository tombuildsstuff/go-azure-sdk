
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/recoverabledatabases` Documentation

The `recoverabledatabases` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/recoverabledatabases"
```


### Client Initialization

```go
client := recoverabledatabases.NewRecoverableDatabasesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RecoverableDatabasesClient.Get`

```go
ctx := context.TODO()
id := recoverabledatabases.NewRecoverableDatabaseID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue", "recoverableDatabaseValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `RecoverableDatabasesClient.ListByServer`

```go
ctx := context.TODO()
id := recoverabledatabases.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

// alternatively `client.ListByServer(ctx, id)` can be used to do batched pagination
items, err := client.ListByServerComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
