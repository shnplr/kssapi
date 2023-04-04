# SubjectAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Verb** | **string** |  | 
**Resource** | **string** |  | 
**ResourceName** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**ResourceAPIGroup** | Pointer to **string** |  | [optional] 
**ResourceAPIVersion** | Pointer to **string** |  | [optional] 

## Methods

### NewSubjectAccessReview

`func NewSubjectAccessReview(verb string, resource string, ) *SubjectAccessReview`

NewSubjectAccessReview instantiates a new SubjectAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectAccessReviewWithDefaults

`func NewSubjectAccessReviewWithDefaults() *SubjectAccessReview`

NewSubjectAccessReviewWithDefaults instantiates a new SubjectAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SubjectAccessReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SubjectAccessReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SubjectAccessReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SubjectAccessReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetNamespace

`func (o *SubjectAccessReview) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SubjectAccessReview) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SubjectAccessReview) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SubjectAccessReview) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetVerb

`func (o *SubjectAccessReview) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *SubjectAccessReview) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *SubjectAccessReview) SetVerb(v string)`

SetVerb sets Verb field to given value.


### GetResource

`func (o *SubjectAccessReview) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SubjectAccessReview) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SubjectAccessReview) SetResource(v string)`

SetResource sets Resource field to given value.


### GetResourceName

`func (o *SubjectAccessReview) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *SubjectAccessReview) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *SubjectAccessReview) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *SubjectAccessReview) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetUser

`func (o *SubjectAccessReview) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SubjectAccessReview) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SubjectAccessReview) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SubjectAccessReview) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetResourceAPIGroup

`func (o *SubjectAccessReview) GetResourceAPIGroup() string`

GetResourceAPIGroup returns the ResourceAPIGroup field if non-nil, zero value otherwise.

### GetResourceAPIGroupOk

`func (o *SubjectAccessReview) GetResourceAPIGroupOk() (*string, bool)`

GetResourceAPIGroupOk returns a tuple with the ResourceAPIGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAPIGroup

`func (o *SubjectAccessReview) SetResourceAPIGroup(v string)`

SetResourceAPIGroup sets ResourceAPIGroup field to given value.

### HasResourceAPIGroup

`func (o *SubjectAccessReview) HasResourceAPIGroup() bool`

HasResourceAPIGroup returns a boolean if a field has been set.

### GetResourceAPIVersion

`func (o *SubjectAccessReview) GetResourceAPIVersion() string`

GetResourceAPIVersion returns the ResourceAPIVersion field if non-nil, zero value otherwise.

### GetResourceAPIVersionOk

`func (o *SubjectAccessReview) GetResourceAPIVersionOk() (*string, bool)`

GetResourceAPIVersionOk returns a tuple with the ResourceAPIVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceAPIVersion

`func (o *SubjectAccessReview) SetResourceAPIVersion(v string)`

SetResourceAPIVersion sets ResourceAPIVersion field to given value.

### HasResourceAPIVersion

`func (o *SubjectAccessReview) HasResourceAPIVersion() bool`

HasResourceAPIVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


