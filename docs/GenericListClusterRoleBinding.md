# GenericListClusterRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ClusterRoleBinding**](ClusterRoleBinding.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericListClusterRoleBinding

`func NewGenericListClusterRoleBinding() *GenericListClusterRoleBinding`

NewGenericListClusterRoleBinding instantiates a new GenericListClusterRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericListClusterRoleBindingWithDefaults

`func NewGenericListClusterRoleBindingWithDefaults() *GenericListClusterRoleBinding`

NewGenericListClusterRoleBindingWithDefaults instantiates a new GenericListClusterRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GenericListClusterRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GenericListClusterRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GenericListClusterRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GenericListClusterRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetItems

`func (o *GenericListClusterRoleBinding) GetItems() []ClusterRoleBinding`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GenericListClusterRoleBinding) GetItemsOk() (*[]ClusterRoleBinding, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GenericListClusterRoleBinding) SetItems(v []ClusterRoleBinding)`

SetItems sets Items field to given value.

### HasItems

`func (o *GenericListClusterRoleBinding) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetName

`func (o *GenericListClusterRoleBinding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericListClusterRoleBinding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericListClusterRoleBinding) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericListClusterRoleBinding) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *GenericListClusterRoleBinding) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *GenericListClusterRoleBinding) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *GenericListClusterRoleBinding) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *GenericListClusterRoleBinding) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


