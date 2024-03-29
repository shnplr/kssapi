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

// checks if the ResourcePattern type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourcePattern{}

// ResourcePattern struct for ResourcePattern
type ResourcePattern struct {
	ResourceType *string `json:"resourceType,omitempty"`
	Name *string `json:"name,omitempty"`
	PatternType *string `json:"patternType,omitempty"`
}

// NewResourcePattern instantiates a new ResourcePattern object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourcePattern() *ResourcePattern {
	this := ResourcePattern{}
	return &this
}

// NewResourcePatternWithDefaults instantiates a new ResourcePattern object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcePatternWithDefaults() *ResourcePattern {
	this := ResourcePattern{}
	return &this
}

// GetResourceType returns the ResourceType field value if set, zero value otherwise.
func (o *ResourcePattern) GetResourceType() string {
	if o == nil || IsNil(o.ResourceType) {
		var ret string
		return ret
	}
	return *o.ResourceType
}

// GetResourceTypeOk returns a tuple with the ResourceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePattern) GetResourceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceType) {
		return nil, false
	}
	return o.ResourceType, true
}

// HasResourceType returns a boolean if a field has been set.
func (o *ResourcePattern) HasResourceType() bool {
	if o != nil && !IsNil(o.ResourceType) {
		return true
	}

	return false
}

// SetResourceType gets a reference to the given string and assigns it to the ResourceType field.
func (o *ResourcePattern) SetResourceType(v string) {
	o.ResourceType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourcePattern) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePattern) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourcePattern) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourcePattern) SetName(v string) {
	o.Name = &v
}

// GetPatternType returns the PatternType field value if set, zero value otherwise.
func (o *ResourcePattern) GetPatternType() string {
	if o == nil || IsNil(o.PatternType) {
		var ret string
		return ret
	}
	return *o.PatternType
}

// GetPatternTypeOk returns a tuple with the PatternType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourcePattern) GetPatternTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PatternType) {
		return nil, false
	}
	return o.PatternType, true
}

// HasPatternType returns a boolean if a field has been set.
func (o *ResourcePattern) HasPatternType() bool {
	if o != nil && !IsNil(o.PatternType) {
		return true
	}

	return false
}

// SetPatternType gets a reference to the given string and assigns it to the PatternType field.
func (o *ResourcePattern) SetPatternType(v string) {
	o.PatternType = &v
}

func (o ResourcePattern) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourcePattern) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourceType) {
		toSerialize["resourceType"] = o.ResourceType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.PatternType) {
		toSerialize["patternType"] = o.PatternType
	}
	return toSerialize, nil
}

type NullableResourcePattern struct {
	value *ResourcePattern
	isSet bool
}

func (v NullableResourcePattern) Get() *ResourcePattern {
	return v.value
}

func (v *NullableResourcePattern) Set(val *ResourcePattern) {
	v.value = val
	v.isSet = true
}

func (v NullableResourcePattern) IsSet() bool {
	return v.isSet
}

func (v *NullableResourcePattern) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourcePattern(val *ResourcePattern) *NullableResourcePattern {
	return &NullableResourcePattern{value: val, isSet: true}
}

func (v NullableResourcePattern) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourcePattern) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


