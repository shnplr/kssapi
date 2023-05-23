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

// checks if the ProjectList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProjectList{}

// ProjectList struct for ProjectList
type ProjectList struct {
	Kind *string `json:"kind,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []Project `json:"items,omitempty"`
}

// NewProjectList instantiates a new ProjectList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectList() *ProjectList {
	this := ProjectList{}
	return &this
}

// NewProjectListWithDefaults instantiates a new ProjectList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectListWithDefaults() *ProjectList {
	this := ProjectList{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ProjectList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ProjectList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ProjectList) SetKind(v string) {
	o.Kind = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ProjectList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ProjectList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ProjectList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ProjectList) GetItems() []Project {
	if o == nil || IsNil(o.Items) {
		var ret []Project
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectList) GetItemsOk() ([]Project, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ProjectList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Project and assigns it to the Items field.
func (o *ProjectList) SetItems(v []Project) {
	o.Items = v
}

func (o ProjectList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProjectList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	return toSerialize, nil
}

type NullableProjectList struct {
	value *ProjectList
	isSet bool
}

func (v NullableProjectList) Get() *ProjectList {
	return v.value
}

func (v *NullableProjectList) Set(val *ProjectList) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectList) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectList(val *ProjectList) *NullableProjectList {
	return &NullableProjectList{value: val, isSet: true}
}

func (v NullableProjectList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

