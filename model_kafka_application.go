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

// checks if the KafkaApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &KafkaApplication{}

// KafkaApplication struct for KafkaApplication
type KafkaApplication struct {
	Kind *string `json:"kind,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	Name *string `json:"name,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
	Type string `json:"type"`
	Principal string `json:"principal"`
	Transactional *bool `json:"transactional,omitempty"`
	Idempotent *bool `json:"idempotent,omitempty"`
}

// NewKafkaApplication instantiates a new KafkaApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaApplication(type_ string, principal string) *KafkaApplication {
	this := KafkaApplication{}
	this.Type = type_
	this.Principal = principal
	return &this
}

// NewKafkaApplicationWithDefaults instantiates a new KafkaApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaApplicationWithDefaults() *KafkaApplication {
	this := KafkaApplication{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KafkaApplication) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KafkaApplication) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KafkaApplication) SetKind(v string) {
	o.Kind = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *KafkaApplication) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *KafkaApplication) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *KafkaApplication) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *KafkaApplication) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *KafkaApplication) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *KafkaApplication) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *KafkaApplication) GetNamespace() string {
	if o == nil || IsNil(o.Namespace) {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetNamespaceOk() (*string, bool) {
	if o == nil || IsNil(o.Namespace) {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *KafkaApplication) HasNamespace() bool {
	if o != nil && !IsNil(o.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *KafkaApplication) SetNamespace(v string) {
	o.Namespace = &v
}

// GetType returns the Type field value
func (o *KafkaApplication) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *KafkaApplication) SetType(v string) {
	o.Type = v
}

// GetPrincipal returns the Principal field value
func (o *KafkaApplication) GetPrincipal() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Principal
}

// GetPrincipalOk returns a tuple with the Principal field value
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetPrincipalOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Principal, true
}

// SetPrincipal sets field value
func (o *KafkaApplication) SetPrincipal(v string) {
	o.Principal = v
}

// GetTransactional returns the Transactional field value if set, zero value otherwise.
func (o *KafkaApplication) GetTransactional() bool {
	if o == nil || IsNil(o.Transactional) {
		var ret bool
		return ret
	}
	return *o.Transactional
}

// GetTransactionalOk returns a tuple with the Transactional field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetTransactionalOk() (*bool, bool) {
	if o == nil || IsNil(o.Transactional) {
		return nil, false
	}
	return o.Transactional, true
}

// HasTransactional returns a boolean if a field has been set.
func (o *KafkaApplication) HasTransactional() bool {
	if o != nil && !IsNil(o.Transactional) {
		return true
	}

	return false
}

// SetTransactional gets a reference to the given bool and assigns it to the Transactional field.
func (o *KafkaApplication) SetTransactional(v bool) {
	o.Transactional = &v
}

// GetIdempotent returns the Idempotent field value if set, zero value otherwise.
func (o *KafkaApplication) GetIdempotent() bool {
	if o == nil || IsNil(o.Idempotent) {
		var ret bool
		return ret
	}
	return *o.Idempotent
}

// GetIdempotentOk returns a tuple with the Idempotent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaApplication) GetIdempotentOk() (*bool, bool) {
	if o == nil || IsNil(o.Idempotent) {
		return nil, false
	}
	return o.Idempotent, true
}

// HasIdempotent returns a boolean if a field has been set.
func (o *KafkaApplication) HasIdempotent() bool {
	if o != nil && !IsNil(o.Idempotent) {
		return true
	}

	return false
}

// SetIdempotent gets a reference to the given bool and assigns it to the Idempotent field.
func (o *KafkaApplication) SetIdempotent(v bool) {
	o.Idempotent = &v
}

func (o KafkaApplication) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o KafkaApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	toSerialize["type"] = o.Type
	toSerialize["principal"] = o.Principal
	if !IsNil(o.Transactional) {
		toSerialize["transactional"] = o.Transactional
	}
	if !IsNil(o.Idempotent) {
		toSerialize["idempotent"] = o.Idempotent
	}
	return toSerialize, nil
}

type NullableKafkaApplication struct {
	value *KafkaApplication
	isSet bool
}

func (v NullableKafkaApplication) Get() *KafkaApplication {
	return v.value
}

func (v *NullableKafkaApplication) Set(val *KafkaApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaApplication(val *KafkaApplication) *NullableKafkaApplication {
	return &NullableKafkaApplication{value: val, isSet: true}
}

func (v NullableKafkaApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


