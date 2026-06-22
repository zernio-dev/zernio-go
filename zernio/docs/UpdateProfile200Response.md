# UpdateProfile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 

## Methods

### NewUpdateProfile200Response

`func NewUpdateProfile200Response() *UpdateProfile200Response`

NewUpdateProfile200Response instantiates a new UpdateProfile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateProfile200ResponseWithDefaults

`func NewUpdateProfile200ResponseWithDefaults() *UpdateProfile200Response`

NewUpdateProfile200ResponseWithDefaults instantiates a new UpdateProfile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *UpdateProfile200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UpdateProfile200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UpdateProfile200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UpdateProfile200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateProfile200Response) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateProfile200Response) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateProfile200Response) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UpdateProfile200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


