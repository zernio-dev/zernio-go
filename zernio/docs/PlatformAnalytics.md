# PlatformAnalytics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PlatformPostId** | Pointer to **NullableString** | The native post ID on the platform (e.g. Instagram media ID, tweet ID) | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **NullableString** |  | [optional] 
**Analytics** | Pointer to [**NullablePostAnalytics**](PostAnalytics.md) |  | [optional] 
**SyncStatus** | Pointer to **string** | Sync state of analytics for this platform | [optional] 
**PlatformPostUrl** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** | Error details when status is failed | [optional] 

## Methods

### NewPlatformAnalytics

`func NewPlatformAnalytics() *PlatformAnalytics`

NewPlatformAnalytics instantiates a new PlatformAnalytics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformAnalyticsWithDefaults

`func NewPlatformAnalyticsWithDefaults() *PlatformAnalytics`

NewPlatformAnalyticsWithDefaults instantiates a new PlatformAnalytics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *PlatformAnalytics) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *PlatformAnalytics) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *PlatformAnalytics) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *PlatformAnalytics) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetStatus

`func (o *PlatformAnalytics) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PlatformAnalytics) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PlatformAnalytics) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PlatformAnalytics) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *PlatformAnalytics) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *PlatformAnalytics) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *PlatformAnalytics) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *PlatformAnalytics) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### SetPlatformPostIdNil

`func (o *PlatformAnalytics) SetPlatformPostIdNil(b bool)`

 SetPlatformPostIdNil sets the value for PlatformPostId to be an explicit nil

### UnsetPlatformPostId
`func (o *PlatformAnalytics) UnsetPlatformPostId()`

UnsetPlatformPostId ensures that no value is present for PlatformPostId, not even an explicit nil
### GetAccountId

`func (o *PlatformAnalytics) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *PlatformAnalytics) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *PlatformAnalytics) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *PlatformAnalytics) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUsername

`func (o *PlatformAnalytics) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *PlatformAnalytics) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *PlatformAnalytics) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *PlatformAnalytics) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### SetAccountUsernameNil

`func (o *PlatformAnalytics) SetAccountUsernameNil(b bool)`

 SetAccountUsernameNil sets the value for AccountUsername to be an explicit nil

### UnsetAccountUsername
`func (o *PlatformAnalytics) UnsetAccountUsername()`

UnsetAccountUsername ensures that no value is present for AccountUsername, not even an explicit nil
### GetAnalytics

`func (o *PlatformAnalytics) GetAnalytics() PostAnalytics`

GetAnalytics returns the Analytics field if non-nil, zero value otherwise.

### GetAnalyticsOk

`func (o *PlatformAnalytics) GetAnalyticsOk() (*PostAnalytics, bool)`

GetAnalyticsOk returns a tuple with the Analytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnalytics

`func (o *PlatformAnalytics) SetAnalytics(v PostAnalytics)`

SetAnalytics sets Analytics field to given value.

### HasAnalytics

`func (o *PlatformAnalytics) HasAnalytics() bool`

HasAnalytics returns a boolean if a field has been set.

### SetAnalyticsNil

`func (o *PlatformAnalytics) SetAnalyticsNil(b bool)`

 SetAnalyticsNil sets the value for Analytics to be an explicit nil

### UnsetAnalytics
`func (o *PlatformAnalytics) UnsetAnalytics()`

UnsetAnalytics ensures that no value is present for Analytics, not even an explicit nil
### GetSyncStatus

`func (o *PlatformAnalytics) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *PlatformAnalytics) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *PlatformAnalytics) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *PlatformAnalytics) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.

### GetPlatformPostUrl

`func (o *PlatformAnalytics) GetPlatformPostUrl() string`

GetPlatformPostUrl returns the PlatformPostUrl field if non-nil, zero value otherwise.

### GetPlatformPostUrlOk

`func (o *PlatformAnalytics) GetPlatformPostUrlOk() (*string, bool)`

GetPlatformPostUrlOk returns a tuple with the PlatformPostUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostUrl

`func (o *PlatformAnalytics) SetPlatformPostUrl(v string)`

SetPlatformPostUrl sets PlatformPostUrl field to given value.

### HasPlatformPostUrl

`func (o *PlatformAnalytics) HasPlatformPostUrl() bool`

HasPlatformPostUrl returns a boolean if a field has been set.

### SetPlatformPostUrlNil

`func (o *PlatformAnalytics) SetPlatformPostUrlNil(b bool)`

 SetPlatformPostUrlNil sets the value for PlatformPostUrl to be an explicit nil

### UnsetPlatformPostUrl
`func (o *PlatformAnalytics) UnsetPlatformPostUrl()`

UnsetPlatformPostUrl ensures that no value is present for PlatformPostUrl, not even an explicit nil
### GetErrorMessage

`func (o *PlatformAnalytics) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *PlatformAnalytics) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *PlatformAnalytics) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *PlatformAnalytics) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *PlatformAnalytics) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *PlatformAnalytics) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


