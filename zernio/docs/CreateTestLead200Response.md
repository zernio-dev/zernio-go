# CreateTestLead200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**TestLead** | Pointer to [**CreateTestLead200ResponseTestLead**](CreateTestLead200ResponseTestLead.md) |  | [optional] 

## Methods

### NewCreateTestLead200Response

`func NewCreateTestLead200Response() *CreateTestLead200Response`

NewCreateTestLead200Response instantiates a new CreateTestLead200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateTestLead200ResponseWithDefaults

`func NewCreateTestLead200ResponseWithDefaults() *CreateTestLead200Response`

NewCreateTestLead200ResponseWithDefaults instantiates a new CreateTestLead200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateTestLead200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateTestLead200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateTestLead200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CreateTestLead200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTestLead

`func (o *CreateTestLead200Response) GetTestLead() CreateTestLead200ResponseTestLead`

GetTestLead returns the TestLead field if non-nil, zero value otherwise.

### GetTestLeadOk

`func (o *CreateTestLead200Response) GetTestLeadOk() (*CreateTestLead200ResponseTestLead, bool)`

GetTestLeadOk returns a tuple with the TestLead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestLead

`func (o *CreateTestLead200Response) SetTestLead(v CreateTestLead200ResponseTestLead)`

SetTestLead sets TestLead field to given value.

### HasTestLead

`func (o *CreateTestLead200Response) HasTestLead() bool`

HasTestLead returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


