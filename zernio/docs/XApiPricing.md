# XApiPricing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Currency** | Pointer to **string** |  | [optional] 
**Markup** | Pointer to **string** | Always 0% — Zernio does not mark up X API rates. | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**LastVerified** | Pointer to **string** | Date the prices were last verified against X&#39;s published rates. | [optional] 
**Tiers** | Pointer to [**[]XApiPricingTiersInner**](XApiPricingTiersInner.md) | Rollup of operations grouped by their per-call price. | [optional] 
**Operations** | Pointer to [**[]XApiOperation**](XApiOperation.md) | Flat list of every X operation Zernio can perform, with its rate. | [optional] 

## Methods

### NewXApiPricing

`func NewXApiPricing() *XApiPricing`

NewXApiPricing instantiates a new XApiPricing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewXApiPricingWithDefaults

`func NewXApiPricingWithDefaults() *XApiPricing`

NewXApiPricingWithDefaults instantiates a new XApiPricing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrency

`func (o *XApiPricing) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *XApiPricing) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *XApiPricing) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *XApiPricing) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetMarkup

`func (o *XApiPricing) GetMarkup() string`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *XApiPricing) GetMarkupOk() (*string, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *XApiPricing) SetMarkup(v string)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *XApiPricing) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetSource

`func (o *XApiPricing) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *XApiPricing) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *XApiPricing) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *XApiPricing) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetLastVerified

`func (o *XApiPricing) GetLastVerified() string`

GetLastVerified returns the LastVerified field if non-nil, zero value otherwise.

### GetLastVerifiedOk

`func (o *XApiPricing) GetLastVerifiedOk() (*string, bool)`

GetLastVerifiedOk returns a tuple with the LastVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastVerified

`func (o *XApiPricing) SetLastVerified(v string)`

SetLastVerified sets LastVerified field to given value.

### HasLastVerified

`func (o *XApiPricing) HasLastVerified() bool`

HasLastVerified returns a boolean if a field has been set.

### GetTiers

`func (o *XApiPricing) GetTiers() []XApiPricingTiersInner`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *XApiPricing) GetTiersOk() (*[]XApiPricingTiersInner, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *XApiPricing) SetTiers(v []XApiPricingTiersInner)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *XApiPricing) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetOperations

`func (o *XApiPricing) GetOperations() []XApiOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *XApiPricing) GetOperationsOk() (*[]XApiOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *XApiPricing) SetOperations(v []XApiOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *XApiPricing) HasOperations() bool`

HasOperations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


