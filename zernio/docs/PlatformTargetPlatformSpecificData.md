# PlatformTargetPlatformSpecificData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyToTweetId** | Pointer to **string** | ID of an existing tweet to reply to. The published tweet will appear as a reply in that tweet&#39;s thread. For threads, only the first tweet replies to the target; subsequent tweets chain normally. | [optional] 
**QuoteTweetId** | Pointer to **string** | ID (or full status URL) of an existing tweet to quote-repost. The published tweet becomes a quote tweet of the target. Mutually exclusive with media and poll. X only permits quoting your own posts or posts you are mentioned in / part of the conversation thread of; quoting an arbitrary other account&#39;s post is rejected by X. Billed at the standard create rate ($0.015), unlike pasting a tweet URL into the text which is billed at the URL rate ($0.20). For threads, applies to the first tweet only. | [optional] 
**ReplySettings** | Pointer to **string** | Controls who can reply to the tweet. \&quot;following\&quot; allows only people you follow, \&quot;mentionedUsers\&quot; allows only mentioned users, \&quot;subscribers\&quot; allows only subscribers, \&quot;verified\&quot; allows only verified users. Omit for default (everyone can reply). For threads, applies to the first tweet only. Cannot be combined with replyToTweetId. | [optional] 
**ThreadItems** | Pointer to [**[]TwitterPlatformDataThreadItemsInner**](TwitterPlatformDataThreadItemsInner.md) | Complete sequence of posts in a Bluesky thread. The first item becomes the root post, subsequent items are chained as replies. When threadItems is provided, the top-level content field is used only for display and search purposes, it is NOT published. You must include your first post as threadItems[0].  | [optional] 
**Poll** | Pointer to [**DiscordPlatformDataPoll**](DiscordPlatformDataPoll.md) |  | [optional] 
**LongVideo** | Pointer to **bool** | Enable long video uploads (over 140 seconds) using amplify_video media category. Requires the connected X account to have an active X Premium subscription. When true, videos are uploaded with the amplify_video category which supports longer durations (up to 10 minutes via API). When false or omitted, the standard tweet_video category is used (140 second limit). Note that not all Premium accounts have API long-video access, as X may require separate allowlisting. | [optional] [default to false]
**GeoRestriction** | Pointer to [**GeoRestriction**](GeoRestriction.md) |  | [optional] 
**PaidPartnership** | Pointer to **bool** | When true, the post is labeled by X as a paid partnership / paid promotion. For threads, applies to the root tweet only. Field availability may depend on your X API access tier. | [optional] [default to false]
**MadeWithAi** | Pointer to **bool** | When true, the post is labeled by X as containing AI-generated media. Per X, this label is for AI-generated media, not AI-written text. For threads, applies to the root tweet only. | [optional] [default to false]
**SensitiveMedia** | Pointer to [**TwitterPlatformDataSensitiveMedia**](TwitterPlatformDataSensitiveMedia.md) |  | [optional] 
**TopicTag** | Pointer to **string** | Topic tag for post categorization and discoverability on Threads. Must be 1-50 characters, cannot contain periods (.) or ampersands (&amp;). Overrides auto-extraction from content hashtags when provided. | [optional] 
**Draft** | Pointer to **bool** | When true, sends the post to the TikTok Creator Inbox as a draft instead of publishing immediately. The creator receives an inbox notification to complete posting via TikTok&#39;s editing flow. Maps to TikTok API post_mode: \&quot;MEDIA_UPLOAD\&quot; (photos) or the dedicated inbox endpoint (videos). When false or omitted, publishes directly via post_mode: \&quot;DIRECT_POST\&quot;. Note: publish_type is not a supported field. Use this field instead.  | [optional] 
**ContentType** | Pointer to **string** | Content type: story (ephemeral 24h, default), saved_story (permanent on Public Profile), spotlight (video feed) | [optional] [default to "story"]
**Title** | Pointer to **string** | Post title. Defaults to the first line of content, truncated to 300 characters. | [optional] 
**FirstComment** | Pointer to **string** | Optional first comment to post immediately after video upload. Up to 10,000 characters (YouTube&#39;s comment limit). | [optional] 
**PageId** | Pointer to **string** | Target Facebook Page ID for multi-page posting. If omitted, uses the default page. Use GET /v1/accounts/{id}/facebook-page to list pages. | [optional] 
**CarouselCards** | Pointer to [**[]FacebookPlatformDataCarouselCardsInner**](FacebookPlatformDataCarouselCardsInner.md) | Renders the post as a multi-link carousel (organic Page post). When set, mediaItems must be provided with the same length and all items must be images (no videos). Each cards[i] adds the click-through link and headline for the image at mediaItems[i]. Mutually exclusive with contentType&#x3D;story|reel. Facebook display truncates name at ~35 chars and description at ~30 chars; longer strings are accepted but get truncated on render.  | [optional] 
**CarouselLink** | Pointer to **string** | Optional top-level \&quot;See more\&quot; destination shown on the carousel end card. Defaults to the first card&#39;s link when omitted. Only used together with carouselCards.  | [optional] 
**ShareToFeed** | Pointer to **bool** | For Reels only. When true (default), the Reel appears on both the Reels tab and your main profile feed. Set to false to post to the Reels tab only. | [optional] [default to true]
**Collaborators** | Pointer to **[]string** | Up to 3 Instagram usernames to invite as collaborators (feed/Reels only) | [optional] 
**TrialParams** | Pointer to [**InstagramPlatformDataTrialParams**](InstagramPlatformDataTrialParams.md) |  | [optional] 
**UserTags** | Pointer to [**[]InstagramPlatformDataUserTagsInner**](InstagramPlatformDataUserTagsInner.md) | Tag Instagram users in photos by username and position. Not supported for stories or videos. For carousels, use mediaIndex to target specific slides (defaults to 0). Tags on video items are silently skipped. | [optional] 
**AudioName** | Pointer to **string** | Custom name for original audio in Reels. Replaces the default \&quot;Original Audio\&quot; label. Can only be set once. | [optional] 
**ThumbOffset** | Pointer to **int32** | Millisecond offset from video start for the Reel cover frame. Ignored when instagramThumbnail or reelCover is provided. Defaults to 0. | [optional] 
**InstagramThumbnail** | Pointer to **string** | Custom cover image URL for Instagram Reels (JPG or PNG, publicly accessible). Overrides thumbOffset when provided. Also accepted as reelCover (alias). | [optional] 
**ReelCover** | Pointer to **string** | Alias for instagramThumbnail. If both are provided, instagramThumbnail takes priority. | [optional] 
**DocumentTitle** | Pointer to **string** | Title displayed on LinkedIn document (PDF/carousel) posts. Required by LinkedIn for document posts. If omitted, falls back to the media item title, then the filename. | [optional] 
**OrganizationUrn** | Pointer to **string** | Target LinkedIn Organization URN (e.g. \&quot;urn:li:organization:123456789\&quot;). If omitted, uses the default org. Use GET /v1/accounts/{id}/linkedin-organizations to list orgs. | [optional] 
**DisableLinkPreview** | Pointer to **bool** | Set to true to disable automatic link previews for URLs in the post content (default is false) | [optional] 
**ReshareUrl** | Pointer to **string** | LinkedIn post link to repost (use the post&#39;s \&quot;Copy link to post\&quot; action), or a urn:li:share / urn:li:ugcPost / urn:li:groupPost URN. The published post becomes a quote-reshare: your content is shown as the commentary and the original post is embedded underneath (LinkedIn&#39;s \&quot;repost with your thoughts\&quot;). Mutually exclusive with media. Works on personal profiles and organization pages. | [optional] 
**BoardId** | Pointer to **string** | Target Pinterest board ID. If omitted, the first available board is used. | [optional] 
**Link** | Pointer to **string** | Destination link (pin URL) | [optional] 
**CoverImageUrl** | Pointer to **string** | Optional cover image for video pins | [optional] 
**CoverImageKeyFrameTime** | Pointer to **int32** | Optional key frame time in seconds for derived video cover | [optional] 
**Visibility** | Pointer to **string** | Video visibility: public (default, anyone can watch), unlisted (link only), private (invite only) | [optional] [default to "public"]
**MadeForKids** | Pointer to **bool** | COPPA compliance flag. Set true for child-directed content (restricts comments, notifications, ad targeting). Defaults to false. YouTube may block views if not explicitly set. | [optional] [default to false]
**ContainsSyntheticMedia** | Pointer to **bool** | AI-generated content disclosure. Set true if the video contains synthetic content that could be mistaken for real. YouTube may add a label. | [optional] [default to false]
**CategoryId** | Pointer to **string** | YouTube video category ID. Defaults to 22 (People &amp; Blogs). Common: 1 (Film), 2 (Autos), 10 (Music), 15 (Pets), 17 (Sports), 20 (Gaming), 23 (Comedy), 24 (Entertainment), 25 (News), 26 (Howto), 27 (Education), 28 (Science &amp; Tech). | [optional] [default to "22"]
**PlaylistId** | Pointer to **string** | Optional YouTube playlist ID to add the video to after upload (e.g. &#39;PLxxxxxxxxxxxxx&#39;). Use GET /v1/accounts/{id}/youtube-playlists to list available playlists. Works for both immediate and scheduled uploads. Quota cost: 50 YouTube API units per call. | [optional] 
**LocationId** | Pointer to **string** | Target GBP location ID (e.g. \&quot;locations/123456789\&quot;). If omitted, uses the default location. Use GET /v1/accounts/{id}/gmb-locations to list locations. | [optional] 
**LanguageCode** | Pointer to **string** | BCP 47 language code (e.g. \&quot;en\&quot;, \&quot;de\&quot;, \&quot;es\&quot;). Auto-detected if omitted. Set explicitly for short or mixed-language posts. | [optional] 
**TopicType** | Pointer to **string** | Post type. STANDARD is a regular update. EVENT requires the event object. OFFER requires the offer object. Defaults to STANDARD if omitted. | [optional] [default to "STANDARD"]
**CallToAction** | Pointer to [**GoogleBusinessPlatformDataCallToAction**](GoogleBusinessPlatformDataCallToAction.md) |  | [optional] 
**Event** | Pointer to [**GoogleBusinessPlatformDataEvent**](GoogleBusinessPlatformDataEvent.md) |  | [optional] 
**Offer** | Pointer to [**GoogleBusinessPlatformDataOffer**](GoogleBusinessPlatformDataOffer.md) |  | [optional] 
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
**ParseMode** | Pointer to **string** | Text formatting mode for the message (default is HTML) | [optional] 
**DisableWebPagePreview** | Pointer to **bool** | Disable link preview generation for URLs in the message | [optional] 
**DisableNotification** | Pointer to **bool** | Send the message silently (users will receive notification without sound) | [optional] 
**ProtectContent** | Pointer to **bool** | Protect message content from forwarding and saving | [optional] 
**Subreddit** | Pointer to **string** | Target subreddit name (without \&quot;r/\&quot; prefix). Overrides the default. Use GET /v1/accounts/{id}/reddit-subreddits to list options. | [optional] 
**Url** | Pointer to **string** | URL for link posts. If provided (and forceSelf is not true), creates a link post instead of a text post. | [optional] 
**ForceSelf** | Pointer to **bool** | When true, creates a text/self post even when a URL or media is provided. | [optional] 
**FlairId** | Pointer to **string** | Flair ID for the post. Required by some subreddits. Use GET /v1/accounts/{id}/reddit-flairs?subreddit&#x3D;name to list flairs. | [optional] 
**FlairText** | Pointer to **string** | Custom flair text, for subreddits that allow free-text flair. Ignored when flairId is provided (flairId wins). | [optional] 
**Nsfw** | Pointer to **bool** | Mark the post as NSFW (Not Safe For Work / over 18). | [optional] [default to false]
**Spoiler** | Pointer to **bool** | Mark the post as a spoiler. The subreddit must have spoiler tagging enabled for this to take effect. | [optional] [default to false]
**Sendreplies** | Pointer to **bool** | Whether to receive inbox replies for comments on this post. Set to false to opt out. | [optional] [default to true]
**NativeVideo** | Pointer to **bool** | Controls Reddit&#39;s native video upload flow. When true (default for video mediaItems), the video is uploaded to Reddit&#39;s CDN and submitted with kind&#x3D;video so it renders as an embedded Reddit video player. Reddit transcodes server-side (1080p/30fps cap). Set to false to fall back to a legacy link post. If the subreddit blocks video posts, the upload falls back to a link post automatically.  | [optional] [default to true]
**Videogif** | Pointer to **bool** | When true (and nativeVideo is active), submits the video as a silent videogif (kind&#x3D;videogif). Use for short looping clips without audio. | [optional] 
**VideoPosterUrl** | Pointer to **string** | Optional poster/thumbnail image URL for native video posts. If omitted, the first frame of the video is extracted and used automatically. | [optional] 
**ChannelId** | **string** | Target channel snowflake ID. Determines which channel in the connected server receives the message. | 
**Embeds** | Pointer to [**[]DiscordPlatformDataEmbedsInner**](DiscordPlatformDataEmbedsInner.md) | Up to 10 Discord embed objects (combined max 6,000 characters across all embeds). Sent alongside or instead of plain-text content. | [optional] 
**Crosspost** | Pointer to **bool** | Auto-crosspost to every server following this announcement channel (type 5). No-op for regular text channels. | [optional] 
**ForumThreadName** | Pointer to **string** | Thread title for forum channel posts (type 15). Required when posting to a forum channel. | [optional] 
**ForumAppliedTags** | Pointer to **[]string** | Tag snowflake IDs to apply to forum posts. Max 5 tags. | [optional] 
**ThreadFromMessage** | Pointer to [**DiscordPlatformDataThreadFromMessage**](DiscordPlatformDataThreadFromMessage.md) |  | [optional] 
**Tts** | Pointer to **bool** | Send as text-to-speech message. Discord reads the message aloud in the channel. | [optional] 
**WebhookUsername** | Pointer to **string** | Override the webhook display name for this post only (1-80 chars). Falls back to the account-level default set via PATCH /v1/connect/discord. | [optional] 
**WebhookAvatarUrl** | Pointer to **string** | Override the webhook avatar URL for this post only. Falls back to the account-level default. | [optional] 

## Methods

### NewPlatformTargetPlatformSpecificData

`func NewPlatformTargetPlatformSpecificData(channelId string, ) *PlatformTargetPlatformSpecificData`

NewPlatformTargetPlatformSpecificData instantiates a new PlatformTargetPlatformSpecificData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlatformTargetPlatformSpecificDataWithDefaults

`func NewPlatformTargetPlatformSpecificDataWithDefaults() *PlatformTargetPlatformSpecificData`

NewPlatformTargetPlatformSpecificDataWithDefaults instantiates a new PlatformTargetPlatformSpecificData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyToTweetId

`func (o *PlatformTargetPlatformSpecificData) GetReplyToTweetId() string`

GetReplyToTweetId returns the ReplyToTweetId field if non-nil, zero value otherwise.

### GetReplyToTweetIdOk

`func (o *PlatformTargetPlatformSpecificData) GetReplyToTweetIdOk() (*string, bool)`

GetReplyToTweetIdOk returns a tuple with the ReplyToTweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToTweetId

`func (o *PlatformTargetPlatformSpecificData) SetReplyToTweetId(v string)`

SetReplyToTweetId sets ReplyToTweetId field to given value.

### HasReplyToTweetId

`func (o *PlatformTargetPlatformSpecificData) HasReplyToTweetId() bool`

HasReplyToTweetId returns a boolean if a field has been set.

### GetQuoteTweetId

`func (o *PlatformTargetPlatformSpecificData) GetQuoteTweetId() string`

GetQuoteTweetId returns the QuoteTweetId field if non-nil, zero value otherwise.

### GetQuoteTweetIdOk

`func (o *PlatformTargetPlatformSpecificData) GetQuoteTweetIdOk() (*string, bool)`

GetQuoteTweetIdOk returns a tuple with the QuoteTweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTweetId

`func (o *PlatformTargetPlatformSpecificData) SetQuoteTweetId(v string)`

SetQuoteTweetId sets QuoteTweetId field to given value.

### HasQuoteTweetId

`func (o *PlatformTargetPlatformSpecificData) HasQuoteTweetId() bool`

HasQuoteTweetId returns a boolean if a field has been set.

### GetReplySettings

`func (o *PlatformTargetPlatformSpecificData) GetReplySettings() string`

GetReplySettings returns the ReplySettings field if non-nil, zero value otherwise.

### GetReplySettingsOk

`func (o *PlatformTargetPlatformSpecificData) GetReplySettingsOk() (*string, bool)`

GetReplySettingsOk returns a tuple with the ReplySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplySettings

`func (o *PlatformTargetPlatformSpecificData) SetReplySettings(v string)`

SetReplySettings sets ReplySettings field to given value.

### HasReplySettings

`func (o *PlatformTargetPlatformSpecificData) HasReplySettings() bool`

HasReplySettings returns a boolean if a field has been set.

### GetThreadItems

`func (o *PlatformTargetPlatformSpecificData) GetThreadItems() []TwitterPlatformDataThreadItemsInner`

GetThreadItems returns the ThreadItems field if non-nil, zero value otherwise.

### GetThreadItemsOk

`func (o *PlatformTargetPlatformSpecificData) GetThreadItemsOk() (*[]TwitterPlatformDataThreadItemsInner, bool)`

GetThreadItemsOk returns a tuple with the ThreadItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadItems

`func (o *PlatformTargetPlatformSpecificData) SetThreadItems(v []TwitterPlatformDataThreadItemsInner)`

SetThreadItems sets ThreadItems field to given value.

### HasThreadItems

`func (o *PlatformTargetPlatformSpecificData) HasThreadItems() bool`

HasThreadItems returns a boolean if a field has been set.

### GetPoll

`func (o *PlatformTargetPlatformSpecificData) GetPoll() DiscordPlatformDataPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *PlatformTargetPlatformSpecificData) GetPollOk() (*DiscordPlatformDataPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *PlatformTargetPlatformSpecificData) SetPoll(v DiscordPlatformDataPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *PlatformTargetPlatformSpecificData) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetLongVideo

`func (o *PlatformTargetPlatformSpecificData) GetLongVideo() bool`

GetLongVideo returns the LongVideo field if non-nil, zero value otherwise.

### GetLongVideoOk

`func (o *PlatformTargetPlatformSpecificData) GetLongVideoOk() (*bool, bool)`

GetLongVideoOk returns a tuple with the LongVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongVideo

`func (o *PlatformTargetPlatformSpecificData) SetLongVideo(v bool)`

SetLongVideo sets LongVideo field to given value.

### HasLongVideo

`func (o *PlatformTargetPlatformSpecificData) HasLongVideo() bool`

HasLongVideo returns a boolean if a field has been set.

### GetGeoRestriction

`func (o *PlatformTargetPlatformSpecificData) GetGeoRestriction() GeoRestriction`

GetGeoRestriction returns the GeoRestriction field if non-nil, zero value otherwise.

### GetGeoRestrictionOk

`func (o *PlatformTargetPlatformSpecificData) GetGeoRestrictionOk() (*GeoRestriction, bool)`

GetGeoRestrictionOk returns a tuple with the GeoRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestriction

`func (o *PlatformTargetPlatformSpecificData) SetGeoRestriction(v GeoRestriction)`

SetGeoRestriction sets GeoRestriction field to given value.

### HasGeoRestriction

`func (o *PlatformTargetPlatformSpecificData) HasGeoRestriction() bool`

HasGeoRestriction returns a boolean if a field has been set.

### GetPaidPartnership

`func (o *PlatformTargetPlatformSpecificData) GetPaidPartnership() bool`

GetPaidPartnership returns the PaidPartnership field if non-nil, zero value otherwise.

### GetPaidPartnershipOk

`func (o *PlatformTargetPlatformSpecificData) GetPaidPartnershipOk() (*bool, bool)`

GetPaidPartnershipOk returns a tuple with the PaidPartnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidPartnership

`func (o *PlatformTargetPlatformSpecificData) SetPaidPartnership(v bool)`

SetPaidPartnership sets PaidPartnership field to given value.

### HasPaidPartnership

`func (o *PlatformTargetPlatformSpecificData) HasPaidPartnership() bool`

HasPaidPartnership returns a boolean if a field has been set.

### GetMadeWithAi

`func (o *PlatformTargetPlatformSpecificData) GetMadeWithAi() bool`

GetMadeWithAi returns the MadeWithAi field if non-nil, zero value otherwise.

### GetMadeWithAiOk

`func (o *PlatformTargetPlatformSpecificData) GetMadeWithAiOk() (*bool, bool)`

GetMadeWithAiOk returns a tuple with the MadeWithAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadeWithAi

`func (o *PlatformTargetPlatformSpecificData) SetMadeWithAi(v bool)`

SetMadeWithAi sets MadeWithAi field to given value.

### HasMadeWithAi

`func (o *PlatformTargetPlatformSpecificData) HasMadeWithAi() bool`

HasMadeWithAi returns a boolean if a field has been set.

### GetSensitiveMedia

`func (o *PlatformTargetPlatformSpecificData) GetSensitiveMedia() TwitterPlatformDataSensitiveMedia`

GetSensitiveMedia returns the SensitiveMedia field if non-nil, zero value otherwise.

### GetSensitiveMediaOk

`func (o *PlatformTargetPlatformSpecificData) GetSensitiveMediaOk() (*TwitterPlatformDataSensitiveMedia, bool)`

GetSensitiveMediaOk returns a tuple with the SensitiveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveMedia

`func (o *PlatformTargetPlatformSpecificData) SetSensitiveMedia(v TwitterPlatformDataSensitiveMedia)`

SetSensitiveMedia sets SensitiveMedia field to given value.

### HasSensitiveMedia

`func (o *PlatformTargetPlatformSpecificData) HasSensitiveMedia() bool`

HasSensitiveMedia returns a boolean if a field has been set.

### GetTopicTag

`func (o *PlatformTargetPlatformSpecificData) GetTopicTag() string`

GetTopicTag returns the TopicTag field if non-nil, zero value otherwise.

### GetTopicTagOk

`func (o *PlatformTargetPlatformSpecificData) GetTopicTagOk() (*string, bool)`

GetTopicTagOk returns a tuple with the TopicTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicTag

`func (o *PlatformTargetPlatformSpecificData) SetTopicTag(v string)`

SetTopicTag sets TopicTag field to given value.

### HasTopicTag

`func (o *PlatformTargetPlatformSpecificData) HasTopicTag() bool`

HasTopicTag returns a boolean if a field has been set.

### GetDraft

`func (o *PlatformTargetPlatformSpecificData) GetDraft() bool`

GetDraft returns the Draft field if non-nil, zero value otherwise.

### GetDraftOk

`func (o *PlatformTargetPlatformSpecificData) GetDraftOk() (*bool, bool)`

GetDraftOk returns a tuple with the Draft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDraft

`func (o *PlatformTargetPlatformSpecificData) SetDraft(v bool)`

SetDraft sets Draft field to given value.

### HasDraft

`func (o *PlatformTargetPlatformSpecificData) HasDraft() bool`

HasDraft returns a boolean if a field has been set.

### GetContentType

`func (o *PlatformTargetPlatformSpecificData) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PlatformTargetPlatformSpecificData) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PlatformTargetPlatformSpecificData) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PlatformTargetPlatformSpecificData) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetTitle

`func (o *PlatformTargetPlatformSpecificData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlatformTargetPlatformSpecificData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlatformTargetPlatformSpecificData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PlatformTargetPlatformSpecificData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstComment

`func (o *PlatformTargetPlatformSpecificData) GetFirstComment() string`

GetFirstComment returns the FirstComment field if non-nil, zero value otherwise.

### GetFirstCommentOk

`func (o *PlatformTargetPlatformSpecificData) GetFirstCommentOk() (*string, bool)`

GetFirstCommentOk returns a tuple with the FirstComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstComment

`func (o *PlatformTargetPlatformSpecificData) SetFirstComment(v string)`

SetFirstComment sets FirstComment field to given value.

### HasFirstComment

`func (o *PlatformTargetPlatformSpecificData) HasFirstComment() bool`

HasFirstComment returns a boolean if a field has been set.

### GetPageId

`func (o *PlatformTargetPlatformSpecificData) GetPageId() string`

GetPageId returns the PageId field if non-nil, zero value otherwise.

### GetPageIdOk

`func (o *PlatformTargetPlatformSpecificData) GetPageIdOk() (*string, bool)`

GetPageIdOk returns a tuple with the PageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageId

`func (o *PlatformTargetPlatformSpecificData) SetPageId(v string)`

SetPageId sets PageId field to given value.

### HasPageId

`func (o *PlatformTargetPlatformSpecificData) HasPageId() bool`

HasPageId returns a boolean if a field has been set.

### GetCarouselCards

`func (o *PlatformTargetPlatformSpecificData) GetCarouselCards() []FacebookPlatformDataCarouselCardsInner`

GetCarouselCards returns the CarouselCards field if non-nil, zero value otherwise.

### GetCarouselCardsOk

`func (o *PlatformTargetPlatformSpecificData) GetCarouselCardsOk() (*[]FacebookPlatformDataCarouselCardsInner, bool)`

GetCarouselCardsOk returns a tuple with the CarouselCards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarouselCards

`func (o *PlatformTargetPlatformSpecificData) SetCarouselCards(v []FacebookPlatformDataCarouselCardsInner)`

SetCarouselCards sets CarouselCards field to given value.

### HasCarouselCards

`func (o *PlatformTargetPlatformSpecificData) HasCarouselCards() bool`

HasCarouselCards returns a boolean if a field has been set.

### GetCarouselLink

`func (o *PlatformTargetPlatformSpecificData) GetCarouselLink() string`

GetCarouselLink returns the CarouselLink field if non-nil, zero value otherwise.

### GetCarouselLinkOk

`func (o *PlatformTargetPlatformSpecificData) GetCarouselLinkOk() (*string, bool)`

GetCarouselLinkOk returns a tuple with the CarouselLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarouselLink

`func (o *PlatformTargetPlatformSpecificData) SetCarouselLink(v string)`

SetCarouselLink sets CarouselLink field to given value.

### HasCarouselLink

`func (o *PlatformTargetPlatformSpecificData) HasCarouselLink() bool`

HasCarouselLink returns a boolean if a field has been set.

### GetShareToFeed

`func (o *PlatformTargetPlatformSpecificData) GetShareToFeed() bool`

GetShareToFeed returns the ShareToFeed field if non-nil, zero value otherwise.

### GetShareToFeedOk

`func (o *PlatformTargetPlatformSpecificData) GetShareToFeedOk() (*bool, bool)`

GetShareToFeedOk returns a tuple with the ShareToFeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToFeed

`func (o *PlatformTargetPlatformSpecificData) SetShareToFeed(v bool)`

SetShareToFeed sets ShareToFeed field to given value.

### HasShareToFeed

`func (o *PlatformTargetPlatformSpecificData) HasShareToFeed() bool`

HasShareToFeed returns a boolean if a field has been set.

### GetCollaborators

`func (o *PlatformTargetPlatformSpecificData) GetCollaborators() []string`

GetCollaborators returns the Collaborators field if non-nil, zero value otherwise.

### GetCollaboratorsOk

`func (o *PlatformTargetPlatformSpecificData) GetCollaboratorsOk() (*[]string, bool)`

GetCollaboratorsOk returns a tuple with the Collaborators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollaborators

`func (o *PlatformTargetPlatformSpecificData) SetCollaborators(v []string)`

SetCollaborators sets Collaborators field to given value.

### HasCollaborators

`func (o *PlatformTargetPlatformSpecificData) HasCollaborators() bool`

HasCollaborators returns a boolean if a field has been set.

### GetTrialParams

`func (o *PlatformTargetPlatformSpecificData) GetTrialParams() InstagramPlatformDataTrialParams`

GetTrialParams returns the TrialParams field if non-nil, zero value otherwise.

### GetTrialParamsOk

`func (o *PlatformTargetPlatformSpecificData) GetTrialParamsOk() (*InstagramPlatformDataTrialParams, bool)`

GetTrialParamsOk returns a tuple with the TrialParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialParams

`func (o *PlatformTargetPlatformSpecificData) SetTrialParams(v InstagramPlatformDataTrialParams)`

SetTrialParams sets TrialParams field to given value.

### HasTrialParams

`func (o *PlatformTargetPlatformSpecificData) HasTrialParams() bool`

HasTrialParams returns a boolean if a field has been set.

### GetUserTags

`func (o *PlatformTargetPlatformSpecificData) GetUserTags() []InstagramPlatformDataUserTagsInner`

GetUserTags returns the UserTags field if non-nil, zero value otherwise.

### GetUserTagsOk

`func (o *PlatformTargetPlatformSpecificData) GetUserTagsOk() (*[]InstagramPlatformDataUserTagsInner, bool)`

GetUserTagsOk returns a tuple with the UserTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserTags

`func (o *PlatformTargetPlatformSpecificData) SetUserTags(v []InstagramPlatformDataUserTagsInner)`

SetUserTags sets UserTags field to given value.

### HasUserTags

`func (o *PlatformTargetPlatformSpecificData) HasUserTags() bool`

HasUserTags returns a boolean if a field has been set.

### GetAudioName

`func (o *PlatformTargetPlatformSpecificData) GetAudioName() string`

GetAudioName returns the AudioName field if non-nil, zero value otherwise.

### GetAudioNameOk

`func (o *PlatformTargetPlatformSpecificData) GetAudioNameOk() (*string, bool)`

GetAudioNameOk returns a tuple with the AudioName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioName

`func (o *PlatformTargetPlatformSpecificData) SetAudioName(v string)`

SetAudioName sets AudioName field to given value.

### HasAudioName

`func (o *PlatformTargetPlatformSpecificData) HasAudioName() bool`

HasAudioName returns a boolean if a field has been set.

### GetThumbOffset

`func (o *PlatformTargetPlatformSpecificData) GetThumbOffset() int32`

GetThumbOffset returns the ThumbOffset field if non-nil, zero value otherwise.

### GetThumbOffsetOk

`func (o *PlatformTargetPlatformSpecificData) GetThumbOffsetOk() (*int32, bool)`

GetThumbOffsetOk returns a tuple with the ThumbOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbOffset

`func (o *PlatformTargetPlatformSpecificData) SetThumbOffset(v int32)`

SetThumbOffset sets ThumbOffset field to given value.

### HasThumbOffset

`func (o *PlatformTargetPlatformSpecificData) HasThumbOffset() bool`

HasThumbOffset returns a boolean if a field has been set.

### GetInstagramThumbnail

`func (o *PlatformTargetPlatformSpecificData) GetInstagramThumbnail() string`

GetInstagramThumbnail returns the InstagramThumbnail field if non-nil, zero value otherwise.

### GetInstagramThumbnailOk

`func (o *PlatformTargetPlatformSpecificData) GetInstagramThumbnailOk() (*string, bool)`

GetInstagramThumbnailOk returns a tuple with the InstagramThumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramThumbnail

`func (o *PlatformTargetPlatformSpecificData) SetInstagramThumbnail(v string)`

SetInstagramThumbnail sets InstagramThumbnail field to given value.

### HasInstagramThumbnail

`func (o *PlatformTargetPlatformSpecificData) HasInstagramThumbnail() bool`

HasInstagramThumbnail returns a boolean if a field has been set.

### GetReelCover

`func (o *PlatformTargetPlatformSpecificData) GetReelCover() string`

GetReelCover returns the ReelCover field if non-nil, zero value otherwise.

### GetReelCoverOk

`func (o *PlatformTargetPlatformSpecificData) GetReelCoverOk() (*string, bool)`

GetReelCoverOk returns a tuple with the ReelCover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReelCover

`func (o *PlatformTargetPlatformSpecificData) SetReelCover(v string)`

SetReelCover sets ReelCover field to given value.

### HasReelCover

`func (o *PlatformTargetPlatformSpecificData) HasReelCover() bool`

HasReelCover returns a boolean if a field has been set.

### GetDocumentTitle

`func (o *PlatformTargetPlatformSpecificData) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *PlatformTargetPlatformSpecificData) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *PlatformTargetPlatformSpecificData) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *PlatformTargetPlatformSpecificData) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetOrganizationUrn

`func (o *PlatformTargetPlatformSpecificData) GetOrganizationUrn() string`

GetOrganizationUrn returns the OrganizationUrn field if non-nil, zero value otherwise.

### GetOrganizationUrnOk

`func (o *PlatformTargetPlatformSpecificData) GetOrganizationUrnOk() (*string, bool)`

GetOrganizationUrnOk returns a tuple with the OrganizationUrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUrn

`func (o *PlatformTargetPlatformSpecificData) SetOrganizationUrn(v string)`

SetOrganizationUrn sets OrganizationUrn field to given value.

### HasOrganizationUrn

`func (o *PlatformTargetPlatformSpecificData) HasOrganizationUrn() bool`

HasOrganizationUrn returns a boolean if a field has been set.

### GetDisableLinkPreview

`func (o *PlatformTargetPlatformSpecificData) GetDisableLinkPreview() bool`

GetDisableLinkPreview returns the DisableLinkPreview field if non-nil, zero value otherwise.

### GetDisableLinkPreviewOk

`func (o *PlatformTargetPlatformSpecificData) GetDisableLinkPreviewOk() (*bool, bool)`

GetDisableLinkPreviewOk returns a tuple with the DisableLinkPreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableLinkPreview

`func (o *PlatformTargetPlatformSpecificData) SetDisableLinkPreview(v bool)`

SetDisableLinkPreview sets DisableLinkPreview field to given value.

### HasDisableLinkPreview

`func (o *PlatformTargetPlatformSpecificData) HasDisableLinkPreview() bool`

HasDisableLinkPreview returns a boolean if a field has been set.

### GetReshareUrl

`func (o *PlatformTargetPlatformSpecificData) GetReshareUrl() string`

GetReshareUrl returns the ReshareUrl field if non-nil, zero value otherwise.

### GetReshareUrlOk

`func (o *PlatformTargetPlatformSpecificData) GetReshareUrlOk() (*string, bool)`

GetReshareUrlOk returns a tuple with the ReshareUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReshareUrl

`func (o *PlatformTargetPlatformSpecificData) SetReshareUrl(v string)`

SetReshareUrl sets ReshareUrl field to given value.

### HasReshareUrl

`func (o *PlatformTargetPlatformSpecificData) HasReshareUrl() bool`

HasReshareUrl returns a boolean if a field has been set.

### GetBoardId

`func (o *PlatformTargetPlatformSpecificData) GetBoardId() string`

GetBoardId returns the BoardId field if non-nil, zero value otherwise.

### GetBoardIdOk

`func (o *PlatformTargetPlatformSpecificData) GetBoardIdOk() (*string, bool)`

GetBoardIdOk returns a tuple with the BoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoardId

`func (o *PlatformTargetPlatformSpecificData) SetBoardId(v string)`

SetBoardId sets BoardId field to given value.

### HasBoardId

`func (o *PlatformTargetPlatformSpecificData) HasBoardId() bool`

HasBoardId returns a boolean if a field has been set.

### GetLink

`func (o *PlatformTargetPlatformSpecificData) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PlatformTargetPlatformSpecificData) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PlatformTargetPlatformSpecificData) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PlatformTargetPlatformSpecificData) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetCoverImageUrl

`func (o *PlatformTargetPlatformSpecificData) GetCoverImageUrl() string`

GetCoverImageUrl returns the CoverImageUrl field if non-nil, zero value otherwise.

### GetCoverImageUrlOk

`func (o *PlatformTargetPlatformSpecificData) GetCoverImageUrlOk() (*string, bool)`

GetCoverImageUrlOk returns a tuple with the CoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageUrl

`func (o *PlatformTargetPlatformSpecificData) SetCoverImageUrl(v string)`

SetCoverImageUrl sets CoverImageUrl field to given value.

### HasCoverImageUrl

`func (o *PlatformTargetPlatformSpecificData) HasCoverImageUrl() bool`

HasCoverImageUrl returns a boolean if a field has been set.

### GetCoverImageKeyFrameTime

`func (o *PlatformTargetPlatformSpecificData) GetCoverImageKeyFrameTime() int32`

GetCoverImageKeyFrameTime returns the CoverImageKeyFrameTime field if non-nil, zero value otherwise.

### GetCoverImageKeyFrameTimeOk

`func (o *PlatformTargetPlatformSpecificData) GetCoverImageKeyFrameTimeOk() (*int32, bool)`

GetCoverImageKeyFrameTimeOk returns a tuple with the CoverImageKeyFrameTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverImageKeyFrameTime

`func (o *PlatformTargetPlatformSpecificData) SetCoverImageKeyFrameTime(v int32)`

SetCoverImageKeyFrameTime sets CoverImageKeyFrameTime field to given value.

### HasCoverImageKeyFrameTime

`func (o *PlatformTargetPlatformSpecificData) HasCoverImageKeyFrameTime() bool`

HasCoverImageKeyFrameTime returns a boolean if a field has been set.

### GetVisibility

`func (o *PlatformTargetPlatformSpecificData) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *PlatformTargetPlatformSpecificData) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *PlatformTargetPlatformSpecificData) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *PlatformTargetPlatformSpecificData) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetMadeForKids

`func (o *PlatformTargetPlatformSpecificData) GetMadeForKids() bool`

GetMadeForKids returns the MadeForKids field if non-nil, zero value otherwise.

### GetMadeForKidsOk

`func (o *PlatformTargetPlatformSpecificData) GetMadeForKidsOk() (*bool, bool)`

GetMadeForKidsOk returns a tuple with the MadeForKids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadeForKids

`func (o *PlatformTargetPlatformSpecificData) SetMadeForKids(v bool)`

SetMadeForKids sets MadeForKids field to given value.

### HasMadeForKids

`func (o *PlatformTargetPlatformSpecificData) HasMadeForKids() bool`

HasMadeForKids returns a boolean if a field has been set.

### GetContainsSyntheticMedia

`func (o *PlatformTargetPlatformSpecificData) GetContainsSyntheticMedia() bool`

GetContainsSyntheticMedia returns the ContainsSyntheticMedia field if non-nil, zero value otherwise.

### GetContainsSyntheticMediaOk

`func (o *PlatformTargetPlatformSpecificData) GetContainsSyntheticMediaOk() (*bool, bool)`

GetContainsSyntheticMediaOk returns a tuple with the ContainsSyntheticMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainsSyntheticMedia

`func (o *PlatformTargetPlatformSpecificData) SetContainsSyntheticMedia(v bool)`

SetContainsSyntheticMedia sets ContainsSyntheticMedia field to given value.

### HasContainsSyntheticMedia

`func (o *PlatformTargetPlatformSpecificData) HasContainsSyntheticMedia() bool`

HasContainsSyntheticMedia returns a boolean if a field has been set.

### GetCategoryId

`func (o *PlatformTargetPlatformSpecificData) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *PlatformTargetPlatformSpecificData) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *PlatformTargetPlatformSpecificData) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *PlatformTargetPlatformSpecificData) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### GetPlaylistId

`func (o *PlatformTargetPlatformSpecificData) GetPlaylistId() string`

GetPlaylistId returns the PlaylistId field if non-nil, zero value otherwise.

### GetPlaylistIdOk

`func (o *PlatformTargetPlatformSpecificData) GetPlaylistIdOk() (*string, bool)`

GetPlaylistIdOk returns a tuple with the PlaylistId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlaylistId

`func (o *PlatformTargetPlatformSpecificData) SetPlaylistId(v string)`

SetPlaylistId sets PlaylistId field to given value.

### HasPlaylistId

`func (o *PlatformTargetPlatformSpecificData) HasPlaylistId() bool`

HasPlaylistId returns a boolean if a field has been set.

### GetLocationId

`func (o *PlatformTargetPlatformSpecificData) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *PlatformTargetPlatformSpecificData) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *PlatformTargetPlatformSpecificData) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *PlatformTargetPlatformSpecificData) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLanguageCode

`func (o *PlatformTargetPlatformSpecificData) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *PlatformTargetPlatformSpecificData) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *PlatformTargetPlatformSpecificData) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *PlatformTargetPlatformSpecificData) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetTopicType

`func (o *PlatformTargetPlatformSpecificData) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *PlatformTargetPlatformSpecificData) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *PlatformTargetPlatformSpecificData) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *PlatformTargetPlatformSpecificData) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetCallToAction

`func (o *PlatformTargetPlatformSpecificData) GetCallToAction() GoogleBusinessPlatformDataCallToAction`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *PlatformTargetPlatformSpecificData) GetCallToActionOk() (*GoogleBusinessPlatformDataCallToAction, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *PlatformTargetPlatformSpecificData) SetCallToAction(v GoogleBusinessPlatformDataCallToAction)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *PlatformTargetPlatformSpecificData) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetEvent

`func (o *PlatformTargetPlatformSpecificData) GetEvent() GoogleBusinessPlatformDataEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *PlatformTargetPlatformSpecificData) GetEventOk() (*GoogleBusinessPlatformDataEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *PlatformTargetPlatformSpecificData) SetEvent(v GoogleBusinessPlatformDataEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *PlatformTargetPlatformSpecificData) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOffer

`func (o *PlatformTargetPlatformSpecificData) GetOffer() GoogleBusinessPlatformDataOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *PlatformTargetPlatformSpecificData) GetOfferOk() (*GoogleBusinessPlatformDataOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *PlatformTargetPlatformSpecificData) SetOffer(v GoogleBusinessPlatformDataOffer)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *PlatformTargetPlatformSpecificData) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### GetPrivacyLevel

`func (o *PlatformTargetPlatformSpecificData) GetPrivacyLevel() string`

GetPrivacyLevel returns the PrivacyLevel field if non-nil, zero value otherwise.

### GetPrivacyLevelOk

`func (o *PlatformTargetPlatformSpecificData) GetPrivacyLevelOk() (*string, bool)`

GetPrivacyLevelOk returns a tuple with the PrivacyLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyLevel

`func (o *PlatformTargetPlatformSpecificData) SetPrivacyLevel(v string)`

SetPrivacyLevel sets PrivacyLevel field to given value.

### HasPrivacyLevel

`func (o *PlatformTargetPlatformSpecificData) HasPrivacyLevel() bool`

HasPrivacyLevel returns a boolean if a field has been set.

### GetAllowComment

`func (o *PlatformTargetPlatformSpecificData) GetAllowComment() bool`

GetAllowComment returns the AllowComment field if non-nil, zero value otherwise.

### GetAllowCommentOk

`func (o *PlatformTargetPlatformSpecificData) GetAllowCommentOk() (*bool, bool)`

GetAllowCommentOk returns a tuple with the AllowComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowComment

`func (o *PlatformTargetPlatformSpecificData) SetAllowComment(v bool)`

SetAllowComment sets AllowComment field to given value.

### HasAllowComment

`func (o *PlatformTargetPlatformSpecificData) HasAllowComment() bool`

HasAllowComment returns a boolean if a field has been set.

### GetAllowDuet

`func (o *PlatformTargetPlatformSpecificData) GetAllowDuet() bool`

GetAllowDuet returns the AllowDuet field if non-nil, zero value otherwise.

### GetAllowDuetOk

`func (o *PlatformTargetPlatformSpecificData) GetAllowDuetOk() (*bool, bool)`

GetAllowDuetOk returns a tuple with the AllowDuet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowDuet

`func (o *PlatformTargetPlatformSpecificData) SetAllowDuet(v bool)`

SetAllowDuet sets AllowDuet field to given value.

### HasAllowDuet

`func (o *PlatformTargetPlatformSpecificData) HasAllowDuet() bool`

HasAllowDuet returns a boolean if a field has been set.

### GetAllowStitch

`func (o *PlatformTargetPlatformSpecificData) GetAllowStitch() bool`

GetAllowStitch returns the AllowStitch field if non-nil, zero value otherwise.

### GetAllowStitchOk

`func (o *PlatformTargetPlatformSpecificData) GetAllowStitchOk() (*bool, bool)`

GetAllowStitchOk returns a tuple with the AllowStitch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowStitch

`func (o *PlatformTargetPlatformSpecificData) SetAllowStitch(v bool)`

SetAllowStitch sets AllowStitch field to given value.

### HasAllowStitch

`func (o *PlatformTargetPlatformSpecificData) HasAllowStitch() bool`

HasAllowStitch returns a boolean if a field has been set.

### GetCommercialContentType

`func (o *PlatformTargetPlatformSpecificData) GetCommercialContentType() string`

GetCommercialContentType returns the CommercialContentType field if non-nil, zero value otherwise.

### GetCommercialContentTypeOk

`func (o *PlatformTargetPlatformSpecificData) GetCommercialContentTypeOk() (*string, bool)`

GetCommercialContentTypeOk returns a tuple with the CommercialContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommercialContentType

`func (o *PlatformTargetPlatformSpecificData) SetCommercialContentType(v string)`

SetCommercialContentType sets CommercialContentType field to given value.

### HasCommercialContentType

`func (o *PlatformTargetPlatformSpecificData) HasCommercialContentType() bool`

HasCommercialContentType returns a boolean if a field has been set.

### GetBrandPartnerPromote

`func (o *PlatformTargetPlatformSpecificData) GetBrandPartnerPromote() bool`

GetBrandPartnerPromote returns the BrandPartnerPromote field if non-nil, zero value otherwise.

### GetBrandPartnerPromoteOk

`func (o *PlatformTargetPlatformSpecificData) GetBrandPartnerPromoteOk() (*bool, bool)`

GetBrandPartnerPromoteOk returns a tuple with the BrandPartnerPromote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandPartnerPromote

`func (o *PlatformTargetPlatformSpecificData) SetBrandPartnerPromote(v bool)`

SetBrandPartnerPromote sets BrandPartnerPromote field to given value.

### HasBrandPartnerPromote

`func (o *PlatformTargetPlatformSpecificData) HasBrandPartnerPromote() bool`

HasBrandPartnerPromote returns a boolean if a field has been set.

### GetIsBrandOrganicPost

`func (o *PlatformTargetPlatformSpecificData) GetIsBrandOrganicPost() bool`

GetIsBrandOrganicPost returns the IsBrandOrganicPost field if non-nil, zero value otherwise.

### GetIsBrandOrganicPostOk

`func (o *PlatformTargetPlatformSpecificData) GetIsBrandOrganicPostOk() (*bool, bool)`

GetIsBrandOrganicPostOk returns a tuple with the IsBrandOrganicPost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBrandOrganicPost

`func (o *PlatformTargetPlatformSpecificData) SetIsBrandOrganicPost(v bool)`

SetIsBrandOrganicPost sets IsBrandOrganicPost field to given value.

### HasIsBrandOrganicPost

`func (o *PlatformTargetPlatformSpecificData) HasIsBrandOrganicPost() bool`

HasIsBrandOrganicPost returns a boolean if a field has been set.

### GetContentPreviewConfirmed

`func (o *PlatformTargetPlatformSpecificData) GetContentPreviewConfirmed() bool`

GetContentPreviewConfirmed returns the ContentPreviewConfirmed field if non-nil, zero value otherwise.

### GetContentPreviewConfirmedOk

`func (o *PlatformTargetPlatformSpecificData) GetContentPreviewConfirmedOk() (*bool, bool)`

GetContentPreviewConfirmedOk returns a tuple with the ContentPreviewConfirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentPreviewConfirmed

`func (o *PlatformTargetPlatformSpecificData) SetContentPreviewConfirmed(v bool)`

SetContentPreviewConfirmed sets ContentPreviewConfirmed field to given value.

### HasContentPreviewConfirmed

`func (o *PlatformTargetPlatformSpecificData) HasContentPreviewConfirmed() bool`

HasContentPreviewConfirmed returns a boolean if a field has been set.

### GetExpressConsentGiven

`func (o *PlatformTargetPlatformSpecificData) GetExpressConsentGiven() bool`

GetExpressConsentGiven returns the ExpressConsentGiven field if non-nil, zero value otherwise.

### GetExpressConsentGivenOk

`func (o *PlatformTargetPlatformSpecificData) GetExpressConsentGivenOk() (*bool, bool)`

GetExpressConsentGivenOk returns a tuple with the ExpressConsentGiven field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressConsentGiven

`func (o *PlatformTargetPlatformSpecificData) SetExpressConsentGiven(v bool)`

SetExpressConsentGiven sets ExpressConsentGiven field to given value.

### HasExpressConsentGiven

`func (o *PlatformTargetPlatformSpecificData) HasExpressConsentGiven() bool`

HasExpressConsentGiven returns a boolean if a field has been set.

### GetMediaType

`func (o *PlatformTargetPlatformSpecificData) GetMediaType() string`

GetMediaType returns the MediaType field if non-nil, zero value otherwise.

### GetMediaTypeOk

`func (o *PlatformTargetPlatformSpecificData) GetMediaTypeOk() (*string, bool)`

GetMediaTypeOk returns a tuple with the MediaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaType

`func (o *PlatformTargetPlatformSpecificData) SetMediaType(v string)`

SetMediaType sets MediaType field to given value.

### HasMediaType

`func (o *PlatformTargetPlatformSpecificData) HasMediaType() bool`

HasMediaType returns a boolean if a field has been set.

### GetVideoCoverTimestampMs

`func (o *PlatformTargetPlatformSpecificData) GetVideoCoverTimestampMs() int32`

GetVideoCoverTimestampMs returns the VideoCoverTimestampMs field if non-nil, zero value otherwise.

### GetVideoCoverTimestampMsOk

`func (o *PlatformTargetPlatformSpecificData) GetVideoCoverTimestampMsOk() (*int32, bool)`

GetVideoCoverTimestampMsOk returns a tuple with the VideoCoverTimestampMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCoverTimestampMs

`func (o *PlatformTargetPlatformSpecificData) SetVideoCoverTimestampMs(v int32)`

SetVideoCoverTimestampMs sets VideoCoverTimestampMs field to given value.

### HasVideoCoverTimestampMs

`func (o *PlatformTargetPlatformSpecificData) HasVideoCoverTimestampMs() bool`

HasVideoCoverTimestampMs returns a boolean if a field has been set.

### GetVideoCoverImageUrl

`func (o *PlatformTargetPlatformSpecificData) GetVideoCoverImageUrl() string`

GetVideoCoverImageUrl returns the VideoCoverImageUrl field if non-nil, zero value otherwise.

### GetVideoCoverImageUrlOk

`func (o *PlatformTargetPlatformSpecificData) GetVideoCoverImageUrlOk() (*string, bool)`

GetVideoCoverImageUrlOk returns a tuple with the VideoCoverImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoCoverImageUrl

`func (o *PlatformTargetPlatformSpecificData) SetVideoCoverImageUrl(v string)`

SetVideoCoverImageUrl sets VideoCoverImageUrl field to given value.

### HasVideoCoverImageUrl

`func (o *PlatformTargetPlatformSpecificData) HasVideoCoverImageUrl() bool`

HasVideoCoverImageUrl returns a boolean if a field has been set.

### GetPhotoCoverIndex

`func (o *PlatformTargetPlatformSpecificData) GetPhotoCoverIndex() int32`

GetPhotoCoverIndex returns the PhotoCoverIndex field if non-nil, zero value otherwise.

### GetPhotoCoverIndexOk

`func (o *PlatformTargetPlatformSpecificData) GetPhotoCoverIndexOk() (*int32, bool)`

GetPhotoCoverIndexOk returns a tuple with the PhotoCoverIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotoCoverIndex

`func (o *PlatformTargetPlatformSpecificData) SetPhotoCoverIndex(v int32)`

SetPhotoCoverIndex sets PhotoCoverIndex field to given value.

### HasPhotoCoverIndex

`func (o *PlatformTargetPlatformSpecificData) HasPhotoCoverIndex() bool`

HasPhotoCoverIndex returns a boolean if a field has been set.

### GetAutoAddMusic

`func (o *PlatformTargetPlatformSpecificData) GetAutoAddMusic() bool`

GetAutoAddMusic returns the AutoAddMusic field if non-nil, zero value otherwise.

### GetAutoAddMusicOk

`func (o *PlatformTargetPlatformSpecificData) GetAutoAddMusicOk() (*bool, bool)`

GetAutoAddMusicOk returns a tuple with the AutoAddMusic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAddMusic

`func (o *PlatformTargetPlatformSpecificData) SetAutoAddMusic(v bool)`

SetAutoAddMusic sets AutoAddMusic field to given value.

### HasAutoAddMusic

`func (o *PlatformTargetPlatformSpecificData) HasAutoAddMusic() bool`

HasAutoAddMusic returns a boolean if a field has been set.

### GetVideoMadeWithAi

`func (o *PlatformTargetPlatformSpecificData) GetVideoMadeWithAi() bool`

GetVideoMadeWithAi returns the VideoMadeWithAi field if non-nil, zero value otherwise.

### GetVideoMadeWithAiOk

`func (o *PlatformTargetPlatformSpecificData) GetVideoMadeWithAiOk() (*bool, bool)`

GetVideoMadeWithAiOk returns a tuple with the VideoMadeWithAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoMadeWithAi

`func (o *PlatformTargetPlatformSpecificData) SetVideoMadeWithAi(v bool)`

SetVideoMadeWithAi sets VideoMadeWithAi field to given value.

### HasVideoMadeWithAi

`func (o *PlatformTargetPlatformSpecificData) HasVideoMadeWithAi() bool`

HasVideoMadeWithAi returns a boolean if a field has been set.

### GetDescription

`func (o *PlatformTargetPlatformSpecificData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PlatformTargetPlatformSpecificData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PlatformTargetPlatformSpecificData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PlatformTargetPlatformSpecificData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParseMode

`func (o *PlatformTargetPlatformSpecificData) GetParseMode() string`

GetParseMode returns the ParseMode field if non-nil, zero value otherwise.

### GetParseModeOk

`func (o *PlatformTargetPlatformSpecificData) GetParseModeOk() (*string, bool)`

GetParseModeOk returns a tuple with the ParseMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParseMode

`func (o *PlatformTargetPlatformSpecificData) SetParseMode(v string)`

SetParseMode sets ParseMode field to given value.

### HasParseMode

`func (o *PlatformTargetPlatformSpecificData) HasParseMode() bool`

HasParseMode returns a boolean if a field has been set.

### GetDisableWebPagePreview

`func (o *PlatformTargetPlatformSpecificData) GetDisableWebPagePreview() bool`

GetDisableWebPagePreview returns the DisableWebPagePreview field if non-nil, zero value otherwise.

### GetDisableWebPagePreviewOk

`func (o *PlatformTargetPlatformSpecificData) GetDisableWebPagePreviewOk() (*bool, bool)`

GetDisableWebPagePreviewOk returns a tuple with the DisableWebPagePreview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableWebPagePreview

`func (o *PlatformTargetPlatformSpecificData) SetDisableWebPagePreview(v bool)`

SetDisableWebPagePreview sets DisableWebPagePreview field to given value.

### HasDisableWebPagePreview

`func (o *PlatformTargetPlatformSpecificData) HasDisableWebPagePreview() bool`

HasDisableWebPagePreview returns a boolean if a field has been set.

### GetDisableNotification

`func (o *PlatformTargetPlatformSpecificData) GetDisableNotification() bool`

GetDisableNotification returns the DisableNotification field if non-nil, zero value otherwise.

### GetDisableNotificationOk

`func (o *PlatformTargetPlatformSpecificData) GetDisableNotificationOk() (*bool, bool)`

GetDisableNotificationOk returns a tuple with the DisableNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableNotification

`func (o *PlatformTargetPlatformSpecificData) SetDisableNotification(v bool)`

SetDisableNotification sets DisableNotification field to given value.

### HasDisableNotification

`func (o *PlatformTargetPlatformSpecificData) HasDisableNotification() bool`

HasDisableNotification returns a boolean if a field has been set.

### GetProtectContent

`func (o *PlatformTargetPlatformSpecificData) GetProtectContent() bool`

GetProtectContent returns the ProtectContent field if non-nil, zero value otherwise.

### GetProtectContentOk

`func (o *PlatformTargetPlatformSpecificData) GetProtectContentOk() (*bool, bool)`

GetProtectContentOk returns a tuple with the ProtectContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectContent

`func (o *PlatformTargetPlatformSpecificData) SetProtectContent(v bool)`

SetProtectContent sets ProtectContent field to given value.

### HasProtectContent

`func (o *PlatformTargetPlatformSpecificData) HasProtectContent() bool`

HasProtectContent returns a boolean if a field has been set.

### GetSubreddit

`func (o *PlatformTargetPlatformSpecificData) GetSubreddit() string`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *PlatformTargetPlatformSpecificData) GetSubredditOk() (*string, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *PlatformTargetPlatformSpecificData) SetSubreddit(v string)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *PlatformTargetPlatformSpecificData) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### GetUrl

`func (o *PlatformTargetPlatformSpecificData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *PlatformTargetPlatformSpecificData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *PlatformTargetPlatformSpecificData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *PlatformTargetPlatformSpecificData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetForceSelf

`func (o *PlatformTargetPlatformSpecificData) GetForceSelf() bool`

GetForceSelf returns the ForceSelf field if non-nil, zero value otherwise.

### GetForceSelfOk

`func (o *PlatformTargetPlatformSpecificData) GetForceSelfOk() (*bool, bool)`

GetForceSelfOk returns a tuple with the ForceSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSelf

`func (o *PlatformTargetPlatformSpecificData) SetForceSelf(v bool)`

SetForceSelf sets ForceSelf field to given value.

### HasForceSelf

`func (o *PlatformTargetPlatformSpecificData) HasForceSelf() bool`

HasForceSelf returns a boolean if a field has been set.

### GetFlairId

`func (o *PlatformTargetPlatformSpecificData) GetFlairId() string`

GetFlairId returns the FlairId field if non-nil, zero value otherwise.

### GetFlairIdOk

`func (o *PlatformTargetPlatformSpecificData) GetFlairIdOk() (*string, bool)`

GetFlairIdOk returns a tuple with the FlairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairId

`func (o *PlatformTargetPlatformSpecificData) SetFlairId(v string)`

SetFlairId sets FlairId field to given value.

### HasFlairId

`func (o *PlatformTargetPlatformSpecificData) HasFlairId() bool`

HasFlairId returns a boolean if a field has been set.

### GetFlairText

`func (o *PlatformTargetPlatformSpecificData) GetFlairText() string`

GetFlairText returns the FlairText field if non-nil, zero value otherwise.

### GetFlairTextOk

`func (o *PlatformTargetPlatformSpecificData) GetFlairTextOk() (*string, bool)`

GetFlairTextOk returns a tuple with the FlairText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairText

`func (o *PlatformTargetPlatformSpecificData) SetFlairText(v string)`

SetFlairText sets FlairText field to given value.

### HasFlairText

`func (o *PlatformTargetPlatformSpecificData) HasFlairText() bool`

HasFlairText returns a boolean if a field has been set.

### GetNsfw

`func (o *PlatformTargetPlatformSpecificData) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *PlatformTargetPlatformSpecificData) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *PlatformTargetPlatformSpecificData) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *PlatformTargetPlatformSpecificData) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetSpoiler

`func (o *PlatformTargetPlatformSpecificData) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *PlatformTargetPlatformSpecificData) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *PlatformTargetPlatformSpecificData) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *PlatformTargetPlatformSpecificData) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### GetSendreplies

`func (o *PlatformTargetPlatformSpecificData) GetSendreplies() bool`

GetSendreplies returns the Sendreplies field if non-nil, zero value otherwise.

### GetSendrepliesOk

`func (o *PlatformTargetPlatformSpecificData) GetSendrepliesOk() (*bool, bool)`

GetSendrepliesOk returns a tuple with the Sendreplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendreplies

`func (o *PlatformTargetPlatformSpecificData) SetSendreplies(v bool)`

SetSendreplies sets Sendreplies field to given value.

### HasSendreplies

`func (o *PlatformTargetPlatformSpecificData) HasSendreplies() bool`

HasSendreplies returns a boolean if a field has been set.

### GetNativeVideo

`func (o *PlatformTargetPlatformSpecificData) GetNativeVideo() bool`

GetNativeVideo returns the NativeVideo field if non-nil, zero value otherwise.

### GetNativeVideoOk

`func (o *PlatformTargetPlatformSpecificData) GetNativeVideoOk() (*bool, bool)`

GetNativeVideoOk returns a tuple with the NativeVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVideo

`func (o *PlatformTargetPlatformSpecificData) SetNativeVideo(v bool)`

SetNativeVideo sets NativeVideo field to given value.

### HasNativeVideo

`func (o *PlatformTargetPlatformSpecificData) HasNativeVideo() bool`

HasNativeVideo returns a boolean if a field has been set.

### GetVideogif

`func (o *PlatformTargetPlatformSpecificData) GetVideogif() bool`

GetVideogif returns the Videogif field if non-nil, zero value otherwise.

### GetVideogifOk

`func (o *PlatformTargetPlatformSpecificData) GetVideogifOk() (*bool, bool)`

GetVideogifOk returns a tuple with the Videogif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideogif

`func (o *PlatformTargetPlatformSpecificData) SetVideogif(v bool)`

SetVideogif sets Videogif field to given value.

### HasVideogif

`func (o *PlatformTargetPlatformSpecificData) HasVideogif() bool`

HasVideogif returns a boolean if a field has been set.

### GetVideoPosterUrl

`func (o *PlatformTargetPlatformSpecificData) GetVideoPosterUrl() string`

GetVideoPosterUrl returns the VideoPosterUrl field if non-nil, zero value otherwise.

### GetVideoPosterUrlOk

`func (o *PlatformTargetPlatformSpecificData) GetVideoPosterUrlOk() (*string, bool)`

GetVideoPosterUrlOk returns a tuple with the VideoPosterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoPosterUrl

`func (o *PlatformTargetPlatformSpecificData) SetVideoPosterUrl(v string)`

SetVideoPosterUrl sets VideoPosterUrl field to given value.

### HasVideoPosterUrl

`func (o *PlatformTargetPlatformSpecificData) HasVideoPosterUrl() bool`

HasVideoPosterUrl returns a boolean if a field has been set.

### GetChannelId

`func (o *PlatformTargetPlatformSpecificData) GetChannelId() string`

GetChannelId returns the ChannelId field if non-nil, zero value otherwise.

### GetChannelIdOk

`func (o *PlatformTargetPlatformSpecificData) GetChannelIdOk() (*string, bool)`

GetChannelIdOk returns a tuple with the ChannelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelId

`func (o *PlatformTargetPlatformSpecificData) SetChannelId(v string)`

SetChannelId sets ChannelId field to given value.


### GetEmbeds

`func (o *PlatformTargetPlatformSpecificData) GetEmbeds() []DiscordPlatformDataEmbedsInner`

GetEmbeds returns the Embeds field if non-nil, zero value otherwise.

### GetEmbedsOk

`func (o *PlatformTargetPlatformSpecificData) GetEmbedsOk() (*[]DiscordPlatformDataEmbedsInner, bool)`

GetEmbedsOk returns a tuple with the Embeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeds

`func (o *PlatformTargetPlatformSpecificData) SetEmbeds(v []DiscordPlatformDataEmbedsInner)`

SetEmbeds sets Embeds field to given value.

### HasEmbeds

`func (o *PlatformTargetPlatformSpecificData) HasEmbeds() bool`

HasEmbeds returns a boolean if a field has been set.

### GetCrosspost

`func (o *PlatformTargetPlatformSpecificData) GetCrosspost() bool`

GetCrosspost returns the Crosspost field if non-nil, zero value otherwise.

### GetCrosspostOk

`func (o *PlatformTargetPlatformSpecificData) GetCrosspostOk() (*bool, bool)`

GetCrosspostOk returns a tuple with the Crosspost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosspost

`func (o *PlatformTargetPlatformSpecificData) SetCrosspost(v bool)`

SetCrosspost sets Crosspost field to given value.

### HasCrosspost

`func (o *PlatformTargetPlatformSpecificData) HasCrosspost() bool`

HasCrosspost returns a boolean if a field has been set.

### GetForumThreadName

`func (o *PlatformTargetPlatformSpecificData) GetForumThreadName() string`

GetForumThreadName returns the ForumThreadName field if non-nil, zero value otherwise.

### GetForumThreadNameOk

`func (o *PlatformTargetPlatformSpecificData) GetForumThreadNameOk() (*string, bool)`

GetForumThreadNameOk returns a tuple with the ForumThreadName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumThreadName

`func (o *PlatformTargetPlatformSpecificData) SetForumThreadName(v string)`

SetForumThreadName sets ForumThreadName field to given value.

### HasForumThreadName

`func (o *PlatformTargetPlatformSpecificData) HasForumThreadName() bool`

HasForumThreadName returns a boolean if a field has been set.

### GetForumAppliedTags

`func (o *PlatformTargetPlatformSpecificData) GetForumAppliedTags() []string`

GetForumAppliedTags returns the ForumAppliedTags field if non-nil, zero value otherwise.

### GetForumAppliedTagsOk

`func (o *PlatformTargetPlatformSpecificData) GetForumAppliedTagsOk() (*[]string, bool)`

GetForumAppliedTagsOk returns a tuple with the ForumAppliedTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForumAppliedTags

`func (o *PlatformTargetPlatformSpecificData) SetForumAppliedTags(v []string)`

SetForumAppliedTags sets ForumAppliedTags field to given value.

### HasForumAppliedTags

`func (o *PlatformTargetPlatformSpecificData) HasForumAppliedTags() bool`

HasForumAppliedTags returns a boolean if a field has been set.

### GetThreadFromMessage

`func (o *PlatformTargetPlatformSpecificData) GetThreadFromMessage() DiscordPlatformDataThreadFromMessage`

GetThreadFromMessage returns the ThreadFromMessage field if non-nil, zero value otherwise.

### GetThreadFromMessageOk

`func (o *PlatformTargetPlatformSpecificData) GetThreadFromMessageOk() (*DiscordPlatformDataThreadFromMessage, bool)`

GetThreadFromMessageOk returns a tuple with the ThreadFromMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadFromMessage

`func (o *PlatformTargetPlatformSpecificData) SetThreadFromMessage(v DiscordPlatformDataThreadFromMessage)`

SetThreadFromMessage sets ThreadFromMessage field to given value.

### HasThreadFromMessage

`func (o *PlatformTargetPlatformSpecificData) HasThreadFromMessage() bool`

HasThreadFromMessage returns a boolean if a field has been set.

### GetTts

`func (o *PlatformTargetPlatformSpecificData) GetTts() bool`

GetTts returns the Tts field if non-nil, zero value otherwise.

### GetTtsOk

`func (o *PlatformTargetPlatformSpecificData) GetTtsOk() (*bool, bool)`

GetTtsOk returns a tuple with the Tts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTts

`func (o *PlatformTargetPlatformSpecificData) SetTts(v bool)`

SetTts sets Tts field to given value.

### HasTts

`func (o *PlatformTargetPlatformSpecificData) HasTts() bool`

HasTts returns a boolean if a field has been set.

### GetWebhookUsername

`func (o *PlatformTargetPlatformSpecificData) GetWebhookUsername() string`

GetWebhookUsername returns the WebhookUsername field if non-nil, zero value otherwise.

### GetWebhookUsernameOk

`func (o *PlatformTargetPlatformSpecificData) GetWebhookUsernameOk() (*string, bool)`

GetWebhookUsernameOk returns a tuple with the WebhookUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUsername

`func (o *PlatformTargetPlatformSpecificData) SetWebhookUsername(v string)`

SetWebhookUsername sets WebhookUsername field to given value.

### HasWebhookUsername

`func (o *PlatformTargetPlatformSpecificData) HasWebhookUsername() bool`

HasWebhookUsername returns a boolean if a field has been set.

### GetWebhookAvatarUrl

`func (o *PlatformTargetPlatformSpecificData) GetWebhookAvatarUrl() string`

GetWebhookAvatarUrl returns the WebhookAvatarUrl field if non-nil, zero value otherwise.

### GetWebhookAvatarUrlOk

`func (o *PlatformTargetPlatformSpecificData) GetWebhookAvatarUrlOk() (*string, bool)`

GetWebhookAvatarUrlOk returns a tuple with the WebhookAvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookAvatarUrl

`func (o *PlatformTargetPlatformSpecificData) SetWebhookAvatarUrl(v string)`

SetWebhookAvatarUrl sets WebhookAvatarUrl field to given value.

### HasWebhookAvatarUrl

`func (o *PlatformTargetPlatformSpecificData) HasWebhookAvatarUrl() bool`

HasWebhookAvatarUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


