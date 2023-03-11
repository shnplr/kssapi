# \RbacAuthResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisRbacAuthorizationV1ClusterrolebindingsDelete**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsDelete) | **Delete** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsGet) | **Get** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsPost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsPost) | **Post** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolesGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolesGet) | **Get** /apis/rbac.authorization/v1/clusterroles | 
[**ApisRbacAuthorizationV1ClusterrolesNameGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolesNameGet) | **Get** /apis/rbac.authorization/v1/clusterroles/{name} | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete) | **Delete** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet) | **Get** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost) | **Post** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
[**ApisRbacAuthorizationV1RolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1RolebindingsGet) | **Get** /apis/rbac.authorization/v1/rolebindings | 



## ApisRbacAuthorizationV1ClusterrolebindingsDelete

> ApiStatus ApisRbacAuthorizationV1ClusterrolebindingsDelete(ctx).ClusterRoleBinding(clusterRoleBinding).Execute()



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
    clusterRoleBinding := *openapiclient.NewClusterRoleBinding("Name_example", *openapiclient.NewRoleRef("Name_example")) // ClusterRoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsDelete(context.Background()).ClusterRoleBinding(clusterRoleBinding).Execute()
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
 **clusterRoleBinding** | [**ClusterRoleBinding**](ClusterRoleBinding.md) |  | 

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

> GenericListClusterRoleBinding ApisRbacAuthorizationV1ClusterrolebindingsGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsGet`: GenericListClusterRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsGetRequest struct via the builder pattern


### Return type

[**GenericListClusterRoleBinding**](GenericListClusterRoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolebindingsPost

> ClusterRoleBinding ApisRbacAuthorizationV1ClusterrolebindingsPost(ctx).ClusterRoleBinding(clusterRoleBinding).Execute()



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
    clusterRoleBinding := *openapiclient.NewClusterRoleBinding("Name_example", *openapiclient.NewRoleRef("Name_example")) // ClusterRoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost(context.Background()).ClusterRoleBinding(clusterRoleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsPost`: ClusterRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **clusterRoleBinding** | [**ClusterRoleBinding**](ClusterRoleBinding.md) |  | 

### Return type

[**ClusterRoleBinding**](ClusterRoleBinding.md)

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
    openapiclient "github.com/shnplr/kssapi"
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
    openapiclient "github.com/shnplr/kssapi"
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


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete

> ApiStatus ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete(ctx, namespace).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding("Name_example", "Namespace_example", *openapiclient.NewRoleRef("Name_example")) // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete(context.Background(), namespace).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsDeleteRequest struct via the builder pattern


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


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet

> GenericListRoleBinding ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet(ctx, namespace).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet`: GenericListRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGetRequest struct via the builder pattern


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


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost

> RoleBinding ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost(ctx, namespace).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding("Name_example", "Namespace_example", *openapiclient.NewRoleRef("Name_example")) // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost(context.Background(), namespace).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost`: %v\n", resp)
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

[**RoleBinding**](RoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1RolebindingsGet

> GenericListRoleBinding ApisRbacAuthorizationV1RolebindingsGet(ctx).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1RolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1RolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1RolebindingsGet`: GenericListRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1RolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1RolebindingsGetRequest struct via the builder pattern


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

