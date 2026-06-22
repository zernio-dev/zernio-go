# EditPost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Id** | Pointer to **string** | New tweet ID assigned by X after edit | [optional] 
**Url** | Pointer to **string** | URL of the edited tweet | [optional] 
**Message** | Pointer to **string** |  | [optional] 

## Methods

### NewEditPost200Response

`func NewEditPost200Response() *EditPost200Response`

NewEditPost200Response instantiates a new EditPost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditPost200ResponseWithDefaults

`func NewEditPost200ResponseWithDefaults() *EditPost200Response`

NewEditPost200ResponseWithDefaults instantiates a new EditPost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *EditPost200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *EditPost200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *EditPost200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *EditPost200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetId

`func (o *EditPost200Response) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EditPost200Response) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EditPost200Response) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EditPost200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *EditPost200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EditPost200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EditPost200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *EditPost200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetMessage

`func (o *EditPost200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EditPost200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EditPost200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EditPost200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


