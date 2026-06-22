# GetAccountHealth200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | Overall health status | [optional] 
**TokenStatus** | Pointer to [**GetAccountHealth200ResponseTokenStatus**](GetAccountHealth200ResponseTokenStatus.md) |  | [optional] 
**Permissions** | Pointer to [**GetAccountHealth200ResponsePermissions**](GetAccountHealth200ResponsePermissions.md) |  | [optional] 
**Issues** | Pointer to **[]string** | List of issues found | [optional] 
**Recommendations** | Pointer to **[]string** | Actionable recommendations to fix issues | [optional] 

## Methods

### NewGetAccountHealth200Response

`func NewGetAccountHealth200Response() *GetAccountHealth200Response`

NewGetAccountHealth200Response instantiates a new GetAccountHealth200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountHealth200ResponseWithDefaults

`func NewGetAccountHealth200ResponseWithDefaults() *GetAccountHealth200Response`

NewGetAccountHealth200ResponseWithDefaults instantiates a new GetAccountHealth200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetAccountHealth200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAccountHealth200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAccountHealth200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetAccountHealth200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *GetAccountHealth200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAccountHealth200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAccountHealth200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetAccountHealth200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUsername

`func (o *GetAccountHealth200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetAccountHealth200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetAccountHealth200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetAccountHealth200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetAccountHealth200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetAccountHealth200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetAccountHealth200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetAccountHealth200Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetStatus

`func (o *GetAccountHealth200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetAccountHealth200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetAccountHealth200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetAccountHealth200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTokenStatus

`func (o *GetAccountHealth200Response) GetTokenStatus() GetAccountHealth200ResponseTokenStatus`

GetTokenStatus returns the TokenStatus field if non-nil, zero value otherwise.

### GetTokenStatusOk

`func (o *GetAccountHealth200Response) GetTokenStatusOk() (*GetAccountHealth200ResponseTokenStatus, bool)`

GetTokenStatusOk returns a tuple with the TokenStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenStatus

`func (o *GetAccountHealth200Response) SetTokenStatus(v GetAccountHealth200ResponseTokenStatus)`

SetTokenStatus sets TokenStatus field to given value.

### HasTokenStatus

`func (o *GetAccountHealth200Response) HasTokenStatus() bool`

HasTokenStatus returns a boolean if a field has been set.

### GetPermissions

`func (o *GetAccountHealth200Response) GetPermissions() GetAccountHealth200ResponsePermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetAccountHealth200Response) GetPermissionsOk() (*GetAccountHealth200ResponsePermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetAccountHealth200Response) SetPermissions(v GetAccountHealth200ResponsePermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetAccountHealth200Response) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetIssues

`func (o *GetAccountHealth200Response) GetIssues() []string`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *GetAccountHealth200Response) GetIssuesOk() (*[]string, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *GetAccountHealth200Response) SetIssues(v []string)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *GetAccountHealth200Response) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetRecommendations

`func (o *GetAccountHealth200Response) GetRecommendations() []string`

GetRecommendations returns the Recommendations field if non-nil, zero value otherwise.

### GetRecommendationsOk

`func (o *GetAccountHealth200Response) GetRecommendationsOk() (*[]string, bool)`

GetRecommendationsOk returns a tuple with the Recommendations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendations

`func (o *GetAccountHealth200Response) SetRecommendations(v []string)`

SetRecommendations sets Recommendations field to given value.

### HasRecommendations

`func (o *GetAccountHealth200Response) HasRecommendations() bool`

HasRecommendations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


