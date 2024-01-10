
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/security/2017-08-01-preview/workspacesettings` Documentation

The `workspacesettings` SDK allows for interaction with the Azure Resource Manager Service `security` (API Version `2017-08-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/security/2017-08-01-preview/workspacesettings"
```


### Client Initialization

```go
client := workspacesettings.NewWorkspaceSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `WorkspaceSettingsClient.Create`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

payload := workspacesettings.WorkspaceSetting{
	// ...
}


read, err := client.Create(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceSettingsClient.Delete`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceSettingsClient.Get`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `WorkspaceSettingsClient.List`

```go
ctx := context.TODO()
id := workspacesettings.NewSubscriptionID("12345678-1234-9876-4563-123456789012")

// alternatively `client.List(ctx, id)` can be used to do batched pagination
items, err := client.ListComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```


### Example Usage: `WorkspaceSettingsClient.Update`

```go
ctx := context.TODO()
id := workspacesettings.NewWorkspaceSettingID("12345678-1234-9876-4563-123456789012", "workspaceSettingValue")

payload := workspacesettings.WorkspaceSetting{
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
