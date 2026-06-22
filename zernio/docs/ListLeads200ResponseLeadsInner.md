# ListLeads200ResponseLeadsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Zernio lead id. | [optional] 
**LeadgenId** | Pointer to **string** | Meta lead id. | [optional] 
**FormId** | Pointer to **string** |  | [optional] 
**FormName** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**AdId** | Pointer to **NullableString** |  | [optional] 
**AdsetId** | Pointer to **NullableString** |  | [optional] 
**CampaignId** | Pointer to **NullableString** |  | [optional] 
**IsOrganic** | Pointer to **bool** |  | [optional] 
**CreatedTime** | Pointer to **NullableString** | ISO 8601. | [optional] 
**Fields** | Pointer to **map[string]string** | Question key → answer. | [optional] 
**FieldData** | Pointer to **[]map[string]interface{}** | Raw Meta field_data. | [optional] 

## Methods

### NewListLeads200ResponseLeadsInner

`func NewListLeads200ResponseLeadsInner() *ListLeads200ResponseLeadsInner`

NewListLeads200ResponseLeadsInner instantiates a new ListLeads200ResponseLeadsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLeads200ResponseLeadsInnerWithDefaults

`func NewListLeads200ResponseLeadsInnerWithDefaults() *ListLeads200ResponseLeadsInner`

NewListLeads200ResponseLeadsInnerWithDefaults instantiates a new ListLeads200ResponseLeadsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLeads200ResponseLeadsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLeads200ResponseLeadsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLeads200ResponseLeadsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListLeads200ResponseLeadsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLeadgenId

`func (o *ListLeads200ResponseLeadsInner) GetLeadgenId() string`

GetLeadgenId returns the LeadgenId field if non-nil, zero value otherwise.

### GetLeadgenIdOk

`func (o *ListLeads200ResponseLeadsInner) GetLeadgenIdOk() (*string, bool)`

GetLeadgenIdOk returns a tuple with the LeadgenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeadgenId

`func (o *ListLeads200ResponseLeadsInner) SetLeadgenId(v string)`

SetLeadgenId sets LeadgenId field to given value.

### HasLeadgenId

`func (o *ListLeads200ResponseLeadsInner) HasLeadgenId() bool`

HasLeadgenId returns a boolean if a field has been set.

### GetFormId

`func (o *ListLeads200ResponseLeadsInner) GetFormId() string`

GetFormId returns the FormId field if non-nil, zero value otherwise.

### GetFormIdOk

`func (o *ListLeads200ResponseLeadsInner) GetFormIdOk() (*string, bool)`

GetFormIdOk returns a tuple with the FormId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormId

`func (o *ListLeads200ResponseLeadsInner) SetFormId(v string)`

SetFormId sets FormId field to given value.

### HasFormId

`func (o *ListLeads200ResponseLeadsInner) HasFormId() bool`

HasFormId returns a boolean if a field has been set.

### GetFormName

`func (o *ListLeads200ResponseLeadsInner) GetFormName() string`

GetFormName returns the FormName field if non-nil, zero value otherwise.

### GetFormNameOk

`func (o *ListLeads200ResponseLeadsInner) GetFormNameOk() (*string, bool)`

GetFormNameOk returns a tuple with the FormName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormName

`func (o *ListLeads200ResponseLeadsInner) SetFormName(v string)`

SetFormName sets FormName field to given value.

### HasFormName

`func (o *ListLeads200ResponseLeadsInner) HasFormName() bool`

HasFormName returns a boolean if a field has been set.

### SetFormNameNil

`func (o *ListLeads200ResponseLeadsInner) SetFormNameNil(b bool)`

 SetFormNameNil sets the value for FormName to be an explicit nil

### UnsetFormName
`func (o *ListLeads200ResponseLeadsInner) UnsetFormName()`

UnsetFormName ensures that no value is present for FormName, not even an explicit nil
### GetAccountId

`func (o *ListLeads200ResponseLeadsInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListLeads200ResponseLeadsInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListLeads200ResponseLeadsInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListLeads200ResponseLeadsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAdId

`func (o *ListLeads200ResponseLeadsInner) GetAdId() string`

GetAdId returns the AdId field if non-nil, zero value otherwise.

### GetAdIdOk

`func (o *ListLeads200ResponseLeadsInner) GetAdIdOk() (*string, bool)`

GetAdIdOk returns a tuple with the AdId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdId

`func (o *ListLeads200ResponseLeadsInner) SetAdId(v string)`

SetAdId sets AdId field to given value.

### HasAdId

`func (o *ListLeads200ResponseLeadsInner) HasAdId() bool`

HasAdId returns a boolean if a field has been set.

### SetAdIdNil

`func (o *ListLeads200ResponseLeadsInner) SetAdIdNil(b bool)`

 SetAdIdNil sets the value for AdId to be an explicit nil

### UnsetAdId
`func (o *ListLeads200ResponseLeadsInner) UnsetAdId()`

UnsetAdId ensures that no value is present for AdId, not even an explicit nil
### GetAdsetId

`func (o *ListLeads200ResponseLeadsInner) GetAdsetId() string`

GetAdsetId returns the AdsetId field if non-nil, zero value otherwise.

### GetAdsetIdOk

`func (o *ListLeads200ResponseLeadsInner) GetAdsetIdOk() (*string, bool)`

GetAdsetIdOk returns a tuple with the AdsetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdsetId

`func (o *ListLeads200ResponseLeadsInner) SetAdsetId(v string)`

SetAdsetId sets AdsetId field to given value.

### HasAdsetId

`func (o *ListLeads200ResponseLeadsInner) HasAdsetId() bool`

HasAdsetId returns a boolean if a field has been set.

### SetAdsetIdNil

`func (o *ListLeads200ResponseLeadsInner) SetAdsetIdNil(b bool)`

 SetAdsetIdNil sets the value for AdsetId to be an explicit nil

### UnsetAdsetId
`func (o *ListLeads200ResponseLeadsInner) UnsetAdsetId()`

UnsetAdsetId ensures that no value is present for AdsetId, not even an explicit nil
### GetCampaignId

`func (o *ListLeads200ResponseLeadsInner) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ListLeads200ResponseLeadsInner) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ListLeads200ResponseLeadsInner) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *ListLeads200ResponseLeadsInner) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *ListLeads200ResponseLeadsInner) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *ListLeads200ResponseLeadsInner) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil
### GetIsOrganic

`func (o *ListLeads200ResponseLeadsInner) GetIsOrganic() bool`

GetIsOrganic returns the IsOrganic field if non-nil, zero value otherwise.

### GetIsOrganicOk

`func (o *ListLeads200ResponseLeadsInner) GetIsOrganicOk() (*bool, bool)`

GetIsOrganicOk returns a tuple with the IsOrganic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOrganic

`func (o *ListLeads200ResponseLeadsInner) SetIsOrganic(v bool)`

SetIsOrganic sets IsOrganic field to given value.

### HasIsOrganic

`func (o *ListLeads200ResponseLeadsInner) HasIsOrganic() bool`

HasIsOrganic returns a boolean if a field has been set.

### GetCreatedTime

`func (o *ListLeads200ResponseLeadsInner) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *ListLeads200ResponseLeadsInner) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *ListLeads200ResponseLeadsInner) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *ListLeads200ResponseLeadsInner) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### SetCreatedTimeNil

`func (o *ListLeads200ResponseLeadsInner) SetCreatedTimeNil(b bool)`

 SetCreatedTimeNil sets the value for CreatedTime to be an explicit nil

### UnsetCreatedTime
`func (o *ListLeads200ResponseLeadsInner) UnsetCreatedTime()`

UnsetCreatedTime ensures that no value is present for CreatedTime, not even an explicit nil
### GetFields

`func (o *ListLeads200ResponseLeadsInner) GetFields() map[string]string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *ListLeads200ResponseLeadsInner) GetFieldsOk() (*map[string]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *ListLeads200ResponseLeadsInner) SetFields(v map[string]string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *ListLeads200ResponseLeadsInner) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetFieldData

`func (o *ListLeads200ResponseLeadsInner) GetFieldData() []map[string]interface{}`

GetFieldData returns the FieldData field if non-nil, zero value otherwise.

### GetFieldDataOk

`func (o *ListLeads200ResponseLeadsInner) GetFieldDataOk() (*[]map[string]interface{}, bool)`

GetFieldDataOk returns a tuple with the FieldData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldData

`func (o *ListLeads200ResponseLeadsInner) SetFieldData(v []map[string]interface{})`

SetFieldData sets FieldData field to given value.

### HasFieldData

`func (o *ListLeads200ResponseLeadsInner) HasFieldData() bool`

HasFieldData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


