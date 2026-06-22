# WhatsAppSandboxSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Session id. Use this to revoke via DELETE. | 
**PhoneE164** | **string** | Digits-only E.164 form (no +, spaces, or dashes). | 
**Status** | **string** | &#x60;pending&#x60; until the phone replies to the activation template, then &#x60;active&#x60;. Expired sessions are pruned by TTL and never appear in list responses.  | 
**ExpiresAt** | **time.Time** | UTC timestamp at which the session becomes invalid. Pending sessions get a 24h window; activated sessions get 7 days.  | 
**ActivatedAt** | Pointer to **NullableTime** | When the session transitioned &#x60;pending → active&#x60;, or null. | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewWhatsAppSandboxSession

`func NewWhatsAppSandboxSession(id string, phoneE164 string, status string, expiresAt time.Time, ) *WhatsAppSandboxSession`

NewWhatsAppSandboxSession instantiates a new WhatsAppSandboxSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWhatsAppSandboxSessionWithDefaults

`func NewWhatsAppSandboxSessionWithDefaults() *WhatsAppSandboxSession`

NewWhatsAppSandboxSessionWithDefaults instantiates a new WhatsAppSandboxSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WhatsAppSandboxSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WhatsAppSandboxSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WhatsAppSandboxSession) SetId(v string)`

SetId sets Id field to given value.


### GetPhoneE164

`func (o *WhatsAppSandboxSession) GetPhoneE164() string`

GetPhoneE164 returns the PhoneE164 field if non-nil, zero value otherwise.

### GetPhoneE164Ok

`func (o *WhatsAppSandboxSession) GetPhoneE164Ok() (*string, bool)`

GetPhoneE164Ok returns a tuple with the PhoneE164 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneE164

`func (o *WhatsAppSandboxSession) SetPhoneE164(v string)`

SetPhoneE164 sets PhoneE164 field to given value.


### GetStatus

`func (o *WhatsAppSandboxSession) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WhatsAppSandboxSession) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WhatsAppSandboxSession) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetExpiresAt

`func (o *WhatsAppSandboxSession) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *WhatsAppSandboxSession) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *WhatsAppSandboxSession) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetActivatedAt

`func (o *WhatsAppSandboxSession) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *WhatsAppSandboxSession) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *WhatsAppSandboxSession) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.

### HasActivatedAt

`func (o *WhatsAppSandboxSession) HasActivatedAt() bool`

HasActivatedAt returns a boolean if a field has been set.

### SetActivatedAtNil

`func (o *WhatsAppSandboxSession) SetActivatedAtNil(b bool)`

 SetActivatedAtNil sets the value for ActivatedAt to be an explicit nil

### UnsetActivatedAt
`func (o *WhatsAppSandboxSession) UnsetActivatedAt()`

UnsetActivatedAt ensures that no value is present for ActivatedAt, not even an explicit nil
### GetCreatedAt

`func (o *WhatsAppSandboxSession) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WhatsAppSandboxSession) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WhatsAppSandboxSession) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *WhatsAppSandboxSession) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *WhatsAppSandboxSession) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *WhatsAppSandboxSession) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


