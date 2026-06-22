# FollowUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**TargetUserId** | Pointer to **string** |  | [optional] 
**Following** | Pointer to **bool** |  | [optional] 
**PendingFollow** | Pointer to **bool** | True if the target account is protected and a follow request was sent | [optional] 
**Platform** | Pointer to **string** |  | [optional] 

## Methods

### NewFollowUser200Response

`func NewFollowUser200Response() *FollowUser200Response`

NewFollowUser200Response instantiates a new FollowUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowUser200ResponseWithDefaults

`func NewFollowUser200ResponseWithDefaults() *FollowUser200Response`

NewFollowUser200ResponseWithDefaults instantiates a new FollowUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *FollowUser200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FollowUser200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FollowUser200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FollowUser200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTargetUserId

`func (o *FollowUser200Response) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *FollowUser200Response) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *FollowUser200Response) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.

### HasTargetUserId

`func (o *FollowUser200Response) HasTargetUserId() bool`

HasTargetUserId returns a boolean if a field has been set.

### GetFollowing

`func (o *FollowUser200Response) GetFollowing() bool`

GetFollowing returns the Following field if non-nil, zero value otherwise.

### GetFollowingOk

`func (o *FollowUser200Response) GetFollowingOk() (*bool, bool)`

GetFollowingOk returns a tuple with the Following field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowing

`func (o *FollowUser200Response) SetFollowing(v bool)`

SetFollowing sets Following field to given value.

### HasFollowing

`func (o *FollowUser200Response) HasFollowing() bool`

HasFollowing returns a boolean if a field has been set.

### GetPendingFollow

`func (o *FollowUser200Response) GetPendingFollow() bool`

GetPendingFollow returns the PendingFollow field if non-nil, zero value otherwise.

### GetPendingFollowOk

`func (o *FollowUser200Response) GetPendingFollowOk() (*bool, bool)`

GetPendingFollowOk returns a tuple with the PendingFollow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingFollow

`func (o *FollowUser200Response) SetPendingFollow(v bool)`

SetPendingFollow sets PendingFollow field to given value.

### HasPendingFollow

`func (o *FollowUser200Response) HasPendingFollow() bool`

HasPendingFollow returns a boolean if a field has been set.

### GetPlatform

`func (o *FollowUser200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *FollowUser200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *FollowUser200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *FollowUser200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


