# ListCommentAutomations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Automations** | Pointer to [**[]ListCommentAutomations200ResponseAutomationsInner**](ListCommentAutomations200ResponseAutomationsInner.md) |  | [optional] 

## Methods

### NewListCommentAutomations200Response

`func NewListCommentAutomations200Response() *ListCommentAutomations200Response`

NewListCommentAutomations200Response instantiates a new ListCommentAutomations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCommentAutomations200ResponseWithDefaults

`func NewListCommentAutomations200ResponseWithDefaults() *ListCommentAutomations200Response`

NewListCommentAutomations200ResponseWithDefaults instantiates a new ListCommentAutomations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListCommentAutomations200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListCommentAutomations200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListCommentAutomations200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListCommentAutomations200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAutomations

`func (o *ListCommentAutomations200Response) GetAutomations() []ListCommentAutomations200ResponseAutomationsInner`

GetAutomations returns the Automations field if non-nil, zero value otherwise.

### GetAutomationsOk

`func (o *ListCommentAutomations200Response) GetAutomationsOk() (*[]ListCommentAutomations200ResponseAutomationsInner, bool)`

GetAutomationsOk returns a tuple with the Automations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomations

`func (o *ListCommentAutomations200Response) SetAutomations(v []ListCommentAutomations200ResponseAutomationsInner)`

SetAutomations sets Automations field to given value.

### HasAutomations

`func (o *ListCommentAutomations200Response) HasAutomations() bool`

HasAutomations returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


