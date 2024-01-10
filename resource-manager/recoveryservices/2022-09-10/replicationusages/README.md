
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/replicationusages` Documentation

The `replicationusages` SDK allows for interaction with the Azure Resource Manager Service `recoveryservices` (API Version `2022-09-10`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/recoveryservices/2022-09-10/replicationusages"
```


### Client Initialization

```go
client := replicationusages.NewReplicationUsagesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ReplicationUsagesClient.List`

```go
ctx := context.TODO()
id := replicationusages.NewVaultID("12345678-1234-9876-4563-123456789012", "example-resource-group", "vaultValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
