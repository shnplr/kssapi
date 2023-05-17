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

// checks if the NonResourceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NonResourceAttributes{}

// NonResourceAttributes struct for NonResourceAttributes
type NonResourceAttributes struct {
	Path *string `json:"path,omitempty"`
	Verb *string `json:"verb,omitempty"`
}

// NewNonResourceAttributes instantiates a new NonResourceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNonResourceAttributes() *NonResourceAttributes {
	this := NonResourceAttributes{}
	return &this
}

// NewNonResourceAttributesWithDefaults instantiates a new NonResourceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNonResourceAttributesWithDefaults() *NonResourceAttributes {
	this := NonResourceAttributes{}
	return &this
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *NonResourceAttributes) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonResourceAttributes) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *NonResourceAttributes) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *NonResourceAttributes) SetPath(v string) {
	o.Path = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *NonResourceAttributes) GetVerb() string {
	if o == nil || IsNil(o.Verb) {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NonResourceAttributes) GetVerbOk() (*string, bool) {
	if o == nil || IsNil(o.Verb) {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *NonResourceAttributes) HasVerb() bool {
	if o != nil && !IsNil(o.Verb) {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *NonResourceAttributes) SetVerb(v string) {
	o.Verb = &v
}

func (o NonResourceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NonResourceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.Verb) {
		toSerialize["verb"] = o.Verb
	}
	return toSerialize, nil
}

type NullableNonResourceAttributes struct {
	value *NonResourceAttributes
	isSet bool
}

func (v NullableNonResourceAttributes) Get() *NonResourceAttributes {
	return v.value
}

func (v *NullableNonResourceAttributes) Set(val *NonResourceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableNonResourceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableNonResourceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNonResourceAttributes(val *NonResourceAttributes) *NullableNonResourceAttributes {
	return &NullableNonResourceAttributes{value: val, isSet: true}
}

func (v NullableNonResourceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNonResourceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


