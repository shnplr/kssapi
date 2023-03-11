# ClusterRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Rules** | Pointer to [**[]Rule**](Rule.md) |  | [optional] 

## Methods

### NewClusterRole

`func NewClusterRole(name string, ) *ClusterRole`

NewClusterRole instantiates a new ClusterRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleWithDefaults

`func NewClusterRoleWithDefaults() *ClusterRole`

NewClusterRoleWithDefaults instantiates a new ClusterRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterRole) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterRole) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterRole) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterRole) HasKind() bool`

HasKind returns a boolean if a field has been set.

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


### GetRules

`func (o *ClusterRole) GetRules() []Rule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ClusterRole) GetRulesOk() (*[]Rule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ClusterRole) SetRules(v []Rule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ClusterRole) HasRules() bool`

HasRules returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


