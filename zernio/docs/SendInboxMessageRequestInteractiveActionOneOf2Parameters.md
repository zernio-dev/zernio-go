# SendInboxMessageRequestInteractiveActionOneOf2Parameters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayText** | Pointer to **string** | Button label. Defaults to \&quot;Call Now\&quot;. | [optional] 
**TtlMinutes** | Pointer to **int32** | How long the button stays tappable. Defaults to 10080 (7 days). | [optional] 
**Payload** | Pointer to **string** | Arbitrary string echoed back as &#x60;cta_payload&#x60; on the &#x60;calls&#x60; webhook (connect/terminate) for attribution. | [optional] 

## Methods

### NewSendInboxMessageRequestInteractiveActionOneOf2Parameters

`func NewSendInboxMessageRequestInteractiveActionOneOf2Parameters() *SendInboxMessageRequestInteractiveActionOneOf2Parameters`

NewSendInboxMessageRequestInteractiveActionOneOf2Parameters instantiates a new SendInboxMessageRequestInteractiveActionOneOf2Parameters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveActionOneOf2ParametersWithDefaults

`func NewSendInboxMessageRequestInteractiveActionOneOf2ParametersWithDefaults() *SendInboxMessageRequestInteractiveActionOneOf2Parameters`

NewSendInboxMessageRequestInteractiveActionOneOf2ParametersWithDefaults instantiates a new SendInboxMessageRequestInteractiveActionOneOf2Parameters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayText

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) GetDisplayText() string`

GetDisplayText returns the DisplayText field if non-nil, zero value otherwise.

### GetDisplayTextOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) GetDisplayTextOk() (*string, bool)`

GetDisplayTextOk returns a tuple with the DisplayText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayText

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) SetDisplayText(v string)`

SetDisplayText sets DisplayText field to given value.

### HasDisplayText

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) HasDisplayText() bool`

HasDisplayText returns a boolean if a field has been set.

### GetTtlMinutes

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) GetTtlMinutes() int32`

GetTtlMinutes returns the TtlMinutes field if non-nil, zero value otherwise.

### GetTtlMinutesOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) GetTtlMinutesOk() (*int32, bool)`

GetTtlMinutesOk returns a tuple with the TtlMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtlMinutes

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) SetTtlMinutes(v int32)`

SetTtlMinutes sets TtlMinutes field to given value.

### HasTtlMinutes

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) HasTtlMinutes() bool`

HasTtlMinutes returns a boolean if a field has been set.

### GetPayload

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SendInboxMessageRequestInteractiveActionOneOf2Parameters) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


