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

// ClusterRoleBinding struct for ClusterRoleBinding
type ClusterRoleBinding struct {
	Kind *string `json:"kind,omitempty"`
	Name string `json:"name"`
	Namespace *string `json:"namespace,omitempty"`
	RoleRef RoleRef `json:"roleRef"`
	Subjects []Subject `json:"subjects,omitempty"`
}

// NewClusterRoleBinding instantiates a new ClusterRoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterRoleBinding(name string, roleRef RoleRef) *ClusterRoleBinding {
	this := ClusterRoleBinding{}
	this.Name = name
	this.RoleRef = roleRef
	return &this
}

// NewClusterRoleBindingWithDefaults instantiates a new ClusterRoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterRoleBindingWithDefaults() *ClusterRoleBinding {
	this := ClusterRoleBinding{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ClusterRoleBinding) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRoleBinding) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ClusterRoleBinding) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ClusterRoleBinding) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value
func (o *ClusterRoleBinding) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ClusterRoleBinding) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ClusterRoleBinding) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ClusterRoleBinding) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRoleBinding) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
    return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ClusterRoleBinding) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ClusterRoleBinding) SetNamespace(v string) {
	o.Namespace = &v
}

// GetRoleRef returns the RoleRef field value
func (o *ClusterRoleBinding) GetRoleRef() RoleRef {
	if o == nil {
		var ret RoleRef
		return ret
	}

	return o.RoleRef
}

// GetRoleRefOk returns a tuple with the RoleRef field value
// and a boolean to check if the value has been set.
func (o *ClusterRoleBinding) GetRoleRefOk() (*RoleRef, bool) {
	if o == nil {
    return nil, false
	}
	return &o.RoleRef, true
}

// SetRoleRef sets field value
func (o *ClusterRoleBinding) SetRoleRef(v RoleRef) {
	o.RoleRef = v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *ClusterRoleBinding) GetSubjects() []Subject {
	if o == nil || isNil(o.Subjects) {
		var ret []Subject
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClusterRoleBinding) GetSubjectsOk() ([]Subject, bool) {
	if o == nil || isNil(o.Subjects) {
    return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *ClusterRoleBinding) HasSubjects() bool {
	if o != nil && !isNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []Subject and assigns it to the Subjects field.
func (o *ClusterRoleBinding) SetSubjects(v []Subject) {
	o.Subjects = v
}

func (o ClusterRoleBinding) MarshalJSON() ([]byte, error) {
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
	if true {
		toSerialize["roleRef"] = o.RoleRef
	}
	if !isNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return json.Marshal(toSerialize)
}

type NullableClusterRoleBinding struct {
	value *ClusterRoleBinding
	isSet bool
}

func (v NullableClusterRoleBinding) Get() *ClusterRoleBinding {
	return v.value
}

func (v *NullableClusterRoleBinding) Set(val *ClusterRoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterRoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterRoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterRoleBinding(val *ClusterRoleBinding) *NullableClusterRoleBinding {
	return &NullableClusterRoleBinding{value: val, isSet: true}
}

func (v NullableClusterRoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterRoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


