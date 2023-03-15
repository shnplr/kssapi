# ApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SingularName** | Pointer to **string** |  | [optional] 
**Namespaced** | Pointer to **bool** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Kind** | Pointer to **string** |  | [optional] 
**Verbs** | Pointer to **[]string** |  | [optional] 
**ShortNames** | Pointer to **[]string** |  | [optional] 

## Methods

### NewApiResource

`func NewApiResource() *ApiResource`

NewApiResource instantiates a new ApiResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiResourceWithDefaults

`func NewApiResourceWithDefaults() *ApiResource`

NewApiResourceWithDefaults instantiates a new ApiResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ApiResource) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApiResource) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApiResource) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApiResource) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSingularName

`func (o *ApiResource) GetSingularName() string`

GetSingularName returns the SingularName field if non-nil, zero value otherwise.

### GetSingularNameOk

`func (o *ApiResource) GetSingularNameOk() (*string, bool)`

GetSingularNameOk returns a tuple with the SingularName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSingularName

`func (o *ApiResource) SetSingularName(v string)`

SetSingularName sets SingularName field to given value.

### HasSingularName

`func (o *ApiResource) HasSingularName() bool`

HasSingularName returns a boolean if a field has been set.

### GetNamespaced

`func (o *ApiResource) GetNamespaced() bool`

GetNamespaced returns the Namespaced field if non-nil, zero value otherwise.

### GetNamespacedOk

`func (o *ApiResource) GetNamespacedOk() (*bool, bool)`

GetNamespacedOk returns a tuple with the Namespaced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespaced

`func (o *ApiResource) SetNamespaced(v bool)`

SetNamespaced sets Namespaced field to given value.

### HasNamespaced

`func (o *ApiResource) HasNamespaced() bool`

HasNamespaced returns a boolean if a field has been set.

### GetGroup

`func (o *ApiResource) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ApiResource) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ApiResource) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ApiResource) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetVersion

`func (o *ApiResource) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ApiResource) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ApiResource) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ApiResource) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetKind

`func (o *ApiResource) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ApiResource) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ApiResource) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *ApiResource) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetVerbs

`func (o *ApiResource) GetVerbs() []string`

GetVerbs returns the Verbs field if non-nil, zero value otherwise.

### GetVerbsOk

`func (o *ApiResource) GetVerbsOk() (*[]string, bool)`

GetVerbsOk returns a tuple with the Verbs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerbs

`func (o *ApiResource) SetVerbs(v []string)`

SetVerbs sets Verbs field to given value.

### HasVerbs

`func (o *ApiResource) HasVerbs() bool`

HasVerbs returns a boolean if a field has been set.

### GetShortNames

`func (o *ApiResource) GetShortNames() []string`

GetShortNames returns the ShortNames field if non-nil, zero value otherwise.

### GetShortNamesOk

`func (o *ApiResource) GetShortNamesOk() (*[]string, bool)`

GetShortNamesOk returns a tuple with the ShortNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortNames

`func (o *ApiResource) SetShortNames(v []string)`

SetShortNames sets ShortNames field to given value.

### HasShortNames

`func (o *ApiResource) HasShortNames() bool`

HasShortNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


