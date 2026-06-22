# ValidateSubreddit200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** |  | [optional] 
**Subreddit** | Pointer to [**ValidateSubreddit200ResponseOneOfSubreddit**](ValidateSubreddit200ResponseOneOfSubreddit.md) |  | [optional] 
**Error** | Pointer to **string** |  | [optional] 

## Methods

### NewValidateSubreddit200Response

`func NewValidateSubreddit200Response() *ValidateSubreddit200Response`

NewValidateSubreddit200Response instantiates a new ValidateSubreddit200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateSubreddit200ResponseWithDefaults

`func NewValidateSubreddit200ResponseWithDefaults() *ValidateSubreddit200Response`

NewValidateSubreddit200ResponseWithDefaults instantiates a new ValidateSubreddit200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *ValidateSubreddit200Response) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *ValidateSubreddit200Response) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *ValidateSubreddit200Response) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *ValidateSubreddit200Response) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetSubreddit

`func (o *ValidateSubreddit200Response) GetSubreddit() ValidateSubreddit200ResponseOneOfSubreddit`

GetSubreddit returns the Subreddit field if non-nil, zero value otherwise.

### GetSubredditOk

`func (o *ValidateSubreddit200Response) GetSubredditOk() (*ValidateSubreddit200ResponseOneOfSubreddit, bool)`

GetSubredditOk returns a tuple with the Subreddit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubreddit

`func (o *ValidateSubreddit200Response) SetSubreddit(v ValidateSubreddit200ResponseOneOfSubreddit)`

SetSubreddit sets Subreddit field to given value.

### HasSubreddit

`func (o *ValidateSubreddit200Response) HasSubreddit() bool`

HasSubreddit returns a boolean if a field has been set.

### GetError

`func (o *ValidateSubreddit200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidateSubreddit200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidateSubreddit200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ValidateSubreddit200Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


