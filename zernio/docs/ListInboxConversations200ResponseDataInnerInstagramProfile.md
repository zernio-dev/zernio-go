# ListInboxConversations200ResponseDataInnerInstagramProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFollower** | Pointer to **NullableBool** | Whether the participant follows your Instagram business account | [optional] 
**IsFollowing** | Pointer to **NullableBool** | Whether your Instagram business account follows the participant | [optional] 
**FollowerCount** | Pointer to **NullableInt32** | The participant&#39;s follower count on Instagram | [optional] 
**IsVerified** | Pointer to **NullableBool** | Whether the participant is a verified Instagram user | [optional] 
**FetchedAt** | Pointer to **NullableTime** | When this profile data was last fetched from Instagram | [optional] 

## Methods

### NewListInboxConversations200ResponseDataInnerInstagramProfile

`func NewListInboxConversations200ResponseDataInnerInstagramProfile() *ListInboxConversations200ResponseDataInnerInstagramProfile`

NewListInboxConversations200ResponseDataInnerInstagramProfile instantiates a new ListInboxConversations200ResponseDataInnerInstagramProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInboxConversations200ResponseDataInnerInstagramProfileWithDefaults

`func NewListInboxConversations200ResponseDataInnerInstagramProfileWithDefaults() *ListInboxConversations200ResponseDataInnerInstagramProfile`

NewListInboxConversations200ResponseDataInnerInstagramProfileWithDefaults instantiates a new ListInboxConversations200ResponseDataInnerInstagramProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFollower

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetIsFollower() bool`

GetIsFollower returns the IsFollower field if non-nil, zero value otherwise.

### GetIsFollowerOk

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetIsFollowerOk() (*bool, bool)`

GetIsFollowerOk returns a tuple with the IsFollower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFollower

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetIsFollower(v bool)`

SetIsFollower sets IsFollower field to given value.

### HasIsFollower

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) HasIsFollower() bool`

HasIsFollower returns a boolean if a field has been set.

### SetIsFollowerNil

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetIsFollowerNil(b bool)`

 SetIsFollowerNil sets the value for IsFollower to be an explicit nil

### UnsetIsFollower
`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) UnsetIsFollower()`

UnsetIsFollower ensures that no value is present for IsFollower, not even an explicit nil
### GetIsFollowing

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetIsFollowing() bool`

GetIsFollowing returns the IsFollowing field if non-nil, zero value otherwise.

### GetIsFollowingOk

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetIsFollowingOk() (*bool, bool)`

GetIsFollowingOk returns a tuple with the IsFollowing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFollowing

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetIsFollowing(v bool)`

SetIsFollowing sets IsFollowing field to given value.

### HasIsFollowing

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) HasIsFollowing() bool`

HasIsFollowing returns a boolean if a field has been set.

### SetIsFollowingNil

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetIsFollowingNil(b bool)`

 SetIsFollowingNil sets the value for IsFollowing to be an explicit nil

### UnsetIsFollowing
`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) UnsetIsFollowing()`

UnsetIsFollowing ensures that no value is present for IsFollowing, not even an explicit nil
### GetFollowerCount

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetFollowerCount() int32`

GetFollowerCount returns the FollowerCount field if non-nil, zero value otherwise.

### GetFollowerCountOk

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetFollowerCountOk() (*int32, bool)`

GetFollowerCountOk returns a tuple with the FollowerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowerCount

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetFollowerCount(v int32)`

SetFollowerCount sets FollowerCount field to given value.

### HasFollowerCount

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) HasFollowerCount() bool`

HasFollowerCount returns a boolean if a field has been set.

### SetFollowerCountNil

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetFollowerCountNil(b bool)`

 SetFollowerCountNil sets the value for FollowerCount to be an explicit nil

### UnsetFollowerCount
`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) UnsetFollowerCount()`

UnsetFollowerCount ensures that no value is present for FollowerCount, not even an explicit nil
### GetIsVerified

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetIsVerified() bool`

GetIsVerified returns the IsVerified field if non-nil, zero value otherwise.

### GetIsVerifiedOk

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetIsVerifiedOk() (*bool, bool)`

GetIsVerifiedOk returns a tuple with the IsVerified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVerified

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetIsVerified(v bool)`

SetIsVerified sets IsVerified field to given value.

### HasIsVerified

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) HasIsVerified() bool`

HasIsVerified returns a boolean if a field has been set.

### SetIsVerifiedNil

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetIsVerifiedNil(b bool)`

 SetIsVerifiedNil sets the value for IsVerified to be an explicit nil

### UnsetIsVerified
`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) UnsetIsVerified()`

UnsetIsVerified ensures that no value is present for IsVerified, not even an explicit nil
### GetFetchedAt

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.

### HasFetchedAt

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) HasFetchedAt() bool`

HasFetchedAt returns a boolean if a field has been set.

### SetFetchedAtNil

`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) SetFetchedAtNil(b bool)`

 SetFetchedAtNil sets the value for FetchedAt to be an explicit nil

### UnsetFetchedAt
`func (o *ListInboxConversations200ResponseDataInnerInstagramProfile) UnsetFetchedAt()`

UnsetFetchedAt ensures that no value is present for FetchedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


