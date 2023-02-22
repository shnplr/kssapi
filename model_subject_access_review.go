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

// checks if the SubjectAccessReview type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubjectAccessReview{}

// SubjectAccessReview struct for SubjectAccessReview
type SubjectAccessReview struct {
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Verb *string `json:"verb,omitempty"`
	Resource *string `json:"resource,omitempty"`
	ResourceName *string `json:"resourceName,omitempty"`
	User *string `json:"user,omitempty"`
}

// NewSubjectAccessReview instantiates a new SubjectAccessReview object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubjectAccessReview() *SubjectAccessReview {
	this := SubjectAccessReview{}
	return &this
}

// NewSubjectAccessReviewWithDefaults instantiates a new SubjectAccessReview object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubjectAccessReviewWithDefaults() *SubjectAccessReview {
	this := SubjectAccessReview{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *SubjectAccessReview) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SubjectAccessReview) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *SubjectAccessReview) SetNamespace(v string) {
	o.Namespace = &v
}

// GetVerb returns the Verb field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetVerb() string {
	if o == nil || IsNil(o.Verb) {
		var ret string
		return ret
	}
	return *o.Verb
}

// GetVerbOk returns a tuple with the Verb field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetVerbOk() (*string, bool) {
	if o == nil || IsNil(o.Verb) {
		return nil, false
	}
	return o.Verb, true
}

// HasVerb returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasVerb() bool {
	if o != nil && !IsNil(o.Verb) {
		return true
	}

	return false
}

// SetVerb gets a reference to the given string and assigns it to the Verb field.
func (o *SubjectAccessReview) SetVerb(v string) {
	o.Verb = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *SubjectAccessReview) SetResource(v string) {
	o.Resource = &v
}

// GetResourceName returns the ResourceName field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetResourceName() string {
	if o == nil || IsNil(o.ResourceName) {
		var ret string
		return ret
	}
	return *o.ResourceName
}

// GetResourceNameOk returns a tuple with the ResourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetResourceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ResourceName) {
		return nil, false
	}
	return o.ResourceName, true
}

// HasResourceName returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasResourceName() bool {
	if o != nil && !IsNil(o.ResourceName) {
		return true
	}

	return false
}

// SetResourceName gets a reference to the given string and assigns it to the ResourceName field.
func (o *SubjectAccessReview) SetResourceName(v string) {
	o.ResourceName = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *SubjectAccessReview) GetUser() string {
	if o == nil || IsNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubjectAccessReview) GetUserOk() (*string, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *SubjectAccessReview) HasUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *SubjectAccessReview) SetUser(v string) {
	o.User = &v
}

func (o SubjectAccessReview) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubjectAccessReview) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !IsNil(o.Verb) {
		toSerialize["verb"] = o.Verb
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	if !IsNil(o.ResourceName) {
		toSerialize["resourceName"] = o.ResourceName
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	return toSerialize, nil
}

type NullableSubjectAccessReview struct {
	value *SubjectAccessReview
	isSet bool
}

func (v NullableSubjectAccessReview) Get() *SubjectAccessReview {
	return v.value
}

func (v *NullableSubjectAccessReview) Set(val *SubjectAccessReview) {
	v.value = val
	v.isSet = true
}

func (v NullableSubjectAccessReview) IsSet() bool {
	return v.isSet
}

func (v *NullableSubjectAccessReview) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubjectAccessReview(val *SubjectAccessReview) *NullableSubjectAccessReview {
	return &NullableSubjectAccessReview{value: val, isSet: true}
}

func (v NullableSubjectAccessReview) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubjectAccessReview) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


