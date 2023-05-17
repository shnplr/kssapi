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

// checks if the RbacRoleBindingRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RbacRoleBindingRequest{}

// RbacRoleBindingRequest struct for RbacRoleBindingRequest
type RbacRoleBindingRequest struct {
	Kind *string `json:"kind,omitempty"`
	ApiVersion *string `json:"apiVersion,omitempty"`
	Role string `json:"role"`
	User string `json:"user"`
	Topics []string `json:"topics,omitempty"`
	Metadata *ObjectMeta `json:"metadata,omitempty"`
	Subjects []string `json:"subjects,omitempty"`
}

// NewRbacRoleBindingRequest instantiates a new RbacRoleBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacRoleBindingRequest(role string, user string) *RbacRoleBindingRequest {
	this := RbacRoleBindingRequest{}
	this.Role = role
	this.User = user
	return &this
}

// NewRbacRoleBindingRequestWithDefaults instantiates a new RbacRoleBindingRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRbacRoleBindingRequestWithDefaults() *RbacRoleBindingRequest {
	this := RbacRoleBindingRequest{}
	return &this
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetKind() string {
	if o == nil || IsNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetKindOk() (*string, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RbacRoleBindingRequest) SetKind(v string) {
	o.Kind = &v
}

// GetApiVersion returns the ApiVersion field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetApiVersion() string {
	if o == nil || IsNil(o.ApiVersion) {
		var ret string
		return ret
	}
	return *o.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetApiVersionOk() (*string, bool) {
	if o == nil || IsNil(o.ApiVersion) {
		return nil, false
	}
	return o.ApiVersion, true
}

// HasApiVersion returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasApiVersion() bool {
	if o != nil && !IsNil(o.ApiVersion) {
		return true
	}

	return false
}

// SetApiVersion gets a reference to the given string and assigns it to the ApiVersion field.
func (o *RbacRoleBindingRequest) SetApiVersion(v string) {
	o.ApiVersion = &v
}

// GetRole returns the Role field value
func (o *RbacRoleBindingRequest) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *RbacRoleBindingRequest) SetRole(v string) {
	o.Role = v
}

// GetUser returns the User field value
func (o *RbacRoleBindingRequest) GetUser() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetUserOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RbacRoleBindingRequest) SetUser(v string) {
	o.User = v
}

// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetTopics() []string {
	if o == nil || IsNil(o.Topics) {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetTopicsOk() ([]string, bool) {
	if o == nil || IsNil(o.Topics) {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasTopics() bool {
	if o != nil && !IsNil(o.Topics) {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []string and assigns it to the Topics field.
func (o *RbacRoleBindingRequest) SetTopics(v []string) {
	o.Topics = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetMetadata() ObjectMeta {
	if o == nil || IsNil(o.Metadata) {
		var ret ObjectMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetMetadataOk() (*ObjectMeta, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given ObjectMeta and assigns it to the Metadata field.
func (o *RbacRoleBindingRequest) SetMetadata(v ObjectMeta) {
	o.Metadata = &v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetSubjects() []string {
	if o == nil || IsNil(o.Subjects) {
		var ret []string
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetSubjectsOk() ([]string, bool) {
	if o == nil || IsNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasSubjects() bool {
	if o != nil && !IsNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []string and assigns it to the Subjects field.
func (o *RbacRoleBindingRequest) SetSubjects(v []string) {
	o.Subjects = v
}

func (o RbacRoleBindingRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RbacRoleBindingRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.ApiVersion) {
		toSerialize["apiVersion"] = o.ApiVersion
	}
	toSerialize["role"] = o.Role
	toSerialize["user"] = o.User
	if !IsNil(o.Topics) {
		toSerialize["topics"] = o.Topics
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	return toSerialize, nil
}

type NullableRbacRoleBindingRequest struct {
	value *RbacRoleBindingRequest
	isSet bool
}

func (v NullableRbacRoleBindingRequest) Get() *RbacRoleBindingRequest {
	return v.value
}

func (v *NullableRbacRoleBindingRequest) Set(val *RbacRoleBindingRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRbacRoleBindingRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRbacRoleBindingRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRbacRoleBindingRequest(val *RbacRoleBindingRequest) *NullableRbacRoleBindingRequest {
	return &NullableRbacRoleBindingRequest{value: val, isSet: true}
}

func (v NullableRbacRoleBindingRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRbacRoleBindingRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


