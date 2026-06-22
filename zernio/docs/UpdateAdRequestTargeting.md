# UpdateAdRequestTargeting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgeMin** | Pointer to **int32** |  | [optional] 
**AgeMax** | Pointer to **int32** |  | [optional] 
**Countries** | Pointer to **[]string** |  | [optional] 
**Interests** | Pointer to [**[]UpdateAdRequestTargetingInterestsInner**](UpdateAdRequestTargetingInterestsInner.md) | Interest objects from /v1/ads/interests. Each must include id and name. | [optional] 
**AdvantageAudience** | Pointer to **int32** | Meta only. Omit to preserve the existing setting on update. 0 &#x3D; disabled, 1 &#x3D; enabled. | [optional] 

## Methods

### NewUpdateAdRequestTargeting

`func NewUpdateAdRequestTargeting() *UpdateAdRequestTargeting`

NewUpdateAdRequestTargeting instantiates a new UpdateAdRequestTargeting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAdRequestTargetingWithDefaults

`func NewUpdateAdRequestTargetingWithDefaults() *UpdateAdRequestTargeting`

NewUpdateAdRequestTargetingWithDefaults instantiates a new UpdateAdRequestTargeting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgeMin

`func (o *UpdateAdRequestTargeting) GetAgeMin() int32`

GetAgeMin returns the AgeMin field if non-nil, zero value otherwise.

### GetAgeMinOk

`func (o *UpdateAdRequestTargeting) GetAgeMinOk() (*int32, bool)`

GetAgeMinOk returns a tuple with the AgeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMin

`func (o *UpdateAdRequestTargeting) SetAgeMin(v int32)`

SetAgeMin sets AgeMin field to given value.

### HasAgeMin

`func (o *UpdateAdRequestTargeting) HasAgeMin() bool`

HasAgeMin returns a boolean if a field has been set.

### GetAgeMax

`func (o *UpdateAdRequestTargeting) GetAgeMax() int32`

GetAgeMax returns the AgeMax field if non-nil, zero value otherwise.

### GetAgeMaxOk

`func (o *UpdateAdRequestTargeting) GetAgeMaxOk() (*int32, bool)`

GetAgeMaxOk returns a tuple with the AgeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeMax

`func (o *UpdateAdRequestTargeting) SetAgeMax(v int32)`

SetAgeMax sets AgeMax field to given value.

### HasAgeMax

`func (o *UpdateAdRequestTargeting) HasAgeMax() bool`

HasAgeMax returns a boolean if a field has been set.

### GetCountries

`func (o *UpdateAdRequestTargeting) GetCountries() []string`

GetCountries returns the Countries field if non-nil, zero value otherwise.

### GetCountriesOk

`func (o *UpdateAdRequestTargeting) GetCountriesOk() (*[]string, bool)`

GetCountriesOk returns a tuple with the Countries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountries

`func (o *UpdateAdRequestTargeting) SetCountries(v []string)`

SetCountries sets Countries field to given value.

### HasCountries

`func (o *UpdateAdRequestTargeting) HasCountries() bool`

HasCountries returns a boolean if a field has been set.

### GetInterests

`func (o *UpdateAdRequestTargeting) GetInterests() []UpdateAdRequestTargetingInterestsInner`

GetInterests returns the Interests field if non-nil, zero value otherwise.

### GetInterestsOk

`func (o *UpdateAdRequestTargeting) GetInterestsOk() (*[]UpdateAdRequestTargetingInterestsInner, bool)`

GetInterestsOk returns a tuple with the Interests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterests

`func (o *UpdateAdRequestTargeting) SetInterests(v []UpdateAdRequestTargetingInterestsInner)`

SetInterests sets Interests field to given value.

### HasInterests

`func (o *UpdateAdRequestTargeting) HasInterests() bool`

HasInterests returns a boolean if a field has been set.

### GetAdvantageAudience

`func (o *UpdateAdRequestTargeting) GetAdvantageAudience() int32`

GetAdvantageAudience returns the AdvantageAudience field if non-nil, zero value otherwise.

### GetAdvantageAudienceOk

`func (o *UpdateAdRequestTargeting) GetAdvantageAudienceOk() (*int32, bool)`

GetAdvantageAudienceOk returns a tuple with the AdvantageAudience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvantageAudience

`func (o *UpdateAdRequestTargeting) SetAdvantageAudience(v int32)`

SetAdvantageAudience sets AdvantageAudience field to given value.

### HasAdvantageAudience

`func (o *UpdateAdRequestTargeting) HasAdvantageAudience() bool`

HasAdvantageAudience returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


