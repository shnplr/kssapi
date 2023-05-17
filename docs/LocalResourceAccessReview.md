# LocalResourceAccessReview

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

### NewLocalResourceAccessReview

`func NewLocalResourceAccessReview(verb string, resource string, ) *LocalResourceAccessReview`

NewLocalResourceAccessReview instantiates a new LocalResourceAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocalResourceAccessReviewWithDefaults

`func NewLocalResourceAccessReviewWithDefaults() *LocalResourceAccessReview`

NewLocalResourceAccessReviewWithDefaults instantiates a new LocalResourceAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *LocalResourceAccessReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LocalResourceAccessReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LocalResourceAccessReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LocalResourceAccessReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *LocalResourceAccessReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *LocalResourceAccessReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *LocalResourceAccessReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *LocalResourceAccessReview) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetNamespace

`func (o *LocalResourceAccessReview) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *LocalResourceAccessReview) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *LocalResourceAccessReview) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *LocalResourceAccessReview) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetVerb

`func (o *LocalResourceAccessReview) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *LocalResourceAccessReview) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *LocalResourceAccessReview) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetResource

`func (o *LocalResourceAccessReview) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *LocalResourceAccessReview) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *LocalResourceAccessReview) SetResource(v string)`

SetResource sets Resource field to given value.


### GetResourceName

`func (o *LocalResourceAccessReview) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *LocalResourceAccessReview) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *LocalResourceAccessReview) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *LocalResourceAccessReview) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetResourceAPIGroup

`func (o *LocalResourceAccessReview) GetResourceAPIGroup() string`

GetResourceAPIGroup returns the ResourceAPIGroup field if non-nil, zero value otherwise.

### GetResourceAPIGroupOk

`func (o *LocalResourceAccessReview) GetResourceAPIGroupOk() (*string, bool)`

GetResourceAPIGroupOk returns a tuple with the ResourceAPIGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAPIGroup

`func (o *LocalResourceAccessReview) SetResourceAPIGroup(v string)`

SetResourceAPIGroup sets ResourceAPIGroup field to given value.

### HasResourceAPIGroup

`func (o *LocalResourceAccessReview) HasResourceAPIGroup() bool`

HasResourceAPIGroup returns a boolean if a field has been set.

### GetResourceAPIVersion

`func (o *LocalResourceAccessReview) GetResourceAPIVersion() string`

GetResourceAPIVersion returns the ResourceAPIVersion field if non-nil, zero value otherwise.

### GetResourceAPIVersionOk

`func (o *LocalResourceAccessReview) GetResourceAPIVersionOk() (*string, bool)`

GetResourceAPIVersionOk returns a tuple with the ResourceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAPIVersion

`func (o *LocalResourceAccessReview) SetResourceAPIVersion(v string)`

SetResourceAPIVersion sets ResourceAPIVersion field to given value.

### HasResourceAPIVersion

`func (o *LocalResourceAccessReview) HasResourceAPIVersion() bool`

HasResourceAPIVersion returns a boolean if a field has been set.

### GetPath

`func (o *LocalResourceAccessReview) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LocalResourceAccessReview) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LocalResourceAccessReview) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LocalResourceAccessReview) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetIsNonResourceURL

`func (o *LocalResourceAccessReview) GetIsNonResourceURL() bool`

GetIsNonResourceURL returns the IsNonResourceURL field if non-nil, zero value otherwise.

### GetIsNonResourceURLOk

`func (o *LocalResourceAccessReview) GetIsNonResourceURLOk() (*bool, bool)`

GetIsNonResourceURLOk returns a tuple with the IsNonResourceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsNonResourceURL

`func (o *LocalResourceAccessReview) SetIsNonResourceURL(v bool)`

SetIsNonResourceURL sets IsNonResourceURL field to given value.

### HasIsNonResourceURL

`func (o *LocalResourceAccessReview) HasIsNonResourceURL() bool`

HasIsNonResourceURL returns a boolean if a field has been set.

### GetNonResourceURL

`func (o *LocalResourceAccessReview) GetNonResourceURL() bool`

GetNonResourceURL returns the NonResourceURL field if non-nil, zero value otherwise.

### GetNonResourceURLOk

`func (o *LocalResourceAccessReview) GetNonResourceURLOk() (*bool, bool)`

GetNonResourceURLOk returns a tuple with the NonResourceURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceURL

`func (o *LocalResourceAccessReview) SetNonResourceURL(v bool)`

SetNonResourceURL sets NonResourceURL field to given value.

### HasNonResourceURL

`func (o *LocalResourceAccessReview) HasNonResourceURL() bool`

HasNonResourceURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


