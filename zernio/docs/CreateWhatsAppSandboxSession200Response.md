# CreateWhatsAppSandboxSession200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**WhatsAppSandboxSession**](WhatsAppSandboxSession.md) |  | [optional] 
**SandboxNumber** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateWhatsAppSandboxSession200Response

`func NewCreateWhatsAppSandboxSession200Response() *CreateWhatsAppSandboxSession200Response`

NewCreateWhatsAppSandboxSession200Response instantiates a new CreateWhatsAppSandboxSession200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhatsAppSandboxSession200ResponseWithDefaults

`func NewCreateWhatsAppSandboxSession200ResponseWithDefaults() *CreateWhatsAppSandboxSession200Response`

NewCreateWhatsAppSandboxSession200ResponseWithDefaults instantiates a new CreateWhatsAppSandboxSession200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *CreateWhatsAppSandboxSession200Response) GetSession() WhatsAppSandboxSession`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *CreateWhatsAppSandboxSession200Response) GetSessionOk() (*WhatsAppSandboxSession, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *CreateWhatsAppSandboxSession200Response) SetSession(v WhatsAppSandboxSession)`

SetSession sets Session field to given value.

### HasSession

`func (o *CreateWhatsAppSandboxSession200Response) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetSandboxNumber

`func (o *CreateWhatsAppSandboxSession200Response) GetSandboxNumber() string`

GetSandboxNumber returns the SandboxNumber field if non-nil, zero value otherwise.

### GetSandboxNumberOk

`func (o *CreateWhatsAppSandboxSession200Response) GetSandboxNumberOk() (*string, bool)`

GetSandboxNumberOk returns a tuple with the SandboxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxNumber

`func (o *CreateWhatsAppSandboxSession200Response) SetSandboxNumber(v string)`

SetSandboxNumber sets SandboxNumber field to given value.

### HasSandboxNumber

`func (o *CreateWhatsAppSandboxSession200Response) HasSandboxNumber() bool`

HasSandboxNumber returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


