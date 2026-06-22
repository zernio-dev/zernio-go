# ListInboxComments200ResponseDataInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AccountUsername** | Pointer to **string** |  | [optional] 
**Content** | Pointer to **string** |  | [optional] 
**Picture** | Pointer to **NullableString** |  | [optional] 
**Permalink** | Pointer to **NullableString** |  | [optional] 
**CreatedTime** | Pointer to **time.Time** |  | [optional] 
**CommentCount** | Pointer to **int32** |  | [optional] 
**LikeCount** | Pointer to **int32** |  | [optional] 
**Cid** | Pointer to **NullableString** | Bluesky content identifier | [optional] 
**Subreddit** | Pointer to **NullableString** | Reddit subreddit name | [optional] 
**IsAd** | Pointer to **bool** | True when this row is an ad (boosted/dark post). &#x60;platform&#x60; is then the placement (facebook &#x3D; the Page dark post / instagram &#x3D; the IG media), &#x60;id&#x60; is &#x60;{adId}:{placement}&#x60;, and the thread is at GET /v1/ads/{adId}/comments?placement&#x3D;{placement}. | [optional] 
**AdId** | Pointer to **string** | Internal Zernio ad id — only on ad rows. | [optional] 
**Placement** | Pointer to **string** | Which side of the ad this row&#39;s comments are on — only on ad rows. | [optional] 

## Methods

### NewListInboxComments200ResponseDataInner

`func NewListInboxComments200ResponseDataInner() *ListInboxComments200ResponseDataInner`

NewListInboxComments200ResponseDataInner instantiates a new ListInboxComments200ResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxComments200ResponseDataInnerWithDefaults

`func NewListInboxComments200ResponseDataInnerWithDefaults() *ListInboxComments200ResponseDataInner`

NewListInboxComments200ResponseDataInnerWithDefaults instantiates a new ListInboxComments200ResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInboxComments200ResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInboxComments200ResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInboxComments200ResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListInboxComments200ResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListInboxComments200ResponseDataInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInboxComments200ResponseDataInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInboxComments200ResponseDataInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInboxComments200ResponseDataInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *ListInboxComments200ResponseDataInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListInboxComments200ResponseDataInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListInboxComments200ResponseDataInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListInboxComments200ResponseDataInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountUsername

`func (o *ListInboxComments200ResponseDataInner) GetAccountUsername() string`

GetAccountUsername returns the AccountUsername field if non-nil, zero value otherwise.

### GetAccountUsernameOk

`func (o *ListInboxComments200ResponseDataInner) GetAccountUsernameOk() (*string, bool)`

GetAccountUsernameOk returns a tuple with the AccountUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountUsername

`func (o *ListInboxComments200ResponseDataInner) SetAccountUsername(v string)`

SetAccountUsername sets AccountUsername field to given value.

### HasAccountUsername

`func (o *ListInboxComments200ResponseDataInner) HasAccountUsername() bool`

HasAccountUsername returns a boolean if a field has been set.

### GetContent

`func (o *ListInboxComments200ResponseDataInner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *ListInboxComments200ResponseDataInner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *ListInboxComments200ResponseDataInner) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *ListInboxComments200ResponseDataInner) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetPicture

`func (o *ListInboxComments200ResponseDataInner) GetPicture() string`

GetPicture returns the Picture field if non-nil, zero value otherwise.

### GetPictureOk

`func (o *ListInboxComments200ResponseDataInner) GetPictureOk() (*string, bool)`

GetPictureOk returns a tuple with the Picture field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPicture

`func (o *ListInboxComments200ResponseDataInner) SetPicture(v string)`

SetPicture sets Picture field to given value.

### HasPicture

`func (o *ListInboxComments200ResponseDataInner) HasPicture() bool`

HasPicture returns a boolean if a field has been set.

### SetPictureNil

`func (o *ListInboxComments200ResponseDataInner) SetPictureNil(b bool)`

 SetPictureNil sets the value for Picture to be an explicit nil

### UnsetPicture
`func (o *ListInboxComments200ResponseDataInner) UnsetPicture()`

UnsetPicture ensures that no value is present for Picture, not even an explicit nil
### GetPermalink

`func (o *ListInboxComments200ResponseDataInner) GetPermalink() string`

GetPermalink returns the Permalink field if non-nil, zero value otherwise.

### GetPermalinkOk

`func (o *ListInboxComments200ResponseDataInner) GetPermalinkOk() (*string, bool)`

GetPermalinkOk returns a tuple with the Permalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermalink

`func (o *ListInboxComments200ResponseDataInner) SetPermalink(v string)`

SetPermalink sets Permalink field to given value.

### HasPermalink

`func (o *ListInboxComments200ResponseDataInner) HasPermalink() bool`

HasPermalink returns a boolean if a field has been set.

### SetPermalinkNil

`func (o *ListInboxComments200ResponseDataInner) SetPermalinkNil(b bool)`

 SetPermalinkNil sets the value for Permalink to be an explicit nil

### UnsetPermalink
`func (o *ListInboxComments200ResponseDataInner) UnsetPermalink()`

UnsetPermalink ensures that no value is present for Permalink, not even an explicit nil
### GetCreatedTime

`func (o *ListInboxComments200ResponseDataInner) GetCreatedTime() time.Time`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ListInboxComments200ResponseDataInner) GetCreatedTimeOk() (*time.Time, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ListInboxComments200ResponseDataInner) SetCreatedTime(v time.Time)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ListInboxComments200ResponseDataInner) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCommentCount

`func (o *ListInboxComments200ResponseDataInner) GetCommentCount() int32`

GetCommentCount returns the CommentCount field if non-nil, zero value otherwise.

### GetCommentCountOk

`func (o *ListInboxComments200ResponseDataInner) GetCommentCountOk() (*int32, bool)`

GetCommentCountOk returns a tuple with the CommentCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentCount

`func (o *ListInboxComments200ResponseDataInner) SetCommentCount(v int32)`

SetCommentCount sets CommentCount field to given value.

### HasCommentCount

`func (o *ListInboxComments200ResponseDataInner) HasCommentCount() bool`

HasCommentCount returns a boolean if a field has been set.

### GetLikeCount

`func (o *ListInboxComments200ResponseDataInner) GetLikeCount() int32`

GetLikeCount returns the LikeCount field if non-nil, zero value otherwise.

### GetLikeCountOk

`func (o *ListInboxComments200ResponseDataInner) GetLikeCountOk() (*int32, bool)`

GetLikeCountOk returns a tuple with the LikeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeCount

`func (o *ListInboxComments200ResponseDataInner) SetLikeCount(v int32)`

SetLikeCount sets LikeCount field to given value.

### HasLikeCount

`func (o *ListInboxComments200ResponseDataInner) HasLikeCount() bool`

HasLikeCount returns a boolean if a field has been set.

### GetCid

`func (o *ListInboxComments200ResponseDataInner) GetCid() string`

GetCid returns the Cid field if non-nil, zero value otherwise.

### GetCidOk

`func (o *ListInboxComments200ResponseDataInner) GetCidOk() (*string, bool)`

GetCidOk returns a tuple with the Cid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCid

`func (o *ListInboxComments200ResponseDataInner) SetCid(v string)`

SetCid sets Cid field to given value.

### HasCid

`func (o *ListInboxComments200ResponseDataInner) HasCid() bool`

HasCid returns a boolean if a field has been set.

### SetCidNil

`func (o *ListInboxComments200ResponseDataInner) SetCidNil(b bool)`

 SetCidNil sets the value for Cid to be an explicit nil

### UnsetCid
`func (o *ListInboxComments200ResponseDataInner) UnsetCid()`

UnsetCid ensures that no value is present for Cid, not even an explicit nil
### GetSubreddit

`func (o *ListInboxComments200ResponseDataInner) GetSubreddit() string`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *ListInboxComments200ResponseDataInner) GetSubredditOk() (*string, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *ListInboxComments200ResponseDataInner) SetSubreddit(v string)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *ListInboxComments200ResponseDataInner) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### SetSubredditNil

`func (o *ListInboxComments200ResponseDataInner) SetSubredditNil(b bool)`

 SetSubredditNil sets the value for Subreddit to be an explicit nil

### UnsetSubreddit
`func (o *ListInboxComments200ResponseDataInner) UnsetSubreddit()`

UnsetSubreddit ensures that no value is present for Subreddit, not even an explicit nil
### GetIsAd

`func (o *ListInboxComments200ResponseDataInner) GetIsAd() bool`

GetIsAd returns the IsAd field if non-nil, zero value otherwise.

### GetIsAdOk

`func (o *ListInboxComments200ResponseDataInner) GetIsAdOk() (*bool, bool)`

GetIsAdOk returns a tuple with the IsAd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAd

`func (o *ListInboxComments200ResponseDataInner) SetIsAd(v bool)`

SetIsAd sets IsAd field to given value.

### HasIsAd

`func (o *ListInboxComments200ResponseDataInner) HasIsAd() bool`

HasIsAd returns a boolean if a field has been set.

### GetAdId

`func (o *ListInboxComments200ResponseDataInner) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *ListInboxComments200ResponseDataInner) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *ListInboxComments200ResponseDataInner) SetAdId(v string)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *ListInboxComments200ResponseDataInner) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### GetPlacement

`func (o *ListInboxComments200ResponseDataInner) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *ListInboxComments200ResponseDataInner) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *ListInboxComments200ResponseDataInner) SetPlacement(v string)`

SetPlacement sets Placement field to given value.

### HasPlacement

`func (o *ListInboxComments200ResponseDataInner) HasPlacement() bool`

HasPlacement returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


