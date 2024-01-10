
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2021-05-01-preview/managementgroupdiagnosticsettings` Documentation

The `managementgroupdiagnosticsettings` SDK allows for interaction with the Azure Resource Manager Service `insights` (API Version `2021-05-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/insights/2021-05-01-preview/managementgroupdiagnosticsettings"
```


### Client Initialization

```go
client := managementgroupdiagnosticsettings.NewManagementGroupDiagnosticSettingsClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `ManagementGroupDiagnosticSettingsClient.CreateOrUpdate`

```go
ctx := context.TODO()
id := managementgroupdiagnosticsettings.NewProviders2DiagnosticSettingID("managementGroupIdValue", "diagnosticSettingValue")

payload := managementgroupdiagnosticsettings.ManagementGroupDiagnosticSettingsResource{
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


### Example Usage: `ManagementGroupDiagnosticSettingsClient.Delete`

```go
ctx := context.TODO()
id := managementgroupdiagnosticsettings.NewProviders2DiagnosticSettingID("managementGroupIdValue", "diagnosticSettingValue")

read, err := client.Delete(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupDiagnosticSettingsClient.Get`

```go
ctx := context.TODO()
id := managementgroupdiagnosticsettings.NewProviders2DiagnosticSettingID("managementGroupIdValue", "diagnosticSettingValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `ManagementGroupDiagnosticSettingsClient.List`

```go
ctx := context.TODO()
id := managementgroupdiagnosticsettings.NewManagementGroupID("groupIdValue")

read, err := client.List(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
