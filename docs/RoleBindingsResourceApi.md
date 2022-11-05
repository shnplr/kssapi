# \RoleBindingsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAuthorizationV1ClusterrolebindingsGet**](RoleBindingsResourceApi.md#ApisAuthorizationV1ClusterrolebindingsGet) | **Get** /apis/authorization/v1/clusterrolebindings | 
[**ApisAuthorizationV1ClusterrolebindingsPost**](RoleBindingsResourceApi.md#ApisAuthorizationV1ClusterrolebindingsPost) | **Post** /apis/authorization/v1/clusterrolebindings | 
[**ApisAuthorizationV1NamespacesNameRolebindingsGet**](RoleBindingsResourceApi.md#ApisAuthorizationV1NamespacesNameRolebindingsGet) | **Get** /apis/authorization/v1/namespaces/{name}/rolebindings | 
[**ApisAuthorizationV1NamespacesNameRolebindingsPost**](RoleBindingsResourceApi.md#ApisAuthorizationV1NamespacesNameRolebindingsPost) | **Post** /apis/authorization/v1/namespaces/{name}/rolebindings | 
[**ApisAuthorizationV1RolebindingsGet**](RoleBindingsResourceApi.md#ApisAuthorizationV1RolebindingsGet) | **Get** /apis/authorization/v1/rolebindings | 



## ApisAuthorizationV1ClusterrolebindingsGet

> []Fact ApisAuthorizationV1ClusterrolebindingsGet(ctx).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisAuthorizationV1ClusterrolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisAuthorizationV1ClusterrolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1ClusterrolebindingsGet`: []Fact
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisAuthorizationV1ClusterrolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1ClusterrolebindingsGetRequest struct via the builder pattern


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


## ApisAuthorizationV1ClusterrolebindingsPost

> ApiStatus ApisAuthorizationV1ClusterrolebindingsPost(ctx).Project(project).Execute()



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
    project := *openapiclient.NewProject() // Project |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisAuthorizationV1ClusterrolebindingsPost(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisAuthorizationV1ClusterrolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1ClusterrolebindingsPost`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisAuthorizationV1ClusterrolebindingsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1ClusterrolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) |  | 

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


## ApisAuthorizationV1NamespacesNameRolebindingsGet

> ApiStatus ApisAuthorizationV1NamespacesNameRolebindingsGet(ctx, name).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisAuthorizationV1NamespacesNameRolebindingsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisAuthorizationV1NamespacesNameRolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1NamespacesNameRolebindingsGet`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisAuthorizationV1NamespacesNameRolebindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1NamespacesNameRolebindingsGetRequest struct via the builder pattern


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


## ApisAuthorizationV1NamespacesNameRolebindingsPost

> ApiStatus ApisAuthorizationV1NamespacesNameRolebindingsPost(ctx, name).Project(project).Execute()



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
    project := *openapiclient.NewProject() // Project |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisAuthorizationV1NamespacesNameRolebindingsPost(context.Background(), name).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisAuthorizationV1NamespacesNameRolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1NamespacesNameRolebindingsPost`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisAuthorizationV1NamespacesNameRolebindingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1NamespacesNameRolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) |  | 

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


## ApisAuthorizationV1RolebindingsGet

> ApiStatus ApisAuthorizationV1RolebindingsGet(ctx).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisAuthorizationV1RolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisAuthorizationV1RolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1RolebindingsGet`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisAuthorizationV1RolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1RolebindingsGetRequest struct via the builder pattern


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

