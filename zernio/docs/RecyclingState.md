# RecyclingState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether recycling is currently active | [optional] 
**Gap** | Pointer to **int32** | Number of interval units between reposts | [optional] 
**GapFreq** | Pointer to **string** | Interval unit (week or month) | [optional] 
**StartDate** | Pointer to **time.Time** |  | [optional] 
**ExpireCount** | Pointer to **int32** |  | [optional] 
**ExpireDate** | Pointer to **time.Time** |  | [optional] 
**ContentVariations** | Pointer to **[]string** | Content variations for recycled copies (if configured) | [optional] 
**ContentVariationIndex** | Pointer to **int32** | Current position in the content variations rotation (read-only) | [optional] 
**RecycleCount** | Pointer to **int32** | How many recycled copies have been created so far (read-only) | [optional] 
**NextRecycleAt** | Pointer to **time.Time** | When the next recycled copy will be created (read-only) | [optional] 
**LastRecycledAt** | Pointer to **time.Time** | When the last recycled copy was created (read-only) | [optional] 

## Methods

### NewRecyclingState

`func NewRecyclingState() *RecyclingState`

NewRecyclingState instantiates a new RecyclingState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecyclingStateWithDefaults

`func NewRecyclingStateWithDefaults() *RecyclingState`

NewRecyclingStateWithDefaults instantiates a new RecyclingState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *RecyclingState) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *RecyclingState) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *RecyclingState) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *RecyclingState) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGap

`func (o *RecyclingState) GetGap() int32`

GetGap returns the Gap field if non-nil, zero value otherwise.

### GetGapOk

`func (o *RecyclingState) GetGapOk() (*int32, bool)`

GetGapOk returns a tuple with the Gap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGap

`func (o *RecyclingState) SetGap(v int32)`

SetGap sets Gap field to given value.

### HasGap

`func (o *RecyclingState) HasGap() bool`

HasGap returns a boolean if a field has been set.

### GetGapFreq

`func (o *RecyclingState) GetGapFreq() string`

GetGapFreq returns the GapFreq field if non-nil, zero value otherwise.

### GetGapFreqOk

`func (o *RecyclingState) GetGapFreqOk() (*string, bool)`

GetGapFreqOk returns a tuple with the GapFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGapFreq

`func (o *RecyclingState) SetGapFreq(v string)`

SetGapFreq sets GapFreq field to given value.

### HasGapFreq

`func (o *RecyclingState) HasGapFreq() bool`

HasGapFreq returns a boolean if a field has been set.

### GetStartDate

`func (o *RecyclingState) GetStartDate() time.Time`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *RecyclingState) GetStartDateOk() (*time.Time, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *RecyclingState) SetStartDate(v time.Time)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *RecyclingState) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetExpireCount

`func (o *RecyclingState) GetExpireCount() int32`

GetExpireCount returns the ExpireCount field if non-nil, zero value otherwise.

### GetExpireCountOk

`func (o *RecyclingState) GetExpireCountOk() (*int32, bool)`

GetExpireCountOk returns a tuple with the ExpireCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireCount

`func (o *RecyclingState) SetExpireCount(v int32)`

SetExpireCount sets ExpireCount field to given value.

### HasExpireCount

`func (o *RecyclingState) HasExpireCount() bool`

HasExpireCount returns a boolean if a field has been set.

### GetExpireDate

`func (o *RecyclingState) GetExpireDate() time.Time`

GetExpireDate returns the ExpireDate field if non-nil, zero value otherwise.

### GetExpireDateOk

`func (o *RecyclingState) GetExpireDateOk() (*time.Time, bool)`

GetExpireDateOk returns a tuple with the ExpireDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireDate

`func (o *RecyclingState) SetExpireDate(v time.Time)`

SetExpireDate sets ExpireDate field to given value.

### HasExpireDate

`func (o *RecyclingState) HasExpireDate() bool`

HasExpireDate returns a boolean if a field has been set.

### GetContentVariations

`func (o *RecyclingState) GetContentVariations() []string`

GetContentVariations returns the ContentVariations field if non-nil, zero value otherwise.

### GetContentVariationsOk

`func (o *RecyclingState) GetContentVariationsOk() (*[]string, bool)`

GetContentVariationsOk returns a tuple with the ContentVariations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentVariations

`func (o *RecyclingState) SetContentVariations(v []string)`

SetContentVariations sets ContentVariations field to given value.

### HasContentVariations

`func (o *RecyclingState) HasContentVariations() bool`

HasContentVariations returns a boolean if a field has been set.

### GetContentVariationIndex

`func (o *RecyclingState) GetContentVariationIndex() int32`

GetContentVariationIndex returns the ContentVariationIndex field if non-nil, zero value otherwise.

### GetContentVariationIndexOk

`func (o *RecyclingState) GetContentVariationIndexOk() (*int32, bool)`

GetContentVariationIndexOk returns a tuple with the ContentVariationIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentVariationIndex

`func (o *RecyclingState) SetContentVariationIndex(v int32)`

SetContentVariationIndex sets ContentVariationIndex field to given value.

### HasContentVariationIndex

`func (o *RecyclingState) HasContentVariationIndex() bool`

HasContentVariationIndex returns a boolean if a field has been set.

### GetRecycleCount

`func (o *RecyclingState) GetRecycleCount() int32`

GetRecycleCount returns the RecycleCount field if non-nil, zero value otherwise.

### GetRecycleCountOk

`func (o *RecyclingState) GetRecycleCountOk() (*int32, bool)`

GetRecycleCountOk returns a tuple with the RecycleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecycleCount

`func (o *RecyclingState) SetRecycleCount(v int32)`

SetRecycleCount sets RecycleCount field to given value.

### HasRecycleCount

`func (o *RecyclingState) HasRecycleCount() bool`

HasRecycleCount returns a boolean if a field has been set.

### GetNextRecycleAt

`func (o *RecyclingState) GetNextRecycleAt() time.Time`

GetNextRecycleAt returns the NextRecycleAt field if non-nil, zero value otherwise.

### GetNextRecycleAtOk

`func (o *RecyclingState) GetNextRecycleAtOk() (*time.Time, bool)`

GetNextRecycleAtOk returns a tuple with the NextRecycleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRecycleAt

`func (o *RecyclingState) SetNextRecycleAt(v time.Time)`

SetNextRecycleAt sets NextRecycleAt field to given value.

### HasNextRecycleAt

`func (o *RecyclingState) HasNextRecycleAt() bool`

HasNextRecycleAt returns a boolean if a field has been set.

### GetLastRecycledAt

`func (o *RecyclingState) GetLastRecycledAt() time.Time`

GetLastRecycledAt returns the LastRecycledAt field if non-nil, zero value otherwise.

### GetLastRecycledAtOk

`func (o *RecyclingState) GetLastRecycledAtOk() (*time.Time, bool)`

GetLastRecycledAtOk returns a tuple with the LastRecycledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRecycledAt

`func (o *RecyclingState) SetLastRecycledAt(v time.Time)`

SetLastRecycledAt sets LastRecycledAt field to given value.

### HasLastRecycledAt

`func (o *RecyclingState) HasLastRecycledAt() bool`

HasLastRecycledAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


