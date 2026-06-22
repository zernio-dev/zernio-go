# RedditPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Subreddit** | Pointer to **string** | Target subreddit name (without \&quot;r/\&quot; prefix). Overrides the default. Use GET /v1/accounts/{id}/reddit-subreddits to list options. | [optional] 
**Title** | Pointer to **string** | Post title. Defaults to the first line of content, truncated to 300 characters. | [optional] 
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

## Methods

### NewRedditPlatformData

`func NewRedditPlatformData() *RedditPlatformData`

NewRedditPlatformData instantiates a new RedditPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedditPlatformDataWithDefaults

`func NewRedditPlatformDataWithDefaults() *RedditPlatformData`

NewRedditPlatformDataWithDefaults instantiates a new RedditPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubreddit

`func (o *RedditPlatformData) GetSubreddit() string`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *RedditPlatformData) GetSubredditOk() (*string, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *RedditPlatformData) SetSubreddit(v string)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *RedditPlatformData) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### GetTitle

`func (o *RedditPlatformData) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RedditPlatformData) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RedditPlatformData) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RedditPlatformData) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *RedditPlatformData) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RedditPlatformData) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RedditPlatformData) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RedditPlatformData) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetForceSelf

`func (o *RedditPlatformData) GetForceSelf() bool`

GetForceSelf returns the ForceSelf field if non-nil, zero value otherwise.

### GetForceSelfOk

`func (o *RedditPlatformData) GetForceSelfOk() (*bool, bool)`

GetForceSelfOk returns a tuple with the ForceSelf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceSelf

`func (o *RedditPlatformData) SetForceSelf(v bool)`

SetForceSelf sets ForceSelf field to given value.

### HasForceSelf

`func (o *RedditPlatformData) HasForceSelf() bool`

HasForceSelf returns a boolean if a field has been set.

### GetFlairId

`func (o *RedditPlatformData) GetFlairId() string`

GetFlairId returns the FlairId field if non-nil, zero value otherwise.

### GetFlairIdOk

`func (o *RedditPlatformData) GetFlairIdOk() (*string, bool)`

GetFlairIdOk returns a tuple with the FlairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairId

`func (o *RedditPlatformData) SetFlairId(v string)`

SetFlairId sets FlairId field to given value.

### HasFlairId

`func (o *RedditPlatformData) HasFlairId() bool`

HasFlairId returns a boolean if a field has been set.

### GetFlairText

`func (o *RedditPlatformData) GetFlairText() string`

GetFlairText returns the FlairText field if non-nil, zero value otherwise.

### GetFlairTextOk

`func (o *RedditPlatformData) GetFlairTextOk() (*string, bool)`

GetFlairTextOk returns a tuple with the FlairText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairText

`func (o *RedditPlatformData) SetFlairText(v string)`

SetFlairText sets FlairText field to given value.

### HasFlairText

`func (o *RedditPlatformData) HasFlairText() bool`

HasFlairText returns a boolean if a field has been set.

### GetNsfw

`func (o *RedditPlatformData) GetNsfw() bool`

GetNsfw returns the Nsfw field if non-nil, zero value otherwise.

### GetNsfwOk

`func (o *RedditPlatformData) GetNsfwOk() (*bool, bool)`

GetNsfwOk returns a tuple with the Nsfw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNsfw

`func (o *RedditPlatformData) SetNsfw(v bool)`

SetNsfw sets Nsfw field to given value.

### HasNsfw

`func (o *RedditPlatformData) HasNsfw() bool`

HasNsfw returns a boolean if a field has been set.

### GetSpoiler

`func (o *RedditPlatformData) GetSpoiler() bool`

GetSpoiler returns the Spoiler field if non-nil, zero value otherwise.

### GetSpoilerOk

`func (o *RedditPlatformData) GetSpoilerOk() (*bool, bool)`

GetSpoilerOk returns a tuple with the Spoiler field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpoiler

`func (o *RedditPlatformData) SetSpoiler(v bool)`

SetSpoiler sets Spoiler field to given value.

### HasSpoiler

`func (o *RedditPlatformData) HasSpoiler() bool`

HasSpoiler returns a boolean if a field has been set.

### GetSendreplies

`func (o *RedditPlatformData) GetSendreplies() bool`

GetSendreplies returns the Sendreplies field if non-nil, zero value otherwise.

### GetSendrepliesOk

`func (o *RedditPlatformData) GetSendrepliesOk() (*bool, bool)`

GetSendrepliesOk returns a tuple with the Sendreplies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendreplies

`func (o *RedditPlatformData) SetSendreplies(v bool)`

SetSendreplies sets Sendreplies field to given value.

### HasSendreplies

`func (o *RedditPlatformData) HasSendreplies() bool`

HasSendreplies returns a boolean if a field has been set.

### GetNativeVideo

`func (o *RedditPlatformData) GetNativeVideo() bool`

GetNativeVideo returns the NativeVideo field if non-nil, zero value otherwise.

### GetNativeVideoOk

`func (o *RedditPlatformData) GetNativeVideoOk() (*bool, bool)`

GetNativeVideoOk returns a tuple with the NativeVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNativeVideo

`func (o *RedditPlatformData) SetNativeVideo(v bool)`

SetNativeVideo sets NativeVideo field to given value.

### HasNativeVideo

`func (o *RedditPlatformData) HasNativeVideo() bool`

HasNativeVideo returns a boolean if a field has been set.

### GetVideogif

`func (o *RedditPlatformData) GetVideogif() bool`

GetVideogif returns the Videogif field if non-nil, zero value otherwise.

### GetVideogifOk

`func (o *RedditPlatformData) GetVideogifOk() (*bool, bool)`

GetVideogifOk returns a tuple with the Videogif field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideogif

`func (o *RedditPlatformData) SetVideogif(v bool)`

SetVideogif sets Videogif field to given value.

### HasVideogif

`func (o *RedditPlatformData) HasVideogif() bool`

HasVideogif returns a boolean if a field has been set.

### GetVideoPosterUrl

`func (o *RedditPlatformData) GetVideoPosterUrl() string`

GetVideoPosterUrl returns the VideoPosterUrl field if non-nil, zero value otherwise.

### GetVideoPosterUrlOk

`func (o *RedditPlatformData) GetVideoPosterUrlOk() (*string, bool)`

GetVideoPosterUrlOk returns a tuple with the VideoPosterUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoPosterUrl

`func (o *RedditPlatformData) SetVideoPosterUrl(v string)`

SetVideoPosterUrl sets VideoPosterUrl field to given value.

### HasVideoPosterUrl

`func (o *RedditPlatformData) HasVideoPosterUrl() bool`

HasVideoPosterUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


