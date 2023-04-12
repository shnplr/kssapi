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

// checks if the GenericListApiResource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GenericListApiResource{}

// GenericListApiResource struct for GenericListApiResource
type GenericListApiResource struct {
	Kind *string `json:"kind,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	Items []ApiResource `json:"items,omitempty"`
}

// NewGenericListApiResource instantiates a new GenericListApiResource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGenericListApiResource() *GenericListApiResource {
	this := GenericListApiResource{}
	return &this
}

// NewGenericListApiResourceWithDefaults instantiates a new GenericListApiResource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGenericListApiResourceWithDefaults() *GenericListApiResource {
	this := GenericListApiResource{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *GenericListApiResource) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListApiResource) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *GenericListApiResource) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *GenericListApiResource) SetKind(v string) {
	o.Kind = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *GenericListApiResource) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListApiResource) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *GenericListApiResource) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *GenericListApiResource) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *GenericListApiResource) GetItems() []ApiResource {
	if o == nil || IsNil(o.Items) {
		var ret []ApiResource
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GenericListApiResource) GetItemsOk() ([]ApiResource, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *GenericListApiResource) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []ApiResource and assigns it to the Items field.
func (o *GenericListApiResource) SetItems(v []ApiResource) {
	o.Items = v
}

func (o GenericListApiResource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GenericListApiResource) ToMap() (map[string]interface{}, error) {
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

type NullableGenericListApiResource struct {
	value *GenericListApiResource
	isSet bool
}

func (v NullableGenericListApiResource) Get() *GenericListApiResource {
	return v.value
}

func (v *NullableGenericListApiResource) Set(val *GenericListApiResource) {
	v.value = val
	v.isSet = true
}

func (v NullableGenericListApiResource) IsSet() bool {
	return v.isSet
}

func (v *NullableGenericListApiResource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGenericListApiResource(val *GenericListApiResource) *NullableGenericListApiResource {
	return &NullableGenericListApiResource{value: val, isSet: true}
}

func (v NullableGenericListApiResource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGenericListApiResource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


