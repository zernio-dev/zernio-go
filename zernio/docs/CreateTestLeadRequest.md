# CreateTestLeadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**FieldData** | [**[]CreateTestLeadRequestFieldDataInner**](CreateTestLeadRequestFieldDataInner.md) |  | 

## Methods

### NewCreateTestLeadRequest

`func NewCreateTestLeadRequest(accountId string, fieldData []CreateTestLeadRequestFieldDataInner, ) *CreateTestLeadRequest`

NewCreateTestLeadRequest instantiates a new CreateTestLeadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestLeadRequestWithDefaults

`func NewCreateTestLeadRequestWithDefaults() *CreateTestLeadRequest`

NewCreateTestLeadRequestWithDefaults instantiates a new CreateTestLeadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateTestLeadRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateTestLeadRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateTestLeadRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetFieldData

`func (o *CreateTestLeadRequest) GetFieldData() []CreateTestLeadRequestFieldDataInner`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *CreateTestLeadRequest) GetFieldDataOk() (*[]CreateTestLeadRequestFieldDataInner, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *CreateTestLeadRequest) SetFieldData(v []CreateTestLeadRequestFieldDataInner)`

SetFieldData sets FieldData field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


