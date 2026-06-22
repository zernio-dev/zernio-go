# SendInboxMessageRequestTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Template type. Required for Instagram/Facebook generic templates; ignored on WhatsApp. | [optional] 
**Elements** | Pointer to [**[]SendInboxMessageRequestTemplateElementsInner**](SendInboxMessageRequestTemplateElementsInner.md) |  | [optional] 

## Methods

### NewSendInboxMessageRequestTemplate

`func NewSendInboxMessageRequestTemplate() *SendInboxMessageRequestTemplate`

NewSendInboxMessageRequestTemplate instantiates a new SendInboxMessageRequestTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestTemplateWithDefaults

`func NewSendInboxMessageRequestTemplateWithDefaults() *SendInboxMessageRequestTemplate`

NewSendInboxMessageRequestTemplateWithDefaults instantiates a new SendInboxMessageRequestTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendInboxMessageRequestTemplate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendInboxMessageRequestTemplate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendInboxMessageRequestTemplate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SendInboxMessageRequestTemplate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetElements

`func (o *SendInboxMessageRequestTemplate) GetElements() []SendInboxMessageRequestTemplateElementsInner`

GetElements returns the Elements field if non-nil, zero value otherwise.

### GetElementsOk

`func (o *SendInboxMessageRequestTemplate) GetElementsOk() (*[]SendInboxMessageRequestTemplateElementsInner, bool)`

GetElementsOk returns a tuple with the Elements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElements

`func (o *SendInboxMessageRequestTemplate) SetElements(v []SendInboxMessageRequestTemplateElementsInner)`

SetElements sets Elements field to given value.

### HasElements

`func (o *SendInboxMessageRequestTemplate) HasElements() bool`

HasElements returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


