# \KafkaTopicsResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisKafkaTopicV1NamespacesNamespaceTopicsGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/configs:alter | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete) | **Delete** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name}/describe | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet) | **Get** /apis/kafka.topic/v1/namespaces/{namespace}/topics/{name} | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsPost**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics | 
[**ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost) | **Post** /apis/kafka.topic/v1/namespaces/{namespace}/topics/sync | 
[**ApisKafkaTopicV1TopicsGet**](KafkaTopicsResourceApi.md#ApisKafkaTopicV1TopicsGet) | **Get** /apis/kafka.topic/v1/topics | 



## ApisKafkaTopicV1NamespacesNamespaceTopicsGet

> KafkaTopicList ApisKafkaTopicV1NamespacesNamespaceTopicsGet(ctx, namespace).Execute()



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
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsGet`: KafkaTopicList
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**KafkaTopicList**](KafkaTopicList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost(ctx, name, namespace).KafkaTopicRequest(kafkaTopicRequest).Execute()



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
    kafkaTopicRequest := *openapiclient.NewKafkaTopicRequest() // KafkaTopicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost(context.Background(), name, namespace).KafkaTopicRequest(kafkaTopicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameConfigsalterPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **kafkaTopicRequest** | [**KafkaTopicRequest**](KafkaTopicRequest.md) |  | 

### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameDescribeGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(ctx, name, namespace).Execute()



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
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet(context.Background(), name, namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsNameGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**name** | **string** |  | 
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsNameGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsPost

> KafkaTopic ApisKafkaTopicV1NamespacesNamespaceTopicsPost(ctx, namespace).KafkaTopicRequest(kafkaTopicRequest).Execute()



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
    kafkaTopicRequest := *openapiclient.NewKafkaTopicRequest() // KafkaTopicRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost(context.Background(), namespace).KafkaTopicRequest(kafkaTopicRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsPost`: KafkaTopic
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **kafkaTopicRequest** | [**KafkaTopicRequest**](KafkaTopicRequest.md) |  | 

### Return type

[**KafkaTopic**](KafkaTopic.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost

> ApiStatus ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost(ctx, namespace).Execute()



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
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost(context.Background(), namespace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost`: ApiStatus
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1NamespacesNamespaceTopicsSyncPost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1NamespacesNamespaceTopicsSyncPostRequest struct via the builder pattern


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


## ApisKafkaTopicV1TopicsGet

> KafkaTopicList ApisKafkaTopicV1TopicsGet(ctx).Execute()



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
    resp, r, err := apiClient.KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApisKafkaTopicV1TopicsGet`: KafkaTopicList
    fmt.Fprintf(os.Stdout, "Response from `KafkaTopicsResourceApi.ApisKafkaTopicV1TopicsGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiApisKafkaTopicV1TopicsGetRequest struct via the builder pattern


### Return type

[**KafkaTopicList**](KafkaTopicList.md)

### Authorization

[SecurityScheme](../README.md#SecurityScheme)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

