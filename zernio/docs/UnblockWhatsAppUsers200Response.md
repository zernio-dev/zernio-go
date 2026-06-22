# UnblockWhatsAppUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Unblocked** | Pointer to [**[]BlockWhatsAppUsers200ResponseBlockedInner**](BlockWhatsAppUsers200ResponseBlockedInner.md) | Users successfully unblocked. | [optional] 
**Failed** | Pointer to [**[]BlockWhatsAppUsers200ResponseFailedInner**](BlockWhatsAppUsers200ResponseFailedInner.md) | Users that could not be unblocked, with reasons. | [optional] 

## Methods

### NewUnblockWhatsAppUsers200Response

`func NewUnblockWhatsAppUsers200Response() *UnblockWhatsAppUsers200Response`

NewUnblockWhatsAppUsers200Response instantiates a new UnblockWhatsAppUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUnblockWhatsAppUsers200ResponseWithDefaults

`func NewUnblockWhatsAppUsers200ResponseWithDefaults() *UnblockWhatsAppUsers200Response`

NewUnblockWhatsAppUsers200ResponseWithDefaults instantiates a new UnblockWhatsAppUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnblocked

`func (o *UnblockWhatsAppUsers200Response) GetUnblocked() []BlockWhatsAppUsers200ResponseBlockedInner`

GetUnblocked returns the Unblocked field if non-nil, zero value otherwise.

### GetUnblockedOk

`func (o *UnblockWhatsAppUsers200Response) GetUnblockedOk() (*[]BlockWhatsAppUsers200ResponseBlockedInner, bool)`

GetUnblockedOk returns a tuple with the Unblocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnblocked

`func (o *UnblockWhatsAppUsers200Response) SetUnblocked(v []BlockWhatsAppUsers200ResponseBlockedInner)`

SetUnblocked sets Unblocked field to given value.

### HasUnblocked

`func (o *UnblockWhatsAppUsers200Response) HasUnblocked() bool`

HasUnblocked returns a boolean if a field has been set.

### GetFailed

`func (o *UnblockWhatsAppUsers200Response) GetFailed() []BlockWhatsAppUsers200ResponseFailedInner`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *UnblockWhatsAppUsers200Response) GetFailedOk() (*[]BlockWhatsAppUsers200ResponseFailedInner, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *UnblockWhatsAppUsers200Response) SetFailed(v []BlockWhatsAppUsers200ResponseFailedInner)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *UnblockWhatsAppUsers200Response) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


