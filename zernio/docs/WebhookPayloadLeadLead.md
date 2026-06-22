# WebhookPayloadLeadLead

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Zernio lead ID (AdLead document ID) | 
**LeadgenId** | **string** | Meta lead ID (the platform&#39;s leadgen_id) | 
**FormId** | **string** | Lead Gen form ID the lead was submitted against | 
**FormName** | Pointer to **NullableString** | Human-readable form name (best-effort; may be null) | [optional] 
**AdId** | Pointer to **NullableString** | Meta ad ID that drove the lead (null for organic/test leads) | [optional] 
**AdsetId** | Pointer to **NullableString** |  | [optional] 
**CampaignId** | Pointer to **NullableString** |  | [optional] 
**Fields** | **map[string]string** | Flattened question key -&gt; answer map. For multiple-choice questions the value is the option key (e.g. \&quot;k1\&quot;), not the display label.  | 
**IsOrganic** | **bool** | True when the lead came from an organic post rather than a paid ad | 
**CreatedAt** | **time.Time** | Meta&#39;s lead creation time (ISO 8601) | 

## Methods

### NewWebhookPayloadLeadLead

`func NewWebhookPayloadLeadLead(id string, leadgenId string, formId string, fields map[string]string, isOrganic bool, createdAt time.Time, ) *WebhookPayloadLeadLead`

NewWebhookPayloadLeadLead instantiates a new WebhookPayloadLeadLead object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPayloadLeadLeadWithDefaults

`func NewWebhookPayloadLeadLeadWithDefaults() *WebhookPayloadLeadLead`

NewWebhookPayloadLeadLeadWithDefaults instantiates a new WebhookPayloadLeadLead object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPayloadLeadLead) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPayloadLeadLead) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPayloadLeadLead) SetId(v string)`

SetId sets Id field to given value.


### GetLeadgenId

`func (o *WebhookPayloadLeadLead) GetLeadgenId() string`

GetLeadgenId returns the LeadgenId field if non-nil, zero value otherwise.

### GetLeadgenIdOk

`func (o *WebhookPayloadLeadLead) GetLeadgenIdOk() (*string, bool)`

GetLeadgenIdOk returns a tuple with the LeadgenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadgenId

`func (o *WebhookPayloadLeadLead) SetLeadgenId(v string)`

SetLeadgenId sets LeadgenId field to given value.


### GetFormId

`func (o *WebhookPayloadLeadLead) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *WebhookPayloadLeadLead) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *WebhookPayloadLeadLead) SetFormId(v string)`

SetFormId sets FormId field to given value.


### GetFormName

`func (o *WebhookPayloadLeadLead) GetFormName() string`

GetFormName returns the FormName field if non-nil, zero value otherwise.

### GetFormNameOk

`func (o *WebhookPayloadLeadLead) GetFormNameOk() (*string, bool)`

GetFormNameOk returns a tuple with the FormName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormName

`func (o *WebhookPayloadLeadLead) SetFormName(v string)`

SetFormName sets FormName field to given value.

### HasFormName

`func (o *WebhookPayloadLeadLead) HasFormName() bool`

HasFormName returns a boolean if a field has been set.

### SetFormNameNil

`func (o *WebhookPayloadLeadLead) SetFormNameNil(b bool)`

 SetFormNameNil sets the value for FormName to be an explicit nil

### UnsetFormName
`func (o *WebhookPayloadLeadLead) UnsetFormName()`

UnsetFormName ensures that no value is present for FormName, not even an explicit nil
### GetAdId

`func (o *WebhookPayloadLeadLead) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *WebhookPayloadLeadLead) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *WebhookPayloadLeadLead) SetAdId(v string)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *WebhookPayloadLeadLead) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### SetAdIdNil

`func (o *WebhookPayloadLeadLead) SetAdIdNil(b bool)`

 SetAdIdNil sets the value for AdId to be an explicit nil

### UnsetAdId
`func (o *WebhookPayloadLeadLead) UnsetAdId()`

UnsetAdId ensures that no value is present for AdId, not even an explicit nil
### GetAdsetId

`func (o *WebhookPayloadLeadLead) GetAdsetId() string`

GetAdsetId returns the AdsetId field if non-nil, zero value otherwise.

### GetAdsetIdOk

`func (o *WebhookPayloadLeadLead) GetAdsetIdOk() (*string, bool)`

GetAdsetIdOk returns a tuple with the AdsetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsetId

`func (o *WebhookPayloadLeadLead) SetAdsetId(v string)`

SetAdsetId sets AdsetId field to given value.

### HasAdsetId

`func (o *WebhookPayloadLeadLead) HasAdsetId() bool`

HasAdsetId returns a boolean if a field has been set.

### SetAdsetIdNil

`func (o *WebhookPayloadLeadLead) SetAdsetIdNil(b bool)`

 SetAdsetIdNil sets the value for AdsetId to be an explicit nil

### UnsetAdsetId
`func (o *WebhookPayloadLeadLead) UnsetAdsetId()`

UnsetAdsetId ensures that no value is present for AdsetId, not even an explicit nil
### GetCampaignId

`func (o *WebhookPayloadLeadLead) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *WebhookPayloadLeadLead) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *WebhookPayloadLeadLead) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *WebhookPayloadLeadLead) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *WebhookPayloadLeadLead) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *WebhookPayloadLeadLead) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetFields

`func (o *WebhookPayloadLeadLead) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *WebhookPayloadLeadLead) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *WebhookPayloadLeadLead) SetFields(v map[string]string)`

SetFields sets Fields field to given value.


### GetIsOrganic

`func (o *WebhookPayloadLeadLead) GetIsOrganic() bool`

GetIsOrganic returns the IsOrganic field if non-nil, zero value otherwise.

### GetIsOrganicOk

`func (o *WebhookPayloadLeadLead) GetIsOrganicOk() (*bool, bool)`

GetIsOrganicOk returns a tuple with the IsOrganic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganic

`func (o *WebhookPayloadLeadLead) SetIsOrganic(v bool)`

SetIsOrganic sets IsOrganic field to given value.


### GetCreatedAt

`func (o *WebhookPayloadLeadLead) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WebhookPayloadLeadLead) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WebhookPayloadLeadLead) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


