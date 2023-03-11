# StatusDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Causes** | Pointer to [**[]StatusCause**](StatusCause.md) |  | [optional] 
**RetryAfterSeconds** | Pointer to **int32** |  | [optional] 

## Methods

### NewStatusDetails

`func NewStatusDetails() *StatusDetails`

NewStatusDetails instantiates a new StatusDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusDetailsWithDefaults

`func NewStatusDetailsWithDefaults() *StatusDetails`

NewStatusDetailsWithDefaults instantiates a new StatusDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *StatusDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatusDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatusDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatusDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKind

`func (o *StatusDetails) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *StatusDetails) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *StatusDetails) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *StatusDetails) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetCauses

`func (o *StatusDetails) GetCauses() []StatusCause`

GetCauses returns the Causes field if non-nil, zero value otherwise.

### GetCausesOk

`func (o *StatusDetails) GetCausesOk() (*[]StatusCause, bool)`

GetCausesOk returns a tuple with the Causes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCauses

`func (o *StatusDetails) SetCauses(v []StatusCause)`

SetCauses sets Causes field to given value.

### HasCauses

`func (o *StatusDetails) HasCauses() bool`

HasCauses returns a boolean if a field has been set.

### GetRetryAfterSeconds

`func (o *StatusDetails) GetRetryAfterSeconds() int32`

GetRetryAfterSeconds returns the RetryAfterSeconds field if non-nil, zero value otherwise.

### GetRetryAfterSecondsOk

`func (o *StatusDetails) GetRetryAfterSecondsOk() (*int32, bool)`

GetRetryAfterSecondsOk returns a tuple with the RetryAfterSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryAfterSeconds

`func (o *StatusDetails) SetRetryAfterSeconds(v int32)`

SetRetryAfterSeconds sets RetryAfterSeconds field to given value.

### HasRetryAfterSeconds

`func (o *StatusDetails) HasRetryAfterSeconds() bool`

HasRetryAfterSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


