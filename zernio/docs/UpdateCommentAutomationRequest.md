# UpdateCommentAutomationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**MatchMode** | Pointer to **string** |  | [optional] 
**DmMessage** | Pointer to **string** |  | [optional] 
**Buttons** | Pointer to [**[]DmButton**](DmButton.md) | Inline DM buttons (1-3). Pass [] to clear all buttons. | [optional] 
**CommentReply** | Pointer to **string** |  | [optional] 
**LinkTracking** | Pointer to **bool** | Wrap link buttons in a tracked redirect to count clicks. Pass false to send links untouched. | [optional] 
**ClickTag** | Pointer to **string** | Tag applied to a contact when they click a tracked link (requires linkTracking). Empty string clears it. | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCommentAutomationRequest

`func NewUpdateCommentAutomationRequest() *UpdateCommentAutomationRequest`

NewUpdateCommentAutomationRequest instantiates a new UpdateCommentAutomationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCommentAutomationRequestWithDefaults

`func NewUpdateCommentAutomationRequestWithDefaults() *UpdateCommentAutomationRequest`

NewUpdateCommentAutomationRequestWithDefaults instantiates a new UpdateCommentAutomationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCommentAutomationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCommentAutomationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCommentAutomationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCommentAutomationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetKeywords

`func (o *UpdateCommentAutomationRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *UpdateCommentAutomationRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *UpdateCommentAutomationRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *UpdateCommentAutomationRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMatchMode

`func (o *UpdateCommentAutomationRequest) GetMatchMode() string`

GetMatchMode returns the MatchMode field if non-nil, zero value otherwise.

### GetMatchModeOk

`func (o *UpdateCommentAutomationRequest) GetMatchModeOk() (*string, bool)`

GetMatchModeOk returns a tuple with the MatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMode

`func (o *UpdateCommentAutomationRequest) SetMatchMode(v string)`

SetMatchMode sets MatchMode field to given value.

### HasMatchMode

`func (o *UpdateCommentAutomationRequest) HasMatchMode() bool`

HasMatchMode returns a boolean if a field has been set.

### GetDmMessage

`func (o *UpdateCommentAutomationRequest) GetDmMessage() string`

GetDmMessage returns the DmMessage field if non-nil, zero value otherwise.

### GetDmMessageOk

`func (o *UpdateCommentAutomationRequest) GetDmMessageOk() (*string, bool)`

GetDmMessageOk returns a tuple with the DmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmMessage

`func (o *UpdateCommentAutomationRequest) SetDmMessage(v string)`

SetDmMessage sets DmMessage field to given value.

### HasDmMessage

`func (o *UpdateCommentAutomationRequest) HasDmMessage() bool`

HasDmMessage returns a boolean if a field has been set.

### GetButtons

`func (o *UpdateCommentAutomationRequest) GetButtons() []DmButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *UpdateCommentAutomationRequest) GetButtonsOk() (*[]DmButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *UpdateCommentAutomationRequest) SetButtons(v []DmButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *UpdateCommentAutomationRequest) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetCommentReply

`func (o *UpdateCommentAutomationRequest) GetCommentReply() string`

GetCommentReply returns the CommentReply field if non-nil, zero value otherwise.

### GetCommentReplyOk

`func (o *UpdateCommentAutomationRequest) GetCommentReplyOk() (*string, bool)`

GetCommentReplyOk returns a tuple with the CommentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReply

`func (o *UpdateCommentAutomationRequest) SetCommentReply(v string)`

SetCommentReply sets CommentReply field to given value.

### HasCommentReply

`func (o *UpdateCommentAutomationRequest) HasCommentReply() bool`

HasCommentReply returns a boolean if a field has been set.

### GetLinkTracking

`func (o *UpdateCommentAutomationRequest) GetLinkTracking() bool`

GetLinkTracking returns the LinkTracking field if non-nil, zero value otherwise.

### GetLinkTrackingOk

`func (o *UpdateCommentAutomationRequest) GetLinkTrackingOk() (*bool, bool)`

GetLinkTrackingOk returns a tuple with the LinkTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTracking

`func (o *UpdateCommentAutomationRequest) SetLinkTracking(v bool)`

SetLinkTracking sets LinkTracking field to given value.

### HasLinkTracking

`func (o *UpdateCommentAutomationRequest) HasLinkTracking() bool`

HasLinkTracking returns a boolean if a field has been set.

### GetClickTag

`func (o *UpdateCommentAutomationRequest) GetClickTag() string`

GetClickTag returns the ClickTag field if non-nil, zero value otherwise.

### GetClickTagOk

`func (o *UpdateCommentAutomationRequest) GetClickTagOk() (*string, bool)`

GetClickTagOk returns a tuple with the ClickTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTag

`func (o *UpdateCommentAutomationRequest) SetClickTag(v string)`

SetClickTag sets ClickTag field to given value.

### HasClickTag

`func (o *UpdateCommentAutomationRequest) HasClickTag() bool`

HasClickTag returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateCommentAutomationRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateCommentAutomationRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateCommentAutomationRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateCommentAutomationRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


