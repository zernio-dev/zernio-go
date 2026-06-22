# ConnectBlueskyCredentialsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | **string** | Your Bluesky handle (e.g. user.bsky.social) or email address | 
**AppPassword** | **string** | App password generated from Bluesky Settings &gt; App Passwords | 
**State** | **string** | Required state formatted as {userId}-{profileId}. Get userId from GET /v1/users and profileId from GET /v1/profiles. | 
**RedirectUri** | Pointer to **string** | Optional URL to redirect to after successful connection | [optional] 

## Methods

### NewConnectBlueskyCredentialsRequest

`func NewConnectBlueskyCredentialsRequest(identifier string, appPassword string, state string, ) *ConnectBlueskyCredentialsRequest`

NewConnectBlueskyCredentialsRequest instantiates a new ConnectBlueskyCredentialsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectBlueskyCredentialsRequestWithDefaults

`func NewConnectBlueskyCredentialsRequestWithDefaults() *ConnectBlueskyCredentialsRequest`

NewConnectBlueskyCredentialsRequestWithDefaults instantiates a new ConnectBlueskyCredentialsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *ConnectBlueskyCredentialsRequest) GetIdentifier() string`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *ConnectBlueskyCredentialsRequest) GetIdentifierOk() (*string, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *ConnectBlueskyCredentialsRequest) SetIdentifier(v string)`

SetIdentifier sets Identifier field to given value.


### GetAppPassword

`func (o *ConnectBlueskyCredentialsRequest) GetAppPassword() string`

GetAppPassword returns the AppPassword field if non-nil, zero value otherwise.

### GetAppPasswordOk

`func (o *ConnectBlueskyCredentialsRequest) GetAppPasswordOk() (*string, bool)`

GetAppPasswordOk returns a tuple with the AppPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppPassword

`func (o *ConnectBlueskyCredentialsRequest) SetAppPassword(v string)`

SetAppPassword sets AppPassword field to given value.


### GetState

`func (o *ConnectBlueskyCredentialsRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ConnectBlueskyCredentialsRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ConnectBlueskyCredentialsRequest) SetState(v string)`

SetState sets State field to given value.


### GetRedirectUri

`func (o *ConnectBlueskyCredentialsRequest) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *ConnectBlueskyCredentialsRequest) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *ConnectBlueskyCredentialsRequest) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.

### HasRedirectUri

`func (o *ConnectBlueskyCredentialsRequest) HasRedirectUri() bool`

HasRedirectUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


