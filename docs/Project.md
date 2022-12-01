# Project

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 

## Methods

### NewProject

`func NewProject() *Project`

NewProject instantiates a new Project object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDefaults

`func NewProjectWithDefaults() *Project`

NewProjectWithDefaults instantiates a new Project object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *Project) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *Project) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *Project) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *Project) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetMetadata

`func (o *Project) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Project) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Project) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Project) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetDescription

`func (o *Project) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Project) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Project) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Project) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


