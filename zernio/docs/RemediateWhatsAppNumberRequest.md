# RemediateWhatsAppNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Values** | Pointer to **map[string]string** |  | [optional] 
**Documents** | Pointer to [**[]SubmitWhatsAppNumberKycRequestDocumentsInner**](SubmitWhatsAppNumberKycRequestDocumentsInner.md) |  | [optional] 
**Address** | Pointer to **map[string]interface{}** | Same shape as the KYC submit address. | [optional] 

## Methods

### NewRemediateWhatsAppNumberRequest

`func NewRemediateWhatsAppNumberRequest() *RemediateWhatsAppNumberRequest`

NewRemediateWhatsAppNumberRequest instantiates a new RemediateWhatsAppNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemediateWhatsAppNumberRequestWithDefaults

`func NewRemediateWhatsAppNumberRequestWithDefaults() *RemediateWhatsAppNumberRequest`

NewRemediateWhatsAppNumberRequestWithDefaults instantiates a new RemediateWhatsAppNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValues

`func (o *RemediateWhatsAppNumberRequest) GetValues() map[string]string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RemediateWhatsAppNumberRequest) GetValuesOk() (*map[string]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RemediateWhatsAppNumberRequest) SetValues(v map[string]string)`

SetValues sets Values field to given value.

### HasValues

`func (o *RemediateWhatsAppNumberRequest) HasValues() bool`

HasValues returns a boolean if a field has been set.

### GetDocuments

`func (o *RemediateWhatsAppNumberRequest) GetDocuments() []SubmitWhatsAppNumberKycRequestDocumentsInner`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *RemediateWhatsAppNumberRequest) GetDocumentsOk() (*[]SubmitWhatsAppNumberKycRequestDocumentsInner, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *RemediateWhatsAppNumberRequest) SetDocuments(v []SubmitWhatsAppNumberKycRequestDocumentsInner)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *RemediateWhatsAppNumberRequest) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetAddress

`func (o *RemediateWhatsAppNumberRequest) GetAddress() map[string]interface{}`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *RemediateWhatsAppNumberRequest) GetAddressOk() (*map[string]interface{}, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *RemediateWhatsAppNumberRequest) SetAddress(v map[string]interface{})`

SetAddress sets Address field to given value.

### HasAddress

`func (o *RemediateWhatsAppNumberRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


