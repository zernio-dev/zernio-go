# ListWhatsAppSandboxSessions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sessions** | Pointer to [**[]WhatsAppSandboxSession**](WhatsAppSandboxSession.md) |  | [optional] 
**SandboxNumber** | Pointer to **NullableString** | The shared sandbox phone number in E.164 form. | [optional] 

## Methods

### NewListWhatsAppSandboxSessions200Response

`func NewListWhatsAppSandboxSessions200Response() *ListWhatsAppSandboxSessions200Response`

NewListWhatsAppSandboxSessions200Response instantiates a new ListWhatsAppSandboxSessions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppSandboxSessions200ResponseWithDefaults

`func NewListWhatsAppSandboxSessions200ResponseWithDefaults() *ListWhatsAppSandboxSessions200Response`

NewListWhatsAppSandboxSessions200ResponseWithDefaults instantiates a new ListWhatsAppSandboxSessions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSessions

`func (o *ListWhatsAppSandboxSessions200Response) GetSessions() []WhatsAppSandboxSession`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *ListWhatsAppSandboxSessions200Response) GetSessionsOk() (*[]WhatsAppSandboxSession, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *ListWhatsAppSandboxSessions200Response) SetSessions(v []WhatsAppSandboxSession)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *ListWhatsAppSandboxSessions200Response) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetSandboxNumber

`func (o *ListWhatsAppSandboxSessions200Response) GetSandboxNumber() string`

GetSandboxNumber returns the SandboxNumber field if non-nil, zero value otherwise.

### GetSandboxNumberOk

`func (o *ListWhatsAppSandboxSessions200Response) GetSandboxNumberOk() (*string, bool)`

GetSandboxNumberOk returns a tuple with the SandboxNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxNumber

`func (o *ListWhatsAppSandboxSessions200Response) SetSandboxNumber(v string)`

SetSandboxNumber sets SandboxNumber field to given value.

### HasSandboxNumber

`func (o *ListWhatsAppSandboxSessions200Response) HasSandboxNumber() bool`

HasSandboxNumber returns a boolean if a field has been set.

### SetSandboxNumberNil

`func (o *ListWhatsAppSandboxSessions200Response) SetSandboxNumberNil(b bool)`

 SetSandboxNumberNil sets the value for SandboxNumber to be an explicit nil

### UnsetSandboxNumber
`func (o *ListWhatsAppSandboxSessions200Response) UnsetSandboxNumber()`

UnsetSandboxNumber ensures that no value is present for SandboxNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


