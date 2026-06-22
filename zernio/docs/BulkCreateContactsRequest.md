# BulkCreateContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** |  | 
**AccountId** | **string** |  | 
**Platform** | **string** |  | 
**Contacts** | [**[]BulkCreateContactsRequestContactsInner**](BulkCreateContactsRequestContactsInner.md) |  | 

## Methods

### NewBulkCreateContactsRequest

`func NewBulkCreateContactsRequest(profileId string, accountId string, platform string, contacts []BulkCreateContactsRequestContactsInner, ) *BulkCreateContactsRequest`

NewBulkCreateContactsRequest instantiates a new BulkCreateContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkCreateContactsRequestWithDefaults

`func NewBulkCreateContactsRequestWithDefaults() *BulkCreateContactsRequest`

NewBulkCreateContactsRequestWithDefaults instantiates a new BulkCreateContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *BulkCreateContactsRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *BulkCreateContactsRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *BulkCreateContactsRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetAccountId

`func (o *BulkCreateContactsRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *BulkCreateContactsRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *BulkCreateContactsRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetPlatform

`func (o *BulkCreateContactsRequest) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *BulkCreateContactsRequest) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *BulkCreateContactsRequest) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetContacts

`func (o *BulkCreateContactsRequest) GetContacts() []BulkCreateContactsRequestContactsInner`

GetContacts returns the Contacts field if non-nil, zero value otherwise.

### GetContactsOk

`func (o *BulkCreateContactsRequest) GetContactsOk() (*[]BulkCreateContactsRequestContactsInner, bool)`

GetContactsOk returns a tuple with the Contacts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContacts

`func (o *BulkCreateContactsRequest) SetContacts(v []BulkCreateContactsRequestContactsInner)`

SetContacts sets Contacts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


