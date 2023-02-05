# KafkaStreams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | **string** |  | 
**Namespace** | **string** |  | 
**Principal** | **string** |  | 
**Transactional** | Pointer to **bool** |  | [optional] 
**Idempotent** | Pointer to **bool** |  | [optional] 

## Methods

### NewKafkaStreams

`func NewKafkaStreams(name string, namespace string, principal string, ) *KafkaStreams`

NewKafkaStreams instantiates a new KafkaStreams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaStreamsWithDefaults

`func NewKafkaStreamsWithDefaults() *KafkaStreams`

NewKafkaStreamsWithDefaults instantiates a new KafkaStreams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaStreams) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaStreams) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaStreams) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaStreams) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *KafkaStreams) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaStreams) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaStreams) SetName(v string)`

SetName sets Name field to given value.


### GetNamespace

`func (o *KafkaStreams) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KafkaStreams) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KafkaStreams) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPrincipal

`func (o *KafkaStreams) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *KafkaStreams) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *KafkaStreams) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetTransactional

`func (o *KafkaStreams) GetTransactional() bool`

GetTransactional returns the Transactional field if non-nil, zero value otherwise.

### GetTransactionalOk

`func (o *KafkaStreams) GetTransactionalOk() (*bool, bool)`

GetTransactionalOk returns a tuple with the Transactional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactional

`func (o *KafkaStreams) SetTransactional(v bool)`

SetTransactional sets Transactional field to given value.

### HasTransactional

`func (o *KafkaStreams) HasTransactional() bool`

HasTransactional returns a boolean if a field has been set.

### GetIdempotent

`func (o *KafkaStreams) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *KafkaStreams) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *KafkaStreams) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.

### HasIdempotent

`func (o *KafkaStreams) HasIdempotent() bool`

HasIdempotent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


