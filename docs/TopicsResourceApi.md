# \TopicsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisTopicV1NamespacesNameTopicsGet**](TopicsResourceApi.md#ApisTopicV1NamespacesNameTopicsGet) | **Get** /apis/topic/v1/namespaces/{name}/topics | 
[**ApisTopicV1NamespacesNameTopicsPost**](TopicsResourceApi.md#ApisTopicV1NamespacesNameTopicsPost) | **Post** /apis/topic/v1/namespaces/{name}/topics | 



## ApisTopicV1NamespacesNameTopicsGet

> []string ApisTopicV1NamespacesNameTopicsGet(ctx, name).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisTopicV1NamespacesNameTopicsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisTopicV1NamespacesNameTopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisTopicV1NamespacesNameTopicsGet`: []string
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisTopicV1NamespacesNameTopicsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisTopicV1NamespacesNameTopicsGetRequest struct via the builder pattern


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


## ApisTopicV1NamespacesNameTopicsPost

> TopicData ApisTopicV1NamespacesNameTopicsPost(ctx, name).CreateTopicRequestData(createTopicRequestData).Execute()



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
    createTopicRequestData := *openapiclient.NewCreateTopicRequestData() // CreateTopicRequestData |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TopicsResourceApi.ApisTopicV1NamespacesNameTopicsPost(context.Background(), name).CreateTopicRequestData(createTopicRequestData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TopicsResourceApi.ApisTopicV1NamespacesNameTopicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisTopicV1NamespacesNameTopicsPost`: TopicData
    fmt.Fprintf(os.Stdout, "Response from `TopicsResourceApi.ApisTopicV1NamespacesNameTopicsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisTopicV1NamespacesNameTopicsPostRequest struct via the builder pattern


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

