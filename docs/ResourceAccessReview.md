# ResourceAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | **string** |  | 
**Verb** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceAccessReview

`func NewResourceAccessReview(namespace string, ) *ResourceAccessReview`

NewResourceAccessReview instantiates a new ResourceAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAccessReviewWithDefaults

`func NewResourceAccessReviewWithDefaults() *ResourceAccessReview`

NewResourceAccessReviewWithDefaults instantiates a new ResourceAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ResourceAccessReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ResourceAccessReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ResourceAccessReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ResourceAccessReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *ResourceAccessReview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ResourceAccessReview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ResourceAccessReview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ResourceAccessReview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *ResourceAccessReview) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ResourceAccessReview) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ResourceAccessReview) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetVerb

`func (o *ResourceAccessReview) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *ResourceAccessReview) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *ResourceAccessReview) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *ResourceAccessReview) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetResource

`func (o *ResourceAccessReview) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ResourceAccessReview) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ResourceAccessReview) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *ResourceAccessReview) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetResourceName

`func (o *ResourceAccessReview) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *ResourceAccessReview) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *ResourceAccessReview) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *ResourceAccessReview) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


