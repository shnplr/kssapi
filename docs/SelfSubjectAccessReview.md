# SelfSubjectAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Verb** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**SubjectAccessReviewStatus**](SubjectAccessReviewStatus.md) |  | [optional] 

## Methods

### NewSelfSubjectAccessReview

`func NewSelfSubjectAccessReview() *SelfSubjectAccessReview`

NewSelfSubjectAccessReview instantiates a new SelfSubjectAccessReview object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelfSubjectAccessReviewWithDefaults

`func NewSelfSubjectAccessReviewWithDefaults() *SelfSubjectAccessReview`

NewSelfSubjectAccessReviewWithDefaults instantiates a new SelfSubjectAccessReview object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *SelfSubjectAccessReview) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *SelfSubjectAccessReview) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *SelfSubjectAccessReview) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *SelfSubjectAccessReview) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *SelfSubjectAccessReview) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *SelfSubjectAccessReview) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *SelfSubjectAccessReview) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *SelfSubjectAccessReview) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetName

`func (o *SelfSubjectAccessReview) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SelfSubjectAccessReview) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SelfSubjectAccessReview) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SelfSubjectAccessReview) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *SelfSubjectAccessReview) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *SelfSubjectAccessReview) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *SelfSubjectAccessReview) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *SelfSubjectAccessReview) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetVerb

`func (o *SelfSubjectAccessReview) GetVerb() string`

GetVerb returns the Verb field if non-nil, zero value otherwise.

### GetVerbOk

`func (o *SelfSubjectAccessReview) GetVerbOk() (*string, bool)`

GetVerbOk returns a tuple with the Verb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerb

`func (o *SelfSubjectAccessReview) SetVerb(v string)`

SetVerb sets Verb field to given value.

### HasVerb

`func (o *SelfSubjectAccessReview) HasVerb() bool`

HasVerb returns a boolean if a field has been set.

### GetResource

`func (o *SelfSubjectAccessReview) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *SelfSubjectAccessReview) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *SelfSubjectAccessReview) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *SelfSubjectAccessReview) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetGroup

`func (o *SelfSubjectAccessReview) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *SelfSubjectAccessReview) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *SelfSubjectAccessReview) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *SelfSubjectAccessReview) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetVersion

`func (o *SelfSubjectAccessReview) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SelfSubjectAccessReview) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SelfSubjectAccessReview) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SelfSubjectAccessReview) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *SelfSubjectAccessReview) GetStatus() SubjectAccessReviewStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SelfSubjectAccessReview) GetStatusOk() (*SubjectAccessReviewStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SelfSubjectAccessReview) SetStatus(v SubjectAccessReviewStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SelfSubjectAccessReview) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


