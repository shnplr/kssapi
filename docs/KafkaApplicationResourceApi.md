# \KafkaApplicationResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaAppsV1NamespacesNamespaceAppsGet**](KafkaApplicationResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsGet) | **Get** /apis/kafka.apps/v1/namespaces/{namespace}/apps | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete**](KafkaApplicationResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete) | **Delete** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsNameGet**](KafkaApplicationResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsNameGet) | **Get** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsNamePut**](KafkaApplicationResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsNamePut) | **Put** /apis/kafka.apps/v1/namespaces/{namespace}/apps/{name} | 
[**ApisKafkaAppsV1NamespacesNamespaceAppsPost**](KafkaApplicationResourceApi.md#ApisKafkaAppsV1NamespacesNamespaceAppsPost) | **Post** /apis/kafka.apps/v1/namespaces/{namespace}/apps | 



## ApisKafkaAppsV1NamespacesNamespaceAppsGet

> GenericListKafkaApplication ApisKafkaAppsV1NamespacesNamespaceAppsGet(ctx, namespace).Execute()



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsGet`: GenericListKafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsGet`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameDelete`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsNameGet`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNameGet`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    name := "name_example" // string | 
    namespace := "namespace_example" // string | 
    kafkaApplication := *openapiclient.NewKafkaApplication("Name_example", "Namespace_example", "Type_example", "Principal_example") // KafkaApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNamePut(context.Background(), name, namespace).KafkaApplication(kafkaApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNamePut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsNamePut`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsNamePut`: %v\n", resp)
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
    openapiclient "./openapi"
)

func main() {
    namespace := "namespace_example" // string | 
    kafkaApplication := *openapiclient.NewKafkaApplication("Name_example", "Namespace_example", "Type_example", "Principal_example") // KafkaApplication |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsPost(context.Background(), namespace).KafkaApplication(kafkaApplication).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaAppsV1NamespacesNamespaceAppsPost`: KafkaApplication
    fmt.Fprintf(os.Stdout, "Response from `KafkaApplicationResourceApi.ApisKafkaAppsV1NamespacesNamespaceAppsPost`: %v\n", resp)
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
