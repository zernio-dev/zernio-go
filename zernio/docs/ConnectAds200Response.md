# ConnectAds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyConnected** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ScopedAdAccountIds** | Pointer to **[]string** | Echo of the persisted ad-account scope when the caller passed &#x60;adAccountId&#x60; / &#x60;adAccountIds&#x60;. Omitted when no scope is set.  | [optional] 
**AuthUrl** | Pointer to **string** | URL to redirect your user to for OAuth authorization | [optional] 
**State** | Pointer to **string** | State parameter for security (handled automatically) | [optional] 

## Methods

### NewConnectAds200Response

`func NewConnectAds200Response() *ConnectAds200Response`

NewConnectAds200Response instantiates a new ConnectAds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectAds200ResponseWithDefaults

`func NewConnectAds200ResponseWithDefaults() *ConnectAds200Response`

NewConnectAds200ResponseWithDefaults instantiates a new ConnectAds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyConnected

`func (o *ConnectAds200Response) GetAlreadyConnected() bool`

GetAlreadyConnected returns the AlreadyConnected field if non-nil, zero value otherwise.

### GetAlreadyConnectedOk

`func (o *ConnectAds200Response) GetAlreadyConnectedOk() (*bool, bool)`

GetAlreadyConnectedOk returns a tuple with the AlreadyConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyConnected

`func (o *ConnectAds200Response) SetAlreadyConnected(v bool)`

SetAlreadyConnected sets AlreadyConnected field to given value.

### HasAlreadyConnected

`func (o *ConnectAds200Response) HasAlreadyConnected() bool`

HasAlreadyConnected returns a boolean if a field has been set.

### GetAccountId

`func (o *ConnectAds200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ConnectAds200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ConnectAds200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ConnectAds200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *ConnectAds200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ConnectAds200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ConnectAds200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ConnectAds200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUsername

`func (o *ConnectAds200Response) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectAds200Response) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectAds200Response) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectAds200Response) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectAds200Response) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectAds200Response) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectAds200Response) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectAds200Response) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetScopedAdAccountIds

`func (o *ConnectAds200Response) GetScopedAdAccountIds() []string`

GetScopedAdAccountIds returns the ScopedAdAccountIds field if non-nil, zero value otherwise.

### GetScopedAdAccountIdsOk

`func (o *ConnectAds200Response) GetScopedAdAccountIdsOk() (*[]string, bool)`

GetScopedAdAccountIdsOk returns a tuple with the ScopedAdAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedAdAccountIds

`func (o *ConnectAds200Response) SetScopedAdAccountIds(v []string)`

SetScopedAdAccountIds sets ScopedAdAccountIds field to given value.

### HasScopedAdAccountIds

`func (o *ConnectAds200Response) HasScopedAdAccountIds() bool`

HasScopedAdAccountIds returns a boolean if a field has been set.

### GetAuthUrl

`func (o *ConnectAds200Response) GetAuthUrl() string`

GetAuthUrl returns the AuthUrl field if non-nil, zero value otherwise.

### GetAuthUrlOk

`func (o *ConnectAds200Response) GetAuthUrlOk() (*string, bool)`

GetAuthUrlOk returns a tuple with the AuthUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthUrl

`func (o *ConnectAds200Response) SetAuthUrl(v string)`

SetAuthUrl sets AuthUrl field to given value.

### HasAuthUrl

`func (o *ConnectAds200Response) HasAuthUrl() bool`

HasAuthUrl returns a boolean if a field has been set.

### GetState

`func (o *ConnectAds200Response) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectAds200Response) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectAds200Response) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *ConnectAds200Response) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


