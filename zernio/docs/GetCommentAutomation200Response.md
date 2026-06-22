# GetCommentAutomation200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Automation** | Pointer to [**GetCommentAutomation200ResponseAutomation**](GetCommentAutomation200ResponseAutomation.md) |  | [optional] 
**Logs** | Pointer to [**[]GetCommentAutomation200ResponseLogsInner**](GetCommentAutomation200ResponseLogsInner.md) |  | [optional] 

## Methods

### NewGetCommentAutomation200Response

`func NewGetCommentAutomation200Response() *GetCommentAutomation200Response`

NewGetCommentAutomation200Response instantiates a new GetCommentAutomation200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCommentAutomation200ResponseWithDefaults

`func NewGetCommentAutomation200ResponseWithDefaults() *GetCommentAutomation200Response`

NewGetCommentAutomation200ResponseWithDefaults instantiates a new GetCommentAutomation200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetCommentAutomation200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetCommentAutomation200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetCommentAutomation200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetCommentAutomation200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAutomation

`func (o *GetCommentAutomation200Response) GetAutomation() GetCommentAutomation200ResponseAutomation`

GetAutomation returns the Automation field if non-nil, zero value otherwise.

### GetAutomationOk

`func (o *GetCommentAutomation200Response) GetAutomationOk() (*GetCommentAutomation200ResponseAutomation, bool)`

GetAutomationOk returns a tuple with the Automation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomation

`func (o *GetCommentAutomation200Response) SetAutomation(v GetCommentAutomation200ResponseAutomation)`

SetAutomation sets Automation field to given value.

### HasAutomation

`func (o *GetCommentAutomation200Response) HasAutomation() bool`

HasAutomation returns a boolean if a field has been set.

### GetLogs

`func (o *GetCommentAutomation200Response) GetLogs() []GetCommentAutomation200ResponseLogsInner`

GetLogs returns the Logs field if non-nil, zero value otherwise.

### GetLogsOk

`func (o *GetCommentAutomation200Response) GetLogsOk() (*[]GetCommentAutomation200ResponseLogsInner, bool)`

GetLogsOk returns a tuple with the Logs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogs

`func (o *GetCommentAutomation200Response) SetLogs(v []GetCommentAutomation200ResponseLogsInner)`

SetLogs sets Logs field to given value.

### HasLogs

`func (o *GetCommentAutomation200Response) HasLogs() bool`

HasLogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


