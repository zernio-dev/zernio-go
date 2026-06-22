# BoostPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PostId** | Pointer to **string** | Zernio post ID (provide this or platformPostId) | [optional] 
**PlatformPostId** | Pointer to **string** | Platform post ID (alternative to postId) | [optional] 
**AccountId** | **string** | Social account ID | 
**AdAccountId** | **string** | Platform ad account ID | 
**Name** | **string** |  | 
**Goal** | **string** | Available goals vary by platform. Meta (Facebook/Instagram) and TikTok support all 7. LinkedIn supports all except app_promotion. Twitter/X supports engagement, traffic, awareness, video_views, app_promotion. Pinterest and Google Ads support only engagement, traffic, awareness, video_views. | 
**Budget** | [**UpdateAdCampaignRequestBudget**](UpdateAdCampaignRequestBudget.md) |  | 
**Currency** | Pointer to **string** |  | [optional] 
**Schedule** | Pointer to [**BoostPostRequestSchedule**](BoostPostRequestSchedule.md) |  | [optional] 
**Targeting** | Pointer to [**UpdateAdRequestTargeting**](UpdateAdRequestTargeting.md) |  | [optional] 
**BidStrategy** | Pointer to [**BidStrategy**](BidStrategy.md) | Meta bid strategy applied to the ad set. On TikTok, mapped to &#x60;bid_type&#x60; / &#x60;bid_price&#x60; / &#x60;deep_bid_type&#x60; automatically.  | [optional] 
**BidAmount** | Pointer to **float32** | Bid cap in WHOLE currency units (USD: 5 &#x3D; $5.00; JPY: 100 &#x3D; ¥100). Required when &#x60;bidStrategy&#x60; is &#x60;LOWEST_COST_WITH_BID_CAP&#x60; or &#x60;COST_CAP&#x60;. Backward-compat: providing &#x60;bidAmount&#x60; without &#x60;bidStrategy&#x60; is treated as &#x60;LOWEST_COST_WITH_BID_CAP&#x60;.  | [optional] 
**RoasAverageFloor** | Pointer to **float32** | Minimum ROAS as a decimal multiplier (e.g. 2.0 &#x3D; 2.0x ROAS). Required when &#x60;bidStrategy&#x60; is &#x60;LOWEST_COST_WITH_MIN_ROAS&#x60;. Sent to Meta as &#x60;bid_constraints.roas_average_floor&#x60; × 10000 (Meta uses fixed-point integers).  | [optional] 
**Tracking** | Pointer to [**BoostPostRequestTracking**](BoostPostRequestTracking.md) |  | [optional] 
**SpecialAdCategories** | Pointer to **[]string** | Meta only. Required for housing, employment, credit, or political ads. | [optional] 
**LinkUrl** | Pointer to **string** | TikTok-only. Custom destination URL for the Spark Ad. Without this, TikTok Spark Ads have no clickable destination — required for traffic / conversion objectives. Maps to &#x60;landing_page_url&#x60; on the creative entry of /v2/ad/create/ (TikTok SDK &#x60;AdcreateCreatives.landing_page_url&#x60;). Ignored on Meta / LinkedIn / Pinterest / X / Google (those infer the destination from the boosted post).  | [optional] 
**CallToAction** | Pointer to **string** | TikTok-only. Call-to-action button label on the Spark Ad creative (e.g. &#x60;LEARN_MORE&#x60;, &#x60;SHOP_NOW&#x60;, &#x60;DOWNLOAD_NOW&#x60;, &#x60;SIGN_UP&#x60;, &#x60;WATCH_NOW&#x60;). Maps to &#x60;call_to_action&#x60; on the creative entry of /v2/ad/create/. Pass-through — the platform validates the value. See TikTok&#39;s \&quot;Enumeration - Call-to-Action\&quot; reference for the full list.  | [optional] 
**SparkAuthCode** | Pointer to **string** | TikTok-only. Spark Code (creator&#39;s &#x60;auth_code&#x60;) authorizing cross-creator Spark Ads — the advertiser can boost a video owned by a DIFFERENT TikTok account. Without this, boosts are limited to videos owned by the same account running the ads (same-BC creators only). The creator generates the code in their TikTok app&#39;s Promote settings and shares it with the advertiser. Maps to &#x60;auth_code&#x60; on the creative entry of /v2/ad/create/.  | [optional] 
**DsaBeneficiary** | Pointer to **string** | Name of the legal entity benefiting from the ad. Required by Meta when targeting EU users (DSA Article 26). Not enforced at schema level; enforced server-side when targeting intersects EU member states.  | [optional] 
**DsaPayor** | Pointer to **string** | Name of the legal entity paying for the ad. Required by Meta when targeting EU users (DSA Article 26). Note Meta API spelling: dsa_payor (not dsa_payer).  | [optional] 

## Methods

### NewBoostPostRequest

`func NewBoostPostRequest(accountId string, adAccountId string, name string, goal string, budget UpdateAdCampaignRequestBudget, ) *BoostPostRequest`

NewBoostPostRequest instantiates a new BoostPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoostPostRequestWithDefaults

`func NewBoostPostRequestWithDefaults() *BoostPostRequest`

NewBoostPostRequestWithDefaults instantiates a new BoostPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPostId

`func (o *BoostPostRequest) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *BoostPostRequest) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *BoostPostRequest) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *BoostPostRequest) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *BoostPostRequest) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *BoostPostRequest) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *BoostPostRequest) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *BoostPostRequest) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetAccountId

`func (o *BoostPostRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BoostPostRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BoostPostRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAdAccountId

`func (o *BoostPostRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *BoostPostRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *BoostPostRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *BoostPostRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BoostPostRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BoostPostRequest) SetName(v string)`

SetName sets Name field to given value.


### GetGoal

`func (o *BoostPostRequest) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *BoostPostRequest) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *BoostPostRequest) SetGoal(v string)`

SetGoal sets Goal field to given value.


### GetBudget

`func (o *BoostPostRequest) GetBudget() UpdateAdCampaignRequestBudget`

GetBudget returns the Budget field if non-nil, zero value otherwise.

### GetBudgetOk

`func (o *BoostPostRequest) GetBudgetOk() (*UpdateAdCampaignRequestBudget, bool)`

GetBudgetOk returns a tuple with the Budget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudget

`func (o *BoostPostRequest) SetBudget(v UpdateAdCampaignRequestBudget)`

SetBudget sets Budget field to given value.


### GetCurrency

`func (o *BoostPostRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BoostPostRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BoostPostRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BoostPostRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSchedule

`func (o *BoostPostRequest) GetSchedule() BoostPostRequestSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *BoostPostRequest) GetScheduleOk() (*BoostPostRequestSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *BoostPostRequest) SetSchedule(v BoostPostRequestSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *BoostPostRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetTargeting

`func (o *BoostPostRequest) GetTargeting() UpdateAdRequestTargeting`

GetTargeting returns the Targeting field if non-nil, zero value otherwise.

### GetTargetingOk

`func (o *BoostPostRequest) GetTargetingOk() (*UpdateAdRequestTargeting, bool)`

GetTargetingOk returns a tuple with the Targeting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargeting

`func (o *BoostPostRequest) SetTargeting(v UpdateAdRequestTargeting)`

SetTargeting sets Targeting field to given value.

### HasTargeting

`func (o *BoostPostRequest) HasTargeting() bool`

HasTargeting returns a boolean if a field has been set.

### GetBidStrategy

`func (o *BoostPostRequest) GetBidStrategy() BidStrategy`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *BoostPostRequest) GetBidStrategyOk() (*BidStrategy, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *BoostPostRequest) SetBidStrategy(v BidStrategy)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *BoostPostRequest) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### GetBidAmount

`func (o *BoostPostRequest) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *BoostPostRequest) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *BoostPostRequest) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *BoostPostRequest) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### GetRoasAverageFloor

`func (o *BoostPostRequest) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *BoostPostRequest) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *BoostPostRequest) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *BoostPostRequest) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### GetTracking

`func (o *BoostPostRequest) GetTracking() BoostPostRequestTracking`

GetTracking returns the Tracking field if non-nil, zero value otherwise.

### GetTrackingOk

`func (o *BoostPostRequest) GetTrackingOk() (*BoostPostRequestTracking, bool)`

GetTrackingOk returns a tuple with the Tracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTracking

`func (o *BoostPostRequest) SetTracking(v BoostPostRequestTracking)`

SetTracking sets Tracking field to given value.

### HasTracking

`func (o *BoostPostRequest) HasTracking() bool`

HasTracking returns a boolean if a field has been set.

### GetSpecialAdCategories

`func (o *BoostPostRequest) GetSpecialAdCategories() []string`

GetSpecialAdCategories returns the SpecialAdCategories field if non-nil, zero value otherwise.

### GetSpecialAdCategoriesOk

`func (o *BoostPostRequest) GetSpecialAdCategoriesOk() (*[]string, bool)`

GetSpecialAdCategoriesOk returns a tuple with the SpecialAdCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialAdCategories

`func (o *BoostPostRequest) SetSpecialAdCategories(v []string)`

SetSpecialAdCategories sets SpecialAdCategories field to given value.

### HasSpecialAdCategories

`func (o *BoostPostRequest) HasSpecialAdCategories() bool`

HasSpecialAdCategories returns a boolean if a field has been set.

### GetLinkUrl

`func (o *BoostPostRequest) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *BoostPostRequest) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *BoostPostRequest) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *BoostPostRequest) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetCallToAction

`func (o *BoostPostRequest) GetCallToAction() string`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *BoostPostRequest) GetCallToActionOk() (*string, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *BoostPostRequest) SetCallToAction(v string)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *BoostPostRequest) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetSparkAuthCode

`func (o *BoostPostRequest) GetSparkAuthCode() string`

GetSparkAuthCode returns the SparkAuthCode field if non-nil, zero value otherwise.

### GetSparkAuthCodeOk

`func (o *BoostPostRequest) GetSparkAuthCodeOk() (*string, bool)`

GetSparkAuthCodeOk returns a tuple with the SparkAuthCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSparkAuthCode

`func (o *BoostPostRequest) SetSparkAuthCode(v string)`

SetSparkAuthCode sets SparkAuthCode field to given value.

### HasSparkAuthCode

`func (o *BoostPostRequest) HasSparkAuthCode() bool`

HasSparkAuthCode returns a boolean if a field has been set.

### GetDsaBeneficiary

`func (o *BoostPostRequest) GetDsaBeneficiary() string`

GetDsaBeneficiary returns the DsaBeneficiary field if non-nil, zero value otherwise.

### GetDsaBeneficiaryOk

`func (o *BoostPostRequest) GetDsaBeneficiaryOk() (*string, bool)`

GetDsaBeneficiaryOk returns a tuple with the DsaBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaBeneficiary

`func (o *BoostPostRequest) SetDsaBeneficiary(v string)`

SetDsaBeneficiary sets DsaBeneficiary field to given value.

### HasDsaBeneficiary

`func (o *BoostPostRequest) HasDsaBeneficiary() bool`

HasDsaBeneficiary returns a boolean if a field has been set.

### GetDsaPayor

`func (o *BoostPostRequest) GetDsaPayor() string`

GetDsaPayor returns the DsaPayor field if non-nil, zero value otherwise.

### GetDsaPayorOk

`func (o *BoostPostRequest) GetDsaPayorOk() (*string, bool)`

GetDsaPayorOk returns a tuple with the DsaPayor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaPayor

`func (o *BoostPostRequest) SetDsaPayor(v string)`

SetDsaPayor sets DsaPayor field to given value.

### HasDsaPayor

`func (o *BoostPostRequest) HasDsaPayor() bool`

HasDsaPayor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


