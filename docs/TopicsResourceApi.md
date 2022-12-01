# \TopicsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisTopicV1NamespacesNamespaceTopicsGet**](TopicsResourceApi.md#ApisTopicV1NamespacesNamespaceTopicsGet) | **Get** /apis/topic/v1/namespaces/{namespace}/topics | 
[**ApisTopicV1NamespacesNamespaceTopicsPost**](TopicsResourceApi.md#ApisTopicV1NamespacesNamespaceTopicsPost) | **Post** /apis/topic/v1/namespaces/{namespace}/topics | 



## ApisTopicV1NamespacesNamespaceTopicsGet

> []string ApisTopicV1NamespacesNamespaceTopicsGet(ctx, namespace).Execute()



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
    resp, r, err := apiClient.TopicsResourceApi.ApisTopicV1NamespacesNamespaceTopicsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisTopicV1NamespacesNamespaceTopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisTopicV1NamespacesNamespaceTopicsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisTopicV1NamespacesNamespaceTopicsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisTopicV1NamespacesNamespaceTopicsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisTopicV1NamespacesNamespaceTopicsPost

> TopicData ApisTopicV1NamespacesNamespaceTopicsPost(ctx, namespace).CreateTopicRequestData(createTopicRequestData).Execute()



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
    createTopicRequestData := *openapiclient.NewCreateTopicRequestData() // CreateTopicRequestData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisTopicV1NamespacesNamespaceTopicsPost(context.Background(), namespace).CreateTopicRequestData(createTopicRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisTopicV1NamespacesNamespaceTopicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisTopicV1NamespacesNamespaceTopicsPost`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisTopicV1NamespacesNamespaceTopicsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisTopicV1NamespacesNamespaceTopicsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createTopicRequestData** | [**CreateTopicRequestData**](CreateTopicRequestData.md) |  | 

### Return type

[**TopicData**](TopicData.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

