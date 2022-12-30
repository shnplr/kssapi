# ApiResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Namespaced** | Pointer to **bool** |  | [optional] 

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


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


