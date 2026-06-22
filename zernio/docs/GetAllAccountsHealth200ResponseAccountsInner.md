# GetAllAccountsHealth200ResponseAccountsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ProfileId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CanPost** | Pointer to **bool** |  | [optional] 
**CanFetchAnalytics** | Pointer to **bool** |  | [optional] 
**TokenValid** | Pointer to **bool** |  | [optional] 
**TokenExpiresAt** | Pointer to **time.Time** |  | [optional] 
**NeedsReconnect** | Pointer to **bool** |  | [optional] 
**Issues** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGetAllAccountsHealth200ResponseAccountsInner

`func NewGetAllAccountsHealth200ResponseAccountsInner() *GetAllAccountsHealth200ResponseAccountsInner`

NewGetAllAccountsHealth200ResponseAccountsInner instantiates a new GetAllAccountsHealth200ResponseAccountsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllAccountsHealth200ResponseAccountsInnerWithDefaults

`func NewGetAllAccountsHealth200ResponseAccountsInnerWithDefaults() *GetAllAccountsHealth200ResponseAccountsInner`

NewGetAllAccountsHealth200ResponseAccountsInnerWithDefaults instantiates a new GetAllAccountsHealth200ResponseAccountsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUsername

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetProfileId

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.

### HasProfileId

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasProfileId() bool`

HasProfileId returns a boolean if a field has been set.

### GetStatus

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCanPost

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetCanPost() bool`

GetCanPost returns the CanPost field if non-nil, zero value otherwise.

### GetCanPostOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetCanPostOk() (*bool, bool)`

GetCanPostOk returns a tuple with the CanPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanPost

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetCanPost(v bool)`

SetCanPost sets CanPost field to given value.

### HasCanPost

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasCanPost() bool`

HasCanPost returns a boolean if a field has been set.

### GetCanFetchAnalytics

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetCanFetchAnalytics() bool`

GetCanFetchAnalytics returns the CanFetchAnalytics field if non-nil, zero value otherwise.

### GetCanFetchAnalyticsOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetCanFetchAnalyticsOk() (*bool, bool)`

GetCanFetchAnalyticsOk returns a tuple with the CanFetchAnalytics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanFetchAnalytics

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetCanFetchAnalytics(v bool)`

SetCanFetchAnalytics sets CanFetchAnalytics field to given value.

### HasCanFetchAnalytics

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasCanFetchAnalytics() bool`

HasCanFetchAnalytics returns a boolean if a field has been set.

### GetTokenValid

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetTokenValid() bool`

GetTokenValid returns the TokenValid field if non-nil, zero value otherwise.

### GetTokenValidOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetTokenValidOk() (*bool, bool)`

GetTokenValidOk returns a tuple with the TokenValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenValid

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetTokenValid(v bool)`

SetTokenValid sets TokenValid field to given value.

### HasTokenValid

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasTokenValid() bool`

HasTokenValid returns a boolean if a field has been set.

### GetTokenExpiresAt

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetTokenExpiresAt() time.Time`

GetTokenExpiresAt returns the TokenExpiresAt field if non-nil, zero value otherwise.

### GetTokenExpiresAtOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetTokenExpiresAtOk() (*time.Time, bool)`

GetTokenExpiresAtOk returns a tuple with the TokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiresAt

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetTokenExpiresAt(v time.Time)`

SetTokenExpiresAt sets TokenExpiresAt field to given value.

### HasTokenExpiresAt

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasTokenExpiresAt() bool`

HasTokenExpiresAt returns a boolean if a field has been set.

### GetNeedsReconnect

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetNeedsReconnect() bool`

GetNeedsReconnect returns the NeedsReconnect field if non-nil, zero value otherwise.

### GetNeedsReconnectOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetNeedsReconnectOk() (*bool, bool)`

GetNeedsReconnectOk returns a tuple with the NeedsReconnect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsReconnect

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetNeedsReconnect(v bool)`

SetNeedsReconnect sets NeedsReconnect field to given value.

### HasNeedsReconnect

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasNeedsReconnect() bool`

HasNeedsReconnect returns a boolean if a field has been set.

### GetIssues

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetAllAccountsHealth200ResponseAccountsInner) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetAllAccountsHealth200ResponseAccountsInner) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *GetAllAccountsHealth200ResponseAccountsInner) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


