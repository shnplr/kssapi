# \KafkaRbacResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaRbacV1NamespacesNamespaceBindingsDelete**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNamespaceBindingsDelete) | **Delete** /apis/kafka.rbac/v1/namespaces/{namespace}/bindings | 
[**ApisKafkaRbacV1NamespacesNamespaceBindingsPost**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNamespaceBindingsPost) | **Post** /apis/kafka.rbac/v1/namespaces/{namespace}/bindings | 
[**ApisKafkaRbacV1NamespacesNamespaceResourcesGet**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNamespaceResourcesGet) | **Get** /apis/kafka.rbac/v1/namespaces/{namespace}/resources | 
[**ApisKafkaRbacV1ResourcesGet**](KafkaRbacResourceApi.md#ApisKafkaRbacV1ResourcesGet) | **Get** /apis/kafka.rbac/v1/resources | 



## ApisKafkaRbacV1NamespacesNamespaceBindingsDelete

> RbacRoleBindingResponse ApisKafkaRbacV1NamespacesNamespaceBindingsDelete(ctx, namespace).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()



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
    rbacRoleBindingRequest := *openapiclient.NewRbacRoleBindingRequest("Role_example", "User_example") // RbacRoleBindingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceBindingsDelete(context.Background(), namespace).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceBindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1NamespacesNamespaceBindingsDelete`: RbacRoleBindingResponse
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceBindingsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1NamespacesNamespaceBindingsDeleteRequest struct via the builder pattern


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


## ApisKafkaRbacV1NamespacesNamespaceBindingsPost

> RbacRoleBindingResponse ApisKafkaRbacV1NamespacesNamespaceBindingsPost(ctx, namespace).RbacRoleBindingRequest(rbacRoleBindingRequest).Execute()



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
    rbacRoleBindingRequest := *openapiclient.NewRbacRoleBindingRequest("Role_example", "User_example") // RbacRoleBindingRequest |  (optional)

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


## ApisKafkaRbacV1NamespacesNamespaceResourcesGet

> KafkaRbacSummaryList ApisKafkaRbacV1NamespacesNamespaceResourcesGet(ctx, namespace).Role(role).User(user).Execute()



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
    role := "role_example" // string |  (optional)
    user := "user_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceResourcesGet(context.Background(), namespace).Role(role).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceResourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1NamespacesNamespaceResourcesGet`: KafkaRbacSummaryList
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceResourcesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1NamespacesNamespaceResourcesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **role** | **string** |  | 
 **user** | **string** |  | 

### Return type

[**KafkaRbacSummaryList**](KafkaRbacSummaryList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaRbacV1ResourcesGet

> KafkaRbacSummaryList ApisKafkaRbacV1ResourcesGet(ctx).Role(role).User(user).Execute()



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
    role := "role_example" // string |  (optional)
    user := "user_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1ResourcesGet(context.Background()).Role(role).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1ResourcesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1ResourcesGet`: KafkaRbacSummaryList
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1ResourcesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1ResourcesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **role** | **string** |  | 
 **user** | **string** |  | 

### Return type

[**KafkaRbacSummaryList**](KafkaRbacSummaryList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

