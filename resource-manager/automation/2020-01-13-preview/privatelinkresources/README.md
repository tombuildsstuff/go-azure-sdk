
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/automation/2020-01-13-preview/privatelinkresources` Documentation

The `privatelinkresources` SDK allows for interaction with the Azure Resource Manager Service `automation` (API Version `2020-01-13-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/automation/2020-01-13-preview/privatelinkresources"
```


### Client Initialization

```go
client := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkResourcesClient.Automation`

```go
ctx := context.TODO()
id := privatelinkresources.NewAutomationAccountID("12345678-1234-9876-4563-123456789012", "example-resource-group", "automationAccountValue")

read, err := client.Automation(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
