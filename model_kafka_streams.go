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

// KafkaStreams struct for KafkaStreams
type KafkaStreams struct {
	Kind *string `json:"kind,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Principal string `json:"principal"`
	Transactional *bool `json:"transactional,omitempty"`
}

// NewKafkaStreams instantiates a new KafkaStreams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaStreams(name string, namespace string, principal string) *KafkaStreams {
	this := KafkaStreams{}
	this.Name = name
	this.Namespace = namespace
	this.Principal = principal
	return &this
}

// NewKafkaStreamsWithDefaults instantiates a new KafkaStreams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaStreamsWithDefaults() *KafkaStreams {
	this := KafkaStreams{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KafkaStreams) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaStreams) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KafkaStreams) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KafkaStreams) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value
func (o *KafkaStreams) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KafkaStreams) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KafkaStreams) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *KafkaStreams) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *KafkaStreams) GetNamespaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *KafkaStreams) SetNamespace(v string) {
	o.Namespace = v
}

// GetPrincipal returns the Principal field value
func (o *KafkaStreams) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *KafkaStreams) GetPrincipalOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *KafkaStreams) SetPrincipal(v string) {
	o.Principal = v
}

// GetTransactional returns the Transactional field value if set, zero value otherwise.
func (o *KafkaStreams) GetTransactional() bool {
	if o == nil || isNil(o.Transactional) {
		var ret bool
		return ret
	}
	return *o.Transactional
}

// GetTransactionalOk returns a tuple with the Transactional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaStreams) GetTransactionalOk() (*bool, bool) {
	if o == nil || isNil(o.Transactional) {
    return nil, false
	}
	return o.Transactional, true
}

// HasTransactional returns a boolean if a field has been set.
func (o *KafkaStreams) HasTransactional() bool {
	if o != nil && !isNil(o.Transactional) {
		return true
	}

	return false
}

// SetTransactional gets a reference to the given bool and assigns it to the Transactional field.
func (o *KafkaStreams) SetTransactional(v bool) {
	o.Transactional = &v
}

func (o KafkaStreams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	if true {
		toSerialize["principal"] = o.Principal
	}
	if !isNil(o.Transactional) {
		toSerialize["transactional"] = o.Transactional
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaStreams struct {
	value *KafkaStreams
	isSet bool
}

func (v NullableKafkaStreams) Get() *KafkaStreams {
	return v.value
}

func (v *NullableKafkaStreams) Set(val *KafkaStreams) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaStreams) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaStreams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaStreams(val *KafkaStreams) *NullableKafkaStreams {
	return &NullableKafkaStreams{value: val, isSet: true}
}

func (v NullableKafkaStreams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaStreams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


