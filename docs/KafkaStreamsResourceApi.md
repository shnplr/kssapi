# \KafkaStreamsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet**](KafkaStreamsResourceApi.md#ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet) | **Get** /apis/kafka.streams/v1/namespaces/{namespace}/streams/{name} | 
[**ApisKafkaStreamsV1NamespacesNamespaceStreamsPost**](KafkaStreamsResourceApi.md#ApisKafkaStreamsV1NamespacesNamespaceStreamsPost) | **Post** /apis/kafka.streams/v1/namespaces/{namespace}/streams | 



## ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet

> KafkaStreams ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaStreamsResourceApi.ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaStreamsResourceApi.ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet`: KafkaStreams
    fmt.Fprintf(os.Stdout, "Response from `KafkaStreamsResourceApi.ApisKafkaStreamsV1NamespacesNamespaceStreamsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaStreamsV1NamespacesNamespaceStreamsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaStreams**](KafkaStreams.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaStreamsV1NamespacesNamespaceStreamsPost

> KafkaStreams ApisKafkaStreamsV1NamespacesNamespaceStreamsPost(ctx, namespace).KafkaStreams(kafkaStreams).Execute()



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
    kafkaStreams := *openapiclient.NewKafkaStreams("Name_example", "Principal_example") // KafkaStreams |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaStreamsResourceApi.ApisKafkaStreamsV1NamespacesNamespaceStreamsPost(context.Background(), namespace).KafkaStreams(kafkaStreams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaStreamsResourceApi.ApisKafkaStreamsV1NamespacesNamespaceStreamsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaStreamsV1NamespacesNamespaceStreamsPost`: KafkaStreams
    fmt.Fprintf(os.Stdout, "Response from `KafkaStreamsResourceApi.ApisKafkaStreamsV1NamespacesNamespaceStreamsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaStreamsV1NamespacesNamespaceStreamsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaStreams** | [**KafkaStreams**](KafkaStreams.md) |  | 

### Return type

[**KafkaStreams**](KafkaStreams.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

