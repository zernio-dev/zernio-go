# GetInboxPostComments200ResponsePost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Reddit post base36 id (e.g. \&quot;1tjtj26\&quot;) | [optional] 
**Fullname** | Pointer to **string** | Fullname with type prefix (e.g. \&quot;t3_1tjtj26\&quot;) | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Selftext** | Pointer to **string** | Body text for self-posts (empty for link posts) | [optional] 
**Author** | Pointer to **string** | Reddit username, without the u/ prefix | [optional] 
**Subreddit** | Pointer to **string** | Subreddit name, without the r/ prefix | [optional] 
**Permalink** | Pointer to **string** | Absolute URL to the post on reddit.com | [optional] 
**Url** | Pointer to **string** | For link posts, the external URL; for self-posts, the Reddit permalink | [optional] 
**Score** | Pointer to **int32** | Net upvotes (upvotes minus downvotes) | [optional] 
**NumComments** | Pointer to **int32** |  | [optional] 
**CreatedUtc** | Pointer to **int32** | Unix timestamp in seconds | [optional] 
**Over18** | Pointer to **bool** |  | [optional] 
**Stickied** | Pointer to **bool** |  | [optional] 
**FlairText** | Pointer to **NullableString** | Link flair text if any | [optional] 
**IsGallery** | Pointer to **bool** | True if the post is a Reddit gallery (multiple images) | [optional] 

## Methods

### NewGetInboxPostComments200ResponsePost

`func NewGetInboxPostComments200ResponsePost() *GetInboxPostComments200ResponsePost`

NewGetInboxPostComments200ResponsePost instantiates a new GetInboxPostComments200ResponsePost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxPostComments200ResponsePostWithDefaults

`func NewGetInboxPostComments200ResponsePostWithDefaults() *GetInboxPostComments200ResponsePost`

NewGetInboxPostComments200ResponsePostWithDefaults instantiates a new GetInboxPostComments200ResponsePost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInboxPostComments200ResponsePost) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxPostComments200ResponsePost) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxPostComments200ResponsePost) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetInboxPostComments200ResponsePost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFullname

`func (o *GetInboxPostComments200ResponsePost) GetFullname() string`

GetFullname returns the Fullname field if non-nil, zero value otherwise.

### GetFullnameOk

`func (o *GetInboxPostComments200ResponsePost) GetFullnameOk() (*string, bool)`

GetFullnameOk returns a tuple with the Fullname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullname

`func (o *GetInboxPostComments200ResponsePost) SetFullname(v string)`

SetFullname sets Fullname field to given value.

### HasFullname

`func (o *GetInboxPostComments200ResponsePost) HasFullname() bool`

HasFullname returns a boolean if a field has been set.

### GetTitle

`func (o *GetInboxPostComments200ResponsePost) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetInboxPostComments200ResponsePost) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetInboxPostComments200ResponsePost) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetInboxPostComments200ResponsePost) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetSelftext

`func (o *GetInboxPostComments200ResponsePost) GetSelftext() string`

GetSelftext returns the Selftext field if non-nil, zero value otherwise.

### GetSelftextOk

`func (o *GetInboxPostComments200ResponsePost) GetSelftextOk() (*string, bool)`

GetSelftextOk returns a tuple with the Selftext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelftext

`func (o *GetInboxPostComments200ResponsePost) SetSelftext(v string)`

SetSelftext sets Selftext field to given value.

### HasSelftext

`func (o *GetInboxPostComments200ResponsePost) HasSelftext() bool`

HasSelftext returns a boolean if a field has been set.

### GetAuthor

`func (o *GetInboxPostComments200ResponsePost) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetInboxPostComments200ResponsePost) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetInboxPostComments200ResponsePost) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetInboxPostComments200ResponsePost) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetSubreddit

`func (o *GetInboxPostComments200ResponsePost) GetSubreddit() string`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *GetInboxPostComments200ResponsePost) GetSubredditOk() (*string, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *GetInboxPostComments200ResponsePost) SetSubreddit(v string)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *GetInboxPostComments200ResponsePost) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### GetPermalink

`func (o *GetInboxPostComments200ResponsePost) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *GetInboxPostComments200ResponsePost) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *GetInboxPostComments200ResponsePost) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *GetInboxPostComments200ResponsePost) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### GetUrl

`func (o *GetInboxPostComments200ResponsePost) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetInboxPostComments200ResponsePost) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetInboxPostComments200ResponsePost) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetInboxPostComments200ResponsePost) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetScore

`func (o *GetInboxPostComments200ResponsePost) GetScore() int32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *GetInboxPostComments200ResponsePost) GetScoreOk() (*int32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *GetInboxPostComments200ResponsePost) SetScore(v int32)`

SetScore sets Score field to given value.

### HasScore

`func (o *GetInboxPostComments200ResponsePost) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetNumComments

`func (o *GetInboxPostComments200ResponsePost) GetNumComments() int32`

GetNumComments returns the NumComments field if non-nil, zero value otherwise.

### GetNumCommentsOk

`func (o *GetInboxPostComments200ResponsePost) GetNumCommentsOk() (*int32, bool)`

GetNumCommentsOk returns a tuple with the NumComments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumComments

`func (o *GetInboxPostComments200ResponsePost) SetNumComments(v int32)`

SetNumComments sets NumComments field to given value.

### HasNumComments

`func (o *GetInboxPostComments200ResponsePost) HasNumComments() bool`

HasNumComments returns a boolean if a field has been set.

### GetCreatedUtc

`func (o *GetInboxPostComments200ResponsePost) GetCreatedUtc() int32`

GetCreatedUtc returns the CreatedUtc field if non-nil, zero value otherwise.

### GetCreatedUtcOk

`func (o *GetInboxPostComments200ResponsePost) GetCreatedUtcOk() (*int32, bool)`

GetCreatedUtcOk returns a tuple with the CreatedUtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedUtc

`func (o *GetInboxPostComments200ResponsePost) SetCreatedUtc(v int32)`

SetCreatedUtc sets CreatedUtc field to given value.

### HasCreatedUtc

`func (o *GetInboxPostComments200ResponsePost) HasCreatedUtc() bool`

HasCreatedUtc returns a boolean if a field has been set.

### GetOver18

`func (o *GetInboxPostComments200ResponsePost) GetOver18() bool`

GetOver18 returns the Over18 field if non-nil, zero value otherwise.

### GetOver18Ok

`func (o *GetInboxPostComments200ResponsePost) GetOver18Ok() (*bool, bool)`

GetOver18Ok returns a tuple with the Over18 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOver18

`func (o *GetInboxPostComments200ResponsePost) SetOver18(v bool)`

SetOver18 sets Over18 field to given value.

### HasOver18

`func (o *GetInboxPostComments200ResponsePost) HasOver18() bool`

HasOver18 returns a boolean if a field has been set.

### GetStickied

`func (o *GetInboxPostComments200ResponsePost) GetStickied() bool`

GetStickied returns the Stickied field if non-nil, zero value otherwise.

### GetStickiedOk

`func (o *GetInboxPostComments200ResponsePost) GetStickiedOk() (*bool, bool)`

GetStickiedOk returns a tuple with the Stickied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickied

`func (o *GetInboxPostComments200ResponsePost) SetStickied(v bool)`

SetStickied sets Stickied field to given value.

### HasStickied

`func (o *GetInboxPostComments200ResponsePost) HasStickied() bool`

HasStickied returns a boolean if a field has been set.

### GetFlairText

`func (o *GetInboxPostComments200ResponsePost) GetFlairText() string`

GetFlairText returns the FlairText field if non-nil, zero value otherwise.

### GetFlairTextOk

`func (o *GetInboxPostComments200ResponsePost) GetFlairTextOk() (*string, bool)`

GetFlairTextOk returns a tuple with the FlairText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlairText

`func (o *GetInboxPostComments200ResponsePost) SetFlairText(v string)`

SetFlairText sets FlairText field to given value.

### HasFlairText

`func (o *GetInboxPostComments200ResponsePost) HasFlairText() bool`

HasFlairText returns a boolean if a field has been set.

### SetFlairTextNil

`func (o *GetInboxPostComments200ResponsePost) SetFlairTextNil(b bool)`

 SetFlairTextNil sets the value for FlairText to be an explicit nil

### UnsetFlairText
`func (o *GetInboxPostComments200ResponsePost) UnsetFlairText()`

UnsetFlairText ensures that no value is present for FlairText, not even an explicit nil
### GetIsGallery

`func (o *GetInboxPostComments200ResponsePost) GetIsGallery() bool`

GetIsGallery returns the IsGallery field if non-nil, zero value otherwise.

### GetIsGalleryOk

`func (o *GetInboxPostComments200ResponsePost) GetIsGalleryOk() (*bool, bool)`

GetIsGalleryOk returns a tuple with the IsGallery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGallery

`func (o *GetInboxPostComments200ResponsePost) SetIsGallery(v bool)`

SetIsGallery sets IsGallery field to given value.

### HasIsGallery

`func (o *GetInboxPostComments200ResponsePost) HasIsGallery() bool`

HasIsGallery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


