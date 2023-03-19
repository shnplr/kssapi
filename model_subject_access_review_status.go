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

// checks if the SubjectAccessReviewStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectAccessReviewStatus{}

// SubjectAccessReviewStatus struct for SubjectAccessReviewStatus
type SubjectAccessReviewStatus struct {
	Allowed bool `json:"allowed"`
	Reason *string `json:"reason,omitempty"`
}

// NewSubjectAccessReviewStatus instantiates a new SubjectAccessReviewStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectAccessReviewStatus(allowed bool) *SubjectAccessReviewStatus {
	this := SubjectAccessReviewStatus{}
	this.Allowed = allowed
	return &this
}

// NewSubjectAccessReviewStatusWithDefaults instantiates a new SubjectAccessReviewStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectAccessReviewStatusWithDefaults() *SubjectAccessReviewStatus {
	this := SubjectAccessReviewStatus{}
	return &this
}

// GetAllowed returns the Allowed field value
func (o *SubjectAccessReviewStatus) GetAllowed() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Allowed
}

// GetAllowedOk returns a tuple with the Allowed field value
// and a boolean to check if the value has been set.
func (o *SubjectAccessReviewStatus) GetAllowedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Allowed, true
}

// SetAllowed sets field value
func (o *SubjectAccessReviewStatus) SetAllowed(v bool) {
	o.Allowed = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *SubjectAccessReviewStatus) GetReason() string {
	if o == nil || IsNil(o.Reason) {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReviewStatus) GetReasonOk() (*string, bool) {
	if o == nil || IsNil(o.Reason) {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *SubjectAccessReviewStatus) HasReason() bool {
	if o != nil && !IsNil(o.Reason) {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *SubjectAccessReviewStatus) SetReason(v string) {
	o.Reason = &v
}

func (o SubjectAccessReviewStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectAccessReviewStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["allowed"] = o.Allowed
	if !IsNil(o.Reason) {
		toSerialize["reason"] = o.Reason
	}
	return toSerialize, nil
}

type NullableSubjectAccessReviewStatus struct {
	value *SubjectAccessReviewStatus
	isSet bool
}

func (v NullableSubjectAccessReviewStatus) Get() *SubjectAccessReviewStatus {
	return v.value
}

func (v *NullableSubjectAccessReviewStatus) Set(val *SubjectAccessReviewStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectAccessReviewStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectAccessReviewStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectAccessReviewStatus(val *SubjectAccessReviewStatus) *NullableSubjectAccessReviewStatus {
	return &NullableSubjectAccessReviewStatus{value: val, isSet: true}
}

func (v NullableSubjectAccessReviewStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectAccessReviewStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


