# RoleRef

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewRoleRef

`func NewRoleRef(name string, ) *RoleRef`

NewRoleRef instantiates a new RoleRef object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoleRefWithDefaults

`func NewRoleRefWithDefaults() *RoleRef`

NewRoleRefWithDefaults instantiates a new RoleRef object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RoleRef) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RoleRef) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RoleRef) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RoleRef) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *RoleRef) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RoleRef) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RoleRef) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *RoleRef) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RoleRef) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RoleRef) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RoleRef) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


