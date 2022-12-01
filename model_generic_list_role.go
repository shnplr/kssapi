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

// GenericListRole struct for GenericListRole
type GenericListRole struct {
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	Items []Role `json:"items,omitempty"`
	Kind *string `json:"kind,omitempty"`
}

// NewGenericListRole instantiates a new GenericListRole object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericListRole() *GenericListRole {
	this := GenericListRole{}
	return &this
}

// NewGenericListRoleWithDefaults instantiates a new GenericListRole object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericListRoleWithDefaults() *GenericListRole {
	this := GenericListRole{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *GenericListRole) GetMetadata() ObjectMeta {
	if o == nil || isNil(o.Metadata) {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRole) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || isNil(o.Metadata) {
    return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *GenericListRole) HasMetadata() bool {
	if o != nil && !isNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *GenericListRole) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GenericListRole) GetItems() []Role {
	if o == nil || isNil(o.Items) {
		var ret []Role
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRole) GetItemsOk() ([]Role, bool) {
	if o == nil || isNil(o.Items) {
    return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GenericListRole) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Role and assigns it to the Items field.
func (o *GenericListRole) SetItems(v []Role) {
	o.Items = v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GenericListRole) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListRole) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GenericListRole) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GenericListRole) SetKind(v string) {
	o.Kind = &v
}

func (o GenericListRole) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	return json.Marshal(toSerialize)
}

type NullableGenericListRole struct {
	value *GenericListRole
	isSet bool
}

func (v NullableGenericListRole) Get() *GenericListRole {
	return v.value
}

func (v *NullableGenericListRole) Set(val *GenericListRole) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericListRole) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericListRole) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericListRole(val *GenericListRole) *NullableGenericListRole {
	return &NullableGenericListRole{value: val, isSet: true}
}

func (v NullableGenericListRole) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericListRole) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


