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

// checks if the StatusCause type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusCause{}

// StatusCause struct for StatusCause
type StatusCause struct {
	Reason *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
	Field *string `json:"field,omitempty"`
}

// NewStatusCause instantiates a new StatusCause object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusCause() *StatusCause {
	this := StatusCause{}
	return &this
}

// NewStatusCauseWithDefaults instantiates a new StatusCause object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusCauseWithDefaults() *StatusCause {
	this := StatusCause{}
	return &this
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *StatusCause) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCause) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *StatusCause) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *StatusCause) SetReason(v string) {
	o.Reason = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *StatusCause) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCause) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *StatusCause) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *StatusCause) SetMessage(v string) {
	o.Message = &v
}

// GetField returns the Field field value if set, zero value otherwise.
func (o *StatusCause) GetField() string {
	if o == nil || IsNil(o.Field) {
		var ret string
		return ret
	}
	return *o.Field
}

// GetFieldOk returns a tuple with the Field field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusCause) GetFieldOk() (*string, bool) {
	if o == nil || IsNil(o.Field) {
		return nil, false
	}
	return o.Field, true
}

// HasField returns a boolean if a field has been set.
func (o *StatusCause) HasField() bool {
	if o != nil && !IsNil(o.Field) {
		return true
	}

	return false
}

// SetField gets a reference to the given string and assigns it to the Field field.
func (o *StatusCause) SetField(v string) {
	o.Field = &v
}

func (o StatusCause) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusCause) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Field) {
		toSerialize["field"] = o.Field
	}
	return toSerialize, nil
}

type NullableStatusCause struct {
	value *StatusCause
	isSet bool
}

func (v NullableStatusCause) Get() *StatusCause {
	return v.value
}

func (v *NullableStatusCause) Set(val *StatusCause) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusCause) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusCause(val *StatusCause) *NullableStatusCause {
	return &NullableStatusCause{value: val, isSet: true}
}

func (v NullableStatusCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


