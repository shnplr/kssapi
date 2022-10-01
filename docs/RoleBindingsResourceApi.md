# \RoleBindingsResourceApi

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisRbacV1ClusterrolebindingsGet**](RoleBindingsResourceApi.md#ApisRbacV1ClusterrolebindingsGet) | **Get** /apis/rbac/v1/clusterrolebindings | 
[**ApisRbacV1ClusterrolebindingsPost**](RoleBindingsResourceApi.md#ApisRbacV1ClusterrolebindingsPost) | **Post** /apis/rbac/v1/clusterrolebindings | 
[**ApisRbacV1NamespacesNameRolebindingsGet**](RoleBindingsResourceApi.md#ApisRbacV1NamespacesNameRolebindingsGet) | **Get** /apis/rbac/v1/namespaces/{name}/rolebindings | 
[**ApisRbacV1NamespacesNameRolebindingsPost**](RoleBindingsResourceApi.md#ApisRbacV1NamespacesNameRolebindingsPost) | **Post** /apis/rbac/v1/namespaces/{name}/rolebindings | 
[**ApisRbacV1RolebindingsGet**](RoleBindingsResourceApi.md#ApisRbacV1RolebindingsGet) | **Get** /apis/rbac/v1/rolebindings | 



## ApisRbacV1ClusterrolebindingsGet

> []Fact ApisRbacV1ClusterrolebindingsGet(ctx).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacV1ClusterrolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacV1ClusterrolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisRbacV1ClusterrolebindingsGet`: []Fact
    fmt.Fprintf(os.Stdout, "Response from `RoleBindingsResourceApi.ApisRbacV1ClusterrolebindingsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacV1ClusterrolebindingsGetRequest struct via the builder pattern


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


## ApisRbacV1ClusterrolebindingsPost

> ApisRbacV1ClusterrolebindingsPost(ctx).Project(project).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacV1ClusterrolebindingsPost(context.Background()).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacV1ClusterrolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacV1ClusterrolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **project** | [**Project**](Project.md) |  | 

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


## ApisRbacV1NamespacesNameRolebindingsGet

> ApisRbacV1NamespacesNameRolebindingsGet(ctx, name).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacV1NamespacesNameRolebindingsGet(context.Background(), name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacV1NamespacesNameRolebindingsGet``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApisRbacV1NamespacesNameRolebindingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisRbacV1NamespacesNameRolebindingsPost

> ApisRbacV1NamespacesNameRolebindingsPost(ctx, name).Project(project).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacV1NamespacesNameRolebindingsPost(context.Background(), name).Project(project).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacV1NamespacesNameRolebindingsPost``: %v\n", err)
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

Other parameters are passed through a pointer to a apiApisRbacV1NamespacesNameRolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **project** | [**Project**](Project.md) |  | 

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


## ApisRbacV1RolebindingsGet

> ApisRbacV1RolebindingsGet(ctx).Execute()



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
    resp, r, err := apiClient.RoleBindingsResourceApi.ApisRbacV1RolebindingsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RoleBindingsResourceApi.ApisRbacV1RolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisRbacV1RolebindingsGetRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

