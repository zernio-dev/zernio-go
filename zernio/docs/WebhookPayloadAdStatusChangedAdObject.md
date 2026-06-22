# WebhookPayloadAdStatusChangedAdObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level** | **string** | Hierarchy level the status applies to. Mirrors Meta&#39;s &#x60;level&#x60;. Creative-level events are not forwarded. | 
**PlatformId** | **string** | Platform-native ID of the campaign / ad set / ad. For Meta this is the bare numeric ID (e.g. &#x60;120244894077860689&#x60;).  | 
**PlatformAdAccountId** | **string** | Platform-native ad-account ID. For Meta this uses the &#x60;act_&lt;id&gt;&#x60; shape.  | 

## Methods

### NewWebhookPayloadAdStatusChangedAdObject

`func NewWebhookPayloadAdStatusChangedAdObject(level string, platformId string, platformAdAccountId string, ) *WebhookPayloadAdStatusChangedAdObject`

NewWebhookPayloadAdStatusChangedAdObject instantiates a new WebhookPayloadAdStatusChangedAdObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadAdStatusChangedAdObjectWithDefaults

`func NewWebhookPayloadAdStatusChangedAdObjectWithDefaults() *WebhookPayloadAdStatusChangedAdObject`

NewWebhookPayloadAdStatusChangedAdObjectWithDefaults instantiates a new WebhookPayloadAdStatusChangedAdObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel

`func (o *WebhookPayloadAdStatusChangedAdObject) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *WebhookPayloadAdStatusChangedAdObject) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *WebhookPayloadAdStatusChangedAdObject) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetPlatformId

`func (o *WebhookPayloadAdStatusChangedAdObject) GetPlatformId() string`

GetPlatformId returns the PlatformId field if non-nil, zero value otherwise.

### GetPlatformIdOk

`func (o *WebhookPayloadAdStatusChangedAdObject) GetPlatformIdOk() (*string, bool)`

GetPlatformIdOk returns a tuple with the PlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformId

`func (o *WebhookPayloadAdStatusChangedAdObject) SetPlatformId(v string)`

SetPlatformId sets PlatformId field to given value.


### GetPlatformAdAccountId

`func (o *WebhookPayloadAdStatusChangedAdObject) GetPlatformAdAccountId() string`

GetPlatformAdAccountId returns the PlatformAdAccountId field if non-nil, zero value otherwise.

### GetPlatformAdAccountIdOk

`func (o *WebhookPayloadAdStatusChangedAdObject) GetPlatformAdAccountIdOk() (*string, bool)`

GetPlatformAdAccountIdOk returns a tuple with the PlatformAdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformAdAccountId

`func (o *WebhookPayloadAdStatusChangedAdObject) SetPlatformAdAccountId(v string)`

SetPlatformAdAccountId sets PlatformAdAccountId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


