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

// checks if the RbacRoleBindingResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacRoleBindingResponse{}

// RbacRoleBindingResponse struct for RbacRoleBindingResponse
type RbacRoleBindingResponse struct {
	Kind *string `json:"kind,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Principal *string `json:"principal,omitempty"`
	Resources []ResourcePattern `json:"resources,omitempty"`
}

// NewRbacRoleBindingResponse instantiates a new RbacRoleBindingResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacRoleBindingResponse() *RbacRoleBindingResponse {
	this := RbacRoleBindingResponse{}
	return &this
}

// NewRbacRoleBindingResponseWithDefaults instantiates a new RbacRoleBindingResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacRoleBindingResponseWithDefaults() *RbacRoleBindingResponse {
	this := RbacRoleBindingResponse{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RbacRoleBindingResponse) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingResponse) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RbacRoleBindingResponse) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RbacRoleBindingResponse) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RbacRoleBindingResponse) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingResponse) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RbacRoleBindingResponse) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RbacRoleBindingResponse) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *RbacRoleBindingResponse) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingResponse) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *RbacRoleBindingResponse) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *RbacRoleBindingResponse) SetNamespace(v string) {
	o.Namespace = &v
}

// GetPrincipal returns the Principal field value if set, zero value otherwise.
func (o *RbacRoleBindingResponse) GetPrincipal() string {
	if o == nil || IsNil(o.Principal) {
		var ret string
		return ret
	}
	return *o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingResponse) GetPrincipalOk() (*string, bool) {
	if o == nil || IsNil(o.Principal) {
		return nil, false
	}
	return o.Principal, true
}

// HasPrincipal returns a boolean if a field has been set.
func (o *RbacRoleBindingResponse) HasPrincipal() bool {
	if o != nil && !IsNil(o.Principal) {
		return true
	}

	return false
}

// SetPrincipal gets a reference to the given string and assigns it to the Principal field.
func (o *RbacRoleBindingResponse) SetPrincipal(v string) {
	o.Principal = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *RbacRoleBindingResponse) GetResources() []ResourcePattern {
	if o == nil || IsNil(o.Resources) {
		var ret []ResourcePattern
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingResponse) GetResourcesOk() ([]ResourcePattern, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *RbacRoleBindingResponse) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []ResourcePattern and assigns it to the Resources field.
func (o *RbacRoleBindingResponse) SetResources(v []ResourcePattern) {
	o.Resources = v
}

func (o RbacRoleBindingResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacRoleBindingResponse) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.Principal) {
		toSerialize["principal"] = o.Principal
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	return toSerialize, nil
}

type NullableRbacRoleBindingResponse struct {
	value *RbacRoleBindingResponse
	isSet bool
}

func (v NullableRbacRoleBindingResponse) Get() *RbacRoleBindingResponse {
	return v.value
}

func (v *NullableRbacRoleBindingResponse) Set(val *RbacRoleBindingResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacRoleBindingResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacRoleBindingResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacRoleBindingResponse(val *RbacRoleBindingResponse) *NullableRbacRoleBindingResponse {
	return &NullableRbacRoleBindingResponse{value: val, isSet: true}
}

func (v NullableRbacRoleBindingResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacRoleBindingResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


