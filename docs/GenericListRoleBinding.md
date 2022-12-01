# GenericListRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Items** | Pointer to [**[]RoleBinding**](RoleBinding.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericListRoleBinding

`func NewGenericListRoleBinding() *GenericListRoleBinding`

NewGenericListRoleBinding instantiates a new GenericListRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericListRoleBindingWithDefaults

`func NewGenericListRoleBindingWithDefaults() *GenericListRoleBinding`

NewGenericListRoleBindingWithDefaults instantiates a new GenericListRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GenericListRoleBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenericListRoleBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenericListRoleBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenericListRoleBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetItems

`func (o *GenericListRoleBinding) GetItems() []RoleBinding`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GenericListRoleBinding) GetItemsOk() (*[]RoleBinding, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GenericListRoleBinding) SetItems(v []RoleBinding)`

SetItems sets Items field to given value.

### HasItems

`func (o *GenericListRoleBinding) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *GenericListRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GenericListRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GenericListRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GenericListRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


