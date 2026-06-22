# TargetingSpec

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Countries** | Pointer to **[]string** | ISO 3166-1 alpha-2 country codes (e.g. [&#39;US&#39;]). | [optional] 
**Regions** | Pointer to [**[]CreateStandaloneAdRequestZipsInner**](CreateStandaloneAdRequestZipsInner.md) | Region/state targeting. &#x60;key&#x60; is the platform location ID from /v1/ads/targeting/search?dimension&#x3D;geo&amp;geoType&#x3D;region. | [optional] 
**Cities** | Pointer to [**[]TargetingSpecCitiesInner**](TargetingSpecCitiesInner.md) | City targeting. Optional &#x60;radius&#x60; + &#x60;distanceUnit&#x60; extend beyond the city limits; both must be set together or both omitted. &#x60;radius&#x60; is only honoured on platforms whose capability map allows city radius (Meta). | [optional] 
**Zips** | Pointer to [**[]CreateStandaloneAdRequestZipsInner**](CreateStandaloneAdRequestZipsInner.md) | Postal/ZIP targeting. &#x60;key&#x60; is the platform&#39;s postal location ID (e.g. Meta &#x60;US:94304&#x60;). Supported on Meta, Google, TikTok, Pinterest, X. | [optional] 
**Metros** | Pointer to [**[]CreateStandaloneAdRequestZipsInner**](CreateStandaloneAdRequestZipsInner.md) | DMA / metro-area targeting. &#x60;key&#x60; is the platform&#39;s metro ID (e.g. Meta &#x60;DMA:807&#x60;). | [optional] 
**CustomLocations** | Pointer to [**[]CreateStandaloneAdRequestCustomLocationsInner**](CreateStandaloneAdRequestCustomLocationsInner.md) | Point-radius (lat/lng) targeting (Meta custom_locations / Google proximity). Honoured only where the capability map allows radius (Meta). | [optional] 
**ExcludedLocations** | Pointer to [**TargetingSpecExcludedLocations**](TargetingSpecExcludedLocations.md) |  | [optional] 
**AgeMin** | Pointer to **int32** |  | [optional] 
**AgeMax** | Pointer to **int32** |  | [optional] 
**Gender** | Pointer to **string** | Restrict by gender. &#39;all&#39; (default) targets everyone. | [optional] 
**IncomeTier** | Pointer to **string** | Normalized household-income tier (ZIP/percentile based). Meta and TikTok express all four. Google maps only &#x60;top_10&#x60; (its INCOME_RANGE_90_UP); other tiers on Google, and any income tier on LinkedIn / X / Pinterest, are rejected. On Meta, income/zip targeting requires the relevant &#x60;specialAdCategories&#x60; to be unset (housing/employment/credit ads cannot use it).  | [optional] 
**Languages** | Pointer to **[]string** | Language codes (e.g. [&#39;en&#39;]). | [optional] 
**Interests** | Pointer to [**[]CreateStandaloneAdRequestBehaviorsInner**](CreateStandaloneAdRequestBehaviorsInner.md) | Interest entities from /v1/ads/targeting/search?dimension&#x3D;interest. Each carries the platform&#39;s opaque id. | [optional] 
**Behaviors** | Pointer to [**[]CreateStandaloneAdRequestBehaviorsInner**](CreateStandaloneAdRequestBehaviorsInner.md) | Behaviour entities from /v1/ads/targeting/search?dimension&#x3D;behavior. Supported on Meta and TikTok. | [optional] 
**Industries** | Pointer to **[]string** | LinkedIn B2B only. Industry URN id fragments. | [optional] 
**CompanySizes** | Pointer to **[]string** | LinkedIn B2B only. | [optional] 
**Seniorities** | Pointer to **[]string** | LinkedIn B2B only. | [optional] 
**JobFunctions** | Pointer to **[]string** | LinkedIn B2B only. | [optional] 
**AudienceInclude** | Pointer to **[]string** | Platform audience IDs to include. | [optional] 
**AudienceExclude** | Pointer to **[]string** | Platform audience IDs to exclude. | [optional] 

## Methods

### NewTargetingSpec

`func NewTargetingSpec() *TargetingSpec`

NewTargetingSpec instantiates a new TargetingSpec object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetingSpecWithDefaults

`func NewTargetingSpecWithDefaults() *TargetingSpec`

NewTargetingSpecWithDefaults instantiates a new TargetingSpec object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountries

`func (o *TargetingSpec) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *TargetingSpec) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *TargetingSpec) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *TargetingSpec) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetRegions

`func (o *TargetingSpec) GetRegions() []CreateStandaloneAdRequestZipsInner`

GetRegions returns the Regions field if non-nil, zero value otherwise.

### GetRegionsOk

`func (o *TargetingSpec) GetRegionsOk() (*[]CreateStandaloneAdRequestZipsInner, bool)`

GetRegionsOk returns a tuple with the Regions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegions

`func (o *TargetingSpec) SetRegions(v []CreateStandaloneAdRequestZipsInner)`

SetRegions sets Regions field to given value.

### HasRegions

`func (o *TargetingSpec) HasRegions() bool`

HasRegions returns a boolean if a field has been set.

### GetCities

`func (o *TargetingSpec) GetCities() []TargetingSpecCitiesInner`

GetCities returns the Cities field if non-nil, zero value otherwise.

### GetCitiesOk

`func (o *TargetingSpec) GetCitiesOk() (*[]TargetingSpecCitiesInner, bool)`

GetCitiesOk returns a tuple with the Cities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCities

`func (o *TargetingSpec) SetCities(v []TargetingSpecCitiesInner)`

SetCities sets Cities field to given value.

### HasCities

`func (o *TargetingSpec) HasCities() bool`

HasCities returns a boolean if a field has been set.

### GetZips

`func (o *TargetingSpec) GetZips() []CreateStandaloneAdRequestZipsInner`

GetZips returns the Zips field if non-nil, zero value otherwise.

### GetZipsOk

`func (o *TargetingSpec) GetZipsOk() (*[]CreateStandaloneAdRequestZipsInner, bool)`

GetZipsOk returns a tuple with the Zips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZips

`func (o *TargetingSpec) SetZips(v []CreateStandaloneAdRequestZipsInner)`

SetZips sets Zips field to given value.

### HasZips

`func (o *TargetingSpec) HasZips() bool`

HasZips returns a boolean if a field has been set.

### GetMetros

`func (o *TargetingSpec) GetMetros() []CreateStandaloneAdRequestZipsInner`

GetMetros returns the Metros field if non-nil, zero value otherwise.

### GetMetrosOk

`func (o *TargetingSpec) GetMetrosOk() (*[]CreateStandaloneAdRequestZipsInner, bool)`

GetMetrosOk returns a tuple with the Metros field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetros

`func (o *TargetingSpec) SetMetros(v []CreateStandaloneAdRequestZipsInner)`

SetMetros sets Metros field to given value.

### HasMetros

`func (o *TargetingSpec) HasMetros() bool`

HasMetros returns a boolean if a field has been set.

### GetCustomLocations

`func (o *TargetingSpec) GetCustomLocations() []CreateStandaloneAdRequestCustomLocationsInner`

GetCustomLocations returns the CustomLocations field if non-nil, zero value otherwise.

### GetCustomLocationsOk

`func (o *TargetingSpec) GetCustomLocationsOk() (*[]CreateStandaloneAdRequestCustomLocationsInner, bool)`

GetCustomLocationsOk returns a tuple with the CustomLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLocations

`func (o *TargetingSpec) SetCustomLocations(v []CreateStandaloneAdRequestCustomLocationsInner)`

SetCustomLocations sets CustomLocations field to given value.

### HasCustomLocations

`func (o *TargetingSpec) HasCustomLocations() bool`

HasCustomLocations returns a boolean if a field has been set.

### GetExcludedLocations

`func (o *TargetingSpec) GetExcludedLocations() TargetingSpecExcludedLocations`

GetExcludedLocations returns the ExcludedLocations field if non-nil, zero value otherwise.

### GetExcludedLocationsOk

`func (o *TargetingSpec) GetExcludedLocationsOk() (*TargetingSpecExcludedLocations, bool)`

GetExcludedLocationsOk returns a tuple with the ExcludedLocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludedLocations

`func (o *TargetingSpec) SetExcludedLocations(v TargetingSpecExcludedLocations)`

SetExcludedLocations sets ExcludedLocations field to given value.

### HasExcludedLocations

`func (o *TargetingSpec) HasExcludedLocations() bool`

HasExcludedLocations returns a boolean if a field has been set.

### GetAgeMin

`func (o *TargetingSpec) GetAgeMin() int32`

GetAgeMin returns the AgeMin field if non-nil, zero value otherwise.

### GetAgeMinOk

`func (o *TargetingSpec) GetAgeMinOk() (*int32, bool)`

GetAgeMinOk returns a tuple with the AgeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMin

`func (o *TargetingSpec) SetAgeMin(v int32)`

SetAgeMin sets AgeMin field to given value.

### HasAgeMin

`func (o *TargetingSpec) HasAgeMin() bool`

HasAgeMin returns a boolean if a field has been set.

### GetAgeMax

`func (o *TargetingSpec) GetAgeMax() int32`

GetAgeMax returns the AgeMax field if non-nil, zero value otherwise.

### GetAgeMaxOk

`func (o *TargetingSpec) GetAgeMaxOk() (*int32, bool)`

GetAgeMaxOk returns a tuple with the AgeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMax

`func (o *TargetingSpec) SetAgeMax(v int32)`

SetAgeMax sets AgeMax field to given value.

### HasAgeMax

`func (o *TargetingSpec) HasAgeMax() bool`

HasAgeMax returns a boolean if a field has been set.

### GetGender

`func (o *TargetingSpec) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *TargetingSpec) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *TargetingSpec) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *TargetingSpec) HasGender() bool`

HasGender returns a boolean if a field has been set.

### GetIncomeTier

`func (o *TargetingSpec) GetIncomeTier() string`

GetIncomeTier returns the IncomeTier field if non-nil, zero value otherwise.

### GetIncomeTierOk

`func (o *TargetingSpec) GetIncomeTierOk() (*string, bool)`

GetIncomeTierOk returns a tuple with the IncomeTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeTier

`func (o *TargetingSpec) SetIncomeTier(v string)`

SetIncomeTier sets IncomeTier field to given value.

### HasIncomeTier

`func (o *TargetingSpec) HasIncomeTier() bool`

HasIncomeTier returns a boolean if a field has been set.

### GetLanguages

`func (o *TargetingSpec) GetLanguages() []string`

GetLanguages returns the Languages field if non-nil, zero value otherwise.

### GetLanguagesOk

`func (o *TargetingSpec) GetLanguagesOk() (*[]string, bool)`

GetLanguagesOk returns a tuple with the Languages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguages

`func (o *TargetingSpec) SetLanguages(v []string)`

SetLanguages sets Languages field to given value.

### HasLanguages

`func (o *TargetingSpec) HasLanguages() bool`

HasLanguages returns a boolean if a field has been set.

### GetInterests

`func (o *TargetingSpec) GetInterests() []CreateStandaloneAdRequestBehaviorsInner`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *TargetingSpec) GetInterestsOk() (*[]CreateStandaloneAdRequestBehaviorsInner, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *TargetingSpec) SetInterests(v []CreateStandaloneAdRequestBehaviorsInner)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *TargetingSpec) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetBehaviors

`func (o *TargetingSpec) GetBehaviors() []CreateStandaloneAdRequestBehaviorsInner`

GetBehaviors returns the Behaviors field if non-nil, zero value otherwise.

### GetBehaviorsOk

`func (o *TargetingSpec) GetBehaviorsOk() (*[]CreateStandaloneAdRequestBehaviorsInner, bool)`

GetBehaviorsOk returns a tuple with the Behaviors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBehaviors

`func (o *TargetingSpec) SetBehaviors(v []CreateStandaloneAdRequestBehaviorsInner)`

SetBehaviors sets Behaviors field to given value.

### HasBehaviors

`func (o *TargetingSpec) HasBehaviors() bool`

HasBehaviors returns a boolean if a field has been set.

### GetIndustries

`func (o *TargetingSpec) GetIndustries() []string`

GetIndustries returns the Industries field if non-nil, zero value otherwise.

### GetIndustriesOk

`func (o *TargetingSpec) GetIndustriesOk() (*[]string, bool)`

GetIndustriesOk returns a tuple with the Industries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIndustries

`func (o *TargetingSpec) SetIndustries(v []string)`

SetIndustries sets Industries field to given value.

### HasIndustries

`func (o *TargetingSpec) HasIndustries() bool`

HasIndustries returns a boolean if a field has been set.

### GetCompanySizes

`func (o *TargetingSpec) GetCompanySizes() []string`

GetCompanySizes returns the CompanySizes field if non-nil, zero value otherwise.

### GetCompanySizesOk

`func (o *TargetingSpec) GetCompanySizesOk() (*[]string, bool)`

GetCompanySizesOk returns a tuple with the CompanySizes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanySizes

`func (o *TargetingSpec) SetCompanySizes(v []string)`

SetCompanySizes sets CompanySizes field to given value.

### HasCompanySizes

`func (o *TargetingSpec) HasCompanySizes() bool`

HasCompanySizes returns a boolean if a field has been set.

### GetSeniorities

`func (o *TargetingSpec) GetSeniorities() []string`

GetSeniorities returns the Seniorities field if non-nil, zero value otherwise.

### GetSenioritiesOk

`func (o *TargetingSpec) GetSenioritiesOk() (*[]string, bool)`

GetSenioritiesOk returns a tuple with the Seniorities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeniorities

`func (o *TargetingSpec) SetSeniorities(v []string)`

SetSeniorities sets Seniorities field to given value.

### HasSeniorities

`func (o *TargetingSpec) HasSeniorities() bool`

HasSeniorities returns a boolean if a field has been set.

### GetJobFunctions

`func (o *TargetingSpec) GetJobFunctions() []string`

GetJobFunctions returns the JobFunctions field if non-nil, zero value otherwise.

### GetJobFunctionsOk

`func (o *TargetingSpec) GetJobFunctionsOk() (*[]string, bool)`

GetJobFunctionsOk returns a tuple with the JobFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobFunctions

`func (o *TargetingSpec) SetJobFunctions(v []string)`

SetJobFunctions sets JobFunctions field to given value.

### HasJobFunctions

`func (o *TargetingSpec) HasJobFunctions() bool`

HasJobFunctions returns a boolean if a field has been set.

### GetAudienceInclude

`func (o *TargetingSpec) GetAudienceInclude() []string`

GetAudienceInclude returns the AudienceInclude field if non-nil, zero value otherwise.

### GetAudienceIncludeOk

`func (o *TargetingSpec) GetAudienceIncludeOk() (*[]string, bool)`

GetAudienceIncludeOk returns a tuple with the AudienceInclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceInclude

`func (o *TargetingSpec) SetAudienceInclude(v []string)`

SetAudienceInclude sets AudienceInclude field to given value.

### HasAudienceInclude

`func (o *TargetingSpec) HasAudienceInclude() bool`

HasAudienceInclude returns a boolean if a field has been set.

### GetAudienceExclude

`func (o *TargetingSpec) GetAudienceExclude() []string`

GetAudienceExclude returns the AudienceExclude field if non-nil, zero value otherwise.

### GetAudienceExcludeOk

`func (o *TargetingSpec) GetAudienceExcludeOk() (*[]string, bool)`

GetAudienceExcludeOk returns a tuple with the AudienceExclude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudienceExclude

`func (o *TargetingSpec) SetAudienceExclude(v []string)`

SetAudienceExclude sets AudienceExclude field to given value.

### HasAudienceExclude

`func (o *TargetingSpec) HasAudienceExclude() bool`

HasAudienceExclude returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


