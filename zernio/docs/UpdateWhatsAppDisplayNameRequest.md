# UpdateWhatsAppDisplayNameRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**DisplayName** | **string** | New display name (must follow WhatsApp naming guidelines) | 

## Methods

### NewUpdateWhatsAppDisplayNameRequest

`func NewUpdateWhatsAppDisplayNameRequest(accountId string, displayName string, ) *UpdateWhatsAppDisplayNameRequest`

NewUpdateWhatsAppDisplayNameRequest instantiates a new UpdateWhatsAppDisplayNameRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWhatsAppDisplayNameRequestWithDefaults

`func NewUpdateWhatsAppDisplayNameRequestWithDefaults() *UpdateWhatsAppDisplayNameRequest`

NewUpdateWhatsAppDisplayNameRequestWithDefaults instantiates a new UpdateWhatsAppDisplayNameRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateWhatsAppDisplayNameRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWhatsAppDisplayNameRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWhatsAppDisplayNameRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDisplayName

`func (o *UpdateWhatsAppDisplayNameRequest) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateWhatsAppDisplayNameRequest) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateWhatsAppDisplayNameRequest) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


