# YouTubeDemographicsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** | The Zernio SocialAccount ID | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Demographics** | Pointer to [**map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner**](array.md) | Object keyed by breakdown dimension (age, gender, country) | [optional] 
**DateRange** | Pointer to [**YouTubeDemographicsResponseDateRange**](YouTubeDemographicsResponseDateRange.md) |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 

## Methods

### NewYouTubeDemographicsResponse

`func NewYouTubeDemographicsResponse() *YouTubeDemographicsResponse`

NewYouTubeDemographicsResponse instantiates a new YouTubeDemographicsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewYouTubeDemographicsResponseWithDefaults

`func NewYouTubeDemographicsResponseWithDefaults() *YouTubeDemographicsResponse`

NewYouTubeDemographicsResponseWithDefaults instantiates a new YouTubeDemographicsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *YouTubeDemographicsResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *YouTubeDemographicsResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *YouTubeDemographicsResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *YouTubeDemographicsResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *YouTubeDemographicsResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *YouTubeDemographicsResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *YouTubeDemographicsResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *YouTubeDemographicsResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *YouTubeDemographicsResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *YouTubeDemographicsResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *YouTubeDemographicsResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *YouTubeDemographicsResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetDemographics

`func (o *YouTubeDemographicsResponse) GetDemographics() map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner`

GetDemographics returns the Demographics field if non-nil, zero value otherwise.

### GetDemographicsOk

`func (o *YouTubeDemographicsResponse) GetDemographicsOk() (*map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner, bool)`

GetDemographicsOk returns a tuple with the Demographics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemographics

`func (o *YouTubeDemographicsResponse) SetDemographics(v map[string][]InstagramAccountInsightsResponseMetricsValueBreakdownsInner)`

SetDemographics sets Demographics field to given value.

### HasDemographics

`func (o *YouTubeDemographicsResponse) HasDemographics() bool`

HasDemographics returns a boolean if a field has been set.

### GetDateRange

`func (o *YouTubeDemographicsResponse) GetDateRange() YouTubeDemographicsResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *YouTubeDemographicsResponse) GetDateRangeOk() (*YouTubeDemographicsResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *YouTubeDemographicsResponse) SetDateRange(v YouTubeDemographicsResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *YouTubeDemographicsResponse) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetNote

`func (o *YouTubeDemographicsResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *YouTubeDemographicsResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *YouTubeDemographicsResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *YouTubeDemographicsResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


