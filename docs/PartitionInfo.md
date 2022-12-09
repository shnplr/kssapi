# PartitionInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Topic** | Pointer to **string** |  | [optional] 
**Partition** | Pointer to **int32** |  | [optional] 
**Leader** | Pointer to **int32** |  | [optional] 
**Replicas** | Pointer to **[]int32** |  | [optional] 
**Isr** | Pointer to **[]int32** |  | [optional] 

## Methods

### NewPartitionInfo

`func NewPartitionInfo() *PartitionInfo`

NewPartitionInfo instantiates a new PartitionInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartitionInfoWithDefaults

`func NewPartitionInfoWithDefaults() *PartitionInfo`

NewPartitionInfoWithDefaults instantiates a new PartitionInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTopic

`func (o *PartitionInfo) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *PartitionInfo) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *PartitionInfo) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *PartitionInfo) HasTopic() bool`

HasTopic returns a boolean if a field has been set.

### GetPartition

`func (o *PartitionInfo) GetPartition() int32`

GetPartition returns the Partition field if non-nil, zero value otherwise.

### GetPartitionOk

`func (o *PartitionInfo) GetPartitionOk() (*int32, bool)`

GetPartitionOk returns a tuple with the Partition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartition

`func (o *PartitionInfo) SetPartition(v int32)`

SetPartition sets Partition field to given value.

### HasPartition

`func (o *PartitionInfo) HasPartition() bool`

HasPartition returns a boolean if a field has been set.

### GetLeader

`func (o *PartitionInfo) GetLeader() int32`

GetLeader returns the Leader field if non-nil, zero value otherwise.

### GetLeaderOk

`func (o *PartitionInfo) GetLeaderOk() (*int32, bool)`

GetLeaderOk returns a tuple with the Leader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeader

`func (o *PartitionInfo) SetLeader(v int32)`

SetLeader sets Leader field to given value.

### HasLeader

`func (o *PartitionInfo) HasLeader() bool`

HasLeader returns a boolean if a field has been set.

### GetReplicas

`func (o *PartitionInfo) GetReplicas() []int32`

GetReplicas returns the Replicas field if non-nil, zero value otherwise.

### GetReplicasOk

`func (o *PartitionInfo) GetReplicasOk() (*[]int32, bool)`

GetReplicasOk returns a tuple with the Replicas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplicas

`func (o *PartitionInfo) SetReplicas(v []int32)`

SetReplicas sets Replicas field to given value.

### HasReplicas

`func (o *PartitionInfo) HasReplicas() bool`

HasReplicas returns a boolean if a field has been set.

### GetIsr

`func (o *PartitionInfo) GetIsr() []int32`

GetIsr returns the Isr field if non-nil, zero value otherwise.

### GetIsrOk

`func (o *PartitionInfo) GetIsrOk() (*[]int32, bool)`

GetIsrOk returns a tuple with the Isr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsr

`func (o *PartitionInfo) SetIsr(v []int32)`

SetIsr sets Isr field to given value.

### HasIsr

`func (o *PartitionInfo) HasIsr() bool`

HasIsr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


