# LikeInboxComment200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**CommentId** | Pointer to **string** |  | [optional] 
**Liked** | Pointer to **bool** |  | [optional] 
**LikeUri** | Pointer to **string** | (Bluesky only) URI to use for unliking | [optional] 
**Platform** | Pointer to **string** |  | [optional] 

## Methods

### NewLikeInboxComment200Response

`func NewLikeInboxComment200Response() *LikeInboxComment200Response`

NewLikeInboxComment200Response instantiates a new LikeInboxComment200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLikeInboxComment200ResponseWithDefaults

`func NewLikeInboxComment200ResponseWithDefaults() *LikeInboxComment200Response`

NewLikeInboxComment200ResponseWithDefaults instantiates a new LikeInboxComment200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *LikeInboxComment200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LikeInboxComment200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LikeInboxComment200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LikeInboxComment200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCommentId

`func (o *LikeInboxComment200Response) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *LikeInboxComment200Response) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *LikeInboxComment200Response) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *LikeInboxComment200Response) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetLiked

`func (o *LikeInboxComment200Response) GetLiked() bool`

GetLiked returns the Liked field if non-nil, zero value otherwise.

### GetLikedOk

`func (o *LikeInboxComment200Response) GetLikedOk() (*bool, bool)`

GetLikedOk returns a tuple with the Liked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiked

`func (o *LikeInboxComment200Response) SetLiked(v bool)`

SetLiked sets Liked field to given value.

### HasLiked

`func (o *LikeInboxComment200Response) HasLiked() bool`

HasLiked returns a boolean if a field has been set.

### GetLikeUri

`func (o *LikeInboxComment200Response) GetLikeUri() string`

GetLikeUri returns the LikeUri field if non-nil, zero value otherwise.

### GetLikeUriOk

`func (o *LikeInboxComment200Response) GetLikeUriOk() (*string, bool)`

GetLikeUriOk returns a tuple with the LikeUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeUri

`func (o *LikeInboxComment200Response) SetLikeUri(v string)`

SetLikeUri sets LikeUri field to given value.

### HasLikeUri

`func (o *LikeInboxComment200Response) HasLikeUri() bool`

HasLikeUri returns a boolean if a field has been set.

### GetPlatform

`func (o *LikeInboxComment200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *LikeInboxComment200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *LikeInboxComment200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *LikeInboxComment200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


