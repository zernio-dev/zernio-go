# PostCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Post** | Pointer to [**Post**](Post.md) |  | [optional] 

## Methods

### NewPostCreateResponse

`func NewPostCreateResponse() *PostCreateResponse`

NewPostCreateResponse instantiates a new PostCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostCreateResponseWithDefaults

`func NewPostCreateResponseWithDefaults() *PostCreateResponse`

NewPostCreateResponseWithDefaults instantiates a new PostCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *PostCreateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *PostCreateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *PostCreateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *PostCreateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetPost

`func (o *PostCreateResponse) GetPost() Post`

GetPost returns the Post field if non-nil, zero value otherwise.

### GetPostOk

`func (o *PostCreateResponse) GetPostOk() (*Post, bool)`

GetPostOk returns a tuple with the Post field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPost

`func (o *PostCreateResponse) SetPost(v Post)`

SetPost sets Post field to given value.

### HasPost

`func (o *PostCreateResponse) HasPost() bool`

HasPost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


