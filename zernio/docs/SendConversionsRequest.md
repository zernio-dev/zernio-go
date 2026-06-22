# SendConversionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | SocialAccount ID (metaads, googleads, linkedinads, or tiktokads). | 
**DestinationId** | **string** | Platform destination identifier. For Meta, the pixel/dataset ID. For Google, the conversion action resource name. For LinkedIn, the conversion rule ID or full &#x60;urn:lla:llaPartnerConversion:{id}&#x60; URN.  | 
**Events** | [**[]ConversionEvent**](ConversionEvent.md) |  | 
**TestCode** | Pointer to **string** | Meta &#x60;test_event_code&#x60; passthrough. Ignored by Google and LinkedIn. | [optional] 
**Consent** | Pointer to [**SendConversionsRequestConsent**](SendConversionsRequestConsent.md) |  | [optional] 

## Methods

### NewSendConversionsRequest

`func NewSendConversionsRequest(accountId string, destinationId string, events []ConversionEvent, ) *SendConversionsRequest`

NewSendConversionsRequest instantiates a new SendConversionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendConversionsRequestWithDefaults

`func NewSendConversionsRequestWithDefaults() *SendConversionsRequest`

NewSendConversionsRequestWithDefaults instantiates a new SendConversionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SendConversionsRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SendConversionsRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SendConversionsRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetDestinationId

`func (o *SendConversionsRequest) GetDestinationId() string`

GetDestinationId returns the DestinationId field if non-nil, zero value otherwise.

### GetDestinationIdOk

`func (o *SendConversionsRequest) GetDestinationIdOk() (*string, bool)`

GetDestinationIdOk returns a tuple with the DestinationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationId

`func (o *SendConversionsRequest) SetDestinationId(v string)`

SetDestinationId sets DestinationId field to given value.


### GetEvents

`func (o *SendConversionsRequest) GetEvents() []ConversionEvent`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *SendConversionsRequest) GetEventsOk() (*[]ConversionEvent, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *SendConversionsRequest) SetEvents(v []ConversionEvent)`

SetEvents sets Events field to given value.


### GetTestCode

`func (o *SendConversionsRequest) GetTestCode() string`

GetTestCode returns the TestCode field if non-nil, zero value otherwise.

### GetTestCodeOk

`func (o *SendConversionsRequest) GetTestCodeOk() (*string, bool)`

GetTestCodeOk returns a tuple with the TestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCode

`func (o *SendConversionsRequest) SetTestCode(v string)`

SetTestCode sets TestCode field to given value.

### HasTestCode

`func (o *SendConversionsRequest) HasTestCode() bool`

HasTestCode returns a boolean if a field has been set.

### GetConsent

`func (o *SendConversionsRequest) GetConsent() SendConversionsRequestConsent`

GetConsent returns the Consent field if non-nil, zero value otherwise.

### GetConsentOk

`func (o *SendConversionsRequest) GetConsentOk() (*SendConversionsRequestConsent, bool)`

GetConsentOk returns a tuple with the Consent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsent

`func (o *SendConversionsRequest) SetConsent(v SendConversionsRequestConsent)`

SetConsent sets Consent field to given value.

### HasConsent

`func (o *SendConversionsRequest) HasConsent() bool`

HasConsent returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


