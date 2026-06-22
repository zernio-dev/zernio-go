# CreateWhatsAppGroupChatRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**Subject** | **string** | Group name (max 128 characters) | 
**Description** | Pointer to **string** | Group description (max 2048 characters) | [optional] 
**JoinApprovalMode** | Pointer to **string** | Whether users need approval to join via invite link | [optional] 

## Methods

### NewCreateWhatsAppGroupChatRequest

`func NewCreateWhatsAppGroupChatRequest(accountId string, subject string, ) *CreateWhatsAppGroupChatRequest`

NewCreateWhatsAppGroupChatRequest instantiates a new CreateWhatsAppGroupChatRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWhatsAppGroupChatRequestWithDefaults

`func NewCreateWhatsAppGroupChatRequestWithDefaults() *CreateWhatsAppGroupChatRequest`

NewCreateWhatsAppGroupChatRequestWithDefaults instantiates a new CreateWhatsAppGroupChatRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateWhatsAppGroupChatRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateWhatsAppGroupChatRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateWhatsAppGroupChatRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetSubject

`func (o *CreateWhatsAppGroupChatRequest) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *CreateWhatsAppGroupChatRequest) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *CreateWhatsAppGroupChatRequest) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetDescription

`func (o *CreateWhatsAppGroupChatRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateWhatsAppGroupChatRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateWhatsAppGroupChatRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateWhatsAppGroupChatRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetJoinApprovalMode

`func (o *CreateWhatsAppGroupChatRequest) GetJoinApprovalMode() string`

GetJoinApprovalMode returns the JoinApprovalMode field if non-nil, zero value otherwise.

### GetJoinApprovalModeOk

`func (o *CreateWhatsAppGroupChatRequest) GetJoinApprovalModeOk() (*string, bool)`

GetJoinApprovalModeOk returns a tuple with the JoinApprovalMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinApprovalMode

`func (o *CreateWhatsAppGroupChatRequest) SetJoinApprovalMode(v string)`

SetJoinApprovalMode sets JoinApprovalMode field to given value.

### HasJoinApprovalMode

`func (o *CreateWhatsAppGroupChatRequest) HasJoinApprovalMode() bool`

HasJoinApprovalMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


