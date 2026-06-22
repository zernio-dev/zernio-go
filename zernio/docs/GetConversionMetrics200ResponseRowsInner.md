# GetConversionMetrics200ResponseRowsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Start** | Pointer to **string** | YYYY-MM-DD | [optional] 
**End** | Pointer to **string** | YYYY-MM-DD (inclusive) | [optional] 
**Metrics** | Pointer to [**map[string]GetConversionMetrics200ResponseRowsInnerMetricsValue**](GetConversionMetrics200ResponseRowsInnerMetricsValue.md) |  | [optional] 

## Methods

### NewGetConversionMetrics200ResponseRowsInner

`func NewGetConversionMetrics200ResponseRowsInner() *GetConversionMetrics200ResponseRowsInner`

NewGetConversionMetrics200ResponseRowsInner instantiates a new GetConversionMetrics200ResponseRowsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConversionMetrics200ResponseRowsInnerWithDefaults

`func NewGetConversionMetrics200ResponseRowsInnerWithDefaults() *GetConversionMetrics200ResponseRowsInner`

NewGetConversionMetrics200ResponseRowsInnerWithDefaults instantiates a new GetConversionMetrics200ResponseRowsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStart

`func (o *GetConversionMetrics200ResponseRowsInner) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *GetConversionMetrics200ResponseRowsInner) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *GetConversionMetrics200ResponseRowsInner) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *GetConversionMetrics200ResponseRowsInner) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *GetConversionMetrics200ResponseRowsInner) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *GetConversionMetrics200ResponseRowsInner) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *GetConversionMetrics200ResponseRowsInner) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *GetConversionMetrics200ResponseRowsInner) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetMetrics

`func (o *GetConversionMetrics200ResponseRowsInner) GetMetrics() map[string]GetConversionMetrics200ResponseRowsInnerMetricsValue`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *GetConversionMetrics200ResponseRowsInner) GetMetricsOk() (*map[string]GetConversionMetrics200ResponseRowsInnerMetricsValue, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *GetConversionMetrics200ResponseRowsInner) SetMetrics(v map[string]GetConversionMetrics200ResponseRowsInnerMetricsValue)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *GetConversionMetrics200ResponseRowsInner) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


