# InboxWebhookMessageAttachmentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Attachment type (image, video, file, sticker, audio) | 
**Url** | **string** | Attachment URL (may expire for Meta platforms) | 
**Payload** | Pointer to **map[string]interface{}** | Additional attachment metadata | [optional] 

## Methods

### NewInboxWebhookMessageAttachmentsInner

`func NewInboxWebhookMessageAttachmentsInner(type_ string, url string, ) *InboxWebhookMessageAttachmentsInner`

NewInboxWebhookMessageAttachmentsInner instantiates a new InboxWebhookMessageAttachmentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxWebhookMessageAttachmentsInnerWithDefaults

`func NewInboxWebhookMessageAttachmentsInnerWithDefaults() *InboxWebhookMessageAttachmentsInner`

NewInboxWebhookMessageAttachmentsInnerWithDefaults instantiates a new InboxWebhookMessageAttachmentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InboxWebhookMessageAttachmentsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InboxWebhookMessageAttachmentsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InboxWebhookMessageAttachmentsInner) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *InboxWebhookMessageAttachmentsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *InboxWebhookMessageAttachmentsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *InboxWebhookMessageAttachmentsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPayload

`func (o *InboxWebhookMessageAttachmentsInner) GetPayload() map[string]interface{}`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *InboxWebhookMessageAttachmentsInner) GetPayloadOk() (*map[string]interface{}, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *InboxWebhookMessageAttachmentsInner) SetPayload(v map[string]interface{})`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *InboxWebhookMessageAttachmentsInner) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


