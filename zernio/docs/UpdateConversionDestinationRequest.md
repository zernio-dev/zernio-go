# UpdateConversionDestinationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdAccountId** | **string** |  | 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** | Setting &#x60;false&#x60; is equivalent to calling DELETE — the rule will appear as &#x60;inactive&#x60; afterwards.  | [optional] 
**AttributionType** | Pointer to **string** |  | [optional] 
**PostClickAttributionWindowSize** | Pointer to **int32** | 365 only allowed for LEAD, PURCHASE, ADD_TO_CART, QUALIFIED_LEAD, SUBMIT_APPLICATION rule types.  | [optional] 
**ViewThroughAttributionWindowSize** | Pointer to **int32** | 365 only allowed for LEAD, PURCHASE, ADD_TO_CART, QUALIFIED_LEAD, SUBMIT_APPLICATION rule types.  | [optional] 
**ValueType** | Pointer to **string** |  | [optional] 
**Value** | Pointer to [**UpdateConversionDestinationRequestValue**](UpdateConversionDestinationRequestValue.md) |  | [optional] 

## Methods

### NewUpdateConversionDestinationRequest

`func NewUpdateConversionDestinationRequest(adAccountId string, ) *UpdateConversionDestinationRequest`

NewUpdateConversionDestinationRequest instantiates a new UpdateConversionDestinationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConversionDestinationRequestWithDefaults

`func NewUpdateConversionDestinationRequestWithDefaults() *UpdateConversionDestinationRequest`

NewUpdateConversionDestinationRequestWithDefaults instantiates a new UpdateConversionDestinationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdAccountId

`func (o *UpdateConversionDestinationRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *UpdateConversionDestinationRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *UpdateConversionDestinationRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *UpdateConversionDestinationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateConversionDestinationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateConversionDestinationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateConversionDestinationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateConversionDestinationRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateConversionDestinationRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateConversionDestinationRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateConversionDestinationRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAttributionType

`func (o *UpdateConversionDestinationRequest) GetAttributionType() string`

GetAttributionType returns the AttributionType field if non-nil, zero value otherwise.

### GetAttributionTypeOk

`func (o *UpdateConversionDestinationRequest) GetAttributionTypeOk() (*string, bool)`

GetAttributionTypeOk returns a tuple with the AttributionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributionType

`func (o *UpdateConversionDestinationRequest) SetAttributionType(v string)`

SetAttributionType sets AttributionType field to given value.

### HasAttributionType

`func (o *UpdateConversionDestinationRequest) HasAttributionType() bool`

HasAttributionType returns a boolean if a field has been set.

### GetPostClickAttributionWindowSize

`func (o *UpdateConversionDestinationRequest) GetPostClickAttributionWindowSize() int32`

GetPostClickAttributionWindowSize returns the PostClickAttributionWindowSize field if non-nil, zero value otherwise.

### GetPostClickAttributionWindowSizeOk

`func (o *UpdateConversionDestinationRequest) GetPostClickAttributionWindowSizeOk() (*int32, bool)`

GetPostClickAttributionWindowSizeOk returns a tuple with the PostClickAttributionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostClickAttributionWindowSize

`func (o *UpdateConversionDestinationRequest) SetPostClickAttributionWindowSize(v int32)`

SetPostClickAttributionWindowSize sets PostClickAttributionWindowSize field to given value.

### HasPostClickAttributionWindowSize

`func (o *UpdateConversionDestinationRequest) HasPostClickAttributionWindowSize() bool`

HasPostClickAttributionWindowSize returns a boolean if a field has been set.

### GetViewThroughAttributionWindowSize

`func (o *UpdateConversionDestinationRequest) GetViewThroughAttributionWindowSize() int32`

GetViewThroughAttributionWindowSize returns the ViewThroughAttributionWindowSize field if non-nil, zero value otherwise.

### GetViewThroughAttributionWindowSizeOk

`func (o *UpdateConversionDestinationRequest) GetViewThroughAttributionWindowSizeOk() (*int32, bool)`

GetViewThroughAttributionWindowSizeOk returns a tuple with the ViewThroughAttributionWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewThroughAttributionWindowSize

`func (o *UpdateConversionDestinationRequest) SetViewThroughAttributionWindowSize(v int32)`

SetViewThroughAttributionWindowSize sets ViewThroughAttributionWindowSize field to given value.

### HasViewThroughAttributionWindowSize

`func (o *UpdateConversionDestinationRequest) HasViewThroughAttributionWindowSize() bool`

HasViewThroughAttributionWindowSize returns a boolean if a field has been set.

### GetValueType

`func (o *UpdateConversionDestinationRequest) GetValueType() string`

GetValueType returns the ValueType field if non-nil, zero value otherwise.

### GetValueTypeOk

`func (o *UpdateConversionDestinationRequest) GetValueTypeOk() (*string, bool)`

GetValueTypeOk returns a tuple with the ValueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueType

`func (o *UpdateConversionDestinationRequest) SetValueType(v string)`

SetValueType sets ValueType field to given value.

### HasValueType

`func (o *UpdateConversionDestinationRequest) HasValueType() bool`

HasValueType returns a boolean if a field has been set.

### GetValue

`func (o *UpdateConversionDestinationRequest) GetValue() UpdateConversionDestinationRequestValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateConversionDestinationRequest) GetValueOk() (*UpdateConversionDestinationRequestValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateConversionDestinationRequest) SetValue(v UpdateConversionDestinationRequestValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateConversionDestinationRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


