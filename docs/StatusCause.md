# StatusCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Field** | Pointer to **string** |  | [optional] 

## Methods

### NewStatusCause

`func NewStatusCause() *StatusCause`

NewStatusCause instantiates a new StatusCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCauseWithDefaults

`func NewStatusCauseWithDefaults() *StatusCause`

NewStatusCauseWithDefaults instantiates a new StatusCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *StatusCause) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *StatusCause) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *StatusCause) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *StatusCause) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetMessage

`func (o *StatusCause) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *StatusCause) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *StatusCause) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *StatusCause) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetField

`func (o *StatusCause) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *StatusCause) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *StatusCause) SetField(v string)`

SetField sets Field field to given value.

### HasField

`func (o *StatusCause) HasField() bool`

HasField returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


