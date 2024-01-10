
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/endpointcertificates` Documentation

The `endpointcertificates` SDK allows for interaction with the Azure Resource Manager Service `sql` (API Version `2021-11-01`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/sql/2021-11-01/endpointcertificates"
```


### Client Initialization

```go
client := endpointcertificates.NewEndpointCertificatesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `EndpointCertificatesClient.Get`

```go
ctx := context.TODO()
id := endpointcertificates.NewEndpointCertificateID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue", "endpointCertificateValue")

read, err := client.Get(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `EndpointCertificatesClient.ListByInstance`

```go
ctx := context.TODO()
id := endpointcertificates.NewSqlManagedInstanceID("12345678-1234-9876-4563-123456789012", "example-resource-group", "managedInstanceValue")

// alternatively `client.ListByInstance(ctx, id)` can be used to do batched pagination
items, err := client.ListByInstanceComplete(ctx, id)
if err != nil {
	// handle the error
}
for _, item := range items {
	// do something
}
```
