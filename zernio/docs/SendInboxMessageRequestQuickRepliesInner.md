# SendInboxMessageRequestQuickRepliesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | Button label (max 20 chars) | 
**Payload** | **string** | Payload sent back on tap | 
**ImageUrl** | Pointer to **string** | Optional icon URL (Meta only) | [optional] 

## Methods

### NewSendInboxMessageRequestQuickRepliesInner

`func NewSendInboxMessageRequestQuickRepliesInner(title string, payload string, ) *SendInboxMessageRequestQuickRepliesInner`

NewSendInboxMessageRequestQuickRepliesInner instantiates a new SendInboxMessageRequestQuickRepliesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestQuickRepliesInnerWithDefaults

`func NewSendInboxMessageRequestQuickRepliesInnerWithDefaults() *SendInboxMessageRequestQuickRepliesInner`

NewSendInboxMessageRequestQuickRepliesInnerWithDefaults instantiates a new SendInboxMessageRequestQuickRepliesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *SendInboxMessageRequestQuickRepliesInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendInboxMessageRequestQuickRepliesInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendInboxMessageRequestQuickRepliesInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPayload

`func (o *SendInboxMessageRequestQuickRepliesInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendInboxMessageRequestQuickRepliesInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendInboxMessageRequestQuickRepliesInner) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetImageUrl

`func (o *SendInboxMessageRequestQuickRepliesInner) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *SendInboxMessageRequestQuickRepliesInner) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *SendInboxMessageRequestQuickRepliesInner) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *SendInboxMessageRequestQuickRepliesInner) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


