# AddConversionAssociations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | Pointer to **string** |  | [optional] 
**Succeeded** | Pointer to **[]string** | Numeric campaign IDs that were successfully associated. | [optional] 
**Failed** | Pointer to [**[]AddConversionAssociations200ResponseFailedInner**](AddConversionAssociations200ResponseFailedInner.md) |  | [optional] 

## Methods

### NewAddConversionAssociations200Response

`func NewAddConversionAssociations200Response() *AddConversionAssociations200Response`

NewAddConversionAssociations200Response instantiates a new AddConversionAssociations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddConversionAssociations200ResponseWithDefaults

`func NewAddConversionAssociations200ResponseWithDefaults() *AddConversionAssociations200Response`

NewAddConversionAssociations200ResponseWithDefaults instantiates a new AddConversionAssociations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *AddConversionAssociations200Response) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddConversionAssociations200Response) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddConversionAssociations200Response) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AddConversionAssociations200Response) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSucceeded

`func (o *AddConversionAssociations200Response) GetSucceeded() []string`

GetSucceeded returns the Succeeded field if non-nil, zero value otherwise.

### GetSucceededOk

`func (o *AddConversionAssociations200Response) GetSucceededOk() (*[]string, bool)`

GetSucceededOk returns a tuple with the Succeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSucceeded

`func (o *AddConversionAssociations200Response) SetSucceeded(v []string)`

SetSucceeded sets Succeeded field to given value.

### HasSucceeded

`func (o *AddConversionAssociations200Response) HasSucceeded() bool`

HasSucceeded returns a boolean if a field has been set.

### GetFailed

`func (o *AddConversionAssociations200Response) GetFailed() []AddConversionAssociations200ResponseFailedInner`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *AddConversionAssociations200Response) GetFailedOk() (*[]AddConversionAssociations200ResponseFailedInner, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *AddConversionAssociations200Response) SetFailed(v []AddConversionAssociations200ResponseFailedInner)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *AddConversionAssociations200Response) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


