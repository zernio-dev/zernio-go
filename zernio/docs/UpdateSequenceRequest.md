# UpdateSequenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]CreateSequenceRequestStepsInner**](CreateSequenceRequestStepsInner.md) | Replace the full step list. Only allowed while the sequence is draft or paused. | [optional] 
**ExitOnReply** | Pointer to **bool** |  | [optional] 
**ExitOnUnsubscribe** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateSequenceRequest

`func NewUpdateSequenceRequest() *UpdateSequenceRequest`

NewUpdateSequenceRequest instantiates a new UpdateSequenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSequenceRequestWithDefaults

`func NewUpdateSequenceRequestWithDefaults() *UpdateSequenceRequest`

NewUpdateSequenceRequestWithDefaults instantiates a new UpdateSequenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSequenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSequenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSequenceRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSequenceRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSequenceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSequenceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSequenceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSequenceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSteps

`func (o *UpdateSequenceRequest) GetSteps() []CreateSequenceRequestStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *UpdateSequenceRequest) GetStepsOk() (*[]CreateSequenceRequestStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *UpdateSequenceRequest) SetSteps(v []CreateSequenceRequestStepsInner)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *UpdateSequenceRequest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetExitOnReply

`func (o *UpdateSequenceRequest) GetExitOnReply() bool`

GetExitOnReply returns the ExitOnReply field if non-nil, zero value otherwise.

### GetExitOnReplyOk

`func (o *UpdateSequenceRequest) GetExitOnReplyOk() (*bool, bool)`

GetExitOnReplyOk returns a tuple with the ExitOnReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnReply

`func (o *UpdateSequenceRequest) SetExitOnReply(v bool)`

SetExitOnReply sets ExitOnReply field to given value.

### HasExitOnReply

`func (o *UpdateSequenceRequest) HasExitOnReply() bool`

HasExitOnReply returns a boolean if a field has been set.

### GetExitOnUnsubscribe

`func (o *UpdateSequenceRequest) GetExitOnUnsubscribe() bool`

GetExitOnUnsubscribe returns the ExitOnUnsubscribe field if non-nil, zero value otherwise.

### GetExitOnUnsubscribeOk

`func (o *UpdateSequenceRequest) GetExitOnUnsubscribeOk() (*bool, bool)`

GetExitOnUnsubscribeOk returns a tuple with the ExitOnUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnUnsubscribe

`func (o *UpdateSequenceRequest) SetExitOnUnsubscribe(v bool)`

SetExitOnUnsubscribe sets ExitOnUnsubscribe field to given value.

### HasExitOnUnsubscribe

`func (o *UpdateSequenceRequest) HasExitOnUnsubscribe() bool`

HasExitOnUnsubscribe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


