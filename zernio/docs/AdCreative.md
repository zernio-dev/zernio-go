# AdCreative

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ThumbnailUrl** | Pointer to **string** | Primary thumbnail/image URL | [optional] 
**ImageUrl** | Pointer to **string** | Alternative image URL | [optional] 
**VideoId** | Pointer to **NullableString** | Meta video ID for VIDEO-type ads. Null for non-video ads. Callers that need an embeddable MP4 can call GET /{videoId}?fields&#x3D;source with the page access token. | [optional] 
**VideoUrl** | Pointer to **NullableString** | Public Facebook watch URL for VIDEO-type ads (https://www.facebook.com/watch/?v&#x3D;{videoId}). Null for non-video ads. | [optional] 
**ObjectType** | Pointer to **string** | Meta creative object_type (e.g. SHARE, VIDEO, PRIVACY_CHECK_FAIL, POST_DELETED). Use this to render state-aware previews — when Meta moderation strips image/video fields, only thumbnailUrl at 64x64 is available. | [optional] 
**ObjectStoryId** | Pointer to **NullableString** | Meta creative &#x60;object_story_id&#x60; (the SHARE reference). Frequently absent — Meta omits it for SHARE creatives. Use effectiveObjectStoryId instead. | [optional] 
**EffectiveObjectStoryId** | Pointer to **NullableString** | Meta &#x60;effective_object_story_id&#x60; — &#x60;{pageId}_{postId}&#x60; of the Facebook post the ad&#39;s engagement (comments) lives on. Pass to GET /v1/ads?effectiveObjectStoryId&#x3D; to map a Business-Manager-visible post back to this ad; GET /v1/ads/{adId}/comments resolves comments against it. | [optional] 
**EffectiveInstagramMediaId** | Pointer to **NullableString** | Meta &#x60;effective_instagram_media_id&#x60; — the Instagram media ID of the boosted post the ad&#39;s engagement lives on. Pass to GET /v1/ads?effectiveInstagramMediaId&#x3D; to map a Business-Manager-visible IG post back to this ad. | [optional] 
**InstagramUserId** | Pointer to **NullableString** | Meta &#x60;instagram_user_id&#x60; — the Instagram-scoped business ID that owns the boosted media. | [optional] 
**InstagramPermalinkUrl** | Pointer to **NullableString** | Meta &#x60;instagram_permalink_url&#x60; — public Instagram post URL of the boosted media. | [optional] 
**MediaUrls** | Pointer to **[]string** | All media URLs for this ad (carousel images, multiple assets). Populated for Meta (carousel child_attachments), Google Ads (responsive display marketing_images), and LinkedIn (multi-image posts). | [optional] 
**Body** | Pointer to **string** | Ad copy/text | [optional] 
**GoogleHeadline** | Pointer to **string** | Google Ads headline | [optional] 
**GoogleDescription** | Pointer to **string** | Google Ads description | [optional] 
**LinkUrl** | Pointer to **string** | Destination URL | [optional] 
**PinterestImageUrl** | Pointer to **string** |  | [optional] 
**PinterestTitle** | Pointer to **string** |  | [optional] 
**PinterestDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewAdCreative

`func NewAdCreative() *AdCreative`

NewAdCreative instantiates a new AdCreative object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdCreativeWithDefaults

`func NewAdCreativeWithDefaults() *AdCreative`

NewAdCreativeWithDefaults instantiates a new AdCreative object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetThumbnailUrl

`func (o *AdCreative) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *AdCreative) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *AdCreative) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *AdCreative) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### GetImageUrl

`func (o *AdCreative) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *AdCreative) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *AdCreative) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *AdCreative) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### GetVideoId

`func (o *AdCreative) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *AdCreative) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *AdCreative) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *AdCreative) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### SetVideoIdNil

`func (o *AdCreative) SetVideoIdNil(b bool)`

 SetVideoIdNil sets the value for VideoId to be an explicit nil

### UnsetVideoId
`func (o *AdCreative) UnsetVideoId()`

UnsetVideoId ensures that no value is present for VideoId, not even an explicit nil
### GetVideoUrl

`func (o *AdCreative) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *AdCreative) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *AdCreative) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *AdCreative) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *AdCreative) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *AdCreative) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetObjectType

`func (o *AdCreative) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *AdCreative) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *AdCreative) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.

### HasObjectType

`func (o *AdCreative) HasObjectType() bool`

HasObjectType returns a boolean if a field has been set.

### GetObjectStoryId

`func (o *AdCreative) GetObjectStoryId() string`

GetObjectStoryId returns the ObjectStoryId field if non-nil, zero value otherwise.

### GetObjectStoryIdOk

`func (o *AdCreative) GetObjectStoryIdOk() (*string, bool)`

GetObjectStoryIdOk returns a tuple with the ObjectStoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectStoryId

`func (o *AdCreative) SetObjectStoryId(v string)`

SetObjectStoryId sets ObjectStoryId field to given value.

### HasObjectStoryId

`func (o *AdCreative) HasObjectStoryId() bool`

HasObjectStoryId returns a boolean if a field has been set.

### SetObjectStoryIdNil

`func (o *AdCreative) SetObjectStoryIdNil(b bool)`

 SetObjectStoryIdNil sets the value for ObjectStoryId to be an explicit nil

### UnsetObjectStoryId
`func (o *AdCreative) UnsetObjectStoryId()`

UnsetObjectStoryId ensures that no value is present for ObjectStoryId, not even an explicit nil
### GetEffectiveObjectStoryId

`func (o *AdCreative) GetEffectiveObjectStoryId() string`

GetEffectiveObjectStoryId returns the EffectiveObjectStoryId field if non-nil, zero value otherwise.

### GetEffectiveObjectStoryIdOk

`func (o *AdCreative) GetEffectiveObjectStoryIdOk() (*string, bool)`

GetEffectiveObjectStoryIdOk returns a tuple with the EffectiveObjectStoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveObjectStoryId

`func (o *AdCreative) SetEffectiveObjectStoryId(v string)`

SetEffectiveObjectStoryId sets EffectiveObjectStoryId field to given value.

### HasEffectiveObjectStoryId

`func (o *AdCreative) HasEffectiveObjectStoryId() bool`

HasEffectiveObjectStoryId returns a boolean if a field has been set.

### SetEffectiveObjectStoryIdNil

`func (o *AdCreative) SetEffectiveObjectStoryIdNil(b bool)`

 SetEffectiveObjectStoryIdNil sets the value for EffectiveObjectStoryId to be an explicit nil

### UnsetEffectiveObjectStoryId
`func (o *AdCreative) UnsetEffectiveObjectStoryId()`

UnsetEffectiveObjectStoryId ensures that no value is present for EffectiveObjectStoryId, not even an explicit nil
### GetEffectiveInstagramMediaId

`func (o *AdCreative) GetEffectiveInstagramMediaId() string`

GetEffectiveInstagramMediaId returns the EffectiveInstagramMediaId field if non-nil, zero value otherwise.

### GetEffectiveInstagramMediaIdOk

`func (o *AdCreative) GetEffectiveInstagramMediaIdOk() (*string, bool)`

GetEffectiveInstagramMediaIdOk returns a tuple with the EffectiveInstagramMediaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveInstagramMediaId

`func (o *AdCreative) SetEffectiveInstagramMediaId(v string)`

SetEffectiveInstagramMediaId sets EffectiveInstagramMediaId field to given value.

### HasEffectiveInstagramMediaId

`func (o *AdCreative) HasEffectiveInstagramMediaId() bool`

HasEffectiveInstagramMediaId returns a boolean if a field has been set.

### SetEffectiveInstagramMediaIdNil

`func (o *AdCreative) SetEffectiveInstagramMediaIdNil(b bool)`

 SetEffectiveInstagramMediaIdNil sets the value for EffectiveInstagramMediaId to be an explicit nil

### UnsetEffectiveInstagramMediaId
`func (o *AdCreative) UnsetEffectiveInstagramMediaId()`

UnsetEffectiveInstagramMediaId ensures that no value is present for EffectiveInstagramMediaId, not even an explicit nil
### GetInstagramUserId

`func (o *AdCreative) GetInstagramUserId() string`

GetInstagramUserId returns the InstagramUserId field if non-nil, zero value otherwise.

### GetInstagramUserIdOk

`func (o *AdCreative) GetInstagramUserIdOk() (*string, bool)`

GetInstagramUserIdOk returns a tuple with the InstagramUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUserId

`func (o *AdCreative) SetInstagramUserId(v string)`

SetInstagramUserId sets InstagramUserId field to given value.

### HasInstagramUserId

`func (o *AdCreative) HasInstagramUserId() bool`

HasInstagramUserId returns a boolean if a field has been set.

### SetInstagramUserIdNil

`func (o *AdCreative) SetInstagramUserIdNil(b bool)`

 SetInstagramUserIdNil sets the value for InstagramUserId to be an explicit nil

### UnsetInstagramUserId
`func (o *AdCreative) UnsetInstagramUserId()`

UnsetInstagramUserId ensures that no value is present for InstagramUserId, not even an explicit nil
### GetInstagramPermalinkUrl

`func (o *AdCreative) GetInstagramPermalinkUrl() string`

GetInstagramPermalinkUrl returns the InstagramPermalinkUrl field if non-nil, zero value otherwise.

### GetInstagramPermalinkUrlOk

`func (o *AdCreative) GetInstagramPermalinkUrlOk() (*string, bool)`

GetInstagramPermalinkUrlOk returns a tuple with the InstagramPermalinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramPermalinkUrl

`func (o *AdCreative) SetInstagramPermalinkUrl(v string)`

SetInstagramPermalinkUrl sets InstagramPermalinkUrl field to given value.

### HasInstagramPermalinkUrl

`func (o *AdCreative) HasInstagramPermalinkUrl() bool`

HasInstagramPermalinkUrl returns a boolean if a field has been set.

### SetInstagramPermalinkUrlNil

`func (o *AdCreative) SetInstagramPermalinkUrlNil(b bool)`

 SetInstagramPermalinkUrlNil sets the value for InstagramPermalinkUrl to be an explicit nil

### UnsetInstagramPermalinkUrl
`func (o *AdCreative) UnsetInstagramPermalinkUrl()`

UnsetInstagramPermalinkUrl ensures that no value is present for InstagramPermalinkUrl, not even an explicit nil
### GetMediaUrls

`func (o *AdCreative) GetMediaUrls() []string`

GetMediaUrls returns the MediaUrls field if non-nil, zero value otherwise.

### GetMediaUrlsOk

`func (o *AdCreative) GetMediaUrlsOk() (*[]string, bool)`

GetMediaUrlsOk returns a tuple with the MediaUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaUrls

`func (o *AdCreative) SetMediaUrls(v []string)`

SetMediaUrls sets MediaUrls field to given value.

### HasMediaUrls

`func (o *AdCreative) HasMediaUrls() bool`

HasMediaUrls returns a boolean if a field has been set.

### GetBody

`func (o *AdCreative) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *AdCreative) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *AdCreative) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *AdCreative) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetGoogleHeadline

`func (o *AdCreative) GetGoogleHeadline() string`

GetGoogleHeadline returns the GoogleHeadline field if non-nil, zero value otherwise.

### GetGoogleHeadlineOk

`func (o *AdCreative) GetGoogleHeadlineOk() (*string, bool)`

GetGoogleHeadlineOk returns a tuple with the GoogleHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleHeadline

`func (o *AdCreative) SetGoogleHeadline(v string)`

SetGoogleHeadline sets GoogleHeadline field to given value.

### HasGoogleHeadline

`func (o *AdCreative) HasGoogleHeadline() bool`

HasGoogleHeadline returns a boolean if a field has been set.

### GetGoogleDescription

`func (o *AdCreative) GetGoogleDescription() string`

GetGoogleDescription returns the GoogleDescription field if non-nil, zero value otherwise.

### GetGoogleDescriptionOk

`func (o *AdCreative) GetGoogleDescriptionOk() (*string, bool)`

GetGoogleDescriptionOk returns a tuple with the GoogleDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoogleDescription

`func (o *AdCreative) SetGoogleDescription(v string)`

SetGoogleDescription sets GoogleDescription field to given value.

### HasGoogleDescription

`func (o *AdCreative) HasGoogleDescription() bool`

HasGoogleDescription returns a boolean if a field has been set.

### GetLinkUrl

`func (o *AdCreative) GetLinkUrl() string`

GetLinkUrl returns the LinkUrl field if non-nil, zero value otherwise.

### GetLinkUrlOk

`func (o *AdCreative) GetLinkUrlOk() (*string, bool)`

GetLinkUrlOk returns a tuple with the LinkUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkUrl

`func (o *AdCreative) SetLinkUrl(v string)`

SetLinkUrl sets LinkUrl field to given value.

### HasLinkUrl

`func (o *AdCreative) HasLinkUrl() bool`

HasLinkUrl returns a boolean if a field has been set.

### GetPinterestImageUrl

`func (o *AdCreative) GetPinterestImageUrl() string`

GetPinterestImageUrl returns the PinterestImageUrl field if non-nil, zero value otherwise.

### GetPinterestImageUrlOk

`func (o *AdCreative) GetPinterestImageUrlOk() (*string, bool)`

GetPinterestImageUrlOk returns a tuple with the PinterestImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinterestImageUrl

`func (o *AdCreative) SetPinterestImageUrl(v string)`

SetPinterestImageUrl sets PinterestImageUrl field to given value.

### HasPinterestImageUrl

`func (o *AdCreative) HasPinterestImageUrl() bool`

HasPinterestImageUrl returns a boolean if a field has been set.

### GetPinterestTitle

`func (o *AdCreative) GetPinterestTitle() string`

GetPinterestTitle returns the PinterestTitle field if non-nil, zero value otherwise.

### GetPinterestTitleOk

`func (o *AdCreative) GetPinterestTitleOk() (*string, bool)`

GetPinterestTitleOk returns a tuple with the PinterestTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinterestTitle

`func (o *AdCreative) SetPinterestTitle(v string)`

SetPinterestTitle sets PinterestTitle field to given value.

### HasPinterestTitle

`func (o *AdCreative) HasPinterestTitle() bool`

HasPinterestTitle returns a boolean if a field has been set.

### GetPinterestDescription

`func (o *AdCreative) GetPinterestDescription() string`

GetPinterestDescription returns the PinterestDescription field if non-nil, zero value otherwise.

### GetPinterestDescriptionOk

`func (o *AdCreative) GetPinterestDescriptionOk() (*string, bool)`

GetPinterestDescriptionOk returns a tuple with the PinterestDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinterestDescription

`func (o *AdCreative) SetPinterestDescription(v string)`

SetPinterestDescription sets PinterestDescription field to given value.

### HasPinterestDescription

`func (o *AdCreative) HasPinterestDescription() bool`

HasPinterestDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


