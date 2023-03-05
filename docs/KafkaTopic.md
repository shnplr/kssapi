# KafkaTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**PartitionCount** | Pointer to **int32** |  | [optional] 
**ReplicationFactor** | Pointer to **int32** |  | [optional] 
**Configs** | Pointer to [**[]ConfigItem**](ConfigItem.md) |  | [optional] 
**Partitions** | Pointer to [**[]PartitionInfo**](PartitionInfo.md) |  | [optional] 

## Methods

### NewKafkaTopic

`func NewKafkaTopic(name string, namespace string, ) *KafkaTopic`

NewKafkaTopic instantiates a new KafkaTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaTopicWithDefaults

`func NewKafkaTopicWithDefaults() *KafkaTopic`

NewKafkaTopicWithDefaults instantiates a new KafkaTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaTopic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaTopic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaTopic) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaTopic) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *KafkaTopic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaTopic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaTopic) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *KafkaTopic) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KafkaTopic) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KafkaTopic) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPartitionCount

`func (o *KafkaTopic) GetPartitionCount() int32`

GetPartitionCount returns the PartitionCount field if non-nil, zero value otherwise.

### GetPartitionCountOk

`func (o *KafkaTopic) GetPartitionCountOk() (*int32, bool)`

GetPartitionCountOk returns a tuple with the PartitionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitionCount

`func (o *KafkaTopic) SetPartitionCount(v int32)`

SetPartitionCount sets PartitionCount field to given value.

### HasPartitionCount

`func (o *KafkaTopic) HasPartitionCount() bool`

HasPartitionCount returns a boolean if a field has been set.

### GetReplicationFactor

`func (o *KafkaTopic) GetReplicationFactor() int32`

GetReplicationFactor returns the ReplicationFactor field if non-nil, zero value otherwise.

### GetReplicationFactorOk

`func (o *KafkaTopic) GetReplicationFactorOk() (*int32, bool)`

GetReplicationFactorOk returns a tuple with the ReplicationFactor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicationFactor

`func (o *KafkaTopic) SetReplicationFactor(v int32)`

SetReplicationFactor sets ReplicationFactor field to given value.

### HasReplicationFactor

`func (o *KafkaTopic) HasReplicationFactor() bool`

HasReplicationFactor returns a boolean if a field has been set.

### GetConfigs

`func (o *KafkaTopic) GetConfigs() []ConfigItem`

GetConfigs returns the Configs field if non-nil, zero value otherwise.

### GetConfigsOk

`func (o *KafkaTopic) GetConfigsOk() (*[]ConfigItem, bool)`

GetConfigsOk returns a tuple with the Configs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigs

`func (o *KafkaTopic) SetConfigs(v []ConfigItem)`

SetConfigs sets Configs field to given value.

### HasConfigs

`func (o *KafkaTopic) HasConfigs() bool`

HasConfigs returns a boolean if a field has been set.

### GetPartitions

`func (o *KafkaTopic) GetPartitions() []PartitionInfo`

GetPartitions returns the Partitions field if non-nil, zero value otherwise.

### GetPartitionsOk

`func (o *KafkaTopic) GetPartitionsOk() (*[]PartitionInfo, bool)`

GetPartitionsOk returns a tuple with the Partitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartitions

`func (o *KafkaTopic) SetPartitions(v []PartitionInfo)`

SetPartitions sets Partitions field to given value.

### HasPartitions

`func (o *KafkaTopic) HasPartitions() bool`

HasPartitions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


