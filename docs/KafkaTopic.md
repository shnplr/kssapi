# KafkaTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**TopicData**](TopicData.md) |  | [optional] 

## Methods

### NewKafkaTopic

`func NewKafkaTopic() *KafkaTopic`

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

### HasName

`func (o *KafkaTopic) HasName() bool`

HasName returns a boolean if a field has been set.

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

### HasNamespace

`func (o *KafkaTopic) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetData

`func (o *KafkaTopic) GetData() TopicData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *KafkaTopic) GetDataOk() (*TopicData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *KafkaTopic) SetData(v TopicData)`

SetData sets Data field to given value.

### HasData

`func (o *KafkaTopic) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


