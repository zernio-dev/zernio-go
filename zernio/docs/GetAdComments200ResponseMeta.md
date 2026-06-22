# GetAdComments200ResponseMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** | Which side these comments are on (same as &#x60;placement&#x60;). | 
**Placement** | **string** | The placement these comments are for — useful when you didn&#39;t pass ?placement&#x3D; and want to know which one you got. | 
**AdId** | **string** | Internal Zernio ad ID. | 
**PlatformAdId** | **string** | Meta ad ID. | 
**EffectiveStoryId** | **string** | Underlying post ID the comments belong to. effective_object_story_id for the Facebook side, effective_instagram_media_id for the Instagram side. | 
**FacebookAccountId** | Pointer to **NullableString** | Facebook-only. The connected Facebook Page SocialAccount these comments were read through — pass it as &#x60;accountId&#x60; (with &#x60;effectiveStoryId&#x60; as the postId) to /v1/inbox/comments to reply/hide/delete. Null when no connected Page was used (then moderation isn&#39;t possible). | [optional] 
**InstagramUserId** | Pointer to **string** | Instagram-only. The Instagram-scoped business ID that owns the boosted media (creative.instagram_user_id). | [optional] 
**InstagramPermalink** | Pointer to **string** | Instagram-only. Public permalink of the boosted IG post (creative.instagram_permalink_url). | [optional] 
**InstagramAccountId** | Pointer to **string** | Instagram-only. The connected Instagram SocialAccount these comments were read through — pass it as &#x60;accountId&#x60; (with &#x60;effectiveStoryId&#x60; as the postId) to /v1/inbox/comments to reply/hide/delete. | [optional] 
**AccountId** | **string** | Social account ID (ads SocialAccount). | 
**LastUpdated** | **time.Time** |  | 

## Methods

### NewGetAdComments200ResponseMeta

`func NewGetAdComments200ResponseMeta(platform string, placement string, adId string, platformAdId string, effectiveStoryId string, accountId string, lastUpdated time.Time, ) *GetAdComments200ResponseMeta`

NewGetAdComments200ResponseMeta instantiates a new GetAdComments200ResponseMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdComments200ResponseMetaWithDefaults

`func NewGetAdComments200ResponseMetaWithDefaults() *GetAdComments200ResponseMeta`

NewGetAdComments200ResponseMetaWithDefaults instantiates a new GetAdComments200ResponseMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetAdComments200ResponseMeta) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAdComments200ResponseMeta) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAdComments200ResponseMeta) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlacement

`func (o *GetAdComments200ResponseMeta) GetPlacement() string`

GetPlacement returns the Placement field if non-nil, zero value otherwise.

### GetPlacementOk

`func (o *GetAdComments200ResponseMeta) GetPlacementOk() (*string, bool)`

GetPlacementOk returns a tuple with the Placement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlacement

`func (o *GetAdComments200ResponseMeta) SetPlacement(v string)`

SetPlacement sets Placement field to given value.


### GetAdId

`func (o *GetAdComments200ResponseMeta) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *GetAdComments200ResponseMeta) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *GetAdComments200ResponseMeta) SetAdId(v string)`

SetAdId sets AdId field to given value.


### GetPlatformAdId

`func (o *GetAdComments200ResponseMeta) GetPlatformAdId() string`

GetPlatformAdId returns the PlatformAdId field if non-nil, zero value otherwise.

### GetPlatformAdIdOk

`func (o *GetAdComments200ResponseMeta) GetPlatformAdIdOk() (*string, bool)`

GetPlatformAdIdOk returns a tuple with the PlatformAdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdId

`func (o *GetAdComments200ResponseMeta) SetPlatformAdId(v string)`

SetPlatformAdId sets PlatformAdId field to given value.


### GetEffectiveStoryId

`func (o *GetAdComments200ResponseMeta) GetEffectiveStoryId() string`

GetEffectiveStoryId returns the EffectiveStoryId field if non-nil, zero value otherwise.

### GetEffectiveStoryIdOk

`func (o *GetAdComments200ResponseMeta) GetEffectiveStoryIdOk() (*string, bool)`

GetEffectiveStoryIdOk returns a tuple with the EffectiveStoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveStoryId

`func (o *GetAdComments200ResponseMeta) SetEffectiveStoryId(v string)`

SetEffectiveStoryId sets EffectiveStoryId field to given value.


### GetFacebookAccountId

`func (o *GetAdComments200ResponseMeta) GetFacebookAccountId() string`

GetFacebookAccountId returns the FacebookAccountId field if non-nil, zero value otherwise.

### GetFacebookAccountIdOk

`func (o *GetAdComments200ResponseMeta) GetFacebookAccountIdOk() (*string, bool)`

GetFacebookAccountIdOk returns a tuple with the FacebookAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacebookAccountId

`func (o *GetAdComments200ResponseMeta) SetFacebookAccountId(v string)`

SetFacebookAccountId sets FacebookAccountId field to given value.

### HasFacebookAccountId

`func (o *GetAdComments200ResponseMeta) HasFacebookAccountId() bool`

HasFacebookAccountId returns a boolean if a field has been set.

### SetFacebookAccountIdNil

`func (o *GetAdComments200ResponseMeta) SetFacebookAccountIdNil(b bool)`

 SetFacebookAccountIdNil sets the value for FacebookAccountId to be an explicit nil

### UnsetFacebookAccountId
`func (o *GetAdComments200ResponseMeta) UnsetFacebookAccountId()`

UnsetFacebookAccountId ensures that no value is present for FacebookAccountId, not even an explicit nil
### GetInstagramUserId

`func (o *GetAdComments200ResponseMeta) GetInstagramUserId() string`

GetInstagramUserId returns the InstagramUserId field if non-nil, zero value otherwise.

### GetInstagramUserIdOk

`func (o *GetAdComments200ResponseMeta) GetInstagramUserIdOk() (*string, bool)`

GetInstagramUserIdOk returns a tuple with the InstagramUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramUserId

`func (o *GetAdComments200ResponseMeta) SetInstagramUserId(v string)`

SetInstagramUserId sets InstagramUserId field to given value.

### HasInstagramUserId

`func (o *GetAdComments200ResponseMeta) HasInstagramUserId() bool`

HasInstagramUserId returns a boolean if a field has been set.

### GetInstagramPermalink

`func (o *GetAdComments200ResponseMeta) GetInstagramPermalink() string`

GetInstagramPermalink returns the InstagramPermalink field if non-nil, zero value otherwise.

### GetInstagramPermalinkOk

`func (o *GetAdComments200ResponseMeta) GetInstagramPermalinkOk() (*string, bool)`

GetInstagramPermalinkOk returns a tuple with the InstagramPermalink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramPermalink

`func (o *GetAdComments200ResponseMeta) SetInstagramPermalink(v string)`

SetInstagramPermalink sets InstagramPermalink field to given value.

### HasInstagramPermalink

`func (o *GetAdComments200ResponseMeta) HasInstagramPermalink() bool`

HasInstagramPermalink returns a boolean if a field has been set.

### GetInstagramAccountId

`func (o *GetAdComments200ResponseMeta) GetInstagramAccountId() string`

GetInstagramAccountId returns the InstagramAccountId field if non-nil, zero value otherwise.

### GetInstagramAccountIdOk

`func (o *GetAdComments200ResponseMeta) GetInstagramAccountIdOk() (*string, bool)`

GetInstagramAccountIdOk returns a tuple with the InstagramAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstagramAccountId

`func (o *GetAdComments200ResponseMeta) SetInstagramAccountId(v string)`

SetInstagramAccountId sets InstagramAccountId field to given value.

### HasInstagramAccountId

`func (o *GetAdComments200ResponseMeta) HasInstagramAccountId() bool`

HasInstagramAccountId returns a boolean if a field has been set.

### GetAccountId

`func (o *GetAdComments200ResponseMeta) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAdComments200ResponseMeta) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAdComments200ResponseMeta) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetLastUpdated

`func (o *GetAdComments200ResponseMeta) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetAdComments200ResponseMeta) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetAdComments200ResponseMeta) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


