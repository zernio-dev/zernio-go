# UpdateWebhookSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Webhook** | Pointer to [**Webhook**](Webhook.md) |  | [optional] 

## Methods

### NewUpdateWebhookSettings200Response

`func NewUpdateWebhookSettings200Response() *UpdateWebhookSettings200Response`

NewUpdateWebhookSettings200Response instantiates a new UpdateWebhookSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWebhookSettings200ResponseWithDefaults

`func NewUpdateWebhookSettings200ResponseWithDefaults() *UpdateWebhookSettings200Response`

NewUpdateWebhookSettings200ResponseWithDefaults instantiates a new UpdateWebhookSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateWebhookSettings200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateWebhookSettings200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateWebhookSettings200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateWebhookSettings200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetWebhook

`func (o *UpdateWebhookSettings200Response) GetWebhook() Webhook`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *UpdateWebhookSettings200Response) GetWebhookOk() (*Webhook, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *UpdateWebhookSettings200Response) SetWebhook(v Webhook)`

SetWebhook sets Webhook field to given value.

### HasWebhook

`func (o *UpdateWebhookSettings200Response) HasWebhook() bool`

HasWebhook returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


