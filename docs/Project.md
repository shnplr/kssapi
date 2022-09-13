# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**AdminUsers** | Pointer to **[]string** |  | [optional] 
**EditUsers** | Pointer to **[]string** |  | [optional] 
**ViewUsers** | Pointer to **[]string** |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Project) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Project) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Project) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Project) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAdminUsers

`func (o *Project) GetAdminUsers() []string`

GetAdminUsers returns the AdminUsers field if non-nil, zero value otherwise.

### GetAdminUsersOk

`func (o *Project) GetAdminUsersOk() (*[]string, bool)`

GetAdminUsersOk returns a tuple with the AdminUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminUsers

`func (o *Project) SetAdminUsers(v []string)`

SetAdminUsers sets AdminUsers field to given value.

### HasAdminUsers

`func (o *Project) HasAdminUsers() bool`

HasAdminUsers returns a boolean if a field has been set.

### GetEditUsers

`func (o *Project) GetEditUsers() []string`

GetEditUsers returns the EditUsers field if non-nil, zero value otherwise.

### GetEditUsersOk

`func (o *Project) GetEditUsersOk() (*[]string, bool)`

GetEditUsersOk returns a tuple with the EditUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditUsers

`func (o *Project) SetEditUsers(v []string)`

SetEditUsers sets EditUsers field to given value.

### HasEditUsers

`func (o *Project) HasEditUsers() bool`

HasEditUsers returns a boolean if a field has been set.

### GetViewUsers

`func (o *Project) GetViewUsers() []string`

GetViewUsers returns the ViewUsers field if non-nil, zero value otherwise.

### GetViewUsersOk

`func (o *Project) GetViewUsersOk() (*[]string, bool)`

GetViewUsersOk returns a tuple with the ViewUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewUsers

`func (o *Project) SetViewUsers(v []string)`

SetViewUsers sets ViewUsers field to given value.

### HasViewUsers

`func (o *Project) HasViewUsers() bool`

HasViewUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


