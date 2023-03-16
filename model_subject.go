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

// checks if the Subject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Subject{}

// Subject struct for Subject
type Subject struct {
	Kind *string `json:"kind,omitempty"`
	Name string `json:"name"`
}

// NewSubject instantiates a new Subject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubject(name string) *Subject {
	this := Subject{}
	this.Name = name
	return &this
}

// NewSubjectWithDefaults instantiates a new Subject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectWithDefaults() *Subject {
	this := Subject{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *Subject) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Subject) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *Subject) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *Subject) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value
func (o *Subject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Subject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Subject) SetName(v string) {
	o.Name = v
}

func (o Subject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Subject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableSubject struct {
	value *Subject
	isSet bool
}

func (v NullableSubject) Get() *Subject {
	return v.value
}

func (v *NullableSubject) Set(val *Subject) {
	v.value = val
	v.isSet = true
}

func (v NullableSubject) IsSet() bool {
	return v.isSet
}

func (v *NullableSubject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubject(val *Subject) *NullableSubject {
	return &NullableSubject{value: val, isSet: true}
}

func (v NullableSubject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


