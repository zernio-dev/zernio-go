# UpdateWhatsAppBusinessProfileRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | WhatsApp social account ID | 
**About** | Pointer to **string** | Short business description (max 139 characters) | [optional] 
**Address** | Pointer to **string** | Business address | [optional] 
**Description** | Pointer to **string** | Full business description (max 512 characters) | [optional] 
**Email** | Pointer to **string** | Business email | [optional] 
**Websites** | Pointer to **[]string** | Business websites (max 2) | [optional] 
**Vertical** | Pointer to **string** | Business category (e.g., RETAIL, ENTERTAINMENT, etc.) | [optional] 
**ProfilePictureHandle** | Pointer to **string** | Handle from resumable upload for profile picture | [optional] 

## Methods

### NewUpdateWhatsAppBusinessProfileRequest

`func NewUpdateWhatsAppBusinessProfileRequest(accountId string, ) *UpdateWhatsAppBusinessProfileRequest`

NewUpdateWhatsAppBusinessProfileRequest instantiates a new UpdateWhatsAppBusinessProfileRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateWhatsAppBusinessProfileRequestWithDefaults

`func NewUpdateWhatsAppBusinessProfileRequestWithDefaults() *UpdateWhatsAppBusinessProfileRequest`

NewUpdateWhatsAppBusinessProfileRequestWithDefaults instantiates a new UpdateWhatsAppBusinessProfileRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateWhatsAppBusinessProfileRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateWhatsAppBusinessProfileRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetAbout

`func (o *UpdateWhatsAppBusinessProfileRequest) GetAbout() string`

GetAbout returns the About field if non-nil, zero value otherwise.

### GetAboutOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetAboutOk() (*string, bool)`

GetAboutOk returns a tuple with the About field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbout

`func (o *UpdateWhatsAppBusinessProfileRequest) SetAbout(v string)`

SetAbout sets About field to given value.

### HasAbout

`func (o *UpdateWhatsAppBusinessProfileRequest) HasAbout() bool`

HasAbout returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateWhatsAppBusinessProfileRequest) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateWhatsAppBusinessProfileRequest) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateWhatsAppBusinessProfileRequest) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateWhatsAppBusinessProfileRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateWhatsAppBusinessProfileRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateWhatsAppBusinessProfileRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateWhatsAppBusinessProfileRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateWhatsAppBusinessProfileRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateWhatsAppBusinessProfileRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetWebsites

`func (o *UpdateWhatsAppBusinessProfileRequest) GetWebsites() []string`

GetWebsites returns the Websites field if non-nil, zero value otherwise.

### GetWebsitesOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetWebsitesOk() (*[]string, bool)`

GetWebsitesOk returns a tuple with the Websites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsites

`func (o *UpdateWhatsAppBusinessProfileRequest) SetWebsites(v []string)`

SetWebsites sets Websites field to given value.

### HasWebsites

`func (o *UpdateWhatsAppBusinessProfileRequest) HasWebsites() bool`

HasWebsites returns a boolean if a field has been set.

### GetVertical

`func (o *UpdateWhatsAppBusinessProfileRequest) GetVertical() string`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetVerticalOk() (*string, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *UpdateWhatsAppBusinessProfileRequest) SetVertical(v string)`

SetVertical sets Vertical field to given value.

### HasVertical

`func (o *UpdateWhatsAppBusinessProfileRequest) HasVertical() bool`

HasVertical returns a boolean if a field has been set.

### GetProfilePictureHandle

`func (o *UpdateWhatsAppBusinessProfileRequest) GetProfilePictureHandle() string`

GetProfilePictureHandle returns the ProfilePictureHandle field if non-nil, zero value otherwise.

### GetProfilePictureHandleOk

`func (o *UpdateWhatsAppBusinessProfileRequest) GetProfilePictureHandleOk() (*string, bool)`

GetProfilePictureHandleOk returns a tuple with the ProfilePictureHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfilePictureHandle

`func (o *UpdateWhatsAppBusinessProfileRequest) SetProfilePictureHandle(v string)`

SetProfilePictureHandle sets ProfilePictureHandle field to given value.

### HasProfilePictureHandle

`func (o *UpdateWhatsAppBusinessProfileRequest) HasProfilePictureHandle() bool`

HasProfilePictureHandle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


