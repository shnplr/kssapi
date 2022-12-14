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
	Kind *string `json:"kind,omitempty"`
	Users []string `json:"users,omitempty"`
	Groups []string `json:"groups,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Name *string `json:"name,omitempty"`
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

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *ResourceAccessReviewResponse) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReviewResponse) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *ResourceAccessReviewResponse) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *ResourceAccessReviewResponse) SetKind(v string) {
	o.Kind = &v
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *ResourceAccessReviewResponse) GetUsers() []string {
	if o == nil || isNil(o.Users) {
		var ret []string
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReviewResponse) GetUsersOk() ([]string, bool) {
	if o == nil || isNil(o.Users) {
    return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *ResourceAccessReviewResponse) HasUsers() bool {
	if o != nil && !isNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []string and assigns it to the Users field.
func (o *ResourceAccessReviewResponse) SetUsers(v []string) {
	o.Users = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *ResourceAccessReviewResponse) GetGroups() []string {
	if o == nil || isNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReviewResponse) GetGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.Groups) {
    return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *ResourceAccessReviewResponse) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *ResourceAccessReviewResponse) SetGroups(v []string) {
	o.Groups = v
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

// GetName returns the Name field value if set, zero value otherwise.
func (o *ResourceAccessReviewResponse) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceAccessReviewResponse) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
    return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ResourceAccessReviewResponse) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ResourceAccessReviewResponse) SetName(v string) {
	o.Name = &v
}

func (o ResourceAccessReviewResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	if !isNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
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


