# SendPrivateReplyToCommentRequestButtonsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Title** | **string** | Label shown on the button. Facebook only. | 
**Url** | **string** | URL opened when the button is tapped. | 
**Payload** | **string** | Opaque value returned in the inbound webhook when the user taps the button. | 
**Phone** | **string** | E.164 phone number dialed when tapped. Facebook only. | 

## Methods

### NewSendPrivateReplyToCommentRequestButtonsInner

`func NewSendPrivateReplyToCommentRequestButtonsInner(type_ string, title string, url string, payload string, phone string, ) *SendPrivateReplyToCommentRequestButtonsInner`

NewSendPrivateReplyToCommentRequestButtonsInner instantiates a new SendPrivateReplyToCommentRequestButtonsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendPrivateReplyToCommentRequestButtonsInnerWithDefaults

`func NewSendPrivateReplyToCommentRequestButtonsInnerWithDefaults() *SendPrivateReplyToCommentRequestButtonsInner`

NewSendPrivateReplyToCommentRequestButtonsInnerWithDefaults instantiates a new SendPrivateReplyToCommentRequestButtonsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendPrivateReplyToCommentRequestButtonsInner) SetType(v string)`

SetType sets Type field to given value.


### GetTitle

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SendPrivateReplyToCommentRequestButtonsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SendPrivateReplyToCommentRequestButtonsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetPayload

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *SendPrivateReplyToCommentRequestButtonsInner) SetPayload(v string)`

SetPayload sets Payload field to given value.


### GetPhone

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SendPrivateReplyToCommentRequestButtonsInner) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SendPrivateReplyToCommentRequestButtonsInner) SetPhone(v string)`

SetPhone sets Phone field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


