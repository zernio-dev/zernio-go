# FetchGoogleBusinessVerificationOptionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LanguageCode** | **string** |  | 
**Context** | Pointer to **map[string]interface{}** | ServiceBusinessContext. Required for service-area businesses (must include the service address). | [optional] 

## Methods

### NewFetchGoogleBusinessVerificationOptionsRequest

`func NewFetchGoogleBusinessVerificationOptionsRequest(languageCode string, ) *FetchGoogleBusinessVerificationOptionsRequest`

NewFetchGoogleBusinessVerificationOptionsRequest instantiates a new FetchGoogleBusinessVerificationOptionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchGoogleBusinessVerificationOptionsRequestWithDefaults

`func NewFetchGoogleBusinessVerificationOptionsRequestWithDefaults() *FetchGoogleBusinessVerificationOptionsRequest`

NewFetchGoogleBusinessVerificationOptionsRequestWithDefaults instantiates a new FetchGoogleBusinessVerificationOptionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLanguageCode

`func (o *FetchGoogleBusinessVerificationOptionsRequest) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *FetchGoogleBusinessVerificationOptionsRequest) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *FetchGoogleBusinessVerificationOptionsRequest) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.


### GetContext

`func (o *FetchGoogleBusinessVerificationOptionsRequest) GetContext() map[string]interface{}`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *FetchGoogleBusinessVerificationOptionsRequest) GetContextOk() (*map[string]interface{}, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *FetchGoogleBusinessVerificationOptionsRequest) SetContext(v map[string]interface{})`

SetContext sets Context field to given value.

### HasContext

`func (o *FetchGoogleBusinessVerificationOptionsRequest) HasContext() bool`

HasContext returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


