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

// GenericListKafkaStreams struct for GenericListKafkaStreams
type GenericListKafkaStreams struct {
	Kind *string `json:"kind,omitempty"`
	Items []KafkaStreams `json:"items,omitempty"`
}

// NewGenericListKafkaStreams instantiates a new GenericListKafkaStreams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericListKafkaStreams() *GenericListKafkaStreams {
	this := GenericListKafkaStreams{}
	return &this
}

// NewGenericListKafkaStreamsWithDefaults instantiates a new GenericListKafkaStreams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericListKafkaStreamsWithDefaults() *GenericListKafkaStreams {
	this := GenericListKafkaStreams{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GenericListKafkaStreams) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListKafkaStreams) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GenericListKafkaStreams) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GenericListKafkaStreams) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GenericListKafkaStreams) GetItems() []KafkaStreams {
	if o == nil || isNil(o.Items) {
		var ret []KafkaStreams
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListKafkaStreams) GetItemsOk() ([]KafkaStreams, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GenericListKafkaStreams) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KafkaStreams and assigns it to the Items field.
func (o *GenericListKafkaStreams) SetItems(v []KafkaStreams) {
	o.Items = v
}

func (o GenericListKafkaStreams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return json.Marshal(toSerialize)
}

type NullableGenericListKafkaStreams struct {
	value *GenericListKafkaStreams
	isSet bool
}

func (v NullableGenericListKafkaStreams) Get() *GenericListKafkaStreams {
	return v.value
}

func (v *NullableGenericListKafkaStreams) Set(val *GenericListKafkaStreams) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericListKafkaStreams) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericListKafkaStreams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericListKafkaStreams(val *GenericListKafkaStreams) *NullableGenericListKafkaStreams {
	return &NullableGenericListKafkaStreams{value: val, isSet: true}
}

func (v NullableGenericListKafkaStreams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericListKafkaStreams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

