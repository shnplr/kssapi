# GenericListRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]Role**](Role.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

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

### GetName

`func (o *GenericListRole) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericListRole) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericListRole) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericListRole) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *GenericListRole) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GenericListRole) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GenericListRole) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GenericListRole) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


