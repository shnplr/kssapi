# \RbacAuthResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisRbacAuthorizationV1ClusterrolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsGet) | **Get** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsNameDelete**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsNameDelete) | **Delete** /apis/rbac.authorization/v1/clusterrolebindings/{name} | 
[**ApisRbacAuthorizationV1ClusterrolebindingsNameGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsNameGet) | **Get** /apis/rbac.authorization/v1/clusterrolebindings/{name} | 
[**ApisRbacAuthorizationV1ClusterrolebindingsNamePut**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsNamePut) | **Put** /apis/rbac.authorization/v1/clusterrolebindings/{name} | 
[**ApisRbacAuthorizationV1ClusterrolebindingsPost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsPost) | **Post** /apis/rbac.authorization/v1/clusterrolebindings | 
[**ApisRbacAuthorizationV1ClusterrolebindingsdeletePost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolebindingsdeletePost) | **Post** /apis/rbac.authorization/v1/clusterrolebindings:delete | 
[**ApisRbacAuthorizationV1ClusterrolesGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolesGet) | **Get** /apis/rbac.authorization/v1/clusterroles | 
[**ApisRbacAuthorizationV1ClusterrolesNameGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1ClusterrolesNameGet) | **Get** /apis/rbac.authorization/v1/clusterroles/{name} | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet) | **Get** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete) | **Delete** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings/{name} | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet) | **Get** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings/{name} | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut) | **Put** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings/{name} | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsPost) | **Post** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings | 
[**ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost) | **Post** /apis/rbac.authorization/v1/namespaces/{namespace}/rolebindings:delete | 
[**ApisRbacAuthorizationV1RolebindingsGet**](RbacAuthResourceApi.md#ApisRbacAuthorizationV1RolebindingsGet) | **Get** /apis/rbac.authorization/v1/rolebindings | 



## ApisRbacAuthorizationV1ClusterrolebindingsGet

> ClusterRoleBindingList ApisRbacAuthorizationV1ClusterrolebindingsGet(ctx).Execute()



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
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsGet`: ClusterRoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsGetRequest struct via the builder pattern


### Return type

[**ClusterRoleBindingList**](ClusterRoleBindingList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolebindingsNameDelete

> Status ApisRbacAuthorizationV1ClusterrolebindingsNameDelete(ctx, name).Execute()



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
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNameDelete(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsNameDelete`: Status
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Status**](Status.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolebindingsNameGet

> ClusterRoleBinding ApisRbacAuthorizationV1ClusterrolebindingsNameGet(ctx, name).Execute()



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
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNameGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsNameGet`: ClusterRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ClusterRoleBinding**](ClusterRoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1ClusterrolebindingsNamePut

> ClusterRoleBinding ApisRbacAuthorizationV1ClusterrolebindingsNamePut(ctx, name).ClusterRoleBinding(clusterRoleBinding).Execute()



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
    clusterRoleBinding := *openapiclient.NewClusterRoleBinding() // ClusterRoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNamePut(context.Background(), name).ClusterRoleBinding(clusterRoleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsNamePut`: ClusterRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsNamePutRequest struct via the builder pattern


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
    clusterRoleBinding := *openapiclient.NewClusterRoleBinding() // ClusterRoleBinding |  (optional)

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


## ApisRbacAuthorizationV1ClusterrolebindingsdeletePost

> ClusterRoleBinding ApisRbacAuthorizationV1ClusterrolebindingsdeletePost(ctx).ClusterRoleBinding(clusterRoleBinding).Execute()



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
    clusterRoleBinding := *openapiclient.NewClusterRoleBinding() // ClusterRoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsdeletePost(context.Background()).ClusterRoleBinding(clusterRoleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsdeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1ClusterrolebindingsdeletePost`: ClusterRoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolebindingsdeletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolebindingsdeletePostRequest struct via the builder pattern


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

> ClusterRoleList ApisRbacAuthorizationV1ClusterrolesGet(ctx).Execute()



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
    // response from `ApisRbacAuthorizationV1ClusterrolesGet`: ClusterRoleList
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1ClusterrolesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1ClusterrolesGetRequest struct via the builder pattern


### Return type

[**ClusterRoleList**](ClusterRoleList.md)

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


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet

> RoleBindingList ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet(ctx, namespace).Execute()



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
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsGet`: RoleBindingList
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

[**RoleBindingList**](RoleBindingList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete

> Status ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete(ctx, name, namespace).Execute()



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
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete`: Status
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Status**](Status.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet

> RoleBinding ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet(ctx, name, namespace).Execute()



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
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RoleBinding**](RoleBinding.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut

> RoleBinding ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut(ctx, name, namespace).RoleBinding(roleBinding).Execute()



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
    namespace := "namespace_example" // string | 
    roleBinding := *openapiclient.NewRoleBinding() // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut(context.Background(), name, namespace).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsNamePutRequest struct via the builder pattern


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
    roleBinding := *openapiclient.NewRoleBinding() // RoleBinding |  (optional)

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


## ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost

> RoleBinding ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost(ctx, namespace).RoleBinding(roleBinding).Execute()



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
    roleBinding := *openapiclient.NewRoleBinding() // RoleBinding |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost(context.Background(), namespace).RoleBinding(roleBinding).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost`: RoleBinding
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1NamespacesNamespaceRolebindingsdeletePostRequest struct via the builder pattern


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

> RoleBindingList ApisRbacAuthorizationV1RolebindingsGet(ctx).Execute()



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
    // response from `ApisRbacAuthorizationV1RolebindingsGet`: RoleBindingList
    fmt.Fprintf(os.Stdout, "Response from `RbacAuthResourceApi.ApisRbacAuthorizationV1RolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacAuthorizationV1RolebindingsGetRequest struct via the builder pattern


### Return type

[**RoleBindingList**](RoleBindingList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

