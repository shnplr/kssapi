# RbacRoleBindingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**Resources** | Pointer to [**[]ResourcePattern**](ResourcePattern.md) |  | [optional] 

## Methods

### NewRbacRoleBindingResponse

`func NewRbacRoleBindingResponse() *RbacRoleBindingResponse`

NewRbacRoleBindingResponse instantiates a new RbacRoleBindingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRbacRoleBindingResponseWithDefaults

`func NewRbacRoleBindingResponseWithDefaults() *RbacRoleBindingResponse`

NewRbacRoleBindingResponseWithDefaults instantiates a new RbacRoleBindingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *RbacRoleBindingResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *RbacRoleBindingResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *RbacRoleBindingResponse) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *RbacRoleBindingResponse) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *RbacRoleBindingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RbacRoleBindingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RbacRoleBindingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *RbacRoleBindingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *RbacRoleBindingResponse) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *RbacRoleBindingResponse) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *RbacRoleBindingResponse) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *RbacRoleBindingResponse) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPrincipal

`func (o *RbacRoleBindingResponse) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *RbacRoleBindingResponse) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *RbacRoleBindingResponse) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *RbacRoleBindingResponse) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetResources

`func (o *RbacRoleBindingResponse) GetResources() []ResourcePattern`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *RbacRoleBindingResponse) GetResourcesOk() (*[]ResourcePattern, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *RbacRoleBindingResponse) SetResources(v []ResourcePattern)`

SetResources sets Resources field to given value.

### HasResources

`func (o *RbacRoleBindingResponse) HasResources() bool`

HasResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


