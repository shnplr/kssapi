# GenericListClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ClusterRole**](ClusterRole.md) |  | [optional] 

## Methods

### NewGenericListClusterRole

`func NewGenericListClusterRole() *GenericListClusterRole`

NewGenericListClusterRole instantiates a new GenericListClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericListClusterRoleWithDefaults

`func NewGenericListClusterRoleWithDefaults() *GenericListClusterRole`

NewGenericListClusterRoleWithDefaults instantiates a new GenericListClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GenericListClusterRole) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GenericListClusterRole) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GenericListClusterRole) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GenericListClusterRole) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *GenericListClusterRole) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *GenericListClusterRole) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *GenericListClusterRole) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *GenericListClusterRole) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *GenericListClusterRole) GetItems() []ClusterRole`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GenericListClusterRole) GetItemsOk() (*[]ClusterRole, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GenericListClusterRole) SetItems(v []ClusterRole)`

SetItems sets Items field to given value.

### HasItems

`func (o *GenericListClusterRole) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


