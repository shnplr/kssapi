# Role

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Actions** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRole

`func NewRole() *Role`

NewRole instantiates a new Role object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleWithDefaults

`func NewRoleWithDefaults() *Role`

NewRoleWithDefaults instantiates a new Role object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *Role) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Role) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Role) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Role) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *Role) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Role) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Role) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Role) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetActions

`func (o *Role) GetActions() []string`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *Role) GetActionsOk() (*[]string, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *Role) SetActions(v []string)`

SetActions sets Actions field to given value.

### HasActions

`func (o *Role) HasActions() bool`

HasActions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


