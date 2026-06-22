# InboxWebhookAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Social account ID | 
**Platform** | **string** |  | 
**Username** | **string** |  | 
**DisplayName** | Pointer to **string** |  | [optional] 

## Methods

### NewInboxWebhookAccount

`func NewInboxWebhookAccount(id string, platform string, username string, ) *InboxWebhookAccount`

NewInboxWebhookAccount instantiates a new InboxWebhookAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInboxWebhookAccountWithDefaults

`func NewInboxWebhookAccountWithDefaults() *InboxWebhookAccount`

NewInboxWebhookAccountWithDefaults instantiates a new InboxWebhookAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InboxWebhookAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InboxWebhookAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InboxWebhookAccount) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *InboxWebhookAccount) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *InboxWebhookAccount) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *InboxWebhookAccount) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUsername

`func (o *InboxWebhookAccount) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *InboxWebhookAccount) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *InboxWebhookAccount) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetDisplayName

`func (o *InboxWebhookAccount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *InboxWebhookAccount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *InboxWebhookAccount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *InboxWebhookAccount) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


