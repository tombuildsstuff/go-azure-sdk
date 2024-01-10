
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2023-02-01-preview/serverautomatictuning` Documentation

The `serverautomatictuning` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2023-02-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2023-02-01-preview/serverautomatictuning"
```


### Client Initialization

```go
client := serverautomatictuning.NewServerAutomaticTuningClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ServerAutomaticTuningClient.Get`

```go
ctx := context.TODO()
id := serverautomatictuning.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ServerAutomaticTuningClient.Update`

```go
ctx := context.TODO()
id := serverautomatictuning.NewSqlServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "serverValue")

payload := serverautomatictuning.ServerAutomaticTuning{
	// ...
}


read, err := client.Update(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
