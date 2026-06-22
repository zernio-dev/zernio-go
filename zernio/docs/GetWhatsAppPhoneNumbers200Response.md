# GetWhatsAppPhoneNumbers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Numbers** | Pointer to [**[]GetWhatsAppPhoneNumbers200ResponseNumbersInner**](GetWhatsAppPhoneNumbers200ResponseNumbersInner.md) |  | [optional] 
**Connected** | Pointer to [**[]GetWhatsAppPhoneNumbers200ResponseConnectedInner**](GetWhatsAppPhoneNumbers200ResponseConnectedInner.md) | Connected (bring-your-own) WhatsApp numbers — your own WABA numbers linked via Embedded Signup. Not provisioned or billed by Zernio, so they are not in &#x60;numbers&#x60;; &#x60;accountId&#x60; is the social-account id used by the messaging and inbox endpoints. Included only on the default and &#x60;status&#x3D;active&#x60; views.  | [optional] 
**Sandbox** | Pointer to [**GetWhatsAppPhoneNumbers200ResponseSandbox**](GetWhatsAppPhoneNumbers200ResponseSandbox.md) |  | [optional] 

## Methods

### NewGetWhatsAppPhoneNumbers200Response

`func NewGetWhatsAppPhoneNumbers200Response() *GetWhatsAppPhoneNumbers200Response`

NewGetWhatsAppPhoneNumbers200Response instantiates a new GetWhatsAppPhoneNumbers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppPhoneNumbers200ResponseWithDefaults

`func NewGetWhatsAppPhoneNumbers200ResponseWithDefaults() *GetWhatsAppPhoneNumbers200Response`

NewGetWhatsAppPhoneNumbers200ResponseWithDefaults instantiates a new GetWhatsAppPhoneNumbers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumbers

`func (o *GetWhatsAppPhoneNumbers200Response) GetNumbers() []GetWhatsAppPhoneNumbers200ResponseNumbersInner`

GetNumbers returns the Numbers field if non-nil, zero value otherwise.

### GetNumbersOk

`func (o *GetWhatsAppPhoneNumbers200Response) GetNumbersOk() (*[]GetWhatsAppPhoneNumbers200ResponseNumbersInner, bool)`

GetNumbersOk returns a tuple with the Numbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumbers

`func (o *GetWhatsAppPhoneNumbers200Response) SetNumbers(v []GetWhatsAppPhoneNumbers200ResponseNumbersInner)`

SetNumbers sets Numbers field to given value.

### HasNumbers

`func (o *GetWhatsAppPhoneNumbers200Response) HasNumbers() bool`

HasNumbers returns a boolean if a field has been set.

### GetConnected

`func (o *GetWhatsAppPhoneNumbers200Response) GetConnected() []GetWhatsAppPhoneNumbers200ResponseConnectedInner`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *GetWhatsAppPhoneNumbers200Response) GetConnectedOk() (*[]GetWhatsAppPhoneNumbers200ResponseConnectedInner, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *GetWhatsAppPhoneNumbers200Response) SetConnected(v []GetWhatsAppPhoneNumbers200ResponseConnectedInner)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *GetWhatsAppPhoneNumbers200Response) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetSandbox

`func (o *GetWhatsAppPhoneNumbers200Response) GetSandbox() GetWhatsAppPhoneNumbers200ResponseSandbox`

GetSandbox returns the Sandbox field if non-nil, zero value otherwise.

### GetSandboxOk

`func (o *GetWhatsAppPhoneNumbers200Response) GetSandboxOk() (*GetWhatsAppPhoneNumbers200ResponseSandbox, bool)`

GetSandboxOk returns a tuple with the Sandbox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandbox

`func (o *GetWhatsAppPhoneNumbers200Response) SetSandbox(v GetWhatsAppPhoneNumbers200ResponseSandbox)`

SetSandbox sets Sandbox field to given value.

### HasSandbox

`func (o *GetWhatsAppPhoneNumbers200Response) HasSandbox() bool`

HasSandbox returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


