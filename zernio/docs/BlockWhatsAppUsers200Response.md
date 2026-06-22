# BlockWhatsAppUsers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | Pointer to [**[]BlockWhatsAppUsers200ResponseBlockedInner**](BlockWhatsAppUsers200ResponseBlockedInner.md) | Users successfully blocked. | [optional] 
**Failed** | Pointer to [**[]BlockWhatsAppUsers200ResponseFailedInner**](BlockWhatsAppUsers200ResponseFailedInner.md) | Users that could not be blocked, with reasons. | [optional] 

## Methods

### NewBlockWhatsAppUsers200Response

`func NewBlockWhatsAppUsers200Response() *BlockWhatsAppUsers200Response`

NewBlockWhatsAppUsers200Response instantiates a new BlockWhatsAppUsers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockWhatsAppUsers200ResponseWithDefaults

`func NewBlockWhatsAppUsers200ResponseWithDefaults() *BlockWhatsAppUsers200Response`

NewBlockWhatsAppUsers200ResponseWithDefaults instantiates a new BlockWhatsAppUsers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *BlockWhatsAppUsers200Response) GetBlocked() []BlockWhatsAppUsers200ResponseBlockedInner`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *BlockWhatsAppUsers200Response) GetBlockedOk() (*[]BlockWhatsAppUsers200ResponseBlockedInner, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *BlockWhatsAppUsers200Response) SetBlocked(v []BlockWhatsAppUsers200ResponseBlockedInner)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *BlockWhatsAppUsers200Response) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetFailed

`func (o *BlockWhatsAppUsers200Response) GetFailed() []BlockWhatsAppUsers200ResponseFailedInner`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BlockWhatsAppUsers200Response) GetFailedOk() (*[]BlockWhatsAppUsers200ResponseFailedInner, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BlockWhatsAppUsers200Response) SetFailed(v []BlockWhatsAppUsers200ResponseFailedInner)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *BlockWhatsAppUsers200Response) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


