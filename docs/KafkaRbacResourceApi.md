# \KafkaRbacResourceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaRbacV1NamespacesNameBindingsPost**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNameBindingsPost) | **Post** /apis/kafka.rbac/v1/namespaces/{name}/bindings | 



## ApisKafkaRbacV1NamespacesNameBindingsPost

> RbacRoleBindingResponse ApisKafkaRbacV1NamespacesNameBindingsPost(ctx, name).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()



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
    rbacRoleBindingRequest := *openapiclient.NewRbacRoleBindingRequest() // RbacRoleBindingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNameBindingsPost(context.Background(), name).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNameBindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1NamespacesNameBindingsPost`: RbacRoleBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNameBindingsPost`: %v\n", resp)
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

 **rbacRoleBindingRequest** | [**RbacRoleBindingRequest**](RbacRoleBindingRequest.md) |  | 

### Return type

[**RbacRoleBindingResponse**](RbacRoleBindingResponse.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

