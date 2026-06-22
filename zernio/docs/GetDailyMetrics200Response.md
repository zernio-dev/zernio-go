# GetDailyMetrics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DailyData** | Pointer to [**[]GetDailyMetrics200ResponseDailyDataInner**](GetDailyMetrics200ResponseDailyDataInner.md) |  | [optional] 
**PlatformBreakdown** | Pointer to [**[]GetDailyMetrics200ResponsePlatformBreakdownInner**](GetDailyMetrics200ResponsePlatformBreakdownInner.md) |  | [optional] 

## Methods

### NewGetDailyMetrics200Response

`func NewGetDailyMetrics200Response() *GetDailyMetrics200Response`

NewGetDailyMetrics200Response instantiates a new GetDailyMetrics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDailyMetrics200ResponseWithDefaults

`func NewGetDailyMetrics200ResponseWithDefaults() *GetDailyMetrics200Response`

NewGetDailyMetrics200ResponseWithDefaults instantiates a new GetDailyMetrics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDailyData

`func (o *GetDailyMetrics200Response) GetDailyData() []GetDailyMetrics200ResponseDailyDataInner`

GetDailyData returns the DailyData field if non-nil, zero value otherwise.

### GetDailyDataOk

`func (o *GetDailyMetrics200Response) GetDailyDataOk() (*[]GetDailyMetrics200ResponseDailyDataInner, bool)`

GetDailyDataOk returns a tuple with the DailyData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyData

`func (o *GetDailyMetrics200Response) SetDailyData(v []GetDailyMetrics200ResponseDailyDataInner)`

SetDailyData sets DailyData field to given value.

### HasDailyData

`func (o *GetDailyMetrics200Response) HasDailyData() bool`

HasDailyData returns a boolean if a field has been set.

### GetPlatformBreakdown

`func (o *GetDailyMetrics200Response) GetPlatformBreakdown() []GetDailyMetrics200ResponsePlatformBreakdownInner`

GetPlatformBreakdown returns the PlatformBreakdown field if non-nil, zero value otherwise.

### GetPlatformBreakdownOk

`func (o *GetDailyMetrics200Response) GetPlatformBreakdownOk() (*[]GetDailyMetrics200ResponsePlatformBreakdownInner, bool)`

GetPlatformBreakdownOk returns a tuple with the PlatformBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBreakdown

`func (o *GetDailyMetrics200Response) SetPlatformBreakdown(v []GetDailyMetrics200ResponsePlatformBreakdownInner)`

SetPlatformBreakdown sets PlatformBreakdown field to given value.

### HasPlatformBreakdown

`func (o *GetDailyMetrics200Response) HasPlatformBreakdown() bool`

HasPlatformBreakdown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


