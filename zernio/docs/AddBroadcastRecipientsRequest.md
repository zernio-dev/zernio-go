# AddBroadcastRecipientsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactIds** | Pointer to **[]string** | Specific contact IDs to add | [optional] 
**Phones** | Pointer to **[]string** | Raw phone numbers (auto-creates contacts). Useful for WhatsApp/Telegram manual entry | [optional] 
**UseSegment** | Pointer to **bool** | Auto-populate from broadcast segment filters | [optional] 

## Methods

### NewAddBroadcastRecipientsRequest

`func NewAddBroadcastRecipientsRequest() *AddBroadcastRecipientsRequest`

NewAddBroadcastRecipientsRequest instantiates a new AddBroadcastRecipientsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBroadcastRecipientsRequestWithDefaults

`func NewAddBroadcastRecipientsRequestWithDefaults() *AddBroadcastRecipientsRequest`

NewAddBroadcastRecipientsRequestWithDefaults instantiates a new AddBroadcastRecipientsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactIds

`func (o *AddBroadcastRecipientsRequest) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *AddBroadcastRecipientsRequest) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *AddBroadcastRecipientsRequest) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.

### HasContactIds

`func (o *AddBroadcastRecipientsRequest) HasContactIds() bool`

HasContactIds returns a boolean if a field has been set.

### GetPhones

`func (o *AddBroadcastRecipientsRequest) GetPhones() []string`

GetPhones returns the Phones field if non-nil, zero value otherwise.

### GetPhonesOk

`func (o *AddBroadcastRecipientsRequest) GetPhonesOk() (*[]string, bool)`

GetPhonesOk returns a tuple with the Phones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhones

`func (o *AddBroadcastRecipientsRequest) SetPhones(v []string)`

SetPhones sets Phones field to given value.

### HasPhones

`func (o *AddBroadcastRecipientsRequest) HasPhones() bool`

HasPhones returns a boolean if a field has been set.

### GetUseSegment

`func (o *AddBroadcastRecipientsRequest) GetUseSegment() bool`

GetUseSegment returns the UseSegment field if non-nil, zero value otherwise.

### GetUseSegmentOk

`func (o *AddBroadcastRecipientsRequest) GetUseSegmentOk() (*bool, bool)`

GetUseSegmentOk returns a tuple with the UseSegment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSegment

`func (o *AddBroadcastRecipientsRequest) SetUseSegment(v bool)`

SetUseSegment sets UseSegment field to given value.

### HasUseSegment

`func (o *AddBroadcastRecipientsRequest) HasUseSegment() bool`

HasUseSegment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


