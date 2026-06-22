# AdjustConversionsRequestAdjustmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdjustmentType** | **string** |  | 
**AdjustmentTime** | **float32** | When the adjustment occurred, unix seconds. | 
**OrderId** | Pointer to **string** | Transaction ID of the original conversion (the &#x60;eventId&#x60; you sent). Recommended; required for ENHANCEMENT. | [optional] 
**Gclid** | Pointer to **string** | Alternative key — the original click ID. Pair with &#x60;conversionTime&#x60;. Not valid for ENHANCEMENT. | [optional] 
**ConversionTime** | Pointer to **float32** | The original conversion&#39;s time, unix seconds. Required when identifying by &#x60;gclid&#x60;. | [optional] 
**RestatementValue** | Pointer to **float32** | RESTATEMENT only — the corrected TOTAL conversion value. | [optional] 
**Currency** | Pointer to **string** | RESTATEMENT only — ISO 4217 currency for &#x60;restatementValue&#x60;. | [optional] 
**User** | Pointer to [**AdjustConversionsRequestAdjustmentsInnerUser**](AdjustConversionsRequestAdjustmentsInnerUser.md) |  | [optional] 
**UserAgent** | Pointer to **string** | ENHANCEMENT only — the original conversion&#39;s user agent (improves match quality). | [optional] 

## Methods

### NewAdjustConversionsRequestAdjustmentsInner

`func NewAdjustConversionsRequestAdjustmentsInner(adjustmentType string, adjustmentTime float32, ) *AdjustConversionsRequestAdjustmentsInner`

NewAdjustConversionsRequestAdjustmentsInner instantiates a new AdjustConversionsRequestAdjustmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdjustConversionsRequestAdjustmentsInnerWithDefaults

`func NewAdjustConversionsRequestAdjustmentsInnerWithDefaults() *AdjustConversionsRequestAdjustmentsInner`

NewAdjustConversionsRequestAdjustmentsInnerWithDefaults instantiates a new AdjustConversionsRequestAdjustmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdjustmentType

`func (o *AdjustConversionsRequestAdjustmentsInner) GetAdjustmentType() string`

GetAdjustmentType returns the AdjustmentType field if non-nil, zero value otherwise.

### GetAdjustmentTypeOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetAdjustmentTypeOk() (*string, bool)`

GetAdjustmentTypeOk returns a tuple with the AdjustmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentType

`func (o *AdjustConversionsRequestAdjustmentsInner) SetAdjustmentType(v string)`

SetAdjustmentType sets AdjustmentType field to given value.


### GetAdjustmentTime

`func (o *AdjustConversionsRequestAdjustmentsInner) GetAdjustmentTime() float32`

GetAdjustmentTime returns the AdjustmentTime field if non-nil, zero value otherwise.

### GetAdjustmentTimeOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetAdjustmentTimeOk() (*float32, bool)`

GetAdjustmentTimeOk returns a tuple with the AdjustmentTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdjustmentTime

`func (o *AdjustConversionsRequestAdjustmentsInner) SetAdjustmentTime(v float32)`

SetAdjustmentTime sets AdjustmentTime field to given value.


### GetOrderId

`func (o *AdjustConversionsRequestAdjustmentsInner) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *AdjustConversionsRequestAdjustmentsInner) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *AdjustConversionsRequestAdjustmentsInner) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetGclid

`func (o *AdjustConversionsRequestAdjustmentsInner) GetGclid() string`

GetGclid returns the Gclid field if non-nil, zero value otherwise.

### GetGclidOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetGclidOk() (*string, bool)`

GetGclidOk returns a tuple with the Gclid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGclid

`func (o *AdjustConversionsRequestAdjustmentsInner) SetGclid(v string)`

SetGclid sets Gclid field to given value.

### HasGclid

`func (o *AdjustConversionsRequestAdjustmentsInner) HasGclid() bool`

HasGclid returns a boolean if a field has been set.

### GetConversionTime

`func (o *AdjustConversionsRequestAdjustmentsInner) GetConversionTime() float32`

GetConversionTime returns the ConversionTime field if non-nil, zero value otherwise.

### GetConversionTimeOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetConversionTimeOk() (*float32, bool)`

GetConversionTimeOk returns a tuple with the ConversionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversionTime

`func (o *AdjustConversionsRequestAdjustmentsInner) SetConversionTime(v float32)`

SetConversionTime sets ConversionTime field to given value.

### HasConversionTime

`func (o *AdjustConversionsRequestAdjustmentsInner) HasConversionTime() bool`

HasConversionTime returns a boolean if a field has been set.

### GetRestatementValue

`func (o *AdjustConversionsRequestAdjustmentsInner) GetRestatementValue() float32`

GetRestatementValue returns the RestatementValue field if non-nil, zero value otherwise.

### GetRestatementValueOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetRestatementValueOk() (*float32, bool)`

GetRestatementValueOk returns a tuple with the RestatementValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestatementValue

`func (o *AdjustConversionsRequestAdjustmentsInner) SetRestatementValue(v float32)`

SetRestatementValue sets RestatementValue field to given value.

### HasRestatementValue

`func (o *AdjustConversionsRequestAdjustmentsInner) HasRestatementValue() bool`

HasRestatementValue returns a boolean if a field has been set.

### GetCurrency

`func (o *AdjustConversionsRequestAdjustmentsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AdjustConversionsRequestAdjustmentsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AdjustConversionsRequestAdjustmentsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUser

`func (o *AdjustConversionsRequestAdjustmentsInner) GetUser() AdjustConversionsRequestAdjustmentsInnerUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetUserOk() (*AdjustConversionsRequestAdjustmentsInnerUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AdjustConversionsRequestAdjustmentsInner) SetUser(v AdjustConversionsRequestAdjustmentsInnerUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AdjustConversionsRequestAdjustmentsInner) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUserAgent

`func (o *AdjustConversionsRequestAdjustmentsInner) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *AdjustConversionsRequestAdjustmentsInner) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *AdjustConversionsRequestAdjustmentsInner) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *AdjustConversionsRequestAdjustmentsInner) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


