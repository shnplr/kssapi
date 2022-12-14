/*
kafka-self-service API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0-SNAPSHOT
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// ApiStatus struct for ApiStatus
type ApiStatus struct {
	Kind *string `json:"kind,omitempty"`
	Status *string `json:"status,omitempty"`
	Message *string `json:"message,omitempty"`
	Reason *string `json:"reason,omitempty"`
	Code *int32 `json:"code,omitempty"`
	DetailedMessage *string `json:"detailedMessage,omitempty"`
}

// NewApiStatus instantiates a new ApiStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiStatus() *ApiStatus {
	this := ApiStatus{}
	return &this
}

// NewApiStatusWithDefaults instantiates a new ApiStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiStatusWithDefaults() *ApiStatus {
	this := ApiStatus{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiStatus) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiStatus) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiStatus) SetKind(v string) {
	o.Kind = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ApiStatus) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ApiStatus) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ApiStatus) SetStatus(v string) {
	o.Status = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ApiStatus) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *ApiStatus) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ApiStatus) SetMessage(v string) {
	o.Message = &v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ApiStatus) GetReason() string {
	if o == nil || isNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetReasonOk() (*string, bool) {
	if o == nil || isNil(o.Reason) {
    return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ApiStatus) HasReason() bool {
	if o != nil && !isNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ApiStatus) SetReason(v string) {
	o.Reason = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *ApiStatus) GetCode() int32 {
	if o == nil || isNil(o.Code) {
		var ret int32
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetCodeOk() (*int32, bool) {
	if o == nil || isNil(o.Code) {
    return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *ApiStatus) HasCode() bool {
	if o != nil && !isNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given int32 and assigns it to the Code field.
func (o *ApiStatus) SetCode(v int32) {
	o.Code = &v
}

// GetDetailedMessage returns the DetailedMessage field value if set, zero value otherwise.
func (o *ApiStatus) GetDetailedMessage() string {
	if o == nil || isNil(o.DetailedMessage) {
		var ret string
		return ret
	}
	return *o.DetailedMessage
}

// GetDetailedMessageOk returns a tuple with the DetailedMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiStatus) GetDetailedMessageOk() (*string, bool) {
	if o == nil || isNil(o.DetailedMessage) {
    return nil, false
	}
	return o.DetailedMessage, true
}

// HasDetailedMessage returns a boolean if a field has been set.
func (o *ApiStatus) HasDetailedMessage() bool {
	if o != nil && !isNil(o.DetailedMessage) {
		return true
	}

	return false
}

// SetDetailedMessage gets a reference to the given string and assigns it to the DetailedMessage field.
func (o *ApiStatus) SetDetailedMessage(v string) {
	o.DetailedMessage = &v
}

func (o ApiStatus) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !isNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !isNil(o.DetailedMessage) {
		toSerialize["detailedMessage"] = o.DetailedMessage
	}
	return json.Marshal(toSerialize)
}

type NullableApiStatus struct {
	value *ApiStatus
	isSet bool
}

func (v NullableApiStatus) Get() *ApiStatus {
	return v.value
}

func (v *NullableApiStatus) Set(val *ApiStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApiStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApiStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiStatus(val *ApiStatus) *NullableApiStatus {
	return &NullableApiStatus{value: val, isSet: true}
}

func (v NullableApiStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


