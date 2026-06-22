# ValidateMedia200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Error** | Pointer to **string** | Error message if valid is false | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**Size** | Pointer to **NullableInt32** | File size in bytes | [optional] 
**SizeFormatted** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**PlatformLimits** | Pointer to [**map[string]ValidateMedia200ResponsePlatformLimitsValue**](ValidateMedia200ResponsePlatformLimitsValue.md) | Per-platform size limit comparison (only present when size and type are known) | [optional] 

## Methods

### NewValidateMedia200Response

`func NewValidateMedia200Response() *ValidateMedia200Response`

NewValidateMedia200Response instantiates a new ValidateMedia200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidateMedia200ResponseWithDefaults

`func NewValidateMedia200ResponseWithDefaults() *ValidateMedia200Response`

NewValidateMedia200ResponseWithDefaults instantiates a new ValidateMedia200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ValidateMedia200Response) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidateMedia200Response) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidateMedia200Response) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ValidateMedia200Response) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetUrl

`func (o *ValidateMedia200Response) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ValidateMedia200Response) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ValidateMedia200Response) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ValidateMedia200Response) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetError

`func (o *ValidateMedia200Response) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ValidateMedia200Response) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ValidateMedia200Response) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ValidateMedia200Response) HasError() bool`

HasError returns a boolean if a field has been set.

### GetContentType

`func (o *ValidateMedia200Response) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ValidateMedia200Response) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ValidateMedia200Response) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ValidateMedia200Response) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *ValidateMedia200Response) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ValidateMedia200Response) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ValidateMedia200Response) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *ValidateMedia200Response) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ValidateMedia200Response) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ValidateMedia200Response) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSizeFormatted

`func (o *ValidateMedia200Response) GetSizeFormatted() string`

GetSizeFormatted returns the SizeFormatted field if non-nil, zero value otherwise.

### GetSizeFormattedOk

`func (o *ValidateMedia200Response) GetSizeFormattedOk() (*string, bool)`

GetSizeFormattedOk returns a tuple with the SizeFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeFormatted

`func (o *ValidateMedia200Response) SetSizeFormatted(v string)`

SetSizeFormatted sets SizeFormatted field to given value.

### HasSizeFormatted

`func (o *ValidateMedia200Response) HasSizeFormatted() bool`

HasSizeFormatted returns a boolean if a field has been set.

### GetType

`func (o *ValidateMedia200Response) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ValidateMedia200Response) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ValidateMedia200Response) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ValidateMedia200Response) HasType() bool`

HasType returns a boolean if a field has been set.

### GetPlatformLimits

`func (o *ValidateMedia200Response) GetPlatformLimits() map[string]ValidateMedia200ResponsePlatformLimitsValue`

GetPlatformLimits returns the PlatformLimits field if non-nil, zero value otherwise.

### GetPlatformLimitsOk

`func (o *ValidateMedia200Response) GetPlatformLimitsOk() (*map[string]ValidateMedia200ResponsePlatformLimitsValue, bool)`

GetPlatformLimitsOk returns a tuple with the PlatformLimits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformLimits

`func (o *ValidateMedia200Response) SetPlatformLimits(v map[string]ValidateMedia200ResponsePlatformLimitsValue)`

SetPlatformLimits sets PlatformLimits field to given value.

### HasPlatformLimits

`func (o *ValidateMedia200Response) HasPlatformLimits() bool`

HasPlatformLimits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


