# CreatePost409Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Details** | Pointer to [**CreatePost409ResponseDetails**](CreatePost409ResponseDetails.md) |  | [optional] 

## Methods

### NewCreatePost409Response

`func NewCreatePost409Response() *CreatePost409Response`

NewCreatePost409Response instantiates a new CreatePost409Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePost409ResponseWithDefaults

`func NewCreatePost409ResponseWithDefaults() *CreatePost409Response`

NewCreatePost409ResponseWithDefaults instantiates a new CreatePost409Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CreatePost409Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreatePost409Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreatePost409Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CreatePost409Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetDetails

`func (o *CreatePost409Response) GetDetails() CreatePost409ResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreatePost409Response) GetDetailsOk() (*CreatePost409ResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreatePost409Response) SetDetails(v CreatePost409ResponseDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CreatePost409Response) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


