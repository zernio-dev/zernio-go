# UpdateWhatsAppFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**Name** | Pointer to **string** | New flow name | [optional] 
**Categories** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateWhatsAppFlowRequest

`func NewUpdateWhatsAppFlowRequest(accountId string, ) *UpdateWhatsAppFlowRequest`

NewUpdateWhatsAppFlowRequest instantiates a new UpdateWhatsAppFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWhatsAppFlowRequestWithDefaults

`func NewUpdateWhatsAppFlowRequestWithDefaults() *UpdateWhatsAppFlowRequest`

NewUpdateWhatsAppFlowRequestWithDefaults instantiates a new UpdateWhatsAppFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateWhatsAppFlowRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWhatsAppFlowRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWhatsAppFlowRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *UpdateWhatsAppFlowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWhatsAppFlowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWhatsAppFlowRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWhatsAppFlowRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCategories

`func (o *UpdateWhatsAppFlowRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UpdateWhatsAppFlowRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UpdateWhatsAppFlowRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UpdateWhatsAppFlowRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


