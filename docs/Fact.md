# Fact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Predicate** | Pointer to **string** |  | [optional] 
**Args** | Pointer to [**[]Value**](Value.md) |  | [optional] 

## Methods

### NewFact

`func NewFact() *Fact`

NewFact instantiates a new Fact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFactWithDefaults

`func NewFactWithDefaults() *Fact`

NewFactWithDefaults instantiates a new Fact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPredicate

`func (o *Fact) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *Fact) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *Fact) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *Fact) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### GetArgs

`func (o *Fact) GetArgs() []Value`

GetArgs returns the Args field if non-nil, zero value otherwise.

### GetArgsOk

`func (o *Fact) GetArgsOk() (*[]Value, bool)`

GetArgsOk returns a tuple with the Args field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgs

`func (o *Fact) SetArgs(v []Value)`

SetArgs sets Args field to given value.

### HasArgs

`func (o *Fact) HasArgs() bool`

HasArgs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


