# ClusterRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Subjects** | Pointer to [**[]Subject**](Subject.md) |  | [optional] 
**Name** | **string** |  | 
**RoleRef** | [**RoleRef**](RoleRef.md) |  | 
**Namespace** | Pointer to **string** |  | [optional] 

## Methods

### NewClusterRoleBinding

`func NewClusterRoleBinding(name string, roleRef RoleRef, ) *ClusterRoleBinding`

NewClusterRoleBinding instantiates a new ClusterRoleBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterRoleBindingWithDefaults

`func NewClusterRoleBindingWithDefaults() *ClusterRoleBinding`

NewClusterRoleBindingWithDefaults instantiates a new ClusterRoleBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *ClusterRoleBinding) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ClusterRoleBinding) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ClusterRoleBinding) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ClusterRoleBinding) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetSubjects

`func (o *ClusterRoleBinding) GetSubjects() []Subject`

GetSubjects returns the Subjects field if non-nil, zero value otherwise.

### GetSubjectsOk

`func (o *ClusterRoleBinding) GetSubjectsOk() (*[]Subject, bool)`

GetSubjectsOk returns a tuple with the Subjects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjects

`func (o *ClusterRoleBinding) SetSubjects(v []Subject)`

SetSubjects sets Subjects field to given value.

### HasSubjects

`func (o *ClusterRoleBinding) HasSubjects() bool`

HasSubjects returns a boolean if a field has been set.

### GetName

`func (o *ClusterRoleBinding) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterRoleBinding) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterRoleBinding) SetName(v string)`

SetName sets Name field to given value.


### GetRoleRef

`func (o *ClusterRoleBinding) GetRoleRef() RoleRef`

GetRoleRef returns the RoleRef field if non-nil, zero value otherwise.

### GetRoleRefOk

`func (o *ClusterRoleBinding) GetRoleRefOk() (*RoleRef, bool)`

GetRoleRefOk returns a tuple with the RoleRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleRef

`func (o *ClusterRoleBinding) SetRoleRef(v RoleRef)`

SetRoleRef sets RoleRef field to given value.


### GetNamespace

`func (o *ClusterRoleBinding) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ClusterRoleBinding) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ClusterRoleBinding) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *ClusterRoleBinding) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


