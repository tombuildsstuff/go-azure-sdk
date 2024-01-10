
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/myworkbooksapis` Documentation

The `myworkbooksapis` SDK allows for interaction with the Azure Resource Manager Service `applicationinsights` (API Version `2015-05-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/applicationinsights/2015-05-01/myworkbooksapis"
```


### Client Initialization

```go
client := myworkbooksapis.NewMyworkbooksAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `MyworkbooksAPIsClient.MyWorkbooksCreateOrUpdate`

```go
ctx := context.TODO()
id := myworkbooksapis.NewMyWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "myWorkbookValue")

payload := myworkbooksapis.MyWorkbook{
	// ...
}


read, err := client.MyWorkbooksCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MyworkbooksAPIsClient.MyWorkbooksDelete`

```go
ctx := context.TODO()
id := myworkbooksapis.NewMyWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "myWorkbookValue")

read, err := client.MyWorkbooksDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MyworkbooksAPIsClient.MyWorkbooksGet`

```go
ctx := context.TODO()
id := myworkbooksapis.NewMyWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "myWorkbookValue")

read, err := client.MyWorkbooksGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MyworkbooksAPIsClient.MyWorkbooksListByResourceGroup`

```go
ctx := context.TODO()
id := myworkbooksapis.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.MyWorkbooksListByResourceGroup(ctx, id, myworkbooksapis.DefaultMyWorkbooksListByResourceGroupOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MyworkbooksAPIsClient.MyWorkbooksListBySubscription`

```go
ctx := context.TODO()
id := myworkbooksapis.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.MyWorkbooksListBySubscription(ctx, id, myworkbooksapis.DefaultMyWorkbooksListBySubscriptionOperationOptions())
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `MyworkbooksAPIsClient.MyWorkbooksUpdate`

```go
ctx := context.TODO()
id := myworkbooksapis.NewMyWorkbookID("12345678-1234-9876-4563-123456789012", "example-resource-group", "myWorkbookValue")

payload := myworkbooksapis.MyWorkbook{
	// ...
}


read, err := client.MyWorkbooksUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
