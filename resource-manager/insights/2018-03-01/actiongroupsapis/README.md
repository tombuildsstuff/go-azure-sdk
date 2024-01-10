
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2018-03-01/actiongroupsapis` Documentation

The `actiongroupsapis` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2018-03-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2018-03-01/actiongroupsapis"
```


### Client Initialization

```go
client := actiongroupsapis.NewActionGroupsAPIsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsCreateOrUpdate`

```go
ctx := context.TODO()
id := actiongroupsapis.NewActionGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "actionGroupValue")

payload := actiongroupsapis.ActionGroupResource{
	// ...
}


read, err := client.ActionGroupsCreateOrUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsDelete`

```go
ctx := context.TODO()
id := actiongroupsapis.NewActionGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "actionGroupValue")

read, err := client.ActionGroupsDelete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsEnableReceiver`

```go
ctx := context.TODO()
id := actiongroupsapis.NewActionGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "actionGroupValue")

payload := actiongroupsapis.EnableRequest{
	// ...
}


read, err := client.ActionGroupsEnableReceiver(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsGet`

```go
ctx := context.TODO()
id := actiongroupsapis.NewActionGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "actionGroupValue")

read, err := client.ActionGroupsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsListByResourceGroup`

```go
ctx := context.TODO()
id := actiongroupsapis.NewResourceGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group")

read, err := client.ActionGroupsListByResourceGroup(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsListBySubscriptionId`

```go
ctx := context.TODO()
id := actiongroupsapis.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

read, err := client.ActionGroupsListBySubscriptionId(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ActionGroupsAPIsClient.ActionGroupsUpdate`

```go
ctx := context.TODO()
id := actiongroupsapis.NewActionGroupID("12345678-1234-9876-4563-123456789012", "example-resource-group", "actionGroupValue")

payload := actiongroupsapis.ActionGroupPatchBody{
	// ...
}


read, err := client.ActionGroupsUpdate(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
