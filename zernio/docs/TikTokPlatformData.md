# TikTokPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Draft** | Pointer to **bool** | When true, sends the post to the TikTok Creator Inbox as a draft instead of publishing immediately. The creator receives an inbox notification to complete posting via TikTok&#39;s editing flow. Maps to TikTok API post_mode: \&quot;MEDIA_UPLOAD\&quot; (photos) or the dedicated inbox endpoint (videos). When false or omitted, publishes directly via post_mode: \&quot;DIRECT_POST\&quot;. Note: publish_type is not a supported field. Use this field instead.  | [optional] 
**PrivacyLevel** | Pointer to **string** | One of the values returned by the TikTok creator info API for the account | [optional] 
**AllowComment** | Pointer to **bool** | Allow comments on the post | [optional] 
**AllowDuet** | Pointer to **bool** | Allow duets (required for video posts) | [optional] 
**AllowStitch** | Pointer to **bool** | Allow stitches (required for video posts) | [optional] 
**CommercialContentType** | Pointer to **string** | Type of commercial content disclosure | [optional] 
**BrandPartnerPromote** | Pointer to **bool** | Whether the post promotes a brand partner | [optional] 
**IsBrandOrganicPost** | Pointer to **bool** | Whether the post is a brand organic post | [optional] 
**ContentPreviewConfirmed** | Pointer to **bool** | User has confirmed they previewed the content | [optional] 
**ExpressConsentGiven** | Pointer to **bool** | User has given express consent for posting | [optional] 
**MediaType** | Pointer to **string** | Optional override. Defaults based on provided media items. | [optional] 
**VideoCoverTimestampMs** | Pointer to **int32** | Optional for video posts. Timestamp in milliseconds to select which frame to use as thumbnail (defaults to 1000ms/1 second). Ignored when videoCoverImageUrl is provided. | [optional] 
**VideoCoverImageUrl** | Pointer to **string** | Optional for video posts. URL of a custom thumbnail image (JPG, PNG, or WebP, max 20MB). The image is stitched as a single frame at the start of the video and used as the cover. Overrides videoCoverTimestampMs when provided. | [optional] 
**PhotoCoverIndex** | Pointer to **int32** | Optional for photo carousels. Index of image to use as cover, 0-based (defaults to 0/first image). | [optional] 
**AutoAddMusic** | Pointer to **bool** | When true, TikTok may add recommended music (photos only) | [optional] 
**VideoMadeWithAi** | Pointer to **bool** | Set true to disclose AI-generated content | [optional] 
**Description** | Pointer to **string** | Optional long-form description for photo posts (max 4000 chars). Recommended when content exceeds 90 chars, as photo titles are auto-truncated. | [optional] 

## Methods

### NewTikTokPlatformData

`func NewTikTokPlatformData() *TikTokPlatformData`

NewTikTokPlatformData instantiates a new TikTokPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTikTokPlatformDataWithDefaults

`func NewTikTokPlatformDataWithDefaults() *TikTokPlatformData`

NewTikTokPlatformDataWithDefaults instantiates a new TikTokPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDraft

`func (o *TikTokPlatformData) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *TikTokPlatformData) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *TikTokPlatformData) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *TikTokPlatformData) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *TikTokPlatformData) GetPrivacyLevel() string`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *TikTokPlatformData) GetPrivacyLevelOk() (*string, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *TikTokPlatformData) SetPrivacyLevel(v string)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *TikTokPlatformData) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### GetAllowComment

`func (o *TikTokPlatformData) GetAllowComment() bool`

GetAllowComment returns the AllowComment field if non-nil, zero value otherwise.

### GetAllowCommentOk

`func (o *TikTokPlatformData) GetAllowCommentOk() (*bool, bool)`

GetAllowCommentOk returns a tuple with the AllowComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowComment

`func (o *TikTokPlatformData) SetAllowComment(v bool)`

SetAllowComment sets AllowComment field to given value.

### HasAllowComment

`func (o *TikTokPlatformData) HasAllowComment() bool`

HasAllowComment returns a boolean if a field has been set.

### GetAllowDuet

`func (o *TikTokPlatformData) GetAllowDuet() bool`

GetAllowDuet returns the AllowDuet field if non-nil, zero value otherwise.

### GetAllowDuetOk

`func (o *TikTokPlatformData) GetAllowDuetOk() (*bool, bool)`

GetAllowDuetOk returns a tuple with the AllowDuet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuet

`func (o *TikTokPlatformData) SetAllowDuet(v bool)`

SetAllowDuet sets AllowDuet field to given value.

### HasAllowDuet

`func (o *TikTokPlatformData) HasAllowDuet() bool`

HasAllowDuet returns a boolean if a field has been set.

### GetAllowStitch

`func (o *TikTokPlatformData) GetAllowStitch() bool`

GetAllowStitch returns the AllowStitch field if non-nil, zero value otherwise.

### GetAllowStitchOk

`func (o *TikTokPlatformData) GetAllowStitchOk() (*bool, bool)`

GetAllowStitchOk returns a tuple with the AllowStitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStitch

`func (o *TikTokPlatformData) SetAllowStitch(v bool)`

SetAllowStitch sets AllowStitch field to given value.

### HasAllowStitch

`func (o *TikTokPlatformData) HasAllowStitch() bool`

HasAllowStitch returns a boolean if a field has been set.

### GetCommercialContentType

`func (o *TikTokPlatformData) GetCommercialContentType() string`

GetCommercialContentType returns the CommercialContentType field if non-nil, zero value otherwise.

### GetCommercialContentTypeOk

`func (o *TikTokPlatformData) GetCommercialContentTypeOk() (*string, bool)`

GetCommercialContentTypeOk returns a tuple with the CommercialContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialContentType

`func (o *TikTokPlatformData) SetCommercialContentType(v string)`

SetCommercialContentType sets CommercialContentType field to given value.

### HasCommercialContentType

`func (o *TikTokPlatformData) HasCommercialContentType() bool`

HasCommercialContentType returns a boolean if a field has been set.

### GetBrandPartnerPromote

`func (o *TikTokPlatformData) GetBrandPartnerPromote() bool`

GetBrandPartnerPromote returns the BrandPartnerPromote field if non-nil, zero value otherwise.

### GetBrandPartnerPromoteOk

`func (o *TikTokPlatformData) GetBrandPartnerPromoteOk() (*bool, bool)`

GetBrandPartnerPromoteOk returns a tuple with the BrandPartnerPromote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandPartnerPromote

`func (o *TikTokPlatformData) SetBrandPartnerPromote(v bool)`

SetBrandPartnerPromote sets BrandPartnerPromote field to given value.

### HasBrandPartnerPromote

`func (o *TikTokPlatformData) HasBrandPartnerPromote() bool`

HasBrandPartnerPromote returns a boolean if a field has been set.

### GetIsBrandOrganicPost

`func (o *TikTokPlatformData) GetIsBrandOrganicPost() bool`

GetIsBrandOrganicPost returns the IsBrandOrganicPost field if non-nil, zero value otherwise.

### GetIsBrandOrganicPostOk

`func (o *TikTokPlatformData) GetIsBrandOrganicPostOk() (*bool, bool)`

GetIsBrandOrganicPostOk returns a tuple with the IsBrandOrganicPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandOrganicPost

`func (o *TikTokPlatformData) SetIsBrandOrganicPost(v bool)`

SetIsBrandOrganicPost sets IsBrandOrganicPost field to given value.

### HasIsBrandOrganicPost

`func (o *TikTokPlatformData) HasIsBrandOrganicPost() bool`

HasIsBrandOrganicPost returns a boolean if a field has been set.

### GetContentPreviewConfirmed

`func (o *TikTokPlatformData) GetContentPreviewConfirmed() bool`

GetContentPreviewConfirmed returns the ContentPreviewConfirmed field if non-nil, zero value otherwise.

### GetContentPreviewConfirmedOk

`func (o *TikTokPlatformData) GetContentPreviewConfirmedOk() (*bool, bool)`

GetContentPreviewConfirmedOk returns a tuple with the ContentPreviewConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPreviewConfirmed

`func (o *TikTokPlatformData) SetContentPreviewConfirmed(v bool)`

SetContentPreviewConfirmed sets ContentPreviewConfirmed field to given value.

### HasContentPreviewConfirmed

`func (o *TikTokPlatformData) HasContentPreviewConfirmed() bool`

HasContentPreviewConfirmed returns a boolean if a field has been set.

### GetExpressConsentGiven

`func (o *TikTokPlatformData) GetExpressConsentGiven() bool`

GetExpressConsentGiven returns the ExpressConsentGiven field if non-nil, zero value otherwise.

### GetExpressConsentGivenOk

`func (o *TikTokPlatformData) GetExpressConsentGivenOk() (*bool, bool)`

GetExpressConsentGivenOk returns a tuple with the ExpressConsentGiven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressConsentGiven

`func (o *TikTokPlatformData) SetExpressConsentGiven(v bool)`

SetExpressConsentGiven sets ExpressConsentGiven field to given value.

### HasExpressConsentGiven

`func (o *TikTokPlatformData) HasExpressConsentGiven() bool`

HasExpressConsentGiven returns a boolean if a field has been set.

### GetMediaType

`func (o *TikTokPlatformData) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *TikTokPlatformData) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *TikTokPlatformData) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *TikTokPlatformData) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetVideoCoverTimestampMs

`func (o *TikTokPlatformData) GetVideoCoverTimestampMs() int32`

GetVideoCoverTimestampMs returns the VideoCoverTimestampMs field if non-nil, zero value otherwise.

### GetVideoCoverTimestampMsOk

`func (o *TikTokPlatformData) GetVideoCoverTimestampMsOk() (*int32, bool)`

GetVideoCoverTimestampMsOk returns a tuple with the VideoCoverTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCoverTimestampMs

`func (o *TikTokPlatformData) SetVideoCoverTimestampMs(v int32)`

SetVideoCoverTimestampMs sets VideoCoverTimestampMs field to given value.

### HasVideoCoverTimestampMs

`func (o *TikTokPlatformData) HasVideoCoverTimestampMs() bool`

HasVideoCoverTimestampMs returns a boolean if a field has been set.

### GetVideoCoverImageUrl

`func (o *TikTokPlatformData) GetVideoCoverImageUrl() string`

GetVideoCoverImageUrl returns the VideoCoverImageUrl field if non-nil, zero value otherwise.

### GetVideoCoverImageUrlOk

`func (o *TikTokPlatformData) GetVideoCoverImageUrlOk() (*string, bool)`

GetVideoCoverImageUrlOk returns a tuple with the VideoCoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCoverImageUrl

`func (o *TikTokPlatformData) SetVideoCoverImageUrl(v string)`

SetVideoCoverImageUrl sets VideoCoverImageUrl field to given value.

### HasVideoCoverImageUrl

`func (o *TikTokPlatformData) HasVideoCoverImageUrl() bool`

HasVideoCoverImageUrl returns a boolean if a field has been set.

### GetPhotoCoverIndex

`func (o *TikTokPlatformData) GetPhotoCoverIndex() int32`

GetPhotoCoverIndex returns the PhotoCoverIndex field if non-nil, zero value otherwise.

### GetPhotoCoverIndexOk

`func (o *TikTokPlatformData) GetPhotoCoverIndexOk() (*int32, bool)`

GetPhotoCoverIndexOk returns a tuple with the PhotoCoverIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCoverIndex

`func (o *TikTokPlatformData) SetPhotoCoverIndex(v int32)`

SetPhotoCoverIndex sets PhotoCoverIndex field to given value.

### HasPhotoCoverIndex

`func (o *TikTokPlatformData) HasPhotoCoverIndex() bool`

HasPhotoCoverIndex returns a boolean if a field has been set.

### GetAutoAddMusic

`func (o *TikTokPlatformData) GetAutoAddMusic() bool`

GetAutoAddMusic returns the AutoAddMusic field if non-nil, zero value otherwise.

### GetAutoAddMusicOk

`func (o *TikTokPlatformData) GetAutoAddMusicOk() (*bool, bool)`

GetAutoAddMusicOk returns a tuple with the AutoAddMusic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddMusic

`func (o *TikTokPlatformData) SetAutoAddMusic(v bool)`

SetAutoAddMusic sets AutoAddMusic field to given value.

### HasAutoAddMusic

`func (o *TikTokPlatformData) HasAutoAddMusic() bool`

HasAutoAddMusic returns a boolean if a field has been set.

### GetVideoMadeWithAi

`func (o *TikTokPlatformData) GetVideoMadeWithAi() bool`

GetVideoMadeWithAi returns the VideoMadeWithAi field if non-nil, zero value otherwise.

### GetVideoMadeWithAiOk

`func (o *TikTokPlatformData) GetVideoMadeWithAiOk() (*bool, bool)`

GetVideoMadeWithAiOk returns a tuple with the VideoMadeWithAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoMadeWithAi

`func (o *TikTokPlatformData) SetVideoMadeWithAi(v bool)`

SetVideoMadeWithAi sets VideoMadeWithAi field to given value.

### HasVideoMadeWithAi

`func (o *TikTokPlatformData) HasVideoMadeWithAi() bool`

HasVideoMadeWithAi returns a boolean if a field has been set.

### GetDescription

`func (o *TikTokPlatformData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TikTokPlatformData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TikTokPlatformData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TikTokPlatformData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


