# CreateCommentAutomation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Automation** | Pointer to [**CreateCommentAutomation200ResponseAutomation**](CreateCommentAutomation200ResponseAutomation.md) |  | [optional] 

## Methods

### NewCreateCommentAutomation200Response

`func NewCreateCommentAutomation200Response() *CreateCommentAutomation200Response`

NewCreateCommentAutomation200Response instantiates a new CreateCommentAutomation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCommentAutomation200ResponseWithDefaults

`func NewCreateCommentAutomation200ResponseWithDefaults() *CreateCommentAutomation200Response`

NewCreateCommentAutomation200ResponseWithDefaults instantiates a new CreateCommentAutomation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateCommentAutomation200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateCommentAutomation200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateCommentAutomation200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateCommentAutomation200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAutomation

`func (o *CreateCommentAutomation200Response) GetAutomation() CreateCommentAutomation200ResponseAutomation`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *CreateCommentAutomation200Response) GetAutomationOk() (*CreateCommentAutomation200ResponseAutomation, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *CreateCommentAutomation200Response) SetAutomation(v CreateCommentAutomation200ResponseAutomation)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *CreateCommentAutomation200Response) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


