# SelfSubjectAccessReview

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Spec** | Pointer to [**SelfSubjectAccessReviewSpec**](SelfSubjectAccessReviewSpec.md) |  | [optional] 
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

### GetMetadata

`func (o *SelfSubjectAccessReview) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SelfSubjectAccessReview) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SelfSubjectAccessReview) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SelfSubjectAccessReview) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSpec

`func (o *SelfSubjectAccessReview) GetSpec() SelfSubjectAccessReviewSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *SelfSubjectAccessReview) GetSpecOk() (*SelfSubjectAccessReviewSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *SelfSubjectAccessReview) SetSpec(v SelfSubjectAccessReviewSpec)`

SetSpec sets Spec field to given value.

### HasSpec

`func (o *SelfSubjectAccessReview) HasSpec() bool`

HasSpec returns a boolean if a field has been set.

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


