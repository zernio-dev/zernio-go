# SelectPinterestBoard200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** | Redirect URL with connection params (if provided) | [optional] 
**Account** | Pointer to [**SelectPinterestBoard200ResponseAccount**](SelectPinterestBoard200ResponseAccount.md) |  | [optional] 

## Methods

### NewSelectPinterestBoard200Response

`func NewSelectPinterestBoard200Response() *SelectPinterestBoard200Response`

NewSelectPinterestBoard200Response instantiates a new SelectPinterestBoard200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectPinterestBoard200ResponseWithDefaults

`func NewSelectPinterestBoard200ResponseWithDefaults() *SelectPinterestBoard200Response`

NewSelectPinterestBoard200ResponseWithDefaults instantiates a new SelectPinterestBoard200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *SelectPinterestBoard200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SelectPinterestBoard200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SelectPinterestBoard200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *SelectPinterestBoard200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *SelectPinterestBoard200Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *SelectPinterestBoard200Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *SelectPinterestBoard200Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *SelectPinterestBoard200Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetAccount

`func (o *SelectPinterestBoard200Response) GetAccount() SelectPinterestBoard200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *SelectPinterestBoard200Response) GetAccountOk() (*SelectPinterestBoard200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *SelectPinterestBoard200Response) SetAccount(v SelectPinterestBoard200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *SelectPinterestBoard200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


