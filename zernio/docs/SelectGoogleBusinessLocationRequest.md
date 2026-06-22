# SelectGoogleBusinessLocationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Profile ID from your connection flow | 
**LocationId** | **string** | The Google Business location ID selected by the user | 
**AccountId** | Pointer to **string** | Optional but recommended. The Google Business Account resource name (\&quot;accounts/123\&quot;) that owns the selected location (returned per-location by GET /v1/connect/googlebusiness/locations). When provided, the location is resolved directly instead of by enumerating the account, which is required for accounts that own many locations. Omit only for small accounts.  | [optional] 
**PendingDataToken** | **string** | Token from the OAuth callback redirect (pendingDataToken query param). Tokens and profile data are retrieved server-side from this token. | 
**RedirectUrl** | Pointer to **string** | Optional custom redirect URL to return to after selection | [optional] 

## Methods

### NewSelectGoogleBusinessLocationRequest

`func NewSelectGoogleBusinessLocationRequest(profileId string, locationId string, pendingDataToken string, ) *SelectGoogleBusinessLocationRequest`

NewSelectGoogleBusinessLocationRequest instantiates a new SelectGoogleBusinessLocationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectGoogleBusinessLocationRequestWithDefaults

`func NewSelectGoogleBusinessLocationRequestWithDefaults() *SelectGoogleBusinessLocationRequest`

NewSelectGoogleBusinessLocationRequestWithDefaults instantiates a new SelectGoogleBusinessLocationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *SelectGoogleBusinessLocationRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *SelectGoogleBusinessLocationRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *SelectGoogleBusinessLocationRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetLocationId

`func (o *SelectGoogleBusinessLocationRequest) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *SelectGoogleBusinessLocationRequest) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *SelectGoogleBusinessLocationRequest) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.


### GetAccountId

`func (o *SelectGoogleBusinessLocationRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SelectGoogleBusinessLocationRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SelectGoogleBusinessLocationRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *SelectGoogleBusinessLocationRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPendingDataToken

`func (o *SelectGoogleBusinessLocationRequest) GetPendingDataToken() string`

GetPendingDataToken returns the PendingDataToken field if non-nil, zero value otherwise.

### GetPendingDataTokenOk

`func (o *SelectGoogleBusinessLocationRequest) GetPendingDataTokenOk() (*string, bool)`

GetPendingDataTokenOk returns a tuple with the PendingDataToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDataToken

`func (o *SelectGoogleBusinessLocationRequest) SetPendingDataToken(v string)`

SetPendingDataToken sets PendingDataToken field to given value.


### GetRedirectUrl

`func (o *SelectGoogleBusinessLocationRequest) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectGoogleBusinessLocationRequest) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectGoogleBusinessLocationRequest) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectGoogleBusinessLocationRequest) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


