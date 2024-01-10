
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/post` Documentation

The `post` SDK allows for interaction with the Azure Resource Manager Service `postgresql` (API Version `2023-03-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/postgresql/2023-03-01-preview/post"
```


### Client Initialization

```go
client := post.NewPOSTClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `POSTClient.CheckMigrationNameAvailability`

```go
ctx := context.TODO()
id := post.NewFlexibleServerID("12345678-1234-9876-4563-123456789012", "example-resource-group", "flexibleServerValue")

payload := post.MigrationNameAvailabilityResource{
	// ...
}


read, err := client.CheckMigrationNameAvailability(ctx, id, payload)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```
