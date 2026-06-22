# WebhookPayloadCallPermissionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Event** | **string** |  | 
**Permission** | [**WebhookPayloadCallPermissionRequestPermission**](WebhookPayloadCallPermissionRequestPermission.md) |  | 
**Account** | [**InboxWebhookAccount**](InboxWebhookAccount.md) |  | 
**Timestamp** | **time.Time** |  | 

## Methods

### NewWebhookPayloadCallPermissionRequest

`func NewWebhookPayloadCallPermissionRequest(id string, event string, permission WebhookPayloadCallPermissionRequestPermission, account InboxWebhookAccount, timestamp time.Time, ) *WebhookPayloadCallPermissionRequest`

NewWebhookPayloadCallPermissionRequest instantiates a new WebhookPayloadCallPermissionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallPermissionRequestWithDefaults

`func NewWebhookPayloadCallPermissionRequestWithDefaults() *WebhookPayloadCallPermissionRequest`

NewWebhookPayloadCallPermissionRequestWithDefaults instantiates a new WebhookPayloadCallPermissionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCallPermissionRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCallPermissionRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCallPermissionRequest) SetId(v string)`

SetId sets Id field to given value.


### GetEvent

`func (o *WebhookPayloadCallPermissionRequest) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *WebhookPayloadCallPermissionRequest) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *WebhookPayloadCallPermissionRequest) SetEvent(v string)`

SetEvent sets Event field to given value.


### GetPermission

`func (o *WebhookPayloadCallPermissionRequest) GetPermission() WebhookPayloadCallPermissionRequestPermission`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *WebhookPayloadCallPermissionRequest) GetPermissionOk() (*WebhookPayloadCallPermissionRequestPermission, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *WebhookPayloadCallPermissionRequest) SetPermission(v WebhookPayloadCallPermissionRequestPermission)`

SetPermission sets Permission field to given value.


### GetAccount

`func (o *WebhookPayloadCallPermissionRequest) GetAccount() InboxWebhookAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *WebhookPayloadCallPermissionRequest) GetAccountOk() (*InboxWebhookAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *WebhookPayloadCallPermissionRequest) SetAccount(v InboxWebhookAccount)`

SetAccount sets Account field to given value.


### GetTimestamp

`func (o *WebhookPayloadCallPermissionRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *WebhookPayloadCallPermissionRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *WebhookPayloadCallPermissionRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


