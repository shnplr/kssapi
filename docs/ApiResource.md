# ApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespaced** | Pointer to **bool** |  | [optional] 
**Verbs** | Pointer to **[]string** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


