# CreateApiKey201Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**ApiKey** | Pointer to [**ApiKey**](ApiKey.md) |  | [optional] 

## Methods

### NewCreateApiKey201Response

`func NewCreateApiKey201Response() *CreateApiKey201Response`

NewCreateApiKey201Response instantiates a new CreateApiKey201Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateApiKey201ResponseWithDefaults

`func NewCreateApiKey201ResponseWithDefaults() *CreateApiKey201Response`

NewCreateApiKey201ResponseWithDefaults instantiates a new CreateApiKey201Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *CreateApiKey201Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateApiKey201Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateApiKey201Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CreateApiKey201Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetApiKey

`func (o *CreateApiKey201Response) GetApiKey() ApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *CreateApiKey201Response) GetApiKeyOk() (*ApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *CreateApiKey201Response) SetApiKey(v ApiKey)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *CreateApiKey201Response) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


