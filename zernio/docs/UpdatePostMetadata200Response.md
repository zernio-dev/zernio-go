# UpdatePostMetadata200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**VideoId** | Pointer to **string** | Only present in direct video ID mode | [optional] 
**UpdatedFields** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdatePostMetadata200Response

`func NewUpdatePostMetadata200Response() *UpdatePostMetadata200Response`

NewUpdatePostMetadata200Response instantiates a new UpdatePostMetadata200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePostMetadata200ResponseWithDefaults

`func NewUpdatePostMetadata200ResponseWithDefaults() *UpdatePostMetadata200Response`

NewUpdatePostMetadata200ResponseWithDefaults instantiates a new UpdatePostMetadata200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdatePostMetadata200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdatePostMetadata200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdatePostMetadata200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdatePostMetadata200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMessage

`func (o *UpdatePostMetadata200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdatePostMetadata200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdatePostMetadata200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdatePostMetadata200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetVideoId

`func (o *UpdatePostMetadata200Response) GetVideoId() string`

GetVideoId returns the VideoId field if non-nil, zero value otherwise.

### GetVideoIdOk

`func (o *UpdatePostMetadata200Response) GetVideoIdOk() (*string, bool)`

GetVideoIdOk returns a tuple with the VideoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoId

`func (o *UpdatePostMetadata200Response) SetVideoId(v string)`

SetVideoId sets VideoId field to given value.

### HasVideoId

`func (o *UpdatePostMetadata200Response) HasVideoId() bool`

HasVideoId returns a boolean if a field has been set.

### GetUpdatedFields

`func (o *UpdatePostMetadata200Response) GetUpdatedFields() []string`

GetUpdatedFields returns the UpdatedFields field if non-nil, zero value otherwise.

### GetUpdatedFieldsOk

`func (o *UpdatePostMetadata200Response) GetUpdatedFieldsOk() (*[]string, bool)`

GetUpdatedFieldsOk returns a tuple with the UpdatedFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedFields

`func (o *UpdatePostMetadata200Response) SetUpdatedFields(v []string)`

SetUpdatedFields sets UpdatedFields field to given value.

### HasUpdatedFields

`func (o *UpdatePostMetadata200Response) HasUpdatedFields() bool`

HasUpdatedFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


