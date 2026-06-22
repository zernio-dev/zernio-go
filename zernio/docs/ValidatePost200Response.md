# ValidatePost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Warnings** | Pointer to [**[]ValidatePost200ResponseOneOfWarningsInner**](ValidatePost200ResponseOneOfWarningsInner.md) |  | [optional] 
**Errors** | Pointer to [**[]ValidatePost200ResponseOneOf1ErrorsInner**](ValidatePost200ResponseOneOf1ErrorsInner.md) |  | [optional] 

## Methods

### NewValidatePost200Response

`func NewValidatePost200Response() *ValidatePost200Response`

NewValidatePost200Response instantiates a new ValidatePost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidatePost200ResponseWithDefaults

`func NewValidatePost200ResponseWithDefaults() *ValidatePost200Response`

NewValidatePost200ResponseWithDefaults instantiates a new ValidatePost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *ValidatePost200Response) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ValidatePost200Response) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ValidatePost200Response) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ValidatePost200Response) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetMessage

`func (o *ValidatePost200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ValidatePost200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ValidatePost200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ValidatePost200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetWarnings

`func (o *ValidatePost200Response) GetWarnings() []ValidatePost200ResponseOneOfWarningsInner`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *ValidatePost200Response) GetWarningsOk() (*[]ValidatePost200ResponseOneOfWarningsInner, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *ValidatePost200Response) SetWarnings(v []ValidatePost200ResponseOneOfWarningsInner)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *ValidatePost200Response) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.

### GetErrors

`func (o *ValidatePost200Response) GetErrors() []ValidatePost200ResponseOneOf1ErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ValidatePost200Response) GetErrorsOk() (*[]ValidatePost200ResponseOneOf1ErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ValidatePost200Response) SetErrors(v []ValidatePost200ResponseOneOf1ErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ValidatePost200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


