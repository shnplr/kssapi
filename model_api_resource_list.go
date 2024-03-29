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

// checks if the ApiResourceList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiResourceList{}

// ApiResourceList struct for ApiResourceList
type ApiResourceList struct {
	Kind *string `json:"kind,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []ApiResource `json:"items,omitempty"`
}

// NewApiResourceList instantiates a new ApiResourceList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiResourceList() *ApiResourceList {
	this := ApiResourceList{}
	return &this
}

// NewApiResourceListWithDefaults instantiates a new ApiResourceList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiResourceListWithDefaults() *ApiResourceList {
	this := ApiResourceList{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ApiResourceList) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResourceList) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ApiResourceList) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ApiResourceList) SetKind(v string) {
	o.Kind = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *ApiResourceList) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResourceList) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *ApiResourceList) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *ApiResourceList) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ApiResourceList) GetItems() []ApiResource {
	if o == nil || IsNil(o.Items) {
		var ret []ApiResource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiResourceList) GetItemsOk() ([]ApiResource, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ApiResourceList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ApiResource and assigns it to the Items field.
func (o *ApiResourceList) SetItems(v []ApiResource) {
	o.Items = v
}

func (o ApiResourceList) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiResourceList) ToMap() (map[string]interface{}, error) {
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

type NullableApiResourceList struct {
	value *ApiResourceList
	isSet bool
}

func (v NullableApiResourceList) Get() *ApiResourceList {
	return v.value
}

func (v *NullableApiResourceList) Set(val *ApiResourceList) {
	v.value = val
	v.isSet = true
}

func (v NullableApiResourceList) IsSet() bool {
	return v.isSet
}

func (v *NullableApiResourceList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiResourceList(val *ApiResourceList) *NullableApiResourceList {
	return &NullableApiResourceList{value: val, isSet: true}
}

func (v NullableApiResourceList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiResourceList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


