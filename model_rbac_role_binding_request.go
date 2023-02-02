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
	Name *string `json:"name,omitempty"`
	Namespace string `json:"namespace"`
	Role *string `json:"role,omitempty"`
	User *string `json:"user,omitempty"`
	Topics []string `json:"topics,omitempty"`
	Subjects []string `json:"subjects,omitempty"`
	Groups []string `json:"groups,omitempty"`
}

// NewRbacRoleBindingRequest instantiates a new RbacRoleBindingRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRbacRoleBindingRequest(namespace string) *RbacRoleBindingRequest {
	this := RbacRoleBindingRequest{}
	this.Namespace = namespace
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
	if o == nil || isNil(o.Kind) {
		var ret string
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetKindOk() (*string, bool) {
	if o == nil || isNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasKind() bool {
	if o != nil && !isNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given string and assigns it to the Kind field.
func (o *RbacRoleBindingRequest) SetKind(v string) {
	o.Kind = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RbacRoleBindingRequest) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value
func (o *RbacRoleBindingRequest) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *RbacRoleBindingRequest) SetNamespace(v string) {
	o.Namespace = v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetRole() string {
	if o == nil || isNil(o.Role) {
		var ret string
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetRoleOk() (*string, bool) {
	if o == nil || isNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// HasRole returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasRole() bool {
	if o != nil && !isNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given string and assigns it to the Role field.
func (o *RbacRoleBindingRequest) SetRole(v string) {
	o.Role = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetUser() string {
	if o == nil || isNil(o.User) {
		var ret string
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetUserOk() (*string, bool) {
	if o == nil || isNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasUser() bool {
	if o != nil && !isNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given string and assigns it to the User field.
func (o *RbacRoleBindingRequest) SetUser(v string) {
	o.User = &v
}

// GetTopics returns the Topics field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetTopics() []string {
	if o == nil || isNil(o.Topics) {
		var ret []string
		return ret
	}
	return o.Topics
}

// GetTopicsOk returns a tuple with the Topics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetTopicsOk() ([]string, bool) {
	if o == nil || isNil(o.Topics) {
		return nil, false
	}
	return o.Topics, true
}

// HasTopics returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasTopics() bool {
	if o != nil && !isNil(o.Topics) {
		return true
	}

	return false
}

// SetTopics gets a reference to the given []string and assigns it to the Topics field.
func (o *RbacRoleBindingRequest) SetTopics(v []string) {
	o.Topics = v
}

// GetSubjects returns the Subjects field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetSubjects() []string {
	if o == nil || isNil(o.Subjects) {
		var ret []string
		return ret
	}
	return o.Subjects
}

// GetSubjectsOk returns a tuple with the Subjects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetSubjectsOk() ([]string, bool) {
	if o == nil || isNil(o.Subjects) {
		return nil, false
	}
	return o.Subjects, true
}

// HasSubjects returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasSubjects() bool {
	if o != nil && !isNil(o.Subjects) {
		return true
	}

	return false
}

// SetSubjects gets a reference to the given []string and assigns it to the Subjects field.
func (o *RbacRoleBindingRequest) SetSubjects(v []string) {
	o.Subjects = v
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *RbacRoleBindingRequest) GetGroups() []string {
	if o == nil || isNil(o.Groups) {
		var ret []string
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RbacRoleBindingRequest) GetGroupsOk() ([]string, bool) {
	if o == nil || isNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *RbacRoleBindingRequest) HasGroups() bool {
	if o != nil && !isNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []string and assigns it to the Groups field.
func (o *RbacRoleBindingRequest) SetGroups(v []string) {
	o.Groups = v
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
	if !isNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["namespace"] = o.Namespace
	if !isNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !isNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !isNil(o.Topics) {
		toSerialize["topics"] = o.Topics
	}
	if !isNil(o.Subjects) {
		toSerialize["subjects"] = o.Subjects
	}
	if !isNil(o.Groups) {
		toSerialize["groups"] = o.Groups
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


