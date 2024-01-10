
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/dataconnectorsconnect` Documentation

The `dataconnectorsconnect` SDK allows for interaction with the Azure Resource Manager Service `securityinsights` (API Version `2021-09-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/securityinsights/2021-09-01-preview/dataconnectorsconnect"
```


### Client Initialization

```go
client := dataconnectorsconnect.NewDataConnectorsConnectClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `DataConnectorsConnectClient.DataConnectorsConnect`

```go
ctx := context.TODO()
id := dataconnectorsconnect.NewDataConnectorID("12345678-1234-9876-4563-123456789012", "example-resource-group", "workspaceValue", "dataConnectorIdValue")

payload := dataconnectorsconnect.DataConnectorConnectBody{
	// ...
}


read, err := client.DataConnectorsConnect(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
