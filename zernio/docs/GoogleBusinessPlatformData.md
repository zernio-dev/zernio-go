# GoogleBusinessPlatformData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** | Target GBP location ID (e.g. \&quot;locations/123456789\&quot;). If omitted, uses the default location. Use GET /v1/accounts/{id}/gmb-locations to list locations. | [optional] 
**LanguageCode** | Pointer to **string** | BCP 47 language code (e.g. \&quot;en\&quot;, \&quot;de\&quot;, \&quot;es\&quot;). Auto-detected if omitted. Set explicitly for short or mixed-language posts. | [optional] 
**TopicType** | Pointer to **string** | Post type. STANDARD is a regular update. EVENT requires the event object. OFFER requires the offer object. Defaults to STANDARD if omitted. | [optional] [default to "STANDARD"]
**CallToAction** | Pointer to [**GoogleBusinessPlatformDataCallToAction**](GoogleBusinessPlatformDataCallToAction.md) |  | [optional] 
**Event** | Pointer to [**GoogleBusinessPlatformDataEvent**](GoogleBusinessPlatformDataEvent.md) |  | [optional] 
**Offer** | Pointer to [**GoogleBusinessPlatformDataOffer**](GoogleBusinessPlatformDataOffer.md) |  | [optional] 

## Methods

### NewGoogleBusinessPlatformData

`func NewGoogleBusinessPlatformData() *GoogleBusinessPlatformData`

NewGoogleBusinessPlatformData instantiates a new GoogleBusinessPlatformData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGoogleBusinessPlatformDataWithDefaults

`func NewGoogleBusinessPlatformDataWithDefaults() *GoogleBusinessPlatformData`

NewGoogleBusinessPlatformDataWithDefaults instantiates a new GoogleBusinessPlatformData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *GoogleBusinessPlatformData) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *GoogleBusinessPlatformData) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *GoogleBusinessPlatformData) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *GoogleBusinessPlatformData) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetLanguageCode

`func (o *GoogleBusinessPlatformData) GetLanguageCode() string`

GetLanguageCode returns the LanguageCode field if non-nil, zero value otherwise.

### GetLanguageCodeOk

`func (o *GoogleBusinessPlatformData) GetLanguageCodeOk() (*string, bool)`

GetLanguageCodeOk returns a tuple with the LanguageCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguageCode

`func (o *GoogleBusinessPlatformData) SetLanguageCode(v string)`

SetLanguageCode sets LanguageCode field to given value.

### HasLanguageCode

`func (o *GoogleBusinessPlatformData) HasLanguageCode() bool`

HasLanguageCode returns a boolean if a field has been set.

### GetTopicType

`func (o *GoogleBusinessPlatformData) GetTopicType() string`

GetTopicType returns the TopicType field if non-nil, zero value otherwise.

### GetTopicTypeOk

`func (o *GoogleBusinessPlatformData) GetTopicTypeOk() (*string, bool)`

GetTopicTypeOk returns a tuple with the TopicType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicType

`func (o *GoogleBusinessPlatformData) SetTopicType(v string)`

SetTopicType sets TopicType field to given value.

### HasTopicType

`func (o *GoogleBusinessPlatformData) HasTopicType() bool`

HasTopicType returns a boolean if a field has been set.

### GetCallToAction

`func (o *GoogleBusinessPlatformData) GetCallToAction() GoogleBusinessPlatformDataCallToAction`

GetCallToAction returns the CallToAction field if non-nil, zero value otherwise.

### GetCallToActionOk

`func (o *GoogleBusinessPlatformData) GetCallToActionOk() (*GoogleBusinessPlatformDataCallToAction, bool)`

GetCallToActionOk returns a tuple with the CallToAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallToAction

`func (o *GoogleBusinessPlatformData) SetCallToAction(v GoogleBusinessPlatformDataCallToAction)`

SetCallToAction sets CallToAction field to given value.

### HasCallToAction

`func (o *GoogleBusinessPlatformData) HasCallToAction() bool`

HasCallToAction returns a boolean if a field has been set.

### GetEvent

`func (o *GoogleBusinessPlatformData) GetEvent() GoogleBusinessPlatformDataEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *GoogleBusinessPlatformData) GetEventOk() (*GoogleBusinessPlatformDataEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *GoogleBusinessPlatformData) SetEvent(v GoogleBusinessPlatformDataEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *GoogleBusinessPlatformData) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetOffer

`func (o *GoogleBusinessPlatformData) GetOffer() GoogleBusinessPlatformDataOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *GoogleBusinessPlatformData) GetOfferOk() (*GoogleBusinessPlatformDataOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *GoogleBusinessPlatformData) SetOffer(v GoogleBusinessPlatformDataOffer)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *GoogleBusinessPlatformData) HasOffer() bool`

HasOffer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


