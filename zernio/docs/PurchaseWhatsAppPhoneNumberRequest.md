# PurchaseWhatsAppPhoneNumberRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProfileId** | **string** | Profile to associate the number with | 
**Country** | Pointer to **string** | ISO 3166-1 alpha-2 country for the number (default US). International numbers require usage-based billing. Tier 3/4 countries return 202 { status: \&quot;kyc_required\&quot;, kycUrl } — the customer must complete KYC at that URL before the number is ordered. See GET /v1/whatsapp/phone-numbers/countries.  | [optional] [default to "US"]
**PurchaseIntentId** | Pointer to **string** | Optional idempotency key. Send the same value when retrying a purchase: if a number was already bought under this key, the API returns { status: \&quot;already_purchased\&quot;, numberId, phoneNumber } instead of provisioning a second number. Generate a fresh key for each genuinely new purchase.  | [optional] 
**AllowMultiple** | Pointer to **bool** | Any second purchase within 10 minutes of a previous one is rejected with 409 code PURCHASE_VELOCITY as duplicate protection. Pass true to confirm the additional purchase is intentional (e.g. bulk provisioning).  | [optional] [default to false]

## Methods

### NewPurchaseWhatsAppPhoneNumberRequest

`func NewPurchaseWhatsAppPhoneNumberRequest(profileId string, ) *PurchaseWhatsAppPhoneNumberRequest`

NewPurchaseWhatsAppPhoneNumberRequest instantiates a new PurchaseWhatsAppPhoneNumberRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPurchaseWhatsAppPhoneNumberRequestWithDefaults

`func NewPurchaseWhatsAppPhoneNumberRequestWithDefaults() *PurchaseWhatsAppPhoneNumberRequest`

NewPurchaseWhatsAppPhoneNumberRequestWithDefaults instantiates a new PurchaseWhatsAppPhoneNumberRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProfileId

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetProfileId() string`

GetProfileId returns the ProfileId field if non-nil, zero value otherwise.

### GetProfileIdOk

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetProfileIdOk() (*string, bool)`

GetProfileIdOk returns a tuple with the ProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileId

`func (o *PurchaseWhatsAppPhoneNumberRequest) SetProfileId(v string)`

SetProfileId sets ProfileId field to given value.


### GetCountry

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PurchaseWhatsAppPhoneNumberRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *PurchaseWhatsAppPhoneNumberRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetPurchaseIntentId

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetPurchaseIntentId() string`

GetPurchaseIntentId returns the PurchaseIntentId field if non-nil, zero value otherwise.

### GetPurchaseIntentIdOk

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetPurchaseIntentIdOk() (*string, bool)`

GetPurchaseIntentIdOk returns a tuple with the PurchaseIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseIntentId

`func (o *PurchaseWhatsAppPhoneNumberRequest) SetPurchaseIntentId(v string)`

SetPurchaseIntentId sets PurchaseIntentId field to given value.

### HasPurchaseIntentId

`func (o *PurchaseWhatsAppPhoneNumberRequest) HasPurchaseIntentId() bool`

HasPurchaseIntentId returns a boolean if a field has been set.

### GetAllowMultiple

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetAllowMultiple() bool`

GetAllowMultiple returns the AllowMultiple field if non-nil, zero value otherwise.

### GetAllowMultipleOk

`func (o *PurchaseWhatsAppPhoneNumberRequest) GetAllowMultipleOk() (*bool, bool)`

GetAllowMultipleOk returns a tuple with the AllowMultiple field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowMultiple

`func (o *PurchaseWhatsAppPhoneNumberRequest) SetAllowMultiple(v bool)`

SetAllowMultiple sets AllowMultiple field to given value.

### HasAllowMultiple

`func (o *PurchaseWhatsAppPhoneNumberRequest) HasAllowMultiple() bool`

HasAllowMultiple returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


