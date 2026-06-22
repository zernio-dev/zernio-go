# WebhookPayloadCommentCommentAd

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Meta ad ID (Instagram only). | [optional] 
**Title** | Pointer to **string** | Ad creative title (Instagram only). | [optional] 
**PromotionStatus** | Pointer to **string** | Facebook promotion status returned by Graph API. Common values: \&quot;active\&quot; (organic post currently boosted), \&quot;ineligible\&quot; (dark post or ad creative, not promotable because it already is an ad).  | [optional] 

## Methods

### NewWebhookPayloadCommentCommentAd

`func NewWebhookPayloadCommentCommentAd() *WebhookPayloadCommentCommentAd`

NewWebhookPayloadCommentCommentAd instantiates a new WebhookPayloadCommentCommentAd object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadCommentCommentAdWithDefaults

`func NewWebhookPayloadCommentCommentAdWithDefaults() *WebhookPayloadCommentCommentAd`

NewWebhookPayloadCommentCommentAdWithDefaults instantiates a new WebhookPayloadCommentCommentAd object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadCommentCommentAd) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadCommentCommentAd) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadCommentCommentAd) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPayloadCommentCommentAd) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *WebhookPayloadCommentCommentAd) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebhookPayloadCommentCommentAd) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebhookPayloadCommentCommentAd) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebhookPayloadCommentCommentAd) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetPromotionStatus

`func (o *WebhookPayloadCommentCommentAd) GetPromotionStatus() string`

GetPromotionStatus returns the PromotionStatus field if non-nil, zero value otherwise.

### GetPromotionStatusOk

`func (o *WebhookPayloadCommentCommentAd) GetPromotionStatusOk() (*string, bool)`

GetPromotionStatusOk returns a tuple with the PromotionStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromotionStatus

`func (o *WebhookPayloadCommentCommentAd) SetPromotionStatus(v string)`

SetPromotionStatus sets PromotionStatus field to given value.

### HasPromotionStatus

`func (o *WebhookPayloadCommentCommentAd) HasPromotionStatus() bool`

HasPromotionStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


