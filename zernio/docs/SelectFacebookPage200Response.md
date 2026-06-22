# SelectFacebookPage200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect URL if custom redirect_url was provided | [optional] 
**Account** | Pointer to [**SelectFacebookPage200ResponseAccount**](SelectFacebookPage200ResponseAccount.md) |  | [optional] 

## Methods

### NewSelectFacebookPage200Response

`func NewSelectFacebookPage200Response() *SelectFacebookPage200Response`

NewSelectFacebookPage200Response instantiates a new SelectFacebookPage200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectFacebookPage200ResponseWithDefaults

`func NewSelectFacebookPage200ResponseWithDefaults() *SelectFacebookPage200Response`

NewSelectFacebookPage200ResponseWithDefaults instantiates a new SelectFacebookPage200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SelectFacebookPage200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SelectFacebookPage200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SelectFacebookPage200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SelectFacebookPage200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectFacebookPage200Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectFacebookPage200Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectFacebookPage200Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectFacebookPage200Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetAccount

`func (o *SelectFacebookPage200Response) GetAccount() SelectFacebookPage200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SelectFacebookPage200Response) GetAccountOk() (*SelectFacebookPage200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SelectFacebookPage200Response) SetAccount(v SelectFacebookPage200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SelectFacebookPage200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


