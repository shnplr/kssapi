# \TopicsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaTopicV1NamespacesNamespaceTopicsGet**](TopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost**](TopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/config | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete**](TopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete) | **Delete** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet**](TopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/describe | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet**](TopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsPost**](TopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 



## ApisKafkaTopicV1NamespacesNamespaceTopicsGet

> GenericListKafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsGet(ctx, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsGet`: GenericListKafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet`: %v\n", resp)
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

[**GenericListKafkaTopic**](GenericListKafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost(ctx, name, namespace).KafkaTopic(kafkaTopic).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    kafkaTopic := *openapiclient.NewKafkaTopic("Name_example", "Namespace_example") // KafkaTopic |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost(context.Background(), name, namespace).KafkaTopic(kafkaTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kafkaTopic** | [**KafkaTopic**](KafkaTopic.md) |  | 

### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete

> ApiStatus ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(ctx, name, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete`: %v\n", resp)
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

[**ApiStatus**](ApiStatus.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet(ctx, name, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(ctx, name, namespace).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet`: %v\n", resp)
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



### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsPost

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsPost(ctx, namespace).KafkaTopic(kafkaTopic).Execute()



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    namespace := "namespace_example" // string | 
    kafkaTopic := *openapiclient.NewKafkaTopic("Name_example", "Namespace_example") // KafkaTopic |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost(context.Background(), namespace).KafkaTopic(kafkaTopic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsPost`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost`: %v\n", resp)
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

 **kafkaTopic** | [**KafkaTopic**](KafkaTopic.md) |  | 

### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

