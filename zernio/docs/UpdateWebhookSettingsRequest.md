# UpdateWebhookSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Webhook ID to update (required) | 
**Name** | Pointer to **string** | Webhook name (1-50 characters). Must be non-empty if provided. | [optional] 
**Url** | Pointer to **string** | Webhook endpoint URL (must be a valid URL, whitespace trimmed). Must be a valid URL if provided. | [optional] 
**Secret** | Pointer to **string** | Secret key for HMAC-SHA256 signature verification | [optional] 
**Events** | Pointer to **[]string** | Events to subscribe to. Must contain at least one event if provided. | [optional] 
**IsActive** | Pointer to **bool** | Enable or disable webhook delivery | [optional] 
**CustomHeaders** | Pointer to **map[string]string** | Custom headers to include in webhook requests | [optional] 

## Methods

### NewUpdateWebhookSettingsRequest

`func NewUpdateWebhookSettingsRequest(id string, ) *UpdateWebhookSettingsRequest`

NewUpdateWebhookSettingsRequest instantiates a new UpdateWebhookSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookSettingsRequestWithDefaults

`func NewUpdateWebhookSettingsRequestWithDefaults() *UpdateWebhookSettingsRequest`

NewUpdateWebhookSettingsRequestWithDefaults instantiates a new UpdateWebhookSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateWebhookSettingsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateWebhookSettingsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateWebhookSettingsRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateWebhookSettingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateWebhookSettingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateWebhookSettingsRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateWebhookSettingsRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUrl

`func (o *UpdateWebhookSettingsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateWebhookSettingsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateWebhookSettingsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateWebhookSettingsRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetSecret

`func (o *UpdateWebhookSettingsRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *UpdateWebhookSettingsRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *UpdateWebhookSettingsRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *UpdateWebhookSettingsRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEvents

`func (o *UpdateWebhookSettingsRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *UpdateWebhookSettingsRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *UpdateWebhookSettingsRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.

### HasEvents

`func (o *UpdateWebhookSettingsRequest) HasEvents() bool`

HasEvents returns a boolean if a field has been set.

### GetIsActive

`func (o *UpdateWebhookSettingsRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *UpdateWebhookSettingsRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *UpdateWebhookSettingsRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *UpdateWebhookSettingsRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *UpdateWebhookSettingsRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *UpdateWebhookSettingsRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *UpdateWebhookSettingsRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *UpdateWebhookSettingsRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


