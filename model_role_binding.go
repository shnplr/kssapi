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

// checks if the RoleBinding type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoleBinding{}

// RoleBinding struct for RoleBinding
type RoleBinding struct {
	Kind *string `json:"kind,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	RoleRef RoleRef `json:"roleRef"`
	Subjects []Subject `json:"subjects,omitempty"`
}

// NewRoleBinding instantiates a new RoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoleBinding(name string, namespace string, roleRef RoleRef) *RoleBinding {
	this := RoleBinding{}
	this.Name = name
	this.Namespace = namespace
	this.RoleRef = roleRef
	return &this
}

// NewRoleBindingWithDefaults instantiates a new RoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoleBindingWithDefaults() *RoleBinding {
	this := RoleBinding{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RoleBinding) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RoleBinding) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RoleBinding) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value
func (o *RoleBinding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *RoleBinding) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *RoleBinding) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *RoleBinding) SetNamespace(v string) {
	o.Namespace = v
}

// GetRoleRef returns the RoleRef field value
func (o *RoleBinding) GetRoleRef() RoleRef {
	if o == nil {
		var ret RoleRef
		return ret
	}

	return o.RoleRef
}

// GetRoleRefOk returns a tuple with the RoleRef field value
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetRoleRefOk() (*RoleRef, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RoleRef, true
}

// SetRoleRef sets field value
func (o *RoleBinding) SetRoleRef(v RoleRef) {
	o.RoleRef = v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *RoleBinding) GetSubjects() []Subject {
	if o == nil || IsNil(o.Subjects) {
		var ret []Subject
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoleBinding) GetSubjectsOk() ([]Subject, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *RoleBinding) HasSubjects() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []Subject and assigns it to the Subjects field.
func (o *RoleBinding) SetSubjects(v []Subject) {
	o.Subjects = v
}

func (o RoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoleBinding) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	toSerialize["roleRef"] = o.RoleRef
	if !IsNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return toSerialize, nil
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


