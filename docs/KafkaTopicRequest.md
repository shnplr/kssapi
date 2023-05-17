# KafkaTopicRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**PartitionCount** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**Configs** | Pointer to [**[]ConfigItem**](ConfigItem.md) |  | [optional] 
**Metadata** | Pointer to [**ObjectMeta**](ObjectMeta.md) |  | [optional] 

## Methods

### NewKafkaTopicRequest

`func NewKafkaTopicRequest() *KafkaTopicRequest`

NewKafkaTopicRequest instantiates a new KafkaTopicRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaTopicRequestWithDefaults

`func NewKafkaTopicRequestWithDefaults() *KafkaTopicRequest`

NewKafkaTopicRequestWithDefaults instantiates a new KafkaTopicRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaTopicRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaTopicRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaTopicRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaTopicRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *KafkaTopicRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaTopicRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaTopicRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaTopicRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetPartitionCount

`func (o *KafkaTopicRequest) GetPartitionCount() int32`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *KafkaTopicRequest) GetPartitionCountOk() (*int32, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *KafkaTopicRequest) SetPartitionCount(v int32)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *KafkaTopicRequest) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *KafkaTopicRequest) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *KafkaTopicRequest) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *KafkaTopicRequest) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *KafkaTopicRequest) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetConfigs

`func (o *KafkaTopicRequest) GetConfigs() []ConfigItem`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *KafkaTopicRequest) GetConfigsOk() (*[]ConfigItem, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *KafkaTopicRequest) SetConfigs(v []ConfigItem)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *KafkaTopicRequest) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetMetadata

`func (o *KafkaTopicRequest) GetMetadata() ObjectMeta`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *KafkaTopicRequest) GetMetadataOk() (*ObjectMeta, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *KafkaTopicRequest) SetMetadata(v ObjectMeta)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *KafkaTopicRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


