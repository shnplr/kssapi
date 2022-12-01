# \KafkaRbacResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaRbacV1NamespacesNamespaceBindingsPost**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNamespaceBindingsPost) | **Post** /apis/kafka.rbac/v1/namespaces/{namespace}/bindings | 



## ApisKafkaRbacV1NamespacesNamespaceBindingsPost

> RbacRoleBindingResponse ApisKafkaRbacV1NamespacesNamespaceBindingsPost(ctx, namespace).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()



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
    rbacRoleBindingRequest := *openapiclient.NewRbacRoleBindingRequest() // RbacRoleBindingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceBindingsPost(context.Background(), namespace).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceBindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1NamespacesNamespaceBindingsPost`: RbacRoleBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceBindingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1NamespacesNamespaceBindingsPostRequest struct via the builder pattern


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

