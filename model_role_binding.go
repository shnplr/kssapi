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

// RoleBinding struct for RoleBinding
type RoleBinding struct {
	Metadata ObjectMeta `json:"metadata"`
	Kind *string `json:"kind,omitempty"`
	Role string `json:"role"`
	Subjects []Subject `json:"subjects,omitempty"`
}

// NewRoleBinding instantiates a new RoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBinding(metadata ObjectMeta, role string) *RoleBinding {
	this := RoleBinding{}
	this.Metadata = metadata
	this.Role = role
	return &this
}

// NewRoleBindingWithDefaults instantiates a new RoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingWithDefaults() *RoleBinding {
	this := RoleBinding{}
	return &this
}

// GetMetadata returns the Metadata field value
func (o *RoleBinding) GetMetadata() ObjectMeta {
	if o == nil {
		var ret ObjectMeta
		return ret
	}

	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Metadata, true
}

// SetMetadata sets field value
func (o *RoleBinding) SetMetadata(v ObjectMeta) {
	o.Metadata = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RoleBinding) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RoleBinding) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RoleBinding) SetKind(v string) {
	o.Kind = &v
}

// GetRole returns the Role field value
func (o *RoleBinding) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetRoleOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RoleBinding) SetRole(v string) {
	o.Role = v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *RoleBinding) GetSubjects() []Subject {
	if o == nil || isNil(o.Subjects) {
		var ret []Subject
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetSubjectsOk() ([]Subject, bool) {
	if o == nil || isNil(o.Subjects) {
    return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *RoleBinding) HasSubjects() bool {
	if o != nil && !isNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []Subject and assigns it to the Subjects field.
func (o *RoleBinding) SetSubjects(v []Subject) {
	o.Subjects = v
}

func (o RoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return json.Marshal(toSerialize)
}

type NullableRoleBinding struct {
	value *RoleBinding
	isSet bool
}

func (v NullableRoleBinding) Get() *RoleBinding {
	return v.value
}

func (v *NullableRoleBinding) Set(val *RoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleBinding(val *RoleBinding) *NullableRoleBinding {
	return &NullableRoleBinding{value: val, isSet: true}
}

func (v NullableRoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


