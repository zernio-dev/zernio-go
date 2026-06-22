# GetSequence200ResponseSequence

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Steps** | Pointer to [**[]GetSequence200ResponseSequenceStepsInner**](GetSequence200ResponseSequenceStepsInner.md) |  | [optional] 
**ExitOnReply** | Pointer to **bool** |  | [optional] 
**ExitOnUnsubscribe** | Pointer to **bool** |  | [optional] 
**TotalEnrolled** | Pointer to **int32** |  | [optional] 
**TotalCompleted** | Pointer to **int32** |  | [optional] 
**TotalExited** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetSequence200ResponseSequence

`func NewGetSequence200ResponseSequence() *GetSequence200ResponseSequence`

NewGetSequence200ResponseSequence instantiates a new GetSequence200ResponseSequence object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSequence200ResponseSequenceWithDefaults

`func NewGetSequence200ResponseSequenceWithDefaults() *GetSequence200ResponseSequence`

NewGetSequence200ResponseSequenceWithDefaults instantiates a new GetSequence200ResponseSequence object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSequence200ResponseSequence) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSequence200ResponseSequence) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSequence200ResponseSequence) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetSequence200ResponseSequence) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSequence200ResponseSequence) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSequence200ResponseSequence) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSequence200ResponseSequence) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSequence200ResponseSequence) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetSequence200ResponseSequence) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSequence200ResponseSequence) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSequence200ResponseSequence) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSequence200ResponseSequence) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *GetSequence200ResponseSequence) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetSequence200ResponseSequence) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetSequence200ResponseSequence) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetSequence200ResponseSequence) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *GetSequence200ResponseSequence) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetSequence200ResponseSequence) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetSequence200ResponseSequence) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetSequence200ResponseSequence) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *GetSequence200ResponseSequence) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSequence200ResponseSequence) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSequence200ResponseSequence) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSequence200ResponseSequence) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSteps

`func (o *GetSequence200ResponseSequence) GetSteps() []GetSequence200ResponseSequenceStepsInner`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GetSequence200ResponseSequence) GetStepsOk() (*[]GetSequence200ResponseSequenceStepsInner, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GetSequence200ResponseSequence) SetSteps(v []GetSequence200ResponseSequenceStepsInner)`

SetSteps sets Steps field to given value.

### HasSteps

`func (o *GetSequence200ResponseSequence) HasSteps() bool`

HasSteps returns a boolean if a field has been set.

### GetExitOnReply

`func (o *GetSequence200ResponseSequence) GetExitOnReply() bool`

GetExitOnReply returns the ExitOnReply field if non-nil, zero value otherwise.

### GetExitOnReplyOk

`func (o *GetSequence200ResponseSequence) GetExitOnReplyOk() (*bool, bool)`

GetExitOnReplyOk returns a tuple with the ExitOnReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnReply

`func (o *GetSequence200ResponseSequence) SetExitOnReply(v bool)`

SetExitOnReply sets ExitOnReply field to given value.

### HasExitOnReply

`func (o *GetSequence200ResponseSequence) HasExitOnReply() bool`

HasExitOnReply returns a boolean if a field has been set.

### GetExitOnUnsubscribe

`func (o *GetSequence200ResponseSequence) GetExitOnUnsubscribe() bool`

GetExitOnUnsubscribe returns the ExitOnUnsubscribe field if non-nil, zero value otherwise.

### GetExitOnUnsubscribeOk

`func (o *GetSequence200ResponseSequence) GetExitOnUnsubscribeOk() (*bool, bool)`

GetExitOnUnsubscribeOk returns a tuple with the ExitOnUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnUnsubscribe

`func (o *GetSequence200ResponseSequence) SetExitOnUnsubscribe(v bool)`

SetExitOnUnsubscribe sets ExitOnUnsubscribe field to given value.

### HasExitOnUnsubscribe

`func (o *GetSequence200ResponseSequence) HasExitOnUnsubscribe() bool`

HasExitOnUnsubscribe returns a boolean if a field has been set.

### GetTotalEnrolled

`func (o *GetSequence200ResponseSequence) GetTotalEnrolled() int32`

GetTotalEnrolled returns the TotalEnrolled field if non-nil, zero value otherwise.

### GetTotalEnrolledOk

`func (o *GetSequence200ResponseSequence) GetTotalEnrolledOk() (*int32, bool)`

GetTotalEnrolledOk returns a tuple with the TotalEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnrolled

`func (o *GetSequence200ResponseSequence) SetTotalEnrolled(v int32)`

SetTotalEnrolled sets TotalEnrolled field to given value.

### HasTotalEnrolled

`func (o *GetSequence200ResponseSequence) HasTotalEnrolled() bool`

HasTotalEnrolled returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *GetSequence200ResponseSequence) GetTotalCompleted() int32`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *GetSequence200ResponseSequence) GetTotalCompletedOk() (*int32, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *GetSequence200ResponseSequence) SetTotalCompleted(v int32)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *GetSequence200ResponseSequence) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetTotalExited

`func (o *GetSequence200ResponseSequence) GetTotalExited() int32`

GetTotalExited returns the TotalExited field if non-nil, zero value otherwise.

### GetTotalExitedOk

`func (o *GetSequence200ResponseSequence) GetTotalExitedOk() (*int32, bool)`

GetTotalExitedOk returns a tuple with the TotalExited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExited

`func (o *GetSequence200ResponseSequence) SetTotalExited(v int32)`

SetTotalExited sets TotalExited field to given value.

### HasTotalExited

`func (o *GetSequence200ResponseSequence) HasTotalExited() bool`

HasTotalExited returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetSequence200ResponseSequence) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetSequence200ResponseSequence) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetSequence200ResponseSequence) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetSequence200ResponseSequence) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetSequence200ResponseSequence) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetSequence200ResponseSequence) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetSequence200ResponseSequence) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetSequence200ResponseSequence) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


