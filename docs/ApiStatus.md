# ApiStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 
**DetailedMessage** | Pointer to **string** |  | [optional] 

## Methods

### NewApiStatus

`func NewApiStatus() *ApiStatus`

NewApiStatus instantiates a new ApiStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiStatusWithDefaults

`func NewApiStatusWithDefaults() *ApiStatus`

NewApiStatusWithDefaults instantiates a new ApiStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ApiStatus) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiStatus) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiStatus) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApiStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMessage

`func (o *ApiStatus) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiStatus) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiStatus) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ApiStatus) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetReason

`func (o *ApiStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ApiStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ApiStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ApiStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCode

`func (o *ApiStatus) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiStatus) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiStatus) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiStatus) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetailedMessage

`func (o *ApiStatus) GetDetailedMessage() string`

GetDetailedMessage returns the DetailedMessage field if non-nil, zero value otherwise.

### GetDetailedMessageOk

`func (o *ApiStatus) GetDetailedMessageOk() (*string, bool)`

GetDetailedMessageOk returns a tuple with the DetailedMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedMessage

`func (o *ApiStatus) SetDetailedMessage(v string)`

SetDetailedMessage sets DetailedMessage field to given value.

### HasDetailedMessage

`func (o *ApiStatus) HasDetailedMessage() bool`

HasDetailedMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


