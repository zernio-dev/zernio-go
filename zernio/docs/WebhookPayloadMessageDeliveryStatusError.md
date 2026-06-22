# WebhookPayloadMessageDeliveryStatusError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Explanation** | Pointer to **NullableString** | Plain-language translation of &#x60;code&#x60; (e.g. for 131026, that the recipient has likely opted out of marketing messages while utility templates are unaffected). Null for unmapped codes; fall back to title/message.  | [optional] 

## Methods

### NewWebhookPayloadMessageDeliveryStatusError

`func NewWebhookPayloadMessageDeliveryStatusError() *WebhookPayloadMessageDeliveryStatusError`

NewWebhookPayloadMessageDeliveryStatusError instantiates a new WebhookPayloadMessageDeliveryStatusError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadMessageDeliveryStatusErrorWithDefaults

`func NewWebhookPayloadMessageDeliveryStatusErrorWithDefaults() *WebhookPayloadMessageDeliveryStatusError`

NewWebhookPayloadMessageDeliveryStatusErrorWithDefaults instantiates a new WebhookPayloadMessageDeliveryStatusError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WebhookPayloadMessageDeliveryStatusError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebhookPayloadMessageDeliveryStatusError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebhookPayloadMessageDeliveryStatusError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *WebhookPayloadMessageDeliveryStatusError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetTitle

`func (o *WebhookPayloadMessageDeliveryStatusError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookPayloadMessageDeliveryStatusError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookPayloadMessageDeliveryStatusError) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebhookPayloadMessageDeliveryStatusError) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetMessage

`func (o *WebhookPayloadMessageDeliveryStatusError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookPayloadMessageDeliveryStatusError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookPayloadMessageDeliveryStatusError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WebhookPayloadMessageDeliveryStatusError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetExplanation

`func (o *WebhookPayloadMessageDeliveryStatusError) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *WebhookPayloadMessageDeliveryStatusError) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *WebhookPayloadMessageDeliveryStatusError) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *WebhookPayloadMessageDeliveryStatusError) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### SetExplanationNil

`func (o *WebhookPayloadMessageDeliveryStatusError) SetExplanationNil(b bool)`

 SetExplanationNil sets the value for Explanation to be an explicit nil

### UnsetExplanation
`func (o *WebhookPayloadMessageDeliveryStatusError) UnsetExplanation()`

UnsetExplanation ensures that no value is present for Explanation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


