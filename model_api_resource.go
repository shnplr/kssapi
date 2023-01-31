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

// ApiResource struct for ApiResource
type ApiResource struct {
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespaced *bool `json:"namespaced,omitempty"`
	Verbs []string `json:"verbs,omitempty"`
}

// NewApiResource instantiates a new ApiResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResource() *ApiResource {
	this := ApiResource{}
	return &this
}

// NewApiResourceWithDefaults instantiates a new ApiResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResourceWithDefaults() *ApiResource {
	this := ApiResource{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiResource) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResource) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiResource) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiResource) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ApiResource) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResource) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ApiResource) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ApiResource) SetName(v string) {
	o.Name = &v
}

// GetNamespaced returns the Namespaced field value if set, zero value otherwise.
func (o *ApiResource) GetNamespaced() bool {
	if o == nil || isNil(o.Namespaced) {
		var ret bool
		return ret
	}
	return *o.Namespaced
}

// GetNamespacedOk returns a tuple with the Namespaced field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResource) GetNamespacedOk() (*bool, bool) {
	if o == nil || isNil(o.Namespaced) {
    return nil, false
	}
	return o.Namespaced, true
}

// HasNamespaced returns a boolean if a field has been set.
func (o *ApiResource) HasNamespaced() bool {
	if o != nil && !isNil(o.Namespaced) {
		return true
	}

	return false
}

// SetNamespaced gets a reference to the given bool and assigns it to the Namespaced field.
func (o *ApiResource) SetNamespaced(v bool) {
	o.Namespaced = &v
}

// GetVerbs returns the Verbs field value if set, zero value otherwise.
func (o *ApiResource) GetVerbs() []string {
	if o == nil || isNil(o.Verbs) {
		var ret []string
		return ret
	}
	return o.Verbs
}

// GetVerbsOk returns a tuple with the Verbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResource) GetVerbsOk() ([]string, bool) {
	if o == nil || isNil(o.Verbs) {
    return nil, false
	}
	return o.Verbs, true
}

// HasVerbs returns a boolean if a field has been set.
func (o *ApiResource) HasVerbs() bool {
	if o != nil && !isNil(o.Verbs) {
		return true
	}

	return false
}

// SetVerbs gets a reference to the given []string and assigns it to the Verbs field.
func (o *ApiResource) SetVerbs(v []string) {
	o.Verbs = v
}

func (o ApiResource) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespaced) {
		toSerialize["namespaced"] = o.Namespaced
	}
	if !isNil(o.Verbs) {
		toSerialize["verbs"] = o.Verbs
	}
	return json.Marshal(toSerialize)
}

type NullableApiResource struct {
	value *ApiResource
	isSet bool
}

func (v NullableApiResource) Get() *ApiResource {
	return v.value
}

func (v *NullableApiResource) Set(val *ApiResource) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResource) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResource(val *ApiResource) *NullableApiResource {
	return &NullableApiResource{value: val, isSet: true}
}

func (v NullableApiResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


