# \KafkaAppsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAppsV1KafkaappsGet**](KafkaAppsResourceApi.md#ApisAppsV1KafkaappsGet) | **Get** /apis/apps/v1/kafkaapps | 
[**ApisAppsV1NamespacesNamespaceKafkaappsGet**](KafkaAppsResourceApi.md#ApisAppsV1NamespacesNamespaceKafkaappsGet) | **Get** /apis/apps/v1/namespaces/{namespace}/kafkaapps | 
[**ApisAppsV1NamespacesNamespaceKafkaappsNameDelete**](KafkaAppsResourceApi.md#ApisAppsV1NamespacesNamespaceKafkaappsNameDelete) | **Delete** /apis/apps/v1/namespaces/{namespace}/kafkaapps/{name} | 
[**ApisAppsV1NamespacesNamespaceKafkaappsNameGet**](KafkaAppsResourceApi.md#ApisAppsV1NamespacesNamespaceKafkaappsNameGet) | **Get** /apis/apps/v1/namespaces/{namespace}/kafkaapps/{name} | 
[**ApisAppsV1NamespacesNamespaceKafkaappsNamePut**](KafkaAppsResourceApi.md#ApisAppsV1NamespacesNamespaceKafkaappsNamePut) | **Put** /apis/apps/v1/namespaces/{namespace}/kafkaapps/{name} | 
[**ApisAppsV1NamespacesNamespaceKafkaappsPost**](KafkaAppsResourceApi.md#ApisAppsV1NamespacesNamespaceKafkaappsPost) | **Post** /apis/apps/v1/namespaces/{namespace}/kafkaapps | 



## ApisAppsV1KafkaappsGet

> KafkaApplicationList ApisAppsV1KafkaappsGet(ctx).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisAppsV1KafkaappsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisAppsV1KafkaappsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsV1KafkaappsGet`: KafkaApplicationList
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisAppsV1KafkaappsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsV1KafkaappsGetRequest struct via the builder pattern


### Return type

[**KafkaApplicationList**](KafkaApplicationList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsV1NamespacesNamespaceKafkaappsGet

> KafkaApplicationList ApisAppsV1NamespacesNamespaceKafkaappsGet(ctx, namespace).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsV1NamespacesNamespaceKafkaappsGet`: KafkaApplicationList
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsV1NamespacesNamespaceKafkaappsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaApplicationList**](KafkaApplicationList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsV1NamespacesNamespaceKafkaappsNameDelete

> Status ApisAppsV1NamespacesNamespaceKafkaappsNameDelete(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsV1NamespacesNamespaceKafkaappsNameDelete`: Status
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsV1NamespacesNamespaceKafkaappsNameDeleteRequest struct via the builder pattern


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


## ApisAppsV1NamespacesNamespaceKafkaappsNameGet

> KafkaApplication ApisAppsV1NamespacesNamespaceKafkaappsNameGet(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsV1NamespacesNamespaceKafkaappsNameGet`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsV1NamespacesNamespaceKafkaappsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaApplication**](KafkaApplication.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsV1NamespacesNamespaceKafkaappsNamePut

> KafkaApplication ApisAppsV1NamespacesNamespaceKafkaappsNamePut(ctx, name, namespace).KafkaApplication(kafkaApplication).Execute()



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
    kafkaApplication := *openapiclient.NewKafkaApplication("Type_example", "Principal_example") // KafkaApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNamePut(context.Background(), name, namespace).KafkaApplication(kafkaApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsV1NamespacesNamespaceKafkaappsNamePut`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsV1NamespacesNamespaceKafkaappsNamePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kafkaApplication** | [**KafkaApplication**](KafkaApplication.md) |  | 

### Return type

[**KafkaApplication**](KafkaApplication.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisAppsV1NamespacesNamespaceKafkaappsPost

> KafkaApplication ApisAppsV1NamespacesNamespaceKafkaappsPost(ctx, namespace).KafkaApplication(kafkaApplication).Execute()



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
    kafkaApplication := *openapiclient.NewKafkaApplication("Type_example", "Principal_example") // KafkaApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsPost(context.Background(), namespace).KafkaApplication(kafkaApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisAppsV1NamespacesNamespaceKafkaappsPost`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisAppsV1NamespacesNamespaceKafkaappsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisAppsV1NamespacesNamespaceKafkaappsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaApplication** | [**KafkaApplication**](KafkaApplication.md) |  | 

### Return type

[**KafkaApplication**](KafkaApplication.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

