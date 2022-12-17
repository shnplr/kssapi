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

// ResourceAccessReview struct for ResourceAccessReview
type ResourceAccessReview struct {
	Verb *string `json:"verb,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
}

// NewResourceAccessReview instantiates a new ResourceAccessReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAccessReview() *ResourceAccessReview {
	this := ResourceAccessReview{}
	return &this
}

// NewResourceAccessReviewWithDefaults instantiates a new ResourceAccessReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAccessReviewWithDefaults() *ResourceAccessReview {
	this := ResourceAccessReview{}
	return &this
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *ResourceAccessReview) GetVerb() string {
	if o == nil || isNil(o.Verb) {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReview) GetVerbOk() (*string, bool) {
	if o == nil || isNil(o.Verb) {
    return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *ResourceAccessReview) HasVerb() bool {
	if o != nil && !isNil(o.Verb) {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *ResourceAccessReview) SetVerb(v string) {
	o.Verb = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *ResourceAccessReview) GetResource() string {
	if o == nil || isNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReview) GetResourceOk() (*string, bool) {
	if o == nil || isNil(o.Resource) {
    return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *ResourceAccessReview) HasResource() bool {
	if o != nil && !isNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *ResourceAccessReview) SetResource(v string) {
	o.Resource = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *ResourceAccessReview) GetResourceName() string {
	if o == nil || isNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReview) GetResourceNameOk() (*string, bool) {
	if o == nil || isNil(o.ResourceName) {
    return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *ResourceAccessReview) HasResourceName() bool {
	if o != nil && !isNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *ResourceAccessReview) SetResourceName(v string) {
	o.ResourceName = &v
}

func (o ResourceAccessReview) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Verb) {
		toSerialize["verb"] = o.Verb
	}
	if !isNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !isNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	return json.Marshal(toSerialize)
}

type NullableResourceAccessReview struct {
	value *ResourceAccessReview
	isSet bool
}

func (v NullableResourceAccessReview) Get() *ResourceAccessReview {
	return v.value
}

func (v *NullableResourceAccessReview) Set(val *ResourceAccessReview) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAccessReview) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAccessReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAccessReview(val *ResourceAccessReview) *NullableResourceAccessReview {
	return &NullableResourceAccessReview{value: val, isSet: true}
}

func (v NullableResourceAccessReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAccessReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


