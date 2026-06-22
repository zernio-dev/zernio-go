# CreateCommentAutomationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**AccountId** | **string** | Instagram or Facebook account ID | 
**Trigger** | Pointer to **string** | What fires the automation. &#39;comment&#39; (keyword comment on a post) or &#39;story_reply&#39; (keyword reply to an Instagram story). For &#39;story_reply&#39;, platformPostId is the story media id (omit for any story). | [optional] [default to "comment"]
**PlatformPostId** | Pointer to **string** | Platform media/post ID (or story media id when trigger&#x3D;story_reply). Omit for an account-wide (any-post / any-story) automation. | [optional] 
**PostId** | Pointer to **string** | Zernio post ID. Required only when also targeting a specific post via platformPostId. | [optional] 
**PostTitle** | Pointer to **string** | Post content snippet for display | [optional] 
**Name** | **string** | Automation label | 
**Keywords** | Pointer to **[]string** | Trigger keywords (empty &#x3D; any comment triggers) | [optional] 
**MatchMode** | Pointer to **string** |  | [optional] [default to "contains"]
**DmMessage** | **string** | DM text to send to commenter. Max 640 chars when buttons are set, otherwise ~1000. | 
**Buttons** | Pointer to [**[]DmButton**](DmButton.md) | Optional inline DM buttons (1-3). Phone buttons are Facebook-only. Omit or pass [] for a plain-text DM. | [optional] 
**CommentReply** | Pointer to **string** | Optional public reply to the comment | [optional] 
**LinkTracking** | Pointer to **bool** | Wrap link buttons in the DM in a tracked redirect so clicks are counted (Link Clicks / CTR). Pass false to send links exactly as written. Defaults to on. | [optional] [default to true]
**ClickTag** | Pointer to **string** | Optional tag applied to a contact when they click a tracked link (requires linkTracking). Lets you segment clickers for broadcasts/sequences. | [optional] 

## Methods

### NewCreateCommentAutomationRequest

`func NewCreateCommentAutomationRequest(profileId string, accountId string, name string, dmMessage string, ) *CreateCommentAutomationRequest`

NewCreateCommentAutomationRequest instantiates a new CreateCommentAutomationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentAutomationRequestWithDefaults

`func NewCreateCommentAutomationRequestWithDefaults() *CreateCommentAutomationRequest`

NewCreateCommentAutomationRequestWithDefaults instantiates a new CreateCommentAutomationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *CreateCommentAutomationRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *CreateCommentAutomationRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *CreateCommentAutomationRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *CreateCommentAutomationRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateCommentAutomationRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateCommentAutomationRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTrigger

`func (o *CreateCommentAutomationRequest) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *CreateCommentAutomationRequest) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *CreateCommentAutomationRequest) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *CreateCommentAutomationRequest) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetPlatformPostId

`func (o *CreateCommentAutomationRequest) GetPlatformPostId() string`

GetPlatformPostId returns the PlatformPostId field if non-nil, zero value otherwise.

### GetPlatformPostIdOk

`func (o *CreateCommentAutomationRequest) GetPlatformPostIdOk() (*string, bool)`

GetPlatformPostIdOk returns a tuple with the PlatformPostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformPostId

`func (o *CreateCommentAutomationRequest) SetPlatformPostId(v string)`

SetPlatformPostId sets PlatformPostId field to given value.

### HasPlatformPostId

`func (o *CreateCommentAutomationRequest) HasPlatformPostId() bool`

HasPlatformPostId returns a boolean if a field has been set.

### GetPostId

`func (o *CreateCommentAutomationRequest) GetPostId() string`

GetPostId returns the PostId field if non-nil, zero value otherwise.

### GetPostIdOk

`func (o *CreateCommentAutomationRequest) GetPostIdOk() (*string, bool)`

GetPostIdOk returns a tuple with the PostId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostId

`func (o *CreateCommentAutomationRequest) SetPostId(v string)`

SetPostId sets PostId field to given value.

### HasPostId

`func (o *CreateCommentAutomationRequest) HasPostId() bool`

HasPostId returns a boolean if a field has been set.

### GetPostTitle

`func (o *CreateCommentAutomationRequest) GetPostTitle() string`

GetPostTitle returns the PostTitle field if non-nil, zero value otherwise.

### GetPostTitleOk

`func (o *CreateCommentAutomationRequest) GetPostTitleOk() (*string, bool)`

GetPostTitleOk returns a tuple with the PostTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostTitle

`func (o *CreateCommentAutomationRequest) SetPostTitle(v string)`

SetPostTitle sets PostTitle field to given value.

### HasPostTitle

`func (o *CreateCommentAutomationRequest) HasPostTitle() bool`

HasPostTitle returns a boolean if a field has been set.

### GetName

`func (o *CreateCommentAutomationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCommentAutomationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCommentAutomationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetKeywords

`func (o *CreateCommentAutomationRequest) GetKeywords() []string`

GetKeywords returns the Keywords field if non-nil, zero value otherwise.

### GetKeywordsOk

`func (o *CreateCommentAutomationRequest) GetKeywordsOk() (*[]string, bool)`

GetKeywordsOk returns a tuple with the Keywords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeywords

`func (o *CreateCommentAutomationRequest) SetKeywords(v []string)`

SetKeywords sets Keywords field to given value.

### HasKeywords

`func (o *CreateCommentAutomationRequest) HasKeywords() bool`

HasKeywords returns a boolean if a field has been set.

### GetMatchMode

`func (o *CreateCommentAutomationRequest) GetMatchMode() string`

GetMatchMode returns the MatchMode field if non-nil, zero value otherwise.

### GetMatchModeOk

`func (o *CreateCommentAutomationRequest) GetMatchModeOk() (*string, bool)`

GetMatchModeOk returns a tuple with the MatchMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMode

`func (o *CreateCommentAutomationRequest) SetMatchMode(v string)`

SetMatchMode sets MatchMode field to given value.

### HasMatchMode

`func (o *CreateCommentAutomationRequest) HasMatchMode() bool`

HasMatchMode returns a boolean if a field has been set.

### GetDmMessage

`func (o *CreateCommentAutomationRequest) GetDmMessage() string`

GetDmMessage returns the DmMessage field if non-nil, zero value otherwise.

### GetDmMessageOk

`func (o *CreateCommentAutomationRequest) GetDmMessageOk() (*string, bool)`

GetDmMessageOk returns a tuple with the DmMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDmMessage

`func (o *CreateCommentAutomationRequest) SetDmMessage(v string)`

SetDmMessage sets DmMessage field to given value.


### GetButtons

`func (o *CreateCommentAutomationRequest) GetButtons() []DmButton`

GetButtons returns the Buttons field if non-nil, zero value otherwise.

### GetButtonsOk

`func (o *CreateCommentAutomationRequest) GetButtonsOk() (*[]DmButton, bool)`

GetButtonsOk returns a tuple with the Buttons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtons

`func (o *CreateCommentAutomationRequest) SetButtons(v []DmButton)`

SetButtons sets Buttons field to given value.

### HasButtons

`func (o *CreateCommentAutomationRequest) HasButtons() bool`

HasButtons returns a boolean if a field has been set.

### GetCommentReply

`func (o *CreateCommentAutomationRequest) GetCommentReply() string`

GetCommentReply returns the CommentReply field if non-nil, zero value otherwise.

### GetCommentReplyOk

`func (o *CreateCommentAutomationRequest) GetCommentReplyOk() (*string, bool)`

GetCommentReplyOk returns a tuple with the CommentReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReply

`func (o *CreateCommentAutomationRequest) SetCommentReply(v string)`

SetCommentReply sets CommentReply field to given value.

### HasCommentReply

`func (o *CreateCommentAutomationRequest) HasCommentReply() bool`

HasCommentReply returns a boolean if a field has been set.

### GetLinkTracking

`func (o *CreateCommentAutomationRequest) GetLinkTracking() bool`

GetLinkTracking returns the LinkTracking field if non-nil, zero value otherwise.

### GetLinkTrackingOk

`func (o *CreateCommentAutomationRequest) GetLinkTrackingOk() (*bool, bool)`

GetLinkTrackingOk returns a tuple with the LinkTracking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkTracking

`func (o *CreateCommentAutomationRequest) SetLinkTracking(v bool)`

SetLinkTracking sets LinkTracking field to given value.

### HasLinkTracking

`func (o *CreateCommentAutomationRequest) HasLinkTracking() bool`

HasLinkTracking returns a boolean if a field has been set.

### GetClickTag

`func (o *CreateCommentAutomationRequest) GetClickTag() string`

GetClickTag returns the ClickTag field if non-nil, zero value otherwise.

### GetClickTagOk

`func (o *CreateCommentAutomationRequest) GetClickTagOk() (*string, bool)`

GetClickTagOk returns a tuple with the ClickTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClickTag

`func (o *CreateCommentAutomationRequest) SetClickTag(v string)`

SetClickTag sets ClickTag field to given value.

### HasClickTag

`func (o *CreateCommentAutomationRequest) HasClickTag() bool`

HasClickTag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


