# SelectPinterestBoardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Your Zernio profile ID | 
**BoardId** | **string** | The Pinterest Board ID selected by the user | 
**BoardName** | Pointer to **string** | The board name (for display purposes) | [optional] 
**TempToken** | **string** | Temporary Pinterest access token from OAuth | 
**UserProfile** | Pointer to **map[string]interface{}** | User profile data from OAuth redirect | [optional] 
**RefreshToken** | Pointer to **string** | Pinterest refresh token (if available) | [optional] 
**ExpiresIn** | Pointer to **int32** | Token expiration time in seconds | [optional] 
**RedirectUrl** | Pointer to **string** | Custom redirect URL after connection completes | [optional] 

## Methods

### NewSelectPinterestBoardRequest

`func NewSelectPinterestBoardRequest(profileId string, boardId string, tempToken string, ) *SelectPinterestBoardRequest`

NewSelectPinterestBoardRequest instantiates a new SelectPinterestBoardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPinterestBoardRequestWithDefaults

`func NewSelectPinterestBoardRequestWithDefaults() *SelectPinterestBoardRequest`

NewSelectPinterestBoardRequestWithDefaults instantiates a new SelectPinterestBoardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SelectPinterestBoardRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SelectPinterestBoardRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SelectPinterestBoardRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetBoardId

`func (o *SelectPinterestBoardRequest) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *SelectPinterestBoardRequest) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *SelectPinterestBoardRequest) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.


### GetBoardName

`func (o *SelectPinterestBoardRequest) GetBoardName() string`

GetBoardName returns the BoardName field if non-nil, zero value otherwise.

### GetBoardNameOk

`func (o *SelectPinterestBoardRequest) GetBoardNameOk() (*string, bool)`

GetBoardNameOk returns a tuple with the BoardName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardName

`func (o *SelectPinterestBoardRequest) SetBoardName(v string)`

SetBoardName sets BoardName field to given value.

### HasBoardName

`func (o *SelectPinterestBoardRequest) HasBoardName() bool`

HasBoardName returns a boolean if a field has been set.

### GetTempToken

`func (o *SelectPinterestBoardRequest) GetTempToken() string`

GetTempToken returns the TempToken field if non-nil, zero value otherwise.

### GetTempTokenOk

`func (o *SelectPinterestBoardRequest) GetTempTokenOk() (*string, bool)`

GetTempTokenOk returns a tuple with the TempToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTempToken

`func (o *SelectPinterestBoardRequest) SetTempToken(v string)`

SetTempToken sets TempToken field to given value.


### GetUserProfile

`func (o *SelectPinterestBoardRequest) GetUserProfile() map[string]interface{}`

GetUserProfile returns the UserProfile field if non-nil, zero value otherwise.

### GetUserProfileOk

`func (o *SelectPinterestBoardRequest) GetUserProfileOk() (*map[string]interface{}, bool)`

GetUserProfileOk returns a tuple with the UserProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserProfile

`func (o *SelectPinterestBoardRequest) SetUserProfile(v map[string]interface{})`

SetUserProfile sets UserProfile field to given value.

### HasUserProfile

`func (o *SelectPinterestBoardRequest) HasUserProfile() bool`

HasUserProfile returns a boolean if a field has been set.

### GetRefreshToken

`func (o *SelectPinterestBoardRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *SelectPinterestBoardRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *SelectPinterestBoardRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.

### HasRefreshToken

`func (o *SelectPinterestBoardRequest) HasRefreshToken() bool`

HasRefreshToken returns a boolean if a field has been set.

### GetExpiresIn

`func (o *SelectPinterestBoardRequest) GetExpiresIn() int32`

GetExpiresIn returns the ExpiresIn field if non-nil, zero value otherwise.

### GetExpiresInOk

`func (o *SelectPinterestBoardRequest) GetExpiresInOk() (*int32, bool)`

GetExpiresInOk returns a tuple with the ExpiresIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresIn

`func (o *SelectPinterestBoardRequest) SetExpiresIn(v int32)`

SetExpiresIn sets ExpiresIn field to given value.

### HasExpiresIn

`func (o *SelectPinterestBoardRequest) HasExpiresIn() bool`

HasExpiresIn returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectPinterestBoardRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectPinterestBoardRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectPinterestBoardRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectPinterestBoardRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


