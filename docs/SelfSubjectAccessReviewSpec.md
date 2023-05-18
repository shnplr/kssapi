# SelfSubjectAccessReviewSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceAttributes** | Pointer to [**ResourceAttributes**](ResourceAttributes.md) |  | [optional] 
**NonResourceAttributes** | Pointer to [**NonResourceAttributes**](NonResourceAttributes.md) |  | [optional] 

## Methods

### NewSelfSubjectAccessReviewSpec

`func NewSelfSubjectAccessReviewSpec() *SelfSubjectAccessReviewSpec`

NewSelfSubjectAccessReviewSpec instantiates a new SelfSubjectAccessReviewSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfSubjectAccessReviewSpecWithDefaults

`func NewSelfSubjectAccessReviewSpecWithDefaults() *SelfSubjectAccessReviewSpec`

NewSelfSubjectAccessReviewSpecWithDefaults instantiates a new SelfSubjectAccessReviewSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceAttributes

`func (o *SelfSubjectAccessReviewSpec) GetResourceAttributes() ResourceAttributes`

GetResourceAttributes returns the ResourceAttributes field if non-nil, zero value otherwise.

### GetResourceAttributesOk

`func (o *SelfSubjectAccessReviewSpec) GetResourceAttributesOk() (*ResourceAttributes, bool)`

GetResourceAttributesOk returns a tuple with the ResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAttributes

`func (o *SelfSubjectAccessReviewSpec) SetResourceAttributes(v ResourceAttributes)`

SetResourceAttributes sets ResourceAttributes field to given value.

### HasResourceAttributes

`func (o *SelfSubjectAccessReviewSpec) HasResourceAttributes() bool`

HasResourceAttributes returns a boolean if a field has been set.

### GetNonResourceAttributes

`func (o *SelfSubjectAccessReviewSpec) GetNonResourceAttributes() NonResourceAttributes`

GetNonResourceAttributes returns the NonResourceAttributes field if non-nil, zero value otherwise.

### GetNonResourceAttributesOk

`func (o *SelfSubjectAccessReviewSpec) GetNonResourceAttributesOk() (*NonResourceAttributes, bool)`

GetNonResourceAttributesOk returns a tuple with the NonResourceAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonResourceAttributes

`func (o *SelfSubjectAccessReviewSpec) SetNonResourceAttributes(v NonResourceAttributes)`

SetNonResourceAttributes sets NonResourceAttributes field to given value.

### HasNonResourceAttributes

`func (o *SelfSubjectAccessReviewSpec) HasNonResourceAttributes() bool`

HasNonResourceAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


