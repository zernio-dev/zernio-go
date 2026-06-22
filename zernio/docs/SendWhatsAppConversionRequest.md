# SendWhatsAppConversionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp SocialAccount ID. | 
**EventName** | **string** | Live-verified allowlist of event names accepted by Meta&#39;s CAPI for Business Messaging (Graph API v25.0). Other standard pixel events including &#x60;Lead&#x60;, &#x60;CompleteRegistration&#x60;, &#x60;Subscribe&#x60;, &#x60;Schedule&#x60;, &#x60;Contact&#x60;, &#x60;StartTrial&#x60;, &#x60;AddPaymentInfo&#x60;, &#x60;Search&#x60;, and &#x60;SubmitApplication&#x60; are rejected with subcode 2804066 (\&quot;Messaging Event Invalid Event Type\&quot;) on &#x60;action_source &#x3D; business_messaging&#x60; events. Custom event names are also rejected.  Use &#x60;LeadSubmitted&#x60; (NOT &#x60;Lead&#x60;) for lead-style conversions.  | 
**EventTime** | Pointer to **float32** | Unix seconds. Defaults to the time of the request when omitted. Meta&#39;s attribution window is 7 days from click; events older than that lose attribution.  | [optional] 
**EventId** | **string** | Stable dedup key. Reuse to suppress duplicate events (Meta dedupes against pixel events with the same id).  | 
**ConversationId** | Pointer to **string** | Zernio Conversation &#x60;_id&#x60; (preferred lookup). The conversation must have a captured &#x60;ctwa_clid&#x60; in metadata (set automatically by the WhatsApp webhook on the first inbound message after a CTWA ad click).  | [optional] 
**PhoneE164** | Pointer to **string** | Contact phone number, digits only with no &#39;+&#39;. When used in lieu of &#x60;conversationId&#x60;, the handler resolves to the most recent CTWA-attributed conversation for this phone on the supplied account.  | [optional] 
**Value** | Pointer to **float32** | Conversion value (e.g. order total). | [optional] 
**Currency** | Pointer to **string** | ISO 4217 currency code (e.g. &#x60;USD&#x60;). | [optional] 
**ContentIds** | Pointer to **[]string** | Optional product / content identifiers. | [optional] 
**Email** | Pointer to **string** | User email. Normalized + SHA-256 hashed before sending to Meta. | [optional] 
**ExternalId** | Pointer to **string** | Stable customer identifier. Lowercased + SHA-256 hashed before sending to Meta.  | [optional] 
**TestCode** | Pointer to **string** | Meta &#x60;test_event_code&#x60; passthrough. Routes the event to the Test Events tab in Events Manager instead of the production dataset, useful for development.  | [optional] 

## Methods

### NewSendWhatsAppConversionRequest

`func NewSendWhatsAppConversionRequest(accountId string, eventName string, eventId string, ) *SendWhatsAppConversionRequest`

NewSendWhatsAppConversionRequest instantiates a new SendWhatsAppConversionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendWhatsAppConversionRequestWithDefaults

`func NewSendWhatsAppConversionRequestWithDefaults() *SendWhatsAppConversionRequest`

NewSendWhatsAppConversionRequestWithDefaults instantiates a new SendWhatsAppConversionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *SendWhatsAppConversionRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *SendWhatsAppConversionRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *SendWhatsAppConversionRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetEventName

`func (o *SendWhatsAppConversionRequest) GetEventName() string`

GetEventName returns the EventName field if non-nil, zero value otherwise.

### GetEventNameOk

`func (o *SendWhatsAppConversionRequest) GetEventNameOk() (*string, bool)`

GetEventNameOk returns a tuple with the EventName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventName

`func (o *SendWhatsAppConversionRequest) SetEventName(v string)`

SetEventName sets EventName field to given value.


### GetEventTime

`func (o *SendWhatsAppConversionRequest) GetEventTime() float32`

GetEventTime returns the EventTime field if non-nil, zero value otherwise.

### GetEventTimeOk

`func (o *SendWhatsAppConversionRequest) GetEventTimeOk() (*float32, bool)`

GetEventTimeOk returns a tuple with the EventTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventTime

`func (o *SendWhatsAppConversionRequest) SetEventTime(v float32)`

SetEventTime sets EventTime field to given value.

### HasEventTime

`func (o *SendWhatsAppConversionRequest) HasEventTime() bool`

HasEventTime returns a boolean if a field has been set.

### GetEventId

`func (o *SendWhatsAppConversionRequest) GetEventId() string`

GetEventId returns the EventId field if non-nil, zero value otherwise.

### GetEventIdOk

`func (o *SendWhatsAppConversionRequest) GetEventIdOk() (*string, bool)`

GetEventIdOk returns a tuple with the EventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventId

`func (o *SendWhatsAppConversionRequest) SetEventId(v string)`

SetEventId sets EventId field to given value.


### GetConversationId

`func (o *SendWhatsAppConversionRequest) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *SendWhatsAppConversionRequest) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *SendWhatsAppConversionRequest) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *SendWhatsAppConversionRequest) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetPhoneE164

`func (o *SendWhatsAppConversionRequest) GetPhoneE164() string`

GetPhoneE164 returns the PhoneE164 field if non-nil, zero value otherwise.

### GetPhoneE164Ok

`func (o *SendWhatsAppConversionRequest) GetPhoneE164Ok() (*string, bool)`

GetPhoneE164Ok returns a tuple with the PhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneE164

`func (o *SendWhatsAppConversionRequest) SetPhoneE164(v string)`

SetPhoneE164 sets PhoneE164 field to given value.

### HasPhoneE164

`func (o *SendWhatsAppConversionRequest) HasPhoneE164() bool`

HasPhoneE164 returns a boolean if a field has been set.

### GetValue

`func (o *SendWhatsAppConversionRequest) GetValue() float32`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SendWhatsAppConversionRequest) GetValueOk() (*float32, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SendWhatsAppConversionRequest) SetValue(v float32)`

SetValue sets Value field to given value.

### HasValue

`func (o *SendWhatsAppConversionRequest) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetCurrency

`func (o *SendWhatsAppConversionRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SendWhatsAppConversionRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SendWhatsAppConversionRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SendWhatsAppConversionRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetContentIds

`func (o *SendWhatsAppConversionRequest) GetContentIds() []string`

GetContentIds returns the ContentIds field if non-nil, zero value otherwise.

### GetContentIdsOk

`func (o *SendWhatsAppConversionRequest) GetContentIdsOk() (*[]string, bool)`

GetContentIdsOk returns a tuple with the ContentIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentIds

`func (o *SendWhatsAppConversionRequest) SetContentIds(v []string)`

SetContentIds sets ContentIds field to given value.

### HasContentIds

`func (o *SendWhatsAppConversionRequest) HasContentIds() bool`

HasContentIds returns a boolean if a field has been set.

### GetEmail

`func (o *SendWhatsAppConversionRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SendWhatsAppConversionRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SendWhatsAppConversionRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SendWhatsAppConversionRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExternalId

`func (o *SendWhatsAppConversionRequest) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SendWhatsAppConversionRequest) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SendWhatsAppConversionRequest) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SendWhatsAppConversionRequest) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetTestCode

`func (o *SendWhatsAppConversionRequest) GetTestCode() string`

GetTestCode returns the TestCode field if non-nil, zero value otherwise.

### GetTestCodeOk

`func (o *SendWhatsAppConversionRequest) GetTestCodeOk() (*string, bool)`

GetTestCodeOk returns a tuple with the TestCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestCode

`func (o *SendWhatsAppConversionRequest) SetTestCode(v string)`

SetTestCode sets TestCode field to given value.

### HasTestCode

`func (o *SendWhatsAppConversionRequest) HasTestCode() bool`

HasTestCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


