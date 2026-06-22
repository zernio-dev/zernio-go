# TriggerWorkflowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**To** | Pointer to **string** | Recipient phone (WhatsApp only) | [optional] 
**ConversationId** | Pointer to **string** | An existing conversation to run in (required for non-WhatsApp workflows) | [optional] 
**Text** | Pointer to **string** | Simulated inbound text, seeded as the run&#39;s lastMessage variable | [optional] 

## Methods

### NewTriggerWorkflowRequest

`func NewTriggerWorkflowRequest() *TriggerWorkflowRequest`

NewTriggerWorkflowRequest instantiates a new TriggerWorkflowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTriggerWorkflowRequestWithDefaults

`func NewTriggerWorkflowRequestWithDefaults() *TriggerWorkflowRequest`

NewTriggerWorkflowRequestWithDefaults instantiates a new TriggerWorkflowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTo

`func (o *TriggerWorkflowRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *TriggerWorkflowRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *TriggerWorkflowRequest) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *TriggerWorkflowRequest) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetConversationId

`func (o *TriggerWorkflowRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *TriggerWorkflowRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *TriggerWorkflowRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *TriggerWorkflowRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetText

`func (o *TriggerWorkflowRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *TriggerWorkflowRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *TriggerWorkflowRequest) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *TriggerWorkflowRequest) HasText() bool`

HasText returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


