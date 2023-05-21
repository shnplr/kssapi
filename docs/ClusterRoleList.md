# ClusterRoleList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ClusterRole**](ClusterRole.md) |  | [optional] 

## Methods

### NewClusterRoleList

`func NewClusterRoleList() *ClusterRoleList`

NewClusterRoleList instantiates a new ClusterRoleList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleListWithDefaults

`func NewClusterRoleListWithDefaults() *ClusterRoleList`

NewClusterRoleListWithDefaults instantiates a new ClusterRoleList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterRoleList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterRoleList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterRoleList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterRoleList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *ClusterRoleList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterRoleList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterRoleList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterRoleList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ClusterRoleList) GetItems() []ClusterRole`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterRoleList) GetItemsOk() (*[]ClusterRole, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterRoleList) SetItems(v []ClusterRole)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterRoleList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


