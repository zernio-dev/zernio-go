# GetCommentAutomation200ResponseLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**CommentId** | Pointer to **string** |  | [optional] 
**CommenterId** | Pointer to **string** |  | [optional] 
**CommenterName** | Pointer to **string** |  | [optional] 
**CommentText** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | DM outcome | [optional] 
**Error** | Pointer to **string** | DM error message if status is failed | [optional] 
**CommentReplyStatus** | Pointer to **string** | Outcome of the optional public reply on the triggering comment. &#39;skipped&#39; if no commentReply was configured or if the DM failed (the public reply is not attempted in that case). | [optional] 
**CommentReplyError** | Pointer to **string** | Public-reply error message if commentReplyStatus is failed | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetCommentAutomation200ResponseLogsInner

`func NewGetCommentAutomation200ResponseLogsInner() *GetCommentAutomation200ResponseLogsInner`

NewGetCommentAutomation200ResponseLogsInner instantiates a new GetCommentAutomation200ResponseLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentAutomation200ResponseLogsInnerWithDefaults

`func NewGetCommentAutomation200ResponseLogsInnerWithDefaults() *GetCommentAutomation200ResponseLogsInner`

NewGetCommentAutomation200ResponseLogsInnerWithDefaults instantiates a new GetCommentAutomation200ResponseLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCommentAutomation200ResponseLogsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCommentAutomation200ResponseLogsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GetCommentAutomation200ResponseLogsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCommentId

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentId() string`

GetCommentId returns the CommentId field if non-nil, zero value otherwise.

### GetCommentIdOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentIdOk() (*string, bool)`

GetCommentIdOk returns a tuple with the CommentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentId

`func (o *GetCommentAutomation200ResponseLogsInner) SetCommentId(v string)`

SetCommentId sets CommentId field to given value.

### HasCommentId

`func (o *GetCommentAutomation200ResponseLogsInner) HasCommentId() bool`

HasCommentId returns a boolean if a field has been set.

### GetCommenterId

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommenterId() string`

GetCommenterId returns the CommenterId field if non-nil, zero value otherwise.

### GetCommenterIdOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommenterIdOk() (*string, bool)`

GetCommenterIdOk returns a tuple with the CommenterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterId

`func (o *GetCommentAutomation200ResponseLogsInner) SetCommenterId(v string)`

SetCommenterId sets CommenterId field to given value.

### HasCommenterId

`func (o *GetCommentAutomation200ResponseLogsInner) HasCommenterId() bool`

HasCommenterId returns a boolean if a field has been set.

### GetCommenterName

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommenterName() string`

GetCommenterName returns the CommenterName field if non-nil, zero value otherwise.

### GetCommenterNameOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommenterNameOk() (*string, bool)`

GetCommenterNameOk returns a tuple with the CommenterName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommenterName

`func (o *GetCommentAutomation200ResponseLogsInner) SetCommenterName(v string)`

SetCommenterName sets CommenterName field to given value.

### HasCommenterName

`func (o *GetCommentAutomation200ResponseLogsInner) HasCommenterName() bool`

HasCommenterName returns a boolean if a field has been set.

### GetCommentText

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentText() string`

GetCommentText returns the CommentText field if non-nil, zero value otherwise.

### GetCommentTextOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentTextOk() (*string, bool)`

GetCommentTextOk returns a tuple with the CommentText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentText

`func (o *GetCommentAutomation200ResponseLogsInner) SetCommentText(v string)`

SetCommentText sets CommentText field to given value.

### HasCommentText

`func (o *GetCommentAutomation200ResponseLogsInner) HasCommentText() bool`

HasCommentText returns a boolean if a field has been set.

### GetStatus

`func (o *GetCommentAutomation200ResponseLogsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCommentAutomation200ResponseLogsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCommentAutomation200ResponseLogsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetError

`func (o *GetCommentAutomation200ResponseLogsInner) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GetCommentAutomation200ResponseLogsInner) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GetCommentAutomation200ResponseLogsInner) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCommentReplyStatus

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentReplyStatus() string`

GetCommentReplyStatus returns the CommentReplyStatus field if non-nil, zero value otherwise.

### GetCommentReplyStatusOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentReplyStatusOk() (*string, bool)`

GetCommentReplyStatusOk returns a tuple with the CommentReplyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReplyStatus

`func (o *GetCommentAutomation200ResponseLogsInner) SetCommentReplyStatus(v string)`

SetCommentReplyStatus sets CommentReplyStatus field to given value.

### HasCommentReplyStatus

`func (o *GetCommentAutomation200ResponseLogsInner) HasCommentReplyStatus() bool`

HasCommentReplyStatus returns a boolean if a field has been set.

### GetCommentReplyError

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentReplyError() string`

GetCommentReplyError returns the CommentReplyError field if non-nil, zero value otherwise.

### GetCommentReplyErrorOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCommentReplyErrorOk() (*string, bool)`

GetCommentReplyErrorOk returns a tuple with the CommentReplyError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentReplyError

`func (o *GetCommentAutomation200ResponseLogsInner) SetCommentReplyError(v string)`

SetCommentReplyError sets CommentReplyError field to given value.

### HasCommentReplyError

`func (o *GetCommentAutomation200ResponseLogsInner) HasCommentReplyError() bool`

HasCommentReplyError returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GetCommentAutomation200ResponseLogsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCommentAutomation200ResponseLogsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCommentAutomation200ResponseLogsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GetCommentAutomation200ResponseLogsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


