# SendWhatsAppFlowMessageRequestFlowActionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Screen** | Pointer to **string** | First screen ID to navigate to | [optional] 
**Data** | Pointer to **map[string]interface{}** | Optional data to pass to the screen | [optional] 

## Methods

### NewSendWhatsAppFlowMessageRequestFlowActionPayload

`func NewSendWhatsAppFlowMessageRequestFlowActionPayload() *SendWhatsAppFlowMessageRequestFlowActionPayload`

NewSendWhatsAppFlowMessageRequestFlowActionPayload instantiates a new SendWhatsAppFlowMessageRequestFlowActionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWhatsAppFlowMessageRequestFlowActionPayloadWithDefaults

`func NewSendWhatsAppFlowMessageRequestFlowActionPayloadWithDefaults() *SendWhatsAppFlowMessageRequestFlowActionPayload`

NewSendWhatsAppFlowMessageRequestFlowActionPayloadWithDefaults instantiates a new SendWhatsAppFlowMessageRequestFlowActionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScreen

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) GetScreen() string`

GetScreen returns the Screen field if non-nil, zero value otherwise.

### GetScreenOk

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) GetScreenOk() (*string, bool)`

GetScreenOk returns a tuple with the Screen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreen

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) SetScreen(v string)`

SetScreen sets Screen field to given value.

### HasScreen

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) HasScreen() bool`

HasScreen returns a boolean if a field has been set.

### GetData

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *SendWhatsAppFlowMessageRequestFlowActionPayload) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


