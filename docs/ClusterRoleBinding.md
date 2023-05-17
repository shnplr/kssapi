# ClusterRoleBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**RoleRef** | [**RoleRef**](RoleRef.md) |  | 
**Subjects** | Pointer to [**[]Subject**](Subject.md) |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 

## Methods

### NewClusterRoleBinding

`func NewClusterRoleBinding(roleRef RoleRef, ) *ClusterRoleBinding`

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

### GetApiVersion

`func (o *ClusterRoleBinding) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *ClusterRoleBinding) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *ClusterRoleBinding) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *ClusterRoleBinding) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

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

### GetMetadata

`func (o *ClusterRoleBinding) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ClusterRoleBinding) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ClusterRoleBinding) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ClusterRoleBinding) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


