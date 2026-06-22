# GetWhatsAppNumberInfo200ResponsePhone

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DisplayPhoneNumber** | Pointer to **string** |  | [optional] 
**VerifiedName** | Pointer to **string** |  | [optional] 
**NameStatus** | Pointer to **string** | APPROVED, AVAILABLE_WITHOUT_REVIEW, PENDING_REVIEW, DECLINED, EXPIRED, NONE | [optional] 
**QualityRating** | Pointer to **string** | GREEN, YELLOW, RED, UNKNOWN | [optional] 
**MessagingLimitTier** | Pointer to **string** | e.g. TIER_250, TIER_1K, TIER_UNLIMITED | [optional] 
**Throughput** | Pointer to [**GetWhatsAppNumberInfo200ResponsePhoneThroughput**](GetWhatsAppNumberInfo200ResponsePhoneThroughput.md) |  | [optional] 
**Status** | Pointer to **string** | e.g. CONNECTED | [optional] 
**IsOfficialBusinessAccount** | Pointer to **bool** |  | [optional] 
**PlatformType** | Pointer to **string** | e.g. CLOUD_API | [optional] 
**HealthStatus** | Pointer to **map[string]interface{}** | Meta&#39;s can_send_message health object (messaging + calling signals) | [optional] 

## Methods

### NewGetWhatsAppNumberInfo200ResponsePhone

`func NewGetWhatsAppNumberInfo200ResponsePhone() *GetWhatsAppNumberInfo200ResponsePhone`

NewGetWhatsAppNumberInfo200ResponsePhone instantiates a new GetWhatsAppNumberInfo200ResponsePhone object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWhatsAppNumberInfo200ResponsePhoneWithDefaults

`func NewGetWhatsAppNumberInfo200ResponsePhoneWithDefaults() *GetWhatsAppNumberInfo200ResponsePhone`

NewGetWhatsAppNumberInfo200ResponsePhoneWithDefaults instantiates a new GetWhatsAppNumberInfo200ResponsePhone object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisplayPhoneNumber

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetDisplayPhoneNumber() string`

GetDisplayPhoneNumber returns the DisplayPhoneNumber field if non-nil, zero value otherwise.

### GetDisplayPhoneNumberOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetDisplayPhoneNumberOk() (*string, bool)`

GetDisplayPhoneNumberOk returns a tuple with the DisplayPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayPhoneNumber

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetDisplayPhoneNumber(v string)`

SetDisplayPhoneNumber sets DisplayPhoneNumber field to given value.

### HasDisplayPhoneNumber

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasDisplayPhoneNumber() bool`

HasDisplayPhoneNumber returns a boolean if a field has been set.

### GetVerifiedName

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetVerifiedName() string`

GetVerifiedName returns the VerifiedName field if non-nil, zero value otherwise.

### GetVerifiedNameOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetVerifiedNameOk() (*string, bool)`

GetVerifiedNameOk returns a tuple with the VerifiedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedName

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetVerifiedName(v string)`

SetVerifiedName sets VerifiedName field to given value.

### HasVerifiedName

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasVerifiedName() bool`

HasVerifiedName returns a boolean if a field has been set.

### GetNameStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetNameStatus() string`

GetNameStatus returns the NameStatus field if non-nil, zero value otherwise.

### GetNameStatusOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetNameStatusOk() (*string, bool)`

GetNameStatusOk returns a tuple with the NameStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetNameStatus(v string)`

SetNameStatus sets NameStatus field to given value.

### HasNameStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasNameStatus() bool`

HasNameStatus returns a boolean if a field has been set.

### GetQualityRating

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetQualityRating() string`

GetQualityRating returns the QualityRating field if non-nil, zero value otherwise.

### GetQualityRatingOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetQualityRatingOk() (*string, bool)`

GetQualityRatingOk returns a tuple with the QualityRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityRating

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetQualityRating(v string)`

SetQualityRating sets QualityRating field to given value.

### HasQualityRating

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasQualityRating() bool`

HasQualityRating returns a boolean if a field has been set.

### GetMessagingLimitTier

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetMessagingLimitTier() string`

GetMessagingLimitTier returns the MessagingLimitTier field if non-nil, zero value otherwise.

### GetMessagingLimitTierOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetMessagingLimitTierOk() (*string, bool)`

GetMessagingLimitTierOk returns a tuple with the MessagingLimitTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingLimitTier

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetMessagingLimitTier(v string)`

SetMessagingLimitTier sets MessagingLimitTier field to given value.

### HasMessagingLimitTier

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasMessagingLimitTier() bool`

HasMessagingLimitTier returns a boolean if a field has been set.

### GetThroughput

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetThroughput() GetWhatsAppNumberInfo200ResponsePhoneThroughput`

GetThroughput returns the Throughput field if non-nil, zero value otherwise.

### GetThroughputOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetThroughputOk() (*GetWhatsAppNumberInfo200ResponsePhoneThroughput, bool)`

GetThroughputOk returns a tuple with the Throughput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThroughput

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetThroughput(v GetWhatsAppNumberInfo200ResponsePhoneThroughput)`

SetThroughput sets Throughput field to given value.

### HasThroughput

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasThroughput() bool`

HasThroughput returns a boolean if a field has been set.

### GetStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsOfficialBusinessAccount

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetIsOfficialBusinessAccount() bool`

GetIsOfficialBusinessAccount returns the IsOfficialBusinessAccount field if non-nil, zero value otherwise.

### GetIsOfficialBusinessAccountOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetIsOfficialBusinessAccountOk() (*bool, bool)`

GetIsOfficialBusinessAccountOk returns a tuple with the IsOfficialBusinessAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOfficialBusinessAccount

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetIsOfficialBusinessAccount(v bool)`

SetIsOfficialBusinessAccount sets IsOfficialBusinessAccount field to given value.

### HasIsOfficialBusinessAccount

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasIsOfficialBusinessAccount() bool`

HasIsOfficialBusinessAccount returns a boolean if a field has been set.

### GetPlatformType

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetPlatformType() string`

GetPlatformType returns the PlatformType field if non-nil, zero value otherwise.

### GetPlatformTypeOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetPlatformTypeOk() (*string, bool)`

GetPlatformTypeOk returns a tuple with the PlatformType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformType

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetPlatformType(v string)`

SetPlatformType sets PlatformType field to given value.

### HasPlatformType

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasPlatformType() bool`

HasPlatformType returns a boolean if a field has been set.

### GetHealthStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetHealthStatus() map[string]interface{}`

GetHealthStatus returns the HealthStatus field if non-nil, zero value otherwise.

### GetHealthStatusOk

`func (o *GetWhatsAppNumberInfo200ResponsePhone) GetHealthStatusOk() (*map[string]interface{}, bool)`

GetHealthStatusOk returns a tuple with the HealthStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) SetHealthStatus(v map[string]interface{})`

SetHealthStatus sets HealthStatus field to given value.

### HasHealthStatus

`func (o *GetWhatsAppNumberInfo200ResponsePhone) HasHealthStatus() bool`

HasHealthStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


