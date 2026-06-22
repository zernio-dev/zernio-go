# SendWhatsAppFlowMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**To** | **string** | Recipient phone number (E.164 format, e.g. +1234567890) | 
**FlowId** | **string** | Published flow ID | 
**FlowCta** | **string** | CTA button text (e.g. &#39;Book Now&#39;, &#39;Sign Up&#39;) | 
**FlowAction** | Pointer to **string** | Action type: navigate opens a screen directly, data_exchange hits your endpoint first | [optional] [default to "navigate"]
**FlowToken** | Pointer to **string** | Unique token to correlate responses. If omitted, auto-generated as &#39;&lt;flowId&gt;:&lt;uuid&gt;&#39; so the response can be attributed to this flow in the Flow Responses view. | [optional] 
**FlowActionPayload** | Pointer to [**SendWhatsAppFlowMessageRequestFlowActionPayload**](SendWhatsAppFlowMessageRequestFlowActionPayload.md) |  | [optional] 
**Body** | **string** | Message body text | 
**Header** | Pointer to [**SendWhatsAppFlowMessageRequestHeader**](SendWhatsAppFlowMessageRequestHeader.md) |  | [optional] 
**Footer** | Pointer to **string** | Optional footer text | [optional] 
**Draft** | Pointer to **bool** | Set true to test an unpublished (DRAFT) flow | [optional] 

## Methods

### NewSendWhatsAppFlowMessageRequest

`func NewSendWhatsAppFlowMessageRequest(accountId string, to string, flowId string, flowCta string, body string, ) *SendWhatsAppFlowMessageRequest`

NewSendWhatsAppFlowMessageRequest instantiates a new SendWhatsAppFlowMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWhatsAppFlowMessageRequestWithDefaults

`func NewSendWhatsAppFlowMessageRequestWithDefaults() *SendWhatsAppFlowMessageRequest`

NewSendWhatsAppFlowMessageRequestWithDefaults instantiates a new SendWhatsAppFlowMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SendWhatsAppFlowMessageRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SendWhatsAppFlowMessageRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SendWhatsAppFlowMessageRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTo

`func (o *SendWhatsAppFlowMessageRequest) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *SendWhatsAppFlowMessageRequest) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *SendWhatsAppFlowMessageRequest) SetTo(v string)`

SetTo sets To field to given value.


### GetFlowId

`func (o *SendWhatsAppFlowMessageRequest) GetFlowId() string`

GetFlowId returns the FlowId field if non-nil, zero value otherwise.

### GetFlowIdOk

`func (o *SendWhatsAppFlowMessageRequest) GetFlowIdOk() (*string, bool)`

GetFlowIdOk returns a tuple with the FlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowId

`func (o *SendWhatsAppFlowMessageRequest) SetFlowId(v string)`

SetFlowId sets FlowId field to given value.


### GetFlowCta

`func (o *SendWhatsAppFlowMessageRequest) GetFlowCta() string`

GetFlowCta returns the FlowCta field if non-nil, zero value otherwise.

### GetFlowCtaOk

`func (o *SendWhatsAppFlowMessageRequest) GetFlowCtaOk() (*string, bool)`

GetFlowCtaOk returns a tuple with the FlowCta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowCta

`func (o *SendWhatsAppFlowMessageRequest) SetFlowCta(v string)`

SetFlowCta sets FlowCta field to given value.


### GetFlowAction

`func (o *SendWhatsAppFlowMessageRequest) GetFlowAction() string`

GetFlowAction returns the FlowAction field if non-nil, zero value otherwise.

### GetFlowActionOk

`func (o *SendWhatsAppFlowMessageRequest) GetFlowActionOk() (*string, bool)`

GetFlowActionOk returns a tuple with the FlowAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowAction

`func (o *SendWhatsAppFlowMessageRequest) SetFlowAction(v string)`

SetFlowAction sets FlowAction field to given value.

### HasFlowAction

`func (o *SendWhatsAppFlowMessageRequest) HasFlowAction() bool`

HasFlowAction returns a boolean if a field has been set.

### GetFlowToken

`func (o *SendWhatsAppFlowMessageRequest) GetFlowToken() string`

GetFlowToken returns the FlowToken field if non-nil, zero value otherwise.

### GetFlowTokenOk

`func (o *SendWhatsAppFlowMessageRequest) GetFlowTokenOk() (*string, bool)`

GetFlowTokenOk returns a tuple with the FlowToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowToken

`func (o *SendWhatsAppFlowMessageRequest) SetFlowToken(v string)`

SetFlowToken sets FlowToken field to given value.

### HasFlowToken

`func (o *SendWhatsAppFlowMessageRequest) HasFlowToken() bool`

HasFlowToken returns a boolean if a field has been set.

### GetFlowActionPayload

`func (o *SendWhatsAppFlowMessageRequest) GetFlowActionPayload() SendWhatsAppFlowMessageRequestFlowActionPayload`

GetFlowActionPayload returns the FlowActionPayload field if non-nil, zero value otherwise.

### GetFlowActionPayloadOk

`func (o *SendWhatsAppFlowMessageRequest) GetFlowActionPayloadOk() (*SendWhatsAppFlowMessageRequestFlowActionPayload, bool)`

GetFlowActionPayloadOk returns a tuple with the FlowActionPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowActionPayload

`func (o *SendWhatsAppFlowMessageRequest) SetFlowActionPayload(v SendWhatsAppFlowMessageRequestFlowActionPayload)`

SetFlowActionPayload sets FlowActionPayload field to given value.

### HasFlowActionPayload

`func (o *SendWhatsAppFlowMessageRequest) HasFlowActionPayload() bool`

HasFlowActionPayload returns a boolean if a field has been set.

### GetBody

`func (o *SendWhatsAppFlowMessageRequest) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendWhatsAppFlowMessageRequest) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendWhatsAppFlowMessageRequest) SetBody(v string)`

SetBody sets Body field to given value.


### GetHeader

`func (o *SendWhatsAppFlowMessageRequest) GetHeader() SendWhatsAppFlowMessageRequestHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SendWhatsAppFlowMessageRequest) GetHeaderOk() (*SendWhatsAppFlowMessageRequestHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SendWhatsAppFlowMessageRequest) SetHeader(v SendWhatsAppFlowMessageRequestHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SendWhatsAppFlowMessageRequest) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetFooter

`func (o *SendWhatsAppFlowMessageRequest) GetFooter() string`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendWhatsAppFlowMessageRequest) GetFooterOk() (*string, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendWhatsAppFlowMessageRequest) SetFooter(v string)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendWhatsAppFlowMessageRequest) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetDraft

`func (o *SendWhatsAppFlowMessageRequest) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *SendWhatsAppFlowMessageRequest) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *SendWhatsAppFlowMessageRequest) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *SendWhatsAppFlowMessageRequest) HasDraft() bool`

HasDraft returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


