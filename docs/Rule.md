# Rule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verbs** | Pointer to **[]string** |  | [optional] 
**Resources** | Pointer to **[]string** |  | [optional] 
**ResourceNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewRule

`func NewRule() *Rule`

NewRule instantiates a new Rule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleWithDefaults

`func NewRuleWithDefaults() *Rule`

NewRuleWithDefaults instantiates a new Rule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerbs

`func (o *Rule) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *Rule) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *Rule) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.

### HasVerbs

`func (o *Rule) HasVerbs() bool`

HasVerbs returns a boolean if a field has been set.

### GetResources

`func (o *Rule) GetResources() []string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Rule) GetResourcesOk() (*[]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Rule) SetResources(v []string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Rule) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetResourceNames

`func (o *Rule) GetResourceNames() []string`

GetResourceNames returns the ResourceNames field if non-nil, zero value otherwise.

### GetResourceNamesOk

`func (o *Rule) GetResourceNamesOk() (*[]string, bool)`

GetResourceNamesOk returns a tuple with the ResourceNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceNames

`func (o *Rule) SetResourceNames(v []string)`

SetResourceNames sets ResourceNames field to given value.

### HasResourceNames

`func (o *Rule) HasResourceNames() bool`

HasResourceNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


