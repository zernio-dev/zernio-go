# GetLinkedInAggregateAnalytics200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Aggregation** | Pointer to **string** |  | [optional] 
**DateRange** | Pointer to [**LinkedInAggregateAnalyticsTotalResponseDateRange**](LinkedInAggregateAnalyticsTotalResponseDateRange.md) |  | [optional] 
**Analytics** | Pointer to [**LinkedInAggregateAnalyticsDailyResponseAnalytics**](LinkedInAggregateAnalyticsDailyResponseAnalytics.md) |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**SkippedMetrics** | Pointer to **[]string** | Metrics that were skipped due to API limitations | [optional] 

## Methods

### NewGetLinkedInAggregateAnalytics200Response

`func NewGetLinkedInAggregateAnalytics200Response() *GetLinkedInAggregateAnalytics200Response`

NewGetLinkedInAggregateAnalytics200Response instantiates a new GetLinkedInAggregateAnalytics200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLinkedInAggregateAnalytics200ResponseWithDefaults

`func NewGetLinkedInAggregateAnalytics200ResponseWithDefaults() *GetLinkedInAggregateAnalytics200Response`

NewGetLinkedInAggregateAnalytics200ResponseWithDefaults instantiates a new GetLinkedInAggregateAnalytics200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetLinkedInAggregateAnalytics200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetLinkedInAggregateAnalytics200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetLinkedInAggregateAnalytics200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetLinkedInAggregateAnalytics200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetLinkedInAggregateAnalytics200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetLinkedInAggregateAnalytics200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountType

`func (o *GetLinkedInAggregateAnalytics200Response) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *GetLinkedInAggregateAnalytics200Response) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *GetLinkedInAggregateAnalytics200Response) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetUsername

`func (o *GetLinkedInAggregateAnalytics200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetLinkedInAggregateAnalytics200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetLinkedInAggregateAnalytics200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAggregation

`func (o *GetLinkedInAggregateAnalytics200Response) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *GetLinkedInAggregateAnalytics200Response) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *GetLinkedInAggregateAnalytics200Response) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetDateRange

`func (o *GetLinkedInAggregateAnalytics200Response) GetDateRange() LinkedInAggregateAnalyticsTotalResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetDateRangeOk() (*LinkedInAggregateAnalyticsTotalResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *GetLinkedInAggregateAnalytics200Response) SetDateRange(v LinkedInAggregateAnalyticsTotalResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *GetLinkedInAggregateAnalytics200Response) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetAnalytics

`func (o *GetLinkedInAggregateAnalytics200Response) GetAnalytics() LinkedInAggregateAnalyticsDailyResponseAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetAnalyticsOk() (*LinkedInAggregateAnalyticsDailyResponseAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *GetLinkedInAggregateAnalytics200Response) SetAnalytics(v LinkedInAggregateAnalyticsDailyResponseAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *GetLinkedInAggregateAnalytics200Response) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetNote

`func (o *GetLinkedInAggregateAnalytics200Response) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *GetLinkedInAggregateAnalytics200Response) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *GetLinkedInAggregateAnalytics200Response) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetLinkedInAggregateAnalytics200Response) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetLinkedInAggregateAnalytics200Response) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetLinkedInAggregateAnalytics200Response) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetSkippedMetrics

`func (o *GetLinkedInAggregateAnalytics200Response) GetSkippedMetrics() []string`

GetSkippedMetrics returns the SkippedMetrics field if non-nil, zero value otherwise.

### GetSkippedMetricsOk

`func (o *GetLinkedInAggregateAnalytics200Response) GetSkippedMetricsOk() (*[]string, bool)`

GetSkippedMetricsOk returns a tuple with the SkippedMetrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedMetrics

`func (o *GetLinkedInAggregateAnalytics200Response) SetSkippedMetrics(v []string)`

SetSkippedMetrics sets SkippedMetrics field to given value.

### HasSkippedMetrics

`func (o *GetLinkedInAggregateAnalytics200Response) HasSkippedMetrics() bool`

HasSkippedMetrics returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


