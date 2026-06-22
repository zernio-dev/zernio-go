# CreateCtwaAdRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Facebook or Instagram SocialAccount ID. | 
**AdAccountId** | **string** | Meta ad account ID, e.g. &#x60;act_123456789&#x60;. | 
**Name** | **string** | Ad display name. Used to derive campaign / ad set names. On the multi-creative shape, each ad&#39;s Meta name gets a \&quot; #N\&quot; suffix (1-indexed) so Ads Manager shows them as a numbered batch.  | 
**Headline** | Pointer to **string** | Single-creative shape only. Mutually exclusive with &#x60;creatives[]&#x60;.  | [optional] 
**Body** | Pointer to **string** | Primary text shown above the image / video. Single-creative shape only. Mutually exclusive with &#x60;creatives[]&#x60;.  | [optional] 
**ImageUrl** | Pointer to **string** | Image asset for single-creative shape. Mutually exclusive with &#x60;video&#x60; and with &#x60;creatives[]&#x60;. Required on the single-creative shape if &#x60;video&#x60; is not supplied.  | [optional] 
**Video** | Pointer to [**CreateCtwaAdRequestVideo**](CreateCtwaAdRequestVideo.md) |  | [optional] 
**Creatives** | Pointer to [**[]CreateCtwaAdRequestCreativesInner**](CreateCtwaAdRequestCreativesInner.md) | Multi-creative shape: N CTWA ads under one campaign + one ad set, sharing budget and targeting. Mutually exclusive with the top-level single-creative fields (&#x60;headline&#x60; / &#x60;body&#x60; / &#x60;imageUrl&#x60; / &#x60;video&#x60;). Each entry must supply its own headline, body, and exactly one of &#x60;imageUrl&#x60; / &#x60;video&#x60;.  | [optional] 
**BudgetAmount** | **float32** | Budget amount in the ad account&#39;s currency major units (e.g. dollars for USD, not cents). Must be &gt; 0.  | 
**BudgetType** | **string** |  | 
**Currency** | Pointer to **string** | ISO 4217 currency code matching the ad account&#39;s currency (e.g. &#x60;USD&#x60;). Optional; Meta infers from the ad account when omitted.  | [optional] 
**EndDate** | Pointer to **time.Time** | ISO 8601 datetime. Required when &#x60;budgetType&#x60; is &#x60;lifetime&#x60;.  | [optional] 
**Countries** | Pointer to **[]string** | ISO 3166-1 alpha-2 country codes. Defaults to &#x60;[\&quot;US\&quot;]&#x60; only when no other geo (&#x60;cities&#x60;, &#x60;regions&#x60;, &#x60;zips&#x60;, &#x60;metros&#x60;, &#x60;customLocations&#x60;) is supplied.  | [optional] 
**Cities** | Pointer to [**[]CreateCtwaAdRequestCitiesInner**](CreateCtwaAdRequestCitiesInner.md) | City-level geo targeting for local CTWA campaigns (e.g. 25km radius around Milan). Each entry maps to Meta&#39;s TargetingGeoLocationCity. &#x60;key&#x60; is Meta&#39;s city ID (lookupable via GET /v1/ads/targeting/search). &#x60;radius&#x60; and &#x60;distance_unit&#x60; are coupled: set both or neither.  | [optional] 
**Regions** | Pointer to [**[]CreateCtwaAdRequestRegionsInner**](CreateCtwaAdRequestRegionsInner.md) | Region / state-level geo targeting. &#x60;key&#x60; is Meta&#39;s region ID (lookupable via GET /v1/ads/targeting/search?type&#x3D;region).  | [optional] 
**Zips** | Pointer to [**[]CreateCtwaAdRequestZipsInner**](CreateCtwaAdRequestZipsInner.md) | ZIP / postal-code geo targeting. &#x60;key&#x60; is the platform&#39;s postal id resolved via /v1/ads/targeting/search.  | [optional] 
**Metros** | Pointer to [**[]CreateCtwaAdRequestZipsInner**](CreateCtwaAdRequestZipsInner.md) | DMA / metro-area geo targeting. &#x60;key&#x60; is Meta&#39;s metro id (e.g. &#x60;DMA:807&#x60;).  | [optional] 
**CustomLocations** | Pointer to [**[]CreateCtwaAdRequestCustomLocationsInner**](CreateCtwaAdRequestCustomLocationsInner.md) | Point-radius geo (Meta &#x60;geo_locations.custom_locations&#x60;). Use for targeting a radius around a specific lat/long when no Meta city/region key fits. &#x60;distanceUnit&#x60; is required.  | [optional] 
**AgeMin** | Pointer to **int32** |  | [optional] 
**AgeMax** | Pointer to **int32** |  | [optional] 
**Interests** | Pointer to [**[]CreateStandaloneAdRequestBehaviorsInner**](CreateStandaloneAdRequestBehaviorsInner.md) |  | [optional] 
**AudienceId** | Pointer to **string** | Custom audience ID to target. | [optional] 
**AdvantageAudience** | Pointer to **int32** | Meta&#39;s Advantage+ audience expansion. &#x60;0&#x60; (default) keeps targeting strict; &#x60;1&#x60; lets Meta expand beyond the supplied targeting when its delivery system finds better matches. Always sent on CREATE (Meta requires it).  | [optional] 
**Objective** | Pointer to **string** | Defaults to &#x60;OUTCOME_ENGAGEMENT&#x60; (the broadly-supported CTWA objective). &#x60;OUTCOME_SALES&#x60; and &#x60;OUTCOME_LEADS&#x60; require additional account configuration (Dataset linked to the WABA for sales) and may be rejected by Meta if missing.  | [optional] 
**BidStrategy** | Pointer to **string** | Meta bid strategy applied to the shared ad set. Defaults to &#x60;LOWEST_COST_WITHOUT_CAP&#x60; (auto-bid) when omitted. &#x60;LOWEST_COST_WITH_BID_CAP&#x60; and &#x60;COST_CAP&#x60; require &#x60;bidAmount&#x60;. &#x60;LOWEST_COST_WITH_MIN_ROAS&#x60; requires &#x60;roasAverageFloor&#x60;. CTWA&#39;s &#x60;optimization_goal&#x60; is fixed to &#x60;CONVERSATIONS&#x60;, but the bid strategy is independent.  | [optional] 
**BidAmount** | Pointer to **float32** | Whole currency units (e.g. &#x60;5&#x60; &#x3D; $5.00 on a USD account). Required when &#x60;bidStrategy&#x60; is &#x60;LOWEST_COST_WITH_BID_CAP&#x60; or &#x60;COST_CAP&#x60;; rejected otherwise.  | [optional] 
**RoasAverageFloor** | Pointer to **float32** | Decimal ROAS multiplier (e.g. &#x60;2.0&#x60; &#x3D; 2.0× ROAS floor). Required when &#x60;bidStrategy&#x60; is &#x60;LOWEST_COST_WITH_MIN_ROAS&#x60;; rejected otherwise. Meta enforces its own upper bound server-side.  | [optional] 
**DsaBeneficiary** | Pointer to **string** | Name of the legal entity benefiting from the ad. Required by Meta when targeting EU users (DSA Article 26). Not enforced at schema level; enforced server-side when targeting intersects EU member states.  | [optional] 
**DsaPayor** | Pointer to **string** | Name of the legal entity paying for the ad. Required by Meta when targeting EU users (DSA Article 26). Note Meta API spelling: dsa_payor (not dsa_payer).  | [optional] 

## Methods

### NewCreateCtwaAdRequest

`func NewCreateCtwaAdRequest(accountId string, adAccountId string, name string, budgetAmount float32, budgetType string, ) *CreateCtwaAdRequest`

NewCreateCtwaAdRequest instantiates a new CreateCtwaAdRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCtwaAdRequestWithDefaults

`func NewCreateCtwaAdRequestWithDefaults() *CreateCtwaAdRequest`

NewCreateCtwaAdRequestWithDefaults instantiates a new CreateCtwaAdRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateCtwaAdRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateCtwaAdRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateCtwaAdRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAdAccountId

`func (o *CreateCtwaAdRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *CreateCtwaAdRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *CreateCtwaAdRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *CreateCtwaAdRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCtwaAdRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCtwaAdRequest) SetName(v string)`

SetName sets Name field to given value.


### GetHeadline

`func (o *CreateCtwaAdRequest) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *CreateCtwaAdRequest) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *CreateCtwaAdRequest) SetHeadline(v string)`

SetHeadline sets Headline field to given value.

### HasHeadline

`func (o *CreateCtwaAdRequest) HasHeadline() bool`

HasHeadline returns a boolean if a field has been set.

### GetBody

`func (o *CreateCtwaAdRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *CreateCtwaAdRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *CreateCtwaAdRequest) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *CreateCtwaAdRequest) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetImageUrl

`func (o *CreateCtwaAdRequest) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateCtwaAdRequest) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateCtwaAdRequest) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateCtwaAdRequest) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideo

`func (o *CreateCtwaAdRequest) GetVideo() CreateCtwaAdRequestVideo`

GetVideo returns the Video field if non-nil, zero value otherwise.

### GetVideoOk

`func (o *CreateCtwaAdRequest) GetVideoOk() (*CreateCtwaAdRequestVideo, bool)`

GetVideoOk returns a tuple with the Video field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideo

`func (o *CreateCtwaAdRequest) SetVideo(v CreateCtwaAdRequestVideo)`

SetVideo sets Video field to given value.

### HasVideo

`func (o *CreateCtwaAdRequest) HasVideo() bool`

HasVideo returns a boolean if a field has been set.

### GetCreatives

`func (o *CreateCtwaAdRequest) GetCreatives() []CreateCtwaAdRequestCreativesInner`

GetCreatives returns the Creatives field if non-nil, zero value otherwise.

### GetCreativesOk

`func (o *CreateCtwaAdRequest) GetCreativesOk() (*[]CreateCtwaAdRequestCreativesInner, bool)`

GetCreativesOk returns a tuple with the Creatives field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatives

`func (o *CreateCtwaAdRequest) SetCreatives(v []CreateCtwaAdRequestCreativesInner)`

SetCreatives sets Creatives field to given value.

### HasCreatives

`func (o *CreateCtwaAdRequest) HasCreatives() bool`

HasCreatives returns a boolean if a field has been set.

### GetBudgetAmount

`func (o *CreateCtwaAdRequest) GetBudgetAmount() float32`

GetBudgetAmount returns the BudgetAmount field if non-nil, zero value otherwise.

### GetBudgetAmountOk

`func (o *CreateCtwaAdRequest) GetBudgetAmountOk() (*float32, bool)`

GetBudgetAmountOk returns a tuple with the BudgetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetAmount

`func (o *CreateCtwaAdRequest) SetBudgetAmount(v float32)`

SetBudgetAmount sets BudgetAmount field to given value.


### GetBudgetType

`func (o *CreateCtwaAdRequest) GetBudgetType() string`

GetBudgetType returns the BudgetType field if non-nil, zero value otherwise.

### GetBudgetTypeOk

`func (o *CreateCtwaAdRequest) GetBudgetTypeOk() (*string, bool)`

GetBudgetTypeOk returns a tuple with the BudgetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetType

`func (o *CreateCtwaAdRequest) SetBudgetType(v string)`

SetBudgetType sets BudgetType field to given value.


### GetCurrency

`func (o *CreateCtwaAdRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateCtwaAdRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateCtwaAdRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateCtwaAdRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetEndDate

`func (o *CreateCtwaAdRequest) GetEndDate() time.Time`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *CreateCtwaAdRequest) GetEndDateOk() (*time.Time, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *CreateCtwaAdRequest) SetEndDate(v time.Time)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *CreateCtwaAdRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetCountries

`func (o *CreateCtwaAdRequest) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *CreateCtwaAdRequest) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *CreateCtwaAdRequest) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *CreateCtwaAdRequest) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetCities

`func (o *CreateCtwaAdRequest) GetCities() []CreateCtwaAdRequestCitiesInner`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *CreateCtwaAdRequest) GetCitiesOk() (*[]CreateCtwaAdRequestCitiesInner, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *CreateCtwaAdRequest) SetCities(v []CreateCtwaAdRequestCitiesInner)`

SetCities sets Cities field to given value.

### HasCities

`func (o *CreateCtwaAdRequest) HasCities() bool`

HasCities returns a boolean if a field has been set.

### GetRegions

`func (o *CreateCtwaAdRequest) GetRegions() []CreateCtwaAdRequestRegionsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *CreateCtwaAdRequest) GetRegionsOk() (*[]CreateCtwaAdRequestRegionsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *CreateCtwaAdRequest) SetRegions(v []CreateCtwaAdRequestRegionsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *CreateCtwaAdRequest) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetZips

`func (o *CreateCtwaAdRequest) GetZips() []CreateCtwaAdRequestZipsInner`

GetZips returns the Zips field if non-nil, zero value otherwise.

### GetZipsOk

`func (o *CreateCtwaAdRequest) GetZipsOk() (*[]CreateCtwaAdRequestZipsInner, bool)`

GetZipsOk returns a tuple with the Zips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZips

`func (o *CreateCtwaAdRequest) SetZips(v []CreateCtwaAdRequestZipsInner)`

SetZips sets Zips field to given value.

### HasZips

`func (o *CreateCtwaAdRequest) HasZips() bool`

HasZips returns a boolean if a field has been set.

### GetMetros

`func (o *CreateCtwaAdRequest) GetMetros() []CreateCtwaAdRequestZipsInner`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *CreateCtwaAdRequest) GetMetrosOk() (*[]CreateCtwaAdRequestZipsInner, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *CreateCtwaAdRequest) SetMetros(v []CreateCtwaAdRequestZipsInner)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *CreateCtwaAdRequest) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetCustomLocations

`func (o *CreateCtwaAdRequest) GetCustomLocations() []CreateCtwaAdRequestCustomLocationsInner`

GetCustomLocations returns the CustomLocations field if non-nil, zero value otherwise.

### GetCustomLocationsOk

`func (o *CreateCtwaAdRequest) GetCustomLocationsOk() (*[]CreateCtwaAdRequestCustomLocationsInner, bool)`

GetCustomLocationsOk returns a tuple with the CustomLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLocations

`func (o *CreateCtwaAdRequest) SetCustomLocations(v []CreateCtwaAdRequestCustomLocationsInner)`

SetCustomLocations sets CustomLocations field to given value.

### HasCustomLocations

`func (o *CreateCtwaAdRequest) HasCustomLocations() bool`

HasCustomLocations returns a boolean if a field has been set.

### GetAgeMin

`func (o *CreateCtwaAdRequest) GetAgeMin() int32`

GetAgeMin returns the AgeMin field if non-nil, zero value otherwise.

### GetAgeMinOk

`func (o *CreateCtwaAdRequest) GetAgeMinOk() (*int32, bool)`

GetAgeMinOk returns a tuple with the AgeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMin

`func (o *CreateCtwaAdRequest) SetAgeMin(v int32)`

SetAgeMin sets AgeMin field to given value.

### HasAgeMin

`func (o *CreateCtwaAdRequest) HasAgeMin() bool`

HasAgeMin returns a boolean if a field has been set.

### GetAgeMax

`func (o *CreateCtwaAdRequest) GetAgeMax() int32`

GetAgeMax returns the AgeMax field if non-nil, zero value otherwise.

### GetAgeMaxOk

`func (o *CreateCtwaAdRequest) GetAgeMaxOk() (*int32, bool)`

GetAgeMaxOk returns a tuple with the AgeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMax

`func (o *CreateCtwaAdRequest) SetAgeMax(v int32)`

SetAgeMax sets AgeMax field to given value.

### HasAgeMax

`func (o *CreateCtwaAdRequest) HasAgeMax() bool`

HasAgeMax returns a boolean if a field has been set.

### GetInterests

`func (o *CreateCtwaAdRequest) GetInterests() []CreateStandaloneAdRequestBehaviorsInner`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *CreateCtwaAdRequest) GetInterestsOk() (*[]CreateStandaloneAdRequestBehaviorsInner, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *CreateCtwaAdRequest) SetInterests(v []CreateStandaloneAdRequestBehaviorsInner)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *CreateCtwaAdRequest) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetAudienceId

`func (o *CreateCtwaAdRequest) GetAudienceId() string`

GetAudienceId returns the AudienceId field if non-nil, zero value otherwise.

### GetAudienceIdOk

`func (o *CreateCtwaAdRequest) GetAudienceIdOk() (*string, bool)`

GetAudienceIdOk returns a tuple with the AudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceId

`func (o *CreateCtwaAdRequest) SetAudienceId(v string)`

SetAudienceId sets AudienceId field to given value.

### HasAudienceId

`func (o *CreateCtwaAdRequest) HasAudienceId() bool`

HasAudienceId returns a boolean if a field has been set.

### GetAdvantageAudience

`func (o *CreateCtwaAdRequest) GetAdvantageAudience() int32`

GetAdvantageAudience returns the AdvantageAudience field if non-nil, zero value otherwise.

### GetAdvantageAudienceOk

`func (o *CreateCtwaAdRequest) GetAdvantageAudienceOk() (*int32, bool)`

GetAdvantageAudienceOk returns a tuple with the AdvantageAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantageAudience

`func (o *CreateCtwaAdRequest) SetAdvantageAudience(v int32)`

SetAdvantageAudience sets AdvantageAudience field to given value.

### HasAdvantageAudience

`func (o *CreateCtwaAdRequest) HasAdvantageAudience() bool`

HasAdvantageAudience returns a boolean if a field has been set.

### GetObjective

`func (o *CreateCtwaAdRequest) GetObjective() string`

GetObjective returns the Objective field if non-nil, zero value otherwise.

### GetObjectiveOk

`func (o *CreateCtwaAdRequest) GetObjectiveOk() (*string, bool)`

GetObjectiveOk returns a tuple with the Objective field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjective

`func (o *CreateCtwaAdRequest) SetObjective(v string)`

SetObjective sets Objective field to given value.

### HasObjective

`func (o *CreateCtwaAdRequest) HasObjective() bool`

HasObjective returns a boolean if a field has been set.

### GetBidStrategy

`func (o *CreateCtwaAdRequest) GetBidStrategy() string`

GetBidStrategy returns the BidStrategy field if non-nil, zero value otherwise.

### GetBidStrategyOk

`func (o *CreateCtwaAdRequest) GetBidStrategyOk() (*string, bool)`

GetBidStrategyOk returns a tuple with the BidStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidStrategy

`func (o *CreateCtwaAdRequest) SetBidStrategy(v string)`

SetBidStrategy sets BidStrategy field to given value.

### HasBidStrategy

`func (o *CreateCtwaAdRequest) HasBidStrategy() bool`

HasBidStrategy returns a boolean if a field has been set.

### GetBidAmount

`func (o *CreateCtwaAdRequest) GetBidAmount() float32`

GetBidAmount returns the BidAmount field if non-nil, zero value otherwise.

### GetBidAmountOk

`func (o *CreateCtwaAdRequest) GetBidAmountOk() (*float32, bool)`

GetBidAmountOk returns a tuple with the BidAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidAmount

`func (o *CreateCtwaAdRequest) SetBidAmount(v float32)`

SetBidAmount sets BidAmount field to given value.

### HasBidAmount

`func (o *CreateCtwaAdRequest) HasBidAmount() bool`

HasBidAmount returns a boolean if a field has been set.

### GetRoasAverageFloor

`func (o *CreateCtwaAdRequest) GetRoasAverageFloor() float32`

GetRoasAverageFloor returns the RoasAverageFloor field if non-nil, zero value otherwise.

### GetRoasAverageFloorOk

`func (o *CreateCtwaAdRequest) GetRoasAverageFloorOk() (*float32, bool)`

GetRoasAverageFloorOk returns a tuple with the RoasAverageFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoasAverageFloor

`func (o *CreateCtwaAdRequest) SetRoasAverageFloor(v float32)`

SetRoasAverageFloor sets RoasAverageFloor field to given value.

### HasRoasAverageFloor

`func (o *CreateCtwaAdRequest) HasRoasAverageFloor() bool`

HasRoasAverageFloor returns a boolean if a field has been set.

### GetDsaBeneficiary

`func (o *CreateCtwaAdRequest) GetDsaBeneficiary() string`

GetDsaBeneficiary returns the DsaBeneficiary field if non-nil, zero value otherwise.

### GetDsaBeneficiaryOk

`func (o *CreateCtwaAdRequest) GetDsaBeneficiaryOk() (*string, bool)`

GetDsaBeneficiaryOk returns a tuple with the DsaBeneficiary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaBeneficiary

`func (o *CreateCtwaAdRequest) SetDsaBeneficiary(v string)`

SetDsaBeneficiary sets DsaBeneficiary field to given value.

### HasDsaBeneficiary

`func (o *CreateCtwaAdRequest) HasDsaBeneficiary() bool`

HasDsaBeneficiary returns a boolean if a field has been set.

### GetDsaPayor

`func (o *CreateCtwaAdRequest) GetDsaPayor() string`

GetDsaPayor returns the DsaPayor field if non-nil, zero value otherwise.

### GetDsaPayorOk

`func (o *CreateCtwaAdRequest) GetDsaPayorOk() (*string, bool)`

GetDsaPayorOk returns a tuple with the DsaPayor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDsaPayor

`func (o *CreateCtwaAdRequest) SetDsaPayor(v string)`

SetDsaPayor sets DsaPayor field to given value.

### HasDsaPayor

`func (o *CreateCtwaAdRequest) HasDsaPayor() bool`

HasDsaPayor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


