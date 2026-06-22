# ListWhatsAppFlowResponses200ResponseResponsesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Message ID | [optional] 
**ReceivedAt** | Pointer to **time.Time** |  | [optional] 
**From** | Pointer to **NullableString** | Sender wa_id / phone | [optional] 
**SenderName** | Pointer to **NullableString** |  | [optional] 
**ConversationId** | Pointer to **NullableString** |  | [optional] 
**FlowToken** | Pointer to **NullableString** |  | [optional] 
**Data** | Pointer to **map[string]interface{}** | Submitted field values (flow_token removed) | [optional] 
**Raw** | Pointer to **NullableString** | Raw response_json string | [optional] 

## Methods

### NewListWhatsAppFlowResponses200ResponseResponsesInner

`func NewListWhatsAppFlowResponses200ResponseResponsesInner() *ListWhatsAppFlowResponses200ResponseResponsesInner`

NewListWhatsAppFlowResponses200ResponseResponsesInner instantiates a new ListWhatsAppFlowResponses200ResponseResponsesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppFlowResponses200ResponseResponsesInnerWithDefaults

`func NewListWhatsAppFlowResponses200ResponseResponsesInnerWithDefaults() *ListWhatsAppFlowResponses200ResponseResponsesInner`

NewListWhatsAppFlowResponses200ResponseResponsesInnerWithDefaults instantiates a new ListWhatsAppFlowResponses200ResponseResponsesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetReceivedAt

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### GetFrom

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### SetFromNil

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetFromNil(b bool)`

 SetFromNil sets the value for From to be an explicit nil

### UnsetFrom
`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) UnsetFrom()`

UnsetFrom ensures that no value is present for From, not even an explicit nil
### GetSenderName

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetSenderName() string`

GetSenderName returns the SenderName field if non-nil, zero value otherwise.

### GetSenderNameOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetSenderNameOk() (*string, bool)`

GetSenderNameOk returns a tuple with the SenderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSenderName

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetSenderName(v string)`

SetSenderName sets SenderName field to given value.

### HasSenderName

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasSenderName() bool`

HasSenderName returns a boolean if a field has been set.

### SetSenderNameNil

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetSenderNameNil(b bool)`

 SetSenderNameNil sets the value for SenderName to be an explicit nil

### UnsetSenderName
`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) UnsetSenderName()`

UnsetSenderName ensures that no value is present for SenderName, not even an explicit nil
### GetConversationId

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### SetConversationIdNil

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetConversationIdNil(b bool)`

 SetConversationIdNil sets the value for ConversationId to be an explicit nil

### UnsetConversationId
`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) UnsetConversationId()`

UnsetConversationId ensures that no value is present for ConversationId, not even an explicit nil
### GetFlowToken

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetFlowToken() string`

GetFlowToken returns the FlowToken field if non-nil, zero value otherwise.

### GetFlowTokenOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetFlowTokenOk() (*string, bool)`

GetFlowTokenOk returns a tuple with the FlowToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowToken

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetFlowToken(v string)`

SetFlowToken sets FlowToken field to given value.

### HasFlowToken

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasFlowToken() bool`

HasFlowToken returns a boolean if a field has been set.

### SetFlowTokenNil

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetFlowTokenNil(b bool)`

 SetFlowTokenNil sets the value for FlowToken to be an explicit nil

### UnsetFlowToken
`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) UnsetFlowToken()`

UnsetFlowToken ensures that no value is present for FlowToken, not even an explicit nil
### GetData

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasData() bool`

HasData returns a boolean if a field has been set.

### GetRaw

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetRaw(v string)`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *ListWhatsAppFlowResponses200ResponseResponsesInner) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


