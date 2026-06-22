# RecyclingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Set to false to disable recycling on this post | [optional] [default to true]
**Gap** | Pointer to **int32** | Number of interval units between each repost. Required when enabling recycling. | [optional] 
**GapFreq** | Pointer to **string** | Interval unit for the gap. Defaults to &#39;month&#39;. | [optional] [default to "month"]
**StartDate** | Pointer to **time.Time** | When to start the recycling cycle. Defaults to the post&#39;s scheduledFor date. | [optional] 
**ExpireCount** | Pointer to **NullableInt32** | Stop recycling after this many copies have been created. Send null on update to clear this limit. | [optional] 
**ExpireDate** | Pointer to **NullableTime** | Stop recycling after this date, regardless of count. Send null on update to clear this limit. | [optional] 
**ContentVariations** | Pointer to **[]string** | Array of content variations for recycled copies. On each recycle, the next variation is used in round-robin order. Recommended for Twitter and Pinterest to avoid duplicate content flags. If omitted, the original post content is used for all recycled copies. Send an empty array [] to clear existing variations. Must have 2+ entries when setting variations. Platform-level customContent still overrides the base content per platform.  | [optional] 

## Methods

### NewRecyclingConfig

`func NewRecyclingConfig() *RecyclingConfig`

NewRecyclingConfig instantiates a new RecyclingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecyclingConfigWithDefaults

`func NewRecyclingConfigWithDefaults() *RecyclingConfig`

NewRecyclingConfigWithDefaults instantiates a new RecyclingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RecyclingConfig) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecyclingConfig) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecyclingConfig) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecyclingConfig) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGap

`func (o *RecyclingConfig) GetGap() int32`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *RecyclingConfig) GetGapOk() (*int32, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *RecyclingConfig) SetGap(v int32)`

SetGap sets Gap field to given value.

### HasGap

`func (o *RecyclingConfig) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetGapFreq

`func (o *RecyclingConfig) GetGapFreq() string`

GetGapFreq returns the GapFreq field if non-nil, zero value otherwise.

### GetGapFreqOk

`func (o *RecyclingConfig) GetGapFreqOk() (*string, bool)`

GetGapFreqOk returns a tuple with the GapFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapFreq

`func (o *RecyclingConfig) SetGapFreq(v string)`

SetGapFreq sets GapFreq field to given value.

### HasGapFreq

`func (o *RecyclingConfig) HasGapFreq() bool`

HasGapFreq returns a boolean if a field has been set.

### GetStartDate

`func (o *RecyclingConfig) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RecyclingConfig) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RecyclingConfig) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RecyclingConfig) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpireCount

`func (o *RecyclingConfig) GetExpireCount() int32`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *RecyclingConfig) GetExpireCountOk() (*int32, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *RecyclingConfig) SetExpireCount(v int32)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *RecyclingConfig) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### SetExpireCountNil

`func (o *RecyclingConfig) SetExpireCountNil(b bool)`

 SetExpireCountNil sets the value for ExpireCount to be an explicit nil

### UnsetExpireCount
`func (o *RecyclingConfig) UnsetExpireCount()`

UnsetExpireCount ensures that no value is present for ExpireCount, not even an explicit nil
### GetExpireDate

`func (o *RecyclingConfig) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *RecyclingConfig) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *RecyclingConfig) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *RecyclingConfig) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### SetExpireDateNil

`func (o *RecyclingConfig) SetExpireDateNil(b bool)`

 SetExpireDateNil sets the value for ExpireDate to be an explicit nil

### UnsetExpireDate
`func (o *RecyclingConfig) UnsetExpireDate()`

UnsetExpireDate ensures that no value is present for ExpireDate, not even an explicit nil
### GetContentVariations

`func (o *RecyclingConfig) GetContentVariations() []string`

GetContentVariations returns the ContentVariations field if non-nil, zero value otherwise.

### GetContentVariationsOk

`func (o *RecyclingConfig) GetContentVariationsOk() (*[]string, bool)`

GetContentVariationsOk returns a tuple with the ContentVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentVariations

`func (o *RecyclingConfig) SetContentVariations(v []string)`

SetContentVariations sets ContentVariations field to given value.

### HasContentVariations

`func (o *RecyclingConfig) HasContentVariations() bool`

HasContentVariations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


