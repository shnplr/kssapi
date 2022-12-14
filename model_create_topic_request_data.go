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

// CreateTopicRequestData struct for CreateTopicRequestData
type CreateTopicRequestData struct {
	TopicName *string `json:"topic_name,omitempty"`
	PartitionsCount *int32 `json:"partitions_count,omitempty"`
	ReplicationFactor *int32 `json:"replication_factor,omitempty"`
	ReplicasAssignments []CreateTopicRequestDataReplicasAssignmentsInner `json:"replicas_assignments,omitempty"`
	Configs []CreateTopicRequestDataConfigsInner `json:"configs,omitempty"`
}

// NewCreateTopicRequestData instantiates a new CreateTopicRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateTopicRequestData() *CreateTopicRequestData {
	this := CreateTopicRequestData{}
	return &this
}

// NewCreateTopicRequestDataWithDefaults instantiates a new CreateTopicRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateTopicRequestDataWithDefaults() *CreateTopicRequestData {
	this := CreateTopicRequestData{}
	return &this
}

// GetTopicName returns the TopicName field value if set, zero value otherwise.
func (o *CreateTopicRequestData) GetTopicName() string {
	if o == nil || isNil(o.TopicName) {
		var ret string
		return ret
	}
	return *o.TopicName
}

// GetTopicNameOk returns a tuple with the TopicName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTopicRequestData) GetTopicNameOk() (*string, bool) {
	if o == nil || isNil(o.TopicName) {
    return nil, false
	}
	return o.TopicName, true
}

// HasTopicName returns a boolean if a field has been set.
func (o *CreateTopicRequestData) HasTopicName() bool {
	if o != nil && !isNil(o.TopicName) {
		return true
	}

	return false
}

// SetTopicName gets a reference to the given string and assigns it to the TopicName field.
func (o *CreateTopicRequestData) SetTopicName(v string) {
	o.TopicName = &v
}

// GetPartitionsCount returns the PartitionsCount field value if set, zero value otherwise.
func (o *CreateTopicRequestData) GetPartitionsCount() int32 {
	if o == nil || isNil(o.PartitionsCount) {
		var ret int32
		return ret
	}
	return *o.PartitionsCount
}

// GetPartitionsCountOk returns a tuple with the PartitionsCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTopicRequestData) GetPartitionsCountOk() (*int32, bool) {
	if o == nil || isNil(o.PartitionsCount) {
    return nil, false
	}
	return o.PartitionsCount, true
}

// HasPartitionsCount returns a boolean if a field has been set.
func (o *CreateTopicRequestData) HasPartitionsCount() bool {
	if o != nil && !isNil(o.PartitionsCount) {
		return true
	}

	return false
}

// SetPartitionsCount gets a reference to the given int32 and assigns it to the PartitionsCount field.
func (o *CreateTopicRequestData) SetPartitionsCount(v int32) {
	o.PartitionsCount = &v
}

// GetReplicationFactor returns the ReplicationFactor field value if set, zero value otherwise.
func (o *CreateTopicRequestData) GetReplicationFactor() int32 {
	if o == nil || isNil(o.ReplicationFactor) {
		var ret int32
		return ret
	}
	return *o.ReplicationFactor
}

// GetReplicationFactorOk returns a tuple with the ReplicationFactor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTopicRequestData) GetReplicationFactorOk() (*int32, bool) {
	if o == nil || isNil(o.ReplicationFactor) {
    return nil, false
	}
	return o.ReplicationFactor, true
}

// HasReplicationFactor returns a boolean if a field has been set.
func (o *CreateTopicRequestData) HasReplicationFactor() bool {
	if o != nil && !isNil(o.ReplicationFactor) {
		return true
	}

	return false
}

// SetReplicationFactor gets a reference to the given int32 and assigns it to the ReplicationFactor field.
func (o *CreateTopicRequestData) SetReplicationFactor(v int32) {
	o.ReplicationFactor = &v
}

// GetReplicasAssignments returns the ReplicasAssignments field value if set, zero value otherwise.
func (o *CreateTopicRequestData) GetReplicasAssignments() []CreateTopicRequestDataReplicasAssignmentsInner {
	if o == nil || isNil(o.ReplicasAssignments) {
		var ret []CreateTopicRequestDataReplicasAssignmentsInner
		return ret
	}
	return o.ReplicasAssignments
}

// GetReplicasAssignmentsOk returns a tuple with the ReplicasAssignments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTopicRequestData) GetReplicasAssignmentsOk() ([]CreateTopicRequestDataReplicasAssignmentsInner, bool) {
	if o == nil || isNil(o.ReplicasAssignments) {
    return nil, false
	}
	return o.ReplicasAssignments, true
}

// HasReplicasAssignments returns a boolean if a field has been set.
func (o *CreateTopicRequestData) HasReplicasAssignments() bool {
	if o != nil && !isNil(o.ReplicasAssignments) {
		return true
	}

	return false
}

// SetReplicasAssignments gets a reference to the given []CreateTopicRequestDataReplicasAssignmentsInner and assigns it to the ReplicasAssignments field.
func (o *CreateTopicRequestData) SetReplicasAssignments(v []CreateTopicRequestDataReplicasAssignmentsInner) {
	o.ReplicasAssignments = v
}

// GetConfigs returns the Configs field value if set, zero value otherwise.
func (o *CreateTopicRequestData) GetConfigs() []CreateTopicRequestDataConfigsInner {
	if o == nil || isNil(o.Configs) {
		var ret []CreateTopicRequestDataConfigsInner
		return ret
	}
	return o.Configs
}

// GetConfigsOk returns a tuple with the Configs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateTopicRequestData) GetConfigsOk() ([]CreateTopicRequestDataConfigsInner, bool) {
	if o == nil || isNil(o.Configs) {
    return nil, false
	}
	return o.Configs, true
}

// HasConfigs returns a boolean if a field has been set.
func (o *CreateTopicRequestData) HasConfigs() bool {
	if o != nil && !isNil(o.Configs) {
		return true
	}

	return false
}

// SetConfigs gets a reference to the given []CreateTopicRequestDataConfigsInner and assigns it to the Configs field.
func (o *CreateTopicRequestData) SetConfigs(v []CreateTopicRequestDataConfigsInner) {
	o.Configs = v
}

func (o CreateTopicRequestData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TopicName) {
		toSerialize["topic_name"] = o.TopicName
	}
	if !isNil(o.PartitionsCount) {
		toSerialize["partitions_count"] = o.PartitionsCount
	}
	if !isNil(o.ReplicationFactor) {
		toSerialize["replication_factor"] = o.ReplicationFactor
	}
	if !isNil(o.ReplicasAssignments) {
		toSerialize["replicas_assignments"] = o.ReplicasAssignments
	}
	if !isNil(o.Configs) {
		toSerialize["configs"] = o.Configs
	}
	return json.Marshal(toSerialize)
}

type NullableCreateTopicRequestData struct {
	value *CreateTopicRequestData
	isSet bool
}

func (v NullableCreateTopicRequestData) Get() *CreateTopicRequestData {
	return v.value
}

func (v *NullableCreateTopicRequestData) Set(val *CreateTopicRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateTopicRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateTopicRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateTopicRequestData(val *CreateTopicRequestData) *NullableCreateTopicRequestData {
	return &NullableCreateTopicRequestData{value: val, isSet: true}
}

func (v NullableCreateTopicRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateTopicRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


