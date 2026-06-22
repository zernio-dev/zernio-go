# CreateSequenceRequestStepsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | **int32** |  | 
**DelayMinutes** | **int32** |  | 
**Message** | Pointer to [**SendInboxMessageRequestInteractiveFooter**](SendInboxMessageRequestInteractiveFooter.md) |  | [optional] 
**Template** | Pointer to [**UpdateBroadcastRequestTemplate**](UpdateBroadcastRequestTemplate.md) |  | [optional] 

## Methods

### NewCreateSequenceRequestStepsInner

`func NewCreateSequenceRequestStepsInner(order int32, delayMinutes int32, ) *CreateSequenceRequestStepsInner`

NewCreateSequenceRequestStepsInner instantiates a new CreateSequenceRequestStepsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSequenceRequestStepsInnerWithDefaults

`func NewCreateSequenceRequestStepsInnerWithDefaults() *CreateSequenceRequestStepsInner`

NewCreateSequenceRequestStepsInnerWithDefaults instantiates a new CreateSequenceRequestStepsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *CreateSequenceRequestStepsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *CreateSequenceRequestStepsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *CreateSequenceRequestStepsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.


### GetDelayMinutes

`func (o *CreateSequenceRequestStepsInner) GetDelayMinutes() int32`

GetDelayMinutes returns the DelayMinutes field if non-nil, zero value otherwise.

### GetDelayMinutesOk

`func (o *CreateSequenceRequestStepsInner) GetDelayMinutesOk() (*int32, bool)`

GetDelayMinutesOk returns a tuple with the DelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMinutes

`func (o *CreateSequenceRequestStepsInner) SetDelayMinutes(v int32)`

SetDelayMinutes sets DelayMinutes field to given value.


### GetMessage

`func (o *CreateSequenceRequestStepsInner) GetMessage() SendInboxMessageRequestInteractiveFooter`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateSequenceRequestStepsInner) GetMessageOk() (*SendInboxMessageRequestInteractiveFooter, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateSequenceRequestStepsInner) SetMessage(v SendInboxMessageRequestInteractiveFooter)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateSequenceRequestStepsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTemplate

`func (o *CreateSequenceRequestStepsInner) GetTemplate() UpdateBroadcastRequestTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CreateSequenceRequestStepsInner) GetTemplateOk() (*UpdateBroadcastRequestTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CreateSequenceRequestStepsInner) SetTemplate(v UpdateBroadcastRequestTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CreateSequenceRequestStepsInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


