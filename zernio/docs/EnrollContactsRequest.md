# EnrollContactsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContactIds** | **[]string** |  | 
**ChannelIds** | Pointer to **[]string** | Optional. Auto-detected if not provided. | [optional] 

## Methods

### NewEnrollContactsRequest

`func NewEnrollContactsRequest(contactIds []string, ) *EnrollContactsRequest`

NewEnrollContactsRequest instantiates a new EnrollContactsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnrollContactsRequestWithDefaults

`func NewEnrollContactsRequestWithDefaults() *EnrollContactsRequest`

NewEnrollContactsRequestWithDefaults instantiates a new EnrollContactsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContactIds

`func (o *EnrollContactsRequest) GetContactIds() []string`

GetContactIds returns the ContactIds field if non-nil, zero value otherwise.

### GetContactIdsOk

`func (o *EnrollContactsRequest) GetContactIdsOk() (*[]string, bool)`

GetContactIdsOk returns a tuple with the ContactIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContactIds

`func (o *EnrollContactsRequest) SetContactIds(v []string)`

SetContactIds sets ContactIds field to given value.


### GetChannelIds

`func (o *EnrollContactsRequest) GetChannelIds() []string`

GetChannelIds returns the ChannelIds field if non-nil, zero value otherwise.

### GetChannelIdsOk

`func (o *EnrollContactsRequest) GetChannelIdsOk() (*[]string, bool)`

GetChannelIdsOk returns a tuple with the ChannelIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelIds

`func (o *EnrollContactsRequest) SetChannelIds(v []string)`

SetChannelIds sets ChannelIds field to given value.

### HasChannelIds

`func (o *EnrollContactsRequest) HasChannelIds() bool`

HasChannelIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


