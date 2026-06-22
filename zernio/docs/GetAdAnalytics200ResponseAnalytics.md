# GetAdAnalytics200ResponseAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Summary** | Pointer to [**AdMetrics**](AdMetrics.md) |  | [optional] 
**Daily** | Pointer to [**[]GetAdAnalytics200ResponseAnalyticsDailyInner**](GetAdAnalytics200ResponseAnalyticsDailyInner.md) |  | [optional] 
**Breakdowns** | Pointer to **map[string][]map[string]interface{}** |  | [optional] 

## Methods

### NewGetAdAnalytics200ResponseAnalytics

`func NewGetAdAnalytics200ResponseAnalytics() *GetAdAnalytics200ResponseAnalytics`

NewGetAdAnalytics200ResponseAnalytics instantiates a new GetAdAnalytics200ResponseAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdAnalytics200ResponseAnalyticsWithDefaults

`func NewGetAdAnalytics200ResponseAnalyticsWithDefaults() *GetAdAnalytics200ResponseAnalytics`

NewGetAdAnalytics200ResponseAnalyticsWithDefaults instantiates a new GetAdAnalytics200ResponseAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSummary

`func (o *GetAdAnalytics200ResponseAnalytics) GetSummary() AdMetrics`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetAdAnalytics200ResponseAnalytics) GetSummaryOk() (*AdMetrics, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetAdAnalytics200ResponseAnalytics) SetSummary(v AdMetrics)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *GetAdAnalytics200ResponseAnalytics) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetDaily

`func (o *GetAdAnalytics200ResponseAnalytics) GetDaily() []GetAdAnalytics200ResponseAnalyticsDailyInner`

GetDaily returns the Daily field if non-nil, zero value otherwise.

### GetDailyOk

`func (o *GetAdAnalytics200ResponseAnalytics) GetDailyOk() (*[]GetAdAnalytics200ResponseAnalyticsDailyInner, bool)`

GetDailyOk returns a tuple with the Daily field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaily

`func (o *GetAdAnalytics200ResponseAnalytics) SetDaily(v []GetAdAnalytics200ResponseAnalyticsDailyInner)`

SetDaily sets Daily field to given value.

### HasDaily

`func (o *GetAdAnalytics200ResponseAnalytics) HasDaily() bool`

HasDaily returns a boolean if a field has been set.

### GetBreakdowns

`func (o *GetAdAnalytics200ResponseAnalytics) GetBreakdowns() map[string][]map[string]interface{}`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *GetAdAnalytics200ResponseAnalytics) GetBreakdownsOk() (*map[string][]map[string]interface{}, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *GetAdAnalytics200ResponseAnalytics) SetBreakdowns(v map[string][]map[string]interface{})`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *GetAdAnalytics200ResponseAnalytics) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


