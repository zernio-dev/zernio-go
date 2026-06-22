# TwitterPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReplyToTweetId** | Pointer to **string** | ID of an existing tweet to reply to. The published tweet will appear as a reply in that tweet&#39;s thread. For threads, only the first tweet replies to the target; subsequent tweets chain normally. | [optional] 
**QuoteTweetId** | Pointer to **string** | ID (or full status URL) of an existing tweet to quote-repost. The published tweet becomes a quote tweet of the target. Mutually exclusive with media and poll. X only permits quoting your own posts or posts you are mentioned in / part of the conversation thread of; quoting an arbitrary other account&#39;s post is rejected by X. Billed at the standard create rate ($0.015), unlike pasting a tweet URL into the text which is billed at the URL rate ($0.20). For threads, applies to the first tweet only. | [optional] 
**ReplySettings** | Pointer to **string** | Controls who can reply to the tweet. \&quot;following\&quot; allows only people you follow, \&quot;mentionedUsers\&quot; allows only mentioned users, \&quot;subscribers\&quot; allows only subscribers, \&quot;verified\&quot; allows only verified users. Omit for default (everyone can reply). For threads, applies to the first tweet only. Cannot be combined with replyToTweetId. | [optional] 
**ThreadItems** | Pointer to [**[]TwitterPlatformDataThreadItemsInner**](TwitterPlatformDataThreadItemsInner.md) | Complete sequence of tweets in a thread. The first item becomes the root tweet, subsequent items are chained as replies. When threadItems is provided, the top-level content field is used only for display and search purposes, it is NOT published. You must include your first tweet as threadItems[0].  | [optional] 
**Poll** | Pointer to [**TwitterPlatformDataPoll**](TwitterPlatformDataPoll.md) |  | [optional] 
**LongVideo** | Pointer to **bool** | Enable long video uploads (over 140 seconds) using amplify_video media category. Requires the connected X account to have an active X Premium subscription. When true, videos are uploaded with the amplify_video category which supports longer durations (up to 10 minutes via API). When false or omitted, the standard tweet_video category is used (140 second limit). Note that not all Premium accounts have API long-video access, as X may require separate allowlisting. | [optional] [default to false]
**GeoRestriction** | Pointer to [**GeoRestriction**](GeoRestriction.md) |  | [optional] 
**PaidPartnership** | Pointer to **bool** | When true, the post is labeled by X as a paid partnership / paid promotion. For threads, applies to the root tweet only. Field availability may depend on your X API access tier. | [optional] [default to false]
**MadeWithAi** | Pointer to **bool** | When true, the post is labeled by X as containing AI-generated media. Per X, this label is for AI-generated media, not AI-written text. For threads, applies to the root tweet only. | [optional] [default to false]
**SensitiveMedia** | Pointer to [**TwitterPlatformDataSensitiveMedia**](TwitterPlatformDataSensitiveMedia.md) |  | [optional] 

## Methods

### NewTwitterPlatformData

`func NewTwitterPlatformData() *TwitterPlatformData`

NewTwitterPlatformData instantiates a new TwitterPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTwitterPlatformDataWithDefaults

`func NewTwitterPlatformDataWithDefaults() *TwitterPlatformData`

NewTwitterPlatformDataWithDefaults instantiates a new TwitterPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReplyToTweetId

`func (o *TwitterPlatformData) GetReplyToTweetId() string`

GetReplyToTweetId returns the ReplyToTweetId field if non-nil, zero value otherwise.

### GetReplyToTweetIdOk

`func (o *TwitterPlatformData) GetReplyToTweetIdOk() (*string, bool)`

GetReplyToTweetIdOk returns a tuple with the ReplyToTweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyToTweetId

`func (o *TwitterPlatformData) SetReplyToTweetId(v string)`

SetReplyToTweetId sets ReplyToTweetId field to given value.

### HasReplyToTweetId

`func (o *TwitterPlatformData) HasReplyToTweetId() bool`

HasReplyToTweetId returns a boolean if a field has been set.

### GetQuoteTweetId

`func (o *TwitterPlatformData) GetQuoteTweetId() string`

GetQuoteTweetId returns the QuoteTweetId field if non-nil, zero value otherwise.

### GetQuoteTweetIdOk

`func (o *TwitterPlatformData) GetQuoteTweetIdOk() (*string, bool)`

GetQuoteTweetIdOk returns a tuple with the QuoteTweetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuoteTweetId

`func (o *TwitterPlatformData) SetQuoteTweetId(v string)`

SetQuoteTweetId sets QuoteTweetId field to given value.

### HasQuoteTweetId

`func (o *TwitterPlatformData) HasQuoteTweetId() bool`

HasQuoteTweetId returns a boolean if a field has been set.

### GetReplySettings

`func (o *TwitterPlatformData) GetReplySettings() string`

GetReplySettings returns the ReplySettings field if non-nil, zero value otherwise.

### GetReplySettingsOk

`func (o *TwitterPlatformData) GetReplySettingsOk() (*string, bool)`

GetReplySettingsOk returns a tuple with the ReplySettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplySettings

`func (o *TwitterPlatformData) SetReplySettings(v string)`

SetReplySettings sets ReplySettings field to given value.

### HasReplySettings

`func (o *TwitterPlatformData) HasReplySettings() bool`

HasReplySettings returns a boolean if a field has been set.

### GetThreadItems

`func (o *TwitterPlatformData) GetThreadItems() []TwitterPlatformDataThreadItemsInner`

GetThreadItems returns the ThreadItems field if non-nil, zero value otherwise.

### GetThreadItemsOk

`func (o *TwitterPlatformData) GetThreadItemsOk() (*[]TwitterPlatformDataThreadItemsInner, bool)`

GetThreadItemsOk returns a tuple with the ThreadItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreadItems

`func (o *TwitterPlatformData) SetThreadItems(v []TwitterPlatformDataThreadItemsInner)`

SetThreadItems sets ThreadItems field to given value.

### HasThreadItems

`func (o *TwitterPlatformData) HasThreadItems() bool`

HasThreadItems returns a boolean if a field has been set.

### GetPoll

`func (o *TwitterPlatformData) GetPoll() TwitterPlatformDataPoll`

GetPoll returns the Poll field if non-nil, zero value otherwise.

### GetPollOk

`func (o *TwitterPlatformData) GetPollOk() (*TwitterPlatformDataPoll, bool)`

GetPollOk returns a tuple with the Poll field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoll

`func (o *TwitterPlatformData) SetPoll(v TwitterPlatformDataPoll)`

SetPoll sets Poll field to given value.

### HasPoll

`func (o *TwitterPlatformData) HasPoll() bool`

HasPoll returns a boolean if a field has been set.

### GetLongVideo

`func (o *TwitterPlatformData) GetLongVideo() bool`

GetLongVideo returns the LongVideo field if non-nil, zero value otherwise.

### GetLongVideoOk

`func (o *TwitterPlatformData) GetLongVideoOk() (*bool, bool)`

GetLongVideoOk returns a tuple with the LongVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongVideo

`func (o *TwitterPlatformData) SetLongVideo(v bool)`

SetLongVideo sets LongVideo field to given value.

### HasLongVideo

`func (o *TwitterPlatformData) HasLongVideo() bool`

HasLongVideo returns a boolean if a field has been set.

### GetGeoRestriction

`func (o *TwitterPlatformData) GetGeoRestriction() GeoRestriction`

GetGeoRestriction returns the GeoRestriction field if non-nil, zero value otherwise.

### GetGeoRestrictionOk

`func (o *TwitterPlatformData) GetGeoRestrictionOk() (*GeoRestriction, bool)`

GetGeoRestrictionOk returns a tuple with the GeoRestriction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoRestriction

`func (o *TwitterPlatformData) SetGeoRestriction(v GeoRestriction)`

SetGeoRestriction sets GeoRestriction field to given value.

### HasGeoRestriction

`func (o *TwitterPlatformData) HasGeoRestriction() bool`

HasGeoRestriction returns a boolean if a field has been set.

### GetPaidPartnership

`func (o *TwitterPlatformData) GetPaidPartnership() bool`

GetPaidPartnership returns the PaidPartnership field if non-nil, zero value otherwise.

### GetPaidPartnershipOk

`func (o *TwitterPlatformData) GetPaidPartnershipOk() (*bool, bool)`

GetPaidPartnershipOk returns a tuple with the PaidPartnership field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidPartnership

`func (o *TwitterPlatformData) SetPaidPartnership(v bool)`

SetPaidPartnership sets PaidPartnership field to given value.

### HasPaidPartnership

`func (o *TwitterPlatformData) HasPaidPartnership() bool`

HasPaidPartnership returns a boolean if a field has been set.

### GetMadeWithAi

`func (o *TwitterPlatformData) GetMadeWithAi() bool`

GetMadeWithAi returns the MadeWithAi field if non-nil, zero value otherwise.

### GetMadeWithAiOk

`func (o *TwitterPlatformData) GetMadeWithAiOk() (*bool, bool)`

GetMadeWithAiOk returns a tuple with the MadeWithAi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadeWithAi

`func (o *TwitterPlatformData) SetMadeWithAi(v bool)`

SetMadeWithAi sets MadeWithAi field to given value.

### HasMadeWithAi

`func (o *TwitterPlatformData) HasMadeWithAi() bool`

HasMadeWithAi returns a boolean if a field has been set.

### GetSensitiveMedia

`func (o *TwitterPlatformData) GetSensitiveMedia() TwitterPlatformDataSensitiveMedia`

GetSensitiveMedia returns the SensitiveMedia field if non-nil, zero value otherwise.

### GetSensitiveMediaOk

`func (o *TwitterPlatformData) GetSensitiveMediaOk() (*TwitterPlatformDataSensitiveMedia, bool)`

GetSensitiveMediaOk returns a tuple with the SensitiveMedia field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSensitiveMedia

`func (o *TwitterPlatformData) SetSensitiveMedia(v TwitterPlatformDataSensitiveMedia)`

SetSensitiveMedia sets SensitiveMedia field to given value.

### HasSensitiveMedia

`func (o *TwitterPlatformData) HasSensitiveMedia() bool`

HasSensitiveMedia returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


