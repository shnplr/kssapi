# KafkaTopicData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**PartitionCount** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**Configs** | Pointer to [**[]TopicConfigData**](TopicConfigData.md) |  | [optional] 
**Partitions** | Pointer to [**[]PartitionInfo**](PartitionInfo.md) |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 

## Methods

### NewKafkaTopicData

`func NewKafkaTopicData() *KafkaTopicData`

NewKafkaTopicData instantiates a new KafkaTopicData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaTopicDataWithDefaults

`func NewKafkaTopicDataWithDefaults() *KafkaTopicData`

NewKafkaTopicDataWithDefaults instantiates a new KafkaTopicData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaTopicData) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaTopicData) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaTopicData) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaTopicData) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *KafkaTopicData) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaTopicData) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaTopicData) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaTopicData) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetName

`func (o *KafkaTopicData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaTopicData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaTopicData) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaTopicData) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KafkaTopicData) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KafkaTopicData) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KafkaTopicData) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KafkaTopicData) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPartitionCount

`func (o *KafkaTopicData) GetPartitionCount() int32`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *KafkaTopicData) GetPartitionCountOk() (*int32, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *KafkaTopicData) SetPartitionCount(v int32)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *KafkaTopicData) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *KafkaTopicData) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *KafkaTopicData) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *KafkaTopicData) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *KafkaTopicData) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetConfigs

`func (o *KafkaTopicData) GetConfigs() []TopicConfigData`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *KafkaTopicData) GetConfigsOk() (*[]TopicConfigData, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *KafkaTopicData) SetConfigs(v []TopicConfigData)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *KafkaTopicData) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetPartitions

`func (o *KafkaTopicData) GetPartitions() []PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *KafkaTopicData) GetPartitionsOk() (*[]PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *KafkaTopicData) SetPartitions(v []PartitionInfo)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *KafkaTopicData) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.

### GetInternal

`func (o *KafkaTopicData) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *KafkaTopicData) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *KafkaTopicData) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *KafkaTopicData) HasInternal() bool`

HasInternal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


