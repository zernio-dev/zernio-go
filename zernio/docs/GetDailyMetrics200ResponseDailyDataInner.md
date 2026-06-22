# GetDailyMetrics200ResponseDailyDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**PostCount** | Pointer to **int32** |  | [optional] 
**Platforms** | Pointer to **map[string]int32** |  | [optional] 
**Metrics** | Pointer to [**GetDailyMetrics200ResponseDailyDataInnerMetrics**](GetDailyMetrics200ResponseDailyDataInnerMetrics.md) |  | [optional] 

## Methods

### NewGetDailyMetrics200ResponseDailyDataInner

`func NewGetDailyMetrics200ResponseDailyDataInner() *GetDailyMetrics200ResponseDailyDataInner`

NewGetDailyMetrics200ResponseDailyDataInner instantiates a new GetDailyMetrics200ResponseDailyDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyMetrics200ResponseDailyDataInnerWithDefaults

`func NewGetDailyMetrics200ResponseDailyDataInnerWithDefaults() *GetDailyMetrics200ResponseDailyDataInner`

NewGetDailyMetrics200ResponseDailyDataInnerWithDefaults instantiates a new GetDailyMetrics200ResponseDailyDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *GetDailyMetrics200ResponseDailyDataInner) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *GetDailyMetrics200ResponseDailyDataInner) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetPostCount

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *GetDailyMetrics200ResponseDailyDataInner) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.

### HasPostCount

`func (o *GetDailyMetrics200ResponseDailyDataInner) HasPostCount() bool`

HasPostCount returns a boolean if a field has been set.

### GetPlatforms

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetPlatforms() map[string]int32`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetPlatformsOk() (*map[string]int32, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *GetDailyMetrics200ResponseDailyDataInner) SetPlatforms(v map[string]int32)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *GetDailyMetrics200ResponseDailyDataInner) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### GetMetrics

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetMetrics() GetDailyMetrics200ResponseDailyDataInnerMetrics`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetDailyMetrics200ResponseDailyDataInner) GetMetricsOk() (*GetDailyMetrics200ResponseDailyDataInnerMetrics, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetDailyMetrics200ResponseDailyDataInner) SetMetrics(v GetDailyMetrics200ResponseDailyDataInnerMetrics)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetDailyMetrics200ResponseDailyDataInner) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


