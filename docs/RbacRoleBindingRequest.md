# RbacRoleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 

## Methods

### NewRbacRoleBindingRequest

`func NewRbacRoleBindingRequest() *RbacRoleBindingRequest`

NewRbacRoleBindingRequest instantiates a new RbacRoleBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacRoleBindingRequestWithDefaults

`func NewRbacRoleBindingRequestWithDefaults() *RbacRoleBindingRequest`

NewRbacRoleBindingRequestWithDefaults instantiates a new RbacRoleBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetResource

`func (o *RbacRoleBindingRequest) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *RbacRoleBindingRequest) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *RbacRoleBindingRequest) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *RbacRoleBindingRequest) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


