# GetInboxConversationAnalytics200ResponseSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Received** | Pointer to **int32** |  | [optional] 
**Sent** | Pointer to **int32** |  | [optional] 
**Read** | Pointer to **int32** |  | [optional] 
**Failed** | Pointer to **int32** |  | [optional] 
**TotalMessages** | Pointer to **int32** |  | [optional] 
**FirstMessageAt** | Pointer to **NullableTime** |  | [optional] 
**LastMessageAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetInboxConversationAnalytics200ResponseSummary

`func NewGetInboxConversationAnalytics200ResponseSummary() *GetInboxConversationAnalytics200ResponseSummary`

NewGetInboxConversationAnalytics200ResponseSummary instantiates a new GetInboxConversationAnalytics200ResponseSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversationAnalytics200ResponseSummaryWithDefaults

`func NewGetInboxConversationAnalytics200ResponseSummaryWithDefaults() *GetInboxConversationAnalytics200ResponseSummary`

NewGetInboxConversationAnalytics200ResponseSummaryWithDefaults instantiates a new GetInboxConversationAnalytics200ResponseSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReceived

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSent

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetRead

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetRead() int32`

GetRead returns the Read field if non-nil, zero value otherwise.

### GetReadOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetReadOk() (*int32, bool)`

GetReadOk returns a tuple with the Read field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRead

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetRead(v int32)`

SetRead sets Read field to given value.

### HasRead

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasRead() bool`

HasRead returns a boolean if a field has been set.

### GetFailed

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetTotalMessages

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetTotalMessages() int32`

GetTotalMessages returns the TotalMessages field if non-nil, zero value otherwise.

### GetTotalMessagesOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetTotalMessagesOk() (*int32, bool)`

GetTotalMessagesOk returns a tuple with the TotalMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalMessages

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetTotalMessages(v int32)`

SetTotalMessages sets TotalMessages field to given value.

### HasTotalMessages

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasTotalMessages() bool`

HasTotalMessages returns a boolean if a field has been set.

### GetFirstMessageAt

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetFirstMessageAt() time.Time`

GetFirstMessageAt returns the FirstMessageAt field if non-nil, zero value otherwise.

### GetFirstMessageAtOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetFirstMessageAtOk() (*time.Time, bool)`

GetFirstMessageAtOk returns a tuple with the FirstMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstMessageAt

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetFirstMessageAt(v time.Time)`

SetFirstMessageAt sets FirstMessageAt field to given value.

### HasFirstMessageAt

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasFirstMessageAt() bool`

HasFirstMessageAt returns a boolean if a field has been set.

### SetFirstMessageAtNil

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetFirstMessageAtNil(b bool)`

 SetFirstMessageAtNil sets the value for FirstMessageAt to be an explicit nil

### UnsetFirstMessageAt
`func (o *GetInboxConversationAnalytics200ResponseSummary) UnsetFirstMessageAt()`

UnsetFirstMessageAt ensures that no value is present for FirstMessageAt, not even an explicit nil
### GetLastMessageAt

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *GetInboxConversationAnalytics200ResponseSummary) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *GetInboxConversationAnalytics200ResponseSummary) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### SetLastMessageAtNil

`func (o *GetInboxConversationAnalytics200ResponseSummary) SetLastMessageAtNil(b bool)`

 SetLastMessageAtNil sets the value for LastMessageAt to be an explicit nil

### UnsetLastMessageAt
`func (o *GetInboxConversationAnalytics200ResponseSummary) UnsetLastMessageAt()`

UnsetLastMessageAt ensures that no value is present for LastMessageAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


