# \KafkaRbacResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaRbacV1KafkarolebindingsGet**](KafkaRbacResourceApi.md#ApisKafkaRbacV1KafkarolebindingsGet) | **Get** /apis/kafka.rbac/v1/kafkarolebindings | 
[**ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet) | **Get** /apis/kafka.rbac/v1/namespaces/{namespace}/kafkarolebindings | 
[**ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost**](KafkaRbacResourceApi.md#ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost) | **Post** /apis/kafka.rbac/v1/namespaces/{namespace}/kafkarolebindings | 



## ApisKafkaRbacV1KafkarolebindingsGet

> KafkaRbacSummaryList ApisKafkaRbacV1KafkarolebindingsGet(ctx).Role(role).User(user).Execute()



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
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1KafkarolebindingsGet(context.Background()).Role(role).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1KafkarolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1KafkarolebindingsGet`: KafkaRbacSummaryList
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1KafkarolebindingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1KafkarolebindingsGetRequest struct via the builder pattern


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


## ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet

> KafkaRbacSummaryList ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet(ctx, namespace).Role(role).User(user).Execute()



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
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet(context.Background(), namespace).Role(role).User(user).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet`: KafkaRbacSummaryList
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsGetRequest struct via the builder pattern


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


## ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost

> Status ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost(ctx, namespace).KafkaRoleBindingRequest(kafkaRoleBindingRequest).Execute()



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
    kafkaRoleBindingRequest := *openapiclient.NewKafkaRoleBindingRequest("Role_example", "User_example") // KafkaRoleBindingRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost(context.Background(), namespace).KafkaRoleBindingRequest(kafkaRoleBindingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost`: Status
    fmt.Fprintf(os.Stdout, "Response from `KafkaRbacResourceApi.ApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaRbacV1NamespacesNamespaceKafkarolebindingsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaRoleBindingRequest** | [**KafkaRoleBindingRequest**](KafkaRoleBindingRequest.md) |  | 

### Return type

[**Status**](Status.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

