# LinkedInAggregateAnalyticsTotalResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountType** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Aggregation** | Pointer to **string** |  | [optional] 
**DateRange** | Pointer to [**LinkedInAggregateAnalyticsTotalResponseDateRange**](LinkedInAggregateAnalyticsTotalResponseDateRange.md) |  | [optional] 
**Analytics** | Pointer to [**LinkedInAggregateAnalyticsTotalResponseAnalytics**](LinkedInAggregateAnalyticsTotalResponseAnalytics.md) |  | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewLinkedInAggregateAnalyticsTotalResponse

`func NewLinkedInAggregateAnalyticsTotalResponse() *LinkedInAggregateAnalyticsTotalResponse`

NewLinkedInAggregateAnalyticsTotalResponse instantiates a new LinkedInAggregateAnalyticsTotalResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkedInAggregateAnalyticsTotalResponseWithDefaults

`func NewLinkedInAggregateAnalyticsTotalResponseWithDefaults() *LinkedInAggregateAnalyticsTotalResponse`

NewLinkedInAggregateAnalyticsTotalResponseWithDefaults instantiates a new LinkedInAggregateAnalyticsTotalResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountType

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetUsername

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAggregation

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAggregation() string`

GetAggregation returns the Aggregation field if non-nil, zero value otherwise.

### GetAggregationOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAggregationOk() (*string, bool)`

GetAggregationOk returns a tuple with the Aggregation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregation

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetAggregation(v string)`

SetAggregation sets Aggregation field to given value.

### HasAggregation

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasAggregation() bool`

HasAggregation returns a boolean if a field has been set.

### GetDateRange

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetDateRange() LinkedInAggregateAnalyticsTotalResponseDateRange`

GetDateRange returns the DateRange field if non-nil, zero value otherwise.

### GetDateRangeOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetDateRangeOk() (*LinkedInAggregateAnalyticsTotalResponseDateRange, bool)`

GetDateRangeOk returns a tuple with the DateRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateRange

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetDateRange(v LinkedInAggregateAnalyticsTotalResponseDateRange)`

SetDateRange sets DateRange field to given value.

### HasDateRange

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasDateRange() bool`

HasDateRange returns a boolean if a field has been set.

### GetAnalytics

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAnalytics() LinkedInAggregateAnalyticsTotalResponseAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetAnalyticsOk() (*LinkedInAggregateAnalyticsTotalResponseAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetAnalytics(v LinkedInAggregateAnalyticsTotalResponseAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### GetNote

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetLastUpdated

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *LinkedInAggregateAnalyticsTotalResponse) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *LinkedInAggregateAnalyticsTotalResponse) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *LinkedInAggregateAnalyticsTotalResponse) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


