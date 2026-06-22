# WebhookPayloadAdStatusChangedStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Raw** | **string** | Platform-native status string, forwarded verbatim. For Meta this is &#x60;status_name&#x60; from &#x60;in_process_ad_objects&#x60; (e.g. &#x60;ACTIVE&#x60;, &#x60;PAUSED&#x60;, &#x60;PENDING_REVIEW&#x60;, &#x60;ARCHIVED&#x60;, &#x60;DELETED&#x60;, &#x60;DISAPPROVED&#x60;), or &#x60;WITH_ISSUES&#x60; when sourced from &#x60;with_issues_ad_objects&#x60;. Not constrained by an &#x60;enum&#x60; — Meta may add new values.  | 

## Methods

### NewWebhookPayloadAdStatusChangedStatus

`func NewWebhookPayloadAdStatusChangedStatus(raw string, ) *WebhookPayloadAdStatusChangedStatus`

NewWebhookPayloadAdStatusChangedStatus instantiates a new WebhookPayloadAdStatusChangedStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAdStatusChangedStatusWithDefaults

`func NewWebhookPayloadAdStatusChangedStatusWithDefaults() *WebhookPayloadAdStatusChangedStatus`

NewWebhookPayloadAdStatusChangedStatusWithDefaults instantiates a new WebhookPayloadAdStatusChangedStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRaw

`func (o *WebhookPayloadAdStatusChangedStatus) GetRaw() string`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *WebhookPayloadAdStatusChangedStatus) GetRawOk() (*string, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *WebhookPayloadAdStatusChangedStatus) SetRaw(v string)`

SetRaw sets Raw field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


