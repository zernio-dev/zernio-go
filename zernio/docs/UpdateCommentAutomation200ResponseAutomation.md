# UpdateCommentAutomation200ResponseAutomation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**MatchMode** | Pointer to **string** |  | [optional] 
**DmMessage** | Pointer to **string** |  | [optional] 
**Buttons** | Pointer to [**[]DmButton**](DmButton.md) | Inline DM buttons (up to 3). Omitted when none are set. | [optional] 
**CommentReply** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateCommentAutomation200ResponseAutomation

`func NewUpdateCommentAutomation200ResponseAutomation() *UpdateCommentAutomation200ResponseAutomation`

NewUpdateCommentAutomation200ResponseAutomation instantiates a new UpdateCommentAutomation200ResponseAutomation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCommentAutomation200ResponseAutomationWithDefaults

`func NewUpdateCommentAutomation200ResponseAutomationWithDefaults() *UpdateCommentAutomation200ResponseAutomation`

NewUpdateCommentAutomation200ResponseAutomationWithDefaults instantiates a new UpdateCommentAutomation200ResponseAutomation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateCommentAutomation200ResponseAutomation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateCommentAutomation200ResponseAutomation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateCommentAutomation200ResponseAutomation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateCommentAutomation200ResponseAutomation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCommentAutomation200ResponseAutomation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCommentAutomation200ResponseAutomation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeywords

`func (o *UpdateCommentAutomation200ResponseAutomation) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *UpdateCommentAutomation200ResponseAutomation) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *UpdateCommentAutomation200ResponseAutomation) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMatchMode

`func (o *UpdateCommentAutomation200ResponseAutomation) GetMatchMode() string`

GetMatchMode returns the MatchMode field if non-nil, zero value otherwise.

### GetMatchModeOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetMatchModeOk() (*string, bool)`

GetMatchModeOk returns a tuple with the MatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMode

`func (o *UpdateCommentAutomation200ResponseAutomation) SetMatchMode(v string)`

SetMatchMode sets MatchMode field to given value.

### HasMatchMode

`func (o *UpdateCommentAutomation200ResponseAutomation) HasMatchMode() bool`

HasMatchMode returns a boolean if a field has been set.

### GetDmMessage

`func (o *UpdateCommentAutomation200ResponseAutomation) GetDmMessage() string`

GetDmMessage returns the DmMessage field if non-nil, zero value otherwise.

### GetDmMessageOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetDmMessageOk() (*string, bool)`

GetDmMessageOk returns a tuple with the DmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmMessage

`func (o *UpdateCommentAutomation200ResponseAutomation) SetDmMessage(v string)`

SetDmMessage sets DmMessage field to given value.

### HasDmMessage

`func (o *UpdateCommentAutomation200ResponseAutomation) HasDmMessage() bool`

HasDmMessage returns a boolean if a field has been set.

### GetButtons

`func (o *UpdateCommentAutomation200ResponseAutomation) GetButtons() []DmButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetButtonsOk() (*[]DmButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *UpdateCommentAutomation200ResponseAutomation) SetButtons(v []DmButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *UpdateCommentAutomation200ResponseAutomation) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetCommentReply

`func (o *UpdateCommentAutomation200ResponseAutomation) GetCommentReply() string`

GetCommentReply returns the CommentReply field if non-nil, zero value otherwise.

### GetCommentReplyOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetCommentReplyOk() (*string, bool)`

GetCommentReplyOk returns a tuple with the CommentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReply

`func (o *UpdateCommentAutomation200ResponseAutomation) SetCommentReply(v string)`

SetCommentReply sets CommentReply field to given value.

### HasCommentReply

`func (o *UpdateCommentAutomation200ResponseAutomation) HasCommentReply() bool`

HasCommentReply returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateCommentAutomation200ResponseAutomation) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateCommentAutomation200ResponseAutomation) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateCommentAutomation200ResponseAutomation) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UpdateCommentAutomation200ResponseAutomation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateCommentAutomation200ResponseAutomation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateCommentAutomation200ResponseAutomation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UpdateCommentAutomation200ResponseAutomation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


