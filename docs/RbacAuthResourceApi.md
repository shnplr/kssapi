# \RbacAuthResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisRbacAuthorizationV1ClusterrolebindingsDelete**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsDelete) | **Delete** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsGet) | **Get** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsPost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsPost) | **Post** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolesGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolesGet) | **Get** /apis/rbac.authorization/v1/clusterroles | 
[**ApisRbacAuthorizationV1ClusterrolesNameGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolesNameGet) | **Get** /apis/rbac.authorization/v1/clusterroles/{name} | 
[**ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete) | **Delete** /apis/rbac.authorization/v1/namespaces/{name}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNameRolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNameRolebindingsGet) | **Get** /apis/rbac.authorization/v1/namespaces/{name}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNameRolebindingsPost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNameRolebindingsPost) | **Post** /apis/rbac.authorization/v1/namespaces/{name}/rolebindings | 



## ApisRbacAuthorizationV1ClusterrolebindingsDelete

> ApiStatus ApisRbacAuthorizationV1ClusterrolebindingsDelete(ctx).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding("Role_example") // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsDelete(context.Background()).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsDelete`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsDeleteRequest struct via the builder pattern


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


## ApisRbacAuthorizationV1ClusterrolebindingsGet

> GenericListRoleBinding ApisRbacAuthorizationV1ClusterrolebindingsGet(ctx).Execute()



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
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsGet`: GenericListRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsGetRequest struct via the builder pattern


### Return type

[**GenericListRoleBinding**](GenericListRoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolebindingsPost

> RoleBinding ApisRbacAuthorizationV1ClusterrolebindingsPost(ctx).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding("Role_example") // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost(context.Background()).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsPost`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **roleBinding** | [**RoleBinding**](RoleBinding.md) |  | 

### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolesGet

> GenericListClusterRole ApisRbacAuthorizationV1ClusterrolesGet(ctx).Execute()



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
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolesGet`: GenericListClusterRole
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolesGetRequest struct via the builder pattern


### Return type

[**GenericListClusterRole**](GenericListClusterRole.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolesNameGet

> ClusterRole ApisRbacAuthorizationV1ClusterrolesNameGet(ctx, name).Execute()



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
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolesNameGet`: ClusterRole
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolesNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterRole**](ClusterRole.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete

> ApiStatus ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete(ctx, name).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding("Role_example") // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete(context.Background(), name).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNameRolebindingsDeleteRequest struct via the builder pattern


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

> GenericListRoleBinding ApisRbacAuthorizationV1NamespacesNameRolebindingsGet(ctx, name).Execute()



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
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNameRolebindingsGet`: GenericListRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsGet`: %v\n", resp)
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

[**GenericListRoleBinding**](GenericListRoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1NamespacesNameRolebindingsPost

> ApiStatus ApisRbacAuthorizationV1NamespacesNameRolebindingsPost(ctx, name).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding("Role_example") // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsPost(context.Background(), name).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNameRolebindingsPost`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNameRolebindingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNameRolebindingsPostRequest struct via the builder pattern


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

