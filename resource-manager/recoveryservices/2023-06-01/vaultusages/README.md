
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/recoveryservices/2023-06-01/vaultusages` Documentation

The `vaultusages` SDK allows for interaction with the Azure Resource Manager Service `recoveryservices` (API Version `2023-06-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/recoveryservices/2023-06-01/vaultusages"
```


### Client Initialization

```go
client := vaultusages.NewVaultUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VaultUsagesClient.UsagesListByVaults`

```go
ctx := context.TODO()
id := vaultusages.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

read, err := client.UsagesListByVaults(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
