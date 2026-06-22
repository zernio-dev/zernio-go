# SendInboxMessageRequestInteractive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Which interactive layout to render. | 
**Header** | Pointer to [**SendInboxMessageRequestInteractiveHeader**](SendInboxMessageRequestInteractiveHeader.md) |  | [optional] 
**Body** | [**ValidatePostLengthRequest**](ValidatePostLengthRequest.md) |  | 
**Footer** | Pointer to [**SendInboxMessageRequestInteractiveFooter**](SendInboxMessageRequestInteractiveFooter.md) |  | [optional] 
**Action** | Pointer to [**SendInboxMessageRequestInteractiveAction**](SendInboxMessageRequestInteractiveAction.md) |  | [optional] 

## Methods

### NewSendInboxMessageRequestInteractive

`func NewSendInboxMessageRequestInteractive(type_ string, body ValidatePostLengthRequest, ) *SendInboxMessageRequestInteractive`

NewSendInboxMessageRequestInteractive instantiates a new SendInboxMessageRequestInteractive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendInboxMessageRequestInteractiveWithDefaults

`func NewSendInboxMessageRequestInteractiveWithDefaults() *SendInboxMessageRequestInteractive`

NewSendInboxMessageRequestInteractiveWithDefaults instantiates a new SendInboxMessageRequestInteractive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SendInboxMessageRequestInteractive) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SendInboxMessageRequestInteractive) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SendInboxMessageRequestInteractive) SetType(v string)`

SetType sets Type field to given value.


### GetHeader

`func (o *SendInboxMessageRequestInteractive) GetHeader() SendInboxMessageRequestInteractiveHeader`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *SendInboxMessageRequestInteractive) GetHeaderOk() (*SendInboxMessageRequestInteractiveHeader, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *SendInboxMessageRequestInteractive) SetHeader(v SendInboxMessageRequestInteractiveHeader)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *SendInboxMessageRequestInteractive) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetBody

`func (o *SendInboxMessageRequestInteractive) GetBody() ValidatePostLengthRequest`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *SendInboxMessageRequestInteractive) GetBodyOk() (*ValidatePostLengthRequest, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *SendInboxMessageRequestInteractive) SetBody(v ValidatePostLengthRequest)`

SetBody sets Body field to given value.


### GetFooter

`func (o *SendInboxMessageRequestInteractive) GetFooter() SendInboxMessageRequestInteractiveFooter`

GetFooter returns the Footer field if non-nil, zero value otherwise.

### GetFooterOk

`func (o *SendInboxMessageRequestInteractive) GetFooterOk() (*SendInboxMessageRequestInteractiveFooter, bool)`

GetFooterOk returns a tuple with the Footer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooter

`func (o *SendInboxMessageRequestInteractive) SetFooter(v SendInboxMessageRequestInteractiveFooter)`

SetFooter sets Footer field to given value.

### HasFooter

`func (o *SendInboxMessageRequestInteractive) HasFooter() bool`

HasFooter returns a boolean if a field has been set.

### GetAction

`func (o *SendInboxMessageRequestInteractive) GetAction() SendInboxMessageRequestInteractiveAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *SendInboxMessageRequestInteractive) GetActionOk() (*SendInboxMessageRequestInteractiveAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *SendInboxMessageRequestInteractive) SetAction(v SendInboxMessageRequestInteractiveAction)`

SetAction sets Action field to given value.

### HasAction

`func (o *SendInboxMessageRequestInteractive) HasAction() bool`

HasAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


