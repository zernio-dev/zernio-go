# WebhookPayloadAdStatusChangedError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Platform-native error code, forwarded verbatim. For Meta this is &#x60;error_code&#x60; as a string. Use as the stable discriminator — &#x60;summary&#x60; and &#x60;message&#x60; are localized.  | 
**Summary** | Pointer to **string** | Short human-readable summary (Meta &#x60;error_summary&#x60;). Localized to the ad-account owner&#39;s Meta locale — display only, do not match on it.  | [optional] 
**Message** | Pointer to **string** | Full human-readable error message (Meta &#x60;error_message&#x60;). Localized — display only.  | [optional] 

## Methods

### NewWebhookPayloadAdStatusChangedError

`func NewWebhookPayloadAdStatusChangedError(code string, ) *WebhookPayloadAdStatusChangedError`

NewWebhookPayloadAdStatusChangedError instantiates a new WebhookPayloadAdStatusChangedError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAdStatusChangedErrorWithDefaults

`func NewWebhookPayloadAdStatusChangedErrorWithDefaults() *WebhookPayloadAdStatusChangedError`

NewWebhookPayloadAdStatusChangedErrorWithDefaults instantiates a new WebhookPayloadAdStatusChangedError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WebhookPayloadAdStatusChangedError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WebhookPayloadAdStatusChangedError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WebhookPayloadAdStatusChangedError) SetCode(v string)`

SetCode sets Code field to given value.


### GetSummary

`func (o *WebhookPayloadAdStatusChangedError) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *WebhookPayloadAdStatusChangedError) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *WebhookPayloadAdStatusChangedError) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *WebhookPayloadAdStatusChangedError) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetMessage

`func (o *WebhookPayloadAdStatusChangedError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WebhookPayloadAdStatusChangedError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WebhookPayloadAdStatusChangedError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WebhookPayloadAdStatusChangedError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


