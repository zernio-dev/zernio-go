# SelectGoogleBusinessLocation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect URL if custom redirect_url was provided | [optional] 
**Account** | Pointer to [**SelectGoogleBusinessLocation200ResponseAccount**](SelectGoogleBusinessLocation200ResponseAccount.md) |  | [optional] 

## Methods

### NewSelectGoogleBusinessLocation200Response

`func NewSelectGoogleBusinessLocation200Response() *SelectGoogleBusinessLocation200Response`

NewSelectGoogleBusinessLocation200Response instantiates a new SelectGoogleBusinessLocation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectGoogleBusinessLocation200ResponseWithDefaults

`func NewSelectGoogleBusinessLocation200ResponseWithDefaults() *SelectGoogleBusinessLocation200Response`

NewSelectGoogleBusinessLocation200ResponseWithDefaults instantiates a new SelectGoogleBusinessLocation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SelectGoogleBusinessLocation200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SelectGoogleBusinessLocation200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SelectGoogleBusinessLocation200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SelectGoogleBusinessLocation200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectGoogleBusinessLocation200Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectGoogleBusinessLocation200Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectGoogleBusinessLocation200Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectGoogleBusinessLocation200Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetAccount

`func (o *SelectGoogleBusinessLocation200Response) GetAccount() SelectGoogleBusinessLocation200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SelectGoogleBusinessLocation200Response) GetAccountOk() (*SelectGoogleBusinessLocation200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SelectGoogleBusinessLocation200Response) SetAccount(v SelectGoogleBusinessLocation200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SelectGoogleBusinessLocation200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


