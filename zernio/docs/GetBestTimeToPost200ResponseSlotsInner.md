# GetBestTimeToPost200ResponseSlotsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DayOfWeek** | Pointer to **int32** | 0&#x3D;Monday, 6&#x3D;Sunday | [optional] 
**Hour** | Pointer to **int32** | Hour in UTC (0-23) | [optional] 
**AvgEngagement** | Pointer to **float32** | Average engagement (likes + comments + shares + saves) | [optional] 
**PostCount** | Pointer to **int32** | Number of posts in this slot | [optional] 

## Methods

### NewGetBestTimeToPost200ResponseSlotsInner

`func NewGetBestTimeToPost200ResponseSlotsInner() *GetBestTimeToPost200ResponseSlotsInner`

NewGetBestTimeToPost200ResponseSlotsInner instantiates a new GetBestTimeToPost200ResponseSlotsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBestTimeToPost200ResponseSlotsInnerWithDefaults

`func NewGetBestTimeToPost200ResponseSlotsInnerWithDefaults() *GetBestTimeToPost200ResponseSlotsInner`

NewGetBestTimeToPost200ResponseSlotsInnerWithDefaults instantiates a new GetBestTimeToPost200ResponseSlotsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDayOfWeek

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetDayOfWeek() int32`

GetDayOfWeek returns the DayOfWeek field if non-nil, zero value otherwise.

### GetDayOfWeekOk

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetDayOfWeekOk() (*int32, bool)`

GetDayOfWeekOk returns a tuple with the DayOfWeek field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayOfWeek

`func (o *GetBestTimeToPost200ResponseSlotsInner) SetDayOfWeek(v int32)`

SetDayOfWeek sets DayOfWeek field to given value.

### HasDayOfWeek

`func (o *GetBestTimeToPost200ResponseSlotsInner) HasDayOfWeek() bool`

HasDayOfWeek returns a boolean if a field has been set.

### GetHour

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetHour() int32`

GetHour returns the Hour field if non-nil, zero value otherwise.

### GetHourOk

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetHourOk() (*int32, bool)`

GetHourOk returns a tuple with the Hour field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHour

`func (o *GetBestTimeToPost200ResponseSlotsInner) SetHour(v int32)`

SetHour sets Hour field to given value.

### HasHour

`func (o *GetBestTimeToPost200ResponseSlotsInner) HasHour() bool`

HasHour returns a boolean if a field has been set.

### GetAvgEngagement

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetAvgEngagement() float32`

GetAvgEngagement returns the AvgEngagement field if non-nil, zero value otherwise.

### GetAvgEngagementOk

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetAvgEngagementOk() (*float32, bool)`

GetAvgEngagementOk returns a tuple with the AvgEngagement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgEngagement

`func (o *GetBestTimeToPost200ResponseSlotsInner) SetAvgEngagement(v float32)`

SetAvgEngagement sets AvgEngagement field to given value.

### HasAvgEngagement

`func (o *GetBestTimeToPost200ResponseSlotsInner) HasAvgEngagement() bool`

HasAvgEngagement returns a boolean if a field has been set.

### GetPostCount

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetPostCount() int32`

GetPostCount returns the PostCount field if non-nil, zero value otherwise.

### GetPostCountOk

`func (o *GetBestTimeToPost200ResponseSlotsInner) GetPostCountOk() (*int32, bool)`

GetPostCountOk returns a tuple with the PostCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostCount

`func (o *GetBestTimeToPost200ResponseSlotsInner) SetPostCount(v int32)`

SetPostCount sets PostCount field to given value.

### HasPostCount

`func (o *GetBestTimeToPost200ResponseSlotsInner) HasPostCount() bool`

HasPostCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


