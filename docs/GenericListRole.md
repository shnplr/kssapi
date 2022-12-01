# GenericListRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Items** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericListRole

`func NewGenericListRole() *GenericListRole`

NewGenericListRole instantiates a new GenericListRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericListRoleWithDefaults

`func NewGenericListRoleWithDefaults() *GenericListRole`

NewGenericListRoleWithDefaults instantiates a new GenericListRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GenericListRole) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenericListRole) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenericListRole) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenericListRole) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *GenericListRole) GetItems() []Role`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GenericListRole) GetItemsOk() (*[]Role, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GenericListRole) SetItems(v []Role)`

SetItems sets Items field to given value.

### HasItems

`func (o *GenericListRole) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *GenericListRole) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GenericListRole) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GenericListRole) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GenericListRole) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


