# KafkaRoleBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Role** | **string** |  | 
**User** | **string** |  | 
**CreateResourcePatterns** | Pointer to [**[]ResourcePattern**](ResourcePattern.md) |  | [optional] 
**DeleteResourcePatterns** | Pointer to [**[]ResourcePattern**](ResourcePattern.md) |  | [optional] 

## Methods

### NewKafkaRoleBindingRequest

`func NewKafkaRoleBindingRequest(role string, user string, ) *KafkaRoleBindingRequest`

NewKafkaRoleBindingRequest instantiates a new KafkaRoleBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRoleBindingRequestWithDefaults

`func NewKafkaRoleBindingRequestWithDefaults() *KafkaRoleBindingRequest`

NewKafkaRoleBindingRequestWithDefaults instantiates a new KafkaRoleBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaRoleBindingRequest) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaRoleBindingRequest) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaRoleBindingRequest) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaRoleBindingRequest) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *KafkaRoleBindingRequest) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaRoleBindingRequest) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaRoleBindingRequest) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaRoleBindingRequest) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetName

`func (o *KafkaRoleBindingRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRoleBindingRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRoleBindingRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaRoleBindingRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KafkaRoleBindingRequest) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KafkaRoleBindingRequest) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KafkaRoleBindingRequest) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KafkaRoleBindingRequest) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetRole

`func (o *KafkaRoleBindingRequest) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *KafkaRoleBindingRequest) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *KafkaRoleBindingRequest) SetRole(v string)`

SetRole sets Role field to given value.


### GetUser

`func (o *KafkaRoleBindingRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *KafkaRoleBindingRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *KafkaRoleBindingRequest) SetUser(v string)`

SetUser sets User field to given value.


### GetCreateResourcePatterns

`func (o *KafkaRoleBindingRequest) GetCreateResourcePatterns() []ResourcePattern`

GetCreateResourcePatterns returns the CreateResourcePatterns field if non-nil, zero value otherwise.

### GetCreateResourcePatternsOk

`func (o *KafkaRoleBindingRequest) GetCreateResourcePatternsOk() (*[]ResourcePattern, bool)`

GetCreateResourcePatternsOk returns a tuple with the CreateResourcePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateResourcePatterns

`func (o *KafkaRoleBindingRequest) SetCreateResourcePatterns(v []ResourcePattern)`

SetCreateResourcePatterns sets CreateResourcePatterns field to given value.

### HasCreateResourcePatterns

`func (o *KafkaRoleBindingRequest) HasCreateResourcePatterns() bool`

HasCreateResourcePatterns returns a boolean if a field has been set.

### GetDeleteResourcePatterns

`func (o *KafkaRoleBindingRequest) GetDeleteResourcePatterns() []ResourcePattern`

GetDeleteResourcePatterns returns the DeleteResourcePatterns field if non-nil, zero value otherwise.

### GetDeleteResourcePatternsOk

`func (o *KafkaRoleBindingRequest) GetDeleteResourcePatternsOk() (*[]ResourcePattern, bool)`

GetDeleteResourcePatternsOk returns a tuple with the DeleteResourcePatterns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteResourcePatterns

`func (o *KafkaRoleBindingRequest) SetDeleteResourcePatterns(v []ResourcePattern)`

SetDeleteResourcePatterns sets DeleteResourcePatterns field to given value.

### HasDeleteResourcePatterns

`func (o *KafkaRoleBindingRequest) HasDeleteResourcePatterns() bool`

HasDeleteResourcePatterns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


