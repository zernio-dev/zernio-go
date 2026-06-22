# SendInboxMessageRequestContactsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | [**SendInboxMessageRequestContactsInnerName**](SendInboxMessageRequestContactsInnerName.md) |  | 
**Phones** | Pointer to [**[]SendInboxMessageRequestContactsInnerPhonesInner**](SendInboxMessageRequestContactsInnerPhonesInner.md) |  | [optional] 
**Emails** | Pointer to [**[]SendInboxMessageRequestContactsInnerEmailsInner**](SendInboxMessageRequestContactsInnerEmailsInner.md) |  | [optional] 

## Methods

### NewSendInboxMessageRequestContactsInner

`func NewSendInboxMessageRequestContactsInner(name SendInboxMessageRequestContactsInnerName, ) *SendInboxMessageRequestContactsInner`

NewSendInboxMessageRequestContactsInner instantiates a new SendInboxMessageRequestContactsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestContactsInnerWithDefaults

`func NewSendInboxMessageRequestContactsInnerWithDefaults() *SendInboxMessageRequestContactsInner`

NewSendInboxMessageRequestContactsInnerWithDefaults instantiates a new SendInboxMessageRequestContactsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SendInboxMessageRequestContactsInner) GetName() SendInboxMessageRequestContactsInnerName`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SendInboxMessageRequestContactsInner) GetNameOk() (*SendInboxMessageRequestContactsInnerName, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SendInboxMessageRequestContactsInner) SetName(v SendInboxMessageRequestContactsInnerName)`

SetName sets Name field to given value.


### GetPhones

`func (o *SendInboxMessageRequestContactsInner) GetPhones() []SendInboxMessageRequestContactsInnerPhonesInner`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *SendInboxMessageRequestContactsInner) GetPhonesOk() (*[]SendInboxMessageRequestContactsInnerPhonesInner, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *SendInboxMessageRequestContactsInner) SetPhones(v []SendInboxMessageRequestContactsInnerPhonesInner)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *SendInboxMessageRequestContactsInner) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetEmails

`func (o *SendInboxMessageRequestContactsInner) GetEmails() []SendInboxMessageRequestContactsInnerEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *SendInboxMessageRequestContactsInner) GetEmailsOk() (*[]SendInboxMessageRequestContactsInnerEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *SendInboxMessageRequestContactsInner) SetEmails(v []SendInboxMessageRequestContactsInnerEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *SendInboxMessageRequestContactsInner) HasEmails() bool`

HasEmails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


