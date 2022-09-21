# KafkaRbacRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Principal** | Pointer to **string** |  | [optional] 
**RoleName** | Pointer to **string** |  | [optional] 
**Topic** | Pointer to **string** |  | [optional] 

## Methods

### NewKafkaRbacRequest

`func NewKafkaRbacRequest() *KafkaRbacRequest`

NewKafkaRbacRequest instantiates a new KafkaRbacRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRbacRequestWithDefaults

`func NewKafkaRbacRequestWithDefaults() *KafkaRbacRequest`

NewKafkaRbacRequestWithDefaults instantiates a new KafkaRbacRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrincipal

`func (o *KafkaRbacRequest) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *KafkaRbacRequest) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *KafkaRbacRequest) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *KafkaRbacRequest) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRoleName

`func (o *KafkaRbacRequest) GetRoleName() string`

GetRoleName returns the RoleName field if non-nil, zero value otherwise.

### GetRoleNameOk

`func (o *KafkaRbacRequest) GetRoleNameOk() (*string, bool)`

GetRoleNameOk returns a tuple with the RoleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleName

`func (o *KafkaRbacRequest) SetRoleName(v string)`

SetRoleName sets RoleName field to given value.

### HasRoleName

`func (o *KafkaRbacRequest) HasRoleName() bool`

HasRoleName returns a boolean if a field has been set.

### GetTopic

`func (o *KafkaRbacRequest) GetTopic() string`

GetTopic returns the Topic field if non-nil, zero value otherwise.

### GetTopicOk

`func (o *KafkaRbacRequest) GetTopicOk() (*string, bool)`

GetTopicOk returns a tuple with the Topic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopic

`func (o *KafkaRbacRequest) SetTopic(v string)`

SetTopic sets Topic field to given value.

### HasTopic

`func (o *KafkaRbacRequest) HasTopic() bool`

HasTopic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


