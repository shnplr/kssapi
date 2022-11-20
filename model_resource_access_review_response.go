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

// ResourceAccessReviewResponse struct for ResourceAccessReviewResponse
type ResourceAccessReviewResponse struct {
	Namespace *string `json:"namespace,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewResourceAccessReviewResponse instantiates a new ResourceAccessReviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceAccessReviewResponse() *ResourceAccessReviewResponse {
	this := ResourceAccessReviewResponse{}
	return &this
}

// NewResourceAccessReviewResponseWithDefaults instantiates a new ResourceAccessReviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceAccessReviewResponseWithDefaults() *ResourceAccessReviewResponse {
	this := ResourceAccessReviewResponse{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ResourceAccessReviewResponse) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReviewResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
    return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ResourceAccessReviewResponse) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *ResourceAccessReviewResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *ResourceAccessReviewResponse) GetAllowed() bool {
	if o == nil || isNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReviewResponse) GetAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.Allowed) {
    return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *ResourceAccessReviewResponse) HasAllowed() bool {
	if o != nil && !isNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *ResourceAccessReviewResponse) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o ResourceAccessReviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !isNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return json.Marshal(toSerialize)
}

type NullableResourceAccessReviewResponse struct {
	value *ResourceAccessReviewResponse
	isSet bool
}

func (v NullableResourceAccessReviewResponse) Get() *ResourceAccessReviewResponse {
	return v.value
}

func (v *NullableResourceAccessReviewResponse) Set(val *ResourceAccessReviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceAccessReviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceAccessReviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceAccessReviewResponse(val *ResourceAccessReviewResponse) *NullableResourceAccessReviewResponse {
	return &NullableResourceAccessReviewResponse{value: val, isSet: true}
}

func (v NullableResourceAccessReviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceAccessReviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


