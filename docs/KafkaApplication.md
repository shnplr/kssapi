# KafkaApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**Principal** | **string** |  | 
**Transactional** | Pointer to **bool** |  | [optional] 
**Idempotent** | Pointer to **bool** |  | [optional] 

## Methods

### NewKafkaApplication

`func NewKafkaApplication(type_ string, principal string, ) *KafkaApplication`

NewKafkaApplication instantiates a new KafkaApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaApplicationWithDefaults

`func NewKafkaApplicationWithDefaults() *KafkaApplication`

NewKafkaApplicationWithDefaults instantiates a new KafkaApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaApplication) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaApplication) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaApplication) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaApplication) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *KafkaApplication) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaApplication) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaApplication) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaApplication) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetName

`func (o *KafkaApplication) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaApplication) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaApplication) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaApplication) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KafkaApplication) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KafkaApplication) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KafkaApplication) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KafkaApplication) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetType

`func (o *KafkaApplication) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *KafkaApplication) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *KafkaApplication) SetType(v string)`

SetType sets Type field to given value.


### GetPrincipal

`func (o *KafkaApplication) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *KafkaApplication) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *KafkaApplication) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.


### GetTransactional

`func (o *KafkaApplication) GetTransactional() bool`

GetTransactional returns the Transactional field if non-nil, zero value otherwise.

### GetTransactionalOk

`func (o *KafkaApplication) GetTransactionalOk() (*bool, bool)`

GetTransactionalOk returns a tuple with the Transactional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactional

`func (o *KafkaApplication) SetTransactional(v bool)`

SetTransactional sets Transactional field to given value.

### HasTransactional

`func (o *KafkaApplication) HasTransactional() bool`

HasTransactional returns a boolean if a field has been set.

### GetIdempotent

`func (o *KafkaApplication) GetIdempotent() bool`

GetIdempotent returns the Idempotent field if non-nil, zero value otherwise.

### GetIdempotentOk

`func (o *KafkaApplication) GetIdempotentOk() (*bool, bool)`

GetIdempotentOk returns a tuple with the Idempotent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotent

`func (o *KafkaApplication) SetIdempotent(v bool)`

SetIdempotent sets Idempotent field to given value.

### HasIdempotent

`func (o *KafkaApplication) HasIdempotent() bool`

HasIdempotent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


