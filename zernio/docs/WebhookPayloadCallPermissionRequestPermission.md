# WebhookPayloadCallPermissionRequestPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Consumer wa_id who replied | [optional] 
**Response** | Pointer to **string** |  | [optional] 
**IsPermanent** | Pointer to **bool** |  | [optional] 
**ExpirationTimestamp** | Pointer to **time.Time** | Present only when temporary | [optional] 
**ResponseSource** | Pointer to **string** | Meta&#39;s response source, typically &#x60;user_action&#x60; | [optional] 

## Methods

### NewWebhookPayloadCallPermissionRequestPermission

`func NewWebhookPayloadCallPermissionRequestPermission() *WebhookPayloadCallPermissionRequestPermission`

NewWebhookPayloadCallPermissionRequestPermission instantiates a new WebhookPayloadCallPermissionRequestPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCallPermissionRequestPermissionWithDefaults

`func NewWebhookPayloadCallPermissionRequestPermissionWithDefaults() *WebhookPayloadCallPermissionRequestPermission`

NewWebhookPayloadCallPermissionRequestPermissionWithDefaults instantiates a new WebhookPayloadCallPermissionRequestPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *WebhookPayloadCallPermissionRequestPermission) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WebhookPayloadCallPermissionRequestPermission) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WebhookPayloadCallPermissionRequestPermission) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WebhookPayloadCallPermissionRequestPermission) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetResponse

`func (o *WebhookPayloadCallPermissionRequestPermission) GetResponse() string`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *WebhookPayloadCallPermissionRequestPermission) GetResponseOk() (*string, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *WebhookPayloadCallPermissionRequestPermission) SetResponse(v string)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *WebhookPayloadCallPermissionRequestPermission) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetIsPermanent

`func (o *WebhookPayloadCallPermissionRequestPermission) GetIsPermanent() bool`

GetIsPermanent returns the IsPermanent field if non-nil, zero value otherwise.

### GetIsPermanentOk

`func (o *WebhookPayloadCallPermissionRequestPermission) GetIsPermanentOk() (*bool, bool)`

GetIsPermanentOk returns a tuple with the IsPermanent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPermanent

`func (o *WebhookPayloadCallPermissionRequestPermission) SetIsPermanent(v bool)`

SetIsPermanent sets IsPermanent field to given value.

### HasIsPermanent

`func (o *WebhookPayloadCallPermissionRequestPermission) HasIsPermanent() bool`

HasIsPermanent returns a boolean if a field has been set.

### GetExpirationTimestamp

`func (o *WebhookPayloadCallPermissionRequestPermission) GetExpirationTimestamp() time.Time`

GetExpirationTimestamp returns the ExpirationTimestamp field if non-nil, zero value otherwise.

### GetExpirationTimestampOk

`func (o *WebhookPayloadCallPermissionRequestPermission) GetExpirationTimestampOk() (*time.Time, bool)`

GetExpirationTimestampOk returns a tuple with the ExpirationTimestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationTimestamp

`func (o *WebhookPayloadCallPermissionRequestPermission) SetExpirationTimestamp(v time.Time)`

SetExpirationTimestamp sets ExpirationTimestamp field to given value.

### HasExpirationTimestamp

`func (o *WebhookPayloadCallPermissionRequestPermission) HasExpirationTimestamp() bool`

HasExpirationTimestamp returns a boolean if a field has been set.

### GetResponseSource

`func (o *WebhookPayloadCallPermissionRequestPermission) GetResponseSource() string`

GetResponseSource returns the ResponseSource field if non-nil, zero value otherwise.

### GetResponseSourceOk

`func (o *WebhookPayloadCallPermissionRequestPermission) GetResponseSourceOk() (*string, bool)`

GetResponseSourceOk returns a tuple with the ResponseSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseSource

`func (o *WebhookPayloadCallPermissionRequestPermission) SetResponseSource(v string)`

SetResponseSource sets ResponseSource field to given value.

### HasResponseSource

`func (o *WebhookPayloadCallPermissionRequestPermission) HasResponseSource() bool`

HasResponseSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


