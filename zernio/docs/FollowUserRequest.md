# FollowUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The social account ID | 
**TargetUserId** | **string** | The Twitter ID of the user to follow | 

## Methods

### NewFollowUserRequest

`func NewFollowUserRequest(accountId string, targetUserId string, ) *FollowUserRequest`

NewFollowUserRequest instantiates a new FollowUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFollowUserRequestWithDefaults

`func NewFollowUserRequestWithDefaults() *FollowUserRequest`

NewFollowUserRequestWithDefaults instantiates a new FollowUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *FollowUserRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *FollowUserRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *FollowUserRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetTargetUserId

`func (o *FollowUserRequest) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *FollowUserRequest) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *FollowUserRequest) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


