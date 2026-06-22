# ErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Human-readable error message. | [optional] 
**Type** | Pointer to **string** | Error class for programmatic handling. | [optional] 
**Code** | Pointer to **string** | Stable machine-readable error code. | [optional] 
**Param** | Pointer to **string** | The request field that caused the error, when applicable. | [optional] 
**Platform** | Pointer to **string** | Upstream platform (e.g. meta, google, tiktok) — present when type is platform_error. | [optional] 
**PlatformError** | Pointer to **map[string]interface{}** | Raw error payload from the upstream platform, passed through verbatim so integrators can read provider-specific codes. For Meta this includes error_subcode, error_user_title, and error_user_msg.  | [optional] 
**Details** | Pointer to **map[string]interface{}** | Additional structured context (e.g. field-level validation errors). | [optional] 

## Methods

### NewErrorResponse

`func NewErrorResponse() *ErrorResponse`

NewErrorResponse instantiates a new ErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWithDefaults

`func NewErrorResponseWithDefaults() *ErrorResponse`

NewErrorResponseWithDefaults instantiates a new ErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *ErrorResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *ErrorResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *ErrorResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *ErrorResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetType

`func (o *ErrorResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ErrorResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ErrorResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ErrorResponse) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCode

`func (o *ErrorResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ErrorResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetParam

`func (o *ErrorResponse) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *ErrorResponse) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *ErrorResponse) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *ErrorResponse) HasParam() bool`

HasParam returns a boolean if a field has been set.

### GetPlatform

`func (o *ErrorResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ErrorResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ErrorResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ErrorResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetPlatformError

`func (o *ErrorResponse) GetPlatformError() map[string]interface{}`

GetPlatformError returns the PlatformError field if non-nil, zero value otherwise.

### GetPlatformErrorOk

`func (o *ErrorResponse) GetPlatformErrorOk() (*map[string]interface{}, bool)`

GetPlatformErrorOk returns a tuple with the PlatformError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformError

`func (o *ErrorResponse) SetPlatformError(v map[string]interface{})`

SetPlatformError sets PlatformError field to given value.

### HasPlatformError

`func (o *ErrorResponse) HasPlatformError() bool`

HasPlatformError returns a boolean if a field has been set.

### GetDetails

`func (o *ErrorResponse) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponse) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponse) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponse) HasDetails() bool`

HasDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


