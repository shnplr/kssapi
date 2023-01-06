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

// RoleRef struct for RoleRef
type RoleRef struct {
	Kind *string `json:"kind,omitempty"`
	Name string `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewRoleRef instantiates a new RoleRef object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleRef(name string) *RoleRef {
	this := RoleRef{}
	this.Name = name
	return &this
}

// NewRoleRefWithDefaults instantiates a new RoleRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleRefWithDefaults() *RoleRef {
	this := RoleRef{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RoleRef) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRef) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RoleRef) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RoleRef) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value
func (o *RoleRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleRef) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleRef) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *RoleRef) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleRef) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
    return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *RoleRef) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *RoleRef) SetNamespace(v string) {
	o.Namespace = &v
}

func (o RoleRef) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableRoleRef struct {
	value *RoleRef
	isSet bool
}

func (v NullableRoleRef) Get() *RoleRef {
	return v.value
}

func (v *NullableRoleRef) Set(val *RoleRef) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleRef) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleRef(val *RoleRef) *NullableRoleRef {
	return &NullableRoleRef{value: val, isSet: true}
}

func (v NullableRoleRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


