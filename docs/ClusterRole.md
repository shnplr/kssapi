# ClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Actions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClusterRole

`func NewClusterRole() *ClusterRole`

NewClusterRole instantiates a new ClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleWithDefaults

`func NewClusterRoleWithDefaults() *ClusterRole`

NewClusterRoleWithDefaults instantiates a new ClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ClusterRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActions

`func (o *ClusterRole) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ClusterRole) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ClusterRole) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ClusterRole) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

