# RedditPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Reddit post ID (without type prefix) | [optional] 
**Fullname** | Pointer to **string** | Reddit fullname (e.g. t3_abc123) | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Subreddit** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** | Post URL (may be a gallery URL, external link, or self-post URL) | [optional] 
**Permalink** | Pointer to **string** | Full permalink to the Reddit post | [optional] 
**Selftext** | Pointer to **string** | Self-post body text (empty string for link posts) | [optional] 
**CreatedUtc** | Pointer to **float32** | Unix timestamp of post creation | [optional] 
**Score** | Pointer to **int32** |  | [optional] 
**NumComments** | Pointer to **int32** |  | [optional] 
**Over18** | Pointer to **bool** | Whether the post is marked NSFW | [optional] 
**Stickied** | Pointer to **bool** |  | [optional] 
**FlairText** | Pointer to **NullableString** | Link flair text if set | [optional] 
**IsGallery** | Pointer to **bool** | Whether the post is a gallery with multiple images | [optional] 
**GalleryImages** | Pointer to **[]string** | Individual image URLs for gallery posts (only present when isGallery is true) | [optional] 

## Methods

### NewRedditPost

`func NewRedditPost() *RedditPost`

NewRedditPost instantiates a new RedditPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedditPostWithDefaults

`func NewRedditPostWithDefaults() *RedditPost`

NewRedditPostWithDefaults instantiates a new RedditPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RedditPost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RedditPost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RedditPost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RedditPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFullname

`func (o *RedditPost) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *RedditPost) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *RedditPost) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *RedditPost) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### GetTitle

`func (o *RedditPost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *RedditPost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *RedditPost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *RedditPost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetAuthor

`func (o *RedditPost) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *RedditPost) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *RedditPost) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *RedditPost) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetSubreddit

`func (o *RedditPost) GetSubreddit() string`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *RedditPost) GetSubredditOk() (*string, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *RedditPost) SetSubreddit(v string)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *RedditPost) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### GetUrl

`func (o *RedditPost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RedditPost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RedditPost) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RedditPost) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPermalink

`func (o *RedditPost) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *RedditPost) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *RedditPost) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *RedditPost) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetSelftext

`func (o *RedditPost) GetSelftext() string`

GetSelftext returns the Selftext field if non-nil, zero value otherwise.

### GetSelftextOk

`func (o *RedditPost) GetSelftextOk() (*string, bool)`

GetSelftextOk returns a tuple with the Selftext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelftext

`func (o *RedditPost) SetSelftext(v string)`

SetSelftext sets Selftext field to given value.

### HasSelftext

`func (o *RedditPost) HasSelftext() bool`

HasSelftext returns a boolean if a field has been set.

### GetCreatedUtc

`func (o *RedditPost) GetCreatedUtc() float32`

GetCreatedUtc returns the CreatedUtc field if non-nil, zero value otherwise.

### GetCreatedUtcOk

`func (o *RedditPost) GetCreatedUtcOk() (*float32, bool)`

GetCreatedUtcOk returns a tuple with the CreatedUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUtc

`func (o *RedditPost) SetCreatedUtc(v float32)`

SetCreatedUtc sets CreatedUtc field to given value.

### HasCreatedUtc

`func (o *RedditPost) HasCreatedUtc() bool`

HasCreatedUtc returns a boolean if a field has been set.

### GetScore

`func (o *RedditPost) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *RedditPost) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *RedditPost) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *RedditPost) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetNumComments

`func (o *RedditPost) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *RedditPost) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *RedditPost) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.

### HasNumComments

`func (o *RedditPost) HasNumComments() bool`

HasNumComments returns a boolean if a field has been set.

### GetOver18

`func (o *RedditPost) GetOver18() bool`

GetOver18 returns the Over18 field if non-nil, zero value otherwise.

### GetOver18Ok

`func (o *RedditPost) GetOver18Ok() (*bool, bool)`

GetOver18Ok returns a tuple with the Over18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver18

`func (o *RedditPost) SetOver18(v bool)`

SetOver18 sets Over18 field to given value.

### HasOver18

`func (o *RedditPost) HasOver18() bool`

HasOver18 returns a boolean if a field has been set.

### GetStickied

`func (o *RedditPost) GetStickied() bool`

GetStickied returns the Stickied field if non-nil, zero value otherwise.

### GetStickiedOk

`func (o *RedditPost) GetStickiedOk() (*bool, bool)`

GetStickiedOk returns a tuple with the Stickied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickied

`func (o *RedditPost) SetStickied(v bool)`

SetStickied sets Stickied field to given value.

### HasStickied

`func (o *RedditPost) HasStickied() bool`

HasStickied returns a boolean if a field has been set.

### GetFlairText

`func (o *RedditPost) GetFlairText() string`

GetFlairText returns the FlairText field if non-nil, zero value otherwise.

### GetFlairTextOk

`func (o *RedditPost) GetFlairTextOk() (*string, bool)`

GetFlairTextOk returns a tuple with the FlairText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairText

`func (o *RedditPost) SetFlairText(v string)`

SetFlairText sets FlairText field to given value.

### HasFlairText

`func (o *RedditPost) HasFlairText() bool`

HasFlairText returns a boolean if a field has been set.

### SetFlairTextNil

`func (o *RedditPost) SetFlairTextNil(b bool)`

 SetFlairTextNil sets the value for FlairText to be an explicit nil

### UnsetFlairText
`func (o *RedditPost) UnsetFlairText()`

UnsetFlairText ensures that no value is present for FlairText, not even an explicit nil
### GetIsGallery

`func (o *RedditPost) GetIsGallery() bool`

GetIsGallery returns the IsGallery field if non-nil, zero value otherwise.

### GetIsGalleryOk

`func (o *RedditPost) GetIsGalleryOk() (*bool, bool)`

GetIsGalleryOk returns a tuple with the IsGallery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGallery

`func (o *RedditPost) SetIsGallery(v bool)`

SetIsGallery sets IsGallery field to given value.

### HasIsGallery

`func (o *RedditPost) HasIsGallery() bool`

HasIsGallery returns a boolean if a field has been set.

### GetGalleryImages

`func (o *RedditPost) GetGalleryImages() []string`

GetGalleryImages returns the GalleryImages field if non-nil, zero value otherwise.

### GetGalleryImagesOk

`func (o *RedditPost) GetGalleryImagesOk() (*[]string, bool)`

GetGalleryImagesOk returns a tuple with the GalleryImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGalleryImages

`func (o *RedditPost) SetGalleryImages(v []string)`

SetGalleryImages sets GalleryImages field to given value.

### HasGalleryImages

`func (o *RedditPost) HasGalleryImages() bool`

HasGalleryImages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


