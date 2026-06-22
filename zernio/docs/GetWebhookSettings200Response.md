# GetWebhookSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Webhooks** | Pointer to [**[]Webhook**](Webhook.md) |  | [optional] 

## Methods

### NewGetWebhookSettings200Response

`func NewGetWebhookSettings200Response() *GetWebhookSettings200Response`

NewGetWebhookSettings200Response instantiates a new GetWebhookSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWebhookSettings200ResponseWithDefaults

`func NewGetWebhookSettings200ResponseWithDefaults() *GetWebhookSettings200Response`

NewGetWebhookSettings200ResponseWithDefaults instantiates a new GetWebhookSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebhooks

`func (o *GetWebhookSettings200Response) GetWebhooks() []Webhook`

GetWebhooks returns the Webhooks field if non-nil, zero value otherwise.

### GetWebhooksOk

`func (o *GetWebhookSettings200Response) GetWebhooksOk() (*[]Webhook, bool)`

GetWebhooksOk returns a tuple with the Webhooks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhooks

`func (o *GetWebhookSettings200Response) SetWebhooks(v []Webhook)`

SetWebhooks sets Webhooks field to given value.

### HasWebhooks

`func (o *GetWebhookSettings200Response) HasWebhooks() bool`

HasWebhooks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


