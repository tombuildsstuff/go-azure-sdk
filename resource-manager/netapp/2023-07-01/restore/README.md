
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-07-01/restore` Documentation

The `restore` SDK allows for interaction with the Azure Resource Manager Service `netapp` (API Version `2023-07-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/netapp/2023-07-01/restore"
```


### Client Initialization

```go
client := restore.NewRestoreClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `RestoreClient.BackupsGetVolumeRestoreStatus`

```go
ctx := context.TODO()
id := restore.NewVolumeID("12345678-1234-9876-4563-123456789012", "example-resource-group", "netAppAccountValue", "capacityPoolValue", "volumeValue")

read, err := client.BackupsGetVolumeRestoreStatus(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
