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

// checks if the GenericListKafkaTopic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericListKafkaTopic{}

// GenericListKafkaTopic struct for GenericListKafkaTopic
type GenericListKafkaTopic struct {
	Kind *string `json:"kind,omitempty"`
	Items []KafkaTopic `json:"items,omitempty"`
}

// NewGenericListKafkaTopic instantiates a new GenericListKafkaTopic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericListKafkaTopic() *GenericListKafkaTopic {
	this := GenericListKafkaTopic{}
	return &this
}

// NewGenericListKafkaTopicWithDefaults instantiates a new GenericListKafkaTopic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericListKafkaTopicWithDefaults() *GenericListKafkaTopic {
	this := GenericListKafkaTopic{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GenericListKafkaTopic) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListKafkaTopic) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GenericListKafkaTopic) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GenericListKafkaTopic) SetKind(v string) {
	o.Kind = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GenericListKafkaTopic) GetItems() []KafkaTopic {
	if o == nil || IsNil(o.Items) {
		var ret []KafkaTopic
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListKafkaTopic) GetItemsOk() ([]KafkaTopic, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GenericListKafkaTopic) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []KafkaTopic and assigns it to the Items field.
func (o *GenericListKafkaTopic) SetItems(v []KafkaTopic) {
	o.Items = v
}

func (o GenericListKafkaTopic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericListKafkaTopic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableGenericListKafkaTopic struct {
	value *GenericListKafkaTopic
	isSet bool
}

func (v NullableGenericListKafkaTopic) Get() *GenericListKafkaTopic {
	return v.value
}

func (v *NullableGenericListKafkaTopic) Set(val *GenericListKafkaTopic) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericListKafkaTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericListKafkaTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericListKafkaTopic(val *GenericListKafkaTopic) *NullableGenericListKafkaTopic {
	return &NullableGenericListKafkaTopic{value: val, isSet: true}
}

func (v NullableGenericListKafkaTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericListKafkaTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


