# DiscordPlatformDataPoll

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Question** | Pointer to [**ValidatePostLengthRequest**](ValidatePostLengthRequest.md) |  | [optional] 
**Answers** | Pointer to [**[]DiscordPlatformDataPollAnswersInner**](DiscordPlatformDataPollAnswersInner.md) | 1-10 answer options | [optional] 
**Duration** | Pointer to **int32** | Poll duration in hours (1-768). Default 24. | [optional] 
**AllowMultiselect** | Pointer to **bool** | Allow users to select multiple answers. Default false. | [optional] 

## Methods

### NewDiscordPlatformDataPoll

`func NewDiscordPlatformDataPoll() *DiscordPlatformDataPoll`

NewDiscordPlatformDataPoll instantiates a new DiscordPlatformDataPoll object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDiscordPlatformDataPollWithDefaults

`func NewDiscordPlatformDataPollWithDefaults() *DiscordPlatformDataPoll`

NewDiscordPlatformDataPollWithDefaults instantiates a new DiscordPlatformDataPoll object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuestion

`func (o *DiscordPlatformDataPoll) GetQuestion() ValidatePostLengthRequest`

GetQuestion returns the Question field if non-nil, zero value otherwise.

### GetQuestionOk

`func (o *DiscordPlatformDataPoll) GetQuestionOk() (*ValidatePostLengthRequest, bool)`

GetQuestionOk returns a tuple with the Question field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestion

`func (o *DiscordPlatformDataPoll) SetQuestion(v ValidatePostLengthRequest)`

SetQuestion sets Question field to given value.

### HasQuestion

`func (o *DiscordPlatformDataPoll) HasQuestion() bool`

HasQuestion returns a boolean if a field has been set.

### GetAnswers

`func (o *DiscordPlatformDataPoll) GetAnswers() []DiscordPlatformDataPollAnswersInner`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *DiscordPlatformDataPoll) GetAnswersOk() (*[]DiscordPlatformDataPollAnswersInner, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *DiscordPlatformDataPoll) SetAnswers(v []DiscordPlatformDataPollAnswersInner)`

SetAnswers sets Answers field to given value.

### HasAnswers

`func (o *DiscordPlatformDataPoll) HasAnswers() bool`

HasAnswers returns a boolean if a field has been set.

### GetDuration

`func (o *DiscordPlatformDataPoll) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *DiscordPlatformDataPoll) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *DiscordPlatformDataPoll) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *DiscordPlatformDataPoll) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetAllowMultiselect

`func (o *DiscordPlatformDataPoll) GetAllowMultiselect() bool`

GetAllowMultiselect returns the AllowMultiselect field if non-nil, zero value otherwise.

### GetAllowMultiselectOk

`func (o *DiscordPlatformDataPoll) GetAllowMultiselectOk() (*bool, bool)`

GetAllowMultiselectOk returns a tuple with the AllowMultiselect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiselect

`func (o *DiscordPlatformDataPoll) SetAllowMultiselect(v bool)`

SetAllowMultiselect sets AllowMultiselect field to given value.

### HasAllowMultiselect

`func (o *DiscordPlatformDataPoll) HasAllowMultiselect() bool`

HasAllowMultiselect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


