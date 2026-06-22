# WhatsAppHeaderComponentExample

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HeaderText** | Pointer to **[]string** | Sample values for header text variables | [optional] 
**HeaderHandle** | Pointer to **[]string** | When the header format is a media type (image, video, gif, document), provide a public URL here. Zernio will download and upload it to WhatsApp on your behalf, replacing it with the internal file handle before creating the template. | [optional] 

## Methods

### NewWhatsAppHeaderComponentExample

`func NewWhatsAppHeaderComponentExample() *WhatsAppHeaderComponentExample`

NewWhatsAppHeaderComponentExample instantiates a new WhatsAppHeaderComponentExample object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppHeaderComponentExampleWithDefaults

`func NewWhatsAppHeaderComponentExampleWithDefaults() *WhatsAppHeaderComponentExample`

NewWhatsAppHeaderComponentExampleWithDefaults instantiates a new WhatsAppHeaderComponentExample object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaderText

`func (o *WhatsAppHeaderComponentExample) GetHeaderText() []string`

GetHeaderText returns the HeaderText field if non-nil, zero value otherwise.

### GetHeaderTextOk

`func (o *WhatsAppHeaderComponentExample) GetHeaderTextOk() (*[]string, bool)`

GetHeaderTextOk returns a tuple with the HeaderText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderText

`func (o *WhatsAppHeaderComponentExample) SetHeaderText(v []string)`

SetHeaderText sets HeaderText field to given value.

### HasHeaderText

`func (o *WhatsAppHeaderComponentExample) HasHeaderText() bool`

HasHeaderText returns a boolean if a field has been set.

### GetHeaderHandle

`func (o *WhatsAppHeaderComponentExample) GetHeaderHandle() []string`

GetHeaderHandle returns the HeaderHandle field if non-nil, zero value otherwise.

### GetHeaderHandleOk

`func (o *WhatsAppHeaderComponentExample) GetHeaderHandleOk() (*[]string, bool)`

GetHeaderHandleOk returns a tuple with the HeaderHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaderHandle

`func (o *WhatsAppHeaderComponentExample) SetHeaderHandle(v []string)`

SetHeaderHandle sets HeaderHandle field to given value.

### HasHeaderHandle

`func (o *WhatsAppHeaderComponentExample) HasHeaderHandle() bool`

HasHeaderHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


