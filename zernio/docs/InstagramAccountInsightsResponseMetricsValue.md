# InstagramAccountInsightsResponseMetricsValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **float32** | Sum or aggregate value for the metric | [optional] 
**Values** | Pointer to [**[]InstagramAccountInsightsResponseMetricsValueValuesInner**](InstagramAccountInsightsResponseMetricsValueValuesInner.md) | Daily values (only for time_series) | [optional] 
**Breakdowns** | Pointer to [**[]InstagramAccountInsightsResponseMetricsValueBreakdownsInner**](InstagramAccountInsightsResponseMetricsValueBreakdownsInner.md) | Breakdown values (only for total_value with breakdown) | [optional] 

## Methods

### NewInstagramAccountInsightsResponseMetricsValue

`func NewInstagramAccountInsightsResponseMetricsValue() *InstagramAccountInsightsResponseMetricsValue`

NewInstagramAccountInsightsResponseMetricsValue instantiates a new InstagramAccountInsightsResponseMetricsValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstagramAccountInsightsResponseMetricsValueWithDefaults

`func NewInstagramAccountInsightsResponseMetricsValueWithDefaults() *InstagramAccountInsightsResponseMetricsValue`

NewInstagramAccountInsightsResponseMetricsValueWithDefaults instantiates a new InstagramAccountInsightsResponseMetricsValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *InstagramAccountInsightsResponseMetricsValue) GetTotal() float32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *InstagramAccountInsightsResponseMetricsValue) GetTotalOk() (*float32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *InstagramAccountInsightsResponseMetricsValue) SetTotal(v float32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *InstagramAccountInsightsResponseMetricsValue) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *InstagramAccountInsightsResponseMetricsValue) GetValues() []InstagramAccountInsightsResponseMetricsValueValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *InstagramAccountInsightsResponseMetricsValue) GetValuesOk() (*[]InstagramAccountInsightsResponseMetricsValueValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *InstagramAccountInsightsResponseMetricsValue) SetValues(v []InstagramAccountInsightsResponseMetricsValueValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *InstagramAccountInsightsResponseMetricsValue) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetBreakdowns

`func (o *InstagramAccountInsightsResponseMetricsValue) GetBreakdowns() []InstagramAccountInsightsResponseMetricsValueBreakdownsInner`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *InstagramAccountInsightsResponseMetricsValue) GetBreakdownsOk() (*[]InstagramAccountInsightsResponseMetricsValueBreakdownsInner, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *InstagramAccountInsightsResponseMetricsValue) SetBreakdowns(v []InstagramAccountInsightsResponseMetricsValueBreakdownsInner)`

SetBreakdowns sets Breakdowns field to given value.

### HasBreakdowns

`func (o *InstagramAccountInsightsResponseMetricsValue) HasBreakdowns() bool`

HasBreakdowns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


