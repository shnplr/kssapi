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

// SubjectAccessReviewResponse struct for SubjectAccessReviewResponse
type SubjectAccessReviewResponse struct {
	Namespace *string `json:"namespace,omitempty"`
	Allowed *bool `json:"allowed,omitempty"`
}

// NewSubjectAccessReviewResponse instantiates a new SubjectAccessReviewResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectAccessReviewResponse() *SubjectAccessReviewResponse {
	this := SubjectAccessReviewResponse{}
	return &this
}

// NewSubjectAccessReviewResponseWithDefaults instantiates a new SubjectAccessReviewResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectAccessReviewResponseWithDefaults() *SubjectAccessReviewResponse {
	this := SubjectAccessReviewResponse{}
	return &this
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SubjectAccessReviewResponse) GetNamespace() string {
	if o == nil || isNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReviewResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || isNil(o.Namespace) {
    return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SubjectAccessReviewResponse) HasNamespace() bool {
	if o != nil && !isNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SubjectAccessReviewResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetAllowed returns the Allowed field value if set, zero value otherwise.
func (o *SubjectAccessReviewResponse) GetAllowed() bool {
	if o == nil || isNil(o.Allowed) {
		var ret bool
		return ret
	}
	return *o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReviewResponse) GetAllowedOk() (*bool, bool) {
	if o == nil || isNil(o.Allowed) {
    return nil, false
	}
	return o.Allowed, true
}

// HasAllowed returns a boolean if a field has been set.
func (o *SubjectAccessReviewResponse) HasAllowed() bool {
	if o != nil && !isNil(o.Allowed) {
		return true
	}

	return false
}

// SetAllowed gets a reference to the given bool and assigns it to the Allowed field.
func (o *SubjectAccessReviewResponse) SetAllowed(v bool) {
	o.Allowed = &v
}

func (o SubjectAccessReviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !isNil(o.Allowed) {
		toSerialize["allowed"] = o.Allowed
	}
	return json.Marshal(toSerialize)
}

type NullableSubjectAccessReviewResponse struct {
	value *SubjectAccessReviewResponse
	isSet bool
}

func (v NullableSubjectAccessReviewResponse) Get() *SubjectAccessReviewResponse {
	return v.value
}

func (v *NullableSubjectAccessReviewResponse) Set(val *SubjectAccessReviewResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectAccessReviewResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectAccessReviewResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectAccessReviewResponse(val *SubjectAccessReviewResponse) *NullableSubjectAccessReviewResponse {
	return &NullableSubjectAccessReviewResponse{value: val, isSet: true}
}

func (v NullableSubjectAccessReviewResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectAccessReviewResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

