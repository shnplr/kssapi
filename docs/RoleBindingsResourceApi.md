# \RoleBindingsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisRbacAuthorizationV1ClusterrolebindingsGet**](RoleBindingsResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsGet) | **Get** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsPost**](RoleBindingsResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsPost) | **Post** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1NamespacesNameRolebindingsGet**](RoleBindingsResourceApi.md#ApisRbacAuthorizationV1NamespacesNameRolebindingsGet) | **Get** /apis/rbac.authorization/v1/namespaces/{name}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost**](RoleBindingsResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost) | **Post** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
[**ApisRbacAuthorizationV1RolebindingsGet**](RoleBindingsResourceApi.md#ApisRbacAuthorizationV1RolebindingsGet) | **Get** /apis/rbac.authorization/v1/rolebindings | 



## ApisRbacAuthorizationV1ClusterrolebindingsGet

> []Fact ApisRbacAuthorizationV1ClusterrolebindingsGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsGet`: []Fact
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsGetRequest struct via the builder pattern


### Return type

[**[]Fact**](Fact.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolebindingsPost

> ApiStatus ApisRbacAuthorizationV1ClusterrolebindingsPost(ctx).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding() // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost(context.Background()).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsPost`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleBinding** | [**RoleBinding**](RoleBinding.md) |  | 

### Return type

[**ApiStatus**](ApiStatus.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1NamespacesNameRolebindingsGet

> ApiStatus ApisRbacAuthorizationV1NamespacesNameRolebindingsGet(ctx, name).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNameRolebindingsGet`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNameRolebindingsGetRequest struct via the builder pattern


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


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost

> ApiStatus ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost(ctx, namespace).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding() // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost(context.Background(), namespace).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **roleBinding** | [**RoleBinding**](RoleBinding.md) |  | 

### Return type

[**ApiStatus**](ApiStatus.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1RolebindingsGet

> ApiStatus ApisRbacAuthorizationV1RolebindingsGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacAuthorizationV1RolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacAuthorizationV1RolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1RolebindingsGet`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisRbacAuthorizationV1RolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1RolebindingsGetRequest struct via the builder pattern


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

