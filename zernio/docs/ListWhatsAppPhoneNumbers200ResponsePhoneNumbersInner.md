# ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Phone Number ID (Meta) | [optional] 
**DisplayPhoneNumber** | Pointer to **string** | E.164-formatted display number | [optional] 
**VerifiedName** | Pointer to **string** | Meta-verified business name | [optional] 
**QualityRating** | Pointer to **string** | GREEN, YELLOW, RED, or UNKNOWN | [optional] 
**NameStatus** | Pointer to **string** | APPROVED, PENDING_REVIEW, DECLINED, or NONE | [optional] 
**MessagingLimitTier** | Pointer to **string** | TIER_250, TIER_1K, TIER_10K, TIER_100K, or TIER_UNLIMITED | [optional] 
**WabaId** | Pointer to **string** | WhatsApp Business Account ID (Zernio enrichment) | [optional] 
**WabaName** | Pointer to **string** | WABA display name (Zernio enrichment) | [optional] 

## Methods

### NewListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner

`func NewListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner() *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner`

NewListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner instantiates a new ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWhatsAppPhoneNumbers200ResponsePhoneNumbersInnerWithDefaults

`func NewListWhatsAppPhoneNumbers200ResponsePhoneNumbersInnerWithDefaults() *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner`

NewListWhatsAppPhoneNumbers200ResponsePhoneNumbersInnerWithDefaults instantiates a new ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDisplayPhoneNumber

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetDisplayPhoneNumber() string`

GetDisplayPhoneNumber returns the DisplayPhoneNumber field if non-nil, zero value otherwise.

### GetDisplayPhoneNumberOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetDisplayPhoneNumberOk() (*string, bool)`

GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPhoneNumber

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetDisplayPhoneNumber(v string)`

SetDisplayPhoneNumber sets DisplayPhoneNumber field to given value.

### HasDisplayPhoneNumber

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasDisplayPhoneNumber() bool`

HasDisplayPhoneNumber returns a boolean if a field has been set.

### GetVerifiedName

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetVerifiedName() string`

GetVerifiedName returns the VerifiedName field if non-nil, zero value otherwise.

### GetVerifiedNameOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetVerifiedNameOk() (*string, bool)`

GetVerifiedNameOk returns a tuple with the VerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedName

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetVerifiedName(v string)`

SetVerifiedName sets VerifiedName field to given value.

### HasVerifiedName

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasVerifiedName() bool`

HasVerifiedName returns a boolean if a field has been set.

### GetQualityRating

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetQualityRating() string`

GetQualityRating returns the QualityRating field if non-nil, zero value otherwise.

### GetQualityRatingOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetQualityRatingOk() (*string, bool)`

GetQualityRatingOk returns a tuple with the QualityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityRating

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetQualityRating(v string)`

SetQualityRating sets QualityRating field to given value.

### HasQualityRating

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasQualityRating() bool`

HasQualityRating returns a boolean if a field has been set.

### GetNameStatus

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetNameStatus() string`

GetNameStatus returns the NameStatus field if non-nil, zero value otherwise.

### GetNameStatusOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetNameStatusOk() (*string, bool)`

GetNameStatusOk returns a tuple with the NameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameStatus

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetNameStatus(v string)`

SetNameStatus sets NameStatus field to given value.

### HasNameStatus

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasNameStatus() bool`

HasNameStatus returns a boolean if a field has been set.

### GetMessagingLimitTier

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetMessagingLimitTier() string`

GetMessagingLimitTier returns the MessagingLimitTier field if non-nil, zero value otherwise.

### GetMessagingLimitTierOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetMessagingLimitTierOk() (*string, bool)`

GetMessagingLimitTierOk returns a tuple with the MessagingLimitTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingLimitTier

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetMessagingLimitTier(v string)`

SetMessagingLimitTier sets MessagingLimitTier field to given value.

### HasMessagingLimitTier

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasMessagingLimitTier() bool`

HasMessagingLimitTier returns a boolean if a field has been set.

### GetWabaId

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetWabaId() string`

GetWabaId returns the WabaId field if non-nil, zero value otherwise.

### GetWabaIdOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetWabaIdOk() (*string, bool)`

GetWabaIdOk returns a tuple with the WabaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaId

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetWabaId(v string)`

SetWabaId sets WabaId field to given value.

### HasWabaId

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasWabaId() bool`

HasWabaId returns a boolean if a field has been set.

### GetWabaName

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetWabaName() string`

GetWabaName returns the WabaName field if non-nil, zero value otherwise.

### GetWabaNameOk

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) GetWabaNameOk() (*string, bool)`

GetWabaNameOk returns a tuple with the WabaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWabaName

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) SetWabaName(v string)`

SetWabaName sets WabaName field to given value.

### HasWabaName

`func (o *ListWhatsAppPhoneNumbers200ResponsePhoneNumbersInner) HasWabaName() bool`

HasWabaName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


