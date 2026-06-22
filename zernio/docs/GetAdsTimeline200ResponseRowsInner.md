# GetAdsTimeline200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Spend** | Pointer to **float32** | Native currency units (matches /ads/tree convention). | [optional] 
**Impressions** | Pointer to **int32** |  | [optional] 
**Reach** | Pointer to **int32** |  | [optional] 
**Clicks** | Pointer to **int32** |  | [optional] 
**Engagement** | Pointer to **int32** |  | [optional] 
**Ctr** | Pointer to **float32** | Click-through rate as a percentage (0–100). | [optional] 
**Cpc** | Pointer to **float32** | Cost per click in native currency. | [optional] 
**Cpm** | Pointer to **float32** | Cost per 1000 impressions in native currency. | [optional] 
**Conversions** | Pointer to **int32** | Sum of conversion events matching the campaign optimization goal. Meta-only at time of writing. | [optional] 
**CostPerConversion** | Pointer to **float32** |  | [optional] 
**Actions** | Pointer to **map[string]float32** | Per-action-type counts merged across all ads on this day. Keys are platform-native action types. | [optional] 
**ActionValues** | Pointer to **map[string]float32** | Monetary mirror of &#x60;actions&#x60; in native currency. | [optional] 
**PurchaseValue** | Pointer to **float32** | Sum of purchase-type action values on this day, native currency. | [optional] 
**Roas** | Pointer to **float32** | Derived purchaseValue / spend. | [optional] 

## Methods

### NewGetAdsTimeline200ResponseRowsInner

`func NewGetAdsTimeline200ResponseRowsInner() *GetAdsTimeline200ResponseRowsInner`

NewGetAdsTimeline200ResponseRowsInner instantiates a new GetAdsTimeline200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdsTimeline200ResponseRowsInnerWithDefaults

`func NewGetAdsTimeline200ResponseRowsInnerWithDefaults() *GetAdsTimeline200ResponseRowsInner`

NewGetAdsTimeline200ResponseRowsInnerWithDefaults instantiates a new GetAdsTimeline200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetAdsTimeline200ResponseRowsInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetAdsTimeline200ResponseRowsInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetAdsTimeline200ResponseRowsInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetSpend

`func (o *GetAdsTimeline200ResponseRowsInner) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *GetAdsTimeline200ResponseRowsInner) SetSpend(v float32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *GetAdsTimeline200ResponseRowsInner) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetImpressions

`func (o *GetAdsTimeline200ResponseRowsInner) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *GetAdsTimeline200ResponseRowsInner) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *GetAdsTimeline200ResponseRowsInner) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetReach

`func (o *GetAdsTimeline200ResponseRowsInner) GetReach() int32`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetReachOk() (*int32, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *GetAdsTimeline200ResponseRowsInner) SetReach(v int32)`

SetReach sets Reach field to given value.

### HasReach

`func (o *GetAdsTimeline200ResponseRowsInner) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetClicks

`func (o *GetAdsTimeline200ResponseRowsInner) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetAdsTimeline200ResponseRowsInner) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *GetAdsTimeline200ResponseRowsInner) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetEngagement

`func (o *GetAdsTimeline200ResponseRowsInner) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *GetAdsTimeline200ResponseRowsInner) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *GetAdsTimeline200ResponseRowsInner) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetCtr

`func (o *GetAdsTimeline200ResponseRowsInner) GetCtr() float32`

GetCtr returns the Ctr field if non-nil, zero value otherwise.

### GetCtrOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetCtrOk() (*float32, bool)`

GetCtrOk returns a tuple with the Ctr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtr

`func (o *GetAdsTimeline200ResponseRowsInner) SetCtr(v float32)`

SetCtr sets Ctr field to given value.

### HasCtr

`func (o *GetAdsTimeline200ResponseRowsInner) HasCtr() bool`

HasCtr returns a boolean if a field has been set.

### GetCpc

`func (o *GetAdsTimeline200ResponseRowsInner) GetCpc() float32`

GetCpc returns the Cpc field if non-nil, zero value otherwise.

### GetCpcOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetCpcOk() (*float32, bool)`

GetCpcOk returns a tuple with the Cpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpc

`func (o *GetAdsTimeline200ResponseRowsInner) SetCpc(v float32)`

SetCpc sets Cpc field to given value.

### HasCpc

`func (o *GetAdsTimeline200ResponseRowsInner) HasCpc() bool`

HasCpc returns a boolean if a field has been set.

### GetCpm

`func (o *GetAdsTimeline200ResponseRowsInner) GetCpm() float32`

GetCpm returns the Cpm field if non-nil, zero value otherwise.

### GetCpmOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetCpmOk() (*float32, bool)`

GetCpmOk returns a tuple with the Cpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpm

`func (o *GetAdsTimeline200ResponseRowsInner) SetCpm(v float32)`

SetCpm sets Cpm field to given value.

### HasCpm

`func (o *GetAdsTimeline200ResponseRowsInner) HasCpm() bool`

HasCpm returns a boolean if a field has been set.

### GetConversions

`func (o *GetAdsTimeline200ResponseRowsInner) GetConversions() int32`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetConversionsOk() (*int32, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *GetAdsTimeline200ResponseRowsInner) SetConversions(v int32)`

SetConversions sets Conversions field to given value.

### HasConversions

`func (o *GetAdsTimeline200ResponseRowsInner) HasConversions() bool`

HasConversions returns a boolean if a field has been set.

### GetCostPerConversion

`func (o *GetAdsTimeline200ResponseRowsInner) GetCostPerConversion() float32`

GetCostPerConversion returns the CostPerConversion field if non-nil, zero value otherwise.

### GetCostPerConversionOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetCostPerConversionOk() (*float32, bool)`

GetCostPerConversionOk returns a tuple with the CostPerConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerConversion

`func (o *GetAdsTimeline200ResponseRowsInner) SetCostPerConversion(v float32)`

SetCostPerConversion sets CostPerConversion field to given value.

### HasCostPerConversion

`func (o *GetAdsTimeline200ResponseRowsInner) HasCostPerConversion() bool`

HasCostPerConversion returns a boolean if a field has been set.

### GetActions

`func (o *GetAdsTimeline200ResponseRowsInner) GetActions() map[string]float32`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetActionsOk() (*map[string]float32, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetAdsTimeline200ResponseRowsInner) SetActions(v map[string]float32)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetAdsTimeline200ResponseRowsInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActionValues

`func (o *GetAdsTimeline200ResponseRowsInner) GetActionValues() map[string]float32`

GetActionValues returns the ActionValues field if non-nil, zero value otherwise.

### GetActionValuesOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetActionValuesOk() (*map[string]float32, bool)`

GetActionValuesOk returns a tuple with the ActionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValues

`func (o *GetAdsTimeline200ResponseRowsInner) SetActionValues(v map[string]float32)`

SetActionValues sets ActionValues field to given value.

### HasActionValues

`func (o *GetAdsTimeline200ResponseRowsInner) HasActionValues() bool`

HasActionValues returns a boolean if a field has been set.

### GetPurchaseValue

`func (o *GetAdsTimeline200ResponseRowsInner) GetPurchaseValue() float32`

GetPurchaseValue returns the PurchaseValue field if non-nil, zero value otherwise.

### GetPurchaseValueOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetPurchaseValueOk() (*float32, bool)`

GetPurchaseValueOk returns a tuple with the PurchaseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseValue

`func (o *GetAdsTimeline200ResponseRowsInner) SetPurchaseValue(v float32)`

SetPurchaseValue sets PurchaseValue field to given value.

### HasPurchaseValue

`func (o *GetAdsTimeline200ResponseRowsInner) HasPurchaseValue() bool`

HasPurchaseValue returns a boolean if a field has been set.

### GetRoas

`func (o *GetAdsTimeline200ResponseRowsInner) GetRoas() float32`

GetRoas returns the Roas field if non-nil, zero value otherwise.

### GetRoasOk

`func (o *GetAdsTimeline200ResponseRowsInner) GetRoasOk() (*float32, bool)`

GetRoasOk returns a tuple with the Roas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoas

`func (o *GetAdsTimeline200ResponseRowsInner) SetRoas(v float32)`

SetRoas sets Roas field to given value.

### HasRoas

`func (o *GetAdsTimeline200ResponseRowsInner) HasRoas() bool`

HasRoas returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


