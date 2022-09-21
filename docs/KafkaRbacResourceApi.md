# \KafkaRbacResourceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaRbacV1NamespacesNameBindingsPost**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNameBindingsPost) | **Post** /apis/kafka.rbac/v1/namespaces/{name}/bindings | 



## ApisKafkaRbacV1NamespacesNameBindingsPost

> ApisKafkaRbacV1NamespacesNameBindingsPost(ctx, name).KafkaRbacRequest(kafkaRbacRequest).Execute()



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
    kafkaRbacRequest := *openapiclient.NewKafkaRbacRequest() // KafkaRbacRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNameBindingsPost(context.Background(), name).KafkaRbacRequest(kafkaRbacRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNameBindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1NamespacesNameBindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaRbacRequest** | [**KafkaRbacRequest**](KafkaRbacRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

