# SendBroadcast200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** | Current broadcast status after processing first batch | [optional] 
**Sent** | Pointer to **int32** | Recipients sent in this batch | [optional] 
**Failed** | Pointer to **int32** | Recipients failed in this batch | [optional] 
**RecipientCount** | Pointer to **int32** | Total recipient count | [optional] 

## Methods

### NewSendBroadcast200Response

`func NewSendBroadcast200Response() *SendBroadcast200Response`

NewSendBroadcast200Response instantiates a new SendBroadcast200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendBroadcast200ResponseWithDefaults

`func NewSendBroadcast200ResponseWithDefaults() *SendBroadcast200Response`

NewSendBroadcast200ResponseWithDefaults instantiates a new SendBroadcast200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SendBroadcast200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SendBroadcast200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SendBroadcast200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SendBroadcast200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetStatus

`func (o *SendBroadcast200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SendBroadcast200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SendBroadcast200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SendBroadcast200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSent

`func (o *SendBroadcast200Response) GetSent() int32`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *SendBroadcast200Response) GetSentOk() (*int32, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *SendBroadcast200Response) SetSent(v int32)`

SetSent sets Sent field to given value.

### HasSent

`func (o *SendBroadcast200Response) HasSent() bool`

HasSent returns a boolean if a field has been set.

### GetFailed

`func (o *SendBroadcast200Response) GetFailed() int32`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *SendBroadcast200Response) GetFailedOk() (*int32, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *SendBroadcast200Response) SetFailed(v int32)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *SendBroadcast200Response) HasFailed() bool`

HasFailed returns a boolean if a field has been set.

### GetRecipientCount

`func (o *SendBroadcast200Response) GetRecipientCount() int32`

GetRecipientCount returns the RecipientCount field if non-nil, zero value otherwise.

### GetRecipientCountOk

`func (o *SendBroadcast200Response) GetRecipientCountOk() (*int32, bool)`

GetRecipientCountOk returns a tuple with the RecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCount

`func (o *SendBroadcast200Response) SetRecipientCount(v int32)`

SetRecipientCount sets RecipientCount field to given value.

### HasRecipientCount

`func (o *SendBroadcast200Response) HasRecipientCount() bool`

HasRecipientCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


