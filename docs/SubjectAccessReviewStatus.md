# SubjectAccessReviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** |  | 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewSubjectAccessReviewStatus

`func NewSubjectAccessReviewStatus(allowed bool, ) *SubjectAccessReviewStatus`

NewSubjectAccessReviewStatus instantiates a new SubjectAccessReviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubjectAccessReviewStatusWithDefaults

`func NewSubjectAccessReviewStatusWithDefaults() *SubjectAccessReviewStatus`

NewSubjectAccessReviewStatusWithDefaults instantiates a new SubjectAccessReviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllowed

`func (o *SubjectAccessReviewStatus) GetAllowed() bool`

GetAllowed returns the Allowed field if non-nil, zero value otherwise.

### GetAllowedOk

`func (o *SubjectAccessReviewStatus) GetAllowedOk() (*bool, bool)`

GetAllowedOk returns a tuple with the Allowed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowed

`func (o *SubjectAccessReviewStatus) SetAllowed(v bool)`

SetAllowed sets Allowed field to given value.


### GetReason

`func (o *SubjectAccessReviewStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *SubjectAccessReviewStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *SubjectAccessReviewStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *SubjectAccessReviewStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


