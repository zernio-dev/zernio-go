# UploadedOrDerivedAudience

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
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

## Methods

### NewUploadedOrDerivedAudience

`func NewUploadedOrDerivedAudience(accountId string, adAccountId string, name string, type_ string, ) *UploadedOrDerivedAudience`

NewUploadedOrDerivedAudience instantiates a new UploadedOrDerivedAudience object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadedOrDerivedAudienceWithDefaults

`func NewUploadedOrDerivedAudienceWithDefaults() *UploadedOrDerivedAudience`

NewUploadedOrDerivedAudienceWithDefaults instantiates a new UploadedOrDerivedAudience object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UploadedOrDerivedAudience) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UploadedOrDerivedAudience) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UploadedOrDerivedAudience) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAdAccountId

`func (o *UploadedOrDerivedAudience) GetAdAccountId() string`

GetAdAccountId returns the AdAccountId field if non-nil, zero value otherwise.

### GetAdAccountIdOk

`func (o *UploadedOrDerivedAudience) GetAdAccountIdOk() (*string, bool)`

GetAdAccountIdOk returns a tuple with the AdAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdAccountId

`func (o *UploadedOrDerivedAudience) SetAdAccountId(v string)`

SetAdAccountId sets AdAccountId field to given value.


### GetName

`func (o *UploadedOrDerivedAudience) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadedOrDerivedAudience) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadedOrDerivedAudience) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *UploadedOrDerivedAudience) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UploadedOrDerivedAudience) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UploadedOrDerivedAudience) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UploadedOrDerivedAudience) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *UploadedOrDerivedAudience) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UploadedOrDerivedAudience) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UploadedOrDerivedAudience) SetType(v string)`

SetType sets Type field to given value.


### GetPixelId

`func (o *UploadedOrDerivedAudience) GetPixelId() string`

GetPixelId returns the PixelId field if non-nil, zero value otherwise.

### GetPixelIdOk

`func (o *UploadedOrDerivedAudience) GetPixelIdOk() (*string, bool)`

GetPixelIdOk returns a tuple with the PixelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPixelId

`func (o *UploadedOrDerivedAudience) SetPixelId(v string)`

SetPixelId sets PixelId field to given value.

### HasPixelId

`func (o *UploadedOrDerivedAudience) HasPixelId() bool`

HasPixelId returns a boolean if a field has been set.

### GetRetentionDays

`func (o *UploadedOrDerivedAudience) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *UploadedOrDerivedAudience) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *UploadedOrDerivedAudience) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.

### HasRetentionDays

`func (o *UploadedOrDerivedAudience) HasRetentionDays() bool`

HasRetentionDays returns a boolean if a field has been set.

### GetSourceAudienceId

`func (o *UploadedOrDerivedAudience) GetSourceAudienceId() string`

GetSourceAudienceId returns the SourceAudienceId field if non-nil, zero value otherwise.

### GetSourceAudienceIdOk

`func (o *UploadedOrDerivedAudience) GetSourceAudienceIdOk() (*string, bool)`

GetSourceAudienceIdOk returns a tuple with the SourceAudienceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceAudienceId

`func (o *UploadedOrDerivedAudience) SetSourceAudienceId(v string)`

SetSourceAudienceId sets SourceAudienceId field to given value.

### HasSourceAudienceId

`func (o *UploadedOrDerivedAudience) HasSourceAudienceId() bool`

HasSourceAudienceId returns a boolean if a field has been set.

### GetCountry

`func (o *UploadedOrDerivedAudience) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UploadedOrDerivedAudience) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UploadedOrDerivedAudience) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *UploadedOrDerivedAudience) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetRatio

`func (o *UploadedOrDerivedAudience) GetRatio() float32`

GetRatio returns the Ratio field if non-nil, zero value otherwise.

### GetRatioOk

`func (o *UploadedOrDerivedAudience) GetRatioOk() (*float32, bool)`

GetRatioOk returns a tuple with the Ratio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRatio

`func (o *UploadedOrDerivedAudience) SetRatio(v float32)`

SetRatio sets Ratio field to given value.

### HasRatio

`func (o *UploadedOrDerivedAudience) HasRatio() bool`

HasRatio returns a boolean if a field has been set.

### GetRule

`func (o *UploadedOrDerivedAudience) GetRule() map[string]interface{}`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UploadedOrDerivedAudience) GetRuleOk() (*map[string]interface{}, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UploadedOrDerivedAudience) SetRule(v map[string]interface{})`

SetRule sets Rule field to given value.

### HasRule

`func (o *UploadedOrDerivedAudience) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetCustomerFileSource

`func (o *UploadedOrDerivedAudience) GetCustomerFileSource() string`

GetCustomerFileSource returns the CustomerFileSource field if non-nil, zero value otherwise.

### GetCustomerFileSourceOk

`func (o *UploadedOrDerivedAudience) GetCustomerFileSourceOk() (*string, bool)`

GetCustomerFileSourceOk returns a tuple with the CustomerFileSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerFileSource

`func (o *UploadedOrDerivedAudience) SetCustomerFileSource(v string)`

SetCustomerFileSource sets CustomerFileSource field to given value.

### HasCustomerFileSource

`func (o *UploadedOrDerivedAudience) HasCustomerFileSource() bool`

HasCustomerFileSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


