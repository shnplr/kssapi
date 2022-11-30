# ClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**ObjectMeta**](ObjectMeta.md) |  | 
**Actions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewClusterRole

`func NewClusterRole(metadata ObjectMeta, ) *ClusterRole`

NewClusterRole instantiates a new ClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleWithDefaults

`func NewClusterRoleWithDefaults() *ClusterRole`

NewClusterRoleWithDefaults instantiates a new ClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *ClusterRole) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterRole) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterRole) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.


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


