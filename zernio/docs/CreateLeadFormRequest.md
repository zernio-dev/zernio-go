# CreateLeadFormRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** |  | 
**Name** | **string** |  | 
**Questions** | [**[]CreateLeadFormRequestQuestionsInner**](CreateLeadFormRequestQuestionsInner.md) |  | 
**PrivacyPolicyUrl** | **string** |  | 
**PrivacyPolicyLinkText** | Pointer to **string** |  | [optional] 
**FollowUpActionUrl** | Pointer to **string** |  | [optional] 
**Locale** | Pointer to **string** |  | [optional] 
**ThankYouTitle** | Pointer to **string** |  | [optional] 
**ThankYouBody** | Pointer to **string** |  | [optional] 
**ThankYouButtonText** | Pointer to **string** |  | [optional] 
**ThankYouButtonType** | Pointer to **string** |  | [optional] 
**ThankYouWebsiteUrl** | Pointer to **string** |  | [optional] 
**IsOptimizedForQuality** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateLeadFormRequest

`func NewCreateLeadFormRequest(accountId string, name string, questions []CreateLeadFormRequestQuestionsInner, privacyPolicyUrl string, ) *CreateLeadFormRequest`

NewCreateLeadFormRequest instantiates a new CreateLeadFormRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLeadFormRequestWithDefaults

`func NewCreateLeadFormRequestWithDefaults() *CreateLeadFormRequest`

NewCreateLeadFormRequestWithDefaults instantiates a new CreateLeadFormRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *CreateLeadFormRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *CreateLeadFormRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *CreateLeadFormRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetName

`func (o *CreateLeadFormRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateLeadFormRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateLeadFormRequest) SetName(v string)`

SetName sets Name field to given value.


### GetQuestions

`func (o *CreateLeadFormRequest) GetQuestions() []CreateLeadFormRequestQuestionsInner`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *CreateLeadFormRequest) GetQuestionsOk() (*[]CreateLeadFormRequestQuestionsInner, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *CreateLeadFormRequest) SetQuestions(v []CreateLeadFormRequestQuestionsInner)`

SetQuestions sets Questions field to given value.


### GetPrivacyPolicyUrl

`func (o *CreateLeadFormRequest) GetPrivacyPolicyUrl() string`

GetPrivacyPolicyUrl returns the PrivacyPolicyUrl field if non-nil, zero value otherwise.

### GetPrivacyPolicyUrlOk

`func (o *CreateLeadFormRequest) GetPrivacyPolicyUrlOk() (*string, bool)`

GetPrivacyPolicyUrlOk returns a tuple with the PrivacyPolicyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyUrl

`func (o *CreateLeadFormRequest) SetPrivacyPolicyUrl(v string)`

SetPrivacyPolicyUrl sets PrivacyPolicyUrl field to given value.


### GetPrivacyPolicyLinkText

`func (o *CreateLeadFormRequest) GetPrivacyPolicyLinkText() string`

GetPrivacyPolicyLinkText returns the PrivacyPolicyLinkText field if non-nil, zero value otherwise.

### GetPrivacyPolicyLinkTextOk

`func (o *CreateLeadFormRequest) GetPrivacyPolicyLinkTextOk() (*string, bool)`

GetPrivacyPolicyLinkTextOk returns a tuple with the PrivacyPolicyLinkText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicyLinkText

`func (o *CreateLeadFormRequest) SetPrivacyPolicyLinkText(v string)`

SetPrivacyPolicyLinkText sets PrivacyPolicyLinkText field to given value.

### HasPrivacyPolicyLinkText

`func (o *CreateLeadFormRequest) HasPrivacyPolicyLinkText() bool`

HasPrivacyPolicyLinkText returns a boolean if a field has been set.

### GetFollowUpActionUrl

`func (o *CreateLeadFormRequest) GetFollowUpActionUrl() string`

GetFollowUpActionUrl returns the FollowUpActionUrl field if non-nil, zero value otherwise.

### GetFollowUpActionUrlOk

`func (o *CreateLeadFormRequest) GetFollowUpActionUrlOk() (*string, bool)`

GetFollowUpActionUrlOk returns a tuple with the FollowUpActionUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowUpActionUrl

`func (o *CreateLeadFormRequest) SetFollowUpActionUrl(v string)`

SetFollowUpActionUrl sets FollowUpActionUrl field to given value.

### HasFollowUpActionUrl

`func (o *CreateLeadFormRequest) HasFollowUpActionUrl() bool`

HasFollowUpActionUrl returns a boolean if a field has been set.

### GetLocale

`func (o *CreateLeadFormRequest) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *CreateLeadFormRequest) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *CreateLeadFormRequest) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *CreateLeadFormRequest) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetThankYouTitle

`func (o *CreateLeadFormRequest) GetThankYouTitle() string`

GetThankYouTitle returns the ThankYouTitle field if non-nil, zero value otherwise.

### GetThankYouTitleOk

`func (o *CreateLeadFormRequest) GetThankYouTitleOk() (*string, bool)`

GetThankYouTitleOk returns a tuple with the ThankYouTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouTitle

`func (o *CreateLeadFormRequest) SetThankYouTitle(v string)`

SetThankYouTitle sets ThankYouTitle field to given value.

### HasThankYouTitle

`func (o *CreateLeadFormRequest) HasThankYouTitle() bool`

HasThankYouTitle returns a boolean if a field has been set.

### GetThankYouBody

`func (o *CreateLeadFormRequest) GetThankYouBody() string`

GetThankYouBody returns the ThankYouBody field if non-nil, zero value otherwise.

### GetThankYouBodyOk

`func (o *CreateLeadFormRequest) GetThankYouBodyOk() (*string, bool)`

GetThankYouBodyOk returns a tuple with the ThankYouBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouBody

`func (o *CreateLeadFormRequest) SetThankYouBody(v string)`

SetThankYouBody sets ThankYouBody field to given value.

### HasThankYouBody

`func (o *CreateLeadFormRequest) HasThankYouBody() bool`

HasThankYouBody returns a boolean if a field has been set.

### GetThankYouButtonText

`func (o *CreateLeadFormRequest) GetThankYouButtonText() string`

GetThankYouButtonText returns the ThankYouButtonText field if non-nil, zero value otherwise.

### GetThankYouButtonTextOk

`func (o *CreateLeadFormRequest) GetThankYouButtonTextOk() (*string, bool)`

GetThankYouButtonTextOk returns a tuple with the ThankYouButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouButtonText

`func (o *CreateLeadFormRequest) SetThankYouButtonText(v string)`

SetThankYouButtonText sets ThankYouButtonText field to given value.

### HasThankYouButtonText

`func (o *CreateLeadFormRequest) HasThankYouButtonText() bool`

HasThankYouButtonText returns a boolean if a field has been set.

### GetThankYouButtonType

`func (o *CreateLeadFormRequest) GetThankYouButtonType() string`

GetThankYouButtonType returns the ThankYouButtonType field if non-nil, zero value otherwise.

### GetThankYouButtonTypeOk

`func (o *CreateLeadFormRequest) GetThankYouButtonTypeOk() (*string, bool)`

GetThankYouButtonTypeOk returns a tuple with the ThankYouButtonType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouButtonType

`func (o *CreateLeadFormRequest) SetThankYouButtonType(v string)`

SetThankYouButtonType sets ThankYouButtonType field to given value.

### HasThankYouButtonType

`func (o *CreateLeadFormRequest) HasThankYouButtonType() bool`

HasThankYouButtonType returns a boolean if a field has been set.

### GetThankYouWebsiteUrl

`func (o *CreateLeadFormRequest) GetThankYouWebsiteUrl() string`

GetThankYouWebsiteUrl returns the ThankYouWebsiteUrl field if non-nil, zero value otherwise.

### GetThankYouWebsiteUrlOk

`func (o *CreateLeadFormRequest) GetThankYouWebsiteUrlOk() (*string, bool)`

GetThankYouWebsiteUrlOk returns a tuple with the ThankYouWebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThankYouWebsiteUrl

`func (o *CreateLeadFormRequest) SetThankYouWebsiteUrl(v string)`

SetThankYouWebsiteUrl sets ThankYouWebsiteUrl field to given value.

### HasThankYouWebsiteUrl

`func (o *CreateLeadFormRequest) HasThankYouWebsiteUrl() bool`

HasThankYouWebsiteUrl returns a boolean if a field has been set.

### GetIsOptimizedForQuality

`func (o *CreateLeadFormRequest) GetIsOptimizedForQuality() bool`

GetIsOptimizedForQuality returns the IsOptimizedForQuality field if non-nil, zero value otherwise.

### GetIsOptimizedForQualityOk

`func (o *CreateLeadFormRequest) GetIsOptimizedForQualityOk() (*bool, bool)`

GetIsOptimizedForQualityOk returns a tuple with the IsOptimizedForQuality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOptimizedForQuality

`func (o *CreateLeadFormRequest) SetIsOptimizedForQuality(v bool)`

SetIsOptimizedForQuality sets IsOptimizedForQuality field to given value.

### HasIsOptimizedForQuality

`func (o *CreateLeadFormRequest) HasIsOptimizedForQuality() bool`

HasIsOptimizedForQuality returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


