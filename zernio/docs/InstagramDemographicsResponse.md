# InstagramDemographicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** | The Zernio SocialAccount ID | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Metric** | Pointer to **string** |  | [optional] 
**Timeframe** | Pointer to **string** | The timeframe used for demographic data | [optional] 
**Demographics** | Pointer to [**map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner**](array.md) | Object keyed by breakdown dimension (age, city, country, gender) | [optional] 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewInstagramDemographicsResponse

`func NewInstagramDemographicsResponse() *InstagramDemographicsResponse`

NewInstagramDemographicsResponse instantiates a new InstagramDemographicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstagramDemographicsResponseWithDefaults

`func NewInstagramDemographicsResponseWithDefaults() *InstagramDemographicsResponse`

NewInstagramDemographicsResponseWithDefaults instantiates a new InstagramDemographicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *InstagramDemographicsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *InstagramDemographicsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *InstagramDemographicsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *InstagramDemographicsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *InstagramDemographicsResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *InstagramDemographicsResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *InstagramDemographicsResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *InstagramDemographicsResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *InstagramDemographicsResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InstagramDemographicsResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InstagramDemographicsResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *InstagramDemographicsResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetMetric

`func (o *InstagramDemographicsResponse) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *InstagramDemographicsResponse) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *InstagramDemographicsResponse) SetMetric(v string)`

SetMetric sets Metric field to given value.

### HasMetric

`func (o *InstagramDemographicsResponse) HasMetric() bool`

HasMetric returns a boolean if a field has been set.

### GetTimeframe

`func (o *InstagramDemographicsResponse) GetTimeframe() string`

GetTimeframe returns the Timeframe field if non-nil, zero value otherwise.

### GetTimeframeOk

`func (o *InstagramDemographicsResponse) GetTimeframeOk() (*string, bool)`

GetTimeframeOk returns a tuple with the Timeframe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeframe

`func (o *InstagramDemographicsResponse) SetTimeframe(v string)`

SetTimeframe sets Timeframe field to given value.

### HasTimeframe

`func (o *InstagramDemographicsResponse) HasTimeframe() bool`

HasTimeframe returns a boolean if a field has been set.

### GetDemographics

`func (o *InstagramDemographicsResponse) GetDemographics() map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner`

GetDemographics returns the Demographics field if non-nil, zero value otherwise.

### GetDemographicsOk

`func (o *InstagramDemographicsResponse) GetDemographicsOk() (*map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner, bool)`

GetDemographicsOk returns a tuple with the Demographics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographics

`func (o *InstagramDemographicsResponse) SetDemographics(v map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner)`

SetDemographics sets Demographics field to given value.

### HasDemographics

`func (o *InstagramDemographicsResponse) HasDemographics() bool`

HasDemographics returns a boolean if a field has been set.

### GetNote

`func (o *InstagramDemographicsResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *InstagramDemographicsResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *InstagramDemographicsResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *InstagramDemographicsResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


