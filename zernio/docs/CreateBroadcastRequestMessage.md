# CreateBroadcastRequestMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to [**[]CreateBroadcastRequestMessageAttachmentsInner**](CreateBroadcastRequestMessageAttachmentsInner.md) |  | [optional] 

## Methods

### NewCreateBroadcastRequestMessage

`func NewCreateBroadcastRequestMessage() *CreateBroadcastRequestMessage`

NewCreateBroadcastRequestMessage instantiates a new CreateBroadcastRequestMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBroadcastRequestMessageWithDefaults

`func NewCreateBroadcastRequestMessageWithDefaults() *CreateBroadcastRequestMessage`

NewCreateBroadcastRequestMessageWithDefaults instantiates a new CreateBroadcastRequestMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *CreateBroadcastRequestMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CreateBroadcastRequestMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CreateBroadcastRequestMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CreateBroadcastRequestMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetAttachments

`func (o *CreateBroadcastRequestMessage) GetAttachments() []CreateBroadcastRequestMessageAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *CreateBroadcastRequestMessage) GetAttachmentsOk() (*[]CreateBroadcastRequestMessageAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *CreateBroadcastRequestMessage) SetAttachments(v []CreateBroadcastRequestMessageAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *CreateBroadcastRequestMessage) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


