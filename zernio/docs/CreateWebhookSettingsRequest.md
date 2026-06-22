# CreateWebhookSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Webhook name (1-50 characters) | 
**Url** | **string** | Webhook endpoint URL (must be a valid URL, whitespace trimmed) | 
**Secret** | Pointer to **string** | Secret key for HMAC-SHA256 signature verification | [optional] 
**Events** | **[]string** | Events to subscribe to (at least one required) | 
**IsActive** | Pointer to **bool** | Enable or disable webhook delivery. Defaults to &#x60;true&#x60; when omitted. | [optional] [default to true]
**CustomHeaders** | Pointer to **map[string]string** | Custom headers to include in webhook requests | [optional] 

## Methods

### NewCreateWebhookSettingsRequest

`func NewCreateWebhookSettingsRequest(name string, url string, events []string, ) *CreateWebhookSettingsRequest`

NewCreateWebhookSettingsRequest instantiates a new CreateWebhookSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWebhookSettingsRequestWithDefaults

`func NewCreateWebhookSettingsRequestWithDefaults() *CreateWebhookSettingsRequest`

NewCreateWebhookSettingsRequestWithDefaults instantiates a new CreateWebhookSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateWebhookSettingsRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateWebhookSettingsRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateWebhookSettingsRequest) SetName(v string)`

SetName sets Name field to given value.


### GetUrl

`func (o *CreateWebhookSettingsRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateWebhookSettingsRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateWebhookSettingsRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSecret

`func (o *CreateWebhookSettingsRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *CreateWebhookSettingsRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *CreateWebhookSettingsRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *CreateWebhookSettingsRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetEvents

`func (o *CreateWebhookSettingsRequest) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *CreateWebhookSettingsRequest) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *CreateWebhookSettingsRequest) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetIsActive

`func (o *CreateWebhookSettingsRequest) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateWebhookSettingsRequest) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateWebhookSettingsRequest) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.

### HasIsActive

`func (o *CreateWebhookSettingsRequest) HasIsActive() bool`

HasIsActive returns a boolean if a field has been set.

### GetCustomHeaders

`func (o *CreateWebhookSettingsRequest) GetCustomHeaders() map[string]string`

GetCustomHeaders returns the CustomHeaders field if non-nil, zero value otherwise.

### GetCustomHeadersOk

`func (o *CreateWebhookSettingsRequest) GetCustomHeadersOk() (*map[string]string, bool)`

GetCustomHeadersOk returns a tuple with the CustomHeaders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomHeaders

`func (o *CreateWebhookSettingsRequest) SetCustomHeaders(v map[string]string)`

SetCustomHeaders sets CustomHeaders field to given value.

### HasCustomHeaders

`func (o *CreateWebhookSettingsRequest) HasCustomHeaders() bool`

HasCustomHeaders returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


