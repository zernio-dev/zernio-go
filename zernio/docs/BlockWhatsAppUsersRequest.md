# BlockWhatsAppUsersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**Users** | **[]string** | Phone numbers (E.164, e.g. \&quot;+16505551234\&quot;) or WhatsApp user IDs to block. | 

## Methods

### NewBlockWhatsAppUsersRequest

`func NewBlockWhatsAppUsersRequest(accountId string, users []string, ) *BlockWhatsAppUsersRequest`

NewBlockWhatsAppUsersRequest instantiates a new BlockWhatsAppUsersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWhatsAppUsersRequestWithDefaults

`func NewBlockWhatsAppUsersRequestWithDefaults() *BlockWhatsAppUsersRequest`

NewBlockWhatsAppUsersRequestWithDefaults instantiates a new BlockWhatsAppUsersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *BlockWhatsAppUsersRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BlockWhatsAppUsersRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BlockWhatsAppUsersRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetUsers

`func (o *BlockWhatsAppUsersRequest) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BlockWhatsAppUsersRequest) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BlockWhatsAppUsersRequest) SetUsers(v []string)`

SetUsers sets Users field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


