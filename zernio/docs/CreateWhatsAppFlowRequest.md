# CreateWhatsAppFlowRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**Name** | **string** | Flow display name | 
**Categories** | **[]string** | Flow categories | 
**CloneFlowId** | Pointer to **string** | Optional: ID of an existing flow to clone the Flow JSON from | [optional] 
**AsVersion** | Pointer to **bool** | When cloning, true keeps the clone in cloneFlowId&#39;s version lineage (auto-numbered next version); false/absent creates an independent flow. Ignored without cloneFlowId. | [optional] 

## Methods

### NewCreateWhatsAppFlowRequest

`func NewCreateWhatsAppFlowRequest(accountId string, name string, categories []string, ) *CreateWhatsAppFlowRequest`

NewCreateWhatsAppFlowRequest instantiates a new CreateWhatsAppFlowRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhatsAppFlowRequestWithDefaults

`func NewCreateWhatsAppFlowRequestWithDefaults() *CreateWhatsAppFlowRequest`

NewCreateWhatsAppFlowRequestWithDefaults instantiates a new CreateWhatsAppFlowRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateWhatsAppFlowRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateWhatsAppFlowRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateWhatsAppFlowRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CreateWhatsAppFlowRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWhatsAppFlowRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWhatsAppFlowRequest) SetName(v string)`

SetName sets Name field to given value.


### GetCategories

`func (o *CreateWhatsAppFlowRequest) GetCategories() []string`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CreateWhatsAppFlowRequest) GetCategoriesOk() (*[]string, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CreateWhatsAppFlowRequest) SetCategories(v []string)`

SetCategories sets Categories field to given value.


### GetCloneFlowId

`func (o *CreateWhatsAppFlowRequest) GetCloneFlowId() string`

GetCloneFlowId returns the CloneFlowId field if non-nil, zero value otherwise.

### GetCloneFlowIdOk

`func (o *CreateWhatsAppFlowRequest) GetCloneFlowIdOk() (*string, bool)`

GetCloneFlowIdOk returns a tuple with the CloneFlowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloneFlowId

`func (o *CreateWhatsAppFlowRequest) SetCloneFlowId(v string)`

SetCloneFlowId sets CloneFlowId field to given value.

### HasCloneFlowId

`func (o *CreateWhatsAppFlowRequest) HasCloneFlowId() bool`

HasCloneFlowId returns a boolean if a field has been set.

### GetAsVersion

`func (o *CreateWhatsAppFlowRequest) GetAsVersion() bool`

GetAsVersion returns the AsVersion field if non-nil, zero value otherwise.

### GetAsVersionOk

`func (o *CreateWhatsAppFlowRequest) GetAsVersionOk() (*bool, bool)`

GetAsVersionOk returns a tuple with the AsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsVersion

`func (o *CreateWhatsAppFlowRequest) SetAsVersion(v bool)`

SetAsVersion sets AsVersion field to given value.

### HasAsVersion

`func (o *CreateWhatsAppFlowRequest) HasAsVersion() bool`

HasAsVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


