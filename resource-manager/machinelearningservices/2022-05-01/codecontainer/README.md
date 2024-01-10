
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/codecontainer` Documentation

The `codecontainer` SDK allows for interaction with the Azure Resource Manager Service `machinelearningservices` (API Version `2022-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/machinelearningservices/2022-05-01/codecontainer"
```


### Client Initialization

```go
client := codecontainer.NewCodeContainerClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `CodeContainerClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := codecontainer.NewCodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "codeValue")

payload := codecontainer.CodeContainerResource{
	// ...
}


read, err := client.CreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CodeContainerClient.Delete`

```go
ctx := context.TODO()
id := codecontainer.NewCodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "codeValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CodeContainerClient.Get`

```go
ctx := context.TODO()
id := codecontainer.NewCodeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "codeValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `CodeContainerClient.List`

```go
ctx := context.TODO()
id := codecontainer.NewWorkspaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue")

// alternatively `client.List(ctx, id, codecontainer.DefaultListOperationOptions())` can be used to do batched pagination
items, err := client.ListComplete(ctx, id, codecontainer.DefaultListOperationOptions())
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
