# ResourceAccessReviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | Pointer to **string** |  | [optional] 
**Allowed** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceAccessReviewResponse

`func NewResourceAccessReviewResponse() *ResourceAccessReviewResponse`

NewResourceAccessReviewResponse instantiates a new ResourceAccessReviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessReviewResponseWithDefaults

`func NewResourceAccessReviewResponseWithDefaults() *ResourceAccessReviewResponse`

NewResourceAccessReviewResponseWithDefaults instantiates a new ResourceAccessReviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ResourceAccessReviewResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResourceAccessReviewResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResourceAccessReviewResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ResourceAccessReviewResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetAllowed

`func (o *ResourceAccessReviewResponse) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *ResourceAccessReviewResponse) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *ResourceAccessReviewResponse) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.

### HasAllowed

`func (o *ResourceAccessReviewResponse) HasAllowed() bool`

HasAllowed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


