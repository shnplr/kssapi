# ResourceAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Verb** | **string** |  | 
**Resource** | **string** |  | 
**ResourceName** | Pointer to **string** |  | [optional] 
**ResourceAPIGroup** | Pointer to **string** |  | [optional] 
**ResourceAPIVersion** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**IsNonResourceURL** | Pointer to **bool** |  | [optional] 
**NonResourceURL** | Pointer to **bool** |  | [optional] 

## Methods

### NewResourceAccessReview

`func NewResourceAccessReview(verb string, resource string, ) *ResourceAccessReview`

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

### GetApiVersion

`func (o *ResourceAccessReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ResourceAccessReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ResourceAccessReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ResourceAccessReview) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

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

### HasNamespace

`func (o *ResourceAccessReview) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

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

### GetResourceAPIGroup

`func (o *ResourceAccessReview) GetResourceAPIGroup() string`

GetResourceAPIGroup returns the ResourceAPIGroup field if non-nil, zero value otherwise.

### GetResourceAPIGroupOk

`func (o *ResourceAccessReview) GetResourceAPIGroupOk() (*string, bool)`

GetResourceAPIGroupOk returns a tuple with the ResourceAPIGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAPIGroup

`func (o *ResourceAccessReview) SetResourceAPIGroup(v string)`

SetResourceAPIGroup sets ResourceAPIGroup field to given value.

### HasResourceAPIGroup

`func (o *ResourceAccessReview) HasResourceAPIGroup() bool`

HasResourceAPIGroup returns a boolean if a field has been set.

### GetResourceAPIVersion

`func (o *ResourceAccessReview) GetResourceAPIVersion() string`

GetResourceAPIVersion returns the ResourceAPIVersion field if non-nil, zero value otherwise.

### GetResourceAPIVersionOk

`func (o *ResourceAccessReview) GetResourceAPIVersionOk() (*string, bool)`

GetResourceAPIVersionOk returns a tuple with the ResourceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAPIVersion

`func (o *ResourceAccessReview) SetResourceAPIVersion(v string)`

SetResourceAPIVersion sets ResourceAPIVersion field to given value.

### HasResourceAPIVersion

`func (o *ResourceAccessReview) HasResourceAPIVersion() bool`

HasResourceAPIVersion returns a boolean if a field has been set.

### GetPath

`func (o *ResourceAccessReview) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ResourceAccessReview) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ResourceAccessReview) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ResourceAccessReview) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetIsNonResourceURL

`func (o *ResourceAccessReview) GetIsNonResourceURL() bool`

GetIsNonResourceURL returns the IsNonResourceURL field if non-nil, zero value otherwise.

### GetIsNonResourceURLOk

`func (o *ResourceAccessReview) GetIsNonResourceURLOk() (*bool, bool)`

GetIsNonResourceURLOk returns a tuple with the IsNonResourceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonResourceURL

`func (o *ResourceAccessReview) SetIsNonResourceURL(v bool)`

SetIsNonResourceURL sets IsNonResourceURL field to given value.

### HasIsNonResourceURL

`func (o *ResourceAccessReview) HasIsNonResourceURL() bool`

HasIsNonResourceURL returns a boolean if a field has been set.

### GetNonResourceURL

`func (o *ResourceAccessReview) GetNonResourceURL() bool`

GetNonResourceURL returns the NonResourceURL field if non-nil, zero value otherwise.

### GetNonResourceURLOk

`func (o *ResourceAccessReview) GetNonResourceURLOk() (*bool, bool)`

GetNonResourceURLOk returns a tuple with the NonResourceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceURL

`func (o *ResourceAccessReview) SetNonResourceURL(v bool)`

SetNonResourceURL sets NonResourceURL field to given value.

### HasNonResourceURL

`func (o *ResourceAccessReview) HasNonResourceURL() bool`

HasNonResourceURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


