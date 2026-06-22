# SendInboxMessageRequestInteractiveActionOneOf3Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FlowMessageVersion** | Pointer to **string** | Defaults to \&quot;3\&quot; when omitted. | [optional] 
**FlowToken** | **string** | Opaque token you choose to correlate Flow responses with your own state (max 200 chars). | 
**FlowId** | **string** | Published Flow ID from Meta Business Manager. | 
**FlowCta** | **string** | Button label that opens the Flow (max 20 chars). | 
**FlowAction** | **string** | &#x60;navigate&#x60; sends the user to &#x60;flow_action_payload.screen&#x60;; &#x60;data_exchange&#x60; posts data to your Flow endpoint. | 
**FlowActionPayload** | Pointer to [**SendInboxMessageRequestInteractiveActionOneOf3ParametersFlowActionPayload**](SendInboxMessageRequestInteractiveActionOneOf3ParametersFlowActionPayload.md) |  | [optional] 
**Mode** | Pointer to **string** | Set to &#x60;draft&#x60; to test an unpublished Flow. | [optional] 

## Methods

### NewSendInboxMessageRequestInteractiveActionOneOf3Parameters

`func NewSendInboxMessageRequestInteractiveActionOneOf3Parameters(flowToken string, flowId string, flowCta string, flowAction string, ) *SendInboxMessageRequestInteractiveActionOneOf3Parameters`

NewSendInboxMessageRequestInteractiveActionOneOf3Parameters instantiates a new SendInboxMessageRequestInteractiveActionOneOf3Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveActionOneOf3ParametersWithDefaults

`func NewSendInboxMessageRequestInteractiveActionOneOf3ParametersWithDefaults() *SendInboxMessageRequestInteractiveActionOneOf3Parameters`

NewSendInboxMessageRequestInteractiveActionOneOf3ParametersWithDefaults instantiates a new SendInboxMessageRequestInteractiveActionOneOf3Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlowMessageVersion

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowMessageVersion() string`

GetFlowMessageVersion returns the FlowMessageVersion field if non-nil, zero value otherwise.

### GetFlowMessageVersionOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowMessageVersionOk() (*string, bool)`

GetFlowMessageVersionOk returns a tuple with the FlowMessageVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowMessageVersion

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetFlowMessageVersion(v string)`

SetFlowMessageVersion sets FlowMessageVersion field to given value.

### HasFlowMessageVersion

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) HasFlowMessageVersion() bool`

HasFlowMessageVersion returns a boolean if a field has been set.

### GetFlowToken

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowToken() string`

GetFlowToken returns the FlowToken field if non-nil, zero value otherwise.

### GetFlowTokenOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowTokenOk() (*string, bool)`

GetFlowTokenOk returns a tuple with the FlowToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowToken

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetFlowToken(v string)`

SetFlowToken sets FlowToken field to given value.


### GetFlowId

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowCta

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowCta() string`

GetFlowCta returns the FlowCta field if non-nil, zero value otherwise.

### GetFlowCtaOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowCtaOk() (*string, bool)`

GetFlowCtaOk returns a tuple with the FlowCta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCta

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetFlowCta(v string)`

SetFlowCta sets FlowCta field to given value.


### GetFlowAction

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowAction() string`

GetFlowAction returns the FlowAction field if non-nil, zero value otherwise.

### GetFlowActionOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowActionOk() (*string, bool)`

GetFlowActionOk returns a tuple with the FlowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAction

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetFlowAction(v string)`

SetFlowAction sets FlowAction field to given value.


### GetFlowActionPayload

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowActionPayload() SendInboxMessageRequestInteractiveActionOneOf3ParametersFlowActionPayload`

GetFlowActionPayload returns the FlowActionPayload field if non-nil, zero value otherwise.

### GetFlowActionPayloadOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetFlowActionPayloadOk() (*SendInboxMessageRequestInteractiveActionOneOf3ParametersFlowActionPayload, bool)`

GetFlowActionPayloadOk returns a tuple with the FlowActionPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowActionPayload

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetFlowActionPayload(v SendInboxMessageRequestInteractiveActionOneOf3ParametersFlowActionPayload)`

SetFlowActionPayload sets FlowActionPayload field to given value.

### HasFlowActionPayload

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) HasFlowActionPayload() bool`

HasFlowActionPayload returns a boolean if a field has been set.

### GetMode

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *SendInboxMessageRequestInteractiveActionOneOf3Parameters) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


