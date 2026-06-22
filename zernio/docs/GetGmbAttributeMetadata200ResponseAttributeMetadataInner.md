# GetGmbAttributeMetadata200ResponseAttributeMetadataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parent** | Pointer to **string** | Resource name of the attribute (e.g. \&quot;attributes/has_delivery\&quot;). | [optional] 
**ValueType** | Pointer to **string** | Value type (e.g. BOOL, ENUM, URL, REPEATED_ENUM). | [optional] 
**DisplayName** | Pointer to **string** | Localized human-readable attribute name. | [optional] 
**GroupDisplayName** | Pointer to **string** | Display name of the attribute group. | [optional] 
**Repeatable** | Pointer to **bool** | True if multiple values can be set simultaneously. | [optional] 
**Deprecated** | Pointer to **bool** | True if this attribute should no longer be used. | [optional] 
**ValueMetadata** | Pointer to [**[]GetGmbAttributeMetadata200ResponseAttributeMetadataInnerValueMetadataInner**](GetGmbAttributeMetadata200ResponseAttributeMetadataInnerValueMetadataInner.md) | Possible enum values (for ENUM / REPEATED_ENUM types). | [optional] 

## Methods

### NewGetGmbAttributeMetadata200ResponseAttributeMetadataInner

`func NewGetGmbAttributeMetadata200ResponseAttributeMetadataInner() *GetGmbAttributeMetadata200ResponseAttributeMetadataInner`

NewGetGmbAttributeMetadata200ResponseAttributeMetadataInner instantiates a new GetGmbAttributeMetadata200ResponseAttributeMetadataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGmbAttributeMetadata200ResponseAttributeMetadataInnerWithDefaults

`func NewGetGmbAttributeMetadata200ResponseAttributeMetadataInnerWithDefaults() *GetGmbAttributeMetadata200ResponseAttributeMetadataInner`

NewGetGmbAttributeMetadata200ResponseAttributeMetadataInnerWithDefaults instantiates a new GetGmbAttributeMetadata200ResponseAttributeMetadataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParent

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetValueType

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetGroupDisplayName

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetGroupDisplayName() string`

GetGroupDisplayName returns the GroupDisplayName field if non-nil, zero value otherwise.

### GetGroupDisplayNameOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetGroupDisplayNameOk() (*string, bool)`

GetGroupDisplayNameOk returns a tuple with the GroupDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDisplayName

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetGroupDisplayName(v string)`

SetGroupDisplayName sets GroupDisplayName field to given value.

### HasGroupDisplayName

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasGroupDisplayName() bool`

HasGroupDisplayName returns a boolean if a field has been set.

### GetRepeatable

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetRepeatable() bool`

GetRepeatable returns the Repeatable field if non-nil, zero value otherwise.

### GetRepeatableOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetRepeatableOk() (*bool, bool)`

GetRepeatableOk returns a tuple with the Repeatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatable

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetRepeatable(v bool)`

SetRepeatable sets Repeatable field to given value.

### HasRepeatable

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasRepeatable() bool`

HasRepeatable returns a boolean if a field has been set.

### GetDeprecated

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetValueMetadata

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetValueMetadata() []GetGmbAttributeMetadata200ResponseAttributeMetadataInnerValueMetadataInner`

GetValueMetadata returns the ValueMetadata field if non-nil, zero value otherwise.

### GetValueMetadataOk

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) GetValueMetadataOk() (*[]GetGmbAttributeMetadata200ResponseAttributeMetadataInnerValueMetadataInner, bool)`

GetValueMetadataOk returns a tuple with the ValueMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueMetadata

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) SetValueMetadata(v []GetGmbAttributeMetadata200ResponseAttributeMetadataInnerValueMetadataInner)`

SetValueMetadata sets ValueMetadata field to given value.

### HasValueMetadata

`func (o *GetGmbAttributeMetadata200ResponseAttributeMetadataInner) HasValueMetadata() bool`

HasValueMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


