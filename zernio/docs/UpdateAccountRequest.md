# UpdateAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**XCapabilities** | Pointer to [**UpdateAccountRequestXCapabilities**](UpdateAccountRequestXCapabilities.md) |  | [optional] 

## Methods

### NewUpdateAccountRequest

`func NewUpdateAccountRequest() *UpdateAccountRequest`

NewUpdateAccountRequest instantiates a new UpdateAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAccountRequestWithDefaults

`func NewUpdateAccountRequestWithDefaults() *UpdateAccountRequest`

NewUpdateAccountRequestWithDefaults instantiates a new UpdateAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UpdateAccountRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateAccountRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateAccountRequest) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateAccountRequest) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetDisplayName

`func (o *UpdateAccountRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateAccountRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateAccountRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UpdateAccountRequest) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetXCapabilities

`func (o *UpdateAccountRequest) GetXCapabilities() UpdateAccountRequestXCapabilities`

GetXCapabilities returns the XCapabilities field if non-nil, zero value otherwise.

### GetXCapabilitiesOk

`func (o *UpdateAccountRequest) GetXCapabilitiesOk() (*UpdateAccountRequestXCapabilities, bool)`

GetXCapabilitiesOk returns a tuple with the XCapabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXCapabilities

`func (o *UpdateAccountRequest) SetXCapabilities(v UpdateAccountRequestXCapabilities)`

SetXCapabilities sets XCapabilities field to given value.

### HasXCapabilities

`func (o *UpdateAccountRequest) HasXCapabilities() bool`

HasXCapabilities returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


