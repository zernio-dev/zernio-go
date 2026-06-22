# ListCommentAutomations200ResponseAutomationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**PlatformPostId** | Pointer to **string** |  | [optional] 
**PostTitle** | Pointer to **string** |  | [optional] 
**Keywords** | Pointer to **[]string** |  | [optional] 
**MatchMode** | Pointer to **string** |  | [optional] 
**DmMessage** | Pointer to **string** |  | [optional] 
**Buttons** | Pointer to [**[]DmButton**](DmButton.md) | Inline DM buttons (up to 3). Omitted when none are set. | [optional] 
**CommentReply** | Pointer to **string** |  | [optional] 
**LinkTracking** | Pointer to **bool** | Whether link buttons in the DM are wrapped in a tracked redirect to count clicks. | [optional] 
**ClickTag** | Pointer to **string** | Tag applied to a contact when they click a tracked link. | [optional] 
**IsActive** | Pointer to **bool** |  | [optional] 
**Stats** | Pointer to [**ListCommentAutomations200ResponseAutomationsInnerStats**](ListCommentAutomations200ResponseAutomationsInnerStats.md) |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListCommentAutomations200ResponseAutomationsInner

`func NewListCommentAutomations200ResponseAutomationsInner() *ListCommentAutomations200ResponseAutomationsInner`

NewListCommentAutomations200ResponseAutomationsInner instantiates a new ListCommentAutomations200ResponseAutomationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentAutomations200ResponseAutomationsInnerWithDefaults

`func NewListCommentAutomations200ResponseAutomationsInnerWithDefaults() *ListCommentAutomations200ResponseAutomationsInner`

NewListCommentAutomations200ResponseAutomationsInnerWithDefaults instantiates a new ListCommentAutomations200ResponseAutomationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetTrigger

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetAccountId

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetPostTitle

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetPostTitle() string`

GetPostTitle returns the PostTitle field if non-nil, zero value otherwise.

### GetPostTitleOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetPostTitleOk() (*string, bool)`

GetPostTitleOk returns a tuple with the PostTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTitle

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetPostTitle(v string)`

SetPostTitle sets PostTitle field to given value.

### HasPostTitle

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasPostTitle() bool`

HasPostTitle returns a boolean if a field has been set.

### GetKeywords

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMatchMode

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetMatchMode() string`

GetMatchMode returns the MatchMode field if non-nil, zero value otherwise.

### GetMatchModeOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetMatchModeOk() (*string, bool)`

GetMatchModeOk returns a tuple with the MatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMode

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetMatchMode(v string)`

SetMatchMode sets MatchMode field to given value.

### HasMatchMode

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasMatchMode() bool`

HasMatchMode returns a boolean if a field has been set.

### GetDmMessage

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetDmMessage() string`

GetDmMessage returns the DmMessage field if non-nil, zero value otherwise.

### GetDmMessageOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetDmMessageOk() (*string, bool)`

GetDmMessageOk returns a tuple with the DmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmMessage

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetDmMessage(v string)`

SetDmMessage sets DmMessage field to given value.

### HasDmMessage

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasDmMessage() bool`

HasDmMessage returns a boolean if a field has been set.

### GetButtons

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetButtons() []DmButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetButtonsOk() (*[]DmButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetButtons(v []DmButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetCommentReply

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetCommentReply() string`

GetCommentReply returns the CommentReply field if non-nil, zero value otherwise.

### GetCommentReplyOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetCommentReplyOk() (*string, bool)`

GetCommentReplyOk returns a tuple with the CommentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReply

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetCommentReply(v string)`

SetCommentReply sets CommentReply field to given value.

### HasCommentReply

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasCommentReply() bool`

HasCommentReply returns a boolean if a field has been set.

### GetLinkTracking

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetLinkTracking() bool`

GetLinkTracking returns the LinkTracking field if non-nil, zero value otherwise.

### GetLinkTrackingOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetLinkTrackingOk() (*bool, bool)`

GetLinkTrackingOk returns a tuple with the LinkTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTracking

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetLinkTracking(v bool)`

SetLinkTracking sets LinkTracking field to given value.

### HasLinkTracking

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasLinkTracking() bool`

HasLinkTracking returns a boolean if a field has been set.

### GetClickTag

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetClickTag() string`

GetClickTag returns the ClickTag field if non-nil, zero value otherwise.

### GetClickTagOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetClickTagOk() (*string, bool)`

GetClickTagOk returns a tuple with the ClickTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTag

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetClickTag(v string)`

SetClickTag sets ClickTag field to given value.

### HasClickTag

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasClickTag() bool`

HasClickTag returns a boolean if a field has been set.

### GetIsActive

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetStats

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetStats() ListCommentAutomations200ResponseAutomationsInnerStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetStatsOk() (*ListCommentAutomations200ResponseAutomationsInnerStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetStats(v ListCommentAutomations200ResponseAutomationsInnerStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListCommentAutomations200ResponseAutomationsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListCommentAutomations200ResponseAutomationsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListCommentAutomations200ResponseAutomationsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


