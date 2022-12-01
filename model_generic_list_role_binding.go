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

// GenericListRoleBinding struct for GenericListRoleBinding
type GenericListRoleBinding struct {
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	Name *string `json:"name,omitempty"`
	Items []RoleBinding `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// NewGenericListRoleBinding instantiates a new GenericListRoleBinding object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericListRoleBinding() *GenericListRoleBinding {
	this := GenericListRoleBinding{}
	return &this
}

// NewGenericListRoleBindingWithDefaults instantiates a new GenericListRoleBinding object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericListRoleBindingWithDefaults() *GenericListRoleBinding {
	this := GenericListRoleBinding{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GenericListRoleBinding) GetMetadata() ObjectMeta {
	if o == nil || isNil(o.Metadata) {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRoleBinding) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GenericListRoleBinding) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *GenericListRoleBinding) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GenericListRoleBinding) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRoleBinding) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *GenericListRoleBinding) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GenericListRoleBinding) SetName(v string) {
	o.Name = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GenericListRoleBinding) GetItems() []RoleBinding {
	if o == nil || isNil(o.Items) {
		var ret []RoleBinding
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRoleBinding) GetItemsOk() ([]RoleBinding, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GenericListRoleBinding) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []RoleBinding and assigns it to the Items field.
func (o *GenericListRoleBinding) SetItems(v []RoleBinding) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GenericListRoleBinding) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRoleBinding) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GenericListRoleBinding) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GenericListRoleBinding) SetKind(v string) {
	o.Kind = &v
}

func (o GenericListRoleBinding) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableGenericListRoleBinding struct {
	value *GenericListRoleBinding
	isSet bool
}

func (v NullableGenericListRoleBinding) Get() *GenericListRoleBinding {
	return v.value
}

func (v *NullableGenericListRoleBinding) Set(val *GenericListRoleBinding) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericListRoleBinding) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericListRoleBinding) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericListRoleBinding(val *GenericListRoleBinding) *NullableGenericListRoleBinding {
	return &NullableGenericListRoleBinding{value: val, isSet: true}
}

func (v NullableGenericListRoleBinding) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericListRoleBinding) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


