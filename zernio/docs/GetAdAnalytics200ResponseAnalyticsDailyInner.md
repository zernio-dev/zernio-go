# GetAdAnalytics200ResponseAnalyticsDailyInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Spend** | Pointer to **float32** |  | [optional] 
**Impressions** | Pointer to **int32** |  | [optional] 
**Reach** | Pointer to **int32** |  | [optional] 
**Clicks** | Pointer to **int32** |  | [optional] 
**Ctr** | Pointer to **float32** | Click-through rate (%) | [optional] 
**Cpc** | Pointer to **float32** | Cost per click | [optional] 
**Cpm** | Pointer to **float32** | Cost per 1000 impressions | [optional] 
**Engagement** | Pointer to **int32** |  | [optional] 
**Conversions** | Pointer to **int32** | Count of conversion events matching the campaign&#39;s promoted_object.custom_event_type (PURCHASE, LEAD, etc.) over the requested date range. 0 for non-conversion campaigns or when no events have fired. Meta-only at time of writing; other platforms return 0. | [optional] 
**CostPerConversion** | Pointer to **float32** | Derived spend / conversions in the same currency as spend. 0 when conversions is 0. | [optional] 
**Actions** | Pointer to **map[string]int32** | Raw per-action-type counts from Meta&#39;s Insights actions[] array, summed over the date range. Keys are Meta action_type strings (e.g. link_click, offsite_conversion.fb_pixel_purchase, onsite_conversion.lead_grouped). Use this to extract any conversion event (purchases, leads, add_to_cart, etc.) without relying on the derived conversions field. Empty object when no actions are reported. | [optional] 
**ActionValues** | Pointer to **map[string]float32** | Monetary mirror of &#x60;actions&#x60;, from Meta&#39;s Insights &#x60;action_values[]&#x60; array. Same keying — values are the revenue attributed to each action_type, in ad-account native currency (same unit as &#x60;spend&#x60;; see the campaign node&#39;s &#x60;currency&#x60; field). Use this to compute revenue-per-event (e.g. avg purchase value). Meta-only; other platforms return {}. | [optional] 
**PurchaseValue** | Pointer to **float32** | Convenience sum of purchase-type action values — picked from &#x60;actionValues&#x60; via the same priority list as &#x60;conversions&#x60; so both fields describe the same events. In ad-account native currency. 0 when the campaign has no purchase event configured. Meta-only. | [optional] 
**Roas** | Pointer to **float32** | Return on ad spend — derived as &#x60;purchaseValue / spend&#x60;. 0 when &#x60;spend&#x60; is 0. Equivalent to Meta&#39;s &#x60;purchase_roas&#x60; under default attribution. At ad-set and campaign levels this is recomputed from summed purchaseValue + spend (NOT averaged across children) so it&#39;s mathematically correct at every rollup level. | [optional] 
**LastSyncedAt** | Pointer to **time.Time** | Present on individual ads only, not on campaign aggregations | [optional] 
**Date** | Pointer to **string** |  | [optional] 

## Methods

### NewGetAdAnalytics200ResponseAnalyticsDailyInner

`func NewGetAdAnalytics200ResponseAnalyticsDailyInner() *GetAdAnalytics200ResponseAnalyticsDailyInner`

NewGetAdAnalytics200ResponseAnalyticsDailyInner instantiates a new GetAdAnalytics200ResponseAnalyticsDailyInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdAnalytics200ResponseAnalyticsDailyInnerWithDefaults

`func NewGetAdAnalytics200ResponseAnalyticsDailyInnerWithDefaults() *GetAdAnalytics200ResponseAnalyticsDailyInner`

NewGetAdAnalytics200ResponseAnalyticsDailyInnerWithDefaults instantiates a new GetAdAnalytics200ResponseAnalyticsDailyInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpend

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetSpend() float32`

GetSpend returns the Spend field if non-nil, zero value otherwise.

### GetSpendOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetSpendOk() (*float32, bool)`

GetSpendOk returns a tuple with the Spend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpend

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetSpend(v float32)`

SetSpend sets Spend field to given value.

### HasSpend

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasSpend() bool`

HasSpend returns a boolean if a field has been set.

### GetImpressions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetImpressions() int32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetImpressionsOk() (*int32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetImpressions(v int32)`

SetImpressions sets Impressions field to given value.

### HasImpressions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasImpressions() bool`

HasImpressions returns a boolean if a field has been set.

### GetReach

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetReach() int32`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetReachOk() (*int32, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetReach(v int32)`

SetReach sets Reach field to given value.

### HasReach

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetClicks

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetClicks() int32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetClicksOk() (*int32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetClicks(v int32)`

SetClicks sets Clicks field to given value.

### HasClicks

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasClicks() bool`

HasClicks returns a boolean if a field has been set.

### GetCtr

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCtr() float32`

GetCtr returns the Ctr field if non-nil, zero value otherwise.

### GetCtrOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCtrOk() (*float32, bool)`

GetCtrOk returns a tuple with the Ctr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCtr

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetCtr(v float32)`

SetCtr sets Ctr field to given value.

### HasCtr

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasCtr() bool`

HasCtr returns a boolean if a field has been set.

### GetCpc

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCpc() float32`

GetCpc returns the Cpc field if non-nil, zero value otherwise.

### GetCpcOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCpcOk() (*float32, bool)`

GetCpcOk returns a tuple with the Cpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpc

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetCpc(v float32)`

SetCpc sets Cpc field to given value.

### HasCpc

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasCpc() bool`

HasCpc returns a boolean if a field has been set.

### GetCpm

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCpm() float32`

GetCpm returns the Cpm field if non-nil, zero value otherwise.

### GetCpmOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCpmOk() (*float32, bool)`

GetCpmOk returns a tuple with the Cpm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpm

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetCpm(v float32)`

SetCpm sets Cpm field to given value.

### HasCpm

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasCpm() bool`

HasCpm returns a boolean if a field has been set.

### GetEngagement

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetEngagement() int32`

GetEngagement returns the Engagement field if non-nil, zero value otherwise.

### GetEngagementOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetEngagementOk() (*int32, bool)`

GetEngagementOk returns a tuple with the Engagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagement

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetEngagement(v int32)`

SetEngagement sets Engagement field to given value.

### HasEngagement

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasEngagement() bool`

HasEngagement returns a boolean if a field has been set.

### GetConversions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetConversions() int32`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetConversionsOk() (*int32, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetConversions(v int32)`

SetConversions sets Conversions field to given value.

### HasConversions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasConversions() bool`

HasConversions returns a boolean if a field has been set.

### GetCostPerConversion

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCostPerConversion() float32`

GetCostPerConversion returns the CostPerConversion field if non-nil, zero value otherwise.

### GetCostPerConversionOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetCostPerConversionOk() (*float32, bool)`

GetCostPerConversionOk returns a tuple with the CostPerConversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostPerConversion

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetCostPerConversion(v float32)`

SetCostPerConversion sets CostPerConversion field to given value.

### HasCostPerConversion

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasCostPerConversion() bool`

HasCostPerConversion returns a boolean if a field has been set.

### GetActions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetActions() map[string]int32`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetActionsOk() (*map[string]int32, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetActions(v map[string]int32)`

SetActions sets Actions field to given value.

### HasActions

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetActionValues

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetActionValues() map[string]float32`

GetActionValues returns the ActionValues field if non-nil, zero value otherwise.

### GetActionValuesOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetActionValuesOk() (*map[string]float32, bool)`

GetActionValuesOk returns a tuple with the ActionValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionValues

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetActionValues(v map[string]float32)`

SetActionValues sets ActionValues field to given value.

### HasActionValues

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasActionValues() bool`

HasActionValues returns a boolean if a field has been set.

### GetPurchaseValue

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetPurchaseValue() float32`

GetPurchaseValue returns the PurchaseValue field if non-nil, zero value otherwise.

### GetPurchaseValueOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetPurchaseValueOk() (*float32, bool)`

GetPurchaseValueOk returns a tuple with the PurchaseValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseValue

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetPurchaseValue(v float32)`

SetPurchaseValue sets PurchaseValue field to given value.

### HasPurchaseValue

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasPurchaseValue() bool`

HasPurchaseValue returns a boolean if a field has been set.

### GetRoas

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetRoas() float32`

GetRoas returns the Roas field if non-nil, zero value otherwise.

### GetRoasOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetRoasOk() (*float32, bool)`

GetRoasOk returns a tuple with the Roas field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoas

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetRoas(v float32)`

SetRoas sets Roas field to given value.

### HasRoas

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasRoas() bool`

HasRoas returns a boolean if a field has been set.

### GetLastSyncedAt

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### GetDate

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetAdAnalytics200ResponseAnalyticsDailyInner) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


