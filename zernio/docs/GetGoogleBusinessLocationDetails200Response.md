# GetGoogleBusinessLocationDetails200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **string** |  | [optional] 
**LocationId** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseLocation**](GetGoogleBusinessLocationDetails200ResponseLocation.md) |  | [optional] 
**Title** | Pointer to **string** | Business name | [optional] 
**RegularHours** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseRegularHours**](GetGoogleBusinessLocationDetails200ResponseRegularHours.md) |  | [optional] 
**SpecialHours** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseSpecialHours**](GetGoogleBusinessLocationDetails200ResponseSpecialHours.md) |  | [optional] 
**Profile** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseProfile**](GetGoogleBusinessLocationDetails200ResponseProfile.md) |  | [optional] 
**WebsiteUri** | Pointer to **string** |  | [optional] 
**PhoneNumbers** | Pointer to [**GetGoogleBusinessLocationDetails200ResponsePhoneNumbers**](GetGoogleBusinessLocationDetails200ResponsePhoneNumbers.md) |  | [optional] 
**Categories** | Pointer to [**GetGoogleBusinessLocationDetails200ResponseCategories**](GetGoogleBusinessLocationDetails200ResponseCategories.md) |  | [optional] 
**ServiceItems** | Pointer to [**[]GetGoogleBusinessLocationDetails200ResponseServiceItemsInner**](GetGoogleBusinessLocationDetails200ResponseServiceItemsInner.md) | Services offered (returned when readMask includes &#39;serviceItems&#39;) | [optional] 

## Methods

### NewGetGoogleBusinessLocationDetails200Response

`func NewGetGoogleBusinessLocationDetails200Response() *GetGoogleBusinessLocationDetails200Response`

NewGetGoogleBusinessLocationDetails200Response instantiates a new GetGoogleBusinessLocationDetails200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGoogleBusinessLocationDetails200ResponseWithDefaults

`func NewGetGoogleBusinessLocationDetails200ResponseWithDefaults() *GetGoogleBusinessLocationDetails200Response`

NewGetGoogleBusinessLocationDetails200ResponseWithDefaults instantiates a new GetGoogleBusinessLocationDetails200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *GetGoogleBusinessLocationDetails200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetGoogleBusinessLocationDetails200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetGoogleBusinessLocationDetails200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetAccountId

`func (o *GetGoogleBusinessLocationDetails200Response) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetGoogleBusinessLocationDetails200Response) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetGoogleBusinessLocationDetails200Response) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetLocationId

`func (o *GetGoogleBusinessLocationDetails200Response) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GetGoogleBusinessLocationDetails200Response) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GetGoogleBusinessLocationDetails200Response) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLocation

`func (o *GetGoogleBusinessLocationDetails200Response) GetLocation() GetGoogleBusinessLocationDetails200ResponseLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetLocationOk() (*GetGoogleBusinessLocationDetails200ResponseLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetGoogleBusinessLocationDetails200Response) SetLocation(v GetGoogleBusinessLocationDetails200ResponseLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetGoogleBusinessLocationDetails200Response) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetTitle

`func (o *GetGoogleBusinessLocationDetails200Response) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetGoogleBusinessLocationDetails200Response) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetGoogleBusinessLocationDetails200Response) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetRegularHours

`func (o *GetGoogleBusinessLocationDetails200Response) GetRegularHours() GetGoogleBusinessLocationDetails200ResponseRegularHours`

GetRegularHours returns the RegularHours field if non-nil, zero value otherwise.

### GetRegularHoursOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetRegularHoursOk() (*GetGoogleBusinessLocationDetails200ResponseRegularHours, bool)`

GetRegularHoursOk returns a tuple with the RegularHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegularHours

`func (o *GetGoogleBusinessLocationDetails200Response) SetRegularHours(v GetGoogleBusinessLocationDetails200ResponseRegularHours)`

SetRegularHours sets RegularHours field to given value.

### HasRegularHours

`func (o *GetGoogleBusinessLocationDetails200Response) HasRegularHours() bool`

HasRegularHours returns a boolean if a field has been set.

### GetSpecialHours

`func (o *GetGoogleBusinessLocationDetails200Response) GetSpecialHours() GetGoogleBusinessLocationDetails200ResponseSpecialHours`

GetSpecialHours returns the SpecialHours field if non-nil, zero value otherwise.

### GetSpecialHoursOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetSpecialHoursOk() (*GetGoogleBusinessLocationDetails200ResponseSpecialHours, bool)`

GetSpecialHoursOk returns a tuple with the SpecialHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialHours

`func (o *GetGoogleBusinessLocationDetails200Response) SetSpecialHours(v GetGoogleBusinessLocationDetails200ResponseSpecialHours)`

SetSpecialHours sets SpecialHours field to given value.

### HasSpecialHours

`func (o *GetGoogleBusinessLocationDetails200Response) HasSpecialHours() bool`

HasSpecialHours returns a boolean if a field has been set.

### GetProfile

`func (o *GetGoogleBusinessLocationDetails200Response) GetProfile() GetGoogleBusinessLocationDetails200ResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetProfileOk() (*GetGoogleBusinessLocationDetails200ResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetGoogleBusinessLocationDetails200Response) SetProfile(v GetGoogleBusinessLocationDetails200ResponseProfile)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *GetGoogleBusinessLocationDetails200Response) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetWebsiteUri

`func (o *GetGoogleBusinessLocationDetails200Response) GetWebsiteUri() string`

GetWebsiteUri returns the WebsiteUri field if non-nil, zero value otherwise.

### GetWebsiteUriOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetWebsiteUriOk() (*string, bool)`

GetWebsiteUriOk returns a tuple with the WebsiteUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUri

`func (o *GetGoogleBusinessLocationDetails200Response) SetWebsiteUri(v string)`

SetWebsiteUri sets WebsiteUri field to given value.

### HasWebsiteUri

`func (o *GetGoogleBusinessLocationDetails200Response) HasWebsiteUri() bool`

HasWebsiteUri returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *GetGoogleBusinessLocationDetails200Response) GetPhoneNumbers() GetGoogleBusinessLocationDetails200ResponsePhoneNumbers`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetPhoneNumbersOk() (*GetGoogleBusinessLocationDetails200ResponsePhoneNumbers, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *GetGoogleBusinessLocationDetails200Response) SetPhoneNumbers(v GetGoogleBusinessLocationDetails200ResponsePhoneNumbers)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *GetGoogleBusinessLocationDetails200Response) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetCategories

`func (o *GetGoogleBusinessLocationDetails200Response) GetCategories() GetGoogleBusinessLocationDetails200ResponseCategories`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetCategoriesOk() (*GetGoogleBusinessLocationDetails200ResponseCategories, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *GetGoogleBusinessLocationDetails200Response) SetCategories(v GetGoogleBusinessLocationDetails200ResponseCategories)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *GetGoogleBusinessLocationDetails200Response) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetServiceItems

`func (o *GetGoogleBusinessLocationDetails200Response) GetServiceItems() []GetGoogleBusinessLocationDetails200ResponseServiceItemsInner`

GetServiceItems returns the ServiceItems field if non-nil, zero value otherwise.

### GetServiceItemsOk

`func (o *GetGoogleBusinessLocationDetails200Response) GetServiceItemsOk() (*[]GetGoogleBusinessLocationDetails200ResponseServiceItemsInner, bool)`

GetServiceItemsOk returns a tuple with the ServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItems

`func (o *GetGoogleBusinessLocationDetails200Response) SetServiceItems(v []GetGoogleBusinessLocationDetails200ResponseServiceItemsInner)`

SetServiceItems sets ServiceItems field to given value.

### HasServiceItems

`func (o *GetGoogleBusinessLocationDetails200Response) HasServiceItems() bool`

HasServiceItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


