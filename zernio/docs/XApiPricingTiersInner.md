# XApiPricingTiersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | Pointer to **string** | Tier key derived from price (e.g. &#x60;x_api_005&#x60; for $0.005, &#x60;x_api_200&#x60; for $0.200). The first three keys map to the legacy &#x60;xApiCalls&#x60; aggregate; new tiers (e.g. &#x60;x_api_200&#x60; for the URL tier added April 2026) are surfaced here but not in the legacy shape.  | [optional] 
**PricePerCallUsd** | Pointer to **float32** |  | [optional] 
**OperationCount** | Pointer to **int32** |  | [optional] 

## Methods

### NewXApiPricingTiersInner

`func NewXApiPricingTiersInner() *XApiPricingTiersInner`

NewXApiPricingTiersInner instantiates a new XApiPricingTiersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXApiPricingTiersInnerWithDefaults

`func NewXApiPricingTiersInnerWithDefaults() *XApiPricingTiersInner`

NewXApiPricingTiersInnerWithDefaults instantiates a new XApiPricingTiersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *XApiPricingTiersInner) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *XApiPricingTiersInner) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *XApiPricingTiersInner) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *XApiPricingTiersInner) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetPricePerCallUsd

`func (o *XApiPricingTiersInner) GetPricePerCallUsd() float32`

GetPricePerCallUsd returns the PricePerCallUsd field if non-nil, zero value otherwise.

### GetPricePerCallUsdOk

`func (o *XApiPricingTiersInner) GetPricePerCallUsdOk() (*float32, bool)`

GetPricePerCallUsdOk returns a tuple with the PricePerCallUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerCallUsd

`func (o *XApiPricingTiersInner) SetPricePerCallUsd(v float32)`

SetPricePerCallUsd sets PricePerCallUsd field to given value.

### HasPricePerCallUsd

`func (o *XApiPricingTiersInner) HasPricePerCallUsd() bool`

HasPricePerCallUsd returns a boolean if a field has been set.

### GetOperationCount

`func (o *XApiPricingTiersInner) GetOperationCount() int32`

GetOperationCount returns the OperationCount field if non-nil, zero value otherwise.

### GetOperationCountOk

`func (o *XApiPricingTiersInner) GetOperationCountOk() (*int32, bool)`

GetOperationCountOk returns a tuple with the OperationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationCount

`func (o *XApiPricingTiersInner) SetOperationCount(v int32)`

SetOperationCount sets OperationCount field to given value.

### HasOperationCount

`func (o *XApiPricingTiersInner) HasOperationCount() bool`

HasOperationCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


