# \KafkaAppsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaAppsV1AppsGet**](KafkaAppsResourceApi.md#ApisKafkaAppsV1AppsGet) | **Get** /apis/kafka.apps/v1/apps | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsGet**](KafkaAppsResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsGet) | **Get** /apis/kafka.apps/v1/namespaces/{namespace}/apps | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete**](KafkaAppsResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete) | **Delete** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsNameGet**](KafkaAppsResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsNameGet) | **Get** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsNamePut**](KafkaAppsResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsNamePut) | **Put** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsPost**](KafkaAppsResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsPost) | **Post** /apis/kafka.apps/v1/namespaces/{namespace}/apps | 



## ApisKafkaAppsV1AppsGet

> GenericListKafkaApplication ApisKafkaAppsV1AppsGet(ctx).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisKafkaAppsV1AppsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisKafkaAppsV1AppsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1AppsGet`: GenericListKafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisKafkaAppsV1AppsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaAppsV1AppsGetRequest struct via the builder pattern


### Return type

[**GenericListKafkaApplication**](GenericListKafkaApplication.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaAppsV1NamespacesNamespaceAppsGet

> GenericListKafkaApplication ApisKafkaAppsV1NamespacesNamespaceAppsGet(ctx, namespace).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsGet`: GenericListKafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaAppsV1NamespacesNamespaceAppsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GenericListKafkaApplication**](GenericListKafkaApplication.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete

> KafkaApplication ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaAppsV1NamespacesNamespaceAppsNameDeleteRequest struct via the builder pattern


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


## ApisKafkaAppsV1NamespacesNamespaceAppsNameGet

> KafkaApplication ApisKafkaAppsV1NamespacesNamespaceAppsNameGet(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsNameGet`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaAppsV1NamespacesNamespaceAppsNameGetRequest struct via the builder pattern


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


## ApisKafkaAppsV1NamespacesNamespaceAppsNamePut

> KafkaApplication ApisKafkaAppsV1NamespacesNamespaceAppsNamePut(ctx, name, namespace).KafkaApplication(kafkaApplication).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNamePut(context.Background(), name, namespace).KafkaApplication(kafkaApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsNamePut`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNamePut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaAppsV1NamespacesNamespaceAppsNamePutRequest struct via the builder pattern


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


## ApisKafkaAppsV1NamespacesNamespaceAppsPost

> KafkaApplication ApisKafkaAppsV1NamespacesNamespaceAppsPost(ctx, namespace).KafkaApplication(kafkaApplication).Execute()



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
    resp, r, err := apiClient.KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsPost(context.Background(), namespace).KafkaApplication(kafkaApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsPost`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaAppsResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaAppsV1NamespacesNamespaceAppsPostRequest struct via the builder pattern


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

