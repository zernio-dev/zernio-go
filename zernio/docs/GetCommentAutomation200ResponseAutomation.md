# GetCommentAutomation200ResponseAutomation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Trigger** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**PlatformPostId** | Pointer to **string** |  | [optional] 
**PostId** | Pointer to **string** |  | [optional] 
**PostTitle** | Pointer to **string** |  | [optional] 
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
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetCommentAutomation200ResponseAutomation

`func NewGetCommentAutomation200ResponseAutomation() *GetCommentAutomation200ResponseAutomation`

NewGetCommentAutomation200ResponseAutomation instantiates a new GetCommentAutomation200ResponseAutomation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentAutomation200ResponseAutomationWithDefaults

`func NewGetCommentAutomation200ResponseAutomationWithDefaults() *GetCommentAutomation200ResponseAutomation`

NewGetCommentAutomation200ResponseAutomationWithDefaults instantiates a new GetCommentAutomation200ResponseAutomation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCommentAutomation200ResponseAutomation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCommentAutomation200ResponseAutomation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCommentAutomation200ResponseAutomation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetCommentAutomation200ResponseAutomation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCommentAutomation200ResponseAutomation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCommentAutomation200ResponseAutomation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCommentAutomation200ResponseAutomation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCommentAutomation200ResponseAutomation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPlatform

`func (o *GetCommentAutomation200ResponseAutomation) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetCommentAutomation200ResponseAutomation) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetCommentAutomation200ResponseAutomation) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetCommentAutomation200ResponseAutomation) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetTrigger

`func (o *GetCommentAutomation200ResponseAutomation) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *GetCommentAutomation200ResponseAutomation) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *GetCommentAutomation200ResponseAutomation) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *GetCommentAutomation200ResponseAutomation) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetAccountId

`func (o *GetCommentAutomation200ResponseAutomation) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetCommentAutomation200ResponseAutomation) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetCommentAutomation200ResponseAutomation) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetCommentAutomation200ResponseAutomation) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *GetCommentAutomation200ResponseAutomation) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *GetCommentAutomation200ResponseAutomation) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *GetCommentAutomation200ResponseAutomation) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *GetCommentAutomation200ResponseAutomation) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetPostId

`func (o *GetCommentAutomation200ResponseAutomation) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *GetCommentAutomation200ResponseAutomation) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *GetCommentAutomation200ResponseAutomation) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *GetCommentAutomation200ResponseAutomation) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetPostTitle

`func (o *GetCommentAutomation200ResponseAutomation) GetPostTitle() string`

GetPostTitle returns the PostTitle field if non-nil, zero value otherwise.

### GetPostTitleOk

`func (o *GetCommentAutomation200ResponseAutomation) GetPostTitleOk() (*string, bool)`

GetPostTitleOk returns a tuple with the PostTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTitle

`func (o *GetCommentAutomation200ResponseAutomation) SetPostTitle(v string)`

SetPostTitle sets PostTitle field to given value.

### HasPostTitle

`func (o *GetCommentAutomation200ResponseAutomation) HasPostTitle() bool`

HasPostTitle returns a boolean if a field has been set.

### GetKeywords

`func (o *GetCommentAutomation200ResponseAutomation) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *GetCommentAutomation200ResponseAutomation) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *GetCommentAutomation200ResponseAutomation) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *GetCommentAutomation200ResponseAutomation) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMatchMode

`func (o *GetCommentAutomation200ResponseAutomation) GetMatchMode() string`

GetMatchMode returns the MatchMode field if non-nil, zero value otherwise.

### GetMatchModeOk

`func (o *GetCommentAutomation200ResponseAutomation) GetMatchModeOk() (*string, bool)`

GetMatchModeOk returns a tuple with the MatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMode

`func (o *GetCommentAutomation200ResponseAutomation) SetMatchMode(v string)`

SetMatchMode sets MatchMode field to given value.

### HasMatchMode

`func (o *GetCommentAutomation200ResponseAutomation) HasMatchMode() bool`

HasMatchMode returns a boolean if a field has been set.

### GetDmMessage

`func (o *GetCommentAutomation200ResponseAutomation) GetDmMessage() string`

GetDmMessage returns the DmMessage field if non-nil, zero value otherwise.

### GetDmMessageOk

`func (o *GetCommentAutomation200ResponseAutomation) GetDmMessageOk() (*string, bool)`

GetDmMessageOk returns a tuple with the DmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmMessage

`func (o *GetCommentAutomation200ResponseAutomation) SetDmMessage(v string)`

SetDmMessage sets DmMessage field to given value.

### HasDmMessage

`func (o *GetCommentAutomation200ResponseAutomation) HasDmMessage() bool`

HasDmMessage returns a boolean if a field has been set.

### GetButtons

`func (o *GetCommentAutomation200ResponseAutomation) GetButtons() []DmButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *GetCommentAutomation200ResponseAutomation) GetButtonsOk() (*[]DmButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *GetCommentAutomation200ResponseAutomation) SetButtons(v []DmButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *GetCommentAutomation200ResponseAutomation) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetCommentReply

`func (o *GetCommentAutomation200ResponseAutomation) GetCommentReply() string`

GetCommentReply returns the CommentReply field if non-nil, zero value otherwise.

### GetCommentReplyOk

`func (o *GetCommentAutomation200ResponseAutomation) GetCommentReplyOk() (*string, bool)`

GetCommentReplyOk returns a tuple with the CommentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReply

`func (o *GetCommentAutomation200ResponseAutomation) SetCommentReply(v string)`

SetCommentReply sets CommentReply field to given value.

### HasCommentReply

`func (o *GetCommentAutomation200ResponseAutomation) HasCommentReply() bool`

HasCommentReply returns a boolean if a field has been set.

### GetLinkTracking

`func (o *GetCommentAutomation200ResponseAutomation) GetLinkTracking() bool`

GetLinkTracking returns the LinkTracking field if non-nil, zero value otherwise.

### GetLinkTrackingOk

`func (o *GetCommentAutomation200ResponseAutomation) GetLinkTrackingOk() (*bool, bool)`

GetLinkTrackingOk returns a tuple with the LinkTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTracking

`func (o *GetCommentAutomation200ResponseAutomation) SetLinkTracking(v bool)`

SetLinkTracking sets LinkTracking field to given value.

### HasLinkTracking

`func (o *GetCommentAutomation200ResponseAutomation) HasLinkTracking() bool`

HasLinkTracking returns a boolean if a field has been set.

### GetClickTag

`func (o *GetCommentAutomation200ResponseAutomation) GetClickTag() string`

GetClickTag returns the ClickTag field if non-nil, zero value otherwise.

### GetClickTagOk

`func (o *GetCommentAutomation200ResponseAutomation) GetClickTagOk() (*string, bool)`

GetClickTagOk returns a tuple with the ClickTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTag

`func (o *GetCommentAutomation200ResponseAutomation) SetClickTag(v string)`

SetClickTag sets ClickTag field to given value.

### HasClickTag

`func (o *GetCommentAutomation200ResponseAutomation) HasClickTag() bool`

HasClickTag returns a boolean if a field has been set.

### GetIsActive

`func (o *GetCommentAutomation200ResponseAutomation) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GetCommentAutomation200ResponseAutomation) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GetCommentAutomation200ResponseAutomation) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *GetCommentAutomation200ResponseAutomation) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetStats

`func (o *GetCommentAutomation200ResponseAutomation) GetStats() CreateCommentAutomation200ResponseAutomationStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetCommentAutomation200ResponseAutomation) GetStatsOk() (*CreateCommentAutomation200ResponseAutomationStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetCommentAutomation200ResponseAutomation) SetStats(v CreateCommentAutomation200ResponseAutomationStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetCommentAutomation200ResponseAutomation) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetCommentAutomation200ResponseAutomation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCommentAutomation200ResponseAutomation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCommentAutomation200ResponseAutomation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetCommentAutomation200ResponseAutomation) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GetCommentAutomation200ResponseAutomation) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetCommentAutomation200ResponseAutomation) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetCommentAutomation200ResponseAutomation) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GetCommentAutomation200ResponseAutomation) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


