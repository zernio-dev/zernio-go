# ListWhatsAppFlows200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Flows** | Pointer to [**[]ListWhatsAppFlows200ResponseFlowsInner**](ListWhatsAppFlows200ResponseFlowsInner.md) |  | [optional] 

## Methods

### NewListWhatsAppFlows200Response

`func NewListWhatsAppFlows200Response() *ListWhatsAppFlows200Response`

NewListWhatsAppFlows200Response instantiates a new ListWhatsAppFlows200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppFlows200ResponseWithDefaults

`func NewListWhatsAppFlows200ResponseWithDefaults() *ListWhatsAppFlows200Response`

NewListWhatsAppFlows200ResponseWithDefaults instantiates a new ListWhatsAppFlows200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListWhatsAppFlows200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListWhatsAppFlows200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListWhatsAppFlows200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListWhatsAppFlows200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFlows

`func (o *ListWhatsAppFlows200Response) GetFlows() []ListWhatsAppFlows200ResponseFlowsInner`

GetFlows returns the Flows field if non-nil, zero value otherwise.

### GetFlowsOk

`func (o *ListWhatsAppFlows200Response) GetFlowsOk() (*[]ListWhatsAppFlows200ResponseFlowsInner, bool)`

GetFlowsOk returns a tuple with the Flows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlows

`func (o *ListWhatsAppFlows200Response) SetFlows(v []ListWhatsAppFlows200ResponseFlowsInner)`

SetFlows sets Flows field to given value.

### HasFlows

`func (o *ListWhatsAppFlows200Response) HasFlows() bool`

HasFlows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


