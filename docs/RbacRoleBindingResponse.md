# RbacRoleBindingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
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

### GetApiVersion

`func (o *RbacRoleBindingResponse) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *RbacRoleBindingResponse) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *RbacRoleBindingResponse) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *RbacRoleBindingResponse) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

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

### GetMetadata

`func (o *RbacRoleBindingResponse) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *RbacRoleBindingResponse) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *RbacRoleBindingResponse) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *RbacRoleBindingResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

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


