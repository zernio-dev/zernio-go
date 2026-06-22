# UploadWhatsAppFlowJson200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ValidationErrors** | Pointer to [**[]UploadWhatsAppFlowJson200ResponseValidationErrorsInner**](UploadWhatsAppFlowJson200ResponseValidationErrorsInner.md) | Empty array if valid; otherwise, contains validation error details from Meta | [optional] 

## Methods

### NewUploadWhatsAppFlowJson200Response

`func NewUploadWhatsAppFlowJson200Response() *UploadWhatsAppFlowJson200Response`

NewUploadWhatsAppFlowJson200Response instantiates a new UploadWhatsAppFlowJson200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWhatsAppFlowJson200ResponseWithDefaults

`func NewUploadWhatsAppFlowJson200ResponseWithDefaults() *UploadWhatsAppFlowJson200Response`

NewUploadWhatsAppFlowJson200ResponseWithDefaults instantiates a new UploadWhatsAppFlowJson200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UploadWhatsAppFlowJson200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadWhatsAppFlowJson200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadWhatsAppFlowJson200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UploadWhatsAppFlowJson200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetValidationErrors

`func (o *UploadWhatsAppFlowJson200Response) GetValidationErrors() []UploadWhatsAppFlowJson200ResponseValidationErrorsInner`

GetValidationErrors returns the ValidationErrors field if non-nil, zero value otherwise.

### GetValidationErrorsOk

`func (o *UploadWhatsAppFlowJson200Response) GetValidationErrorsOk() (*[]UploadWhatsAppFlowJson200ResponseValidationErrorsInner, bool)`

GetValidationErrorsOk returns a tuple with the ValidationErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationErrors

`func (o *UploadWhatsAppFlowJson200Response) SetValidationErrors(v []UploadWhatsAppFlowJson200ResponseValidationErrorsInner)`

SetValidationErrors sets ValidationErrors field to given value.

### HasValidationErrors

`func (o *UploadWhatsAppFlowJson200Response) HasValidationErrors() bool`

HasValidationErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


