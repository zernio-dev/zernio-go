# GetSequence200ResponseSequenceStepsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to **int32** |  | [optional] 
**DelayMinutes** | Pointer to **int32** |  | [optional] 
**Message** | Pointer to [**SendInboxMessageRequestInteractiveFooter**](SendInboxMessageRequestInteractiveFooter.md) |  | [optional] 
**Template** | Pointer to [**GetSequence200ResponseSequenceStepsInnerTemplate**](GetSequence200ResponseSequenceStepsInnerTemplate.md) |  | [optional] 

## Methods

### NewGetSequence200ResponseSequenceStepsInner

`func NewGetSequence200ResponseSequenceStepsInner() *GetSequence200ResponseSequenceStepsInner`

NewGetSequence200ResponseSequenceStepsInner instantiates a new GetSequence200ResponseSequenceStepsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSequence200ResponseSequenceStepsInnerWithDefaults

`func NewGetSequence200ResponseSequenceStepsInnerWithDefaults() *GetSequence200ResponseSequenceStepsInner`

NewGetSequence200ResponseSequenceStepsInnerWithDefaults instantiates a new GetSequence200ResponseSequenceStepsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *GetSequence200ResponseSequenceStepsInner) GetOrder() int32`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *GetSequence200ResponseSequenceStepsInner) GetOrderOk() (*int32, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *GetSequence200ResponseSequenceStepsInner) SetOrder(v int32)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *GetSequence200ResponseSequenceStepsInner) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetDelayMinutes

`func (o *GetSequence200ResponseSequenceStepsInner) GetDelayMinutes() int32`

GetDelayMinutes returns the DelayMinutes field if non-nil, zero value otherwise.

### GetDelayMinutesOk

`func (o *GetSequence200ResponseSequenceStepsInner) GetDelayMinutesOk() (*int32, bool)`

GetDelayMinutesOk returns a tuple with the DelayMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayMinutes

`func (o *GetSequence200ResponseSequenceStepsInner) SetDelayMinutes(v int32)`

SetDelayMinutes sets DelayMinutes field to given value.

### HasDelayMinutes

`func (o *GetSequence200ResponseSequenceStepsInner) HasDelayMinutes() bool`

HasDelayMinutes returns a boolean if a field has been set.

### GetMessage

`func (o *GetSequence200ResponseSequenceStepsInner) GetMessage() SendInboxMessageRequestInteractiveFooter`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetSequence200ResponseSequenceStepsInner) GetMessageOk() (*SendInboxMessageRequestInteractiveFooter, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetSequence200ResponseSequenceStepsInner) SetMessage(v SendInboxMessageRequestInteractiveFooter)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetSequence200ResponseSequenceStepsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTemplate

`func (o *GetSequence200ResponseSequenceStepsInner) GetTemplate() GetSequence200ResponseSequenceStepsInnerTemplate`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *GetSequence200ResponseSequenceStepsInner) GetTemplateOk() (*GetSequence200ResponseSequenceStepsInnerTemplate, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *GetSequence200ResponseSequenceStepsInner) SetTemplate(v GetSequence200ResponseSequenceStepsInnerTemplate)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *GetSequence200ResponseSequenceStepsInner) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


