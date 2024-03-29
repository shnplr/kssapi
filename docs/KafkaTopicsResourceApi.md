# \KafkaTopicsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaTopicV1NamespacesNamespaceTopicsGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/configs:alter | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete) | **Delete** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsPost**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/sync | 
[**ApisKafkaTopicV1TopicsGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1TopicsGet) | **Get** /apis/kafka.topic/v1/topics | 



## ApisKafkaTopicV1NamespacesNamespaceTopicsGet

> KafkaTopicDataList ApisKafkaTopicV1NamespacesNamespaceTopicsGet(ctx, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsGet`: KafkaTopicDataList
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaTopicDataList**](KafkaTopicDataList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost

> KafkaTopicData ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost(ctx, name, namespace).AlterConfigBatchRequest(alterConfigBatchRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    alterConfigBatchRequest := *openapiclient.NewAlterConfigBatchRequest() // AlterConfigBatchRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost(context.Background(), name, namespace).AlterConfigBatchRequest(alterConfigBatchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost`: KafkaTopicData
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **alterConfigBatchRequest** | [**AlterConfigBatchRequest**](AlterConfigBatchRequest.md) |  | 

### Return type

[**KafkaTopicData**](KafkaTopicData.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete

> Status ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(ctx, name, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete`: Status
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Status**](Status.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet

> KafkaTopicData ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(ctx, name, namespace).IncludeConfigs(includeConfigs).IncludePartitions(includePartitions).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    includeConfigs := true // bool |  (optional)
    includePartitions := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(context.Background(), name, namespace).IncludeConfigs(includeConfigs).IncludePartitions(includePartitions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet`: KafkaTopicData
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeConfigs** | **bool** |  | 
 **includePartitions** | **bool** |  | 

### Return type

[**KafkaTopicData**](KafkaTopicData.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsPost

> KafkaTopicData ApisKafkaTopicV1NamespacesNamespaceTopicsPost(ctx, namespace).CreateTopicRequest(createTopicRequest).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {
    namespace := "namespace_example" // string | 
    createTopicRequest := *openapiclient.NewCreateTopicRequest() // CreateTopicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost(context.Background(), namespace).CreateTopicRequest(createTopicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsPost`: KafkaTopicData
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTopicRequest** | [**CreateTopicRequest**](CreateTopicRequest.md) |  | 

### Return type

[**KafkaTopicData**](KafkaTopicData.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost

> Status ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost(ctx, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost`: Status
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsSyncPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1TopicsGet

> KafkaTopicDataList ApisKafkaTopicV1TopicsGet(ctx).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/shnplr/kssapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1TopicsGet`: KafkaTopicDataList
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1TopicsGetRequest struct via the builder pattern


### Return type

[**KafkaTopicDataList**](KafkaTopicDataList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

