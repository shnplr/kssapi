# ClusterRoleBindingList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ClusterRoleBinding**](ClusterRoleBinding.md) |  | [optional] 

## Methods

### NewClusterRoleBindingList

`func NewClusterRoleBindingList() *ClusterRoleBindingList`

NewClusterRoleBindingList instantiates a new ClusterRoleBindingList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleBindingListWithDefaults

`func NewClusterRoleBindingListWithDefaults() *ClusterRoleBindingList`

NewClusterRoleBindingListWithDefaults instantiates a new ClusterRoleBindingList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterRoleBindingList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterRoleBindingList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterRoleBindingList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterRoleBindingList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *ClusterRoleBindingList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterRoleBindingList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterRoleBindingList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterRoleBindingList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *ClusterRoleBindingList) GetItems() []ClusterRoleBinding`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ClusterRoleBindingList) GetItemsOk() (*[]ClusterRoleBinding, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ClusterRoleBindingList) SetItems(v []ClusterRoleBinding)`

SetItems sets Items field to given value.

### HasItems

`func (o *ClusterRoleBindingList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


