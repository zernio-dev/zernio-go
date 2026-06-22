# ConnectAds200ResponseOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AlreadyConnected** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**ScopedAdAccountIds** | Pointer to **[]string** | Echo of the persisted ad-account scope when the caller passed &#x60;adAccountId&#x60; / &#x60;adAccountIds&#x60;. Omitted when no scope is set.  | [optional] 

## Methods

### NewConnectAds200ResponseOneOf

`func NewConnectAds200ResponseOneOf() *ConnectAds200ResponseOneOf`

NewConnectAds200ResponseOneOf instantiates a new ConnectAds200ResponseOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectAds200ResponseOneOfWithDefaults

`func NewConnectAds200ResponseOneOfWithDefaults() *ConnectAds200ResponseOneOf`

NewConnectAds200ResponseOneOfWithDefaults instantiates a new ConnectAds200ResponseOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlreadyConnected

`func (o *ConnectAds200ResponseOneOf) GetAlreadyConnected() bool`

GetAlreadyConnected returns the AlreadyConnected field if non-nil, zero value otherwise.

### GetAlreadyConnectedOk

`func (o *ConnectAds200ResponseOneOf) GetAlreadyConnectedOk() (*bool, bool)`

GetAlreadyConnectedOk returns a tuple with the AlreadyConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlreadyConnected

`func (o *ConnectAds200ResponseOneOf) SetAlreadyConnected(v bool)`

SetAlreadyConnected sets AlreadyConnected field to given value.

### HasAlreadyConnected

`func (o *ConnectAds200ResponseOneOf) HasAlreadyConnected() bool`

HasAlreadyConnected returns a boolean if a field has been set.

### GetAccountId

`func (o *ConnectAds200ResponseOneOf) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ConnectAds200ResponseOneOf) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ConnectAds200ResponseOneOf) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ConnectAds200ResponseOneOf) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatform

`func (o *ConnectAds200ResponseOneOf) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ConnectAds200ResponseOneOf) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ConnectAds200ResponseOneOf) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ConnectAds200ResponseOneOf) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetUsername

`func (o *ConnectAds200ResponseOneOf) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ConnectAds200ResponseOneOf) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ConnectAds200ResponseOneOf) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ConnectAds200ResponseOneOf) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *ConnectAds200ResponseOneOf) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *ConnectAds200ResponseOneOf) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *ConnectAds200ResponseOneOf) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *ConnectAds200ResponseOneOf) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetScopedAdAccountIds

`func (o *ConnectAds200ResponseOneOf) GetScopedAdAccountIds() []string`

GetScopedAdAccountIds returns the ScopedAdAccountIds field if non-nil, zero value otherwise.

### GetScopedAdAccountIdsOk

`func (o *ConnectAds200ResponseOneOf) GetScopedAdAccountIdsOk() (*[]string, bool)`

GetScopedAdAccountIdsOk returns a tuple with the ScopedAdAccountIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopedAdAccountIds

`func (o *ConnectAds200ResponseOneOf) SetScopedAdAccountIds(v []string)`

SetScopedAdAccountIds sets ScopedAdAccountIds field to given value.

### HasScopedAdAccountIds

`func (o *ConnectAds200ResponseOneOf) HasScopedAdAccountIds() bool`

HasScopedAdAccountIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


