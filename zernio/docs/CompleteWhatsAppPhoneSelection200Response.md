# CompleteWhatsAppPhoneSelection200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**RedirectUrl** | Pointer to **string** | Present only if redirect_url was provided in the request | [optional] 
**Account** | Pointer to [**ConnectWhatsAppCredentials200ResponseAccount**](ConnectWhatsAppCredentials200ResponseAccount.md) |  | [optional] 

## Methods

### NewCompleteWhatsAppPhoneSelection200Response

`func NewCompleteWhatsAppPhoneSelection200Response() *CompleteWhatsAppPhoneSelection200Response`

NewCompleteWhatsAppPhoneSelection200Response instantiates a new CompleteWhatsAppPhoneSelection200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCompleteWhatsAppPhoneSelection200ResponseWithDefaults

`func NewCompleteWhatsAppPhoneSelection200ResponseWithDefaults() *CompleteWhatsAppPhoneSelection200Response`

NewCompleteWhatsAppPhoneSelection200ResponseWithDefaults instantiates a new CompleteWhatsAppPhoneSelection200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CompleteWhatsAppPhoneSelection200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CompleteWhatsAppPhoneSelection200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CompleteWhatsAppPhoneSelection200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CompleteWhatsAppPhoneSelection200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetRedirectUrl

`func (o *CompleteWhatsAppPhoneSelection200Response) GetRedirectUrl() string`

GetRedirectUrl returns the RedirectUrl field if non-nil, zero value otherwise.

### GetRedirectUrlOk

`func (o *CompleteWhatsAppPhoneSelection200Response) GetRedirectUrlOk() (*string, bool)`

GetRedirectUrlOk returns a tuple with the RedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUrl

`func (o *CompleteWhatsAppPhoneSelection200Response) SetRedirectUrl(v string)`

SetRedirectUrl sets RedirectUrl field to given value.

### HasRedirectUrl

`func (o *CompleteWhatsAppPhoneSelection200Response) HasRedirectUrl() bool`

HasRedirectUrl returns a boolean if a field has been set.

### GetAccount

`func (o *CompleteWhatsAppPhoneSelection200Response) GetAccount() ConnectWhatsAppCredentials200ResponseAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CompleteWhatsAppPhoneSelection200Response) GetAccountOk() (*ConnectWhatsAppCredentials200ResponseAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CompleteWhatsAppPhoneSelection200Response) SetAccount(v ConnectWhatsAppCredentials200ResponseAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CompleteWhatsAppPhoneSelection200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


