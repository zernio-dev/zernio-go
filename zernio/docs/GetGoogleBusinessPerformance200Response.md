# GetGoogleBusinessPerformance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**DateRange** | Pointer to [**GetGoogleBusinessPerformance200ResponseDateRange**](GetGoogleBusinessPerformance200ResponseDateRange.md) |  | [optional] 
**Metrics** | Pointer to [**map[string]GetGoogleBusinessPerformance200ResponseMetricsValue**](GetGoogleBusinessPerformance200ResponseMetricsValue.md) | Each key is a metric name containing total and daily values. | [optional] 
**DataDelay** | Pointer to **string** |  | [optional] 

## Methods

### NewGetGoogleBusinessPerformance200Response

`func NewGetGoogleBusinessPerformance200Response() *GetGoogleBusinessPerformance200Response`

NewGetGoogleBusinessPerformance200Response instantiates a new GetGoogleBusinessPerformance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessPerformance200ResponseWithDefaults

`func NewGetGoogleBusinessPerformance200ResponseWithDefaults() *GetGoogleBusinessPerformance200Response`

NewGetGoogleBusinessPerformance200ResponseWithDefaults instantiates a new GetGoogleBusinessPerformance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGoogleBusinessPerformance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGoogleBusinessPerformance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGoogleBusinessPerformance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGoogleBusinessPerformance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGoogleBusinessPerformance200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGoogleBusinessPerformance200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGoogleBusinessPerformance200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGoogleBusinessPerformance200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetGoogleBusinessPerformance200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetGoogleBusinessPerformance200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetGoogleBusinessPerformance200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetGoogleBusinessPerformance200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDateRange

`func (o *GetGoogleBusinessPerformance200Response) GetDateRange() GetGoogleBusinessPerformance200ResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *GetGoogleBusinessPerformance200Response) GetDateRangeOk() (*GetGoogleBusinessPerformance200ResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *GetGoogleBusinessPerformance200Response) SetDateRange(v GetGoogleBusinessPerformance200ResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *GetGoogleBusinessPerformance200Response) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetMetrics

`func (o *GetGoogleBusinessPerformance200Response) GetMetrics() map[string]GetGoogleBusinessPerformance200ResponseMetricsValue`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetGoogleBusinessPerformance200Response) GetMetricsOk() (*map[string]GetGoogleBusinessPerformance200ResponseMetricsValue, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetGoogleBusinessPerformance200Response) SetMetrics(v map[string]GetGoogleBusinessPerformance200ResponseMetricsValue)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetGoogleBusinessPerformance200Response) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetDataDelay

`func (o *GetGoogleBusinessPerformance200Response) GetDataDelay() string`

GetDataDelay returns the DataDelay field if non-nil, zero value otherwise.

### GetDataDelayOk

`func (o *GetGoogleBusinessPerformance200Response) GetDataDelayOk() (*string, bool)`

GetDataDelayOk returns a tuple with the DataDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDelay

`func (o *GetGoogleBusinessPerformance200Response) SetDataDelay(v string)`

SetDataDelay sets DataDelay field to given value.

### HasDataDelay

`func (o *GetGoogleBusinessPerformance200Response) HasDataDelay() bool`

HasDataDelay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


