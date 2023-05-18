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

// checks if the PolicyRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PolicyRule{}

// PolicyRule struct for PolicyRule
type PolicyRule struct {
	Verbs []string `json:"verbs,omitempty"`
	Resources []string `json:"resources,omitempty"`
	ResourceNames []string `json:"resourceNames,omitempty"`
}

// NewPolicyRule instantiates a new PolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPolicyRule() *PolicyRule {
	this := PolicyRule{}
	return &this
}

// NewPolicyRuleWithDefaults instantiates a new PolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPolicyRuleWithDefaults() *PolicyRule {
	this := PolicyRule{}
	return &this
}

// GetVerbs returns the Verbs field value if set, zero value otherwise.
func (o *PolicyRule) GetVerbs() []string {
	if o == nil || IsNil(o.Verbs) {
		var ret []string
		return ret
	}
	return o.Verbs
}

// GetVerbsOk returns a tuple with the Verbs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetVerbsOk() ([]string, bool) {
	if o == nil || IsNil(o.Verbs) {
		return nil, false
	}
	return o.Verbs, true
}

// HasVerbs returns a boolean if a field has been set.
func (o *PolicyRule) HasVerbs() bool {
	if o != nil && !IsNil(o.Verbs) {
		return true
	}

	return false
}

// SetVerbs gets a reference to the given []string and assigns it to the Verbs field.
func (o *PolicyRule) SetVerbs(v []string) {
	o.Verbs = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *PolicyRule) GetResources() []string {
	if o == nil || IsNil(o.Resources) {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetResourcesOk() ([]string, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *PolicyRule) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *PolicyRule) SetResources(v []string) {
	o.Resources = v
}

// GetResourceNames returns the ResourceNames field value if set, zero value otherwise.
func (o *PolicyRule) GetResourceNames() []string {
	if o == nil || IsNil(o.ResourceNames) {
		var ret []string
		return ret
	}
	return o.ResourceNames
}

// GetResourceNamesOk returns a tuple with the ResourceNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PolicyRule) GetResourceNamesOk() ([]string, bool) {
	if o == nil || IsNil(o.ResourceNames) {
		return nil, false
	}
	return o.ResourceNames, true
}

// HasResourceNames returns a boolean if a field has been set.
func (o *PolicyRule) HasResourceNames() bool {
	if o != nil && !IsNil(o.ResourceNames) {
		return true
	}

	return false
}

// SetResourceNames gets a reference to the given []string and assigns it to the ResourceNames field.
func (o *PolicyRule) SetResourceNames(v []string) {
	o.ResourceNames = v
}

func (o PolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PolicyRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Verbs) {
		toSerialize["verbs"] = o.Verbs
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.ResourceNames) {
		toSerialize["resourceNames"] = o.ResourceNames
	}
	return toSerialize, nil
}

type NullablePolicyRule struct {
	value *PolicyRule
	isSet bool
}

func (v NullablePolicyRule) Get() *PolicyRule {
	return v.value
}

func (v *NullablePolicyRule) Set(val *PolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePolicyRule(val *PolicyRule) *NullablePolicyRule {
	return &NullablePolicyRule{value: val, isSet: true}
}

func (v NullablePolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

