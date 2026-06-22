# SendInboxMessageRequestButtonsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Button type. phone is Facebook only. | 
**Title** | **string** | Button label (max 20 chars) | 
**Url** | Pointer to **string** | URL for url-type buttons | [optional] 
**Payload** | Pointer to **string** | Payload for postback-type buttons | [optional] 
**Phone** | Pointer to **string** | Phone number for phone-type buttons (Facebook only) | [optional] 

## Methods

### NewSendInboxMessageRequestButtonsInner

`func NewSendInboxMessageRequestButtonsInner(type_ string, title string, ) *SendInboxMessageRequestButtonsInner`

NewSendInboxMessageRequestButtonsInner instantiates a new SendInboxMessageRequestButtonsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestButtonsInnerWithDefaults

`func NewSendInboxMessageRequestButtonsInnerWithDefaults() *SendInboxMessageRequestButtonsInner`

NewSendInboxMessageRequestButtonsInnerWithDefaults instantiates a new SendInboxMessageRequestButtonsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendInboxMessageRequestButtonsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendInboxMessageRequestButtonsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendInboxMessageRequestButtonsInner) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *SendInboxMessageRequestButtonsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendInboxMessageRequestButtonsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendInboxMessageRequestButtonsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *SendInboxMessageRequestButtonsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SendInboxMessageRequestButtonsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SendInboxMessageRequestButtonsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *SendInboxMessageRequestButtonsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPayload

`func (o *SendInboxMessageRequestButtonsInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendInboxMessageRequestButtonsInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendInboxMessageRequestButtonsInner) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *SendInboxMessageRequestButtonsInner) HasPayload() bool`

HasPayload returns a boolean if a field has been set.

### GetPhone

`func (o *SendInboxMessageRequestButtonsInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendInboxMessageRequestButtonsInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendInboxMessageRequestButtonsInner) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SendInboxMessageRequestButtonsInner) HasPhone() bool`

HasPhone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


