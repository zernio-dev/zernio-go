# BulkCreateContacts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Skipped** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewBulkCreateContacts200Response

`func NewBulkCreateContacts200Response() *BulkCreateContacts200Response`

NewBulkCreateContacts200Response instantiates a new BulkCreateContacts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateContacts200ResponseWithDefaults

`func NewBulkCreateContacts200ResponseWithDefaults() *BulkCreateContacts200Response`

NewBulkCreateContacts200ResponseWithDefaults instantiates a new BulkCreateContacts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *BulkCreateContacts200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkCreateContacts200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkCreateContacts200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BulkCreateContacts200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetCreated

`func (o *BulkCreateContacts200Response) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *BulkCreateContacts200Response) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *BulkCreateContacts200Response) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *BulkCreateContacts200Response) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetSkipped

`func (o *BulkCreateContacts200Response) GetSkipped() int32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *BulkCreateContacts200Response) GetSkippedOk() (*int32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *BulkCreateContacts200Response) SetSkipped(v int32)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *BulkCreateContacts200Response) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.

### GetErrors

`func (o *BulkCreateContacts200Response) GetErrors() []map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkCreateContacts200Response) GetErrorsOk() (*[]map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkCreateContacts200Response) SetErrors(v []map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkCreateContacts200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetTotal

`func (o *BulkCreateContacts200Response) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BulkCreateContacts200Response) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BulkCreateContacts200Response) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BulkCreateContacts200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


