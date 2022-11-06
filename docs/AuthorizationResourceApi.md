# \AuthorizationResourceApi

All URIs are relative to *http://localhost:9080*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ApisAuthorizationV1SubjectaccessreviewsPost**](AuthorizationResourceApi.md#ApisAuthorizationV1SubjectaccessreviewsPost) | **Post** /apis/authorization/v1/subjectaccessreviews | 



## ApisAuthorizationV1SubjectaccessreviewsPost

> SubjectAccessReviewResponse ApisAuthorizationV1SubjectaccessreviewsPost(ctx).SubjectAccessReview(subjectAccessReview).Execute()



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
    subjectAccessReview := *openapiclient.NewSubjectAccessReview() // SubjectAccessReview |  (optional)

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

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

