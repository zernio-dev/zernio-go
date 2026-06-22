# CreateAdAudienceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | Social account ID on the target ad platform. | 
**AdAccountId** | **string** | Platform ad account ID. Must start with act_ for Meta; bare platform id for others (Google customer id, X/TikTok/LinkedIn/Pinterest account id). | 
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** |  | 
**PixelId** | Pointer to **string** | Required for website audiences | [optional] 
**RetentionDays** | Pointer to **int32** | Required for website audiences | [optional] 
**SourceAudienceId** | Pointer to **string** | Required for lookalike audiences | [optional] 
**Country** | Pointer to **string** | 2-letter code, required for lookalike audiences | [optional] 
**Ratio** | Pointer to **float32** | Required for lookalike audiences | [optional] 
**Rule** | Pointer to **map[string]interface{}** | Pixel event rule for website audiences (optional) | [optional] 
**CustomerFileSource** | Pointer to **string** | Data source declaration for GDPR compliance (customer_list only) | [optional] 
**Spec** | [**TargetingSpec**](TargetingSpec.md) | The targeting spec to store. | 

## Methods

### NewCreateAdAudienceRequest

`func NewCreateAdAudienceRequest(accountId string, adAccountId string, name string, type_ string, spec TargetingSpec, ) *CreateAdAudienceRequest`

NewCreateAdAudienceRequest instantiates a new CreateAdAudienceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdAudienceRequestWithDefaults

`func NewCreateAdAudienceRequestWithDefaults() *CreateAdAudienceRequest`

NewCreateAdAudienceRequestWithDefaults instantiates a new CreateAdAudienceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateAdAudienceRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateAdAudienceRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateAdAudienceRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAdAccountId

`func (o *CreateAdAudienceRequest) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *CreateAdAudienceRequest) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *CreateAdAudienceRequest) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *CreateAdAudienceRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAdAudienceRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAdAudienceRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateAdAudienceRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAdAudienceRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAdAudienceRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAdAudienceRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *CreateAdAudienceRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateAdAudienceRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateAdAudienceRequest) SetType(v string)`

SetType sets Type field to given value.


### GetPixelId

`func (o *CreateAdAudienceRequest) GetPixelId() string`

GetPixelId returns the PixelId field if non-nil, zero value otherwise.

### GetPixelIdOk

`func (o *CreateAdAudienceRequest) GetPixelIdOk() (*string, bool)`

GetPixelIdOk returns a tuple with the PixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelId

`func (o *CreateAdAudienceRequest) SetPixelId(v string)`

SetPixelId sets PixelId field to given value.

### HasPixelId

`func (o *CreateAdAudienceRequest) HasPixelId() bool`

HasPixelId returns a boolean if a field has been set.

### GetRetentionDays

`func (o *CreateAdAudienceRequest) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *CreateAdAudienceRequest) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *CreateAdAudienceRequest) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *CreateAdAudienceRequest) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetSourceAudienceId

`func (o *CreateAdAudienceRequest) GetSourceAudienceId() string`

GetSourceAudienceId returns the SourceAudienceId field if non-nil, zero value otherwise.

### GetSourceAudienceIdOk

`func (o *CreateAdAudienceRequest) GetSourceAudienceIdOk() (*string, bool)`

GetSourceAudienceIdOk returns a tuple with the SourceAudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAudienceId

`func (o *CreateAdAudienceRequest) SetSourceAudienceId(v string)`

SetSourceAudienceId sets SourceAudienceId field to given value.

### HasSourceAudienceId

`func (o *CreateAdAudienceRequest) HasSourceAudienceId() bool`

HasSourceAudienceId returns a boolean if a field has been set.

### GetCountry

`func (o *CreateAdAudienceRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateAdAudienceRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateAdAudienceRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *CreateAdAudienceRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRatio

`func (o *CreateAdAudienceRequest) GetRatio() float32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *CreateAdAudienceRequest) GetRatioOk() (*float32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *CreateAdAudienceRequest) SetRatio(v float32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *CreateAdAudienceRequest) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetRule

`func (o *CreateAdAudienceRequest) GetRule() map[string]interface{}`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *CreateAdAudienceRequest) GetRuleOk() (*map[string]interface{}, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *CreateAdAudienceRequest) SetRule(v map[string]interface{})`

SetRule sets Rule field to given value.

### HasRule

`func (o *CreateAdAudienceRequest) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetCustomerFileSource

`func (o *CreateAdAudienceRequest) GetCustomerFileSource() string`

GetCustomerFileSource returns the CustomerFileSource field if non-nil, zero value otherwise.

### GetCustomerFileSourceOk

`func (o *CreateAdAudienceRequest) GetCustomerFileSourceOk() (*string, bool)`

GetCustomerFileSourceOk returns a tuple with the CustomerFileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFileSource

`func (o *CreateAdAudienceRequest) SetCustomerFileSource(v string)`

SetCustomerFileSource sets CustomerFileSource field to given value.

### HasCustomerFileSource

`func (o *CreateAdAudienceRequest) HasCustomerFileSource() bool`

HasCustomerFileSource returns a boolean if a field has been set.

### GetSpec

`func (o *CreateAdAudienceRequest) GetSpec() TargetingSpec`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *CreateAdAudienceRequest) GetSpecOk() (*TargetingSpec, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *CreateAdAudienceRequest) SetSpec(v TargetingSpec)`

SetSpec sets Spec field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


