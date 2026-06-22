# ListWorkflowExecutions200ResponseExecutionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**CurrentNodeId** | Pointer to **string** |  | [optional] 
**WaitingFor** | Pointer to [**ListWorkflowExecutions200ResponseExecutionsInnerWaitingFor**](ListWorkflowExecutions200ResponseExecutionsInnerWaitingFor.md) |  | [optional] 
**Variables** | Pointer to **map[string]interface{}** |  | [optional] 
**PlatformIdentifier** | Pointer to **string** |  | [optional] 
**ConversationId** | Pointer to **string** |  | [optional] 
**StepCount** | Pointer to **int32** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**ResumeAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListWorkflowExecutions200ResponseExecutionsInner

`func NewListWorkflowExecutions200ResponseExecutionsInner() *ListWorkflowExecutions200ResponseExecutionsInner`

NewListWorkflowExecutions200ResponseExecutionsInner instantiates a new ListWorkflowExecutions200ResponseExecutionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowExecutions200ResponseExecutionsInnerWithDefaults

`func NewListWorkflowExecutions200ResponseExecutionsInnerWithDefaults() *ListWorkflowExecutions200ResponseExecutionsInner`

NewListWorkflowExecutions200ResponseExecutionsInnerWithDefaults instantiates a new ListWorkflowExecutions200ResponseExecutionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCurrentNodeId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetCurrentNodeId() string`

GetCurrentNodeId returns the CurrentNodeId field if non-nil, zero value otherwise.

### GetCurrentNodeIdOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetCurrentNodeIdOk() (*string, bool)`

GetCurrentNodeIdOk returns a tuple with the CurrentNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentNodeId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetCurrentNodeId(v string)`

SetCurrentNodeId sets CurrentNodeId field to given value.

### HasCurrentNodeId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasCurrentNodeId() bool`

HasCurrentNodeId returns a boolean if a field has been set.

### GetWaitingFor

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetWaitingFor() ListWorkflowExecutions200ResponseExecutionsInnerWaitingFor`

GetWaitingFor returns the WaitingFor field if non-nil, zero value otherwise.

### GetWaitingForOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetWaitingForOk() (*ListWorkflowExecutions200ResponseExecutionsInnerWaitingFor, bool)`

GetWaitingForOk returns a tuple with the WaitingFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingFor

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetWaitingFor(v ListWorkflowExecutions200ResponseExecutionsInnerWaitingFor)`

SetWaitingFor sets WaitingFor field to given value.

### HasWaitingFor

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasWaitingFor() bool`

HasWaitingFor returns a boolean if a field has been set.

### GetVariables

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

### GetPlatformIdentifier

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetPlatformIdentifier() string`

GetPlatformIdentifier returns the PlatformIdentifier field if non-nil, zero value otherwise.

### GetPlatformIdentifierOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetPlatformIdentifierOk() (*string, bool)`

GetPlatformIdentifierOk returns a tuple with the PlatformIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformIdentifier

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetPlatformIdentifier(v string)`

SetPlatformIdentifier sets PlatformIdentifier field to given value.

### HasPlatformIdentifier

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasPlatformIdentifier() bool`

HasPlatformIdentifier returns a boolean if a field has been set.

### GetConversationId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetStepCount

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetStepCount() int32`

GetStepCount returns the StepCount field if non-nil, zero value otherwise.

### GetStepCountOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetStepCountOk() (*int32, bool)`

GetStepCountOk returns a tuple with the StepCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepCount

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetStepCount(v int32)`

SetStepCount sets StepCount field to given value.

### HasStepCount

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasStepCount() bool`

HasStepCount returns a boolean if a field has been set.

### GetLastError

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListWorkflowExecutions200ResponseExecutionsInner) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetResumeAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetResumeAt() time.Time`

GetResumeAt returns the ResumeAt field if non-nil, zero value otherwise.

### GetResumeAtOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetResumeAtOk() (*time.Time, bool)`

GetResumeAtOk returns a tuple with the ResumeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResumeAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetResumeAt(v time.Time)`

SetResumeAt sets ResumeAt field to given value.

### HasResumeAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasResumeAt() bool`

HasResumeAt returns a boolean if a field has been set.

### SetResumeAtNil

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetResumeAtNil(b bool)`

 SetResumeAtNil sets the value for ResumeAt to be an explicit nil

### UnsetResumeAt
`func (o *ListWorkflowExecutions200ResponseExecutionsInner) UnsetResumeAt()`

UnsetResumeAt ensures that no value is present for ResumeAt, not even an explicit nil
### GetCreatedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCompletedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ListWorkflowExecutions200ResponseExecutionsInner) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ListWorkflowExecutions200ResponseExecutionsInner) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


