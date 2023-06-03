# AlterConfigBatchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Configs** | Pointer to [**[]ConfigItem**](ConfigItem.md) |  | [optional] 

## Methods

### NewAlterConfigBatchRequest

`func NewAlterConfigBatchRequest() *AlterConfigBatchRequest`

NewAlterConfigBatchRequest instantiates a new AlterConfigBatchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAlterConfigBatchRequestWithDefaults

`func NewAlterConfigBatchRequestWithDefaults() *AlterConfigBatchRequest`

NewAlterConfigBatchRequestWithDefaults instantiates a new AlterConfigBatchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *AlterConfigBatchRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *AlterConfigBatchRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *AlterConfigBatchRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *AlterConfigBatchRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *AlterConfigBatchRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *AlterConfigBatchRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *AlterConfigBatchRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *AlterConfigBatchRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetConfigs

`func (o *AlterConfigBatchRequest) GetConfigs() []ConfigItem`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *AlterConfigBatchRequest) GetConfigsOk() (*[]ConfigItem, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *AlterConfigBatchRequest) SetConfigs(v []ConfigItem)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *AlterConfigBatchRequest) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


