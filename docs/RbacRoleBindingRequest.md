# RbacRoleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 
**User** | **string** |  | 
**Topics** | Pointer to **[]string** |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Subjects** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRbacRoleBindingRequest

`func NewRbacRoleBindingRequest(role string, user string, ) *RbacRoleBindingRequest`

NewRbacRoleBindingRequest instantiates a new RbacRoleBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacRoleBindingRequestWithDefaults

`func NewRbacRoleBindingRequestWithDefaults() *RbacRoleBindingRequest`

NewRbacRoleBindingRequestWithDefaults instantiates a new RbacRoleBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RbacRoleBindingRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RbacRoleBindingRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RbacRoleBindingRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RbacRoleBindingRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *RbacRoleBindingRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RbacRoleBindingRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RbacRoleBindingRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RbacRoleBindingRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetRole

`func (o *RbacRoleBindingRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *RbacRoleBindingRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *RbacRoleBindingRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetUser

`func (o *RbacRoleBindingRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RbacRoleBindingRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RbacRoleBindingRequest) SetUser(v string)`

SetUser sets User field to given value.


### GetTopics

`func (o *RbacRoleBindingRequest) GetTopics() []string`

GetTopics returns the Topics field if non-nil, zero value otherwise.

### GetTopicsOk

`func (o *RbacRoleBindingRequest) GetTopicsOk() (*[]string, bool)`

GetTopicsOk returns a tuple with the Topics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopics

`func (o *RbacRoleBindingRequest) SetTopics(v []string)`

SetTopics sets Topics field to given value.

### HasTopics

`func (o *RbacRoleBindingRequest) HasTopics() bool`

HasTopics returns a boolean if a field has been set.

### GetMetadata

`func (o *RbacRoleBindingRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RbacRoleBindingRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RbacRoleBindingRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RbacRoleBindingRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetSubjects

`func (o *RbacRoleBindingRequest) GetSubjects() []string`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *RbacRoleBindingRequest) GetSubjectsOk() (*[]string, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *RbacRoleBindingRequest) SetSubjects(v []string)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *RbacRoleBindingRequest) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


