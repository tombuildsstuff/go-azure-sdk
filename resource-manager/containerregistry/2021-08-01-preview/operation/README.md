
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerregistry/2021-08-01-preview/operation` Documentation

The `operation` SDK allows for interaction with the Azure Resource Manager Service `containerregistry` (API Version `2021-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/containerregistry/2021-08-01-preview/operation"
```


### Client Initialization

```go
client := operation.NewOperationClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `OperationClient.RegistriesCheckNameAvailability`

```go
ctx := context.TODO()
id := operation.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

payload := operation.RegistryNameCheckRequest{
	// ...
}


read, err := client.RegistriesCheckNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
