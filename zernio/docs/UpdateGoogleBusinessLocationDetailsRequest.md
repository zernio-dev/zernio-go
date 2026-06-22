# UpdateGoogleBusinessLocationDetailsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateMask** | **string** | Required. Comma-separated fields to update (e.g. &#39;regularHours&#39;, &#39;specialHours&#39;, &#39;profile.description&#39;, &#39;categories&#39;, &#39;serviceItems&#39;). Any valid Google Business Information API updateMask field is supported. | 
**RegularHours** | Pointer to [**UpdateGoogleBusinessLocationDetailsRequestRegularHours**](UpdateGoogleBusinessLocationDetailsRequestRegularHours.md) |  | [optional] 
**SpecialHours** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseSpecialHours**](GetGoogleBusinessLocationDetails200ResponseSpecialHours.md) |  | [optional] 
**Profile** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseProfile**](GetGoogleBusinessLocationDetails200ResponseProfile.md) |  | [optional] 
**WebsiteUri** | Pointer to **string** |  | [optional] 
**PhoneNumbers** | Pointer to [**GetGoogleBusinessLocationDetails200ResponsePhoneNumbers**](GetGoogleBusinessLocationDetails200ResponsePhoneNumbers.md) |  | [optional] 
**Categories** | Pointer to [**UpdateGoogleBusinessLocationDetailsRequestCategories**](UpdateGoogleBusinessLocationDetailsRequestCategories.md) |  | [optional] 
**ServiceItems** | Pointer to [**[]GetGoogleBusinessLocationDetails200ResponseServiceItemsInner**](GetGoogleBusinessLocationDetails200ResponseServiceItemsInner.md) | Services offered by the business. Use updateMask&#x3D;&#39;serviceItems&#39; to update. | [optional] 

## Methods

### NewUpdateGoogleBusinessLocationDetailsRequest

`func NewUpdateGoogleBusinessLocationDetailsRequest(updateMask string, ) *UpdateGoogleBusinessLocationDetailsRequest`

NewUpdateGoogleBusinessLocationDetailsRequest instantiates a new UpdateGoogleBusinessLocationDetailsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGoogleBusinessLocationDetailsRequestWithDefaults

`func NewUpdateGoogleBusinessLocationDetailsRequestWithDefaults() *UpdateGoogleBusinessLocationDetailsRequest`

NewUpdateGoogleBusinessLocationDetailsRequestWithDefaults instantiates a new UpdateGoogleBusinessLocationDetailsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateMask

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetUpdateMask() string`

GetUpdateMask returns the UpdateMask field if non-nil, zero value otherwise.

### GetUpdateMaskOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetUpdateMaskOk() (*string, bool)`

GetUpdateMaskOk returns a tuple with the UpdateMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateMask

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetUpdateMask(v string)`

SetUpdateMask sets UpdateMask field to given value.


### GetRegularHours

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetRegularHours() UpdateGoogleBusinessLocationDetailsRequestRegularHours`

GetRegularHours returns the RegularHours field if non-nil, zero value otherwise.

### GetRegularHoursOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetRegularHoursOk() (*UpdateGoogleBusinessLocationDetailsRequestRegularHours, bool)`

GetRegularHoursOk returns a tuple with the RegularHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularHours

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetRegularHours(v UpdateGoogleBusinessLocationDetailsRequestRegularHours)`

SetRegularHours sets RegularHours field to given value.

### HasRegularHours

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasRegularHours() bool`

HasRegularHours returns a boolean if a field has been set.

### GetSpecialHours

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetSpecialHours() GetGoogleBusinessLocationDetails200ResponseSpecialHours`

GetSpecialHours returns the SpecialHours field if non-nil, zero value otherwise.

### GetSpecialHoursOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetSpecialHoursOk() (*GetGoogleBusinessLocationDetails200ResponseSpecialHours, bool)`

GetSpecialHoursOk returns a tuple with the SpecialHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialHours

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetSpecialHours(v GetGoogleBusinessLocationDetails200ResponseSpecialHours)`

SetSpecialHours sets SpecialHours field to given value.

### HasSpecialHours

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasSpecialHours() bool`

HasSpecialHours returns a boolean if a field has been set.

### GetProfile

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetProfile() GetGoogleBusinessLocationDetails200ResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetProfileOk() (*GetGoogleBusinessLocationDetails200ResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetProfile(v GetGoogleBusinessLocationDetails200ResponseProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetWebsiteUri

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetWebsiteUri() string`

GetWebsiteUri returns the WebsiteUri field if non-nil, zero value otherwise.

### GetWebsiteUriOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetWebsiteUriOk() (*string, bool)`

GetWebsiteUriOk returns a tuple with the WebsiteUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUri

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetWebsiteUri(v string)`

SetWebsiteUri sets WebsiteUri field to given value.

### HasWebsiteUri

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasWebsiteUri() bool`

HasWebsiteUri returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetPhoneNumbers() GetGoogleBusinessLocationDetails200ResponsePhoneNumbers`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetPhoneNumbersOk() (*GetGoogleBusinessLocationDetails200ResponsePhoneNumbers, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetPhoneNumbers(v GetGoogleBusinessLocationDetails200ResponsePhoneNumbers)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetCategories

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetCategories() UpdateGoogleBusinessLocationDetailsRequestCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetCategoriesOk() (*UpdateGoogleBusinessLocationDetailsRequestCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetCategories(v UpdateGoogleBusinessLocationDetailsRequestCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetServiceItems

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetServiceItems() []GetGoogleBusinessLocationDetails200ResponseServiceItemsInner`

GetServiceItems returns the ServiceItems field if non-nil, zero value otherwise.

### GetServiceItemsOk

`func (o *UpdateGoogleBusinessLocationDetailsRequest) GetServiceItemsOk() (*[]GetGoogleBusinessLocationDetails200ResponseServiceItemsInner, bool)`

GetServiceItemsOk returns a tuple with the ServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItems

`func (o *UpdateGoogleBusinessLocationDetailsRequest) SetServiceItems(v []GetGoogleBusinessLocationDetails200ResponseServiceItemsInner)`

SetServiceItems sets ServiceItems field to given value.

### HasServiceItems

`func (o *UpdateGoogleBusinessLocationDetailsRequest) HasServiceItems() bool`

HasServiceItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


