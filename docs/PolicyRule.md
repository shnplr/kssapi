# PolicyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourceNames** | Pointer to **[]string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**Verbs** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPolicyRule

`func NewPolicyRule() *PolicyRule`

NewPolicyRule instantiates a new PolicyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRuleWithDefaults

`func NewPolicyRuleWithDefaults() *PolicyRule`

NewPolicyRuleWithDefaults instantiates a new PolicyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourceNames

`func (o *PolicyRule) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *PolicyRule) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *PolicyRule) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.

### HasResourceNames

`func (o *PolicyRule) HasResourceNames() bool`

HasResourceNames returns a boolean if a field has been set.

### GetResources

`func (o *PolicyRule) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *PolicyRule) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *PolicyRule) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *PolicyRule) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetVerbs

`func (o *PolicyRule) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *PolicyRule) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *PolicyRule) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.

### HasVerbs

`func (o *PolicyRule) HasVerbs() bool`

HasVerbs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


