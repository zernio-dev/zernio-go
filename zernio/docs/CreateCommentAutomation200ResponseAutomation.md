# CreateCommentAutomation200ResponseAutomation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**PlatformPostId** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**MatchMode** | Pointer to **string** |  | [optional] 
**DmMessage** | Pointer to **string** |  | [optional] 
**Buttons** | Pointer to [**[]DmButton**](DmButton.md) | Inline DM buttons (up to 3). Omitted when none are set. | [optional] 
**CommentReply** | Pointer to **string** |  | [optional] 
**LinkTracking** | Pointer to **bool** |  | [optional] 
**ClickTag** | Pointer to **string** |  | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Stats** | Pointer to [**CreateCommentAutomation200ResponseAutomationStats**](CreateCommentAutomation200ResponseAutomationStats.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCreateCommentAutomation200ResponseAutomation

`func NewCreateCommentAutomation200ResponseAutomation() *CreateCommentAutomation200ResponseAutomation`

NewCreateCommentAutomation200ResponseAutomation instantiates a new CreateCommentAutomation200ResponseAutomation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentAutomation200ResponseAutomationWithDefaults

`func NewCreateCommentAutomation200ResponseAutomationWithDefaults() *CreateCommentAutomation200ResponseAutomation`

NewCreateCommentAutomation200ResponseAutomationWithDefaults instantiates a new CreateCommentAutomation200ResponseAutomation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCommentAutomation200ResponseAutomation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCommentAutomation200ResponseAutomation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CreateCommentAutomation200ResponseAutomation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CreateCommentAutomation200ResponseAutomation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCommentAutomation200ResponseAutomation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateCommentAutomation200ResponseAutomation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *CreateCommentAutomation200ResponseAutomation) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateCommentAutomation200ResponseAutomation) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateCommentAutomation200ResponseAutomation) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetTrigger

`func (o *CreateCommentAutomation200ResponseAutomation) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateCommentAutomation200ResponseAutomation) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CreateCommentAutomation200ResponseAutomation) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *CreateCommentAutomation200ResponseAutomation) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *CreateCommentAutomation200ResponseAutomation) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *CreateCommentAutomation200ResponseAutomation) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetKeywords

`func (o *CreateCommentAutomation200ResponseAutomation) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CreateCommentAutomation200ResponseAutomation) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CreateCommentAutomation200ResponseAutomation) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMatchMode

`func (o *CreateCommentAutomation200ResponseAutomation) GetMatchMode() string`

GetMatchMode returns the MatchMode field if non-nil, zero value otherwise.

### GetMatchModeOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetMatchModeOk() (*string, bool)`

GetMatchModeOk returns a tuple with the MatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMode

`func (o *CreateCommentAutomation200ResponseAutomation) SetMatchMode(v string)`

SetMatchMode sets MatchMode field to given value.

### HasMatchMode

`func (o *CreateCommentAutomation200ResponseAutomation) HasMatchMode() bool`

HasMatchMode returns a boolean if a field has been set.

### GetDmMessage

`func (o *CreateCommentAutomation200ResponseAutomation) GetDmMessage() string`

GetDmMessage returns the DmMessage field if non-nil, zero value otherwise.

### GetDmMessageOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetDmMessageOk() (*string, bool)`

GetDmMessageOk returns a tuple with the DmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmMessage

`func (o *CreateCommentAutomation200ResponseAutomation) SetDmMessage(v string)`

SetDmMessage sets DmMessage field to given value.

### HasDmMessage

`func (o *CreateCommentAutomation200ResponseAutomation) HasDmMessage() bool`

HasDmMessage returns a boolean if a field has been set.

### GetButtons

`func (o *CreateCommentAutomation200ResponseAutomation) GetButtons() []DmButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetButtonsOk() (*[]DmButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *CreateCommentAutomation200ResponseAutomation) SetButtons(v []DmButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *CreateCommentAutomation200ResponseAutomation) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetCommentReply

`func (o *CreateCommentAutomation200ResponseAutomation) GetCommentReply() string`

GetCommentReply returns the CommentReply field if non-nil, zero value otherwise.

### GetCommentReplyOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetCommentReplyOk() (*string, bool)`

GetCommentReplyOk returns a tuple with the CommentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReply

`func (o *CreateCommentAutomation200ResponseAutomation) SetCommentReply(v string)`

SetCommentReply sets CommentReply field to given value.

### HasCommentReply

`func (o *CreateCommentAutomation200ResponseAutomation) HasCommentReply() bool`

HasCommentReply returns a boolean if a field has been set.

### GetLinkTracking

`func (o *CreateCommentAutomation200ResponseAutomation) GetLinkTracking() bool`

GetLinkTracking returns the LinkTracking field if non-nil, zero value otherwise.

### GetLinkTrackingOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetLinkTrackingOk() (*bool, bool)`

GetLinkTrackingOk returns a tuple with the LinkTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTracking

`func (o *CreateCommentAutomation200ResponseAutomation) SetLinkTracking(v bool)`

SetLinkTracking sets LinkTracking field to given value.

### HasLinkTracking

`func (o *CreateCommentAutomation200ResponseAutomation) HasLinkTracking() bool`

HasLinkTracking returns a boolean if a field has been set.

### GetClickTag

`func (o *CreateCommentAutomation200ResponseAutomation) GetClickTag() string`

GetClickTag returns the ClickTag field if non-nil, zero value otherwise.

### GetClickTagOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetClickTagOk() (*string, bool)`

GetClickTagOk returns a tuple with the ClickTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTag

`func (o *CreateCommentAutomation200ResponseAutomation) SetClickTag(v string)`

SetClickTag sets ClickTag field to given value.

### HasClickTag

`func (o *CreateCommentAutomation200ResponseAutomation) HasClickTag() bool`

HasClickTag returns a boolean if a field has been set.

### GetIsActive

`func (o *CreateCommentAutomation200ResponseAutomation) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateCommentAutomation200ResponseAutomation) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateCommentAutomation200ResponseAutomation) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetStats

`func (o *CreateCommentAutomation200ResponseAutomation) GetStats() CreateCommentAutomation200ResponseAutomationStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetStatsOk() (*CreateCommentAutomation200ResponseAutomationStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CreateCommentAutomation200ResponseAutomation) SetStats(v CreateCommentAutomation200ResponseAutomationStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CreateCommentAutomation200ResponseAutomation) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreatedAt

`func (o *CreateCommentAutomation200ResponseAutomation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCommentAutomation200ResponseAutomation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCommentAutomation200ResponseAutomation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *CreateCommentAutomation200ResponseAutomation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


