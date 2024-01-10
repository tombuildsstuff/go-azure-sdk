
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/privatelink` Documentation

The `privatelink` SDK allows for interaction with the Azure Resource Manager Service `notificationhubs` (API Version `2023-09-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/notificationhubs/2023-09-01/privatelink"
```


### Client Initialization

```go
client := privatelink.NewPrivateLinkClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsDelete`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "privateEndpointConnectionValue")

if err := client.PrivateEndpointConnectionsDeleteThenPoll(ctx, id); err != nil {
	// handle the error
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsGet`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "privateEndpointConnectionValue")

read, err := client.PrivateEndpointConnectionsGet(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsGetGroupId`

```go
ctx := context.TODO()
id := privatelink.NewPrivateLinkResourceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "privateLinkResourceValue")

read, err := client.PrivateEndpointConnectionsGetGroupId(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsList`

```go
ctx := context.TODO()
id := privatelink.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.PrivateEndpointConnectionsList(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsListGroupIds`

```go
ctx := context.TODO()
id := privatelink.NewNamespaceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue")

read, err := client.PrivateEndpointConnectionsListGroupIds(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `PrivateLinkClient.PrivateEndpointConnectionsUpdate`

```go
ctx := context.TODO()
id := privatelink.NewPrivateEndpointConnectionID("12345678-1234-9876-4563-123456789012", "example-resource-group", "namespaceValue", "privateEndpointConnectionValue")

payload := privatelink.PrivateEndpointConnectionResource{
	// ...
}


if err := client.PrivateEndpointConnectionsUpdateThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
