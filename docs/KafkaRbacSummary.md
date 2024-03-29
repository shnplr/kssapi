# KafkaRbacSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**ApiVersion** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Namespace** | Pointer to **string** |  | [optional] 
**Principal** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 
**ResourceType** | Pointer to **string** |  | [optional] 
**ResourceName** | Pointer to **string** |  | [optional] 
**PatternType** | Pointer to **string** |  | [optional] 

## Methods

### NewKafkaRbacSummary

`func NewKafkaRbacSummary() *KafkaRbacSummary`

NewKafkaRbacSummary instantiates a new KafkaRbacSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKafkaRbacSummaryWithDefaults

`func NewKafkaRbacSummaryWithDefaults() *KafkaRbacSummary`

NewKafkaRbacSummaryWithDefaults instantiates a new KafkaRbacSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *KafkaRbacSummary) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *KafkaRbacSummary) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *KafkaRbacSummary) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *KafkaRbacSummary) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetApiVersion

`func (o *KafkaRbacSummary) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *KafkaRbacSummary) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *KafkaRbacSummary) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *KafkaRbacSummary) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetName

`func (o *KafkaRbacSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KafkaRbacSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KafkaRbacSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KafkaRbacSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *KafkaRbacSummary) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *KafkaRbacSummary) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *KafkaRbacSummary) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.

### HasNamespace

`func (o *KafkaRbacSummary) HasNamespace() bool`

HasNamespace returns a boolean if a field has been set.

### GetPrincipal

`func (o *KafkaRbacSummary) GetPrincipal() string`

GetPrincipal returns the Principal field if non-nil, zero value otherwise.

### GetPrincipalOk

`func (o *KafkaRbacSummary) GetPrincipalOk() (*string, bool)`

GetPrincipalOk returns a tuple with the Principal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrincipal

`func (o *KafkaRbacSummary) SetPrincipal(v string)`

SetPrincipal sets Principal field to given value.

### HasPrincipal

`func (o *KafkaRbacSummary) HasPrincipal() bool`

HasPrincipal returns a boolean if a field has been set.

### GetRole

`func (o *KafkaRbacSummary) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *KafkaRbacSummary) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *KafkaRbacSummary) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *KafkaRbacSummary) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetResourceType

`func (o *KafkaRbacSummary) GetResourceType() string`

GetResourceType returns the ResourceType field if non-nil, zero value otherwise.

### GetResourceTypeOk

`func (o *KafkaRbacSummary) GetResourceTypeOk() (*string, bool)`

GetResourceTypeOk returns a tuple with the ResourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceType

`func (o *KafkaRbacSummary) SetResourceType(v string)`

SetResourceType sets ResourceType field to given value.

### HasResourceType

`func (o *KafkaRbacSummary) HasResourceType() bool`

HasResourceType returns a boolean if a field has been set.

### GetResourceName

`func (o *KafkaRbacSummary) GetResourceName() string`

GetResourceName returns the ResourceName field if non-nil, zero value otherwise.

### GetResourceNameOk

`func (o *KafkaRbacSummary) GetResourceNameOk() (*string, bool)`

GetResourceNameOk returns a tuple with the ResourceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceName

`func (o *KafkaRbacSummary) SetResourceName(v string)`

SetResourceName sets ResourceName field to given value.

### HasResourceName

`func (o *KafkaRbacSummary) HasResourceName() bool`

HasResourceName returns a boolean if a field has been set.

### GetPatternType

`func (o *KafkaRbacSummary) GetPatternType() string`

GetPatternType returns the PatternType field if non-nil, zero value otherwise.

### GetPatternTypeOk

`func (o *KafkaRbacSummary) GetPatternTypeOk() (*string, bool)`

GetPatternTypeOk returns a tuple with the PatternType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPatternType

`func (o *KafkaRbacSummary) SetPatternType(v string)`

SetPatternType sets PatternType field to given value.

### HasPatternType

`func (o *KafkaRbacSummary) HasPatternType() bool`

HasPatternType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


