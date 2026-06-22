# ListSequences200ResponseSequencesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountName** | Pointer to **string** | Display name of the sending account | [optional] 
**MessagePreview** | Pointer to **string** | First step template name or message text snippet | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StepsCount** | Pointer to **int32** |  | [optional] 
**ExitOnReply** | Pointer to **bool** |  | [optional] 
**ExitOnUnsubscribe** | Pointer to **bool** |  | [optional] 
**TotalEnrolled** | Pointer to **int32** |  | [optional] 
**TotalCompleted** | Pointer to **int32** |  | [optional] 
**TotalExited** | Pointer to **int32** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListSequences200ResponseSequencesInner

`func NewListSequences200ResponseSequencesInner() *ListSequences200ResponseSequencesInner`

NewListSequences200ResponseSequencesInner instantiates a new ListSequences200ResponseSequencesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSequences200ResponseSequencesInnerWithDefaults

`func NewListSequences200ResponseSequencesInnerWithDefaults() *ListSequences200ResponseSequencesInner`

NewListSequences200ResponseSequencesInnerWithDefaults instantiates a new ListSequences200ResponseSequencesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSequences200ResponseSequencesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSequences200ResponseSequencesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSequences200ResponseSequencesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListSequences200ResponseSequencesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSequences200ResponseSequencesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSequences200ResponseSequencesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSequences200ResponseSequencesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSequences200ResponseSequencesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListSequences200ResponseSequencesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSequences200ResponseSequencesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSequences200ResponseSequencesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSequences200ResponseSequencesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetPlatform

`func (o *ListSequences200ResponseSequencesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListSequences200ResponseSequencesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListSequences200ResponseSequencesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListSequences200ResponseSequencesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *ListSequences200ResponseSequencesInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListSequences200ResponseSequencesInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListSequences200ResponseSequencesInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListSequences200ResponseSequencesInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountName

`func (o *ListSequences200ResponseSequencesInner) GetAccountName() string`

GetAccountName returns the AccountName field if non-nil, zero value otherwise.

### GetAccountNameOk

`func (o *ListSequences200ResponseSequencesInner) GetAccountNameOk() (*string, bool)`

GetAccountNameOk returns a tuple with the AccountName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountName

`func (o *ListSequences200ResponseSequencesInner) SetAccountName(v string)`

SetAccountName sets AccountName field to given value.

### HasAccountName

`func (o *ListSequences200ResponseSequencesInner) HasAccountName() bool`

HasAccountName returns a boolean if a field has been set.

### GetMessagePreview

`func (o *ListSequences200ResponseSequencesInner) GetMessagePreview() string`

GetMessagePreview returns the MessagePreview field if non-nil, zero value otherwise.

### GetMessagePreviewOk

`func (o *ListSequences200ResponseSequencesInner) GetMessagePreviewOk() (*string, bool)`

GetMessagePreviewOk returns a tuple with the MessagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagePreview

`func (o *ListSequences200ResponseSequencesInner) SetMessagePreview(v string)`

SetMessagePreview sets MessagePreview field to given value.

### HasMessagePreview

`func (o *ListSequences200ResponseSequencesInner) HasMessagePreview() bool`

HasMessagePreview returns a boolean if a field has been set.

### GetStatus

`func (o *ListSequences200ResponseSequencesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSequences200ResponseSequencesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSequences200ResponseSequencesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSequences200ResponseSequencesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStepsCount

`func (o *ListSequences200ResponseSequencesInner) GetStepsCount() int32`

GetStepsCount returns the StepsCount field if non-nil, zero value otherwise.

### GetStepsCountOk

`func (o *ListSequences200ResponseSequencesInner) GetStepsCountOk() (*int32, bool)`

GetStepsCountOk returns a tuple with the StepsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepsCount

`func (o *ListSequences200ResponseSequencesInner) SetStepsCount(v int32)`

SetStepsCount sets StepsCount field to given value.

### HasStepsCount

`func (o *ListSequences200ResponseSequencesInner) HasStepsCount() bool`

HasStepsCount returns a boolean if a field has been set.

### GetExitOnReply

`func (o *ListSequences200ResponseSequencesInner) GetExitOnReply() bool`

GetExitOnReply returns the ExitOnReply field if non-nil, zero value otherwise.

### GetExitOnReplyOk

`func (o *ListSequences200ResponseSequencesInner) GetExitOnReplyOk() (*bool, bool)`

GetExitOnReplyOk returns a tuple with the ExitOnReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnReply

`func (o *ListSequences200ResponseSequencesInner) SetExitOnReply(v bool)`

SetExitOnReply sets ExitOnReply field to given value.

### HasExitOnReply

`func (o *ListSequences200ResponseSequencesInner) HasExitOnReply() bool`

HasExitOnReply returns a boolean if a field has been set.

### GetExitOnUnsubscribe

`func (o *ListSequences200ResponseSequencesInner) GetExitOnUnsubscribe() bool`

GetExitOnUnsubscribe returns the ExitOnUnsubscribe field if non-nil, zero value otherwise.

### GetExitOnUnsubscribeOk

`func (o *ListSequences200ResponseSequencesInner) GetExitOnUnsubscribeOk() (*bool, bool)`

GetExitOnUnsubscribeOk returns a tuple with the ExitOnUnsubscribe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExitOnUnsubscribe

`func (o *ListSequences200ResponseSequencesInner) SetExitOnUnsubscribe(v bool)`

SetExitOnUnsubscribe sets ExitOnUnsubscribe field to given value.

### HasExitOnUnsubscribe

`func (o *ListSequences200ResponseSequencesInner) HasExitOnUnsubscribe() bool`

HasExitOnUnsubscribe returns a boolean if a field has been set.

### GetTotalEnrolled

`func (o *ListSequences200ResponseSequencesInner) GetTotalEnrolled() int32`

GetTotalEnrolled returns the TotalEnrolled field if non-nil, zero value otherwise.

### GetTotalEnrolledOk

`func (o *ListSequences200ResponseSequencesInner) GetTotalEnrolledOk() (*int32, bool)`

GetTotalEnrolledOk returns a tuple with the TotalEnrolled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalEnrolled

`func (o *ListSequences200ResponseSequencesInner) SetTotalEnrolled(v int32)`

SetTotalEnrolled sets TotalEnrolled field to given value.

### HasTotalEnrolled

`func (o *ListSequences200ResponseSequencesInner) HasTotalEnrolled() bool`

HasTotalEnrolled returns a boolean if a field has been set.

### GetTotalCompleted

`func (o *ListSequences200ResponseSequencesInner) GetTotalCompleted() int32`

GetTotalCompleted returns the TotalCompleted field if non-nil, zero value otherwise.

### GetTotalCompletedOk

`func (o *ListSequences200ResponseSequencesInner) GetTotalCompletedOk() (*int32, bool)`

GetTotalCompletedOk returns a tuple with the TotalCompleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCompleted

`func (o *ListSequences200ResponseSequencesInner) SetTotalCompleted(v int32)`

SetTotalCompleted sets TotalCompleted field to given value.

### HasTotalCompleted

`func (o *ListSequences200ResponseSequencesInner) HasTotalCompleted() bool`

HasTotalCompleted returns a boolean if a field has been set.

### GetTotalExited

`func (o *ListSequences200ResponseSequencesInner) GetTotalExited() int32`

GetTotalExited returns the TotalExited field if non-nil, zero value otherwise.

### GetTotalExitedOk

`func (o *ListSequences200ResponseSequencesInner) GetTotalExitedOk() (*int32, bool)`

GetTotalExitedOk returns a tuple with the TotalExited field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalExited

`func (o *ListSequences200ResponseSequencesInner) SetTotalExited(v int32)`

SetTotalExited sets TotalExited field to given value.

### HasTotalExited

`func (o *ListSequences200ResponseSequencesInner) HasTotalExited() bool`

HasTotalExited returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListSequences200ResponseSequencesInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListSequences200ResponseSequencesInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListSequences200ResponseSequencesInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListSequences200ResponseSequencesInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


