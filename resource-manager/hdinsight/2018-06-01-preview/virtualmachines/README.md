
## `github.com/tombuildsstuff/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/virtualmachines` Documentation

The `virtualmachines` SDK allows for interaction with the Azure Resource Manager Service `hdinsight` (API Version `2018-06-01-preview`).

This readme covers example usages, but further information on [using this SDK can be found in the project root](https://github.com/tombuildsstuff/go-azure-sdk/tree/main/docs).

### Import Path

```go
import "github.com/tombuildsstuff/go-azure-sdk/resource-manager/hdinsight/2018-06-01-preview/virtualmachines"
```


### Client Initialization

```go
client := virtualmachines.NewVirtualMachinesClientWithBaseURI("https://management.azure.com")
client.Client.Authorizer = authorizer
```


### Example Usage: `VirtualMachinesClient.ListHosts`

```go
ctx := context.TODO()
id := virtualmachines.NewHDInsightClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")

read, err := client.ListHosts(ctx, id)
if err != nil {
	// handle the error
}
if model := read.Model; model != nil {
	// do something with the model/response object
}
```


### Example Usage: `VirtualMachinesClient.RestartHosts`

```go
ctx := context.TODO()
id := virtualmachines.NewHDInsightClusterID("12345678-1234-9876-4563-123456789012", "example-resource-group", "clusterValue")
var payload []string

if err := client.RestartHostsThenPoll(ctx, id, payload); err != nil {
	// handle the error
}
```
