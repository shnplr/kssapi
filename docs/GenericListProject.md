# GenericListProject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]Project**](Project.md) |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 

## Methods

### NewGenericListProject

`func NewGenericListProject() *GenericListProject`

NewGenericListProject instantiates a new GenericListProject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericListProjectWithDefaults

`func NewGenericListProjectWithDefaults() *GenericListProject`

NewGenericListProjectWithDefaults instantiates a new GenericListProject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *GenericListProject) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GenericListProject) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GenericListProject) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GenericListProject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *GenericListProject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenericListProject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenericListProject) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenericListProject) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *GenericListProject) GetItems() []Project`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GenericListProject) GetItemsOk() (*[]Project, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GenericListProject) SetItems(v []Project)`

SetItems sets Items field to given value.

### HasItems

`func (o *GenericListProject) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetKind

`func (o *GenericListProject) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GenericListProject) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GenericListProject) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *GenericListProject) HasKind() bool`

HasKind returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


