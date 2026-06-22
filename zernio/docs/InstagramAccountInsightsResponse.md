# InstagramAccountInsightsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** | The Zernio SocialAccount ID | [optional] 
**Platform** | Pointer to **string** | Platform that served this response. | [optional] 
**DateRange** | Pointer to [**InstagramAccountInsightsResponseDateRange**](InstagramAccountInsightsResponseDateRange.md) |  | [optional] 
**MetricType** | Pointer to **string** |  | [optional] 
**Breakdown** | Pointer to **string** | Breakdown dimension used (only present when breakdown was requested) | [optional] 
**Metrics** | Pointer to [**map[string]InstagramAccountInsightsResponseMetricsValue**](InstagramAccountInsightsResponseMetricsValue.md) | Object keyed by metric name. For time_series: each metric has \&quot;total\&quot; (number) and \&quot;values\&quot; (array of {date, value}). For total_value: each metric has \&quot;total\&quot; (number) and optionally \&quot;breakdowns\&quot; (array of {dimension, value}).  | [optional] 
**DataDelay** | Pointer to **string** |  | [optional] 

## Methods

### NewInstagramAccountInsightsResponse

`func NewInstagramAccountInsightsResponse() *InstagramAccountInsightsResponse`

NewInstagramAccountInsightsResponse instantiates a new InstagramAccountInsightsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstagramAccountInsightsResponseWithDefaults

`func NewInstagramAccountInsightsResponseWithDefaults() *InstagramAccountInsightsResponse`

NewInstagramAccountInsightsResponseWithDefaults instantiates a new InstagramAccountInsightsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InstagramAccountInsightsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InstagramAccountInsightsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InstagramAccountInsightsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InstagramAccountInsightsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *InstagramAccountInsightsResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstagramAccountInsightsResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstagramAccountInsightsResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstagramAccountInsightsResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *InstagramAccountInsightsResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InstagramAccountInsightsResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InstagramAccountInsightsResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *InstagramAccountInsightsResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDateRange

`func (o *InstagramAccountInsightsResponse) GetDateRange() InstagramAccountInsightsResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *InstagramAccountInsightsResponse) GetDateRangeOk() (*InstagramAccountInsightsResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *InstagramAccountInsightsResponse) SetDateRange(v InstagramAccountInsightsResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *InstagramAccountInsightsResponse) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetMetricType

`func (o *InstagramAccountInsightsResponse) GetMetricType() string`

GetMetricType returns the MetricType field if non-nil, zero value otherwise.

### GetMetricTypeOk

`func (o *InstagramAccountInsightsResponse) GetMetricTypeOk() (*string, bool)`

GetMetricTypeOk returns a tuple with the MetricType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricType

`func (o *InstagramAccountInsightsResponse) SetMetricType(v string)`

SetMetricType sets MetricType field to given value.

### HasMetricType

`func (o *InstagramAccountInsightsResponse) HasMetricType() bool`

HasMetricType returns a boolean if a field has been set.

### GetBreakdown

`func (o *InstagramAccountInsightsResponse) GetBreakdown() string`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *InstagramAccountInsightsResponse) GetBreakdownOk() (*string, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *InstagramAccountInsightsResponse) SetBreakdown(v string)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *InstagramAccountInsightsResponse) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### GetMetrics

`func (o *InstagramAccountInsightsResponse) GetMetrics() map[string]InstagramAccountInsightsResponseMetricsValue`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *InstagramAccountInsightsResponse) GetMetricsOk() (*map[string]InstagramAccountInsightsResponseMetricsValue, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *InstagramAccountInsightsResponse) SetMetrics(v map[string]InstagramAccountInsightsResponseMetricsValue)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *InstagramAccountInsightsResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetDataDelay

`func (o *InstagramAccountInsightsResponse) GetDataDelay() string`

GetDataDelay returns the DataDelay field if non-nil, zero value otherwise.

### GetDataDelayOk

`func (o *InstagramAccountInsightsResponse) GetDataDelayOk() (*string, bool)`

GetDataDelayOk returns a tuple with the DataDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDelay

`func (o *InstagramAccountInsightsResponse) SetDataDelay(v string)`

SetDataDelay sets DataDelay field to given value.

### HasDataDelay

`func (o *InstagramAccountInsightsResponse) HasDataDelay() bool`

HasDataDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


