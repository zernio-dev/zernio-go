# CreateSequenceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**AccountId** | **string** |  | 
**Platform** | **string** |  | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]CreateSequenceRequestStepsInner**](CreateSequenceRequestStepsInner.md) |  | [optional] 
**ExitOnReply** | Pointer to **bool** |  | [optional] [default to true]
**ExitOnUnsubscribe** | Pointer to **bool** |  | [optional] [default to true]

## Methods

### NewCreateSequenceRequest

`func NewCreateSequenceRequest(profileId string, accountId string, platform string, name string, ) *CreateSequenceRequest`

NewCreateSequenceRequest instantiates a new CreateSequenceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSequenceRequestWithDefaults

`func NewCreateSequenceRequestWithDefaults() *CreateSequenceRequest`

NewCreateSequenceRequestWithDefaults instantiates a new CreateSequenceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateSequenceRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateSequenceRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateSequenceRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *CreateSequenceRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateSequenceRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateSequenceRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPlatform

`func (o *CreateSequenceRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateSequenceRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateSequenceRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetName

`func (o *CreateSequenceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSequenceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSequenceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateSequenceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateSequenceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateSequenceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateSequenceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSteps

`func (o *CreateSequenceRequest) GetSteps() []CreateSequenceRequestStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *CreateSequenceRequest) GetStepsOk() (*[]CreateSequenceRequestStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *CreateSequenceRequest) SetSteps(v []CreateSequenceRequestStepsInner)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *CreateSequenceRequest) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetExitOnReply

`func (o *CreateSequenceRequest) GetExitOnReply() bool`

GetExitOnReply returns the ExitOnReply field if non-nil, zero value otherwise.

### GetExitOnReplyOk

`func (o *CreateSequenceRequest) GetExitOnReplyOk() (*bool, bool)`

GetExitOnReplyOk returns a tuple with the ExitOnReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnReply

`func (o *CreateSequenceRequest) SetExitOnReply(v bool)`

SetExitOnReply sets ExitOnReply field to given value.

### HasExitOnReply

`func (o *CreateSequenceRequest) HasExitOnReply() bool`

HasExitOnReply returns a boolean if a field has been set.

### GetExitOnUnsubscribe

`func (o *CreateSequenceRequest) GetExitOnUnsubscribe() bool`

GetExitOnUnsubscribe returns the ExitOnUnsubscribe field if non-nil, zero value otherwise.

### GetExitOnUnsubscribeOk

`func (o *CreateSequenceRequest) GetExitOnUnsubscribeOk() (*bool, bool)`

GetExitOnUnsubscribeOk returns a tuple with the ExitOnUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnUnsubscribe

`func (o *CreateSequenceRequest) SetExitOnUnsubscribe(v bool)`

SetExitOnUnsubscribe sets ExitOnUnsubscribe field to given value.

### HasExitOnUnsubscribe

`func (o *CreateSequenceRequest) HasExitOnUnsubscribe() bool`

HasExitOnUnsubscribe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


