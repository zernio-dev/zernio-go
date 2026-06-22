# Webhook

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique webhook identifier | [optional] 
**Name** | Pointer to **string** | Webhook name (for identification) | [optional] 
**Url** | Pointer to **string** | Webhook endpoint URL | [optional] 
**Secret** | Pointer to **string** | Secret key for HMAC-SHA256 signature verification. | [optional] 
**Events** | Pointer to **[]string** | Events subscribed to | [optional] 
**IsActive** | Pointer to **bool** | Whether webhook delivery is enabled | [optional] 
**LastFiredAt** | Pointer to **time.Time** | Timestamp of last successful webhook delivery | [optional] 
**FailureCount** | Pointer to **int32** | Consecutive delivery failures (resets on success, webhook disabled at 10) | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | Custom headers included in webhook requests | [optional] 

## Methods

### NewWebhook

`func NewWebhook() *Webhook`

NewWebhook instantiates a new Webhook object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookWithDefaults

`func NewWebhookWithDefaults() *Webhook`

NewWebhookWithDefaults instantiates a new Webhook object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Webhook) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Webhook) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Webhook) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Webhook) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Webhook) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Webhook) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Webhook) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Webhook) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *Webhook) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Webhook) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Webhook) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Webhook) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSecret

`func (o *Webhook) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *Webhook) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *Webhook) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *Webhook) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEvents

`func (o *Webhook) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *Webhook) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *Webhook) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *Webhook) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIsActive

`func (o *Webhook) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *Webhook) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *Webhook) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *Webhook) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetLastFiredAt

`func (o *Webhook) GetLastFiredAt() time.Time`

GetLastFiredAt returns the LastFiredAt field if non-nil, zero value otherwise.

### GetLastFiredAtOk

`func (o *Webhook) GetLastFiredAtOk() (*time.Time, bool)`

GetLastFiredAtOk returns a tuple with the LastFiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredAt

`func (o *Webhook) SetLastFiredAt(v time.Time)`

SetLastFiredAt sets LastFiredAt field to given value.

### HasLastFiredAt

`func (o *Webhook) HasLastFiredAt() bool`

HasLastFiredAt returns a boolean if a field has been set.

### GetFailureCount

`func (o *Webhook) GetFailureCount() int32`

GetFailureCount returns the FailureCount field if non-nil, zero value otherwise.

### GetFailureCountOk

`func (o *Webhook) GetFailureCountOk() (*int32, bool)`

GetFailureCountOk returns a tuple with the FailureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureCount

`func (o *Webhook) SetFailureCount(v int32)`

SetFailureCount sets FailureCount field to given value.

### HasFailureCount

`func (o *Webhook) HasFailureCount() bool`

HasFailureCount returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *Webhook) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *Webhook) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *Webhook) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *Webhook) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


