# \AuthorizationResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost**](AuthorizationResourceApi.md#ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost) | **Post** /apis/authorization/v1/namespaces/{namespace}/localresourceaccessreviews | 
[**ApisAuthorizationV1ResourceaccessreviewsPost**](AuthorizationResourceApi.md#ApisAuthorizationV1ResourceaccessreviewsPost) | **Post** /apis/authorization/v1/resourceaccessreviews | 
[**ApisAuthorizationV1SelfsubjectaccessreviewsPost**](AuthorizationResourceApi.md#ApisAuthorizationV1SelfsubjectaccessreviewsPost) | **Post** /apis/authorization/v1/selfsubjectaccessreviews | 
[**ApisAuthorizationV1SubjectaccessreviewsPost**](AuthorizationResourceApi.md#ApisAuthorizationV1SubjectaccessreviewsPost) | **Post** /apis/authorization/v1/subjectaccessreviews | 



## ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost

> ResourceAccessReviewResponse ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost(ctx, namespace).ResourceAccessReview(resourceAccessReview).Execute()



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
    resourceAccessReview := *openapiclient.NewResourceAccessReview("Verb_example", "Resource_example") // ResourceAccessReview |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationResourceApi.ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost(context.Background(), namespace).ResourceAccessReview(resourceAccessReview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationResourceApi.ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost`: ResourceAccessReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationResourceApi.ApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1NamespacesNamespaceLocalresourceaccessreviewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resourceAccessReview** | [**ResourceAccessReview**](ResourceAccessReview.md) |  | 

### Return type

[**ResourceAccessReviewResponse**](ResourceAccessReviewResponse.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAuthorizationV1ResourceaccessreviewsPost

> ResourceAccessReviewResponse ApisAuthorizationV1ResourceaccessreviewsPost(ctx).ResourceAccessReview(resourceAccessReview).Execute()



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
    resourceAccessReview := *openapiclient.NewResourceAccessReview("Verb_example", "Resource_example") // ResourceAccessReview |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationResourceApi.ApisAuthorizationV1ResourceaccessreviewsPost(context.Background()).ResourceAccessReview(resourceAccessReview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationResourceApi.ApisAuthorizationV1ResourceaccessreviewsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1ResourceaccessreviewsPost`: ResourceAccessReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationResourceApi.ApisAuthorizationV1ResourceaccessreviewsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1ResourceaccessreviewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **resourceAccessReview** | [**ResourceAccessReview**](ResourceAccessReview.md) |  | 

### Return type

[**ResourceAccessReviewResponse**](ResourceAccessReviewResponse.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAuthorizationV1SelfsubjectaccessreviewsPost

> SelfSubjectAccessReview ApisAuthorizationV1SelfsubjectaccessreviewsPost(ctx).SelfSubjectAccessReview(selfSubjectAccessReview).Execute()



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
    selfSubjectAccessReview := *openapiclient.NewSelfSubjectAccessReview() // SelfSubjectAccessReview |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationResourceApi.ApisAuthorizationV1SelfsubjectaccessreviewsPost(context.Background()).SelfSubjectAccessReview(selfSubjectAccessReview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationResourceApi.ApisAuthorizationV1SelfsubjectaccessreviewsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1SelfsubjectaccessreviewsPost`: SelfSubjectAccessReview
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationResourceApi.ApisAuthorizationV1SelfsubjectaccessreviewsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1SelfsubjectaccessreviewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **selfSubjectAccessReview** | [**SelfSubjectAccessReview**](SelfSubjectAccessReview.md) |  | 

### Return type

[**SelfSubjectAccessReview**](SelfSubjectAccessReview.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAuthorizationV1SubjectaccessreviewsPost

> SubjectAccessReviewResponse ApisAuthorizationV1SubjectaccessreviewsPost(ctx).SubjectAccessReview(subjectAccessReview).Execute()



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
    subjectAccessReview := *openapiclient.NewSubjectAccessReview("Verb_example", "Resource_example") // SubjectAccessReview |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthorizationResourceApi.ApisAuthorizationV1SubjectaccessreviewsPost(context.Background()).SubjectAccessReview(subjectAccessReview).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthorizationResourceApi.ApisAuthorizationV1SubjectaccessreviewsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAuthorizationV1SubjectaccessreviewsPost`: SubjectAccessReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AuthorizationResourceApi.ApisAuthorizationV1SubjectaccessreviewsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisAuthorizationV1SubjectaccessreviewsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subjectAccessReview** | [**SubjectAccessReview**](SubjectAccessReview.md) |  | 

### Return type

[**SubjectAccessReviewResponse**](SubjectAccessReviewResponse.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

