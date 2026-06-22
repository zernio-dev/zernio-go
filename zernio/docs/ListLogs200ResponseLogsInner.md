# ListLogs200ResponseLogsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Log category (publishing, connections, webhooks, messaging) | [optional] 
**Action** | Pointer to **string** | Specific action (post.published, message.sent, account.connected, etc.) | [optional] 
**UserId** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusCode** | Pointer to **int32** |  | [optional] 
**ErrorMessage** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**DurationMs** | Pointer to **int32** |  | [optional] 
**Endpoint** | Pointer to **string** | The API endpoint that triggered this log | [optional] 
**RequestBody** | Pointer to **string** | Request JSON (truncated to 5KB) | [optional] 
**ResponseBody** | Pointer to **string** | Response JSON (truncated to 10KB) | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**Metadata** | Pointer to **string** | Additional context as JSON string | [optional] 

## Methods

### NewListLogs200ResponseLogsInner

`func NewListLogs200ResponseLogsInner() *ListLogs200ResponseLogsInner`

NewListLogs200ResponseLogsInner instantiates a new ListLogs200ResponseLogsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLogs200ResponseLogsInnerWithDefaults

`func NewListLogs200ResponseLogsInnerWithDefaults() *ListLogs200ResponseLogsInner`

NewListLogs200ResponseLogsInnerWithDefaults instantiates a new ListLogs200ResponseLogsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ListLogs200ResponseLogsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListLogs200ResponseLogsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListLogs200ResponseLogsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListLogs200ResponseLogsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAction

`func (o *ListLogs200ResponseLogsInner) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListLogs200ResponseLogsInner) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListLogs200ResponseLogsInner) SetAction(v string)`

SetAction sets Action field to given value.

### HasAction

`func (o *ListLogs200ResponseLogsInner) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetUserId

`func (o *ListLogs200ResponseLogsInner) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListLogs200ResponseLogsInner) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListLogs200ResponseLogsInner) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListLogs200ResponseLogsInner) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPlatform

`func (o *ListLogs200ResponseLogsInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListLogs200ResponseLogsInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListLogs200ResponseLogsInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListLogs200ResponseLogsInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetAccountId

`func (o *ListLogs200ResponseLogsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListLogs200ResponseLogsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListLogs200ResponseLogsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListLogs200ResponseLogsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetStatus

`func (o *ListLogs200ResponseLogsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListLogs200ResponseLogsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListLogs200ResponseLogsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListLogs200ResponseLogsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCode

`func (o *ListLogs200ResponseLogsInner) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ListLogs200ResponseLogsInner) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ListLogs200ResponseLogsInner) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ListLogs200ResponseLogsInner) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetErrorMessage

`func (o *ListLogs200ResponseLogsInner) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListLogs200ResponseLogsInner) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListLogs200ResponseLogsInner) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListLogs200ResponseLogsInner) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetErrorCode

`func (o *ListLogs200ResponseLogsInner) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ListLogs200ResponseLogsInner) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ListLogs200ResponseLogsInner) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ListLogs200ResponseLogsInner) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetDurationMs

`func (o *ListLogs200ResponseLogsInner) GetDurationMs() int32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ListLogs200ResponseLogsInner) GetDurationMsOk() (*int32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ListLogs200ResponseLogsInner) SetDurationMs(v int32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ListLogs200ResponseLogsInner) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### GetEndpoint

`func (o *ListLogs200ResponseLogsInner) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *ListLogs200ResponseLogsInner) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *ListLogs200ResponseLogsInner) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *ListLogs200ResponseLogsInner) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetRequestBody

`func (o *ListLogs200ResponseLogsInner) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *ListLogs200ResponseLogsInner) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *ListLogs200ResponseLogsInner) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *ListLogs200ResponseLogsInner) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetResponseBody

`func (o *ListLogs200ResponseLogsInner) GetResponseBody() string`

GetResponseBody returns the ResponseBody field if non-nil, zero value otherwise.

### GetResponseBodyOk

`func (o *ListLogs200ResponseLogsInner) GetResponseBodyOk() (*string, bool)`

GetResponseBodyOk returns a tuple with the ResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBody

`func (o *ListLogs200ResponseLogsInner) SetResponseBody(v string)`

SetResponseBody sets ResponseBody field to given value.

### HasResponseBody

`func (o *ListLogs200ResponseLogsInner) HasResponseBody() bool`

HasResponseBody returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListLogs200ResponseLogsInner) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListLogs200ResponseLogsInner) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListLogs200ResponseLogsInner) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListLogs200ResponseLogsInner) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMetadata

`func (o *ListLogs200ResponseLogsInner) GetMetadata() string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ListLogs200ResponseLogsInner) GetMetadataOk() (*string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ListLogs200ResponseLogsInner) SetMetadata(v string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ListLogs200ResponseLogsInner) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


