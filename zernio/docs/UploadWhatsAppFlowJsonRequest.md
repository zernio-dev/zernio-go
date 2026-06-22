# UploadWhatsAppFlowJsonRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**FlowJson** | [**UploadWhatsAppFlowJsonRequestFlowJson**](UploadWhatsAppFlowJsonRequestFlowJson.md) |  | 

## Methods

### NewUploadWhatsAppFlowJsonRequest

`func NewUploadWhatsAppFlowJsonRequest(accountId string, flowJson UploadWhatsAppFlowJsonRequestFlowJson, ) *UploadWhatsAppFlowJsonRequest`

NewUploadWhatsAppFlowJsonRequest instantiates a new UploadWhatsAppFlowJsonRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWhatsAppFlowJsonRequestWithDefaults

`func NewUploadWhatsAppFlowJsonRequestWithDefaults() *UploadWhatsAppFlowJsonRequest`

NewUploadWhatsAppFlowJsonRequestWithDefaults instantiates a new UploadWhatsAppFlowJsonRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UploadWhatsAppFlowJsonRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UploadWhatsAppFlowJsonRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UploadWhatsAppFlowJsonRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetFlowJson

`func (o *UploadWhatsAppFlowJsonRequest) GetFlowJson() UploadWhatsAppFlowJsonRequestFlowJson`

GetFlowJson returns the FlowJson field if non-nil, zero value otherwise.

### GetFlowJsonOk

`func (o *UploadWhatsAppFlowJsonRequest) GetFlowJsonOk() (*UploadWhatsAppFlowJsonRequestFlowJson, bool)`

GetFlowJsonOk returns a tuple with the FlowJson field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlowJson

`func (o *UploadWhatsAppFlowJsonRequest) SetFlowJson(v UploadWhatsAppFlowJsonRequestFlowJson)`

SetFlowJson sets FlowJson field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


