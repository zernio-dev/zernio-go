# ProfileCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Profile** | Pointer to [**Profile**](Profile.md) |  | [optional] 

## Methods

### NewProfileCreateResponse

`func NewProfileCreateResponse() *ProfileCreateResponse`

NewProfileCreateResponse instantiates a new ProfileCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProfileCreateResponseWithDefaults

`func NewProfileCreateResponseWithDefaults() *ProfileCreateResponse`

NewProfileCreateResponseWithDefaults instantiates a new ProfileCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ProfileCreateResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ProfileCreateResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ProfileCreateResponse) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ProfileCreateResponse) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetProfile

`func (o *ProfileCreateResponse) GetProfile() Profile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *ProfileCreateResponse) GetProfileOk() (*Profile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *ProfileCreateResponse) SetProfile(v Profile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *ProfileCreateResponse) HasProfile() bool`

HasProfile returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


