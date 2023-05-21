# KafkaApplicationList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]KafkaApplication**](KafkaApplication.md) |  | [optional] 

## Methods

### NewKafkaApplicationList

`func NewKafkaApplicationList() *KafkaApplicationList`

NewKafkaApplicationList instantiates a new KafkaApplicationList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaApplicationListWithDefaults

`func NewKafkaApplicationListWithDefaults() *KafkaApplicationList`

NewKafkaApplicationListWithDefaults instantiates a new KafkaApplicationList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaApplicationList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaApplicationList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaApplicationList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaApplicationList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *KafkaApplicationList) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaApplicationList) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaApplicationList) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaApplicationList) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetItems

`func (o *KafkaApplicationList) GetItems() []KafkaApplication`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *KafkaApplicationList) GetItemsOk() (*[]KafkaApplication, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *KafkaApplicationList) SetItems(v []KafkaApplication)`

SetItems sets Items field to given value.

### HasItems

`func (o *KafkaApplicationList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


