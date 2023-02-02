# RbacRoleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | **string** |  | 
**Role** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Topics** | Pointer to **[]string** |  | [optional] 
**Subjects** | Pointer to **[]string** |  | [optional] 
**Groups** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRbacRoleBindingRequest

`func NewRbacRoleBindingRequest(namespace string, ) *RbacRoleBindingRequest`

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

### GetName

`func (o *RbacRoleBindingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbacRoleBindingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbacRoleBindingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RbacRoleBindingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *RbacRoleBindingRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RbacRoleBindingRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RbacRoleBindingRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


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

### HasRole

`func (o *RbacRoleBindingRequest) HasRole() bool`

HasRole returns a boolean if a field has been set.

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

### HasUser

`func (o *RbacRoleBindingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

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

### GetGroups

`func (o *RbacRoleBindingRequest) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *RbacRoleBindingRequest) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *RbacRoleBindingRequest) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *RbacRoleBindingRequest) HasGroups() bool`

HasGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


