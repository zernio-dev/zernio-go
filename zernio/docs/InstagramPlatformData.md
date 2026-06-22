# InstagramPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | Set to &#39;story&#39; to publish as a Story. Default posts become Reels or feed depending on media. | [optional] 
**ShareToFeed** | Pointer to **bool** | For Reels only. When true (default), the Reel appears on both the Reels tab and your main profile feed. Set to false to post to the Reels tab only. | [optional] [default to true]
**Collaborators** | Pointer to **[]string** | Up to 3 Instagram usernames to invite as collaborators (feed/Reels only) | [optional] 
**FirstComment** | Pointer to **string** | Optional first comment to add after the post is created (not applied to Stories) | [optional] 
**TrialParams** | Pointer to [**InstagramPlatformDataTrialParams**](InstagramPlatformDataTrialParams.md) |  | [optional] 
**UserTags** | Pointer to [**[]InstagramPlatformDataUserTagsInner**](InstagramPlatformDataUserTagsInner.md) | Tag Instagram users in photos by username and position. Not supported for stories or videos. For carousels, use mediaIndex to target specific slides (defaults to 0). Tags on video items are silently skipped. | [optional] 
**AudioName** | Pointer to **string** | Custom name for original audio in Reels. Replaces the default \&quot;Original Audio\&quot; label. Can only be set once. | [optional] 
**ThumbOffset** | Pointer to **int32** | Millisecond offset from video start for the Reel cover frame. Ignored when instagramThumbnail or reelCover is provided. Defaults to 0. | [optional] 
**InstagramThumbnail** | Pointer to **string** | Custom cover image URL for Instagram Reels (JPG or PNG, publicly accessible). Overrides thumbOffset when provided. Also accepted as reelCover (alias). | [optional] 
**ReelCover** | Pointer to **string** | Alias for instagramThumbnail. If both are provided, instagramThumbnail takes priority. | [optional] 

## Methods

### NewInstagramPlatformData

`func NewInstagramPlatformData() *InstagramPlatformData`

NewInstagramPlatformData instantiates a new InstagramPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstagramPlatformDataWithDefaults

`func NewInstagramPlatformDataWithDefaults() *InstagramPlatformData`

NewInstagramPlatformDataWithDefaults instantiates a new InstagramPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *InstagramPlatformData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *InstagramPlatformData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *InstagramPlatformData) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *InstagramPlatformData) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetShareToFeed

`func (o *InstagramPlatformData) GetShareToFeed() bool`

GetShareToFeed returns the ShareToFeed field if non-nil, zero value otherwise.

### GetShareToFeedOk

`func (o *InstagramPlatformData) GetShareToFeedOk() (*bool, bool)`

GetShareToFeedOk returns a tuple with the ShareToFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToFeed

`func (o *InstagramPlatformData) SetShareToFeed(v bool)`

SetShareToFeed sets ShareToFeed field to given value.

### HasShareToFeed

`func (o *InstagramPlatformData) HasShareToFeed() bool`

HasShareToFeed returns a boolean if a field has been set.

### GetCollaborators

`func (o *InstagramPlatformData) GetCollaborators() []string`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *InstagramPlatformData) GetCollaboratorsOk() (*[]string, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *InstagramPlatformData) SetCollaborators(v []string)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *InstagramPlatformData) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### GetFirstComment

`func (o *InstagramPlatformData) GetFirstComment() string`

GetFirstComment returns the FirstComment field if non-nil, zero value otherwise.

### GetFirstCommentOk

`func (o *InstagramPlatformData) GetFirstCommentOk() (*string, bool)`

GetFirstCommentOk returns a tuple with the FirstComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstComment

`func (o *InstagramPlatformData) SetFirstComment(v string)`

SetFirstComment sets FirstComment field to given value.

### HasFirstComment

`func (o *InstagramPlatformData) HasFirstComment() bool`

HasFirstComment returns a boolean if a field has been set.

### GetTrialParams

`func (o *InstagramPlatformData) GetTrialParams() InstagramPlatformDataTrialParams`

GetTrialParams returns the TrialParams field if non-nil, zero value otherwise.

### GetTrialParamsOk

`func (o *InstagramPlatformData) GetTrialParamsOk() (*InstagramPlatformDataTrialParams, bool)`

GetTrialParamsOk returns a tuple with the TrialParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialParams

`func (o *InstagramPlatformData) SetTrialParams(v InstagramPlatformDataTrialParams)`

SetTrialParams sets TrialParams field to given value.

### HasTrialParams

`func (o *InstagramPlatformData) HasTrialParams() bool`

HasTrialParams returns a boolean if a field has been set.

### GetUserTags

`func (o *InstagramPlatformData) GetUserTags() []InstagramPlatformDataUserTagsInner`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *InstagramPlatformData) GetUserTagsOk() (*[]InstagramPlatformDataUserTagsInner, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *InstagramPlatformData) SetUserTags(v []InstagramPlatformDataUserTagsInner)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *InstagramPlatformData) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### GetAudioName

`func (o *InstagramPlatformData) GetAudioName() string`

GetAudioName returns the AudioName field if non-nil, zero value otherwise.

### GetAudioNameOk

`func (o *InstagramPlatformData) GetAudioNameOk() (*string, bool)`

GetAudioNameOk returns a tuple with the AudioName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioName

`func (o *InstagramPlatformData) SetAudioName(v string)`

SetAudioName sets AudioName field to given value.

### HasAudioName

`func (o *InstagramPlatformData) HasAudioName() bool`

HasAudioName returns a boolean if a field has been set.

### GetThumbOffset

`func (o *InstagramPlatformData) GetThumbOffset() int32`

GetThumbOffset returns the ThumbOffset field if non-nil, zero value otherwise.

### GetThumbOffsetOk

`func (o *InstagramPlatformData) GetThumbOffsetOk() (*int32, bool)`

GetThumbOffsetOk returns a tuple with the ThumbOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbOffset

`func (o *InstagramPlatformData) SetThumbOffset(v int32)`

SetThumbOffset sets ThumbOffset field to given value.

### HasThumbOffset

`func (o *InstagramPlatformData) HasThumbOffset() bool`

HasThumbOffset returns a boolean if a field has been set.

### GetInstagramThumbnail

`func (o *InstagramPlatformData) GetInstagramThumbnail() string`

GetInstagramThumbnail returns the InstagramThumbnail field if non-nil, zero value otherwise.

### GetInstagramThumbnailOk

`func (o *InstagramPlatformData) GetInstagramThumbnailOk() (*string, bool)`

GetInstagramThumbnailOk returns a tuple with the InstagramThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramThumbnail

`func (o *InstagramPlatformData) SetInstagramThumbnail(v string)`

SetInstagramThumbnail sets InstagramThumbnail field to given value.

### HasInstagramThumbnail

`func (o *InstagramPlatformData) HasInstagramThumbnail() bool`

HasInstagramThumbnail returns a boolean if a field has been set.

### GetReelCover

`func (o *InstagramPlatformData) GetReelCover() string`

GetReelCover returns the ReelCover field if non-nil, zero value otherwise.

### GetReelCoverOk

`func (o *InstagramPlatformData) GetReelCoverOk() (*string, bool)`

GetReelCoverOk returns a tuple with the ReelCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReelCover

`func (o *InstagramPlatformData) SetReelCover(v string)`

SetReelCover sets ReelCover field to given value.

### HasReelCover

`func (o *InstagramPlatformData) HasReelCover() bool`

HasReelCover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


