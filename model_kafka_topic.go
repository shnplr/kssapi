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

// KafkaTopic struct for KafkaTopic
type KafkaTopic struct {
	Kind *string `json:"kind,omitempty"`
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	PartitionCount *int32 `json:"partition_count,omitempty"`
	ReplicationFactor *int32 `json:"replication_factor,omitempty"`
	Config *map[string]ConfigItem `json:"config,omitempty"`
	Partitions []PartitionInfo `json:"partitions,omitempty"`
}

// NewKafkaTopic instantiates a new KafkaTopic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKafkaTopic(name string, namespace string) *KafkaTopic {
	this := KafkaTopic{}
	this.Name = name
	this.Namespace = namespace
	return &this
}

// NewKafkaTopicWithDefaults instantiates a new KafkaTopic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKafkaTopicWithDefaults() *KafkaTopic {
	this := KafkaTopic{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *KafkaTopic) GetKind() string {
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
    return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *KafkaTopic) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *KafkaTopic) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value
func (o *KafkaTopic) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetNameOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *KafkaTopic) SetName(v string) {
	o.Name = v
}

// GetNamespace returns the Namespace field value
func (o *KafkaTopic) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetNamespaceOk() (*string, bool) {
	if o == nil {
    return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *KafkaTopic) SetNamespace(v string) {
	o.Namespace = v
}

// GetPartitionCount returns the PartitionCount field value if set, zero value otherwise.
func (o *KafkaTopic) GetPartitionCount() int32 {
	if o == nil || isNil(o.PartitionCount) {
		var ret int32
		return ret
	}
	return *o.PartitionCount
}

// GetPartitionCountOk returns a tuple with the PartitionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetPartitionCountOk() (*int32, bool) {
	if o == nil || isNil(o.PartitionCount) {
    return nil, false
	}
	return o.PartitionCount, true
}

// HasPartitionCount returns a boolean if a field has been set.
func (o *KafkaTopic) HasPartitionCount() bool {
	if o != nil && !isNil(o.PartitionCount) {
		return true
	}

	return false
}

// SetPartitionCount gets a reference to the given int32 and assigns it to the PartitionCount field.
func (o *KafkaTopic) SetPartitionCount(v int32) {
	o.PartitionCount = &v
}

// GetReplicationFactor returns the ReplicationFactor field value if set, zero value otherwise.
func (o *KafkaTopic) GetReplicationFactor() int32 {
	if o == nil || isNil(o.ReplicationFactor) {
		var ret int32
		return ret
	}
	return *o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetReplicationFactorOk() (*int32, bool) {
	if o == nil || isNil(o.ReplicationFactor) {
    return nil, false
	}
	return o.ReplicationFactor, true
}

// HasReplicationFactor returns a boolean if a field has been set.
func (o *KafkaTopic) HasReplicationFactor() bool {
	if o != nil && !isNil(o.ReplicationFactor) {
		return true
	}

	return false
}

// SetReplicationFactor gets a reference to the given int32 and assigns it to the ReplicationFactor field.
func (o *KafkaTopic) SetReplicationFactor(v int32) {
	o.ReplicationFactor = &v
}

// GetConfig returns the Config field value if set, zero value otherwise.
func (o *KafkaTopic) GetConfig() map[string]ConfigItem {
	if o == nil || isNil(o.Config) {
		var ret map[string]ConfigItem
		return ret
	}
	return *o.Config
}

// GetConfigOk returns a tuple with the Config field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetConfigOk() (*map[string]ConfigItem, bool) {
	if o == nil || isNil(o.Config) {
    return nil, false
	}
	return o.Config, true
}

// HasConfig returns a boolean if a field has been set.
func (o *KafkaTopic) HasConfig() bool {
	if o != nil && !isNil(o.Config) {
		return true
	}

	return false
}

// SetConfig gets a reference to the given map[string]ConfigItem and assigns it to the Config field.
func (o *KafkaTopic) SetConfig(v map[string]ConfigItem) {
	o.Config = &v
}

// GetPartitions returns the Partitions field value if set, zero value otherwise.
func (o *KafkaTopic) GetPartitions() []PartitionInfo {
	if o == nil || isNil(o.Partitions) {
		var ret []PartitionInfo
		return ret
	}
	return o.Partitions
}

// GetPartitionsOk returns a tuple with the Partitions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KafkaTopic) GetPartitionsOk() ([]PartitionInfo, bool) {
	if o == nil || isNil(o.Partitions) {
    return nil, false
	}
	return o.Partitions, true
}

// HasPartitions returns a boolean if a field has been set.
func (o *KafkaTopic) HasPartitions() bool {
	if o != nil && !isNil(o.Partitions) {
		return true
	}

	return false
}

// SetPartitions gets a reference to the given []PartitionInfo and assigns it to the Partitions field.
func (o *KafkaTopic) SetPartitions(v []PartitionInfo) {
	o.Partitions = v
}

func (o KafkaTopic) MarshalJSON() ([]byte, error) {
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
	if !isNil(o.PartitionCount) {
		toSerialize["partition_count"] = o.PartitionCount
	}
	if !isNil(o.ReplicationFactor) {
		toSerialize["replication_factor"] = o.ReplicationFactor
	}
	if !isNil(o.Config) {
		toSerialize["config"] = o.Config
	}
	if !isNil(o.Partitions) {
		toSerialize["partitions"] = o.Partitions
	}
	return json.Marshal(toSerialize)
}

type NullableKafkaTopic struct {
	value *KafkaTopic
	isSet bool
}

func (v NullableKafkaTopic) Get() *KafkaTopic {
	return v.value
}

func (v *NullableKafkaTopic) Set(val *KafkaTopic) {
	v.value = val
	v.isSet = true
}

func (v NullableKafkaTopic) IsSet() bool {
	return v.isSet
}

func (v *NullableKafkaTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKafkaTopic(val *KafkaTopic) *NullableKafkaTopic {
	return &NullableKafkaTopic{value: val, isSet: true}
}

func (v NullableKafkaTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKafkaTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


